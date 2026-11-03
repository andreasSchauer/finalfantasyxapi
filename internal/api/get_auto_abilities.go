package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
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
		i.queryEnumList(cfg, r, cfg.T.AutoAbilityCategory, ids, QpnCategory, cfg.db.GetAutoAbilityIDsByCategory),
		i.queryEnum(r, cfg.T.EquipType, ids, QpnType, cfg.db.GetAutoAbilityIDsByEquipType),
		i.queryIDWrapper(cfg, r, ids, QpnMonster, cfg.L.Monsters, getAutoAbilitiesByMonster),
		i.queryID(r, ids, "monster_items", cfg.L.Monsters, cfg.db.GetAutoAbilityIDsByMonsterItems),
		i.queryIDWrapper(cfg, r, ids, QpnShop, cfg.L.Shops, getAutoAbilitiesByShop),
		i.queryValueList(cfg, r, ids, QpnMethods, cfg.db.GetAutoAbilityIDsByMethods),
		i.queryID(r, ids, QpnLocation, cfg.L.Locations, cfg.db.GetAutoAbilityIDsByLocation),
		i.queryID(r, ids, QpnSublocation, cfg.L.Sublocations, cfg.db.GetAutoAbilityIDsBySublocation),
		i.queryID(r, ids, "areas", cfg.L.Areas, cfg.db.GetAutoAbilityIDsByArea),
	})
}
