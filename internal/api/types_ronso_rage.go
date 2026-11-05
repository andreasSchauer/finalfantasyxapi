package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type RonsoRage struct {
	ID          int32                          `json:"id"`
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	Effect      string                         `json:"effect"`
	Overdrive   apiresource.NamedAPIResource   `json:"overdrive"`
	Monsters    []apiresource.NamedAPIResource `json:"monsters"`
}
