package handler

import (
	"log/slog"
	"net/http"

	"github.com/martinmunillas/otter-example/component"
	"github.com/martinmunillas/otter-example/service"
	"github.com/martinmunillas/otter/response/send"
)

func GetIndex() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := service.GetUser()
		if err != nil {
			err = send.Html.Ok(w, r, component.IndexPage(user, err))
			if err != nil {
				slog.Error(err.Error())
			}
			return
		}
		err = send.Html.Ok(w, r, component.IndexPage(user, err))
		if err != nil {
			slog.Error(err.Error())
		}
	})
}
