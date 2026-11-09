package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/assemble"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getAutoAbility(r *http.Request, i handlerinput.Main[seedtypes.AutoAbility, apitypes.AutoAbility, apiresource.Named], id int32) (apitypes.AutoAbility, error) {
	autoAbility, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.AutoAbility{}, err
	}

	rel, err := assemble.GetAutoAbilityRelationships(r, autoAbility)
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
		RelatedStats:           handlerinput.NamesToNamedAPIResources(apilookups.E.Stats, autoAbility.RelatedStats),
		AbilityValue:           autoAbility.AbilityValue,
		RequiredItem:           apilookups.E.AllItems.ToResAmtFromNamePtr(autoAbility.RequiredItem),
		LockedOutAutoAbilities: handlerinput.NamesToNamedAPIResources(apilookups.E.AutoAbilities, autoAbility.LockedOutAbilities),
		ActivationCondition:    autoAbility.ActivationCondition,
		Counter:                autoAbility.Counter,
		GradualRecovery:        handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Stats, autoAbility.GradualRecovery, nil),
		AutoItemUse:            handlerinput.NamesToNamedAPIResources(apilookups.E.Items, autoAbility.AutoItemUse),
		OnHitElement:           handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Elements, autoAbility.OnHitElement, nil),
		AddedElemResist:        ah.ConvertObjPtr(autoAbility.AddedElemResist, apiconv.ConvertElemResist),
		OnHitStatus:            ah.ConvertObjPtr(autoAbility.OnHitStatus, apiconv.ConvertInflictedStatus),
		AddedStatusses:         handlerinput.NamesToNamedAPIResources(apilookups.E.StatusConditions, autoAbility.AddedStatusses),
		AddedStatusResists:     ah.ConvertObjSlice(autoAbility.AddedStatusResists, apiconv.ConvertStatusResist),
		AddedProperty:          handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Properties, autoAbility.AddedProperty, nil),
		ConversionTo:           handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Modifiers, autoAbility.ConversionTo, nil),
		StatChanges:            ah.ConvertObjSlice(autoAbility.StatChanges, apiconv.ConvertStatChange),
		ModifierChanges:        ah.ConvertObjSlice(autoAbility.ModifierChanges, apiconv.ConvertModifierChange),
		MonstersDrop:           rel.MonstersDrop,
		MonstersItems:          rel.MonstersItems,
		ShopsPreAirship:        rel.ShopsPreAirship,
		ShopsPostAirship:       rel.ShopsPostAirship,
		Treasures:              rel.Treasures,
		EquipmentTables:        rel.EquipmentTables,
	}

	return response, nil
}

func retrieveAutoAbilities(r *http.Request, i handlerinput.Main[seedtypes.AutoAbility, apitypes.AutoAbility, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.AutoAbilityCategory, ids, ts.QpnCategory, config.Cfg.DB.GetAutoAbilityIDsByCategory),
		i.QueryEnum(r, enums.T.EquipType, ids, ts.QpnType, config.Cfg.DB.GetAutoAbilityIDsByEquipType),
		i.QueryIDWrapper(r, ids, ts.QpnMonster, lookup.L.Monsters, getAutoAbilitiesByMonster),
		i.QueryID(r, ids, "monster_items", lookup.L.Monsters, config.Cfg.DB.GetAutoAbilityIDsByMonsterItems),
		i.QueryIDWrapper(r, ids, ts.QpnShop, lookup.L.Shops, getAutoAbilitiesByShop),
		i.QueryValueList(r, ids, ts.QpnMethods, config.Cfg.DB.GetAutoAbilityIDsByMethods),
		i.QueryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetAutoAbilityIDsByLocation),
		i.QueryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetAutoAbilityIDsBySublocation),
		i.QueryID(r, ids, "areas", lookup.L.Areas, config.Cfg.DB.GetAutoAbilityIDsByArea),
	})
}
