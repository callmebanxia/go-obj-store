package main

import (
	"log"
	"net/http"
	"os"

	"github.com/callmebanxia/go-obj-store/internal/apiserver/heartbeat"
	"github.com/callmebanxia/go-obj-store/internal/apiserver/locate"
	"github.com/callmebanxia/go-obj-store/internal/apiserver/objects"
)

func main() {
	go heartbeat.ListenHeartBeat()
	http.HandleFunc("/object/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
