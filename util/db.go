package singleton

import (
	"database/sql"
	"fmt"
	"sirclo/layered/relation/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

// func InitDB(connectionString string) (*sql.DB, error) {
// 	return sql.Open("mysql", connectionString)
// }

// err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("error load .env file")
// 	}

func MysqlDriver(config *config.AppConfig) *sql.DB {
	// var url string

	url := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name)

	db, err := sql.Open("mysql",url)

	if err != nil {
		log.Info("failed to connect database: ", err)
		panic(err)
	}

	return db
}