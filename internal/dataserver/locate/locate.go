package locate

import (
	"os"
	"strconv"

	"github.com/callmebanxia/go-obj-store/pkg/mq/rabbitmq"
)

// Locate 定位
// 判断文件名对应文件是否存在
func Locate(name string) bool {
	_, e := os.Stat(name)
	return os.IsNotExist(e)
}

// StartLocate
// 从消息队列中获取需要定位的文件名 并判断其是否存在
func StartLocate() {

	// XXX 该处使用读取环境变量的方式获取 mq 的地址，后期可以升级为读取配置
	q := rabbitmq.New(os.Getenv("RABBITMQ_LISTEN"))
	q.Bind("dataServer")
	c := q.Consume()

	for msg := range c {
		object, e := strconv.Unquote(string(msg.Body))

		// HACK 使用日志抛出错误等方式代替 panic
		if e != nil {
			panic(e)
		}

		if Locate(os.Getenv("STORE_ROOT" + "/objects/" + object)) {
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}
	}
}
