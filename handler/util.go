package handler

import (
	"net/http"
	"strconv"
	"strings"
)

// ここらへんはもっとうまいことやりたい
func splitPath(path string) map[int]string {
	strs := strings.Split(path, "/")
	strs = append(strs[:0], strs[1:]...) // 先頭が空白なので削除
	m := make(map[int]string)
	for i, s := range strs {
		m[i] = s
	}
	return m
}

func errorToStatusCode(err error) int {
	code, err := strconv.Atoi(err.Error())
	if err != nil {
		return http.StatusInternalServerError
	}
	return code
}

func httpError(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func responseJSON(w http.ResponseWriter, resp []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
