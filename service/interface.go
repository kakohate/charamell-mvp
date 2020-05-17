package service

import (
	"net/http"

	"github.com/google/uuid"
)

// ProfileService プロフィールの作成, 取得, 削除
type ProfileService interface {
	CreateProfile([]byte) (sid *uuid.UUID, err error)
	GetProfile(uid uuid.UUID) (resp []byte, err error)
	DeleteProfile(sid uuid.UUID) error
}

// ListService リストの取得
type ListService interface {
	GetList(*http.Request) error
}
