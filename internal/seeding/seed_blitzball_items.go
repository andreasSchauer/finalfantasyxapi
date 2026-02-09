package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
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

func (b BlitzballPosition) Error() string {
	return fmt.Sprintf("blitzball position %s in %s", b.Slot, b.Category)
}

func (b BlitzballPosition) GetResParamsUnnamed() h.ResParamsUnnamed {
	return h.ResParamsUnnamed{
		ID: b.ID,
	}
}

type PossibleItem struct {
	ID         int32
	ItemAmount ItemAmount `json:"item"`
	Chance     int32      `json:"chance"`
}

func (i PossibleItem) ToHashFields() []any {
	return []any{
		i.ItemAmount.ID,
		i.Chance,
	}
}

func (i PossibleItem) GetID() int32 {
	return i.ID
}

func (i PossibleItem) Error() string {
	return fmt.Sprintf("possible item %s, amount: %d, chance: %d", i.ItemAmount.ItemName, i.ItemAmount.Amount, i.Chance)
}

type BlitzballItem struct {
	PositionID int32
	PossibleItem
}

func (b BlitzballItem) ToHashFields() []any {
	return []any{
		b.PositionID,
		b.PossibleItem.ID,
	}
}

func (b BlitzballItem) Error() string {
	return fmt.Sprintf("blitzball item %s, chance %d, position id %d", b.ItemAmount.ItemName, b.Chance, b.PositionID)
}

func (l *Lookup) seedBlitzballItems(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/blitzball_items.json"

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
				return h.NewErr(position.Error(), err, "couldn't create blitzball position")
			}
			position.ID = dbPosition.ID
			key := CreateLookupKey(position)
			l.Positions[key] = position
			l.PositionsID[position.ID] = position
		}
		return nil
	})
}

func (l *Lookup) seedBlitzballItemsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/blitzball_items.json"

	var blitzballPositions []BlitzballPosition
	err := loadJSONFile(string(srcPath), &blitzballPositions)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonPosition := range blitzballPositions {
			key := CreateLookupKey(jsonPosition)
			position, err := GetResource(key, l.Positions)
			if err != nil {
				return err
			}

			for _, item := range position.Items {
				var err error
				item.PositionID = position.ID

				item.PossibleItem, err = seedObjAssignID(qtx, item.PossibleItem, l.seedPossibleItem)
				if err != nil {
					return h.NewErr(item.Error(), err)
				}

				err = qtx.CreateBlitzballItem(context.Background(), database.CreateBlitzballItemParams{
					DataHash:       generateDataHash(item),
					PositionID:     item.PositionID,
					PossibleItemID: item.PossibleItem.ID,
				})
				if err != nil {
					return h.NewErr(item.Error(), err, "couldn't create blitzball item")
				}
			}
		}

		return nil
	})
}

func (l *Lookup) seedPossibleItem(qtx *database.Queries, item PossibleItem) (PossibleItem, error) {
	var err error

	item.ItemAmount, err = seedObjAssignID(qtx, item.ItemAmount, l.seedItemAmount)
	if err != nil {
		return PossibleItem{}, h.NewErr(item.Error(), err)
	}

	dbPossibleItem, err := qtx.CreatePossibleItem(context.Background(), database.CreatePossibleItemParams{
		DataHash:     generateDataHash(item),
		ItemAmountID: item.ItemAmount.ID,
		Chance:       item.Chance,
	})
	if err != nil {
		return PossibleItem{}, h.NewErr(item.Error(), err, "couldn't create possible item")
	}

	item.ID = dbPossibleItem.ID

	return item, nil
}
