package apiconv

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func ConvertCreatedNode(n seedtypes.CreatedNode) apitypes.CreatedNode {
	return apitypes.CreatedNode{
		Node:  n.Node,
		Value: n.Value,
	}
}
