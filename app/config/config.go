package config

import (
	"fmt"
	"github.com/anatofuz/BookmarkTarou/infra/store/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"strconv"

	"github.com/anatofuz/BookmarkTarou/infra/store"
	"golang.org/x/xerrors"
)

// AppComponent is a DI container of Bookmark app.
type AppComponent interface {
	UserStore() store.UserStore
}

type appComponentImpl struct {
	bookmarkDB *sqlx.DB
}

type Config struct {
	port int
	dsn  string
}

func CreateDiConfig() (*Config, error) {
	dsn, ok := os.LookupEnv("DATABASE_DSN")
	if !ok {
		return nil, fmt.Errorf("failed set DATABASE_DSN")
	}

	portStr, ok := os.LookupEnv("PORT")

	if !ok {
		return nil, fmt.Errorf("failed set PORT")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, xerrors.Errorf("failed convert integer at port : %w", err)
	}

	return &Config{
		port: port,
		dsn:  dsn,
	}, nil
}

// CreateAppComponent DIcontainerのコンストラクタ
func CreateAppComponent(config *Config) (AppComponent, error) {
	db, err := sqlx.Open("mysql", config.dsn)
	if err != nil {
		return nil, xerrors.Errorf("failed open sqlx : %w", err)
	}

	return &appComponentImpl{bookmarkDB: db}, nil
}

func (app *appComponentImpl) UserStore() store.UserStore {
	return mysql.NewUserStore(app.bookmarkDB)
}
