package api

type Location struct {
	ID                 int32              `json:"id"`
	Name               string             `json:"name"`
	ConnectedLocations []NamedAPIResource `json:"connected_locations"`
	Sublocations       []NamedAPIResource `json:"sublocations"`
	LocRel
}


type Sublocation struct {
	ID                    int32              `json:"id"`
	Name                  string             `json:"name"`
	ParentLocation        NamedAPIResource   `json:"parent_location"`
	ConnectedSublocations []NamedAPIResource `json:"connected_sublocations"`
	Areas                 []AreaAPIResource  `json:"areas"`
	LocRel
}


type Area struct {
	ID                int32            `json:"id"`
	Name              string           `json:"name"`
	Version           *int32           `json:"version,omitempty"`
	Specification     *string          `json:"specification,omitempty"`
	DisplayName       string           `json:"display_name"`
	ParentLocation    NamedAPIResource `json:"parent_location"`
	ParentSublocation NamedAPIResource `json:"parent_sublocation"`
	StoryOnly         bool             `json:"story_only"`
	HasSaveSphere     bool             `json:"has_save_sphere"`
	AirshipDropOff    bool             `json:"airship_drop_off"`
	HasCompSphere     bool             `json:"has_comp_sphere"`
	CanRideChocobo    bool             `json:"can_ride_chocobo"`
	ConnectedAreas    []AreaConnection `json:"connected_areas"`
	LocRel
}