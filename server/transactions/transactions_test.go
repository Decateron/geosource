package transactions

import (
	"log"
	"os"
	"testing"

	"github.com/joshheinrichs/geosource/server/config"
)

func TestMain(m *testing.M) {
	testConfig := config.New()
	testConfig.ReadFile("../config_test.gcfg")
	err := Init(testConfig)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}
