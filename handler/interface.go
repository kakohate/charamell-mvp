package handler

import "net/http"

type ProfileHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type ListHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
