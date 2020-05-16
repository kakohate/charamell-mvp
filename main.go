package main

import (
	"log"
	"net/http"

	"github.com/kakohate/charamell-mvp/wire"
)

func main() {
	app := wire.NewApp()

	if err := http.ListenAndServe(app.Addr(), app.Mux()); err != nil {
		log.Fatal(err)
	}
}
