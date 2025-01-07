package main

import (
	"log"
	"net/http"
	"os"

	"github.com/callmebanxia/go-obj-store/v1/objects"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
