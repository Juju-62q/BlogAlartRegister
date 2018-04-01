package main

import (
	"log"

	"time"

	"math/rand"

	"fmt"

	"github.com/Juju-62q/BlogAlartRegister/db"
	"github.com/Juju-62q/BlogAlartRegister/g_calendar"
)

const location = "Asia/Tokyo"

func init() {
	//set rand
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

	timeZone, err := time.LoadLocation(location)
	if err != err {
		timeZone = time.FixedZone(location, 9*60*60)
	}

	date := time.Now()
	date = time.Date(date.Year(), date.Month(), date.Day(), 20, 0, 0, 0, timeZone)
	date = nextWeek(date)
	for _, member := range allMember {
		if err != nil {
			log.Fatal(err)
		}
		date = nextWeek(date)
		gcalendar.AddEvent("ブログ"+member.Name, "OthloBlog", member.SlackName, date, date.Add(1*time.Hour))

		fmt.Println(date.Format("2006/01/02") + " " + member.SlackName)
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

func nextWeek(date time.Time) time.Time {
	return date.AddDate(0, 0, 7)
}
