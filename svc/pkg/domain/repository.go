package domain

import (
	"lgtm-gen/svc/pkg/domain/model"
)

type ILGTMTableRepository interface {
	List() ([]model.LGTM, error)
	Create(id string, url string) error
}

type ILGTMBucketRepository interface {
	Create(objectName string, data []byte) error
}
