// Package transactions provides a set of functions which allow for interaction
// with the database.
package transactions

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/joshheinrichs/geosource/server/config"
	"github.com/joshheinrichs/httperr"

	// This is not imported in main to keep all logic about the database inside the transactions package
	_ "github.com/lib/pq"
)

var db *gorm.DB

var ErrInsufficientPermission = httperr.New("Insufficient permission.", http.StatusForbidden)
var ErrNotImplemented = httperr.New("function has not yet been implemented", http.StatusInternalServerError)

// Init opens a connection to the database based on the information in the
// given config. Returns an error if the connection  could not be established.
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
