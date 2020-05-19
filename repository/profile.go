package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/kakohate/charamell-mvp/model"
)

// NewProfileRepository ProfileRepositoryの初期化
func NewProfileRepository(db *sql.DB) ProfileRepository {
	return &profileRepository{
		db: db,
	}
}

type profileRepository struct {
	db *sql.DB
}

func (r *profileRepository) transaction(txFunc func(*sql.Tx) error) (err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return
}

func (r *profileRepository) Create(profile *model.Profile) error {
	return r.transaction(func(tx *sql.Tx) error {
		_, err := tx.Exec(
			`INSERT INTO profile(id, sid, created_at, expires, deleted, name, message, time_limit, color, avatar_url)
			VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			profile.ID,
			profile.SID,
			profile.CreatedAt,
			profile.Expires,
			profile.Deleted,
			profile.Name,
			profile.Message,
			profile.Limit,
			profile.Color,
			profile.AvatarURL,
		)
		if err != nil {
			return err
		}
		tagStmt, err := tx.Prepare(
			`INSERT INTO tag(id, profile_id, category, detail)
			VALUES(?, ?, ?, ?)`,
		)
		if err != nil {
			return err
		}
		defer tagStmt.Close()
		for _, tag := range profile.Tag {
			_, err := tagStmt.Exec(
				tag.ID,
				tag.ProfileID,
				tag.Category,
				tag.Detail,
			)
			if err != nil {
				return err
			}

		}
		pictureStmt, err := tx.Prepare(
			`INSERT INTO picture(id, profile_id, display_order, url)
			VALUES(?, ?, ?, ?)`,
		)
		if err != nil {
			return err
		}
		defer pictureStmt.Close()
		for _, picture := range profile.Pictures {
			_, err := pictureStmt.Exec(
				picture.ID,
				picture.ProfileID,
				picture.Order,
				picture.URL,
			)
			if err != nil {
				return err
			}
		}
		coordinateStmt, err := tx.Prepare(
			`INSERT INTO coordinate(id, profile_id, lat, lng)
			VALUES(?, ?, ?, ?)`,
		)
		if err != nil {
			return err
		}
		defer coordinateStmt.Close()
		_, err = coordinateStmt.Exec(
			profile.Coordinate.ID,
			profile.Coordinate.ProfileID,
			profile.Coordinate.Lat,
			profile.Coordinate.Lng,
		)
		return err
	})
}

func (r *profileRepository) GetOne(uid uuid.UUID) (*model.Profile, error) {
	profile := new(model.Profile)
	if err := r.db.QueryRow(
		`SELECT profile.id, sid, created_at, expires, deleted, name, message, time_limit, color, avatar_url, coordinate.id, profile_id, lat, lng
		FROM profile
		WHERE profile.id = ?
		INNER JOIN coordinate ON
			profile.id = coordinate.profile_id`,
		uid,
	).Scan(
		&profile.ID,
		&profile.SID,
		&profile.CreatedAt,
		&profile.Expires,
		&profile.Deleted,
		&profile.Name,
		&profile.Message,
		&profile.Limit,
		&profile.Color,
		&profile.AvatarURL,
	); err != nil {
		return nil, err
	}
	return profile, nil
}

func (r *profileRepository) GetList(sid uuid.UUID) ([]*model.Profile, error) {
	return nil, nil
}

func (r *profileRepository) Delete(sid uuid.UUID) error {
	return r.transaction(func(tx *sql.Tx) error {
		_, err := tx.Exec(
			`UPDATE profile SET deleted = TRUE WHERE sid = ?;`,
			sid.ID(),
		)
		return err
	})
}
