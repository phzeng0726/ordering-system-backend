package notification

import (
	"context"
	"fmt"
	firebase "ordering-system-backend/pkg/firebase_core"

	"firebase.google.com/go/messaging"
)

func Init(userType int) (*messaging.Client, error) {
	client, err := firebase.Init(userType, "messaging")
	if err != nil {
		return nil, err
	}

	return client.(*messaging.Client), nil
}

func SendNotification(client *messaging.Client, deviceTokens []string, notification *messaging.Notification) error {
	response, err := client.SendMulticast(context.Background(), &messaging.MulticastMessage{
		Notification: notification,
		Tokens:       deviceTokens,
	})

	if err != nil {
		return err
	}

	fmt.Println("Response success count : ", response.SuccessCount)
	fmt.Println("Response failure count : ", response.FailureCount)
	return nil
}
