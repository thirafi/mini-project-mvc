package config

import (
	"fmt"
	"os"
	"project_structure/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// "DB_Username": "root",
// "DB_Password": "Qoala202!",
// "DB_Port":     "3306",
// "DB_Host":     "altadb.c3fzsopw9aux.ap-southeast-1.rds.amazonaws.com",
// "DB_Name":     "structure_project",
func InitDB() *gorm.DB {
	// config := map[string]string{
	// 	"DB_Username": "root",
	// 	"DB_Password": "Qoala202!",
	// 	"DB_Port":     "3306",
	// 	"DB_Host":     "database-2.c3fzsopw9aux.ap-southeast-1.rds.amazonaws.com",
	// 	"DB_Name":     "mini_project",
	// }

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_Username"),
		os.Getenv("DB_Password"),
		os.Getenv("DB_Host"),
		os.Getenv("DB_Port"),
		os.Getenv("DB_Name"))

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
	return DB
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{}, &model.Book{}, &model.Blog{})
}
