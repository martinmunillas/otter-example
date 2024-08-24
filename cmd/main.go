package main

import (
	"log/slog"
	"net/http"

	"github.com/martinmunillas/otter/api"
	"github.com/martinmunillas/otter/env"
	"github.com/martinmunillas/otter/i18n"

	"github.com/martinmunillas/otter-example/handler"
	"github.com/martinmunillas/otter-example/translations"
)

var port = env.OptionalIntEnvVar("PORT", 8080)

func main() {

	translations.Setup()

	mux := http.NewServeMux()

	mux.Handle("GET /{$}", handler.GetIndex())

	api.ServeStatic("./static", mux)

	err := http.ListenAndServe(api.PortString(port), i18n.Middleware(mux))
	if err != nil {
		slog.Error(err.Error())
	}
}
