package blog

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type EventBus struct {
	Channel *amqp.Channel
}

func NewAMQP() (*amqp.Channel, error) {
	dsn := viper.GetString("rmq_dsn")
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Error("could not connect to the dsn")
		return &amqp.Channel{}, err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Error("could not create channel: ", err)
		return &amqp.Channel{}, err
	}

	_, err = ch.QueueDeclare("blog", true, false, false, false, nil)
	if err != nil {
		log.Error("could not declare queue: ", err)
		return &amqp.Channel{}, err
	}

	return ch, nil

}
