package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/go-martini/martini"
	"github.com/pivotal-golang/bytefmt"
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
	m.Get("/disk", func() (int, string) {
		var stat syscall.Statfs_t
		err := syscall.Statfs("/", &stat)
		if err != nil {
			return 500, "Unable to stat root filesystem."
		}
		bytes := stat.Blocks * uint64(stat.Bsize)
		return 200, bytefmt.ByteSize(bytes)
	})
	m.Run()
}
