package main

import (
	"net/http"

	"github.com/martinmunillas/otter/env"
	"github.com/martinmunillas/otter/server"

	"github.com/martinmunillas/otter-example/handler"
	"github.com/martinmunillas/otter-example/translations"
)

var port = env.OptionalIntEnvVar("PORT", 8080)

func main() {
	translations.Setup()

	server.NewServer([]server.Endpoint{{
		Path:    "/{$}",
		Method:  http.MethodGet,
		Handler: handler.GetIndex(),
	}}).ServeStatic("./static").Listen((port))

}
