package handler

import (
	"net/http"

	"github.com/kakohate/charamell-mvp/service"
)

// NewProfileHandler ProfileHandlerの初期化
func NewProfileHandler(s service.ProfileService) ProfileHandler {
	h := new(profileHandler)
	h.profileService = s
	return h
}

type profileHandler struct {
	profileService service.ProfileService
}

func (h *profileHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
}
