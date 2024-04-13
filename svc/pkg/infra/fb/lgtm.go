package infra

import (
	"context"
	"errors"
	"fmt"
	"lgtm-gen/pkg/fb"
	"lgtm-gen/svc/pkg/domain/model"
	"time"
)

const LGTMFolderName = "LGTMs"

type LGTMTable struct {
	f     *fb.Firebase
	tokyo *time.Location
}

func NewLGTMTable(f *fb.Firebase) *LGTMTable {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	return &LGTMTable{
		f:     f,
		tokyo: loc,
	}
}

// List get list of lgtm images data
func (l LGTMTable) List() ([]model.LGTM, error) {
	ctx := context.Background()
	// TODO: pagination
	qs, err := l.f.Client.NewRef(LGTMFolderName).OrderByChild("createdAt").
		LimitToLast(20).GetOrdered(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query result: %w", err)
	}

	results := make([]model.LGTM, 0, len(qs))
	for _, doc := range qs {
		var lgtm = model.LGTM{}
		if err := doc.Unmarshal(&lgtm); err != nil {
			return nil, fmt.Errorf("failed to unmarshal some result: %w", err)
		}
		lgtm.ID = doc.Key()
		results = append(results, lgtm)
	}
	return results, nil
}

func (l LGTMTable) Set(target model.LGTM) error {
	if target.ID == "" {
		return errors.New("id is required for model.LGTM")
	}
	ctx := context.Background()
	if err := l.f.Client.NewRef(LGTMFolderName).Child(target.ID).Set(ctx, target); err != nil {
		return fmt.Errorf("failed to set LGTM data, target: %v, err: %w", target, err)
	}
	return nil
}

func (l LGTMTable) Create(target model.LGTM) error {
	if !target.CreatedAt.IsZero() {
		return errors.New("CreatedAt should be Zero value")
	}
	target.CreatedAt = time.Now().In(l.tokyo)
	return l.Set(target)
}
