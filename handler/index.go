package handler

import (
	"net/http"

	"github.com/martinmunillas/otter-example/component"
	"github.com/martinmunillas/otter-example/service"
	"github.com/martinmunillas/otter/server"
	"github.com/martinmunillas/otter/server/tools"
)

func GetIndex() server.Handler {
	return func(r *http.Request, t tools.Tools) {
		user, err := service.GetUser()
		if err != nil {
			t.Send.InternalError.HTML(err, component.IndexPage(t, user, err))
			return
		}
		t.Send.Ok.HTML(component.IndexPage(t, user, err))
	}
}
