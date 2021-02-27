package models

import(
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"database/sql"
	_ "github.com/lib/pq"
)

var MPosDB *sql.DB
var MPosGORM *gorm.DB
var err error

func InitGormPostgres(){
	host := "127.0.0.1"
	port := "5432"
	user := "postgres"
	password := "qazwsxpol"
	dbname := "simple-db"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	MPosGORM, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Cannot Connect To DB : %s", err)
	}
}