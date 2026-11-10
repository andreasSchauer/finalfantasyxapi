package apiresource

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

type Quest struct {
	ID        int32   `json:"-"`
	Sidequest string  `json:"sidequest"`
	Subquest  *string `json:"subquest,omitempty"`
	URL       string  `json:"url"`
}

func (r Quest) IsZero() bool {
	return r.Sidequest == "" && r.Subquest == nil
}

func (r Quest) GetID() int32 {
	return r.ID
}

func (r Quest) GetURL() string {
	return r.URL
}

func (r Quest) GetTestKey() string {
	var subquestStr string

	if r.Subquest != nil {
		subquestStr = fmt.Sprintf(" - %s", *r.Subquest)
	}
	return r.Sidequest + subquestStr
}

func (r Quest) String() string {
	return fmt.Sprintf("quest api resource '%s', url: %s", h.PtrToString(r.Subquest), r.URL)
}

func (r Quest) GetAPIResource() APIResource {
	return r
}
