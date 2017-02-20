package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type User struct {
	id     int
	name   string
	mail   string
	phone  string
	token  string
	otp    string
	access boolean
}

func signin_handler(w http.ResponseWriter, request []string) error {
	fmt.Fprintf(w, "Inside Signin handler")
	return nil
}

func signup_handler(w http.ResponseWriter, request []string) error {
	fmt.Fprintf(w, "Inside Signup handler")
	return nil
}

func api_handler(w http.ResponseWriter, r *http.Request) {
	request := strings.Split(r.URL.Path[8:], "/")
	var err error
	switch request[0] {
	case "signin":
		err = signin_handler(w, request[1:])
		break
	case "signup":
		err = signup_handler(w, request[1:])
		break
	default:
		fmt.Fprintf(w, "Invalid resource!", r.URL.Path[1:])
		break
	}
	if err != nil {
		http.Error(w, "Something went wrong", 500)
	}
}

func database_init() {
	//Connect to Database server
	db, err := sql.open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Printf("Connection established with Database Server")

	//Check for Database
	users, err := db.Query("select * from Users")
	if err != nil {
		panic(err)
	}
}

func register_apis() {
	//serve static content
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/api/v1/", api_handler)
}

func serve() {
	port := os.Getenv("PORT")
	if port == "" {
		//panic(fmt.Errorf("$PORT not set"))
		port = strconv.FormatInt(5000, 10)
	}

	addr := ":" + port

	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func main() {
	database_init()

	register_apis()

	serve()
}
