package configs

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
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	MPosGORM, err = gorm.Open(dialect, psqlInfo)
	if err != nil {
		log.Fatalf("Cannot Connect To DB : %s", err)
	}
}