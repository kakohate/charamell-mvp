package repository

// ProfileRepository プロフィールの作成, 取得, 削除
type ProfileRepository interface {
	Create()
	GetOne()
	GetList()
	Delete()
}
