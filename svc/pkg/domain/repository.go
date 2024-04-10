package domain

import (
	"lgtm-gen/svc/pkg/domain/model"
)

type ILGTMTableRepository interface {
	List() ([]model.LGTM, error)
	Create(id string) error
}
