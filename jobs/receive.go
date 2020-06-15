package jobs

import (
	"github.com/streadway/amqp"
)

func reciveJob() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
}
