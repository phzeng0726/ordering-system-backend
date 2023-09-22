package auth

import (
	"context"
	"fmt"
	"ordering-system-backend/internal/domain"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

func Init() (*auth.Client, error) {
	opt := option.WithCredentialsFile("firebase_credential.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	// Access Auth service from default app
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CreateUser(uq domain.UserRequest, uid string, client *auth.Client) error {
	fullName := uq.LastName + " " + uq.FirstName
	params := (&auth.UserToCreate{}).
		Email(uq.Email).
		Password(uq.Password).
		DisplayName(fullName).
		UID(uid)

	u, err := client.CreateUser(context.Background(), params)

	if err != nil {
		return err
	}

	fmt.Printf("Successfully created user: %v\n", u)
	return nil
}

func DeleteUser(uid string, client *auth.Client) error {
	err := client.DeleteUser(context.Background(), uid)

	if err != nil {
		return err
	}

	fmt.Printf("Successfully deleted user: %s\n", uid)
	return nil
}
