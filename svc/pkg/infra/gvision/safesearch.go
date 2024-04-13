package gvision

import (
	"context"
	"lgtm-gen/pkg/gvision"

	"cloud.google.com/go/vision/v2/apiv1/visionpb"
)

type SafeSearch struct {
	gv *gvision.GVision
}

func NewSafeSearch(gv *gvision.GVision) *SafeSearch {
	return &SafeSearch{
		gv: gv,
	}
}

func (s SafeSearch) Detect(data []byte) ([]string, error) {
	ctx := context.Background()

	image := &visionpb.Image{
		Content: data,
	}

	annotations, err := s.gv.Client.DetectSafeSearch(ctx, image, nil)
	if err != nil {
		return nil, err
	}

	var labels []string
	if annotations.Adult == visionpb.Likelihood_VERY_LIKELY || annotations.Adult == visionpb.Likelihood_LIKELY {
		labels = append(labels, "adult")
	}
	if annotations.Spoof == visionpb.Likelihood_VERY_LIKELY || annotations.Spoof == visionpb.Likelihood_LIKELY {
		labels = append(labels, "spoof")
	}
	if annotations.Medical == visionpb.Likelihood_VERY_LIKELY || annotations.Medical == visionpb.Likelihood_LIKELY {
		labels = append(labels, "medical")
	}
	if annotations.Violence == visionpb.Likelihood_VERY_LIKELY || annotations.Violence == visionpb.Likelihood_LIKELY {
		labels = append(labels, "violence")
	}
	if annotations.Racy == visionpb.Likelihood_VERY_LIKELY || annotations.Racy == visionpb.Likelihood_LIKELY {
		labels = append(labels, "racy")
	}

	return labels, nil
}
