package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getModifier(r *http.Request, i handlerInput[seedtypes.Modifier, Modifier, apiresource.NamedAPIResource], id int32) (Modifier, error) {
	modifier, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Modifier{}, err
	}

	rel, err := getModifierRelationships(r, modifier)
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

func retrieveModifiers(r *http.Request, i handlerInput[seedtypes.Modifier, Modifier, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryEnumList(r, enums.T.ModifierCategory, ids, ts.QpnCategory, config.Cfg.DB.GetModifierIDsByCategory),
	})
}
