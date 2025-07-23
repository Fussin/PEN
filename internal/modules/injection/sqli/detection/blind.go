package detection

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type Blind struct{}

func (b *Blind) Detect(target string, ctx context.Context) ([]common.Finding, error) {
	// In a real implementation, we would use a combination of boolean-based
	// and time-based techniques to exfiltrate data.
	return nil, nil
}
