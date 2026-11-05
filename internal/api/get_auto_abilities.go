package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getAutoAbility(r *http.Request, i handlerInput[seedtypes.AutoAbility, AutoAbility, NamedAPIResource], id int32) (AutoAbility, error) {
	autoAbility, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return AutoAbility{}, err
	}

	rel, err := getAutoAbilityRelationships(srv, r, autoAbility)
	if err != nil {
		return AutoAbility{}, err
	}

	response := AutoAbility{
		ID:                     autoAbility.ID,
		Name:                   autoAbility.Name,
		Description:            autoAbility.Description,
		Effect:                 autoAbility.Effect,
		EquipType:              autoAbility.Type,
		Category:               autoAbility.Category,
		RelatedStats:           namesToNamedAPIResources(srv, E.Stats, autoAbility.RelatedStats),
		AbilityValue:           autoAbility.AbilityValue,
		RequiredItem:           E.AllItems.toResAmtFromNamePtr(srv, autoAbility.RequiredItem),
		LockedOutAutoAbilities: namesToNamedAPIResources(srv, E.AutoAbilities, autoAbility.LockedOutAbilities),
		ActivationCondition:    autoAbility.ActivationCondition,
		Counter:                autoAbility.Counter,
		GradualRecovery:        namePtrToNamedAPIResPtr(srv, E.Stats, autoAbility.GradualRecovery, nil),
		AutoItemUse:            namesToNamedAPIResources(srv, E.Items, autoAbility.AutoItemUse),
		OnHitElement:           namePtrToNamedAPIResPtr(srv, E.Elements, autoAbility.OnHitElement, nil),
		AddedElemResist:        convertObjPtr(srv, autoAbility.AddedElemResist, convertElemResist),
		OnHitStatus:            convertObjPtr(srv, autoAbility.OnHitStatus, convertInflictedStatus),
		AddedStatusses:         namesToNamedAPIResources(srv, E.StatusConditions, autoAbility.AddedStatusses),
		AddedStatusResists:     convertObjSlice(srv, autoAbility.AddedStatusResists, convertStatusResist),
		AddedProperty:          namePtrToNamedAPIResPtr(srv, E.Properties, autoAbility.AddedProperty, nil),
		ConversionTo:           namePtrToNamedAPIResPtr(srv, E.Modifiers, autoAbility.ConversionTo, nil),
		StatChanges:            convertObjSlice(srv, autoAbility.StatChanges, convertStatChange),
		ModifierChanges:        convertObjSlice(srv, autoAbility.ModifierChanges, convertModifierChange),
		MonstersDrop:           rel.MonstersDrop,
		MonstersItems:          rel.MonstersItems,
		ShopsPreAirship:        rel.ShopsPreAirship,
		ShopsPostAirship:       rel.ShopsPostAirship,
		Treasures:              rel.Treasures,
		EquipmentTables:        rel.EquipmentTables,
	}

	return response, nil
}

func (srv *Server) retrieveAutoAbilities(r *http.Request, i handlerInput[seedtypes.AutoAbility, AutoAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryEnumList(srv, r, enums.T.AutoAbilityCategory, ids, ts.QpnCategory, config.Cfg.DB.GetAutoAbilityIDsByCategory),
		i.queryEnum(r, enums.T.EquipType, ids, ts.QpnType, config.Cfg.DB.GetAutoAbilityIDsByEquipType),
		i.queryIDWrapper(srv, r, ids, ts.QpnMonster, lookup.L.Monsters, getAutoAbilitiesByMonster),
		i.queryID(r, ids, "monster_items", lookup.L.Monsters, config.Cfg.DB.GetAutoAbilityIDsByMonsterItems),
		i.queryIDWrapper(srv, r, ids, ts.QpnShop, lookup.L.Shops, getAutoAbilitiesByShop),
		i.queryValueList(srv, r, ids, ts.QpnMethods, config.Cfg.DB.GetAutoAbilityIDsByMethods),
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetAutoAbilityIDsByLocation),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetAutoAbilityIDsBySublocation),
		i.queryID(r, ids, "areas", lookup.L.Areas, config.Cfg.DB.GetAutoAbilityIDsByArea),
	})
}
