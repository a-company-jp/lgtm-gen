package firestore

import (
	"context"
	"lgtm-gen/svc/pkg/infra/entity"

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
func (l *LGTMTable) List() ([]*entity.LGTM, error) {
	ctx := context.Background()
	// TODO: pagination
	docs, err := l.fsClient.Collection(LGTMCollectionName).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	lgtms := make([]*entity.LGTM, len(docs))
	for i, doc := range docs {
		var lgtm = entity.LGTM{}
		if err := doc.DataTo(&lgtm); err != nil {
			return nil, err
		}
		lgtms[i] = &lgtm
	}

	return lgtms, nil
}
