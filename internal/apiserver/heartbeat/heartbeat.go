package heartbeat

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/callmebanxia/go-obj-store/pkg/mq/rabbitmq"
)

var dataServers map[string]time.Time
var mutex sync.Mutex

// ListenHeartBeat  监听心跳
// 记录已接入的dataServers 并
func ListenHeartBeat() {

	// XXX 该处使用读取环境变量的方式获取 mq 的地址，后期可以升级为读取配置
	q := rabbitmq.New(os.Getenv("RABBITMQ_LISTEN"))
	defer q.Close()

	q.Bind("apiServers")
	c := q.Consume()
	go removeExpiredDataServer()
	for msg := range c {
		dataServer, e := strconv.Unquote(string(msg.Body))

		// HACK  使用日志以及抛出异常堆 pinic 进行替换
		if e != nil {
			panic(e)
		}
		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}

}

// removeExpiredDataServer 移除超时未收到心跳的数据服务
func removeExpiredDataServer() {
	for {
		time.Sleep(5 * time.Second)
		for s, t := range dataServers {
			if t.Add(10 * time.Second).Before(time.Now()) {
				mutex.Lock()
				delete(dataServers, s)
				mutex.Unlock()
			}
		}
	}
}

func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	ds := make([]string, 0)
	for s := range dataServers {
		ds = append(ds, s)
	}

	return ds
}
