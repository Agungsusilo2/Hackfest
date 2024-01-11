package middleware

import (
	"encoding/json"
	"github.com/Agungsusilo2/Hackfest/model/helper"
	"github.com/Agungsusilo2/Hackfest/model/web"
	"net/http"
)

type MiddlewareAuth struct {
	Handler http.Handler
}

func (a MiddlewareAuth) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "Pukulele" == request.Header.Get("X-API-KEY") {
		a.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		res := web.ResponseTemp{Code: http.StatusUnauthorized, Status: "Unauthorized"}

		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(res)
		helper.PanicErr(err)
	}
}

func NewMiddlewareAuth(handler http.Handler) *MiddlewareAuth {
	return &MiddlewareAuth{Handler: handler}
}
