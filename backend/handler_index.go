package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
	"html/template"
	"net/http"
	"strconv"
)

//go:embed ui/index.html
var indexPage string

type indexData struct {
	IdentityID, Username, Results, Email string
}

func (h *handler) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, err := h.repo.List(r.Context())
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	if err := enc.Encode(res); err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	t, err := template.New("index").Parse(indexPage)
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	if err := t.Execute(w, &indexData{
		IdentityID: r.Header.Get("X-User"),
		Username:   r.Header.Get("X-Username"),
		Email:      r.Header.Get("X-Useremail"),
		Results:    buf.String(),
	}); err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
}

func (h *handler) submit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	getCount := func(name string) uint {
		count, err := strconv.ParseUint(r.FormValue(name), 10, 0)
		if err != nil {
			return 0
		}
		return uint(count)
	}
	res := &result{
		District:    r.FormValue("district"),
		Democrats:   getCount("democrats"),
		Republicans: getCount("republicans"),
		Invalid:     getCount("invalid"),
	}

	resp, err := h.ketoCheck.Check(r.Context(), &rts.CheckRequest{
		Tuple: &rts.RelationTuple{
			Namespace: "districts",
			Object:    res.District,
			Relation:  "submit",
			Subject:   rts.NewSubjectID(r.Header.Get("X-User")),
		},
	})
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	if !resp.Allowed {
		h.jw.WriteError(w, r, herodot.ErrForbidden.WithError("not allowed"))
		return
	}

	if err := h.repo.Submit(r.Context(), res.District, res); err != nil {
		h.jw.WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
		return
	}
	h.jw.WriteCreated(w, r, "/results", "")
}
