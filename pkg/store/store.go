package store

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/la4zen/dogma-intwrview/pkg/models"
	_ "github.com/lib/pq"
)

type Store struct {
	DB *sqlx.DB
}

/*
Очень хотелось заюзать GORM, но это было бы очень просто.
Поэтому я использовал sqlx. Похожий на стандарт либу sql,
просто со своими плюшками, например сканирование структуры
*/

func NewStore() *Store {
	db, err := sqlx.Open("postgres", "host=database user=postgres password=asfdsadfas database=dogma sslmode=disable")
	if err != nil {
		log.Fatalf(err.Error())
	}
	return &Store{
		DB: db,
	}
}

func (s *Store) GetUsers() (*[]models.User, error) {
	users := &[]models.User{}
	err := s.DB.Select(users, "SELECT id, login FROM users")
	if err != nil {
		if err == sql.ErrNoRows {
			return users, nil
		}
		return nil, err
	}
	return users, nil
}

func (s *Store) EditUser(user *models.User) error {
	if _, err := s.DB.Query("UPDATE users SET login=$1 WHERE id=$2", user.Login, user.ID); err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteUser(user *models.User) error {
	if _, err := s.DB.Query("DELETE FROM users WHERE id=$1 OR login=$2", user.ID, user.Login); err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateUser(user *models.User) error {
	if _, err := s.DB.Exec("INSERT INTO users(login) VALUES($1)", user.Login); err != nil {
		return err
	}
	return nil
}
