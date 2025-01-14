package objects

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/callmebanxia/go-obj-store/internal/apiserver/locate"
	"github.com/callmebanxia/go-obj-store/internal/pkg/objectstream"
)

func get(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	stream, e := getStream(object)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, err := io.Copy(w, stream)
	if err != nil {
		log.Println(err)
	}
}

func getStream(object string) (*objectstream.GetStream, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return objectstream.NewGetStream(server, object)
}
