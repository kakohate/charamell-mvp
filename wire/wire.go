//+build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/kakohate/charamell-mvp/app"
	"github.com/kakohate/charamell-mvp/handler"
	"github.com/kakohate/charamell-mvp/router"
)

func NewApp() app.App {
	wire.Build(
		app.New,
		router.New,
		handler.NewProfileHandler,
		handler.NewListHandler,
	)
	return nil
}
