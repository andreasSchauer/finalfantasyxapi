package apiresource

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type QuestAPIResource struct {
	ID        int32   `json:"-"`
	Sidequest string  `json:"sidequest"`
	Subquest  *string `json:"subquest,omitempty"`
	URL       string  `json:"url"`
}

func (r QuestAPIResource) IsZero() bool {
	return r.Sidequest == "" && r.Subquest == nil
}

func (r QuestAPIResource) GetID() int32 {
	return r.ID
}

func (r QuestAPIResource) GetURL() string {
	return r.URL
}

func (r QuestAPIResource) GetTestKey() string {
	var subquestStr string

	if r.Subquest != nil {
		subquestStr = fmt.Sprintf(" - %s", *r.Subquest)
	}
	return r.Sidequest + subquestStr
}

func (r QuestAPIResource) String() string {
	return fmt.Sprintf("quest api resource '%s', url: %s", h.PtrToString(r.Subquest), r.URL)
}

func (r QuestAPIResource) GetAPIResource() APIResource {
	return r
}
