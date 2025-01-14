package locate

import (
	"os"
	"strconv"
	"time"

	"github.com/callmebanxia/go-obj-store/pkg/mq/rabbitmq"
)

func Locate(name string) string {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	q.Publish("dataServer", name)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()

	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s
}

func Exist(name string) bool {
	return Locate(name) != ""
}
