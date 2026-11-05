package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

type SubquestSimple struct {
	ID        int32    `json:"id"`
	URL       string   `json:"url"`
	Name      string   `json:"name"`
	Condition *string  `json:"condition"`
	Areas     []string `json:"areas"`
	Reward    string   `json:"reward"`
}

func (s SubquestSimple) GetURL() string {
	return s.URL
}

func createSubquestSimple(_ *http.Request, id int32, _ Subsection) (apiresource.SimpleResource, error) {
	i := E.Subquests
	subquest, _ := lookup.GetResourceByID(id, i.objLookupID)

	subquestSimple := SubquestSimple{
		ID:        subquest.ID,
		URL:       ah.CreateResourceURL(i.Endpoint, id),
		Name:      subquest.Name,
		Condition: subquest.Completion.Condition,
		Areas:     locAreaStrings(subquest.Completion.Areas),
		Reward:    convertItemAmountSimple(subquest.GetItemAmount()),
	}

	return subquestSimple, nil
}
