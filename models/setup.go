package models
import (
"log"
_ "github.com/go-sql-driver/mysql"
"github.com/jinzhu/gorm"
)
// SetupDB : initializing mysql database
func SetupDB() *gorm.DB{

db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/dts_task?charset=utf8&parseTime=True&loc=Local")
if err != nil {
	log.Fatal(err.Error())
}
return db
}