package user

import (
	"database/sql"
	"log"

	"github.com/agungcandra/dao/entity"
)

// CreateUser to database
func CreateUser(db *sql.DB, user entity.User) (entity.User, error) {
	result, err := db.Exec("INSERT INTO users(username, password, name) VALUES(?, ?, ?)", user.Username, user.Password, user.Name)

	if err != nil {
		log.Println("error" + err.Error())
	}

	userID, _ := result.LastInsertId()
	user.ID = userID
	return user, err
}
