package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getCharacterRelationships(cfg *Config, r *http.Request, char seeding.Character) (Character, error) {
	celestialWeapon, err := getResPtrDB(cfg, r, cfg.e.celestialWeapons, char, cfg.db.GetCharacterCelestialWeaponID)
	if err != nil {
		return Character{}, err
	}

	overdriveCommand, err := getResPtrDB(cfg, r, cfg.e.overdriveCommands, char, cfg.db.GetCharacterOverdriveCommandID)
	if err != nil {
		return Character{}, err
	}

	characterClasses, err := getResourcesDbItem(cfg, r, cfg.e.characterClasses, char, cfg.db.GetCharacterCharClassIDs)
	if err != nil {
		return Character{}, err
	}

	defaultAbilities, err := getResourcesDbItem(cfg, r, cfg.e.playerAbilities, char, cfg.db.GetCharacterDefaultAbilityIDs)
	if err != nil {
		return Character{}, err
	}

	stdSgAbilities, err := getResourcesDbItem(cfg, r, cfg.e.playerAbilities, char, cfg.db.GetCharacterSgAbilityIDs)
	if err != nil {
		return Character{}, err
	}
	expSgAbilities, err := getResourcesDbItem(cfg, r, cfg.e.playerAbilities, char, cfg.db.GetCharacterEgAbilityIDs)
	if err != nil {
		return Character{}, err
	}

	modeAmounts, err := getCharacterModes(cfg, r, char)
	if err != nil {
		return Character{}, err
	}

	character := Character{
		CelestialWeapon:        celestialWeapon,
		OverdriveCommand:       overdriveCommand,
		CharacterClasses:       characterClasses,
		DefaultPlayerAbilities: defaultAbilities,
		StdSgPlayerAbilities:   stdSgAbilities,
		ExpSgPlayerAbilities:   expSgAbilities,
		OverdriveModes:         modeAmounts,
	}

	return character, nil
}

func getCharacterModes(cfg *Config, r *http.Request, char seeding.Character) ([]ResourceAmount[NamedAPIResource], error) {
	modes := []ResourceAmount[NamedAPIResource]{}

	if char.IsStoryBased {
		return modes, nil
	}

	modeIDs, err := cfg.db.GetOverdriveModeIDs(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get overdrive modes", err)
	}

	for _, id := range modeIDs {
		i := cfg.e.overdriveModes
		modeLookup, _ := seeding.GetResourceByID(id, i.objLookupID)
		if len(modeLookup.ActionsToLearn) == 0 {
			continue
		}

		amount := modeLookup.ActionsToLearn[char.GetID()-1].Amount
		modeAmount := idAmountToResourceAmount(cfg, i, id, amount)
		modes = append(modes, modeAmount)
	}

	return modes, nil
}
