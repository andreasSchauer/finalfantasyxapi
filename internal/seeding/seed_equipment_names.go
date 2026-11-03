package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop5SeedEquipmentNames(qtx *database.Queries, ctx context.Context) error {
	names, err := extractEquipmentNames()
	if err != nil {
		return err
	}

	params := database.CreateEquipmentNameBulkParams{
		DataHash:    make([]string, len(names)),
		CharacterID: make([]int32, len(names)),
		Name:        make([]string, len(names)),
	}

	for i, n := range names {
		params.DataHash[i] = generateDataHash(n)
		params.CharacterID[i] = n.CharacterID
		params.Name[i] = n.Name
	}

	dbRows, err := qtx.CreateEquipmentNameBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create equipment names: %v", err)
	}

	for i, row := range dbRows {
		names[i].ID = row.ID
		L.EquipmentNames[seedtypes.Key(names[i])] = names[i]
		L.EquipmentNamesID[row.ID] = names[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractEquipmentNames() ([]seedtypes.EquipmentName, error) {
	names := []seedtypes.EquipmentName{}
	var err error

	for i := range rawJSON.equipment {
		table := &rawJSON.equipment[i]

		for j := range table.EquipmentNames {
			name := &table.EquipmentNames[j]

			name.CharacterID, err = assignFK(name.CharacterName, L.Characters)
			if err != nil {
				return nil, err
			}

			names = append(names, *name)
		}
	}

	return dedupeRows(names, hashes), nil
}
