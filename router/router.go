package router

import (
	"net/http"

	"github.com/kakohate/charamell-mvp/handler"
)

type Router interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	Route(http.ResponseWriter, *http.Request)
}

func New(
	profileHandler handler.ProfileHandler,
	listHandler handler.ListHandler,
) Router {
	r := new(router)
	r.mux = http.NewServeMux()
	r.profileHandler = profileHandler
	r.listHandler = listHandler
	return r
}

type router struct {
	mux            *http.ServeMux
	profileHandler handler.ProfileHandler
	listHandler    handler.ListHandler
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Route(w, req)
}

func (r *router) Route(w http.ResponseWriter, req *http.Request) {

}
