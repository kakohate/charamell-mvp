package app

import (
	"net/http"
	"os"

	"github.com/kakohate/charamell-mvp/router"
)

type App interface {
	Addr() string // ":8080"
	Mux() *http.ServeMux
}

func New(r router.Router) App {
	a := new(app)
	addr := os.Getenv("PORT")
	a.addr = addr
	if addr == "" {
		a.addr = ":8080"
	}
	a.router = r
	a.router.Route()
	return a
}

type app struct {
	addr   string
	router router.Router
}

func (a *app) Addr() string {
	return a.addr
}

func (a *app) Mux() *http.ServeMux {
	return a.router.Mux()
}
