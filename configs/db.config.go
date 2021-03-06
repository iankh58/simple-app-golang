package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"database/sql"

	_ "github.com/lib/pq"
)

var MPosDB *sql.DB
var MPosGORM *gorm.DB
var err error

func InitGormPostgres() {
	//using static
	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	// MPosGORM, err = gorm.Open(dialect, psqlInfo)

	//using environment
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dialect := os.Getenv("DB_DIALECT")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	MPosGORM, err = gorm.Open(dialect, psqlInfo)

	if err != nil {
		log.Fatalf("Cannot Connect To DB : %s", err)
	}
}
