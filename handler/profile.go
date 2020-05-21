package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
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
	m := splitPath(req.URL.Path)
	switch m[1] {
	case "new":
		if req.Method != http.MethodPost {
			httpError(w, http.StatusMethodNotAllowed)
			return
		}
		contentType := req.Header.Get("Content-Type")
		if contentType != "application/json" &&
			contentType != "application/json;charset=UTF-8" {
			log.Println("Content-Type:", contentType)
			httpError(w, http.StatusBadRequest)
			return
		}
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			httpError(w, http.StatusInternalServerError)
			return
		}
		sid, err := h.profileService.CreateProfile(b)
		if err != nil {
			httpError(w, errorToStatusCode(err))
			return
		}
		w.Header().Add("Session-ID", sid.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 OK"))
		return
	default:
		switch req.Method {
		case http.MethodGet:
			uid, err := uuid.Parse(m[1])
			if err != nil {
				http.NotFound(w, req)
				return
			}
			resp, err := h.profileService.GetProfile(uid)
			if err != nil {
				httpError(w, errorToStatusCode(err))
				return
			}
			responseJSON(w, resp)
			return
		case http.MethodDelete:
			sidStr := req.Header.Get("Session-ID")
			if sidStr == "" {
				httpError(w, http.StatusBadRequest)
				return
			}
			sid, err := uuid.Parse(sidStr)
			if err != nil {
				httpError(w, http.StatusBadRequest)
				return
			}
			if err := h.profileService.DeleteProfile(sid); err != nil {
				httpError(w, errorToStatusCode(err))
				return
			}
			c := new(http.Cookie)
			c.Name = "sid"
			c.MaxAge = 0
			http.SetCookie(w, c)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, http.StatusText(http.StatusOK))
			return
		default:
			httpError(w, http.StatusMethodNotAllowed)
			return
		}
	}
}
