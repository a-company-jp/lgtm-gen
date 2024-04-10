package infra

import (
	"context"
	"lgtm-gen/pkg/gcs"
)

type LGTMBucket struct {
	g *gcs.GCS
}

func NewLGTMBucket(g *gcs.GCS) *LGTMBucket {
	return &LGTMBucket{
		g: g,
	}
}

func (l LGTMBucket) Create(objectName string, data []byte) error {
	ctx := context.Background()
	err := l.g.Upload(ctx, objectName, data)
	if err != nil {
		return err
	}
	return nil
}
