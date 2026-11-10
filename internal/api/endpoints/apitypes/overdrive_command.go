package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"

type OverdriveCommand struct {
	ID          int32               `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Rank        int32               `json:"rank"`
	User        apiresource.Named   `json:"user"`
	Topmenu     *apiresource.Named  `json:"topmenu"`
	OpenSubmenu apiresource.Named   `json:"open_submenu"`
	Overdrives  []apiresource.Named `json:"overdrives"`
}
