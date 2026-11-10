package apiresource

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

type Named struct {
	ID            int32   `json:"-"`
	Name          string  `json:"name"`
	Version       *int32  `json:"version,omitempty"`
	Specification *string `json:"specification,omitempty"`
	URL           string  `json:"url"`
}

func (r Named) IsZero() bool {
	return r.Name == ""
}

func (r Named) GetID() int32 {
	return r.ID
}

func (r Named) GetURL() string {
	return r.URL
}

func (r Named) GetTestKey() string {
	return h.NameToString(r.Name, r.Version, nil)
}

func (r Named) String() string {
	return fmt.Sprintf("named api resource '%s', url: %s", h.NameToString(r.Name, r.Version, r.Specification), r.URL)
}

func (r Named) GetAPIResource() APIResource {
	return r
}

func (r Named) NameIs(name string, version *int32) bool {
	return r.Name == name && h.PtrsEqual(r.Version, version, h.DerefInt32Ptr)
}
