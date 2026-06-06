package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getAutoAbilityRelationships(cfg *Config, r *http.Request, autoAbility seeding.AutoAbility) (AutoAbility, error) {
	availabilityParams, err := getRelAvailabilityParams(cfg, r, cfg.e.autoAbilities, autoAbility.ID)
	if err != nil {
		return AutoAbility{}, err
	}

	monsterItems, err := runRelAvailabilityQuery(cfg, r, cfg.e.monsters, autoAbility, availabilityParams, getAutoAbilityItemMonsterIDs(cfg))
	if err != nil {
		return AutoAbility{}, err
	}

	monstersDrop, err := runRelAvailabilityQuery(cfg, r, cfg.e.monsters, autoAbility, availabilityParams, getAutoAbilitySourceIDs(cfg, ViewSourceTypeMonster, nil))
	if err != nil {
		return AutoAbility{}, err
	}

	preAirship := database.ShopTypePreAirship
	shopsPre, err := runRelAvailabilityQuery(cfg, r, cfg.e.shops, autoAbility, availabilityParams, getAutoAbilitySourceIDs(cfg, ViewSourceTypeShop, &preAirship))
	if err != nil {
		return AutoAbility{}, err
	}

	postAirship := database.ShopTypePostAirship
	shopsPost, err := runRelAvailabilityQuery(cfg, r, cfg.e.shops, autoAbility, availabilityParams, getAutoAbilitySourceIDs(cfg, ViewSourceTypeShop, &postAirship))
	if err != nil {
		return AutoAbility{}, err
	}

	treasures, err := runRelAvailabilityQuery(cfg, r, cfg.e.treasures, autoAbility, availabilityParams, getAutoAbilitySourceIDs(cfg, ViewSourceTypeTreasure, nil))
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
