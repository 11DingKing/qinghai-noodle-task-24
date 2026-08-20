package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask24(t *testing.T) {
	now := time.Now()
	s := NewService(NewRegistry(), func() time.Time { return now })
	c := SubsidyClaim{StoreID: "store-1", Period: "2026-08", EligibleSales: 10000, EvidenceHashes: []string{"sha"}, SubmittedAt: now}
	require.NoError(t, s.CheckSubsidyClaim(context.Background(), c, compliantStore(now), activeLicense(now)))
}
