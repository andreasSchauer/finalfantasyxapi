package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

type OverdriveSimple struct {
	ID                 int32    `json:"id"`
	URL                string   `json:"url"`
	Name               string   `json:"name"`
	Version            *int32   `json:"version,omitempty"`
	Specification      *string  `json:"specification,omitempty"`
	Rank               *int32   `json:"rank"`
	OverdriveAbilities []string `json:"overdrive_abilities"`
}

func (o OverdriveSimple) GetURL() string {
	return o.URL
}

func createOverdriveSimple(_ *http.Request, id int32, _ Subsection) (SimpleResource, error) {
	i := E.Overdrives
	overdrive, _ := lookup.GetResourceByID(id, i.objLookupID)

	overdriveSimple := OverdriveSimple{
		ID:                 overdrive.ID,
		URL:                CreateResourceURL(i.Endpoint, id),
		Name:               overdrive.Name,
		Version:            overdrive.Version,
		Specification:      overdrive.Specification,
		Rank:               overdrive.Rank,
		OverdriveAbilities: ConvertObjSlice(overdrive.OverdriveAbilities, abilityRefString),
	}

	return overdriveSimple, nil
}
