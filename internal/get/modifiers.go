package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/assemble"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func GetModifier(r *http.Request, i handlerinput.Main[seedtypes.Modifier, apitypes.Modifier, apiresource.Named], id int32) (apitypes.Modifier, error) {
	modifier, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Modifier{}, err
	}

	rel, err := assemble.GetModifierRelationships(r, modifier)
	if err != nil {
		return apitypes.Modifier{}, err
	}

	response := apitypes.Modifier{
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

func RetrieveModifiers(r *http.Request, i handlerinput.Main[seedtypes.Modifier, apitypes.Modifier, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.ModifierCategory, ids, ts.QpnCategory, config.Cfg.DB.GetModifierIDsByCategory),
	})
}
