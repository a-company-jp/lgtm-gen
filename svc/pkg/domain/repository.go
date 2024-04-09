package domain

import (
	"lgtm-gen/svc/pkg/domain/model"
)

type ILGTMRepository interface {
	List() ([]*model.LGTM, error)
	Create(id string) error
}
