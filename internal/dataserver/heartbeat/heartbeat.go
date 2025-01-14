package heartbeat

import (
	"os"
	"time"

	"github.com/callmebanxia/go-obj-store/pkg/mq/rabbitmq"
)

// StartHeartBeat 发起心跳
// 定时向接口服务发送心跳，帮助接口服务了解数据服务状态
func StartHeartBeat() {

	// XXX 该处使用读取环境变量的方式获取 mq 的地址，后期可以升级为读取配置
	q := rabbitmq.New(os.Getenv("RABBITMQ_LISTEN"))
	defer q.Close()

	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)
	}

}
