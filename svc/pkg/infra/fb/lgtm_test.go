package infra

import (
	"lgtm-gen/pkg/testutil"
	"lgtm-gen/svc/pkg/domain/model"
	"testing"
	"time"
)

func TestLGTMTable_Set(t *testing.T) {
	fb := testutil.NewFirebaseTest()
	defer fb.Reset()
	client := NewLGTMTable(fb.GetClient())
	testdata := []model.LGTM{
		{
			ID:        "1",
			Url:       "https://example.com/img/1",
			CreatedAt: time.Now().Add(-2 * time.Hour),
		},
		{
			ID:        "2",
			Url:       "https://example.com/img/2",
			CreatedAt: time.Now().Add(-time.Hour),
		},
		{
			ID:        "3",
			Url:       "https://example.com/img/3",
			CreatedAt: time.Now().Add(-30 * time.Minute),
		},
	}
	for _, d := range testdata {
		if err := client.Set(d); err != nil {
			t.Fatalf("failed to set data, err: %v", err)
		}
	}
	results, err := client.List()
	if err != nil {
		t.Fatalf("failed to list data, err: %v", err)
	}
	if len(results) != len(testdata) {
		t.Fatalf("len(results) got: %d, want: %d", len(results), len(testdata))
	}
	for i, d := range testdata {
		if d.ID != results[i].ID {
			t.Fatalf("results[%d].ID got: %s, want: %s", i, results[i].ID, d.ID)
		}
		if d.Url != results[i].Url {
			t.Fatalf("results[%d].Url got: %s, want: %s", i, results[i].Url, d.Url)
		}
		if d.CreatedAt != results[i].CreatedAt {
			t.Fatalf("results[%d].CreatedAt got: %v, want: %v", i, results[i].CreatedAt, d.CreatedAt)
		}
	}
}
