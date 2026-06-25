package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getModifier(r *http.Request, i handlerInput[seeding.Modifier, Modifier, NamedAPIResource, NamedApiResourceList], id int32) (Modifier, error) {
	modifier, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Modifier{}, err
	}

	rel, err := getModifierRelationships(cfg, r, modifier)
	if err != nil {
		return Modifier{}, err
	}

	response := Modifier{
		ID:                 modifier.ID,
		Name:               modifier.Name,
		Effect:             modifier.Effect,
		Category:           modifier.Category,
		DefaultValue:       modifier.DefaultValue,
		AutoAbilities:      rel.AutoAbilities,
		PlayerAbilities:    rel.PlayerAbilities,
		OverdriveAbilities: rel.OverdriveAbilities,
		ItemAbilities:      rel.ItemAbilities,
		TriggerCommands:    rel.TriggerCommands,
		EnemyAbilities:     rel.EnemyAbilities,
		StatusConditions:   rel.StatusConditions,
		Properties:         rel.Properties,
	}

	return response, nil
}

func (cfg *Config) retrieveModifiers(r *http.Request, i handlerInput[seeding.Modifier, Modifier, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumListQuery(cfg, r, i, cfg.t.ModifierCategory, ids, qpnCategory, cfg.db.GetModifierIDsByCategory)),
	})
}
