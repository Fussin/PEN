package detection

import (
	"context"
	"github.com/autonomouspen/scanner/internal/common"
)

type Engine interface {
	Detect(target string, ctx context.Context) ([]common.Finding, error)
}
