package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getCelestialWeapon(r *http.Request, i handlerInput[seedtypes.CelestialWeapon, CelestialWeapon, NamedAPIResource], id int32) (CelestialWeapon, error) {
	celestialWeapon, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return CelestialWeapon{}, err
	}

	rel, err := getCelestialWeaponRelationships(srv, r, celestialWeapon)
	if err != nil {
		return CelestialWeapon{}, err
	}

	response := CelestialWeapon{
		ID:            celestialWeapon.ID,
		Name:          celestialWeapon.Name,
		Formula:       celestialWeapon.Formula,
		Character:     nameToNamedAPIResource(srv, E.Characters, celestialWeapon.Character, nil),
		Aeon:          namePtrToNamedAPIResPtr(srv, E.Aeons, celestialWeapon.Aeon, nil),
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

func (srv *Server) retrieveCelestialWeapons(r *http.Request, i handlerInput[seedtypes.CelestialWeapon, CelestialWeapon, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryEnum(r, enums.T.CelestialFormula, ids, ts.QpnFormula, config.Cfg.DB.GetCelestialWeaponIDsByFormula),
	})
}
