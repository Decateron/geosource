package transactions

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/joshheinrichs/geosource/server/config"
	_ "github.com/lib/pq"
)

var db *gorm.DB

var ErrInsufficientPermission error = errors.New("Insufficient permission.")

func Init(config *config.Config) (err error) {
	arguments := ""
	if len(config.Database.Host) > 0 {
		arguments += fmt.Sprintf("host=%s ", config.Database.Host)
	}
	if len(config.Database.Database) > 0 {
		arguments += fmt.Sprintf("dbname=%s ", config.Database.Database)
	}
	if len(config.Database.User) > 0 {
		arguments += fmt.Sprintf("user=%s ", config.Database.User)
	}
	if len(config.Database.Password) > 0 {
		arguments += fmt.Sprintf("password=%s ", config.Database.Password)
	}

	db, err = gorm.Open("postgres", arguments)
	if err != nil {
		return err
	}
	return nil
}
