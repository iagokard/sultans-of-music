package services

import (
	"database/sql"
	"som/internal/database"
)

type inventoryLogTypes struct {
	Entry string
	Exit  string
}

var db *sql.DB
var dbq *database.Queries

var InventoryLogType inventoryLogTypes = inventoryLogTypes{
	Entry: "entry",
	Exit:  "exit",
}

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
