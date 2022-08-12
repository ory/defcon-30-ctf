package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/client-go"
	"github.com/ory/herodot"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type handler struct {
	repo      *sqlRepo
	config    *Config
	jw        *herodot.JSONWriter
	tw        *herodot.TextWriter
	c         *client.APIClient
	ketoCheck rts.CheckServiceClient
	ketoWrite rts.WriteServiceClient
}

type logReporter struct{}

func (logReporter) ReportError(r *http.Request, code int, err error, args ...interface{}) {
	log.Printf("ERROR: %s\n  Request: %v\n  Response Code: %d\n  Further Info: %v\n", err, r, code, args)
}

func NewHandler(repo *sqlRepo, config *Config) (http.Handler, error) {
	// TODO cluster-internal mTLS
	ketoRead, err := grpc.Dial("keto:4466", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	// TODO cluster-internal mTLS
	ketoWrite, err := grpc.Dial("keto:4467", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	r := httprouter.New()
	h := &handler{
		repo:   repo,
		config: config,
		jw:     herodot.NewJSONWriter(logReporter{}),
		tw:     herodot.NewTextWriter(logReporter{}, "html"),
		c: client.NewAPIClient(&client.Configuration{
			Servers: client.ServerConfigurations{{
				// TODO cluster-internal mTLS
				URL: "http://kratos:4433",
			}},
		}),
		ketoCheck: rts.NewCheckServiceClient(ketoRead),
		ketoWrite: rts.NewWriteServiceClient(ketoWrite),
	}

	r.GET("/login", requireFlow(h.login, "login"))
	r.GET("/register", requireFlow(h.register, "registration"))
	r.GET("/error", h.error)
	r.GET("/", h.index)
	r.POST("/results", h.submit)
	r.POST("/grant-access", h.grantAccess)
	r.GET("/flag", h.getFlag)
	r.POST("/flag", h.submitFlag)
	r.GET("/static/*filepath", h.static)
	r.GET("/leaderboard", h.leaderboard)

	return withAccessLog(r), nil
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

//go:embed static/*
var staticFiles embed.FS

func (h *handler) static(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	statics, err := fs.Sub(staticFiles, "static")
	if err != nil {
		h.jw.WriteError(w, r, err)
		return
	}
	http.StripPrefix("/static/", http.FileServer(http.FS(statics))).ServeHTTP(w, r)
}
