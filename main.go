package main

import (
	"log"

	"time"

	"encoding/json"
	"fmt"

	"math/rand"

	"github.com/Juju-62q/BlogAlartRegister/db"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Member is struct of OthloMember
type Member struct {
	ID           string    `gorm:"column:id"`
	Name         string    `gorm:"column:name"`
	SlackName    string    `gorm:"column:slackName"`
	GraduateDate time.Time `gorm:"column:graduateDate"`
}

// TableName is name of DB table in Django
func (m *Member) TableName() string {
	return "OthloMember_member"
}

func main() {
	database, err := db.GetDB()
	defer database.Close()
	if err != nil {
		log.Fatal(err)
	}
	var allMember []Member
	database.Find(&allMember)
	shuffle(allMember)
	for _, member := range allMember {
		bytes, err := json.Marshal(member)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(bytes))
	}
}

func shuffle(data []Member) {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
	return
}
