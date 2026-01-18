package main

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type LocationArea struct {
	Location    string `json:"location"`
	Sublocation string `json:"sublocation"`
	Area        string `json:"area"`
	Version     *int32 `json:"version,omitempty"`
}


func (la LocationArea) Error() string {
	return fmt.Sprintf("location area with location: '%s', sublocation: '%s', area: '%s', version: '%v'", la.Location, la.Sublocation, la.Area, h.DerefOrNil(la.Version))
}