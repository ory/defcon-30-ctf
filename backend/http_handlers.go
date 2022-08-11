package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
)

type handler struct {
	repo repository
	jw   *herodot.JSONWriter
	tw   *herodot.TextWriter
}

type logReporter struct{}

func (logReporter) ReportError(r *http.Request, code int, err error, args ...interface{}) {
	log.Printf("ERROR: %s\n  Request: %v\n  Response Code: %d\n  Further Info: %v\n", err, r, code, args)
}

func NewHandler(repo repository) http.Handler {
	r := httprouter.New()
	h := &handler{
		repo: repo,
		jw:   herodot.NewJSONWriter(logReporter{}),
		tw:   herodot.NewTextWriter(logReporter{}, "html"),
	}

	r.GET("/results", h.getResults)
	r.POST("/results/:district", h.submit)
	r.GET("/login", requireFlow(h.login, "login"))
	r.GET("/register", requireFlow(h.register, "registration"))
	r.GET("/", h.index)

	return withAccessLog(r)
}

func withAccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func requireFlow(next httprouter.Handle, flow string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if r.URL.Query().Has("flow") {
			next(w, r, nil)
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/self-service/%s/browser", flow), http.StatusSeeOther)
	}
}

func (h *handler) getResults(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	results, err := h.repo.List(r.Context())
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	b, _ := json.Marshal(results)
	_, _ = w.Write(b)
}

func (h *handler) submit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r.Body == nil {
		h.jw.WriteError(w, r, herodot.ErrBadRequest.WithError("no body"))
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.jw.WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
		return
	}

	res := &result{}
	if err = json.Unmarshal(body, res); err != nil {
		h.jw.WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
		return
	}

	err = h.repo.Submit(r.Context(), params.ByName("district"), res)
	if err != nil {
		h.jw.WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
		return
	}
	h.jw.WriteCreated(w, r, "/results/"+res.District, "")
}

//go:embed ui/login.html
var login []byte

//go:embed ui/register.html
var register []byte

//go:embed ui/index.html
var index []byte

func (h *handler) login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.tw.Write(w, r, login)
}

func (h *handler) register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.tw.Write(w, r, register)
}

func (h *handler) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.tw.Write(w, r, index)
}

type kratosFlow struct {
	ID string `json:"id"`
}

func getFlow(r *http.Request, flow string) (*string, error) {
	flowID := r.URL.Query().Get("flow")
	flowReq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://kratos:4433/self-service/%s/flows/%s", flow, flowID), nil)
	if err != nil {
		return nil, err
	}
	flowReq.Header.Set("Cookie", r.Header.Get("Cookie"))
	flowReq.Header.Set("Accept", "application/json")
	flowRes, err := http.DefaultClient.Do(flowReq)
	if err != nil {
		return nil, err
	}
	defer flowRes.Body.Close()
	if flowRes.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not fetch flow: %s", flowRes.Status)
	}

}
