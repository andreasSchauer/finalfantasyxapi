package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop4SeedTreasures(qtx *database.Queries, ctx context.Context) error {
	treasures, err := l.extractTreasures()
	if err != nil {
		return err
	}

	params := database.CreateTreasureBulkParams{
		DataHash:        make([]string, len(treasures)),
		AreaID:          make([]int32, len(treasures)),
		Version:         make([]int32, len(treasures)),
		TreasureType:    make([]database.TreasureType, len(treasures)),
		LootType:        make([]database.LootType, len(treasures)),
		Availability:    make([]database.AvailabilityType, len(treasures)),
		IsAnimaTreasure: make([]bool, len(treasures)),
		Notes:           make([]sql.NullString, len(treasures)),
		GilAmount:       make([]sql.NullInt32, len(treasures)),
	}

	for i, t := range treasures {
		params.DataHash[i] = generateDataHash(t)
		params.AreaID[i] = t.AreaID
		params.Version[i] = t.Version
		params.TreasureType[i] = database.TreasureType(t.TreasureType)
		params.LootType[i] = database.LootType(t.LootType)
		params.Availability[i] = database.AvailabilityType(t.Availability)
		params.IsAnimaTreasure[i] = t.IsAnimaTreasure
		params.Notes[i] = h.GetNullString(t.Notes)
		params.GilAmount[i] = h.GetNullInt32(t.GilAmount)
	}

	dbRows, err := qtx.CreateTreasureBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create treasures: %v", err)
	}

	for i, row := range dbRows {
		treasures[i].ID = row.ID
		l.Treasures[Key(treasures[i])] = treasures[i]
		l.TreasuresID[row.ID] = treasures[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractTreasures() ([]Treasure, error) {
	treasures := []Treasure{}
	var err error

	for i := range l.json.treasureLists {
		list := &l.json.treasureLists[i]

		for j := range list.Treasures {
			treasure := &list.Treasures[j]
			treasure.Version = int32(j + 1)

			treasure.AreaID, err = assignFK(list.LocationArea, l.Areas)
			if err != nil {
				return nil, err
			}

			treasures = append(treasures, *treasure)
		}
	}

	return dedupeRows(treasures, l.Hashes), nil
}

func (l *Lookup) completeTreasureLists() error {
	for i := range l.json.treasureLists {
		var err error
		list := &l.json.treasureLists[i]

		err = l.completeTreasures(list.Treasures)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lookup) completeTreasures(treasures []Treasure) error {
	for i := range treasures {
		treasure := &treasures[i]

		err := assignIDs(l, treasure.Items)
		if err != nil {
			return err
		}

		err = l.assignID(treasure)
		if err != nil {
			return err
		}

		if treasure.Equipment != nil {
			err := l.assignID(treasure.Equipment)
			if err != nil {
				return err
			}
		}

		l.Treasures[Key(*treasure)] = *treasure
		l.TreasuresID[treasure.ID] = *treasure
	}

	return nil
}
