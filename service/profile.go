package service

import (
	"database/sql"
	"encoding/json"
	"log"
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

type request struct {
	Profile struct {
		Name       string           `json:"name"`
		Tag        []profileTag     `json:"tag"`
		Message    string           `json:"message"`
		Limit      uint             `json:"limit"`
		Color      string           `json:"color"`
		AvatarURL  string           `json:"avatar_url"`
		Pictures   []profilePicture `json:"pictures"`
		Coordinate profileCoodinate `json:"coordinate"`
	} `json:"profile"`
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
	req := new(request)
	if err := json.Unmarshal(b, req); err != nil {
		log.Println("service", 1, err)
		return nil, status(http.StatusBadRequest)
	}
	id := uuid.New()
	sid := uuid.New()
	now := time.Now()
	expires := now.Add(time.Duration(req.Profile.Limit) * time.Hour)
	profile := &model.Profile{
		ID:        id,
		SID:       sid,
		CreatedAt: &now,
		Expires:   &expires,
		Deleted:   false,
		Name:      req.Profile.Name,
		Message:   req.Profile.Message,
		Limit:     req.Profile.Limit,
		Color:     req.Profile.Color,
		AvatarURL: req.Profile.AvatarURL,
		Coordinate: &model.Coordinate{
			ID:        uuid.New(),
			ProfileID: id,
			Lat:       req.Profile.Coordinate.Lat,
			Lng:       req.Profile.Coordinate.Lng,
		},
	}
	for _, tag := range req.Profile.Tag {
		profile.Tag = append(profile.Tag, &model.Tag{
			ID:        uuid.New(),
			ProfileID: profile.ID,
			Category:  tag.Category,
			Detail:    tag.Detail,
		})
	}
	for _, picture := range req.Profile.Pictures {
		profile.Pictures = append(profile.Pictures, &model.Picture{
			ID:        uuid.New(),
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

type profileResp struct {
	Name      string     `json:"name"`
	Message   string     `json:"message"`
	TagResp   []*tagResp `json:"tag"`
	Limit     int        `json:"limit"`
	Color     string     `json:"color"`
	AvatarURL string     `json:"avatar_url"`
}

type tagResp struct {
	Category string `json:"category"`
	Detail   string `json:"detail"`
}

func (s *profileService) GetProfile(uid uuid.UUID) ([]byte, error) {
	profile, err := s.profileRepository.GetOne(uid)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, status(http.StatusNotFound)
		}
		return nil, status(http.StatusInternalServerError)
	}
	resp := &profileResp{
		Name:      profile.Name,
		Message:   profile.Message,
		Limit:     int(time.Until(*profile.Expires).Seconds()),
		Color:     profile.Color,
		AvatarURL: profile.AvatarURL,
	}
	for _, tag := range profile.Tag {
		tr := &tagResp{
			Category: tag.Category,
			Detail:   tag.Detail,
		}
		resp.TagResp = append(resp.TagResp, tr)
	}
	return json.Marshal(resp)
}

func (s *profileService) DeleteProfile(sid uuid.UUID) error {
	return nil
}
