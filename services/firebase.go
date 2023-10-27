package services

import (
	"context"
	"log"

	"firebase.google.com/go/messaging"
	"github.com/hosseinmirzapur/sdb/config"
)

func NotifyViaFirebase(regTok, title, body string) (string, error) {

	app, _, _ := config.SetupFirebase()

	ctx := context.Background()

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting messaging client: %v\n", err)
	}

	return client.Send(ctx, &messaging.Message{
		Token: regTok,
		Data: map[string]string{
			"title": title,
			"body":  body,
		},
	})

}
