package natsx

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

func Connect(url string) (*nats.Conn, error) {
	return nats.Connect(url)
}

func Publish(conn *nats.Conn, subject string, data any) error {

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return conn.Publish(subject, payload)
}

func Subscribe[T any](
	conn *nats.Conn,
	subject string,
	handler func(T),
) error {

	_, err := conn.Subscribe(subject, func(msg *nats.Msg) {

		var payload T

		err := json.Unmarshal(msg.Data, &payload)
		if err != nil {
			return
		}

		handler(payload)
	})

	return err
}