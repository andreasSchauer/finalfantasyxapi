package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

type StatQueries struct {
	AutoAbilities      DbQueryIntMany
	PlayerAbilities    DbQueryIntMany
	OverdriveAbilities DbQueryIntMany
	ItemAbilities      DbQueryIntMany
	TriggerCommands    DbQueryIntMany
	StatusConditions   DbQueryIntMany
	Properties         DbQueryIntMany
}

func getStatRelationships(cfg *Config, r *http.Request, stat seeding.Stat) (Stat, error) {
	var rel Stat
	g, ctx := errgroup.WithContext(r.Context())

	queries, err := getStatQueries(cfg, r)
	if err != nil {
		return Stat{}, err
	}

	g.Go(func() error {
		var err error
		rel.Spheres, err = getResourcesDbItem(cfg, ctx, cfg.E.Spheres, stat, cfg.db.GetStatSphereIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = getResourcesDbItem(cfg, ctx, cfg.E.AutoAbilities, stat, queries.AutoAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = getResourcesDbItem(cfg, ctx, cfg.E.PlayerAbilities, stat, queries.PlayerAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = getResourcesDbItem(cfg, ctx, cfg.E.OverdriveAbilities, stat, queries.OverdriveAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = getResourcesDbItem(cfg, ctx, cfg.E.ItemAbilities, stat, queries.ItemAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.TriggerCommands, err = getResourcesDbItem(cfg, ctx, cfg.E.TriggerCommands, stat, queries.TriggerCommands)
		return err
	})

	g.Go(func() error {
		var err error
		rel.StatusConditions, err = getResourcesDbItem(cfg, ctx, cfg.E.StatusConditions, stat, queries.StatusConditions)
		return err
	})

	g.Go(func() error {
		var err error
		if queries.Properties != nil {
			rel.Properties, err = getResourcesDbItem(cfg, ctx, cfg.E.Properties, stat, queries.Properties)
			return err
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return Stat{}, err
	}

	return rel, nil
}

func getStatQueries(cfg *Config, r *http.Request) (StatQueries, error) {
	changesOnly, err := parseBooleanQuery(r, cfg.Q.stats[QpnChangesOnly])
	if errExceptEmptyQuery(err) {
		return StatQueries{}, err
	}

	var queries StatQueries

	if changesOnly {
		queries = StatQueries{
			AutoAbilities:      cfg.db.GetStatAutoAbilityIDsStatChange,
			PlayerAbilities:    cfg.db.GetStatPlayerAbilityIDsStatChange,
			OverdriveAbilities: cfg.db.GetStatOverdriveAbilityIDsStatChange,
			ItemAbilities:      cfg.db.GetStatItemAbilityIDsStatChange,
			TriggerCommands:    cfg.db.GetStatTriggerCommandIDsStatChange,
			StatusConditions:   cfg.db.GetStatStatusConditionIDsStatChange,
			Properties:         nil,
		}
	} else {
		queries = StatQueries{
			AutoAbilities:      cfg.db.GetStatAutoAbilityIDs,
			PlayerAbilities:    cfg.db.GetStatPlayerAbilityIDs,
			OverdriveAbilities: cfg.db.GetStatOverdriveAbilityIDs,
			ItemAbilities:      cfg.db.GetStatItemAbilityIDs,
			TriggerCommands:    cfg.db.GetStatTriggerCommandIDs,
			StatusConditions:   cfg.db.GetStatStatusConditionIDs,
			Properties:         cfg.db.GetStatPropertyIDs,
		}
	}

	return queries, nil
}
