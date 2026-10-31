package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getCelestialWeapon(r *http.Request, i handlerInput[seeding.CelestialWeapon, CelestialWeapon, NamedAPIResource, NamedAPIResourceList], id int32) (CelestialWeapon, error) {
	celestialWeapon, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return CelestialWeapon{}, err
	}

	rel, err := getCelestialWeaponRelationships(cfg, r, celestialWeapon)
	if err != nil {
		return CelestialWeapon{}, err
	}

	response := CelestialWeapon{
		ID:            celestialWeapon.ID,
		Name:          celestialWeapon.Name,
		Formula:       celestialWeapon.Formula,
		Character:     nameToNamedAPIResource(cfg, cfg.E.Characters, celestialWeapon.Character, nil),
		Aeon:          namePtrToNamedAPIResPtr(cfg, cfg.E.Aeons, celestialWeapon.Aeon, nil),
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

func (cfg *Config) retrieveCelestialWeapons(r *http.Request, i handlerInput[seeding.CelestialWeapon, CelestialWeapon, NamedAPIResource, NamedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IDFilter{
		enumQuery(r, i, cfg.T.CelestialFormula, ids, QpnFormula, cfg.db.GetCelestialWeaponIDsByFormula),
	})
}
