package seedhandlers

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

func Loop4SeedTreasures(qtx *database.Queries, ctx context.Context) error {
	treasures, err := extractTreasures()
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
		params.DataHash[i] = sh.GenerateDataHash(t)
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
		lookup.L.Treasures[seedtypes.Key(treasures[i])] = treasures[i]
		lookup.L.TreasuresID[row.ID] = treasures[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractTreasures() ([]seedtypes.Treasure, error) {
	treasures := []seedtypes.Treasure{}
	var err error

	for i := range sh.RawJSON.TreasureLists {
		list := &sh.RawJSON.TreasureLists[i]

		for j := range list.Treasures {
			treasure := &list.Treasures[j]
			treasure.Version = int32(j + 1)

			treasure.AreaID, err = sh.AssignFK(list.LocationArea, lookup.L.Areas)
			if err != nil {
				return nil, err
			}

			treasures = append(treasures, *treasure)
		}
	}

	return sh.DedupeRows(treasures, sh.Hashes), nil
}

func CompleteTreasureLists() error {
	for i := range sh.RawJSON.TreasureLists {
		var err error
		list := &sh.RawJSON.TreasureLists[i]

		err = completeTreasures(list.Treasures)
		if err != nil {
			return err
		}
	}

	return nil
}

func completeTreasures(treasures []seedtypes.Treasure) error {
	for i := range treasures {
		treasure := &treasures[i]

		err := assignIDs(treasure.Items)
		if err != nil {
			return err
		}

		err = assignID(treasure)
		if err != nil {
			return err
		}

		if treasure.Equipment != nil {
			err := assignID(treasure.Equipment)
			if err != nil {
				return err
			}
		}

		lookup.L.Treasures[seedtypes.Key(*treasure)] = *treasure
		lookup.L.TreasuresID[treasure.ID] = *treasure
	}

	return nil
}
