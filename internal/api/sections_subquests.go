package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type SubquestSimple struct {
	ID          int32                   `json:"id"`
	URL         string                  `json:"url"`
	Name        string                  `json:"name"`
	Condition *string           `json:"condition"`
	Areas     []string         	`json:"areas"`
	Reward    string 			`json:"reward"`
}

func (s SubquestSimple) GetURL() string {
	return s.URL
}


func createSubquestSimple(cfg *Config, _ *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.subquests
	subquest, _ := seeding.GetResourceByID(id, i.objLookupID)

	subquestSimple := SubquestSimple{
		ID:          subquest.ID,
		URL:         createResourceURL(cfg, i.endpoint, id),
		Name:        subquest.Name,
		Condition: 	 subquest.Completion.Condition,
		Areas: 		 locAreaStrings(cfg, subquest.Completion.Areas),
		Reward: 	 convertItemAmountSimple(cfg, subquest.GetItemAmount()),
	}

	return subquestSimple, nil
}
