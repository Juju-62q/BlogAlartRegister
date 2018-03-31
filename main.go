package main

import (
	"log"

	"time"

	"encoding/json"
	"fmt"

	"github.com/Juju-62q/BlogAlartRegister/db"
)

type Member struct {
	ID           string    `gorm:"column:id"`
	Name         string    `gorm:"column:name"`
	SlackName    string    `gorm:"column:slackName"`
	GraduateDate time.Time `gorm:"column:graduateDate"`
}

func (m *Member) TableName() string {
	return "OthloMember_member"
}

func main() {
	database, err := db.GetDB()
	defer database.Close()
	if err != nil {
		log.Fatal(err)
	}
	var allMember []*Member
	database.Find(&allMember)
	for _, member := range allMember {
		bytes, err := json.Marshal(member)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(bytes))
	}
}
