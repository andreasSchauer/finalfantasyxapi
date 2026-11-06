package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getCelestialWeapon(r *http.Request, i handlerInput[seedtypes.CelestialWeapon, apitypes.CelestialWeapon, apiresource.NamedAPIResource], id int32) (apitypes.CelestialWeapon, error) {
	celestialWeapon, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.CelestialWeapon{}, err
	}

	rel, err := getCelestialWeaponRelationships(r, celestialWeapon)
	if err != nil {
		return apitypes.CelestialWeapon{}, err
	}

	response := apitypes.CelestialWeapon{
		ID:            celestialWeapon.ID,
		Name:          celestialWeapon.Name,
		Formula:       celestialWeapon.Formula,
		Character:     nameToNamedAPIResource(E.Characters, celestialWeapon.Character, nil),
		Aeon:          namePtrToNamedAPIResPtr(E.Aeons, celestialWeapon.Aeon, nil),
		AutoAbilities: rel.AutoAbilities,
		Equipment:     rel.Equipment,
		Crest:         rel.Crest,
		Sigil:         rel.Sigil,
		WpnTreasure:   rel.WpnTreasure,
		CrestTreasure: rel.CrestTreasure,
		SigilQuest:    rel.SigilQuest,
	}

	return response, nil
}

func retrieveCelestialWeapons(r *http.Request, i handlerInput[seedtypes.CelestialWeapon, apitypes.CelestialWeapon, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryEnum(r, enums.T.CelestialFormula, ids, ts.QpnFormula, config.Cfg.DB.GetCelestialWeaponIDsByFormula),
	})
}
