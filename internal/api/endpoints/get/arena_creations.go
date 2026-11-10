package get

import (
	"net/http"
	"sync"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"golang.org/x/sync/errgroup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func GetArenaCreation(r *http.Request, i handlerinput.Main[seedtypes.ArenaCreation, apitypes.ArenaCreation, apiresource.Named], id int32) (apitypes.ArenaCreation, error) {
	creation, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.ArenaCreation{}, err
	}

	rel, err := getArenaCreationRelationships(r, creation)
	if err != nil {
		return apitypes.ArenaCreation{}, err
	}

	response := apitypes.ArenaCreation{
		ID:                        creation.ID,
		Name:                      creation.Name,
		Category:                  creation.Category,
		Monster:                   handlerinput.IDToNamedAPIResource(apilookups.E.Monsters, *creation.MonsterID),
		ParentSubquest:            handlerinput.IDToQuestAPIResource(apilookups.E.Subquests, creation.SubquestID),
		Reward:                    rel.Reward,
		RequiredCatchAmount:       creation.Amount,
		UnlockedCreationsCategory: creation.CreationsUnlockedCategory,
		RequiredMonsters:          rel.RequiredMonsters,
	}

	return response, nil
}

func getArenaCreationRelationships(r *http.Request, creation seedtypes.ArenaCreation) (apitypes.ArenaCreation, error) {
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

	if mf.IsZero() || mf.creationsUnlockedCategory != nil {
		return nil, nil
	}

	monsterIDSlices := [][]int32{}
	g, ctx := errgroup.WithContext(r.Context())
	var mu sync.Mutex

	if mf.requiredArea != nil {
		g.Go(func() error {
			area := database.ToNullMaCreationArea(mf.requiredArea)
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

	if mf.requiredSpecies != nil {
		g.Go(func() error {
			species := database.MonsterSpecies(*mf.requiredSpecies)
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

	if mf.underwaterOnly {
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

type monsterFilter struct {
	requiredArea              *string
	requiredSpecies           *string
	creationsUnlockedCategory *string
	underwaterOnly            bool
}

func (mf monsterFilter) IsZero() bool {
	return mf.requiredArea == nil &&
		mf.requiredSpecies == nil &&
		mf.creationsUnlockedCategory == nil &&
		!mf.underwaterOnly
}

func createMonsterFilter(creation seedtypes.ArenaCreation) monsterFilter {
	return monsterFilter{
		requiredArea:              creation.RequiredArea,
		requiredSpecies:           creation.RequiredSpecies,
		creationsUnlockedCategory: creation.CreationsUnlockedCategory,
		underwaterOnly:            creation.UnderwaterOnly,
	}
}
