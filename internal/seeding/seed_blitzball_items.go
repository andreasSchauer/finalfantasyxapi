package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop1SeedBlitzballPositions(qtx *database.Queries, ctx context.Context) error {
	positions := dedupeRows(l.json.blitzballPositions, l.Hashes)

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
		l.json.blitzballPositions[i].ID = row.ID
		l.Positions[Key(positions[i])] = positions[i]
		l.PositionsID[row.ID] = positions[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) loop4SeedBlitzballItems(qtx *database.Queries, ctx context.Context) error {
	items, err := l.extractBlitzballItems()
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
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractBlitzballItems() ([]BlitzballItem, error) {
	items := []BlitzballItem{}
	var err error

	for i := range l.json.blitzballPositions {
		pos := &l.json.blitzballPositions[i]

		for j := range pos.Items {
			item := &pos.Items[j]
			item.PositionID = pos.ID

			item.PossibleItem.ID, err = l.getHashID(item.PossibleItem)
			if err != nil {
				return nil, err
			}

			items = append(items, *item)
		}
	}

	return dedupeRows(items, l.Hashes), nil
}

func (l *Lookup) completeBlitzballPositions() error {
	for i := range l.json.blitzballPositions {
		pos := &l.json.blitzballPositions[i]
		err := assignIDs(l, pos.Items)
		if err != nil {
			return err
		}

		l.Positions[Key(*pos)] = *pos
		l.PositionsID[pos.ID] = *pos
	}

	return nil
}
