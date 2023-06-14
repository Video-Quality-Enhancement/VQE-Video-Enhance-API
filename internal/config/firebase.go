package config

import (
	"context"
	"os"

	firebase "firebase.google.com/go/v4"
	"golang.org/x/exp/slog"
	"google.golang.org/api/option"
)

type FirebaseClient interface {
	VerifyIDToken(ctx context.Context, idToken string) (string, error)
}

type firebaseClient struct {
	app *firebase.App
}

func NewFirebaseClient() FirebaseClient {

	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_SA_KEY_PATH"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		slog.Error("error initializing app: %v\n", err)
	}

	return &firebaseClient{app}
}

func (c *firebaseClient) VerifyIDToken(ctx context.Context, idToken string) (string, error) {

	client, err := c.app.Auth(ctx)
	if err != nil {
		slog.Error("error getting Auth client", "error", err)
		return "", err
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		slog.Error("error verifying ID token", "error", err)
		return "", err
	}

	return token.UID, nil

}
