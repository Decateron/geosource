package transactions

import (
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/joshheinrichs/geosource/server/config"
	"github.com/joshheinrichs/geosource/server/types"
)

func AddPosts() error {
	// Adding a user
	user := types.User{ID: "admin"}
	err := AddUser(&user)
	if err != nil {
		return err
	}
	// Adding a channel
	channel := types.Channel{
		ChannelInfo: types.ChannelInfo{
			Name:       "temp",
			CreatorID:  "admin",
			Visibility: "public",
		},
	}
	err = AddChannel(&channel)
	if err != nil {
		return err
	}
	// Adding 100,000 posts at random locations
	post := types.Post{
		PostInfo: types.PostInfo{
			CreatorID: "admin",
			Channel:   "temp",
			Title:     "",
			Thumbnail: "",
		},
	}
	for i := 0; i < 10000; i++ {
		post.ID = strconv.Itoa(i)
		post.Time = time.Now()
		post.Location = &types.Location{
			Latitude:  math.Mod(rand.Float64(), 180) - 90,
			Longitude: math.Mod(rand.Float64(), 360) - 180,
		}
		err = AddPost("admin", &post)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestMain(m *testing.M) {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	testConfig := config.New()
	err := testConfig.ReadFile("../config_test.gcfg")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = Init(testConfig)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = AddPosts()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}
