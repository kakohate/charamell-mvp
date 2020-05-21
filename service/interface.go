package service

import (
	"github.com/google/uuid"
)

// ProfileService プロフィールの作成, 取得, 削除, 期限の取得
type ProfileService interface {
	CreateProfile([]byte) (sid *uuid.UUID, err error)
	GetProfile(uid uuid.UUID) (resp []byte, err error)
	DeleteProfile(sid uuid.UUID) error
	GetProfileExpires(sid uuid.UUID) (resp []byte, err error)
}

// ListService リストの取得
type ListService interface {
	GetList(sid uuid.UUID) (resp []byte, err error)
}
