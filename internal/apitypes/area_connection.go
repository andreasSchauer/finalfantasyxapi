package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type AreaConnection struct {
	Area           apiresource.AreaAPIResource `json:"area"`
	ConnectionType string                      `json:"connection_type"`
	IsStoryBased   bool                        `json:"is_story_based"`
	Notes          *string                     `json:"notes,omitempty"`
}

func (ac AreaConnection) GetAPIResource() apiresource.APIResource {
	return ac.Area
}