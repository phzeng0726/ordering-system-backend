package notification

import (
	"context"
	"fmt"
	"ordering-system-backend/internal/config"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func Init() (*messaging.Client, error) {
	var projectPath string

	if config.Env.IsOnCloud {
		projectPath = ""
	} else {
		projectPath = ".."
	}

	filePath := fmt.Sprintf(projectPath + "/assets/firebase_credential.json")
	opt := option.WithCredentialsFile(filePath)

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

func SendPushNotification(client *messaging.Client, deviceTokens []string, storeId string) error {
	response, err := client.SendMulticast(context.Background(), &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: "NEW_ORDER_TICKET",
			Body:  storeId,
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
