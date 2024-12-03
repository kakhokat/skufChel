package store

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Dbname   string
	Sslmode  string
}

func NewConfig(Host string, Port int, Username, Password, Dbname, Sslmode string) Config {
	return Config{
		Host:     Host,
		Port:     Port,
		Username: Username,
		Password: Password,
		Dbname:   Dbname,
		Sslmode:  Sslmode,
	}
}

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

type Notifications interface {
	GetNotifications(userId int)
}

func InitPostgres(cfg Config) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("postgres", fmt.Sprintf("host =%s port =%d user =%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Dbname, cfg.Password, cfg.Sslmode))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (r *Store) GetNotifications(id int) {

}
