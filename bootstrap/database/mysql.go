package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/agungcandra/dao/bootstrap/environment"
)

// Connection : Connection variable
var connection *sql.DB

// Connect : Connect to database pool
func Connect() {
	var datasource = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", environment.DATABASEUSERNAME, environment.DATABASEPASSWORD, environment.DATABASEHOST, environment.DATABASEPORT, environment.DATABASENAME)
	conn, err := sql.Open("mysql", datasource)

	if err != nil {
		log.Fatalln(err)
	}

	connection = conn
	defer connection.Close()
}

// GetConnection : Get connection
func GetConnection() *sql.DB {
	if connection == nil {
		Connect()
		return connection
	}

	return connection
}
