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

	modeAmounts, err := getCharacterModeAmounts(cfg, r, char)
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

func getCharacterModeAmounts(cfg *Config, r *http.Request, char seeding.Character) ([]ModeAmount, error) {
	modeAmounts := []ModeAmount{}

	if char.IsStoryBased {
		return modeAmounts, nil
	}

	modeIDs, err := cfg.db.GetOverdriveModeIDs(r.Context())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get overdrive modes", err)
	}

	for _, id := range modeIDs {
		modeLookup, _ := seeding.GetResourceByID(id, cfg.l.OverdriveModesID)
		if len(modeLookup.ActionsToLearn) == 0 {
			continue
		}

		mode := idToNamedAPIResource(cfg, cfg.e.overdriveModes, id)
		amount := modeLookup.ActionsToLearn[char.ID-1].Amount

		modeAmount := convertModeAmount(mode, amount)
		modeAmounts = append(modeAmounts, modeAmount)
	}

	return modeAmounts, nil
}
