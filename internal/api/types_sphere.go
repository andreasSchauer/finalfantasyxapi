package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertCreatedNode(n seedtypes.CreatedNode) apitypes.CreatedNode {
	return apitypes.CreatedNode{
		Node:  n.Node,
		Value: n.Value,
	}
}
