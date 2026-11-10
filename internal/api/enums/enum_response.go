package enums

import ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"

type EnumResponse struct {
	Name               ts.EnumName       `json:"name"`
	Description        string            `json:"description"`
	UsedByEndpointsInt []ts.EndpointName `json:"-"`
	UsedByEndpoints    []string          `json:"used_by_endpoints"`
	Values             []EnumVal         `json:"values"`
}
