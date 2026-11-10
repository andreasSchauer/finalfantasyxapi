package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"

type AreaConnection struct {
	Area           apiresource.Area `json:"area"`
	ConnectionType string           `json:"connection_type"`
	IsStoryBased   bool             `json:"is_story_based"`
	Notes          *string          `json:"notes,omitempty"`
}

func (ac AreaConnection) GetAPIResource() apiresource.APIResource {
	return ac.Area
}
