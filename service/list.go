package service

import "net/http"

// NewListService ListServiceの初期
func NewListService() ListService {
	return new(listService)
}

type listService struct{}

func (s *listService) GetList(req *http.Request) error { return nil }
