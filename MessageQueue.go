package MFTCommon

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

const DownloadedTopic = "downloaded"
const URLQueueTopic = "url"
const DeleteTopic = "delete"
const FlashImages = "flashimages"
const MEImages = "meimages"
const BiosImages = "biosimages"
const ExtractedQeueTopic = "extracted"

type eventFunc func(payload string) error

type MessageBundle struct {
	URLQueue         MessageQueue
	DownloadedQueue  MessageQueue
	ExtractedQueue   MessageQueue
	DeleteQueue      MessageQueue
	FlashImagesQueue MessageQueue
	BiosImagesQueue  MessageQueue
	MEImagesQueue    MessageQueue
	TestQueue        MessageQueue
}

type MessageQueue struct {
	Connection *amqp.Channel
	name       string
	topic      string
	log        *logrus.Logger
}

func MQConnect(log *logrus.Logger, serviceName string) MessageBundle {
	conn, err := amqp.Dial("amqp://***REMOVED***:***REMOVED***@localhost:5672/")
	if err != nil {
		log.WithError(err).Panicf("Could not connect to rabbitMQ: %v", err)
	}

	return MessageBundle{
		URLQueue:         buildMessageQueue(log, conn, serviceName, URLQueueTopic),
		DownloadedQueue:  buildMessageQueue(log, conn, serviceName, DownloadedTopic),
		DeleteQueue:      buildMessageQueue(log, conn, serviceName, DeleteTopic),
		FlashImagesQueue: buildMessageQueue(log, conn, serviceName, FlashImages),
		BiosImagesQueue:  buildMessageQueue(log, conn, serviceName, BiosImages),
		MEImagesQueue:    buildMessageQueue(log, conn, serviceName, MEImages),
		ExtractedQueue:   buildMessageQueue(log, conn, serviceName, ExtractedQeueTopic),
		TestQueue:        buildMessageQueue(log, conn, serviceName, "TESTQUEUE"),
	}
}

func buildMessageQueue(log *logrus.Logger, connection *amqp.Connection, serviceName string, queueName string) MessageQueue {

	ch, err := connection.Channel()

	if err != nil {
		log.WithField("MessageQueue", queueName).WithError(err).Error("Could not open Channel")
		return MessageQueue{}
	}

	err = ch.ExchangeDeclare(
		queueName,   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	if err != nil {
		log.WithField("MessageQueue", queueName).WithError(err).Error("Could not declare exchange")
		return MessageQueue{}
	}


	mq := MessageQueue{
		Connection: ch,
		name:       serviceName,
		topic:      queueName,
		log:        log,
	}

	return mq
}

func (mq MessageQueue) MarshalAndSend(data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		logrus.WithField("MessageQueue", mq.topic).WithError(err).Error("Could not marshall json")
		return err
	}

	return mq.Send(string(bytes))
}

func (mq MessageQueue) Send(data string) error {
	mq.log.Debugf("sending '%s' to %s", data, mq.topic)

	err := mq.Connection.Publish(
		mq.topic, // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		})
	if err != nil {
		logrus.WithField("MessageQueue", mq.topic).WithError(err).Errorf("Could not sent to amqp: %v\n", err)
	}
	return err
}

func (mq MessageQueue) RegisterCallback(consumerName string, callback eventFunc) {

	q, err := mq.Connection.QueueDeclare(
		mq.topic+"->"+mq.name,    // name
		true, // durable
		false, // delete when unused
		false,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		mq.log.WithField("MessageQueue", mq.topic).WithError(err).Error("Could not declare queue")
		return
	}

	err = mq.Connection.QueueBind(
		q.Name, // queue name
		"",     // routing key
		mq.topic, // exchange
		false,
		nil)
	if err != nil {
		mq.log.WithField("MessageQueue", mq.topic).WithError(err).Error("Could not bind queue to exchange")
		return
	}

	mq.log.WithField("MessageQueue", mq.topic).Infof("Created Queue: %v", q.Name)


	msgs, err := mq.Connection.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logrus.WithField("MessageQueue", mq.topic).WithError(err).Errorf("Could not create consumer for %s", mq.topic)
	}

	go func() {
		for d := range msgs {
			err := callback(string(d.Body))
			if(err != nil){
				mq.log.Errorf("Callback failed: %v", err)
				d.Acknowledger.Nack(d.DeliveryTag, false, true)
			}
			d.Acknowledger.Ack(d.DeliveryTag, false)
		}
	}()

	logrus.WithField("MessageQueue", mq.topic).Debug("Callback registered")
}
