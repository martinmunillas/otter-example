package handler

import (
	"net/http"

	"github.com/martinmunillas/otter-example/component"
	"github.com/martinmunillas/otter-example/service"
	"github.com/martinmunillas/otter/response/send"
)

func GetIndex() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := service.GetUser()
		if err != nil {
			send.Html.Ok(w, r, component.IndexPage(user, err))
			return
		}
		send.Html.Ok(w, r, component.IndexPage(user, err))
	})
}
