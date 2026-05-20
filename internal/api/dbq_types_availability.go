package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AvailabilityDbQuery func(ctx context.Context, p AvailabilityParams) ([]int32, error)
type AvailabilityDbQueryBool func(ctx context.Context, p AvailabilityBoolParams) ([]int32, error)

type AvailabilityBoolParams struct {
	Boolean      bool
	Availability []database.AvailabilityType
}


func getMasterItemIDsByLocation(cfg *Config, r *http.Request, id int32) ([]TypedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetMasterItemIDsByLocation(ctx, database.GetMasterItemIDsByLocationParams{
			LocationID:   p.ParentID,
			Availability: p.Availability,
			Method:       p.Method,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.allItems, id, cfg.e.locations.resourceType, dbQuery)
}

func getMasterItemIDsBySublocation(cfg *Config, r *http.Request, id int32) ([]TypedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetMasterItemIDsBySublocation(ctx, database.GetMasterItemIDsBySublocationParams{
			SublocationID: p.ParentID,
			Availability:  p.Availability,
			Method:        p.Method,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.allItems, id, cfg.e.sublocations.resourceType, dbQuery)
}

func getMasterItemIDsByArea(cfg *Config, r *http.Request, id int32) ([]TypedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetMasterItemIDsByArea(ctx, database.GetMasterItemIDsByAreaParams{
			AreaID:       p.ParentID,
			Availability: p.Availability,
			Method:       p.Method,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.allItems, id, cfg.e.areas.resourceType, dbQuery)
}



func getItemIDsByLocation(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetItemIDsByLocation(ctx, database.GetItemIDsByLocationParams{
			LocationID:   p.ParentID,
			Availability: p.Availability,
			Method:       p.Method,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.items, id, cfg.e.locations.resourceType, dbQuery)
}

func getItemIDsBySublocation(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetItemIDsBySublocation(ctx, database.GetItemIDsBySublocationParams{
			SublocationID: p.ParentID,
			Availability:  p.Availability,
			Method:        p.Method,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.items, id, cfg.e.sublocations.resourceType, dbQuery)
}

func getItemIDsByArea(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetItemIDsByArea(ctx, database.GetItemIDsByAreaParams{
			AreaID:       p.ParentID,
			Availability: p.Availability,
			Method:       p.Method,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.items, id, cfg.e.areas.resourceType, dbQuery)
}

func getMonsterIDsByArea(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetMonsterIDsByArea(ctx, database.GetMonsterIDsByAreaParams{
			AreaID:       p.ParentID,
			Availability: p.Availability,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.monsters, id, cfg.e.areas.resourceType, dbQuery)
}

func getMonsterIDsBySublocation(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetMonsterIDsBySublocation(ctx, database.GetMonsterIDsBySublocationParams{
			SublocationID:  p.ParentID,
			Availability: 	p.Availability,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.monsters, id, cfg.e.sublocations.resourceType, dbQuery)
}

func getMonsterIDsByLocation(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetMonsterIDsByLocation(ctx, database.GetMonsterIDsByLocationParams{
			LocationID:  	p.ParentID,
			Availability: 	p.Availability,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.monsters, id, cfg.e.locations.resourceType, dbQuery)
}

func getShopIDsBySublocation(cfg *Config, r *http.Request, id int32) ([]UnnamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetShopIDsBySublocation(ctx, database.GetShopIDsBySublocationParams{
			SublocationID:  p.ParentID,
			Availability: 	p.Availability,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.shops, id, cfg.e.sublocations.resourceType, dbQuery)
}

func getShopIDsByLocation(cfg *Config, r *http.Request, id int32) ([]UnnamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetShopIDsByLocation(ctx, database.GetShopIDsByLocationParams{
			LocationID:  	p.ParentID,
			Availability: 	p.Availability,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.shops, id, cfg.e.locations.resourceType, dbQuery)
}

func getShopIDsWithItems(cfg *Config, r *http.Request, boolean bool) ([]UnnamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityBoolParams) ([]int32, error) {
		return cfg.db.GetShopIDsWithItems(ctx, database.GetShopIDsWithItemsParams{
			HasItems: 		p.Boolean,
			Availability: 	p.Availability,
		})
	}

	return runAvlBoolQuery(cfg, r, cfg.e.shops, boolean, dbQuery)
}

func getShopIDsWithEquipment(cfg *Config, r *http.Request, boolean bool) ([]UnnamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityBoolParams) ([]int32, error) {
		return cfg.db.GetShopIDsWithEquipment(ctx, database.GetShopIDsWithEquipmentParams{
			HasEquipment: 	p.Boolean,
			Availability: 	p.Availability,
		})
	}

	return runAvlBoolQuery(cfg, r, cfg.e.shops, boolean, dbQuery)
}

func runAvlBoolQuery[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput [T, R, A, L], boolean bool, dbQuery AvailabilityDbQueryBool) ([]A, error) {
	queryParam := i.queryLookup["availability"]

	availabilities, err := parseEnumListQuery(cfg, r, i.endpoint, queryParam, cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}

	params := AvailabilityBoolParams{
		Boolean: 		boolean,
		Availability: 	availabilities,
	}

	dbIDs, err := dbQuery(r.Context(), params)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss.", i.resourceType), err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}

func getMonsterFormationIDsByMonster(cfg *Config, r *http.Request, id int32) ([]UnnamedAPIResource, error) {
	i := cfg.e.monsterFormations
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		_, err := checkEmptyQuery(r, i.queryLookup["area"])
		if err != nil {
			p.Availability = nil
		}

		return cfg.db.GetMonsterFormationIDsByMonster(ctx, database.GetMonsterFormationIDsByMonsterParams{
			MonsterID:  	p.ParentID,
			Availability: 	p.Availability,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.monsterFormations, id, cfg.e.monsters.resourceType, dbQuery)
}

func getMonsterFormationIDsByLocation(cfg *Config, r *http.Request, id int32) ([]UnnamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetMonsterFormationIDsByLocation(ctx, database.GetMonsterFormationIDsByLocationParams{
			LocationID:  	p.ParentID,
			Availability: 	p.Availability,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.monsterFormations, id, cfg.e.locations.resourceType, dbQuery)
}

func getMonsterFormationIDsBySublocation(cfg *Config, r *http.Request, id int32) ([]UnnamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetMonsterFormationIDsBySublocation(ctx, database.GetMonsterFormationIDsBySublocationParams{
			SublocationID:  p.ParentID,
			Availability: 	p.Availability,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.monsterFormations, id, cfg.e.sublocations.resourceType, dbQuery)
}

func getMonsterFormationIDsByArea(cfg *Config, r *http.Request, id int32) ([]UnnamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetMonsterFormationIDsByArea(ctx, database.GetMonsterFormationIDsByAreaParams{
			AreaID:  		p.ParentID,
			Availability: 	p.Availability,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.monsterFormations, id, cfg.e.areas.resourceType, dbQuery)
}