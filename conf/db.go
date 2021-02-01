package conf

import (
	"bee.com/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		C.Mysql.Username, C.Mysql.Password, C.Mysql.Path, C.Mysql.Dbname, C.Mysql.Config)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db connect error: ", err)
	}
	db.AutoMigrate(&models.Product{})

	DB = db
}
