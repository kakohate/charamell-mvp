package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kakohate/charamell-mvp/model"
	"github.com/kakohate/charamell-mvp/repository"
)

// NewProfileService ProfileServiceの初期化
func NewProfileService(
	profileRepository repository.ProfileRepository,
) ProfileService {
	return &profileService{
		profileRepository: profileRepository,
	}
}

type profileService struct {
	profileRepository repository.ProfileRepository
}

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

func (s *profileService) CreateProfile(b []byte) (*uuid.UUID, error) {
	req := new(createProfileRequest)
	json.Unmarshal(b, req)
	id, _ := uuid.NewUUID()
	sid, _ := uuid.NewUUID()
	now := time.Now()
	profile := &model.Profile{
		ID:        id,
		SID:       sid,
		CreatedAt: &now,
		Deleted:   false,
		Name:      req.Name,
		Message:   req.Message,
		Limit:     req.Limit,
		Color:     req.Color,
		AvatarURL: req.AvatarURL,
		Coordinate: &model.Coordinate{
			Lat: req.Coordinate.Lat,
			Lng: req.Coordinate.Lng,
		},
	}
	for _, tag := range req.Tag {
		id, _ := uuid.NewUUID()
		profile.Tag = append(profile.Tag, &model.Tag{
			ID:        id,
			ProfileID: profile.ID,
			Category:  tag.Category,
			Detail:    tag.Detail,
		})
	}
	for _, picture := range req.Pictures {
		id, _ := uuid.NewUUID()
		profile.Pictures = append(profile.Pictures, &model.Picture{
			ID:        id,
			ProfileID: profile.ID,
			Order:     picture.Order,
			URL:       picture.URL,
		})
	}
	if err := s.profileRepository.Create(profile); err != nil {
		return nil, status(http.StatusBadRequest)
	}
	return &profile.SID, nil
}

func (s *profileService) GetProfile(uid uuid.UUID) ([]byte, error) {
	return nil, nil
}

func (s *profileService) DeleteProfile(sid uuid.UUID) error {
	return nil
}
