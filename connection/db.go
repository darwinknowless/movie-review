package connection

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func Connect() {
	psqlHost := os.Getenv("PSQL_HOST")
	psqlUser := os.Getenv("PSQL_USER")
	psqlPass := os.Getenv("PSQL_PASSWORD")
	psqlDb := os.Getenv("PSQL_DB")
	psqlPort := os.Getenv("PSQL_PORT")

	dsn := "host=" + psqlHost +
		" user=" + psqlUser +
		" password" + psqlPass +
		" dbName=" + psqlDb +
		" port=" + psqlPort +
		" sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	/*// https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
	  DSN: dsn,
	  PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	*/
	if err != nil {
		log.Fatal("Cannot connect to db...")
	} else {
		log.Println("Succesfully Connected")
	}
	dbConn = db
}

func GetConnection() *gorm.DB {
	return dbConn
}
