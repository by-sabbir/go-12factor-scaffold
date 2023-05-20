package blog

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Article struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Slug   string `json:"slug"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

func (a *Article) Publish(ch *amqp.Channel) (errCh chan error) {
	msg, err := json.Marshal(a)
	if err != nil {
		log.Error("could not encode article: ", err)
		errCh <- err
	}

	var event amqp.Publishing

	event.ContentType = "application/json"
	event.Body = msg

	if err := ch.Publish("", "blog", false, false, event); err != nil {
		log.Error("could not publish msg: ", err)
		errCh <- err
	}

	return nil
}
