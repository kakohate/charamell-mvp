package model

import (
	"time"

	"github.com/google/uuid"
)

// Profile プロフィール
type Profile struct {
	ID         uuid.UUID
	SID        uuid.UUID
	CreatedAt  *time.Time
	Expires    *time.Time
	Deleted    bool
	Name       string
	Message    string
	Limit      uint // time_limit
	Color      string
	AvatarURL  string
	Tag        []*Tag
	Pictures   []*Picture
	Coordinate *Coordinate
}

// Tag タグ
type Tag struct {
	ID        uuid.UUID
	ProfileID uuid.UUID
	Category  string
	Detail    string
	IsMatch   bool
}

// Picture プロフィール画像
type Picture struct {
	ID        uuid.UUID
	ProfileID uuid.UUID
	Order     uint // display_order
	URL       string
}

// Coordinate 座標
type Coordinate struct {
	ID        uuid.UUID
	ProfileID uuid.UUID
	Lat       float64
	Lng       float64
}
