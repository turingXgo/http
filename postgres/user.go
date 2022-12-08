package postgres

import (
	"database/sql"

	"github.com/turingXgo/http/model"
)

type PG struct {
	db *sql.DB
}

func New(db *sql.DB) *PG {
	return &PG{
		db: db,
	}
}

func (pg PG) GetUser(name string) (*model.User, error) {
	var user model.User
	err := pg.db.QueryRow("SELECT name, age FROM data WHERE name = $1", name).
		Scan(&user.Name, &user.Age)
	return &user, err
}

func (pg PG) InsertUser(user *model.User) error {
	_, err := pg.db.Exec("INSERT INTO data (name, age) VALUES ($1, $2)", user.Name, user.Age)
	return err
}
