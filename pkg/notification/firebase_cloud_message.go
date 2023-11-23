package notification

import (
	"context"
	"errors"
	"fmt"
	"ordering-system-backend/internal/config"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func Init(userType int) (*messaging.Client, error) {
	var projectPath string
	var filePath string

	if config.Env.IsOnCloud {
		projectPath = ""
	} else {
		projectPath = ".."
	}

	if userType == 0 {
		filePath = fmt.Sprintf(projectPath + "/assets/firebase_credential.json")
	} else if userType == 1 {
		filePath = fmt.Sprintf(projectPath + "/assets/client_firebase_credential.json")
	} else {
		return nil, errors.New("userType not available for firebase init method")
	}

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

func SendNotificationToStore(client *messaging.Client, deviceTokens []string, storeId string) error {
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

func SendNotificationToClient(client *messaging.Client, deviceTokens []string) error {
	response, err := client.SendMulticast(context.Background(), &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: "NEW_ORDER_TICKET_STATUS",
			Body:  "Update order ticket status",
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
