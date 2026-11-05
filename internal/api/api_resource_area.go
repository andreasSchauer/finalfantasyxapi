package api

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type AreaAPIResource struct {
	AreaID int32 `json:"-"`
	seedtypes.LocationArea
	Specification *string `json:"specification,omitempty"`
	URL           string  `json:"url"`
}

func (r AreaAPIResource) IsZero() bool {
	return r.Area == ""
}

func (r AreaAPIResource) GetID() int32 {
	return r.AreaID
}

func (r AreaAPIResource) GetURL() string {
	return r.URL
}

func (r AreaAPIResource) GetTestKey() string {
	return h.NameToString(r.Area, r.Version, nil)
}

func (r AreaAPIResource) String() string {
	return fmt.Sprintf("area api resource: %s, url: %s", r.LocationArea, r.URL)
}

func (r AreaAPIResource) GetAPIResource() APIResource {
	return r
}
