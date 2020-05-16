package service

// NewListService ListServiceの初期
func NewListService() ListService {
	return new(listService)
}

type listService struct{}
