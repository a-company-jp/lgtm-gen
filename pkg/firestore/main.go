package firestore

import (
	"context"
	"fmt"
	"lgtm-gen/pkg/config"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type FireStore struct {
	client *firestore.Client
}

func NewFireStore() (*FireStore, error) {
	conf := config.Get()
	ctx := context.Background()

	var app *firebase.App
	var err error
	if conf.Infrastructure.GoogleCloud.UseCredentialsFile {
		app, err = firebase.NewApp(ctx, nil, option.WithCredentialsFile(conf.Infrastructure.GoogleCloud.CredentialsFilePath))
		if err != nil {
			return nil, fmt.Errorf("failed to create Firestore app with credentials file: %w", err)
		} else {
			app, err = firebase.NewApp(ctx, nil)
			if err != nil {
				return nil, fmt.Errorf("failed to create Firestore app: %w", err)
			}
		}
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create Firestore client: %w", err)
	}

	return &FireStore{
		client: client,
	}, nil
}
