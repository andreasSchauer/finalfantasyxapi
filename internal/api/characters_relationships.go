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

	character := Character{
		CelestialWeapon:        celestialWeapon,
		OverdriveCommand:       overdriveCommand,
		CharacterClasses:       characterClasses,
		DefaultPlayerAbilities: defaultAbilities,
		StdSgPlayerAbilities:   stdSgAbilities,
		ExpSgPlayerAbilities:   expSgAbilities,
		OverdriveModes:         getForeignSliceResAmts(cfg, cfg.e.overdriveModes, char, char.IsStoryBased, getCharModeAmount),
	}

	return character, nil
}


func getCharModeAmount(cfg *Config, char seeding.Character, id int32) (ResourceAmount[NamedAPIResource], error) {
	i := cfg.e.overdriveModes

	modeLookup, _ := seeding.GetResourceByID(id, i.objLookupID)
	if len(modeLookup.ActionsToLearn) == 0 {
		return ResourceAmount[NamedAPIResource]{}, errContinue
	}

	amount := modeLookup.ActionsToLearn[char.GetID()-1].Amount
	modeAmount := idAmountToResourceAmount(cfg, i, id, amount)

	return modeAmount, nil
}