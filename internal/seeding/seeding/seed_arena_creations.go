package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop7SeedArenaCreations(qtx *database.Queries, ctx context.Context) error {
	creations, err := extractArenaCreations()
	if err != nil {
		return err
	}

	params := database.CreateMonsterArenaCreationBulkParams{
		DataHash:                  make([]string, len(creations)),
		SubquestID:                make([]int32, len(creations)),
		Category:                  make([]database.MaCreationCategory, len(creations)),
		RequiredArea:              make([]database.NullMaCreationArea, len(creations)),
		RequiredSpecies:           make([]database.NullMaCreationSpecies, len(creations)),
		UnderwaterOnly:            make([]bool, len(creations)),
		CreationsUnlockedCategory: make([]database.NullCreationsUnlockedCategory, len(creations)),
		Amount:                    make([]int32, len(creations)),
		MonsterID:                 make([]sql.NullInt32, len(creations)),
	}

	for i, c := range creations {
		params.DataHash[i] = sh.GenerateDataHash(c)
		params.SubquestID[i] = c.SubquestID
		params.Category[i] = database.MaCreationCategory(c.Category)
		params.RequiredArea[i] = database.ToNullMaCreationArea(c.RequiredArea)
		params.RequiredSpecies[i] = database.ToNullMaCreationSpecies(c.RequiredSpecies)
		params.UnderwaterOnly[i] = c.UnderwaterOnly
		params.CreationsUnlockedCategory[i] = database.ToNullCreationsUnlockedCategory(c.CreationsUnlockedCategory)
		params.Amount[i] = c.Amount
		params.MonsterID[i] = h.GetNullInt32(c.MonsterID)
	}

	dbRows, err := qtx.CreateMonsterArenaCreationBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create arena creations: %v", err)
	}

	for i, row := range dbRows {
		creations[i].ID = row.ID
		sh.RawJSON.MonsterArenaCreations[i].ID = row.ID
		lookup.L.ArenaCreations[seedtypes.Key(creations[i])] = creations[i]
		lookup.L.ArenaCreationsID[row.ID] = creations[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractArenaCreations() ([]seedtypes.ArenaCreation, error) {
	creations := []seedtypes.ArenaCreation{}
	var err error

	for i := range sh.RawJSON.MonsterArenaCreations {
		creation := &sh.RawJSON.MonsterArenaCreations[i]

		creation.SubquestID, err = sh.AssignFK(creation.Name, lookup.L.Subquests)
		if err != nil {
			return nil, err
		}

		obj := seedtypes.LookupObject{
			Name: creation.Name,
		}

		creation.MonsterID, err = sh.AssignFKPtr(&obj, lookup.L.Monsters)
		if err != nil {
			return nil, err
		}

		creations = append(creations, *creation)
	}

	return dedupeRows(creations, sh.Hashes), nil
}
