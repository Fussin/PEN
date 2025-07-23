package detection

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type SecondOrder struct{}

func (s *SecondOrder) Detect(target string, ctx context.Context) ([]common.Finding, error) {
	// In a real implementation, we would need to submit a payload to one
	// endpoint and then visit another endpoint to check for its reflection.
	return nil, nil
}
