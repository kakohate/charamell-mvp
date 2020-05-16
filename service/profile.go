package service

import (
	"github.com/google/uuid"
)

// NewProfileService ProfileServiceの初期化
func NewProfileService() ProfileService {
	return new(profileService)
}

type profileService struct{}

func (s *profileService) CreateProfile(b []byte) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (s *profileService) GetProfile(uid uuid.UUID) ([]byte, error) {
	return nil, nil
}

func (s *profileService) DeleteProfile(sid uuid.UUID) error {
	return nil
}
