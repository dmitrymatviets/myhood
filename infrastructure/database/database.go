package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DatabaseConfig struct {
	Server       string `envconfig:"server"`
	User         string `envconfig:"user"`
	Password     string `envconfig:"password"`
	Port         int    `envconfig:"port"`
	DatabaseName string `envconfig:"database"`
	MaxOpenConns int    `envconfig:"maxopenconns"`
	MaxIdleConns int    `envconfig:"maxidleconns"`
	AppName      string `envconfig:"appname"`
}

type Database struct {
	Db *sqlx.DB
}

func NewDatabase(c DatabaseConfig) (*Database, error) {
	cp := new(Database)
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;app name=%s",
		c.Server,
		c.User,
		c.Password,
		c.Port,
		c.DatabaseName,
		c.AppName,
	)

	db, err := sqlx.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	cp.Db = db
	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)

	return cp, nil
}
