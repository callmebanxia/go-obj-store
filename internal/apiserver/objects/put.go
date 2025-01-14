package objects

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/callmebanxia/go-obj-store/internal/apiserver/heartbeat"
	"github.com/callmebanxia/go-obj-store/internal/pkg/objectstream"
)

func put(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	c, e := storeObject(r.Body, object)
	if e != nil {
		log.Println(e)
	}
	w.WriteHeader(c)

}

// putStream
// 随机选择在线的数据服务器 创建上传流
func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("connt find any dataServer")
	}
	return objectstream.NewPutStream(server, object), nil
}

// storeObject
// 将文件内容复制到接收者
func storeObject(r io.Reader, object string) (int, error) {
	stream, e := putStream(object)
	if e != nil {
		return http.StatusServiceUnavailable, e
	}
	_, e = io.Copy(stream, r)
	if e != nil {
		log.Println(e)
	}
	e = stream.Close()
	if e != nil {
		return http.StatusInternalServerError, e
	}

	return http.StatusOK, nil
}
