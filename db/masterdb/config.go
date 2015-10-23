package masterdb

import (
	"github.com/FoxComm/libs/configs"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"log"
)

var DB *gorm.DB

func Db() *gorm.DB {
	if DB == nil {
		if db, err := gorm.Open("postgres", configs.Get("FC_CORE_DB_URL")); err != nil {
			log.Printf("databaseErr=%s", err)
			panic("Could not open database")
		} else {
			DB = &db
			DB.LogMode(true)
		}
	}
	return DB
}
