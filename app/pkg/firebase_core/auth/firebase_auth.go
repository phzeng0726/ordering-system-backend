package auth

import (
	"context"
	"fmt"
	firebase "ordering-system-backend/pkg/firebase_core"

	"firebase.google.com/go/auth"
)

func Init(userType int) (*auth.Client, error) {
	client, err := firebase.Init(userType, "auth")
	if err != nil {
		return nil, err
	}

	return client.(*auth.Client), nil
}

func CreateUser(client *auth.Client, uid string, email string, password string) error {
	params := (&auth.UserToCreate{}).
		UID(uid).
		Email(email).
		Password(password).
		Disabled(false)

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

	fmt.Printf("Successfully updated user password: %v\n", u)
	return nil
}

func DeleteUser(uid string, client *auth.Client) error {
	if err := client.DeleteUser(context.Background(), uid); err != nil {
		return err
	}

	return nil
}
