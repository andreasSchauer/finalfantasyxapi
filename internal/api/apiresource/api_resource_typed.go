package apiresource

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

type Typed struct {
	ID            int32   `json:"-"`
	Name          string  `json:"name"`
	Version       *int32  `json:"version,omitempty"`
	Specification *string `json:"specification,omitempty"`
	Type          string  `json:"type"`
	URL           string  `json:"url"`
}

func (r Typed) IsZero() bool {
	return r.Name == ""
}

func (r Typed) GetID() int32 {
	return r.ID
}

func (r Typed) GetURL() string {
	return r.URL
}

func (r Typed) GetTestKey() string {
	nameStr := h.NameToString(r.Name, r.Version, nil)
	return fmt.Sprintf("%s - %s", r.Type, nameStr)
}

func (r Typed) String() string {
	return fmt.Sprintf("typed api resource '%s', type: %s, url: %s", h.NameToString(r.Name, r.Version, r.Specification), r.Type, r.URL)
}

func (r Typed) GetAPIResource() APIResource {
	return r
}
