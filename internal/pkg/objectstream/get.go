package objectstream

import (
	"fmt"
	"io"
	"net/http"
)

type GetStream struct {
	reader io.Reader
}

// newGetStream
// 通过文件url从服务器获取文件
func newGetStream(url string) (*GetStream, error) {
	r, e := http.Get(url)
	if e != nil {
		return nil, e
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dataServer return http code %d", r.StatusCode)
	}
	return &GetStream{r.Body}, nil
}

// NewGetStream
// 创建新的获取流 判断文件定位地址并访问
func NewGetStream(server, object string) (*GetStream, error) {
	if server == "" || object == "" {
		return nil, fmt.Errorf("invalid server %s object %s", server, object)
	}
	return newGetStream(fmt.Sprintf("http://%s/objects/%s", server, object))
}

func (r *GetStream) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}
