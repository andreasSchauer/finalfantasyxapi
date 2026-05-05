package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop5SeedEquipmentNames(qtx *database.Queries, ctx context.Context) error {
	names, err := l.extractEquipmentNames()
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
		l.EquipmentNames[Key(names[i])] = names[i]
		l.EquipmentNamesID[row.ID] = names[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractEquipmentNames() ([]EquipmentName, error) {
	names := []EquipmentName{}
	var err error

	for i := range l.json.equipment {
		table := &l.json.equipment[i]

		for j := range table.EquipmentNames {
			name := &table.EquipmentNames[j]

			name.CharacterID, err = assignFK(name.CharacterName, l.Characters)
			if err != nil {
				return nil, err
			}

			names = append(names, *name)
		}
	}

	return dedupeRows(names, l.Hashes), nil
}
