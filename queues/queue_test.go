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

	jobManagementError := StartJobManagement(queueConfig, client)
	if jobManagementError != nil {
		t.Errorf("StartJobManagement should return no errors when die is processed.")
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

	jobManagementError := StartJobManagement(queueConfig, client)

	if jobManagementError != nil {
		t.Errorf("StartJobManagement should return no errors when die is processed.")
	}

	outgoingCh, err := conn.Channel()
	failOnError(err, "Failed to open outgoing channel in TestSendNoArtistsFound.")
	defer outgoingCh.Close()

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
		t.Errorf("Decoded job type should be ArtistInfoRetrieval in TestSendNoArtistsFound. It's %d.", decodedJob.Type)
	}
	if decodedJobErr != nil {
		t.Errorf("DecodeJob should return no errors.")
	}

	if decodedJob.Error != "Artist retrieval failed: No artist was found." {
		t.Errorf("DecodeJob error should be 'Artist retrieval failed: No artist was found.', not '%s'.", decodedJob.Error)
	}
}

func TestSendArtistsFound(t *testing.T) {

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
	infoRetrieval.Artist = "Hypocrisy"

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
	"iTotalRecords": 5,
	"iTotalDisplayRecords": 5,
	"sEcho": 0,
	"aaData": [
				[
			"<a href=\"https://www.metal-archives.com/bands/Hypocrisy/96\">Hypocrisy</a>  <!-- 10.740315 -->" ,
			"Death Metal (early), Melodic Death Metal (later)" ,
			"Sweden"     		]
				,
						[
			"<a href=\"https://www.metal-archives.com/bands/Hypocrisy/56165\">Hypocrisy</a>  <!-- 10.740315 -->" ,
			"Power/Thrash Metal" ,
			"United States"     		]
				,
						[
			"<a href=\"https://www.metal-archives.com/bands/Sermon_of_Hypocrisy/7033\">Sermon of Hypocrisy</a>  <!-- 5.3701577 -->" ,
			"Black Metal" ,
			"United Kingdom"     		]
				,
						[
			"<a href=\"https://www.metal-archives.com/bands/The_Polo_Hypocrisy/47897\">The Polo Hypocrisy</a> (<strong>a.k.a.</strong> T.P.H.) <!-- 5.3701577 -->" ,
			"Melodic Death Metal with Hardcore elements" ,
			"Canada"     		]
				,
						[
			"<a href=\"https://www.metal-archives.com/bands/Torture_of_Hypocrisy/3540316100\">Torture of Hypocrisy</a> (<strong>a.k.a.</strong> ToH) <!-- 5.3701577 -->" ,
			"Technical Thrash Metal" ,
			"Poland"     		]
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

	jobManagementError := StartJobManagement(queueConfig, client)

	if jobManagementError != nil {
		t.Errorf("StartJobManagement should return no errors when die is processed.")
	}

	outgoingCh, err := conn.Channel()
	failOnError(err, "Failed to open outgoing channel in TestSendNoArtistsFound.")
	defer outgoingCh.Close()

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
		t.Errorf("Decoded job type should be ArtistInfoRetrieval in TestSendNoArtistsFound. It's %d.", decodedJob.Type)
	}
	if decodedJobErr != nil {
		t.Errorf("DecodeJob should return no errors.")
	}

	if decodedJob.Error != "" {
		t.Errorf("DecodeJob error should be empty, found '%s'.", decodedJob.Error)
	}

	if len(decodedJob.Result) == 0 {
		t.Errorf("DecodeJob  should have result.")
	}

	retrievedData, _ := commontypes.DecodeArtistInfo(decodedJob.Result)

	if retrievedData.Data.Name != "Hypocrisy" {
		t.Errorf("Retrieved data name should be Hyporcisy, not %s.", retrievedData.Data.Name)
	}

	if retrievedData.Data.Country != "Sweden" {
		t.Errorf("Retrieved data country should be Sweden, not %s.", retrievedData.Data.Country)
	}

	if len(retrievedData.ExtraData) != 1 {
		t.Errorf("Retrieved extradata should contain 1 entry, not %d.", len(retrievedData.ExtraData))
	}

}

func TestFailedConfig(t *testing.T) {

	var queueConfig config.Config

	queueConfig.Server.Host = "127.0.0.1"
	queueConfig.Server.Port = 5672
	queueConfig.Server.User = "guest"
	queueConfig.Server.Password = "nopassword"

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

	jobManagementError := StartJobManagement(queueConfig, client)
	if jobManagementError == nil {
		t.Errorf("StartJobManagement should return an error when credentials are invalid.")
	}

}
