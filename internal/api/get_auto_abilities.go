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
		ID:                     autoAbility.ID,
		Name:                   autoAbility.Name,
		Description:            autoAbility.Description,
		Effect:                 autoAbility.Effect,
		EquipType:              autoAbility.Type,
		Category:               enumToNamedAPIResource(cfg, cfg.e.autoAbilityCategory.endpoint, autoAbility.Category, cfg.t.AutoAbilityCategory),
		RelatedStats:           namesToNamedAPIResources(cfg, cfg.e.stats, autoAbility.RelatedStats),
		AbilityValue:           autoAbility.AbilityValue,
		RequiredItem:           nameAmountPtrToResAmtPtr(cfg, cfg.e.allItems, autoAbility.RequiredItem),
		LockedOutAutoAbilities: namesToNamedAPIResources(cfg, cfg.e.autoAbilities, autoAbility.LockedOutAbilities),
		ActivationCondition:    autoAbility.ActivationCondition,
		Counter:                autoAbility.Counter,
		GradualRecovery:        namePtrToNamedAPIResPtr(cfg, cfg.e.stats, autoAbility.GradualRecovery, nil),
		AutoItemUse:            namesToNamedAPIResources(cfg, cfg.e.items, autoAbility.AutoItemUse),
		OnHitElement:           namePtrToNamedAPIResPtr(cfg, cfg.e.elements, autoAbility.OnHitElement, nil),
		AddedElemResist:        convertObjPtr(cfg, autoAbility.AddedElemResist, convertElemResist),
		OnHitStatus:            convertObjPtr(cfg, autoAbility.OnHitStatus, convertInflictedStatus),
		AddedStatusses:         namesToNamedAPIResources(cfg, cfg.e.statusConditions, autoAbility.AddedStatusses),
		AddedStatusResists:     convertObjSlice(cfg, autoAbility.AddedStatusResists, convertStatusResist),
		AddedProperty:          namePtrToNamedAPIResPtr(cfg, cfg.e.properties, autoAbility.AddedProperty, nil),
		ConversionTo:           namePtrToNamedAPIResPtr(cfg, cfg.e.modifiers, autoAbility.ConversionTo, nil),
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

func (cfg *Config) retrieveAutoAbilities(r *http.Request, i handlerInput[seeding.AutoAbility, AutoAbility, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumListQuery(cfg, r, i, cfg.t.AutoAbilityCategory, ids, "category", cfg.db.GetAutoAbilityIDsByCategory)),
		fidl(enumQuery(r, i, cfg.t.EquipType, ids, "type", cfg.db.GetAutoAbilityIDsByEquipType)),
		fidl(idQueryWrapper(cfg, r, i, ids, "monster", cfg.l.Monsters, getAutoAbilitiesByMonster)),
		fidl(idQuery(r, i, ids, "monster_items", cfg.l.Monsters, cfg.db.GetAutoAbilityIDsByMonsterItems)),
		fidl(idQueryWrapper(cfg, r, i, ids, "shop", cfg.l.Shops, getAutoAbilitiesByShop)),
		fidl(valueListQuery(cfg, r, i, ids, "methods", cfg.db.GetAutoAbilityIDsByMethods)),
		fidl(idQuery(r, i, ids, "location", cfg.l.Locations, cfg.db.GetAutoAbilityIDsByLocation)),
		fidl(idQuery(r, i, ids, "sublocation", cfg.l.Sublocations, cfg.db.GetAutoAbilityIDsBySublocation)),
		fidl(idQuery(r, i, ids, "areas", cfg.l.Areas, cfg.db.GetAutoAbilityIDsByArea)),
	})
}
