package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	tarouSQL "github.com/anatofuz/BookmarkTarou/infra/store/mysql"

	"github.com/anatofuz/BookmarkTarou/infra/store"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/xerrors"
)

// AppComponent is a DI container of Bookmark app.
type AppComponent interface {
	UserStore() store.UserStore
}

type appComponentImpl struct {
	bookmarkDB *sql.DB
}

type config struct {
	port int
	dsn  string
}

func createDiConfig() (*config, error) {
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

	return &config{
		port: port,
		dsn:  dsn,
	}, nil
}

// CreateAppComponent DIcontainerのコンストラクタ
func CreateAppComponent() (AppComponent, error) {
	config, err := createDiConfig()
	if err != nil {
		return nil, fmt.Errorf("failed create DiConfig : %w", err)
	}
	fmt.Println(config.dsn)
	db, err := sql.Open("mysql", config.dsn)
	if err != nil {
		return nil, xerrors.Errorf("failed open sqlx : %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &appComponentImpl{bookmarkDB: db}, nil
}

func (app *appComponentImpl) UserStore() store.UserStore {
	return tarouSQL.NewUserStore(app.bookmarkDB)
}
