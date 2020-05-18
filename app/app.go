package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kakohate/charamell-mvp/router"
)

// App アドレスとハンドラを返す機能のみ
type App interface {
	Addr() string // ":8080"
	Mux() *http.ServeMux
}

// New Appの初期化
func New(r router.Router) App {
	a := new(app)
	addr := os.Getenv("PORT")
	if addr == "" {
		addr = "8080"
	}
	a.addr = fmt.Sprintf(":%s", addr)
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
