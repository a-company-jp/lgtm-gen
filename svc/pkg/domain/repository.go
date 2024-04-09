package domain

import "lgtm-gen/svc/pkg/infra/entity"

type ILGTMRepository interface {
	List() ([]*entity.LGTM, error)
}
