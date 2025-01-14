package objectstream

import (
	"fmt"
	"io"
	"net/http"
)

type PutStream struct {
	w *io.PipeWriter
	c chan error
}

func NewPutStream(server, object string) *PutStream {
	r, w := io.Pipe()
	c := make(chan error)
	go func() {
		request, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%s/objects/%s", server, object), r)
		client := http.Client{}
		r, e := client.Do(request)
		if e == nil && r.StatusCode != http.StatusOK {
			e = fmt.Errorf("dataServer return http code %d", r.StatusCode)
		}
		c <- e
	}()
	return &PutStream{w, c}
}

func (c *PutStream) Write(p []byte) (n int, err error) {
	return c.w.Write(p)
}

func (c *PutStream) Close() error {
	c.w.Close()
	return <-c.c
}
