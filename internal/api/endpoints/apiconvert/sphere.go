package apiconvert

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func CreatedNode(n seedtypes.CreatedNode) apitypes.CreatedNode {
	return apitypes.CreatedNode{
		Node:  n.Node,
		Value: n.Value,
	}
}
