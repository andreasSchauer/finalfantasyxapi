package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getAutoAbility(r *http.Request, i handlerInput[seedtypes.AutoAbility, apitypes.AutoAbility, apiresource.NamedAPIResource], id int32) (apitypes.AutoAbility, error) {
	autoAbility, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.AutoAbility{}, err
	}

	rel, err := getAutoAbilityRelationships(r, autoAbility)
	if err != nil {
		return apitypes.AutoAbility{}, err
	}

	response := apitypes.AutoAbility{
		ID:                     autoAbility.ID,
		Name:                   autoAbility.Name,
		Description:            autoAbility.Description,
		Effect:                 autoAbility.Effect,
		EquipType:              autoAbility.Type,
		Category:               autoAbility.Category,
		RelatedStats:           namesToNamedAPIResources(E.Stats, autoAbility.RelatedStats),
		AbilityValue:           autoAbility.AbilityValue,
		RequiredItem:           E.AllItems.toResAmtFromNamePtr(autoAbility.RequiredItem),
		LockedOutAutoAbilities: namesToNamedAPIResources(E.AutoAbilities, autoAbility.LockedOutAbilities),
		ActivationCondition:    autoAbility.ActivationCondition,
		Counter:                autoAbility.Counter,
		GradualRecovery:        namePtrToNamedAPIResPtr(E.Stats, autoAbility.GradualRecovery, nil),
		AutoItemUse:            namesToNamedAPIResources(E.Items, autoAbility.AutoItemUse),
		OnHitElement:           namePtrToNamedAPIResPtr(E.Elements, autoAbility.OnHitElement, nil),
		AddedElemResist:        ah.ConvertObjPtr(autoAbility.AddedElemResist, convertElemResist),
		OnHitStatus:            ah.ConvertObjPtr(autoAbility.OnHitStatus, convertInflictedStatus),
		AddedStatusses:         namesToNamedAPIResources(E.StatusConditions, autoAbility.AddedStatusses),
		AddedStatusResists:     ah.ConvertObjSlice(autoAbility.AddedStatusResists, convertStatusResist),
		AddedProperty:          namePtrToNamedAPIResPtr(E.Properties, autoAbility.AddedProperty, nil),
		ConversionTo:           namePtrToNamedAPIResPtr(E.Modifiers, autoAbility.ConversionTo, nil),
		StatChanges:            ah.ConvertObjSlice(autoAbility.StatChanges, convertStatChange),
		ModifierChanges:        ah.ConvertObjSlice(autoAbility.ModifierChanges, convertModifierChange),
		MonstersDrop:           rel.MonstersDrop,
		MonstersItems:          rel.MonstersItems,
		ShopsPreAirship:        rel.ShopsPreAirship,
		ShopsPostAirship:       rel.ShopsPostAirship,
		Treasures:              rel.Treasures,
		EquipmentTables:        rel.EquipmentTables,
	}

	return response, nil
}

func retrieveAutoAbilities(r *http.Request, i handlerInput[seedtypes.AutoAbility, apitypes.AutoAbility, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryEnumList(r, enums.T.AutoAbilityCategory, ids, ts.QpnCategory, config.Cfg.DB.GetAutoAbilityIDsByCategory),
		i.queryEnum(r, enums.T.EquipType, ids, ts.QpnType, config.Cfg.DB.GetAutoAbilityIDsByEquipType),
		i.queryIDWrapper(r, ids, ts.QpnMonster, lookup.L.Monsters, getAutoAbilitiesByMonster),
		i.queryID(r, ids, "monster_items", lookup.L.Monsters, config.Cfg.DB.GetAutoAbilityIDsByMonsterItems),
		i.queryIDWrapper(r, ids, ts.QpnShop, lookup.L.Shops, getAutoAbilitiesByShop),
		i.queryValueList(r, ids, ts.QpnMethods, config.Cfg.DB.GetAutoAbilityIDsByMethods),
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetAutoAbilityIDsByLocation),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetAutoAbilityIDsBySublocation),
		i.queryID(r, ids, "areas", lookup.L.Areas, config.Cfg.DB.GetAutoAbilityIDsByArea),
	})
}
