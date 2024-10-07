package postgres

import "gorm.io/gorm"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = ConnectDB()
	return Db
}

func ConnectDB() *gorm.DB {
	//proceed env variables
}
