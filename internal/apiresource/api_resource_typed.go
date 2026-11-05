package apiresource

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type TypedAPIResource struct {
	ID            int32   `json:"-"`
	Name          string  `json:"name"`
	Version       *int32  `json:"version,omitempty"`
	Specification *string `json:"specification,omitempty"`
	Type          string  `json:"type"`
	URL           string  `json:"url"`
}

func (r TypedAPIResource) IsZero() bool {
	return r.Name == ""
}

func (r TypedAPIResource) GetID() int32 {
	return r.ID
}

func (r TypedAPIResource) GetURL() string {
	return r.URL
}

func (r TypedAPIResource) GetTestKey() string {
	nameStr := h.NameToString(r.Name, r.Version, nil)
	return fmt.Sprintf("%s - %s", r.Type, nameStr)
}

func (r TypedAPIResource) String() string {
	return fmt.Sprintf("typed api resource '%s', type: %s, url: %s", h.NameToString(r.Name, r.Version, r.Specification), r.Type, r.URL)
}

func (r TypedAPIResource) GetAPIResource() APIResource {
	return r
}
