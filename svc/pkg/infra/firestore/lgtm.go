package infra

import (
	"context"
	"lgtm-gen/pkg/fs"
	"lgtm-gen/svc/pkg/domain/model"
	"time"
)

const LGTMCollectionName = "lgtms"

type LGTMTable struct {
	f *fs.Firestore
}

func NewLGTMTable(f *fs.Firestore) *LGTMTable {
	return &LGTMTable{
		f: f,
	}
}

// List get list of lgtm images data
func (l LGTMTable) List() ([]model.LGTM, error) {
	ctx := context.Background()
	// TODO: pagination
	docs, err := l.f.Client.Collection(LGTMCollectionName).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	lgtms := make([]model.LGTM, len(docs))
	for i, doc := range docs {
		var lgtm = model.LGTM{}
		if err := doc.DataTo(&lgtm); err != nil {
			return nil, err
		}
		lgtm.ID = doc.Ref.ID
		lgtms[i] = lgtm
	}

	return lgtms, nil
}

// Create add item to firebase
func (l LGTMTable) Create(id string, url string) error {
	ctx := context.Background()
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}
	_, err = l.f.Client.Collection(LGTMCollectionName).Doc(id).Set(ctx, map[string]interface{}{
		"createdAt": time.Now().In(loc),
		"url":       url,
	})
	if err != nil {
		return err
	}
	return nil
}
