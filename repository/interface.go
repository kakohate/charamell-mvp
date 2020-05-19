package repository

import (
	"github.com/google/uuid"
	"github.com/kakohate/charamell-mvp/model"
)

// ProfileRepository プロフィールの作成, 取得, 削除
type ProfileRepository interface {
	Create(*model.Profile) error
	GetOne(uuid.UUID) (*model.Profile, error)
	GetOneBySID(uuid.UUID) (*model.Profile, error)
	GetList(uuid.UUID) ([]*model.Profile, error)
	Delete(uuid.UUID) error
}
