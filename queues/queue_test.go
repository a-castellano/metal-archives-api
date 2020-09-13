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
	"strings"
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
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel in TestSendDie")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"incoming", // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue in TestSendDie")

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

	jobManagementError := startJobManagement(queueConfig, client)
	if jobManagementError != nil {
		t.Errorf("startJobManagement should return no errors when die is processed.")
	}

}

func TestSendNoArtistsFound(t *testing.T) {

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

	var infoRetrieval commontypes.InfoRetrieval
	var job commontypes.Job

	infoRetrieval.Type = commontypes.ArtistName
	infoRetrieval.Artist = "AnyArtist"

	retrievalData, _ := commontypes.EncodeInfoRetrieval(infoRetrieval)

	job.Data = retrievalData
	job.ID = 0
	job.Status = true
	job.Finished = false
	job.Type = 1

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
	failOnError(err, "Failed to connect to RabbitMQ in TestSendNoArtistsFound.")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel in TestSendNoArtistsFound.")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"incoming", // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare incoming queue in TestSendNoArtistsFound.")

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

	failOnError(err, "Failed to send first job in TestSendNoArtistsFound.")
	var dieJob commontypes.Job

	dieJob.ID = 0
	dieJob.Status = true
	dieJob.Finished = false
	dieJob.Type = commontypes.Die

	encodedDieJob, _ := commontypes.EncodeJob(dieJob)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         encodedDieJob,
		})

	failOnError(err, "Failed to send die job in TestSendNoArtistsFound.")

	jobManagementError := startJobManagement(queueConfig, client)

	if jobManagementError != nil {
		t.Errorf("startJobManagement should return no errors when die is processed.")
	}

	outgoingCh, err := conn.Channel()
	failOnError(err, "Failed to open outgoing channel in TestSendNoArtistsFound.")
	defer outgoingCh.Close()
	//
	//	outgoingQueue, err := outgoingCh.QueueDeclare(
	//		"outgoing",
	//		true,  // durable
	//		false, // delete when unused
	//		false, // exclusive
	//		false, // no-wait
	//		nil,   // arguments
	//	)
	//	failOnError(err, "Failed to declare a outgoing queue in TestSendNoArtistsFound.")
	//
	//	err = outgoingCh.Qos(
	//		1,     // prefetch count
	//		0,     // prefetch size
	//		false, // global
	//	)
	//	failOnError(err, "Failed to set QoS outgoingCh in TestSendNoArtistsFound.")

	msgs, err := outgoingCh.Consume(
		"outgoing", // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	forever := make(chan bool)
	var receivedData []byte

	go func() {
		for d := range msgs {

			receivedData = d.Body
			d.Ack(false)
			forever <- false
		}
	}()

	<-forever
	decodedJob, decodedJobErr := commontypes.DecodeJob(receivedData)

	if decodedJob.Type != commontypes.ArtistInfoRetrieval {
		t.Errorf("Decoed job type should be ArtistInfoRetrieval in TestSendNoArtistsFound.")
	}
	if decodedJobErr != nil {
		t.Errorf("DecodeJob should return no errors.")
	}
	_, decodedResultErr := commontypes.DecodeArtistInfo(decodedJob.Result)

	if decodedResultErr != nil {
		if !strings.HasPrefix(decodedResultErr.Error(), "Artist r_etrieval failed:") {
			t.Errorf("Message with failed data should return 'Empty data received.' error, not '%s'.", decodedResultErr.Error())
		}
	}
}
