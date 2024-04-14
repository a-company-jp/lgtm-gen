package response

import (
	"lgtm-gen/svc/pkg/domain/model"
	"time"
)

type CreateLGTMResponse struct {
	ImageUrl string `json:"imageUrl"`
}

type LGTM struct {
	ID        string    `json:"id"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewGetListResponse(lgtms []model.LGTM) []LGTM {
	var lgtmsView = make([]LGTM, len(lgtms))
	for i, lgtm := range lgtms {
		lgtmsView[len(lgtms)-i] = LGTM{
			ID:        lgtm.ID,
			Url:       lgtm.Url,
			CreatedAt: lgtm.CreatedAt,
		}
	}
	return lgtmsView
}
