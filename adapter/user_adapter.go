package adapter

import (
	"context"
	"database/sql"

	"github.com/agungcandra/dao/bootstrap/database"
	"github.com/agungcandra/dao/model"
)

// UserAdapter hold function for user
type UserAdapter interface {
	InsertUser(context.Context, model.User) (sql.Result, error)
}

// InsertUser to database
func InsertUser(ctx context.Context, user model.User) (sql.Result, error) {
	db := database.GetConnection()

	statement, err := db.Exec("INSERT INTO users VALUES(?, ?, ?)", user.Username, user.Password, user.Name)
	return statement, err
}

// DeleteUser blabla
func DeleteUser(ctx context.Context, id int) (sql.Result, error) {
	db := database.GetConnection()

	statement, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	return statement, err
}
