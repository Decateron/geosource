package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/garyburd/redigo/redis"
    "encoding/json"
    "strings"
    "flag"
    "time"
    "regexp"
    // "fmt"
)

var (
    pool *redis.Pool
    redisServer = flag.String("127.0.0.1", ":6379", "")
    redisPassword = flag.String("", "", "")
)

type Error struct {
	Error string `json:"error"`
}

type Post struct {
	Title    string   `json:"title"`
	User     string   `json:"user"`
	Channel  string   `json:"channel"`
	Location Location `json:"location"`
	Fields   []Field  `json:"fields"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type User struct {
	Username *string `json:"username"`
}

type Channel struct {
	Channelname *string `json:"channelname"`
	User        *string `json:"user"`
}

type Authentication struct {
	Username *string `json:"username"`
}

type Checkbox struct {
	Label string `json:"label"`
	Value bool   `json:"value"`
}

func newPool(server, password string) *redis.Pool {
    return &redis.Pool{
        MaxIdle: 3,
        IdleTimeout: 240 * time.Second,
        Dial: func () (redis.Conn, error) {
            c, err := redis.Dial("tcp", server)
            if err != nil {
                return nil, err
            }
            // if _, err := c.Do("AUTH", password); err != nil {
            //     c.Close()
            //     return nil, err
            // }
            return c, err
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            _, err := c.Do("PING")
            return err
        },
    }
}

func addPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var post Post
	err := decoder.Decode(&post)
	if err != nil {
		json.NewEncoder(w).Encode(Error{err.Error()})
		return
	}
	json.NewEncoder(w).Encode(post)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	// id := mux.Vars(r)["id"]
	post := Post {
		Title: "Check out this cat!",
		User: "Decateron",
		Channel: "Animals",
		Location: Location{100.0, 50.0},
		Fields: []Field {
			Field {
				Type: "CHECKBOXES",
				Label: "Animal",
				Value: []Checkbox{Checkbox{"Cat", true}, Checkbox{"Dog", false}},
			},
			Field {
				Type: "IMAGE",
				Label: "Picture of Animal",
				Value: "http://www.google.ca/",
			},
		},
	}
	// fmt.Println(post.Fields[0].Value[0].Label)
	json.NewEncoder(w).Encode(post)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var user User

	err := decoder.Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(Error{err.Error()})
		return
	}

	match, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", *user.Username)
	if !match {
		json.NewEncoder(w).Encode(Error{"Invalid character(s)."})
		return
	}

	id := "user:" + strings.ToLower(*user.Username)

	conn := pool.Get()
    exists, err := redis.Bool(conn.Do("EXISTS", id))
	if err != nil {
		json.NewEncoder(w).Encode(Error{err.Error()})
		return
	}
	if exists {
		json.NewEncoder(w).Encode(Error{"A user with that name already exists."})
		return
	}

	_, err = conn.Do("SET", id, user.Username)
	if err != nil {
		json.NewEncoder(w).Encode(Error{err.Error()})
		return
	}

	json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userid := mux.Vars(r)["id"]

	match, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", userid)
	if !match {
		json.NewEncoder(w).Encode(Error{"Invalid character(s)."})
		return
	}

	id := "user:" + strings.ToLower(userid)

	conn := pool.Get()
	resp, err := conn.Do("GET", id)
	if err != nil {
		json.NewEncoder(w).Encode(Error{err.Error()})
		return
	}
	if resp == nil {
		json.NewEncoder(w).Encode(Error{"User does not exist."})
		return
	}
	username, _ := redis.String(resp, err)
	json.NewEncoder(w).Encode(User{&username})
}

func addChannel(w http.ResponseWriter, r *http.Request) {
}

func getChannel(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	json.NewEncoder(w).Encode(id)
}

func main() {
	flag.Parse()
    pool = newPool(*redisServer, *redisPassword)

    r := mux.NewRouter()

    r.HandleFunc("/api/add/post", addPost)
    r.HandleFunc("/api/get/post/", getPost)
    r.HandleFunc("/api/get/post/{id}", getPost)

    r.HandleFunc("/api/add/user", addUser)
    r.HandleFunc("/api/get/user/{id}", getUser)

    r.HandleFunc("/api/add/channel", addChannel)
    r.HandleFunc("/api/get/channel/{id}", getChannel)

    r.PathPrefix("/").Handler(http.FileServer(http.Dir("app/")))
    http.Handle("/", r)
    http.ListenAndServe(":8000", nil)
}
