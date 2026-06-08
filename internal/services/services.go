package services

import (
	"database/sql"
	"som/internal/database"
)

var db *sql.DB
var dbq *database.Queries

func SetDatabase(sqlDB *sql.DB, databaseQueries *database.Queries) {
	db = sqlDB
	dbq = databaseQueries
}

func GetDatabaseQueries() *database.Queries {
	return dbq
}

func GetSQLDB() *sql.DB {
	return db
}
