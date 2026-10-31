package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getAutoAbility(r *http.Request, i handlerInput[seeding.AutoAbility, AutoAbility, NamedAPIResource, NamedAPIResourceList], id int32) (AutoAbility, error) {
	autoAbility, err := verifyParamsAndGet(r, i, id)
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
		RequiredItem:           nameAmountPtrToResAmtPtr(cfg, cfg.E.AllItems, autoAbility.RequiredItem),
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

func (cfg *Config) retrieveAutoAbilities(r *http.Request, i handlerInput[seeding.AutoAbility, AutoAbility, NamedAPIResource, NamedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IDFilter{
		enumListQuery(cfg, r, i, cfg.T.AutoAbilityCategory, ids, QpnCategory, cfg.db.GetAutoAbilityIDsByCategory),
		enumQuery(r, i, cfg.T.EquipType, ids, QpnType, cfg.db.GetAutoAbilityIDsByEquipType),
		idQueryWrapper(cfg, r, i, ids, QpnMonster, cfg.L.Monsters, getAutoAbilitiesByMonster),
		idQuery(r, i, ids, "monster_items", cfg.L.Monsters, cfg.db.GetAutoAbilityIDsByMonsterItems),
		idQueryWrapper(cfg, r, i, ids, QpnShop, cfg.L.Shops, getAutoAbilitiesByShop),
		valueListQuery(cfg, r, i, ids, QpnMethods, cfg.db.GetAutoAbilityIDsByMethods),
		idQuery(r, i, ids, QpnLocation, cfg.L.Locations, cfg.db.GetAutoAbilityIDsByLocation),
		idQuery(r, i, ids, QpnSublocation, cfg.L.Sublocations, cfg.db.GetAutoAbilityIDsBySublocation),
		idQuery(r, i, ids, "areas", cfg.L.Areas, cfg.db.GetAutoAbilityIDsByArea),
	})
}
