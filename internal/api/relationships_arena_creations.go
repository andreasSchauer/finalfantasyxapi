package api

import (
	"net/http"
	"sync"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func GetArenaCreationRelationships(r *http.Request, creation seedtypes.ArenaCreation) (apitypes.ArenaCreation, error) {
	var rel apitypes.ArenaCreation
	var err error
	
	subquest, _ := lookup.GetResourceByID(creation.SubquestID, lookup.L.SubquestsID)
	ia := subquest.Completion.Reward
	rel.Reward = apilookups.E.AllItems.ToResAmtFromName(ia)

	rel.RequiredMonsters, err = getArenaCreationMonsters(r, creation)
	if err != nil {
		return apitypes.ArenaCreation{}, err
	}

	return rel, nil
}



func getArenaCreationMonsters(r *http.Request, creation seedtypes.ArenaCreation) ([]apiresource.Named, error) {
	mf := createMonsterFilter(creation)

	if mf.IsZero() || mf.CreationsUnlockedCategory != nil {
		return nil, nil
	}

	monsterIDSlices := [][]int32{}
	g, ctx := errgroup.WithContext(r.Context())
	var mu sync.Mutex

	if mf.RequiredArea != nil {
		g.Go(func() error {
			area := database.ToNullMaCreationArea(mf.RequiredArea)
			idSlice, err := config.Cfg.DB.GetCaptureMonsterIDsByMaCreationArea(ctx, area)
			if err != nil {
				return err
			}

			mu.Lock()
			monsterIDSlices = append(monsterIDSlices, idSlice)
			mu.Unlock()
			return nil
		})
	}

	if mf.RequiredSpecies != nil {
		g.Go(func() error {
			species := database.MonsterSpecies(*mf.RequiredSpecies)
			idSlice, err := config.Cfg.DB.GetCaptureMonsterIDsBySpecies(ctx, species)
			if err != nil {
				return err
			}

			mu.Lock()
			monsterIDSlices = append(monsterIDSlices, idSlice)
			mu.Unlock()
			return nil
		})
	}

	if mf.UnderwaterOnly {
		g.Go(func() error {
			idSlice, err := config.Cfg.DB.GetCaptureMonsterIDsByIsUnderwater(ctx)
			if err != nil {
				return err
			}

			mu.Lock()
			monsterIDSlices = append(monsterIDSlices, idSlice)
			mu.Unlock()
			return nil
		})
	}

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	monsterIDs, err := ah.FilterIDSlices(monsterIDSlices)
	if err != nil {
		return nil, err
	}

	monsters := apilookups.E.Monsters.ToAPIResources(monsterIDs)
	return monsters, nil
}

type MonsterFilter struct {
	RequiredArea              *string
	RequiredSpecies           *string
	CreationsUnlockedCategory *string
	UnderwaterOnly            bool
}

func (mf MonsterFilter) IsZero() bool {
	return mf.RequiredArea == nil &&
		mf.RequiredSpecies == nil &&
		mf.CreationsUnlockedCategory == nil &&
		!mf.UnderwaterOnly
}

func createMonsterFilter(creation seedtypes.ArenaCreation) MonsterFilter {
	return MonsterFilter{
		RequiredArea:              creation.RequiredArea,
		RequiredSpecies:           creation.RequiredSpecies,
		CreationsUnlockedCategory: creation.CreationsUnlockedCategory,
		UnderwaterOnly:            creation.UnderwaterOnly,
	}
}
