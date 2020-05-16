package handler

import "net/http"

// ProfileHandler プロフィールに関する操作をハンドル
type ProfileHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// ListHandler リストに関する操作をハンドル
type ListHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
