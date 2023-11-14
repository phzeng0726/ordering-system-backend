package notification

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func Init() (*messaging.Client, error) {
	opt := option.WithCredentialsFile("firebase_credential.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	// Access Auth service from default app
	client, err := app.Messaging(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func SendPushNotification(client *messaging.Client, deviceTokens []string) error {
	response, err := client.SendMulticast(context.Background(), &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: "Congratulations!!",
			Body:  "You have just implement push notification",
		},
		Tokens: deviceTokens,
	})

	if err != nil {
		return err
	}

	fmt.Println("Response success count : ", response.SuccessCount)
	fmt.Println("Response failure count : ", response.FailureCount)
	return nil
}
