package api

import (
	"net/http"
	"sync"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func (cfg *Config) getArenaCreation(r *http.Request, i handlerInput[seeding.ArenaCreation, ArenaCreation, NamedAPIResource, NamedApiResourceList], id int32) (ArenaCreation, error) {
	creation, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return ArenaCreation{}, err
	}

	subquest, _ := seeding.GetResourceByID(creation.SubquestID, cfg.l.SubquestsID)
	ia := subquest.Completion.Reward

	monsters, err := getArenaCreationMonsters(cfg, r, creation)
	if err != nil {
		return ArenaCreation{}, err
	}

	response := ArenaCreation{
		ID:                        creation.ID,
		Name:                      creation.Name,
		Category:                  creation.Category,
		Monster:                   idToNamedAPIResource(cfg, cfg.e.monsters, *creation.MonsterID),
		ParentSubquest:            idToQuestAPIResource(cfg, cfg.e.subquests, creation.SubquestID),
		Reward:                    nameAmountToResourceAmount(cfg, cfg.e.allItems, ia),
		RequiredCatchAmount:       creation.Amount,
		UnlockedCreationsCategory: creation.CreationsUnlockedCategory,
		RequiredMonsters:          monsters,
	}

	return response, nil
}

func (cfg *Config) retrieveArenaCreations(r *http.Request, i handlerInput[seeding.ArenaCreation, ArenaCreation, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		enumListQuery(cfg, r, i, cfg.t.ArenaCreationCategory, ids, qpnCategory, cfg.db.GetArenaCreationIDsByCategory),
	})
}

func getArenaCreationMonsters(cfg *Config, r *http.Request, creation seeding.ArenaCreation) ([]NamedAPIResource, error) {
	mf := createMonsterFilter(creation)

	if mf.IsZero() || mf.CreationsUnlockedCategory != nil {
		return nil, nil
	}

	monsterIdSlices := [][]int32{}
	g, ctx := errgroup.WithContext(r.Context())
	var mu sync.Mutex

	if mf.RequiredArea != nil {
		g.Go(func() error {
			area := database.ToNullMaCreationArea(mf.RequiredArea)
			idSlice, err := cfg.db.GetCaptureMonsterIDsByMaCreationArea(ctx, area)
			if err != nil {
				return err
			}

			mu.Lock()
			monsterIdSlices = append(monsterIdSlices, idSlice)
			mu.Unlock()
			return nil
		})
	}

	if mf.RequiredSpecies != nil {
		g.Go(func() error {
			species := database.MonsterSpecies(*mf.RequiredSpecies)
			idSlice, err := cfg.db.GetCaptureMonsterIDsBySpecies(ctx, species)
			if err != nil {
				return err
			}

			mu.Lock()
			monsterIdSlices = append(monsterIdSlices, idSlice)
			mu.Unlock()
			return nil
		})
	}

	if mf.UnderwaterOnly {
		g.Go(func() error {
			idSlice, err := cfg.db.GetCaptureMonsterIDsByIsUnderwater(ctx)
			if err != nil {
				return err
			}

			mu.Lock()
			monsterIdSlices = append(monsterIdSlices, idSlice)
			mu.Unlock()
			return nil
		})
	}

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	monsterIDs, err := filterIdSlices(monsterIdSlices)
	if err != nil {
		return nil, err
	}

	monsters := idsToAPIResources(cfg, cfg.e.monsters, monsterIDs)
	return monsters, nil
}
