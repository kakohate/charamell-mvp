package repository

import (
	"database/sql"
	"log"
	"strings"

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
			log.Println("repository", 1, err)
			return err
		}
		tagStmt, err := tx.Prepare(
			`INSERT INTO tag(id, profile_id, category, detail)
			VALUES(?, ?, ?, ?)`,
		)
		if err != nil {
			log.Println("repository", 2, err)
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
				log.Println("repository", 3, err)
				return err
			}

		}
		pictureStmt, err := tx.Prepare(
			`INSERT INTO picture(id, profile_id, display_order, url)
			VALUES(?, ?, ?, ?)`,
		)
		if err != nil {
			log.Println("repository", 4, err)
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
				log.Println("repository", 5, err)
				return err
			}
		}
		coordinateStmt, err := tx.Prepare(
			`INSERT INTO coordinate(id, profile_id, lat, lng)
			VALUES(?, ?, ?, ?)`,
		)
		if err != nil {
			log.Println("repository", 6, err)
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
		`SELECT profile.id, sid, profile.created_at, expires, deleted, name, message, time_limit, color, avatar_url, coordinate.id, profile_id, lat, lng
		FROM profile
		INNER JOIN coordinate ON
			profile.id = coordinate.profile_id
		WHERE profile.id = ?`,
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
		&profile.Coordinate.ID,
		&profile.Coordinate.ProfileID,
		&profile.Coordinate.Lat,
		&profile.Coordinate.Lng,
	); err != nil {
		log.Println("repository", 1, err)
		return nil, err
	}
	return profile, nil
}

func (r *profileRepository) GetOneBySID(sid uuid.UUID) (*model.Profile, error) {
	profile := new(model.Profile)
	if err := r.db.QueryRow(
		`SELECT id, sid, created_at, expires, deleted, name, message, time_limit, color, avatar_url
		FROM profile
		WHERE sid = ?`,
		sid,
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
		log.Println("repository", 1, err)
		return nil, err
	}
	rows, err := r.db.Query(
		`SELECT category, detail
		FROM tag
		WHERE profile_id = ?`,
		profile.ID,
	)
	if err != nil {
		log.Println("repository", 2, err)
		return nil, err
	}
	for rows.Next() {
		tag := new(model.Tag)
		if err := rows.Scan(&tag.Category, &tag.Detail); err != nil {
			log.Println("repository", 3, err)
			return nil, err
		}
		profile.Tag = append(profile.Tag, tag)
	}
	return profile, nil
}

func (r *profileRepository) GetList(uid uuid.UUID) ([]*model.Profile, error) {
	var lat, lng float64
	var ids = make([]interface{}, 0)
	if err := r.db.QueryRow(
		`SELECT lat, lng
		FROM coordinate
		WHERE profile_id = ?`,
		uid,
	).Scan(&lat, &lng); err != nil {
		log.Println("repository", 1, err)
		return nil, err
	}
	args := make([]interface{}, 0)
	rows, err := r.db.Query(
		`SELECT category
		FROM tag
		WHERE profile_id = ?`,
		uid,
	)
	for rows.Next() {
		var category interface{}
		if err := rows.Scan(&category); err != nil {
			log.Println("repository", 3, err)
			return nil, err
		}
		args = append(args, category)
	}
	if err != nil {
		log.Println("repository", 2, err)
	}
	stmt := `SELECT profile.id, profile.created_at, expires, deleted, name, message, time_limit, color, avatar_url
		FROM profile
		INNER JOIN coordinate ON
			profile.id = coordinate.profile_id
		INNER JOIN tag ON
			profile.id = tag.profile_id
		WHERE NOW() < profile.expires
			AND tag.category IN (?` + strings.Repeat(`, ?`, len(args)-1) + `)` + `
			AND ? < lat
			AND lat < ?
			AND ? < lng
			AND lng < ?`
	args = append(args, lat-0.2, lat+0.2, lng-0.4, lng+0.4)
	profilesMap := make(map[uuid.UUID]*model.Profile)
	rows, err = r.db.Query(stmt, args...)
	if err != nil {
		log.Println("repository", 4, err)
		return nil, err
	}
	for rows.Next() {
		profile := new(model.Profile)
		if err := rows.Scan(
			&profile.ID,
			&profile.CreatedAt,
			&profile.Expires,
			&profile.Deleted,
			&profile.Name,
			&profile.Message,
			&profile.Limit,
			&profile.Color,
			&profile.AvatarURL,
		); err != nil {
			log.Println("repository", 5, err)
			return nil, err
		}
		profilesMap[profile.ID] = profile
		ids = append(ids, profile.ID)
	}
	if len(ids) == 0 {
		return nil, nil
	}
	rows, err = r.db.Query(
		`SELECT id, profile_id, category, detail
		FROM tag
		WHERE profile_id in (?`+strings.Repeat(`, ?`, len(ids)-1)+`)`,
		ids...,
	)
	if err != nil {
		log.Println("repository", 6, err)
		return nil, err
	}
	var tags = make([]*model.Tag, 0)
	for rows.Next() {
		tag := new(model.Tag)
		if err := rows.Scan(
			&tag.ID,
			&tag.ProfileID,
			&tag.Category,
			&tag.Detail,
		); err != nil {
			log.Println(4, err)
			return nil, err
		}
		tags = append(tags, tag)
	}
	for _, tag := range tags {
		profilesMap[tag.ProfileID].Tag = append(profilesMap[tag.ProfileID].Tag, tag)
	}
	profiles := make([]*model.Profile, 0)
	for _, profile := range profilesMap {
		profiles = append(profiles, profile)
	}
	return profiles, nil
}

func (r *profileRepository) Delete(sid uuid.UUID) error {
	return r.transaction(func(tx *sql.Tx) error {
		_, err := tx.Exec(
			`UPDATE profile SET deleted = TRUE WHERE sid = ?`,
			sid.ID(),
		)
		return err
	})
}
