package main

import (
	"github.com/Juju-62q/BlogAlartRegister/g_calendar"
	"context"
	"io/ioutil"
	"log"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"fmt"
)

func main(){
	ctx := context.Background()

	authData, err := ioutil.ReadFile("/home/kenya/client_secret.json")
	if err != nil{
		log.Fatal(err)
	}

	config, err := google.ConfigFromJSON(authData, calendar.CalendarReadonlyScope)
	if err != nil{
		log.Fatal(err)
	}


	client := g_calendar.GetClient(ctx, config)
	if client != nil{
		fmt.Println("success")
	}else{
		fmt.Println("failed")
	}
}
