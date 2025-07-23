package detection

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type ErrorBased struct{}

func (e *ErrorBased) Detect(target string, ctx context.Context) ([]common.Finding, error) {
	// In a real implementation, we would send a request with an error-based
	// payload and check for a database error in the response.
	return nil, nil
}
