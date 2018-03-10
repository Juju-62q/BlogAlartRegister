package main

import (
	"github.com/Juju-62q/BlogAlartRegister/g_calendar"
	"time"
	"log"
)

func main(){
	err := g_calendar.AddEvent("test", "Toyohashi", "testing golang api", time.Now(), time.Now().AddDate(0,0,1))
	if err != nil{
		log.Fatal(err)
	}
}
