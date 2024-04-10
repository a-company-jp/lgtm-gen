package fs

import (
	"context"
	"fmt"
	"lgtm-gen/pkg/config"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type Firestore struct {
	Client *firestore.Client
}

func NewFirestore() (*Firestore, error) {
	conf := config.Get()
	ctx := context.Background()

	var client *firestore.Client
	var err error
	if conf.Infrastructure.GoogleCloud.UseCredentialsFile {
		client, err = firestore.NewClient(ctx, conf.Infrastructure.GoogleCloud.ProjectID, option.WithCredentialsFile(conf.Infrastructure.GoogleCloud.CredentialsFilePath))
		if err != nil {
			return nil, fmt.Errorf("failed to create Firestore client with credentials file: %w", err)
		}
	} else {
		client, err = firestore.NewClient(ctx, conf.Infrastructure.GoogleCloud.ProjectID)
		if err != nil {
			return nil, fmt.Errorf("failed to create Firestore client: %w", err)
		}
	}

	return &Firestore{
		Client: client,
	}, nil
}
