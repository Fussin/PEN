package detection

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/autonomouspen/scanner/internal/common"
)

type BooleanBased struct{}

func (b *BooleanBased) Detect(target string, ctx context.Context) ([]common.Finding, error) {
	// In a real implementation, we would send requests with true and false
	// payloads and compare the responses.
	truePayload := "' AND 1=1 --"
	falsePayload := "' AND 1=2 --"

	originalResp, err := http.Get(target)
	if err != nil {
		return nil, err
	}
	originalBody, _ := ioutil.ReadAll(originalResp.Body)

	trueResp, err := http.Get(target + truePayload)
	if err != nil {
		return nil, err
	}
	trueBody, _ := ioutil.ReadAll(trueResp.Body)

	falseResp, err := http.Get(target + falsePayload)
	if err != nil {
		return nil, err
	}
	falseBody, _ := ioutil.ReadAll(falseResp.Body)

	if !strings.EqualFold(string(originalBody), string(trueBody)) && strings.EqualFold(string(trueBody), string(falseBody)) {
		// This is a potential boolean-based SQLi
	}

	return nil, nil
}
