package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type OverdriveCommand struct {
	ID          int32                          `json:"id"`
	Name        string                         `json:"name"`
	Description string                         `json:"description"`
	Rank        int32                          `json:"rank"`
	User        apiresource.NamedAPIResource   `json:"user"`
	Topmenu     *apiresource.NamedAPIResource  `json:"topmenu"`
	OpenSubmenu apiresource.NamedAPIResource   `json:"open_submenu"`
	Overdrives  []apiresource.NamedAPIResource `json:"overdrives"`
}
