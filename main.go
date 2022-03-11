package main

import (
	"net/http"
	"time"

	"server/app"
)

func main() {
	app := app.New()

	srv := &http.Server{
		Addr:           app.Cnf.String("server.port"),
		Handler:        app,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	srv.ListenAndServe()

}
