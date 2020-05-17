package env

import "os"

// Env 環境変数を返す関数まとめ
type Env interface {
	DBUser() string
	DBPass() string
	DBProtocol() string
	DBAddress() string
	DBName() string
}

// NewEnv Envの初期化(*env)
func NewEnv() Env {
	e := &env{
		dbUser:     os.Getenv("DB_USER"),
		dbPass:     os.Getenv("DB_PASS"),
		dbProtocol: os.Getenv("DB_PROTOCOL"),
		dbAddress:  os.Getenv("DB_ADDRESS"),
		dbName:     os.Getenv("DB_NAME"),
	}
	return e
}

type env struct {
	dbUser     string
	dbPass     string
	dbProtocol string
	dbAddress  string
	dbName     string
}

func (e *env) DBUser() string { return e.dbUser }

func (e *env) DBPass() string { return e.dbPass }

func (e *env) DBProtocol() string { return e.dbProtocol }

func (e *env) DBAddress() string { return e.dbAddress }

func (e *env) DBName() string { return e.dbName }
