package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getAutoAbilityRelationships(cfg *Config, r *http.Request, autoAbility seeding.AutoAbility) (AutoAbility, error) {
	availabilityParams, err := getRelAvailabilityParams(cfg, r, cfg.e.autoAbilities, autoAbility.ID)
	if err != nil {
		return AutoAbility{}, err
	}

	monsterItems, err := runRelAvailabilityQuery(cfg, r, cfg.e.monsters, autoAbility, availabilityParams, convGetAutoAbilityItemMonsterIDs(cfg))
	if err != nil {
		return AutoAbility{}, err
	}

	monstersDrop, err := runRelAvailabilityQuery(cfg, r, cfg.e.monsters, autoAbility, availabilityParams, convGetAutoAbilityMonsterIDs(cfg))
	if err != nil {
		return AutoAbility{}, err
	}

	shopsPre, err := runRelAvailabilityQuery(cfg, r, cfg.e.shops, autoAbility, availabilityParams, convGetAutoAbilityShopIDsPre(cfg))
	if err != nil {
		return AutoAbility{}, err
	}

	shopsPost, err := runRelAvailabilityQuery(cfg, r, cfg.e.shops, autoAbility, availabilityParams, convGetAutoAbilityShopIDsPost(cfg))
	if err != nil {
		return AutoAbility{}, err
	}

	treasures, err := runRelAvailabilityQuery(cfg, r, cfg.e.treasures, autoAbility, availabilityParams, convGetAutoAbilityTreasuresIDs(cfg))
	if err != nil {
		return AutoAbility{}, err
	}

	equipmentTables, err := getResourcesDbItem(cfg, r, cfg.e.equipmentTables, autoAbility, cfg.db.GetAutoAbilityEquipmentTableIDs)
	if err != nil {
		return AutoAbility{}, err
	}

	rel := AutoAbility{
		MonstersDrop:     monstersDrop,
		MonstersItems:    []MonItemAmts{},
		ShopsPreAirship:  shopsPre,
		ShopsPostAirship: shopsPost,
		Treasures:        treasures,
		EquipmentTables:  equipmentTables,
	}

	if autoAbility.RequiredItem != nil {
		rel.MonstersItems = getMonItemAmts(cfg, monsterItems, autoAbility.RequiredItem.ItemName)
	}

	return rel, nil
}
