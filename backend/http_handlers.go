package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/ory/client-go"
	"github.com/ory/herodot"
	"html/template"
	"io"
	"log"
	"net/http"
)

type handler struct {
	repo repository
	jw   *herodot.JSONWriter
	tw   *herodot.TextWriter
	c    *client.APIClient
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
		c: client.NewAPIClient(&client.Configuration{
			Servers: client.ServerConfigurations{{
				URL: "http://kratos:4433",
			}},
		}),
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
	flow, err := getFlow(r, h.c.V0alpha2Api.GetSelfServiceLoginFlowExecute)
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	t, err := template.New("login").Parse(string(login))
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	if err := t.Execute(w, flow); err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
}

func (h *handler) register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	flow, err := getFlow(r, h.c.V0alpha2Api.GetSelfServiceRegistrationFlowExecute)
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	t, err := template.New("register").Parse(string(register))
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	if err := t.Execute(w, flow); err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
}

func (h *handler) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.tw.Write(w, r, index)
}

type flowData struct {
	Action    string
	CSRFToken string
	Messages  string
}

func getFlow[F interface {
	GetUi() client.UiContainer
}, R interface {
	Cookie(string) R
	Id(string) R
}](r *http.Request, fetch func(R) (F, *http.Response, error)) (*flowData, error) {
	flowID := r.URL.Query().Get("flow")
	var req R
	flow, _, err := fetch(req.
		Cookie(r.Header.Get("Cookie")).
		Id(flowID))
	if err != nil {
		return nil, err
	}
	data := flowData{
		Action: flow.GetUi().Action,
	}
	msg := flow.GetUi().Messages
	for _, node := range flow.GetUi().Nodes {
		if attrs := node.Attributes.UiNodeInputAttributes; attrs != nil && attrs.Name == "csrf_token" {
			data.CSRFToken = attrs.Value.(string)
		}
		msg = append(msg, node.Messages...)
	}
	msgRaw, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	data.Messages = string(msgRaw)
	if len(msg) == 0 {
		data.Messages = "no messages right now"
	}
	return &data, nil
}
