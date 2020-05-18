package database

import (
	"database/sql"
	"fmt"

	// mysqlドライバ
	_ "github.com/go-sql-driver/mysql"
	"github.com/kakohate/charamell-mvp/env"
)

// New *sql.DBを初期化するだけ
func New(e env.Env) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@%s(%s)/%s",
		e.DBUser(), e.DBPass(), e.DBProtocol(), e.DBAddress(), e.DBName(),
	)
	return sql.Open("mysql", dsn)
}

// Init データベース作成
func Init(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS charamell.profile( id VARCHAR(36) NOT NULL PRIMARY KEY,
	sid VARCHAR(36) NOT NULL,
	created_at datetime NOT NULL DEFAULT current_timestamp,
	deleted bool NOT NULL DEFAULT FALSE,
	name text NOT NULL,
	message text NOT NULL,
	time_limit int NOT NULL DEFAULT 1,
	color varchar(36),
	avatar_url text NOT NULL );`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS charamell.tag( id VARCHAR(36) NOT NULL PRIMARY KEY,
	profile_id VARCHAR(36) NOT NULL,
	created_at datetime NOT NULL DEFAULT current_timestamp,
	category varchar(36) NOT NULL,
	detail text NOT NULL );`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS charamell.picture( id VARCHAR(36) NOT NULL PRIMARY KEY,
	profile_id VARCHAR(36) NOT NULL,
	created_at datetime NOT NULL DEFAULT current_timestamp,
	display_order int NOT NULL,
	url text NOT NULL );`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS charamell.coordinate( id VARCHAR(36) NOT NULL PRIMARY KEY,
	profile_id VARCHAR(36) NOT NULL,
	created_at datetime NOT NULL DEFAULT current_timestamp,
	lat double NOT NULL,
	lng double NOT NULL );`)
	return err
}
