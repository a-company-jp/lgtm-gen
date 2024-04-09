package fs

import (
	"context"
	"fmt"
	"lgtm-gen/pkg/config"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type FireStore struct {
	Client *firestore.Client
}

func NewFireStore() (*FireStore, error) {
	conf := config.Get()
	ctx := context.Background()

	var client *firestore.Client
	var err error
	if conf.Infrastructure.GoogleCloud.UseCredentialsFile {
		client, err = firestore.NewClient(ctx, conf.Infrastructure.GoogleCloud.ProjectID, option.WithCredentialsFile(conf.Infrastructure.GoogleCloud.CredentialsFilePath))
		if err != nil {
			return nil, fmt.Errorf("failed to create FireStore client with credentials file: %w", err)
		}
	} else {
		client, err = firestore.NewClient(ctx, conf.Infrastructure.GoogleCloud.ProjectID)
		if err != nil {
			return nil, fmt.Errorf("failed to create FireStore client: %w", err)
		}
	}

	return &FireStore{
		Client: client,
	}, nil
}
