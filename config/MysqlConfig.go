package config

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewMysqlConfigurator() *gorm.DB {
	user := os.Getenv("USER_DB")
	password := os.Getenv("PASSWORD_DB")
	host := os.Getenv("HOST_DB")
	port := os.Getenv("PORT_DB")
	name := os.Getenv("NAME_DB")

	db, err := gorm.Open(
		"mysql",
		user+":"+password+"@tcp("+host+":"+port+")/"+name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	return db
}
