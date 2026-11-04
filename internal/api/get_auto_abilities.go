package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getAutoAbility(r *http.Request, i handlerInput[seedtypes.AutoAbility, AutoAbility, NamedAPIResource], id int32) (AutoAbility, error) {
	autoAbility, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return AutoAbility{}, err
	}

	rel, err := getAutoAbilityRelationships(cfg, r, autoAbility)
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
		RelatedStats:           namesToNamedAPIResources(cfg, cfg.E.Stats, autoAbility.RelatedStats),
		AbilityValue:           autoAbility.AbilityValue,
		RequiredItem:           cfg.E.AllItems.toResAmtFromNamePtr(cfg, autoAbility.RequiredItem),
		LockedOutAutoAbilities: namesToNamedAPIResources(cfg, cfg.E.AutoAbilities, autoAbility.LockedOutAbilities),
		ActivationCondition:    autoAbility.ActivationCondition,
		Counter:                autoAbility.Counter,
		GradualRecovery:        namePtrToNamedAPIResPtr(cfg, cfg.E.Stats, autoAbility.GradualRecovery, nil),
		AutoItemUse:            namesToNamedAPIResources(cfg, cfg.E.Items, autoAbility.AutoItemUse),
		OnHitElement:           namePtrToNamedAPIResPtr(cfg, cfg.E.Elements, autoAbility.OnHitElement, nil),
		AddedElemResist:        convertObjPtr(cfg, autoAbility.AddedElemResist, convertElemResist),
		OnHitStatus:            convertObjPtr(cfg, autoAbility.OnHitStatus, convertInflictedStatus),
		AddedStatusses:         namesToNamedAPIResources(cfg, cfg.E.StatusConditions, autoAbility.AddedStatusses),
		AddedStatusResists:     convertObjSlice(cfg, autoAbility.AddedStatusResists, convertStatusResist),
		AddedProperty:          namePtrToNamedAPIResPtr(cfg, cfg.E.Properties, autoAbility.AddedProperty, nil),
		ConversionTo:           namePtrToNamedAPIResPtr(cfg, cfg.E.Modifiers, autoAbility.ConversionTo, nil),
		StatChanges:            convertObjSlice(cfg, autoAbility.StatChanges, convertStatChange),
		ModifierChanges:        convertObjSlice(cfg, autoAbility.ModifierChanges, convertModifierChange),
		MonstersDrop:           rel.MonstersDrop,
		MonstersItems:          rel.MonstersItems,
		ShopsPreAirship:        rel.ShopsPreAirship,
		ShopsPostAirship:       rel.ShopsPostAirship,
		Treasures:              rel.Treasures,
		EquipmentTables:        rel.EquipmentTables,
	}

	return response, nil
}

func (cfg *Config) retrieveAutoAbilities(r *http.Request, i handlerInput[seedtypes.AutoAbility, AutoAbility, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryEnumList(cfg, r, enums.T.AutoAbilityCategory, ids, ts.QpnCategory, cfg.db.GetAutoAbilityIDsByCategory),
		i.queryEnum(r, enums.T.EquipType, ids, ts.QpnType, cfg.db.GetAutoAbilityIDsByEquipType),
		i.queryIDWrapper(cfg, r, ids, ts.QpnMonster, lookup.L.Monsters, getAutoAbilitiesByMonster),
		i.queryID(r, ids, "monster_items", lookup.L.Monsters, cfg.db.GetAutoAbilityIDsByMonsterItems),
		i.queryIDWrapper(cfg, r, ids, ts.QpnShop, lookup.L.Shops, getAutoAbilitiesByShop),
		i.queryValueList(cfg, r, ids, ts.QpnMethods, cfg.db.GetAutoAbilityIDsByMethods),
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, cfg.db.GetAutoAbilityIDsByLocation),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, cfg.db.GetAutoAbilityIDsBySublocation),
		i.queryID(r, ids, "areas", lookup.L.Areas, cfg.db.GetAutoAbilityIDsByArea),
	})
}
