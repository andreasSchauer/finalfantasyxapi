package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type BlitzballPosition struct {
	ID       int32
	Category string          `json:"category"`
	Slot     string          `json:"slot"`
	Items    []BlitzballItem `json:"items"`
}

func (b BlitzballPosition) ToHashFields() []any {
	return []any{
		b.Category,
		b.Slot,
	}
}

func (b BlitzballPosition) GetID() int32 {
	return b.ID
}

func (b BlitzballPosition) ToKeyFields() []any {
	return []any{
		b.Category,
		b.Slot,
	}
}

type BlitzballItem struct {
	PositionID int32
	ItemAmount ItemAmount `json:"item"`
	Chance     int32      `json:"chance"`
}

func (b BlitzballItem) ToHashFields() []any {
	return []any{
		b.PositionID,
		b.ItemAmount.ID,
		b.Chance,
	}
}

func (l *lookup) seedBlitzballItems(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/blitzball_items.json"

	var blitzballPositions []BlitzballPosition
	err := loadJSONFile(string(srcPath), &blitzballPositions)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, position := range blitzballPositions {
			dbPosition, err := qtx.CreateBlitzballPosition(context.Background(), database.CreateBlitzballPositionParams{
				DataHash: generateDataHash(position),
				Category: database.BlitzballTournamentCategory((position.Category)),
				Slot:     database.BlitzballPositionSlot(position.Slot),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Blitzball Position: %s: %v", createLookupKey(position), err)
			}
			position.ID = dbPosition.ID
			key := createLookupKey(position)
			l.positions[key] = position
		}
		return nil
	})
}

func (l *lookup) seedBlitzballItemsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/blitzball_items.json"

	var blitzballPositions []BlitzballPosition
	err := loadJSONFile(string(srcPath), &blitzballPositions)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonPosition := range blitzballPositions {
			key := createLookupKey(jsonPosition)
			position, err := l.getPosition(key)
			if err != nil {
				return err
			}

			for _, item := range position.Items {
				var err error
				item.PositionID = position.ID

				item.ItemAmount, err = seedObjAssignID(qtx, item.ItemAmount, l.seedItemAmount)
				if err != nil {
					return fmt.Errorf("%s: %v", key, err)
				}

				err = qtx.CreateBlitzballItem(context.Background(), database.CreateBlitzballItemParams{
					DataHash:     generateDataHash(item),
					PositionID:   item.PositionID,
					ItemAmountID: item.ItemAmount.ID,
					Chance:       item.Chance,
				})
				if err != nil {
					return fmt.Errorf("%s: couldn't create Blitzball Item: %v", key, err)
				}
			}
		}

		return nil
	})
}
