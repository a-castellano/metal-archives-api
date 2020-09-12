package queues

import (
	"fmt"
	config "github.com/a-castellano/music-manager-metal-archives-wrapper/config_reader"
	"github.com/a-castellano/music-manager-metal-archives-wrapper/jobs"
	"github.com/streadway/amqp"
	"net/http"
	"strconv"
)

func startJobManagement(config config.Config, client http.Client) error {

	connection_string := "amqp://" + config.Server.User + ":" + config.Server.Password + "@" + config.Server.Host + ":" + strconv.Itoa(config.Server.Port) + "/"
	conn, err := amqp.Dial(connection_string)

	defer conn.Close()

	incoming_ch, err := conn.Channel()
	defer incoming_ch.Close()

	if err != nil {
		return fmt.Errorf("Failed to open incoming channel: %w", err)
	}

	outgoing_ch, err := conn.Channel()
	defer outgoing_ch.Close()

	if err != nil {
		return fmt.Errorf("Failed to open outgoing channel: %w", err)
	}

	incoming_q, err := incoming_ch.QueueDeclare(
		config.Incoming.Name,
		config.Incoming.Durable,
		config.Incoming.DeleteWhenUnused,
		config.Incoming.Exclusive,
		config.Incoming.NoWait,
		nil, // arguments
	)

	if err != nil {
		return fmt.Errorf("Failed to declare incoming queue: %w", err)
	}

	outgoing_q, err := outgoing_ch.QueueDeclare(
		config.Outgoing.Name,
		config.Outgoing.Durable,
		config.Outgoing.DeleteWhenUnused,
		config.Outgoing.Exclusive,
		config.Outgoing.NoWait,
		nil, // arguments
	)

	if err != nil {
		return fmt.Errorf("Failed to declare outgoing queue: %w", err)
	}

	err = incoming_ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	if err != nil {
		return fmt.Errorf("Failed to set incoming QoS: %w", err)
	}

	jobsToProcess, err := incoming_ch.Consume(
		incoming_q.Name,
		"",                        // consumer
		config.Incoming.AutoACK,   // auto-ack
		config.Incoming.Exclusive, // exclusive
		config.Incoming.NoLocal,   // no-local
		config.Incoming.NoWait,    // no-wait
		nil,                       // args
	)

	if err != nil {
		return fmt.Errorf("Failed to register a consumer: %w", err)
	}

	processJobs := make(chan bool)

	go func() {
		for job := range jobsToProcess {

			fmt.Println("FUNC JOB")
			die, jobResult, _ := jobs.ProcessJob(job.Body, client)

			if die {
				job.Ack(false)
				processJobs <- false
				return
			}
			err = outgoing_ch.Publish(
				"",              // exchange
				outgoing_q.Name, // routing key
				false,           // mandatory
				false,
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         jobResult,
				})
			if err != nil {
				//return fmt.Errorf("Failed to send job result: %w", err)
				return
			}

			job.Ack(false)
		}
		return
	}()

	<-processJobs

	return nil
}
