package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getAutoAbility(r *http.Request, i handlerInput[seedtypes.AutoAbility, AutoAbility, NamedAPIResource], id int32) (AutoAbility, error) {
	autoAbility, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return AutoAbility{}, err
	}

	rel, err := getAutoAbilityRelationships(s, r, autoAbility)
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
		RelatedStats:           namesToNamedAPIResources(s, E.Stats, autoAbility.RelatedStats),
		AbilityValue:           autoAbility.AbilityValue,
		RequiredItem:           E.AllItems.toResAmtFromNamePtr(s, autoAbility.RequiredItem),
		LockedOutAutoAbilities: namesToNamedAPIResources(s, E.AutoAbilities, autoAbility.LockedOutAbilities),
		ActivationCondition:    autoAbility.ActivationCondition,
		Counter:                autoAbility.Counter,
		GradualRecovery:        namePtrToNamedAPIResPtr(s, E.Stats, autoAbility.GradualRecovery, nil),
		AutoItemUse:            namesToNamedAPIResources(s, E.Items, autoAbility.AutoItemUse),
		OnHitElement:           namePtrToNamedAPIResPtr(s, E.Elements, autoAbility.OnHitElement, nil),
		AddedElemResist:        convertObjPtr(s, autoAbility.AddedElemResist, convertElemResist),
		OnHitStatus:            convertObjPtr(s, autoAbility.OnHitStatus, convertInflictedStatus),
		AddedStatusses:         namesToNamedAPIResources(s, E.StatusConditions, autoAbility.AddedStatusses),
		AddedStatusResists:     convertObjSlice(s, autoAbility.AddedStatusResists, convertStatusResist),
		AddedProperty:          namePtrToNamedAPIResPtr(s, E.Properties, autoAbility.AddedProperty, nil),
		ConversionTo:           namePtrToNamedAPIResPtr(s, E.Modifiers, autoAbility.ConversionTo, nil),
		StatChanges:            convertObjSlice(s, autoAbility.StatChanges, convertStatChange),
		ModifierChanges:        convertObjSlice(s, autoAbility.ModifierChanges, convertModifierChange),
		MonstersDrop:           rel.MonstersDrop,
		MonstersItems:          rel.MonstersItems,
		ShopsPreAirship:        rel.ShopsPreAirship,
		ShopsPostAirship:       rel.ShopsPostAirship,
		Treasures:              rel.Treasures,
		EquipmentTables:        rel.EquipmentTables,
	}

	return response, nil
}

func (s *Server) retrieveAutoAbilities(r *http.Request, i handlerInput[seedtypes.AutoAbility, AutoAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryEnumList(s, r, enums.T.AutoAbilityCategory, ids, ts.QpnCategory, s.Cfg.DB.GetAutoAbilityIDsByCategory),
		i.queryEnum(r, enums.T.EquipType, ids, ts.QpnType, s.Cfg.DB.GetAutoAbilityIDsByEquipType),
		i.queryIDWrapper(s, r, ids, ts.QpnMonster, lookup.L.Monsters, getAutoAbilitiesByMonster),
		i.queryID(r, ids, "monster_items", lookup.L.Monsters, s.Cfg.DB.GetAutoAbilityIDsByMonsterItems),
		i.queryIDWrapper(s, r, ids, ts.QpnShop, lookup.L.Shops, getAutoAbilitiesByShop),
		i.queryValueList(s, r, ids, ts.QpnMethods, s.Cfg.DB.GetAutoAbilityIDsByMethods),
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, s.Cfg.DB.GetAutoAbilityIDsByLocation),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, s.Cfg.DB.GetAutoAbilityIDsBySublocation),
		i.queryID(r, ids, "areas", lookup.L.Areas, s.Cfg.DB.GetAutoAbilityIDsByArea),
	})
}
