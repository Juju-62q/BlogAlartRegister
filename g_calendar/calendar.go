package gcalendar

import (
	"io/ioutil"
	"strings"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

const mailAddressFile = "/home/kenya/mail"

// AddEvent adds event to google calendar
func AddEvent(title string, locate string, description string, startTime time.Time, endTime time.Time) error {

	event := &calendar.Event{
		Summary:     title,
		Location:    locate,
		Description: description,
		Start: &calendar.EventDateTime{
			DateTime: startTime.Format("2006-01-02T15:04:05") + "+09:00",
			TimeZone: "Asia/Tokyo",
		},
		End: &calendar.EventDateTime{
			DateTime: endTime.Format("2006-01-02T15:04:05") + "+09:00",
			TimeZone: "Asia/Tokyo",
		},
	}

	ctx := context.Background()

	authData, err := ioutil.ReadFile("/home/kenya/client_secret.json")
	if err != nil {
		return err
	}

	config, err := google.ConfigFromJSON(authData, calendar.CalendarScope)
	if err != nil {
		return err
	}

	client := GetClient(ctx, config)

	srv, err := calendar.New(client)
	if err != nil {
		return err
	}

	mailAddress, err := getMailAddress()
	if err != nil {
		return err
	}

	_, err = srv.Events.Insert(mailAddress, event).Do()
	if err != nil {
		return err
	}

	return nil
}

func getMailAddress() (string, error) {
	mailAddress, err := ioutil.ReadFile(mailAddressFile)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(mailAddress), "\n"), nil
}
