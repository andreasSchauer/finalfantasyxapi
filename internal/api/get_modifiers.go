package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getModifier(r *http.Request, i handlerInput[seedtypes.Modifier, Modifier, NamedAPIResource], id int32) (Modifier, error) {
	modifier, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Modifier{}, err
	}

	rel, err := getModifierRelationships(srv, r, modifier)
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

func (srv *Server) retrieveModifiers(r *http.Request, i handlerInput[seedtypes.Modifier, Modifier, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryEnumList(srv, r, enums.T.ModifierCategory, ids, ts.QpnCategory, srv.Cfg.DB.GetModifierIDsByCategory),
	})
}
