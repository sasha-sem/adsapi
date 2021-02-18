package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//InitDb initialises db
func InitDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	if !db.Migrator().HasTable(&Advertisement{}) {
		db.Migrator().CreateTable(&Advertisement{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&Advertisement{})
	}

	return db
}
