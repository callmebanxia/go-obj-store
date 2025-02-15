package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	f, e := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])

	if e != nil {
		log.Println(e)
	}

	w.WriteHeader(http.StatusInternalServerError)

	defer f.Close()

	if _, e = io.Copy(f, r.Body); e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
