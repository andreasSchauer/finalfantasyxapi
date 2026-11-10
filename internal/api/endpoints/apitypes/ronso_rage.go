package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"

type RonsoRage struct {
	ID          int32               `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Effect      string              `json:"effect"`
	Overdrive   apiresource.Named   `json:"overdrive"`
	Monsters    []apiresource.Named `json:"monsters"`
}
