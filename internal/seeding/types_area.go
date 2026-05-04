package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Area struct {
	ID                   int32
	Name                 string           `json:"area"`
	Version              *int32           `json:"version"`
	Specification        *string          `json:"specification"`
	Availability         string           `json:"availability"`
	HasSaveSphere        bool             `json:"has_save_sphere"`
	AirshipDropOff       bool             `json:"airship_drop_off"`
	HasCompilationSphere bool             `json:"has_compilation_sphere"`
	CanRideChocobo       bool             `json:"can_ride_chocobo"`
	ConnectedAreas       []AreaConnection `json:"connected_areas"`
	Sublocation          Sublocation
}

func (a Area) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.Sublocation.ID,
		a.Name,
		h.DerefOrNil(a.Version),
		h.DerefOrNil(a.Specification),
		a.Availability,
		a.HasSaveSphere,
		a.AirshipDropOff,
		a.HasCompilationSphere,
		a.CanRideChocobo,
	}
}

func (a Area) GetID() int32 {
	return a.ID
}

func (a *Area) SetID(id int32) {
	a.ID = id
}

func (a Area) Error() string {
	return fmt.Sprintf("area '%s'", h.NameToString(a.Name, a.Version, a.Specification))
}

func (a Area) GetLocationArea() LocationArea {
	return LocationArea{
		Location:    a.Sublocation.Location.Name,
		Sublocation: a.Sublocation.Name,
		Area:        a.Name,
		Version:     a.Version,
	}
}

func (a Area) GetResParamsLocation() h.ResParamsLocation {
	return h.ResParamsLocation{
		AreaID:        a.ID,
		Location:      a.Sublocation.Location.Name,
		Sublocation:   a.Sublocation.Name,
		Area:          a.Name,
		Version:       a.Version,
		Specification: a.Specification,
	}
}

type AreaConnection struct {
	ID             int32
	AreaID         int32
	LocationArea   LocationArea `json:"location_area"`
	ConnectionType string       `json:"connection_type"`
	IsStoryBased   bool         `json:"is_story_based"`
	Notes          *string      `json:"notes"`
}

func (ac AreaConnection) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", ac),
		ac.AreaID,
		ac.ConnectionType,
		ac.IsStoryBased,
		h.DerefOrNil(ac.Notes),
	}
}

func (ac AreaConnection) GetID() int32 {
	return ac.ID
}

func (ac *AreaConnection) SetID(id int32) {
	ac.ID = id
}

func (ac AreaConnection) Error() string {
	return fmt.Sprintf("area connection with %s", ac.LocationArea)
}
