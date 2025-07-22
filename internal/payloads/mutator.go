package payloads

import (
	"html"
	"net/url"
	"strings"

	"github.com/autonomouspen/scanner/internal/common"
)

type Mutator struct{}

func NewMutator() *Mutator {
	return &Mutator{}
}

func (m *Mutator) Mutate(payload common.Payload) []common.Payload {
	var mutations []common.Payload
	mutations = append(mutations, payload)
	mutations = append(mutations, common.Payload{Value: url.QueryEscape(payload.Value), Context: payload.Context})
	mutations = append(mutations, common.Payload{Value: html.EscapeString(payload.Value), Context: payload.Context})
	mutations = append(mutations, common.Payload{Value: strings.ToUpper(payload.Value), Context: payload.Context})
	mutations = append(mutations, common.Payload{Value: strings.ToLower(payload.Value), Context: payload.Context})
	return mutations
}
