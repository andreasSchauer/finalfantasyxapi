package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop7SeedArenaCreations(qtx *database.Queries, ctx context.Context) error {
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
		params.DataHash[i] = generateDataHash(c)
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
		rawJSON.monsterArenaCreations[i].ID = row.ID
		L.ArenaCreations[seedtypes.Key(creations[i])] = creations[i]
		L.ArenaCreationsID[row.ID] = creations[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractArenaCreations() ([]seedtypes.ArenaCreation, error) {
	creations := []seedtypes.ArenaCreation{}
	var err error

	for i := range rawJSON.monsterArenaCreations {
		creation := &rawJSON.monsterArenaCreations[i]

		creation.SubquestID, err = assignFK(creation.Name, L.Subquests)
		if err != nil {
			return nil, err
		}

		obj := seedtypes.LookupObject{
			Name: creation.Name,
		}

		creation.MonsterID, err = assignFKPtr(&obj, L.Monsters)
		if err != nil {
			return nil, err
		}

		creations = append(creations, *creation)
	}

	return dedupeRows(creations, hashes), nil
}
