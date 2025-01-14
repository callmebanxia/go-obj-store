package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func get(w http.ResponseWriter, r *http.Request) {

	f, e := os.Open(os.Getenv("STROE_ROOT" + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2]))

	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer f.Close()
	_, err := io.Copy(w, f)
	if err != nil {
		log.Println(err)
	}
}
