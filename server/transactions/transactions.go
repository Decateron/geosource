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
	db, err = gorm.Open("postgres", fmt.Sprintf("host=%s dbname=%s user=%s password=%s",
		config.Database.Host, config.Database.Database, config.Database.User, config.Database.Password))
	if err != nil {
		return err
	}
	return nil
}
