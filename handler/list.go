package handler

import (
	"net/http"

	"github.com/google/uuid"
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
	m := splitPath(req.URL.Path)
	switch m[1] {
	case "":
		if req.Method != http.MethodGet {
			httpError(w, http.StatusMethodNotAllowed)
		}
		c, err := req.Cookie("sid")
		if err != nil {
			httpError(w, http.StatusBadRequest)
			return
		}
		sid, err := uuid.Parse(c.Value)
		if err != nil {
			httpError(w, http.StatusBadRequest)
			return
		}
		resp, err := h.listService.GetList(sid)
		if err != nil {
			http.Error(w, string(resp), errorToStatusCode(err))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
		return
	}
}
