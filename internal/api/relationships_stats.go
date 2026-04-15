package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type StatQueries struct {
	AutoAbilities		DbQueryIntMany
	PlayerAbilities		DbQueryIntMany
	OverdriveAbilities 	DbQueryIntMany
	ItemAbilities		DbQueryIntMany
	TriggerCommands		DbQueryIntMany
	StatusConditions	DbQueryIntMany
	Properties			DbQueryIntMany
}

func getStatRelationships(cfg *Config, r *http.Request, stat seeding.Stat) (Stat, error) {
	queries, err := getStatQueries(cfg, r)
	if err != nil {
		return Stat{}, err
	}

	spheres, err := getResourcesDbItem(cfg, r, cfg.e.spheres, stat, cfg.db.GetStatSphereIDs)
	if err != nil {
		return Stat{}, err
	}

	autoAbilities, err := getResourcesDbItem(cfg, r, cfg.e.autoAbilities, stat, queries.AutoAbilities)
	if err != nil {
		return Stat{}, err
	}

	playerAbilities, err := getResourcesDbItem(cfg, r, cfg.e.playerAbilities, stat, queries.PlayerAbilities)
	if err != nil {
		return Stat{}, err
	}

	overdriveAbilities, err := getResourcesDbItem(cfg, r, cfg.e.overdriveAbilities, stat, queries.OverdriveAbilities)
	if err != nil {
		return Stat{}, err
	}

	itemAbilities, err := getResourcesDbItem(cfg, r, cfg.e.itemAbilities, stat, queries.ItemAbilities)
	if err != nil {
		return Stat{}, err
	}

	triggerCommands, err := getResourcesDbItem(cfg, r, cfg.e.triggerCommands, stat, queries.TriggerCommands)
	if err != nil {
		return Stat{}, err
	}

	statusConditions, err := getResourcesDbItem(cfg, r, cfg.e.statusConditions, stat, queries.StatusConditions)
	if err != nil {
		return Stat{}, err
	}

	properties, err := getResourcesDbItem(cfg, r, cfg.e.properties, stat, queries.Properties)
	if err != nil {
		return Stat{}, err
	}

	rel := Stat{
		Spheres: 			spheres,
		AutoAbilities: 		autoAbilities,
		PlayerAbilities: 	playerAbilities,
		OverdriveAbilities: overdriveAbilities,
		ItemAbilities: 		itemAbilities,
		TriggerCommands: 	triggerCommands,
		StatusConditions: 	statusConditions,
		Properties: 		properties,
	}

	rel.Spheres = spheres

	return rel, nil
}


func getStatQueries(cfg *Config, r *http.Request) (StatQueries, error) {
	changesOnly, err := parseBooleanQuery(r, cfg.q.stats["changes_only"])
	if errIsNotEmptyQuery(err) {
		return StatQueries{}, err
	}

	var queries StatQueries
	
	if changesOnly {
		queries = StatQueries{
			AutoAbilities: 		cfg.db.GetStatAutoAbilityIDsStatChange,
			PlayerAbilities: 	cfg.db.GetStatPlayerAbilityIDsStatChange,
			OverdriveAbilities: cfg.db.GetStatOverdriveAbilityIDsStatChange,
			ItemAbilities: 		cfg.db.GetStatItemAbilityIDsStatChange,
			TriggerCommands: 	cfg.db.GetStatTriggerCommandIDsStatChange,
			StatusConditions: 	cfg.db.GetStatStatusConditionIDsStatChange,
			Properties: 		cfg.db.GetStatPropertyIDsStatChange,
		}
	} else {
		queries = StatQueries{
			AutoAbilities: 		cfg.db.GetStatAutoAbilityIDs,
			PlayerAbilities: 	cfg.db.GetStatPlayerAbilityIDs,
			OverdriveAbilities: cfg.db.GetStatOverdriveAbilityIDs,
			ItemAbilities: 		cfg.db.GetStatItemAbilityIDs,
			TriggerCommands: 	cfg.db.GetStatTriggerCommandIDs,
			StatusConditions: 	cfg.db.GetStatStatusConditionIDs,
			Properties: 		cfg.db.GetStatPropertyIDs,
		}
	}

	return queries, nil
}