package dao

import (
	"database/sql"
)

// Dao hold all business logic and process
type Dao struct {
	database *sql.DB
}

// NewDao Create Main Container
func NewDao(database *sql.DB) *Dao {
	return &Dao{
		database: database,
	}
}

func (dao *Dao) Database() *sql.DB {
	return dao.database
}
