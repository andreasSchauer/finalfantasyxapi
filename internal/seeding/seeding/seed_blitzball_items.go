package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func loop1SeedBlitzballPositions(qtx *database.Queries, ctx context.Context) error {
	positions := dedupeRows(rawJSON.blitzballPositions, hashes)

	params := database.CreateBlitzballPositionBulkParams{
		DataHash: make([]string, len(positions)),
		Category: make([]database.BlitzballTournamentCategory, len(positions)),
		Slot:     make([]database.BlitzballPositionSlot, len(positions)),
	}

	for i, p := range positions {
		params.DataHash[i] = generateDataHash(p)
		params.Category[i] = database.BlitzballTournamentCategory(p.Category)
		params.Slot[i] = database.BlitzballPositionSlot(p.Slot)
	}

	dbRows, err := qtx.CreateBlitzballPositionBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create blitzball position: %v", err)
	}

	for i, row := range dbRows {
		positions[i].ID = row.ID
		rawJSON.blitzballPositions[i].ID = row.ID
		lookup.L.Positions[seedtypes.Key(positions[i])] = positions[i]
		lookup.L.PositionsID[row.ID] = positions[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func loop4SeedBlitzballItems(qtx *database.Queries, ctx context.Context) error {
	items, err := extractBlitzballItems()
	if err != nil {
		return err
	}

	params := database.CreateBlitzballItemBulkParams{
		DataHash:       make([]string, len(items)),
		PositionID:     make([]int32, len(items)),
		PossibleItemID: make([]int32, len(items)),
	}

	for i, bi := range items {
		params.DataHash[i] = generateDataHash(bi)
		params.PositionID[i] = bi.PositionID
		params.PossibleItemID[i] = bi.PossibleItem.ID
	}

	dbRows, err := qtx.CreateBlitzballItemBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create blitzball items: %v", err)
	}

	for _, row := range dbRows {
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractBlitzballItems() ([]seedtypes.BlitzballItem, error) {
	items := []seedtypes.BlitzballItem{}
	var err error

	for i := range rawJSON.blitzballPositions {
		pos := &rawJSON.blitzballPositions[i]

		for j := range pos.Items {
			item := &pos.Items[j]
			item.PositionID = pos.ID

			item.PossibleItem.ID, err = getHashID(item.PossibleItem)
			if err != nil {
				return nil, err
			}

			items = append(items, *item)
		}
	}

	return dedupeRows(items, hashes), nil
}

func completeBlitzballPositions() error {
	for i := range rawJSON.blitzballPositions {
		pos := &rawJSON.blitzballPositions[i]
		err := assignIDs(pos.Items)
		if err != nil {
			return err
		}

		lookup.L.Positions[seedtypes.Key(*pos)] = *pos
		lookup.L.PositionsID[pos.ID] = *pos
	}

	return nil
}
