package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// put
// 将客户端上传的内容存储到服务器
func put(w http.ResponseWriter, r *http.Request) {
	f, e := os.Create(os.Getenv("STORE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])

	// HACK  使用更加隔离的错误处理方式代替 panic
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer f.Close()
	_, err := io.Copy(f, r.Body)
	if err != nil {
		log.Println(e)
	}
}
