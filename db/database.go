package db

import (
	"io/ioutil"

	"encoding/json"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

// Confidential is struct of db confidential
type Confidential struct {
	User     string
	Password string
	DB       string
}

var instance *gorm.DB

// GetDB gets db instance
func GetDB() (*gorm.DB, error) {
	if instance != nil {
		return instance, nil
	}

	confidential, err := readEnv()
	if err != nil {
		return nil, err
	}

	args := fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable", confidential.User, confidential.DB, confidential.Password)
	instance, err = gorm.Open("postgres", args)
	return instance, err
}

func readEnv() (Confidential, error) {
	raw, err := ioutil.ReadFile("./.env")
	if err != nil {
		return Confidential{}, errors.Wrap(err, "can't read dotenv")
	}

	var confidential Confidential
	err = json.Unmarshal(raw, &confidential)
	if err != nil {
		return Confidential{}, errors.Wrap(err, "can't serialize to json")
	}

	return confidential, nil
}
