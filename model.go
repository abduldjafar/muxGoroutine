package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// structur of Userhara table
type BukaLapakQu struct {
	Id        string
	Condition string
	Name      string
}

func DBMigrationAccount(db *gorm.DB, dbs interface{}) *gorm.DB {
	db.AutoMigrate(dbs)
	return db
}
