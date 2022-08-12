package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	_ "embed"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

func encodeFlag(now, user, seed string) string {
	buf := bytes.Buffer{}
	_, _ = buf.WriteString("flag_")
	enc := base64.NewEncoder(base64.URLEncoding, &buf)
	sum := hmac.New(sha512.New, []byte(seed)).Sum([]byte(now + user))
	fmt.Printf("sum enc: %v\n", sum)
	_, _ = enc.Write([]byte(now + "_"))
	_, _ = enc.Write([]byte(user + "_"))
	_, _ = enc.Write(sum)
	_ = enc.Close()
	return buf.String()
}

func decodeFlag(flag, seed string) (when, user string, err error) {
	dec, err := io.ReadAll(base64.NewDecoder(base64.URLEncoding, bytes.NewBufferString(strings.TrimPrefix(flag, "flag_"))))
	if err != nil {
		return "", "", err
	}
	parts := bytes.Split(dec, []byte("_"))
	if len(parts) != 3 {
		return "", "", herodot.ErrBadRequest.WithError("Invalid flag format")
	}
	whenB, userB, sum := parts[0], parts[1], parts[2]
	sumDec := hmac.New(sha512.New, []byte(seed)).Sum(append(whenB, userB...))
	fmt.Printf("sum dec: %v\nact sum: %v\n", sumDec, sum)
	if !hmac.Equal(sumDec, sum) {
		return "", "", herodot.ErrBadRequest.WithError("Invalid flag")
	}
	return string(whenB), string(userB), nil
}

func (h *handler) getFlag(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := r.Header.Get("X-User")
	res, err := h.ketoCheck.Check(r.Context(), &rts.CheckRequest{
		Tuple: &rts.RelationTuple{
			Namespace: "flags",
			Object:    "new_flag",
			Relation:  "create",
			Subject:   rts.NewSubjectID(user),
		},
	})
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	if !res.Allowed {
		h.jw.WriteError(w, r, herodot.ErrForbidden.WithError("Sorry, no flag for you..."))
		return
	}

	flag := encodeFlag(strconv.Itoa(int(time.Now().Unix())), user, h.config.FlagSeed)
	h.tw.Write(w, r, flag+"\n\nCongrats! Go and submit your flag. Please keep a copy in case the data get lost because of unforeseen events.\n")
}

func (h *handler) submitFlag(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	flag := r.FormValue("flag")
	when, user, err := decodeFlag(flag, h.config.FlagSeed)
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	if user != r.Header.Get("X-User") {
		h.jw.WriteError(w, r, herodot.ErrBadRequest.WithError("This flag was issued for another user"))
		return
	}
	if err := h.repo.FlagSubmission(r.Context(), when, user, flag, r.FormValue("email")); err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	h.jw.Write(w, r, "Thanks, we will reach out to you about your swag!")
}

//go:embed ui/leaderboard.html
var leaderboardPage string

type leaderboardData struct {
	Submissions []*flagSubmission
	Total       int
}

func (h *handler) leaderboard(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	submissions, err := h.repo.ListFlags(r.Context())
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	total, err := h.repo.TotalFlags(r.Context())
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	t, err := template.New("leaderboard").Parse(leaderboardPage)
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	if err := t.Execute(w, &leaderboardData{
		Submissions: submissions,
		Total:       total,
	}); err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
}
