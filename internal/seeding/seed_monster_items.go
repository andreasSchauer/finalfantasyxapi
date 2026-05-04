package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop3SeedMonsterItems(qtx *database.Queries, ctx context.Context) error {
	items, err := l.extractMonsterItems()
	if err != nil {
		return err
	}

	params := database.CreateMonsterItemBulkParams{
		DataHash:              make([]string, len(items)),
		MonsterID:             make([]int32, len(items)),
		DropChance:            make([]int32, len(items)),
		DropCondition:         make([]sql.NullString, len(items)),
		OtherItemsCondition:   make([]sql.NullString, len(items)),
		StealCommonID:         make([]sql.NullInt32, len(items)),
		StealRareID:           make([]sql.NullInt32, len(items)),
		DropCommonID:          make([]sql.NullInt32, len(items)),
		DropRareID:            make([]sql.NullInt32, len(items)),
		SecondaryDropCommonID: make([]sql.NullInt32, len(items)),
		SecondaryDropRareID:   make([]sql.NullInt32, len(items)),
		BribeID:               make([]sql.NullInt32, len(items)),
	}

	for i, mi := range items {
		params.DataHash[i] = generateDataHash(mi)
		params.MonsterID[i] = mi.MonsterID
		params.DropChance[i] = mi.DropChance
		params.DropCondition[i] = h.GetNullString(mi.DropCondition)
		params.OtherItemsCondition[i] = h.GetNullString(mi.OtherItemsCondition)
		params.StealCommonID[i] = h.ObjPtrToNullInt32ID(mi.StealCommon)
		params.StealRareID[i] = h.ObjPtrToNullInt32ID(mi.StealRare)
		params.DropCommonID[i] = h.ObjPtrToNullInt32ID(mi.DropCommon)
		params.DropRareID[i] = h.ObjPtrToNullInt32ID(mi.DropRare)
		params.SecondaryDropCommonID[i] = h.ObjPtrToNullInt32ID(mi.SecondaryDropCommon)
		params.SecondaryDropRareID[i] = h.ObjPtrToNullInt32ID(mi.SecondaryDropRare)
		params.BribeID[i] = h.ObjPtrToNullInt32ID(mi.Bribe)
	}

	dbRows, err := qtx.CreateMonsterItemBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create monster items: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractMonsterItems() ([]MonsterItems, error) {
	items := []MonsterItems{}
	var err error

	for i := range l.json.monsters {
		mon := &l.json.monsters[i]
		monItems := mon.Items

		if monItems == nil {
			continue
		}

		monItems.MonsterID = mon.ID

		if monItems.StealCommon != nil {
			monItems.StealCommon.ID, err = l.getHashID(monItems.StealCommon)
			if err != nil {
				return nil, err
			}
		}

		if monItems.StealRare != nil {
			monItems.StealRare.ID, err = l.getHashID(monItems.StealRare)
			if err != nil {
				return nil, err
			}
		}

		if monItems.DropCommon != nil {
			monItems.DropCommon.ID, err = l.getHashID(monItems.DropCommon)
			if err != nil {
				return nil, err
			}
		}

		if monItems.DropRare != nil {
			monItems.DropRare.ID, err = l.getHashID(monItems.DropRare)
			if err != nil {
				return nil, err
			}
		}

		if monItems.SecondaryDropCommon != nil {
			monItems.SecondaryDropCommon.ID, err = l.getHashID(monItems.SecondaryDropCommon)
			if err != nil {
				return nil, err
			}
		}

		if monItems.SecondaryDropRare != nil {
			monItems.SecondaryDropRare.ID, err = l.getHashID(monItems.SecondaryDropRare)
			if err != nil {
				return nil, err
			}
		}

		if monItems.Bribe != nil {
			monItems.Bribe.ID, err = l.getHashID(monItems.Bribe)
			if err != nil {
				return nil, err
			}
		}

		items = append(items, *monItems)
	}

	return dedupeRows(items, l.Hashes), nil
}

func (l *Lookup) completeMonsterItems(items *MonsterItems) error {
	if items == nil {
		return nil
	}

	err := l.assignID(items)
	if err != nil {
		return err
	}

	err = assignIDs(l, items.OtherItems)
	if err != nil {
		return err
	}

	return nil
}
