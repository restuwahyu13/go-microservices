package packages

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	ampq "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type queueResponse struct {
	Items interface{}
}

type rabbitmqInterface interface {
	Publisher(ctx context.Context, eventName string, eventMessage interface{}) bool
	Consumer(eventName string) (queueResponse, bool)
}

type rabbitmq struct {
	connection *ampq.Connection
}

var (
	rbConfig     = ampq.Config{}
	rbResponse   = queueResponse{}
	exchangeName = "tokserba.direct"
	routingKey   = "a.b.c"
	isPublished  = true
	isConsumed   = true
	json         = jsoniter.ConfigCompatibleWithStandardLibrary
)

func NewRabbit() rabbitmqInterface {
	var (
		protocol = viper.GetString("RABBITMQ_PROTOCOL")
		vhost    = viper.GetString("RABBITMQ_VHOST")
		host     = viper.GetString("RABBITMQ_HOST")
		port     = viper.GetString("RABBITMQ_PORT")
		user     = viper.GetString("RABBITMQ_USER")
		pass     = viper.GetString("RABBITMQ_PASSWORD")
	)

	rbns := protocol + "://" + user + ":" + pass + "@" + host + ":" + port + vhost // example url host like this -> amqp://admin:admin@localhost:5672/
	rbConfig.Heartbeat = time.Duration(time.Second * 30)

	connection, err := ampq.DialConfig(rbns, rbConfig)
	if err != nil {
		defer log.Fatalf("RABBITMQ NOT CONNECTED: %s", err.Error())
		defer connection.Close()
		return nil
	}

	defer log.Println("RABBITMQ CONNECTED")
	return &rabbitmq{connection: connection}
}

func (h *rabbitmq) Publisher(ctx context.Context, eventName string, eventMessage interface{}) bool {
	logrus.Printf(`START RABBITMQ PUBLISHER WITH EVENTNAME: %s`, eventName)

	channel, err := h.connection.Channel()
	if err != nil {
		defer func() {
			logrus.Printf("Rabbitmq channel error: %s", err.Error())
			channel.Close()
		}()
		isPublished = false
	}

	// overwrite routing key default name
	routingKey = eventName

	if err := channel.ExchangeDeclare(exchangeName, ampq.ExchangeDirect, true, false, false, false, nil); err != nil {
		defer logrus.Printf("Prepare exchange into broker error: %s", err.Error())
		isPublished = false
	}

	queue, err := channel.QueueDeclare(eventName, true, false, false, false, nil)
	if err != nil {
		defer logrus.Printf("Prepare queue into broker error: %s", err.Error())
		isPublished = false
	}

	if err := channel.QueueBind(queue.Name, routingKey, exchangeName, false, nil); err != nil {
		defer logrus.Printf("Binding queue into routing key error: %s", err.Error())
		isPublished = false
	}

	bodyRequest, err := json.Marshal(eventMessage)
	if err != nil {
		defer logrus.Printf("Parse request body into request queue error: %s", err.Error())
		isPublished = false
	}

	pubRequest := ampq.Publishing{
		AppId:           uuid.NewString(),
		ContentType:     "application/json",
		ContentEncoding: "application/json",
		Timestamp:       time.Now().Local(),
		Body:            bodyRequest,
	}

	if err := channel.PublishWithContext(ctx, exchangeName, routingKey, false, false, pubRequest); err != nil {
		defer logrus.Printf("Published data into queue error: %s", err.Error())
		isPublished = false
	}

	// talk to server when published msg is not delivery, then requeue message
	if err := channel.Nack(0, true, true); err != nil {
		defer logrus.Printf("Published data into queue error: %s", err.Error())
		isPublished = false
	}

	defer logrus.Printf(`RABBITMQ PUBLISHER LOGS:
	Exchange Name: %s
	Routing Name: %s
	Event Name: %s
	Request: %v
`, exchangeName, routingKey, eventName, eventMessage)

	return isPublished
}

func (h *rabbitmq) Consumer(eventName string) (queueResponse, bool) {
	logrus.Printf(`START RABBITMQ CONSUMER WITH EVENTNAME: %s`, eventName)

	channel, err := h.connection.Channel()
	if err != nil {
		defer func() {
			logrus.Printf("Rabbitmq channel error: %s", err.Error())
			channel.Close()
		}()
		isConsumed = false
	}

	// overwrite routing key default name
	routingKey = eventName

	queue, err := channel.QueueDeclare(exchangeName, true, false, true, false, nil)
	if err != nil {
		defer logrus.Printf("Prepare queue into broker error: %s", err.Error())
		isConsumed = false
	}

	if err := channel.QueueBind(queue.Name, routingKey, exchangeName, false, nil); err != nil {
		defer logrus.Printf("Binding queue into routing key error: %s", err.Error())
		isConsumed = false
	}

	res, err := channel.Consume(eventName, "", true, true, false, false, nil)
	if err != nil {
		defer logrus.Printf("Consumed message from publisher error: %s", err.Error())
		isConsumed = false
	}

	// check publishing data from queue is exist or not before consuming
	if err := channel.Ack(0, true); err != nil {
		defer logrus.Printf("Consumed data into queue error: %s", err.Error())
		isPublished = false
	}

	msg := <-res
	if err := json.Unmarshal(msg.Body, &rbResponse.Items); err != nil {
		defer logrus.Printf("Parse response queue into queueResponse error: %s", err.Error())
		isConsumed = false
	}

	defer logrus.Printf(`RABBITMQ CONSUMER LOGS:
	Exchange Name: %s
	Routing Name: %s
	Event Name: %s
	Response: %v
`, exchangeName, routingKey, eventName, rbResponse.Items)

	return rbResponse, isConsumed
}
