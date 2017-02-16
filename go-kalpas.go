package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

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
