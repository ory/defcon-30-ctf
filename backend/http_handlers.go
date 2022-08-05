package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
)

type handler struct {
	repo repository
	w    *herodot.JSONWriter
}

type logReporter struct{}

func (logReporter) ReportError(r *http.Request, code int, err error, args ...interface{}) {
	log.Printf("ERROR: %s\n  Request: %v\n  Response Code: %d\n  Further Info: %v\n", err, r, code, args)
}

func newHandler(repo repository) http.Handler {
	r := httprouter.New()
	h := &handler{
		repo: repo,
		w:    herodot.NewJSONWriter(logReporter{}),
	}

	r.GET("/results", withAccessLog(h.getResults))
	r.POST("/results", withAccessLog(h.submit))

	return r
}

func withAccessLog(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		log.Printf("%s %s\n", r.Method, r.URL)
		next(w, r, p)
	}
}

func (h *handler) getResults(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	results, err := h.repo.List(r.Context())
	if err != nil {
		h.w.WriteError(w, r, err)
		return
	}
	h.w.Write(w, r, results)
}

func (h *handler) submit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.w.WriteError(w, r, err)
		return
	}

	result := &result{}
	if err = json.Unmarshal(body, result); err != nil {
		h.w.WriteError(w, r, err)
		return
	}

	err = h.repo.Submit(r.Context(), result)
	if err != nil {
		h.w.WriteError(w, r, err)
		return
	}
	h.w.WriteCreated(w, r, "/results", "")
}
