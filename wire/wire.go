//+build wireinject

package wire

import (
	"net/http"

	"github.com/google/wire"
	"github.com/kakohate/charamell-mvp/app"
	"github.com/kakohate/charamell-mvp/handler"
	"github.com/kakohate/charamell-mvp/router"
	"github.com/kakohate/charamell-mvp/service"
)

func NewApp() app.App {
	wire.Build(
		app.New,
		router.New,
		http.NewServeMux,
		handler.NewProfileHandler,
		handler.NewListHandler,
		service.NewProfileService,
		service.NewListService,
	)
	return nil
}
