package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kakohate/charamell-mvp/repository"
)

// NewListService ListServiceの初期
func NewListService() ListService {
	return new(listService)
}

type listService struct {
	profileRepository repository.ProfileRepository
}

type listResponse struct {
	List []*listProfile `json:"list"`
}

type listProfile struct {
	ID        uuid.UUID         `json:"id"`
	Color     string            `json:"color"`
	AvatarURL string            `json:"avatar_url"`
	Limit     uint              `json:"limit"`
	Tag       []*listProfileTag `json:"tag"`
}

type listProfileTag struct {
	Category string `json:"category"`
	IsMatch  bool   `json:"is_match"`
}

func (s *listService) GetList(sid uuid.UUID) ([]byte, error) {
	profile, err := s.profileRepository.GetOneBySID(sid)
	if err != nil {
		return nil, err
	}
	if profile.Deleted || profile.Expires.After(time.Now()) {
		return []byte("Session expired"), status(http.StatusBadRequest)
	}
	profiles, err := s.profileRepository.GetList(sid)
	if err != nil {
		return nil, err
	}
	list := new(listResponse)
	for _, profile := range profiles {
		limit := time.Until(*profile.Expires).Hours()
		lp := &listProfile{
			ID:        profile.ID,
			Color:     profile.Color,
			AvatarURL: profile.AvatarURL,
			Limit:     uint(limit),
		}
		for _, tag := range profile.Tag {
			lp.Tag = append(lp.Tag, &listProfileTag{
				Category: tag.Category,
				IsMatch:  tag.IsMatch,
			})
		}
		list.List = append(list.List, lp)
	}
	return json.Marshal(list)
}
