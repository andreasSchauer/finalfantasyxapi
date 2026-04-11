package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getAutoAbility(r *http.Request, i handlerInput[seeding.AutoAbility, AutoAbility, NamedAPIResource, NamedApiResourceList], id int32) (AutoAbility, error) {
	autoAbility, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return AutoAbility{}, err
	}

	rel, err := getAutoAbilityRelationships(cfg, r, autoAbility)
	if err != nil {
		return AutoAbility{}, err
	}

	response := AutoAbility{
		ID:                 	autoAbility.ID,
		Name:               	autoAbility.Name,
		Description: 			autoAbility.Description,
		Effect: 				autoAbility.Effect,
		EquipType: 				autoAbility.Type,
		Category: 				newNamedAPIResourceFromEnum(cfg, cfg.e.autoAbilityCategory.endpoint, autoAbility.Category, cfg.t.AutoAbilityCategory),
		RelatedStats: 			namesToNamedAPIResources(cfg, cfg.e.stats, autoAbility.RelatedStats),
		AbilityValue: 			autoAbility.AbilityValue,
		RequiredItem: 			nameAmountPtrToResAmtPtr(cfg, cfg.e.allItems, autoAbility.RequiredItem),
		LockedOutAutoAbilities: namesToNamedAPIResources(cfg, cfg.e.autoAbilities, autoAbility.LockedOutAbilities),
		ActivationCondition: 	autoAbility.ActivationCondition,
		Counter: 				autoAbility.Counter,
		GradualRecovery: 		namePtrToNamedAPIResPtr(cfg, cfg.e.stats, autoAbility.GradualRecovery, nil),
		AutoItemUse: 			namesToNamedAPIResources(cfg, cfg.e.items, autoAbility.AutoItemUse),
		OnHitElement: 			namePtrToNamedAPIResPtr(cfg, cfg.e.elements, autoAbility.OnHitElement, nil),
		AddedElemResist: 		convertObjPtr(cfg, autoAbility.AddedElemResist, convertElemResist),
		OnHitStatus: 			convertObjPtr(cfg, autoAbility.OnHitStatus, convertInflictedStatus),
		AddedStatusses: 		namesToNamedAPIResources(cfg, cfg.e.statusConditions, autoAbility.AddedStatusses),
		AddedStatusResists: 	convertObjSlice(cfg, autoAbility.AddedStatusResists, convertStatusResist),
		AddedProperty: 			namePtrToNamedAPIResPtr(cfg, cfg.e.properties, autoAbility.AddedProperty, nil),
		ConversionTo: 			namePtrToNamedAPIResPtr(cfg, cfg.e.modifiers, autoAbility.ConversionTo, nil),
		StatChanges: 			convertObjSlice(cfg, autoAbility.StatChanges, convertStatChange),
		ModifierChanges: 		convertObjSlice(cfg, autoAbility.ModifierChanges, convertModifierChange),
		MonstersDrop: 			rel.MonstersDrop,
		MonstersItems: 			rel.MonstersItems,
		ShopsPreAirship: 		rel.ShopsPreAirship,
		ShopsPostAirship: 		rel.ShopsPostAirship,
		Treasures: 				rel.Treasures,
		EquipmentTables: 		rel.EquipmentTables,
	}

	return response, nil
}



func (cfg *Config) retrieveAutoAbilities(r *http.Request, i handlerInput[seeding.AutoAbility, AutoAbility, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(enumListQuery(cfg, r, i, cfg.t.AutoAbilityCategory, resources, "category", cfg.db.GetAutoAbilityIDsByCategory)),
		frl(enumQuery(cfg, r, i, cfg.t.EquipType, resources, "type", cfg.db.GetAutoAbilityIDsByEquipType)),
		frl(idQuery(cfg, r, i, resources, "monster", len(cfg.l.Monsters), cfg.db.GetAutoAbilityIDsByMonster)),
		frl(idQuery(cfg, r, i, resources, "monster_items", len(cfg.l.Monsters), cfg.db.GetAutoAbilityIDsByMonsterItems)),
	})
}