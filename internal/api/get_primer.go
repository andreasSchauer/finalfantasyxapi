package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getPrimer(r *http.Request, i handlerInput[seeding.Primer, Primer, NamedAPIResource, NamedApiResourceList], id int32) (Primer, error) {
	primer, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Primer{}, err
	}

	rel, err := getPrimerRelationships(cfg, r, primer)
	if err != nil {
		return Primer{}, err
	}

	keyItem, _ := seeding.GetResource(primer.Name, cfg.l.KeyItems)

	response := Primer{
		ID:                 primer.ID,
		Name:               primer.Name,
		KeyItem: 			nameToNamedAPIResource(cfg, cfg.e.keyItems, keyItem.Name, nil),
		Description: 		keyItem.Description,
		AlBhedLetter: 		primer.AlBhedLetter,
		EnglishLetter: 		primer.EnglishLetter,
		Treasures:          rel.Treasures,
		Areas: 				rel.Areas,
	}

	return response, nil
}

func (cfg *Config) retrievePrimers(r *http.Request, i handlerInput[seeding.Primer, Primer, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(enumListQuery(cfg, r, i, cfg.t.AvailabilityType, resources, "availability", cfg.db.GetPrimerIDsByAvailability)),
	})
}
