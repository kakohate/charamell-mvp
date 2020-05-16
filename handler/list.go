package handler

import (
	"net/http"

	"github.com/kakohate/charamell-mvp/service"
)

// NewListHandler ListHandlerの初期化
func NewListHandler(s service.ListService) ListHandler {
	h := new(listHandler)
	h.listService = s
	return h
}

type listHandler struct {
	listService service.ListService
}

func (h *listHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
}
