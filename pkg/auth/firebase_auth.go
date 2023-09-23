package auth

import (
	"context"
	"fmt"

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

func CreateUser(client *auth.Client, uid string, email string, password string) error {
	params := (&auth.UserToCreate{}).
		UID(uid).
		Email(email).
		Password(password)

	_, err := client.CreateUser(context.Background(), params)

	if err != nil {
		return err
	}

	return nil
}

func ResetPassword(client *auth.Client, uid string, newPassword string) error {
	params := (&auth.UserToUpdate{}).
		Password(newPassword)

	u, err := client.UpdateUser(context.Background(), uid, params)

	if err != nil {
		return err
	}

	fmt.Printf("Successfully created user: %v\n", u)
	return nil
}

func DeleteUser(uid string, client *auth.Client) error {
	if err := client.DeleteUser(context.Background(), uid); err != nil {
		return err
	}

	return nil
}
