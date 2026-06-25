package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
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

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumListQuery(cfg, r, i, cfg.t.ArenaCreationCategory, ids, qpnCategory, cfg.db.GetArenaCreationIDsByCategory)),
	})
}

func getArenaCreationMonsters(cfg *Config, r *http.Request, creation seeding.ArenaCreation) ([]NamedAPIResource, error) {
	mf := createMonsterFilter(creation)

	if mf.IsZero() || mf.CreationsUnlockedCategory != nil {
		return nil, nil
	}

	monsterIdSlices := []filteredIdList{}

	if mf.RequiredArea != nil {
		area := database.ToNullMaCreationArea(mf.RequiredArea)
		idSlice := fidl(cfg.db.GetCaptureMonsterIDsByMaCreationArea(r.Context(), area))
		monsterIdSlices = append(monsterIdSlices, idSlice)
	}

	if mf.RequiredSpecies != nil {
		species := database.MonsterSpecies(*mf.RequiredSpecies)
		idSlice := fidl(cfg.db.GetCaptureMonsterIDsBySpecies(r.Context(), species))
		monsterIdSlices = append(monsterIdSlices, idSlice)
	}

	if mf.UnderwaterOnly {
		idSlice := fidl(cfg.db.GetCaptureMonsterIDsByIsUnderwater(r.Context()))
		monsterIdSlices = append(monsterIdSlices, idSlice)
	}

	monsterIDs, err := filterIdSlices(monsterIdSlices)
	if err != nil {
		return nil, err
	}

	monsters := idsToAPIResources(cfg, cfg.e.monsters, monsterIDs)

	return monsters, nil
}
