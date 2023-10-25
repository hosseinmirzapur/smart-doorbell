package services

import (
	"context"
	"errors"
	"log"

	"github.com/hosseinmirzapur/sdb/config"
)

func validateEvent(event string) error {
	valid_events := []string{"ENTER", "LEAVE", "RING"}

	for _, valid_event := range valid_events {
		if event == valid_event {
			return nil
		}
	}
	return errors.New("invalid event")

}

func notify(mobile, event string) error {

	app, _, _ := config.SetupFirebase()

	ctx := context.Background()

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	panic(client) // todo: delete this line and implement the real code

	return validateEvent(event)
}
