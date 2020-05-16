package app

import (
	"net/http"
	"os"

	"github.com/kakohate/charamell-mvp/router"
)

type App interface {
	Addr() string // ":8080"
	Mux() http.Handler
}

func New(r router.Router) App {
	a := new(app)
	addr := os.Getenv("PORT")
	a.addr = addr
	if addr == "" {
		a.addr = ":8080"
	}
	a.mux = r
	return a
}

type app struct {
	addr string
	mux  http.Handler
}

func (a *app) Addr() string {
	return a.addr
}

func (a *app) Mux() http.Handler {
	return a.mux
}
