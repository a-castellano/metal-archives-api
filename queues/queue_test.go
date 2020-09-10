// +build integration_tests

package queues

import (
	"bytes"
	commontypes "github.com/a-castellano/music-manager-common-types/types"
	config "github.com/a-castellano/music-manager-metal-archives-wrapper/config_reader"
	"github.com/streadway/amqp"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

type RoundTripperMock struct {
	Response *http.Response
	RespErr  error
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (rtm *RoundTripperMock) RoundTrip(*http.Request) (*http.Response, error) {
	return rtm.Response, rtm.RespErr
}

func TestSendDie(t *testing.T) {

	var queueConfig config.Config

	queueConfig.Server.Host = "127.0.0.1"
	queueConfig.Server.Port = 5672
	queueConfig.Server.User = "guest"
	queueConfig.Server.Password = "guest"

	queueConfig.Incoming.Name = "incoming"
	queueConfig.Incoming.Durable = true
	queueConfig.Incoming.DeleteWhenUnused = false
	queueConfig.Incoming.Exclusive = false
	queueConfig.Incoming.NoWait = false
	queueConfig.Incoming.NoLocal = false
	queueConfig.Incoming.AutoACK = false

	queueConfig.Outgoing.Name = "outgoing"
	queueConfig.Outgoing.Durable = true
	queueConfig.Outgoing.DeleteWhenUnused = false
	queueConfig.Outgoing.Exclusive = false
	queueConfig.Outgoing.NoWait = false
	queueConfig.Outgoing.NoLocal = false
	queueConfig.Outgoing.AutoACK = true

	var job commontypes.Job

	job.ID = 0
	job.Status = true
	job.Finished = false
	job.Type = commontypes.Die

	encodedJob, _ := commontypes.EncodeJob(job)

	client := http.Client{Transport: &RoundTripperMock{Response: &http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`
{
	"error": "",
	"iTotalRecords": 0,
	"iTotalDisplayRecords": 0,
	"sEcho": 0,
	"aaData": [
		]
}
	`))}}}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "TEST PRE Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"incoming", // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         encodedJob,
		})

	startJobManagement(queueConfig, client)

}
