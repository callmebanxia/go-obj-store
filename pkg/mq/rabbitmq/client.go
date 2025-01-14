package rabbitmq

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

// RabbitMQ  rabbitMq 链接实例
type RabbitMQ struct {
	channel  *amqp.Channel
	conn     *amqp.Connection
	Name     string
	exchange string
}

// New 新建一个RabbitMq链接实例
// Param: s string  rabbitmq 监听端口
func New(s string) *RabbitMQ {

	conn, e := amqp.Dial(s)
	if e != nil {
		panic(e)
	}

	ch, e := conn.Channel()
	if e != nil {
		panic(e)
	}

	q, e := ch.QueueDeclare(
		"", // name
		false,
		true,
		false,
		false,
		nil,
	)

	if e != nil {
		panic(e)
	}

	mq := new(RabbitMQ)
	mq.channel = ch
	mq.conn = conn
	mq.Name = q.Name
	return mq

}

func (q *RabbitMQ) Bind(exchange string) {
	e := q.channel.QueueBind(
		q.Name,
		"",
		exchange,
		false,
		nil)

	if e != nil {
		panic(e)
	}

	q.exchange = exchange
}

func (q *RabbitMQ) Send(queue string, body any) {
	str, e := json.Marshal(body)
	if e != nil {
		panic(e)
	}

	e = q.channel.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ReplyTo: q.Name,
			Body:    []byte(str),
		})

	if e != nil {
		panic(e)
	}
}

func (q *RabbitMQ) Publish(exchange string, body any) {
	str, e := json.Marshal(body)
	if e != nil {
		panic(e)
	}

	e = q.channel.Publish(exchange,
		"", false, false, amqp.Publishing{
			ReplyTo: q.Name,
			Body:    []byte(str),
		})

	if e != nil {
		panic(e)
	}
}

func (q *RabbitMQ) Consume() <-chan amqp.Delivery {
	c, e := q.channel.Consume(q.Name, "", true, false, false, false, nil)
	if e != nil {
		panic(e)
	}

	return c
}

func (q *RabbitMQ) Close() {
	q.channel.Close()
	q.conn.Close()
}
