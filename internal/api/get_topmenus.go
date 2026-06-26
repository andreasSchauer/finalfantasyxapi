package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getTopmenu(r *http.Request, i handlerInput[seeding.Topmenu, Topmenu, NamedAPIResource, NamedApiResourceList], id int32) (Topmenu, error) {
	topmenu, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Topmenu{}, err
	}

	rel, err := getTopmenuRelationships(cfg, r, topmenu)
	if err != nil {
		return Topmenu{}, err
	}

	response := Topmenu{
		ID:                topmenu.ID,
		Name:              topmenu.Name,
		Submenus:          rel.Submenus,
		OverdriveCommands: rel.OverdriveCommands,
		Overdrives:        rel.Overdrives,
		AeonCommands:      rel.AeonCommands,
		Abilities:         rel.Abilities,
	}

	return response, nil
}

func (cfg *Config) retrieveTopmenus(r *http.Request, i handlerInput[seeding.Topmenu, Topmenu, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	return verifyParamsAndRetrieve(r, i)
}
