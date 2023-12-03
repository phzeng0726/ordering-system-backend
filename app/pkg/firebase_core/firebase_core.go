package firebase_core

import (
	"context"
	"errors"
	"fmt"

	"ordering-system-backend/internal/config"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func Init(userType int, serviceType string) (interface{}, error) {
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

	switch serviceType {
	case "auth":
		authClient, err := app.Auth(context.Background())
		if err != nil {
			return nil, err
		}
		return authClient, nil
	case "messaging":
		messagingClient, err := app.Messaging(context.Background())
		if err != nil {
			return nil, err
		}
		return messagingClient, nil
	default:
		return nil, errors.New("serviceType not available for firebase init method")
	}
}
