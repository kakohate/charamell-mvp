package handler

import "net/http"

type listHandler struct{}

func (*listHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func NewListHandler() ListHandler {
	return new(listHandler)
}
