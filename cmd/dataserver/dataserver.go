package dataserver

import (
	"log"
	"net/http"
	"os"

	"github.com/callmebanxia/go-obj-store/internal/dataserver/heartbeat"
	"github.com/callmebanxia/go-obj-store/internal/dataserver/locate"
	"github.com/callmebanxia/go-obj-store/internal/dataserver/objects"
)

func main() {
	go heartbeat.StartHeartBeat()
	go locate.StartLocate()
	http.HandleFunc("/object", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
