package queue

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type Task struct {
	Queue   string
	Message []byte
}

type AMQPError struct {
	Err error
	Retries int
}

type ChannelConnectError struct {
	Name string
	Retries int
	Err error
}

type Connection struct {
	Queue string
	Conn *amqp.Connection
	Channel *amqp.Channel
}

var (
	connectionPool = make(map[string]*Connection)
	errChan = make(chan AMQPError)
	delayTime = 1 * time.Second
	queues = make(map[string]func(m []byte) error)
)

func connect(dsn string, queueName string) (*Connection, error) {
	if c, ok := connectionPool[queueName]; ok {
		return c, nil
	}
	connRabbitMQ, err := amqp.Dial(dsn)
	if err != nil {
		time.Sleep(delayTime)
		return connect(dsn, queueName)
	}
	ch, err := connRabbitMQ.Channel()
	if err != nil {
		defer connRabbitMQ.Close()
		defer ch.Close()
		return nil, err
	}
	c := &Connection{
		Queue:   queueName,
		Conn:    connRabbitMQ,
		Channel: ch,
	}
	connectionPool[queueName] = c
	go func() {
		<-connRabbitMQ.NotifyClose(make(chan *amqp.Error)) //Listen to NotifyClose
		errChan <- AMQPError{Err: errors.New("Connection Closed"), Retries: 0}
	}()

	return c, nil
}

func connectPub(dsn string, queueName string) (*Connection, error) {
	connRabbitMQ, err := amqp.Dial(dsn)
	if err != nil {
		time.Sleep(delayTime)
		return connect(dsn, queueName)
	}
	ch, err := connRabbitMQ.Channel()
	if err != nil {
		defer connRabbitMQ.Close()
		defer ch.Close()
		return nil, err
	}
	c := &Connection{
		Queue:   queueName,
		Conn:    connRabbitMQ,
		Channel: ch,
	}
	return c, nil
}

func Publish(dsn string, t *Task) error {
	go func() {
		c, err := connectPub(dsn, t.Queue)
		if err != nil {
			return
		}
		defer c.Conn.Close()
		defer c.Channel.Close()
		// Attempt to publish a message to the queue.
		_ = c.Channel.Publish(
			"",
			t.Queue,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        t.Message,
			},
		)
	}()
	return nil
}

func CreateQueue(dsn string, q string) error {
	if dsn == "" || q == "" {
		return errors.New("dsn and queue name must be initiated")
	}
	c, err := connect(dsn, q)
	if err != nil {
		return err
	}
	_, err = c.Channel.QueueDeclare(
		q,
		true,
		false,
		false,
		false,
		nil,
	)
	return nil
}

func Consume(dsn string, q string, f func(m []byte) error) error {
	if q == "" {
		return errors.New("queue name must be initiated")
	}
	errDeclare := CreateQueue(dsn, q)
	if errDeclare != nil {
		return errDeclare
	}
	c, err := connect(dsn, q)
	if err != nil {
		return err
	}
	// Start delivering queued messages.
	messages, err := c.Channel.Consume(
		q,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	// Build a welcome message.
	log.Println(fmt.Sprintf("Successfully connected to RabbitMQ channel %s", q))
	forever := make(chan bool)
	for message := range messages {
		mg := message
		go func() error {
			err = f(mg.Body)
			if err != nil {
				return err
			}
			return nil
		}()
	}
	if connErr := <-errChan; connErr.Err != nil {
		fmt.Println("reconnect to RabbitMQ")
		if con, ok := connectionPool[q]; ok {
			con.Conn.Close()
			delete(connectionPool, q)
			delete(queues, q)
		}
		Consume(dsn, q, f)
	}
	<-forever
	return nil
}
