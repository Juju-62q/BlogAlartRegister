package DB

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func TestConnect() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=othloblog dbname=postgres password=hoslocoffee1 sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
}