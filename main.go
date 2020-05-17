package main

import (
	"log"
	"net/http"

	"github.com/kakohate/charamell-mvp/database"
	"github.com/kakohate/charamell-mvp/env"
	"github.com/kakohate/charamell-mvp/wire"
)

func init() {
	db, err := database.New(env.NewEnv())
	defer func() {
		db.Close()
	}()
	if err != nil {
		panic(err)
	}
	if err := database.Init(db); err != nil {
		panic(err)
	}
}

func main() {
	app, err := wire.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(app.Addr(), app.Mux()); err != nil {
		log.Fatal(err)
	}
}
