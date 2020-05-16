package service

// NewProfileService ProfileServiceの初期化
func NewProfileService() ProfileService {
	return new(profileService)
}

type profileService struct{}
