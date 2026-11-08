package apiresource

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type Area struct {
	AreaID int32 `json:"-"`
	seedtypes.LocationArea
	Specification *string `json:"specification,omitempty"`
	URL           string  `json:"url"`
}

func (r Area) IsZero() bool {
	return r.Area == ""
}

func (r Area) GetID() int32 {
	return r.AreaID
}

func (r Area) GetURL() string {
	return r.URL
}

func (r Area) GetTestKey() string {
	return h.NameToString(r.Area, r.Version, nil)
}

func (r Area) String() string {
	return fmt.Sprintf("area api resource: %s, url: %s", r.LocationArea, r.URL)
}

func (r Area) GetAPIResource() APIResource {
	return r
}
