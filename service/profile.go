package service

import (
	"encoding/json"

	"github.com/google/uuid"
)

// NewProfileService ProfileServiceの初期化
func NewProfileService() ProfileService {
	return new(profileService)
}

type profileService struct{}

type createProfileRequest struct {
	Name       string           `json:"name"`
	Tag        []profileTag     `json:"tag"`
	Message    string           `json:"message"`
	Limit      uint             `json:"limit"`
	Color      string           `json:"color"`
	AvatarURL  string           `json:"avatar_url"`
	Pictures   []profilePicture `json:"pictures"`
	Coordinate profileCoodinate `json:"coordinate"`
}

type profileTag struct {
	Category string `json:"category"`
	Detail   string `json:"detail"`
}

type profilePicture struct {
	Order uint   `json:"order"`
	URL   string `json:"url"`
}
type profileCoodinate struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

func (s *profileService) CreateProfile(b []byte) (uuid.UUID, error) {
	req := new(createProfileRequest)
	json.Unmarshal(b, req)
	return uuid.UUID{}, nil
}

func (s *profileService) GetProfile(uid uuid.UUID) ([]byte, error) {
	return nil, nil
}

func (s *profileService) DeleteProfile(sid uuid.UUID) error {
	return nil
}
