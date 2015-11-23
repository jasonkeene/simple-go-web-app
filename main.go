package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-martini/martini"
)

func main() {
	message := os.Getenv("MESSAGE")
	if message == "" {
		message = "Hello world"
	}

	m := martini.Classic()
	m.Get("/", func() string {
		fmt.Println(message)
		return message
	})
	m.Get("/env", func(rw http.ResponseWriter) (int, []byte) {
		bytes, err := json.MarshalIndent(os.Environ(), "", "    ")
		if err != nil {
			return 500, []byte("Unable to marshal environment into JSON.")
		}
		rw.Header().Add("Content-Type", "application/json")
		return 200, bytes
	})
	m.Run()
}
