package router

import (
	"net/http"

	"github.com/kakohate/charamell-mvp/handler"
)

type Router interface {
	Mux() *http.ServeMux
	Route()
}

func New(
	mux *http.ServeMux,
	profileHandler handler.ProfileHandler,
	listHandler handler.ListHandler,
) Router {
	r := new(router)
	r.mux = mux
	r.profileHandler = profileHandler
	r.listHandler = listHandler
	return r
}

type router struct {
	mux            *http.ServeMux
	profileHandler handler.ProfileHandler
	listHandler    handler.ListHandler
}

func (r *router) Mux() *http.ServeMux {
	return r.mux
}

func (r *router) Route() {
	r.mux.Handle("/list", r.listHandler)
	r.mux.Handle("/profile", r.profileHandler)
}
