package main

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type LocationArea struct {
	Location    string `json:"location"`
	SubLocation string `json:"sublocation"`
	Area        string `json:"area"`
	Version     *int32 `json:"version,omitempty"`
}

func (la LocationArea) ToKeyFields() []any {
	return []any{
		la.Location,
		la.SubLocation,
		la.Area,
		h.DerefOrNil(la.Version),
	}
}

type IsLocationArea interface {
	getLocationArea() LocationArea
}

func (la LocationArea) Error() string {
	return fmt.Sprintf("location area with location: '%s', sublocation: '%s', area: '%s', version: '%v'", la.Location, la.SubLocation, la.Area, h.DerefOrNil(la.Version))
}

func newLocationArea(location, sublocation, area string, version *int32) LocationArea {
	return LocationArea{
		Location:    location,
		SubLocation: sublocation,
		Area:        area,
		Version:     version,
	}
}
