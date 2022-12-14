package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB Global DB connection
var DB *gorm.DB

func init() {
	var err error
	// dsn := fmt.Sprintf(
	// 	"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// )
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		"rails",
		"rails",
		"gin_qor",
	)

	DB, err = gorm.Open("postgres", dsn)

	if err == nil {
		if os.Getenv("DEBUG") != "" {
			DB.LogMode(true)
		}

	} else {
		panic(err)
	}
}
