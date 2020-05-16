package handler

import "net/http"

type profileHandler struct{}

func (*profileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func NewProfileHandler() ProfileHandler {
	return new(profileHandler)
}
