package gvision

import (
	"context"
	"fmt"
	"lgtm-gen/pkg/config"

	vision "cloud.google.com/go/vision/apiv1"
	"google.golang.org/api/option"
)

type GVision struct {
	Client *vision.ImageAnnotatorClient
}

func NewGVision() (*GVision, error) {
	conf := config.Get()
	ctx := context.Background()

	var client *vision.ImageAnnotatorClient
	var err error
	if conf.Infrastructure.GoogleCloud.UseCredentialsFile {
		client, err = vision.NewImageAnnotatorClient(ctx, option.WithCredentialsFile(conf.Infrastructure.GoogleCloud.CredentialsFilePath))
		if err != nil {
			return nil, fmt.Errorf("failed to create GVision client with credentials file: %w", err)
		}
	} else {
		client, err = vision.NewImageAnnotatorClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to create GVision client: %w", err)
		}
	}

	return &GVision{
		Client: client,
	}, nil
}
