package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kakohate/charamell-mvp/router"
	"github.com/rs/cors"
)

// App アドレスとハンドラを返す機能のみ
type App interface {
	Addr() string // ":8080"
	Mux() http.Handler
}

// New Appの初期化
func New(r router.Router) App {
	a := new(app)
	addr := os.Getenv("PORT")
	if addr == "" {
		addr = "8080"
	}
	a.addr = fmt.Sprintf(":%s", addr)
	r.Route()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
	})
	a.mux = c.Handler(r.Mux())
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
