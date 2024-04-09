package firestore

import (
	"context"
	"lgtm-gen/svc/pkg/domain/model"
	"time"

	"cloud.google.com/go/firestore"
)

const LGTMCollectionName = "lgtms"

type LGTMTable struct {
	fsClient *firestore.Client
}

func NewLGTMTable() *LGTMTable {
	return &LGTMTable{}
}

// List get list of lgtm images data
func (l *LGTMTable) List() ([]*model.LGTM, error) {
	ctx := context.Background()
	// TODO: pagination
	docs, err := l.fsClient.Collection(LGTMCollectionName).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	lgtms := make([]*model.LGTM, len(docs))
	for i, doc := range docs {
		var lgtm = model.LGTM{}
		lgtm.ID = doc.Ref.ID
		if err := doc.DataTo(&lgtm); err != nil {
			return nil, err
		}
		lgtms[i] = &lgtm
	}

	return lgtms, nil
}

// Create add item to firebase
func (l *LGTMTable) Create(id string, title string) error {
	ctx := context.Background()
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}
	_, err = l.fsClient.Collection(LGTMCollectionName).Doc(id).Set(ctx, map[string]interface{}{
		"title":     title,
		"createdAt": time.Now().In(loc),
	})
	if err != nil {
		return err
	}
	return nil
}
