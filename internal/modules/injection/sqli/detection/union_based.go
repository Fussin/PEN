package detection

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type UnionBased struct{}

func (u *UnionBased) Detect(target string, ctx context.Context) ([]common.Finding, error) {
	// In a real implementation, we would send requests with union-based
	// payloads and check for the injected data in the response.
	return nil, nil
}
