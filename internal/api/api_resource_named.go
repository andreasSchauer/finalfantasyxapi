package api

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type NamedAPIResource struct {
	ID            int32   `json:"-"`
	Name          string  `json:"name"`
	Version       *int32  `json:"version,omitempty"`
	Specification *string `json:"specification,omitempty"`
	URL           string  `json:"url"`
}

func (r NamedAPIResource) IsZero() bool {
	return r.Name == ""
}

func (r NamedAPIResource) GetID() int32 {
	return r.ID
}

func (r NamedAPIResource) GetURL() string {
	return r.URL
}

func (r NamedAPIResource) GetTestKey() string {
	return h.NameToString(r.Name, r.Version, nil)
}

func (r NamedAPIResource) String() string {
	return fmt.Sprintf("named api resource '%s', url: %s", h.NameToString(r.Name, r.Version, r.Specification), r.URL)
}

func (r NamedAPIResource) GetAPIResource() APIResource {
	return r
}

func (r NamedAPIResource) NameIs(name string, version *int32) bool {
	return r.Name == name && h.PtrsEqual(r.Version, version, h.DerefInt32Ptr)
}
