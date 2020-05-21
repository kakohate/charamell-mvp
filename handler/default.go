package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/kakohate/charamell-mvp/service"
)

// NewDefaultHandler DefaultHandlerの初期化
func NewDefaultHandler(
	profileService service.ProfileService,
) DefaultHandler {
	return &defaultHandler{
		profileService: profileService,
	}
}

type defaultHandler struct {
	profileService service.ProfileService
}

func (h *defaultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	switch req.Method {
	case http.MethodGet:
		sidStr := req.Header.Get("Session-ID")
		sid, err := uuid.Parse(sidStr)
		if err != nil {
			httpError(w, http.StatusBadRequest)
		}
		resp, err := h.profileService.GetProfileExpires(sid)
		if err != nil {
			httpError(w, errorToStatusCode(err))
		}
		responseJSON(w, resp)
		return
	default:
		http.NotFound(w, req)
		return
	}
}
