package main

import (
	"net/http"
	"os"
	"time"

	"github.com/apex/httplog"
	"github.com/apex/log"
	"github.com/riazXrazor/github-commit-chart/router"
)

// ServerInit start the server
func ServerInit() {
	PORT := os.Getenv("PORT")
	var ctx = log.WithField("port", PORT)

	r := router.Init()

	var srv = &http.Server{
		Handler:      httplog.New(r),
		Addr:         ":" + PORT,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	ctx.Info("starting up at " + ":" + PORT + "...")
	ctx.WithError(srv.ListenAndServe()).Error("failed to start up server")
}
