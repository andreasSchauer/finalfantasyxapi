package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"

type Location struct {
	ID                 int32               `json:"id"`
	Name               string              `json:"name"`
	Availability       string              `json:"availability"`
	ConnectedLocations []apiresource.Named `json:"connected_locations"`
	Sublocations       []apiresource.Named `json:"sublocations"`
	LocRel
}

type Sublocation struct {
	ID                    int32               `json:"id"`
	Name                  string              `json:"name"`
	ParentLocation        apiresource.Named   `json:"parent_location"`
	Availability          string              `json:"availability"`
	ConnectedSublocations []apiresource.Named `json:"connected_sublocations"`
	Areas                 []apiresource.Area  `json:"areas"`
	LocRel
}

type Area struct {
	ID                int32             `json:"id"`
	Name              string            `json:"name"`
	Version           *int32            `json:"version,omitempty"`
	Specification     *string           `json:"specification,omitempty"`
	DisplayName       string            `json:"display_name"`
	ParentLocation    apiresource.Named `json:"parent_location"`
	ParentSublocation apiresource.Named `json:"parent_sublocation"`
	Availability      string            `json:"availability"`
	HasSaveSphere     bool              `json:"has_save_sphere"`
	AirshipDropOff    bool              `json:"airship_drop_off"`
	HasCompSphere     bool              `json:"has_comp_sphere"`
	CanRideChocobo    bool              `json:"can_ride_chocobo"`
	ConnectedAreas    []AreaConnection  `json:"connected_areas"`
	LocRel
}
