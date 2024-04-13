package fb

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"fmt"
	"google.golang.org/api/option"
	"lgtm-gen/pkg/config"
	"log"
)

type Firebase struct {
	Client *db.Client
}

func NewFirebase() (*Firebase, error) {
	conf := config.Get()
	ctx := context.Background()

	var client *firebase.App
	var err error

	fconf := &firebase.Config{
		DatabaseURL: conf.Infrastructure.GoogleCloud.Firebase.DatabaseURL,
	}
	if conf.Infrastructure.GoogleCloud.UseCredentialsFile {
		client, err = firebase.NewApp(ctx, fconf)
		if err != nil {
			return nil, fmt.Errorf("failed to create Firestore client without credential: %w", err)
		}
	} else {
		opt := option.WithCredentialsFile(
			conf.Infrastructure.GoogleCloud.CredentialsFilePath,
		)
		client, err = firebase.NewApp(ctx, fconf, opt)
		if err != nil {
			return nil, fmt.Errorf("failed to create Firebase client with credential: %w", err)
		}
	}

	c, err := client.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	return &Firebase{
		Client: c,
	}, nil
}
