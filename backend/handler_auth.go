package main

import (
	_ "embed"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/ory/client-go"
	"github.com/ory/herodot"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
	"html/template"
	"net/http"
	"os"
)

//go:embed ui/login.html
var loginPage string

//go:embed ui/register.html
var registerPage string

//go:embed ui/error.html
var errorPage string

func (h *handler) login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	flow, err := getFlow(r, h.c.V0alpha2Api.GetSelfServiceLoginFlowExecute)
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	t, err := template.New("login").Parse(loginPage)
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
	t, err := template.New("register").Parse(registerPage)
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	if err := t.Execute(w, flow); err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
}

func (h *handler) error(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	e, _, err := h.c.V0alpha2Api.GetSelfServiceErrorExecute(client.V0alpha2ApiGetSelfServiceErrorRequest{}.Id(r.URL.Query().Get("id")))
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	t, err := template.New("error").Parse(errorPage)
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	if err := t.Execute(w, struct{ Error *client.SelfServiceError }{Error: e}); err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
}

type flowData struct {
	Action, CSRFToken, Messages, WebAuthNScript, Identifier string
	WebAuthNCallback                                        template.JS
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
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(flow.GetUi())
	data := flowData{
		Action: flow.GetUi().Action,
	}
	msg := flow.GetUi().Messages
	for _, node := range flow.GetUi().Nodes {
		if attrs := node.Attributes.UiNodeInputAttributes; attrs != nil {
			if attrs.Name == "csrf_token" {
				data.CSRFToken = attrs.Value.(string)
			} else if attrs.Name == "webauthn_register_trigger" || attrs.Name == "webauthn_login_trigger" {
				data.WebAuthNCallback = template.JS(*attrs.Onclick)
			} else if attrs.Name == "identifier" {
				data.Identifier = attrs.Value.(string)
			}
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

func (h *handler) grantAccess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sub, rel, nspace, obj := r.FormValue("subject"), r.FormValue("relation"), r.FormValue("namespace"), r.FormValue("object")

	self, err := h.ketoCheck.Check(r.Context(), &rts.CheckRequest{
		Tuple: &rts.RelationTuple{
			Namespace: nspace,
			Object:    obj,
			Relation:  rel,
			Subject:   rts.NewSubjectID(r.Header.Get("X-Username")),
		},
	})
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	if !self.Allowed {
		h.jw.WriteError(w, r, herodot.ErrForbidden.WithError("You don't have the relation yourself, how could you grant it to someone else?"))
		return
	}

	_, err = h.ketoWrite.TransactRelationTuples(r.Context(), &rts.TransactRelationTuplesRequest{
		RelationTupleDeltas: []*rts.RelationTupleDelta{{
			Action: rts.RelationTupleDelta_ACTION_INSERT,
			RelationTuple: &rts.RelationTuple{
				Namespace: nspace,
				Object:    obj,
				Relation:  rel,
				Subject:   rts.NewSubjectID(sub),
			},
		}},
	})
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	h.tw.Write(w, r, "Ok, done.")
}
