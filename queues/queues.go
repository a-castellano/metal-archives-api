package queues

import (
	"fmt"
	config "github.com/a-castellano/metal-archives-wrapper/config_reader"
	"github.com/streadway/amqp"
	"strconv"
)

func startJobManagement(config config.Config) error {

	connection_string := "amqp://" + config.Server.User + ":" + config.Server.Password + "@" + config.Server.Host + ":" + strconv.Itoa(config.Server.Port) + "/"
	conn, err := amqp.Dial(connection_string)

	defer conn.Close()

	incoming_ch, err := conn.Channel()
	defer incoming_ch.Close()

	if err != nil {
		return fmt.Errorf("Failed to open incoming channel: %w", err)
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

	err = incoming_ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	if err != nil {
		return fmt.Errorf("Failed to set incoming QoS: %w", err)
	}

	//msgs, err := incoming_ch.Consume(
	incoming_ch.Consume(
		incoming_q.Name,
		"",                        // consumer
		config.Incoming.AutoACK,   // auto-ack
		config.Incoming.Exclusive, // exclusive
		//		config.Incoming.NoLocal,   // no-local
		false,
		config.Incoming.NoWait, // no-wait
		nil,                    // args
	)

	if err != nil {
		return fmt.Errorf("Failed to register a consumer: %w", err)
	}

	return nil
}
