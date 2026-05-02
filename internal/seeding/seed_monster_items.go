package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type MonsterItems struct {
	ID                  int32
	MonsterID           int32
	DropChance          int32          `json:"drop_chance"`
	DropCondition       *string        `json:"drop_condition"`
	OtherItemsCondition *string        `json:"other_items_condition"`
	OtherItems          []PossibleItem `json:"other_items"`
	StealCommon         *ItemAmount    `json:"steal_common"`
	StealRare           *ItemAmount    `json:"steal_rare"`
	DropCommon          *ItemAmount    `json:"drop_common"`
	DropRare            *ItemAmount    `json:"drop_rare"`
	SecondaryDropCommon *ItemAmount    `json:"secondary_drop_common"`
	SecondaryDropRare   *ItemAmount    `json:"secondary_drop_rare"`
	Bribe               *ItemAmount    `json:"bribe"`
}

func (m MonsterItems) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", m),
		m.MonsterID,
		m.DropChance,
		h.DerefOrNil(m.DropCondition),
		h.DerefOrNil(m.OtherItemsCondition),
		h.ObjPtrToID(m.StealCommon),
		h.ObjPtrToID(m.StealRare),
		h.ObjPtrToID(m.DropCommon),
		h.ObjPtrToID(m.DropRare),
		h.ObjPtrToID(m.SecondaryDropCommon),
		h.ObjPtrToID(m.SecondaryDropRare),
		h.ObjPtrToID(m.Bribe),
	}
}

func (m MonsterItems) GetID() int32 {
	return m.ID
}

func (m *MonsterItems) SetID(id int32) {
	m.ID = id
}

func (m MonsterItems) Error() string {
	return fmt.Sprintf("monster items of monster with id %d", m.MonsterID)
}

func (l *Lookup) seedMonsterItems(qtx *database.Queries, monsterItems MonsterItems) (MonsterItems, error) {
	var err error

	monsterItems.StealCommon, err = seedObjPtrAssignFK(qtx, monsterItems.StealCommon, l.seedItemAmount)
	if err != nil {
		return MonsterItems{}, h.NewErr(monsterItems.Error(), err, "steal common")
	}

	monsterItems.StealRare, err = seedObjPtrAssignFK(qtx, monsterItems.StealRare, l.seedItemAmount)
	if err != nil {
		return MonsterItems{}, h.NewErr(monsterItems.Error(), err, "steal rare")
	}

	monsterItems.DropCommon, err = seedObjPtrAssignFK(qtx, monsterItems.DropCommon, l.seedItemAmount)
	if err != nil {
		return MonsterItems{}, h.NewErr(monsterItems.Error(), err, "drop common")
	}

	monsterItems.DropRare, err = seedObjPtrAssignFK(qtx, monsterItems.DropRare, l.seedItemAmount)
	if err != nil {
		return MonsterItems{}, h.NewErr(monsterItems.Error(), err, "drop rare")
	}

	monsterItems.SecondaryDropCommon, err = seedObjPtrAssignFK(qtx, monsterItems.SecondaryDropCommon, l.seedItemAmount)
	if err != nil {
		return MonsterItems{}, h.NewErr(monsterItems.Error(), err, "secondary drop common")
	}

	monsterItems.SecondaryDropRare, err = seedObjPtrAssignFK(qtx, monsterItems.SecondaryDropRare, l.seedItemAmount)
	if err != nil {
		return MonsterItems{}, h.NewErr(monsterItems.Error(), err, "secondary drop rare")
	}

	monsterItems.Bribe, err = seedObjPtrAssignFK(qtx, monsterItems.Bribe, l.seedItemAmount)
	if err != nil {
		return MonsterItems{}, h.NewErr(monsterItems.Error(), err, "bribe")
	}

	dbMonsterItems, err := qtx.CreateMonsterItem(context.Background(), database.CreateMonsterItemParams{
		DataHash:              generateDataHash(monsterItems),
		MonsterID:             monsterItems.MonsterID,
		DropChance:            monsterItems.DropChance,
		DropCondition:         h.GetNullString(monsterItems.DropCondition),
		OtherItemsCondition:   h.GetNullString(monsterItems.OtherItemsCondition),
		StealCommonID:         h.ObjPtrToNullInt32ID(monsterItems.StealCommon),
		StealRareID:           h.ObjPtrToNullInt32ID(monsterItems.StealRare),
		DropCommonID:          h.ObjPtrToNullInt32ID(monsterItems.DropCommon),
		DropRareID:            h.ObjPtrToNullInt32ID(monsterItems.DropRare),
		SecondaryDropCommonID: h.ObjPtrToNullInt32ID(monsterItems.SecondaryDropCommon),
		SecondaryDropRareID:   h.ObjPtrToNullInt32ID(monsterItems.SecondaryDropRare),
		BribeID:               h.ObjPtrToNullInt32ID(monsterItems.Bribe),
	})
	if err != nil {
		return MonsterItems{}, h.NewErr(monsterItems.Error(), err, "couldn't create monster items")
	}

	monsterItems.ID = dbMonsterItems.ID

	err = l.seedMonsterOtherItems(qtx, monsterItems)
	if err != nil {
		return MonsterItems{}, h.NewErr(monsterItems.Error(), err)
	}

	return monsterItems, nil
}

func (l *Lookup) seedMonsterOtherItems(qtx *database.Queries, monsterItems MonsterItems) error {
	for _, posItem := range monsterItems.OtherItems {
		junction, err := createJunctionSeed(qtx, monsterItems, posItem, l.seedPossibleItem)
		if err != nil {
			return err
		}

		err = qtx.CreateMonsterItemsOtherItemsJunction(context.Background(), database.CreateMonsterItemsOtherItemsJunctionParams{
			DataHash:       generateDataHash(junction),
			MonsterItemsID: junction.ParentID,
			PossibleItemID: junction.ChildID,
		})
		if err != nil {
			return h.NewErr(posItem.Error(), err, "couldn't junction other item")
		}
	}

	return nil
}


func (l *Lookup) loop3SeedMonsterItems(qtx *database.Queries, ctx context.Context) error {
	items, err := l.extractMonsterItems()
	if err != nil {
		return err
	}

	params := database.CreateMonsterItemBulkParams{
		DataHash: 				make([]string, len(items)),
		MonsterID: 				make([]int32, len(items)),
		DropChance: 			make([]int32, len(items)),
		DropCondition: 			make([]sql.NullString, len(items)),
		OtherItemsCondition: 	make([]sql.NullString, len(items)),
		StealCommonID: 			make([]sql.NullInt32, len(items)),
		StealRareID: 			make([]sql.NullInt32, len(items)),
		DropCommonID: 			make([]sql.NullInt32, len(items)),
		DropRareID: 			make([]sql.NullInt32, len(items)),
		SecondaryDropCommonID: 	make([]sql.NullInt32, len(items)),
		SecondaryDropRareID: 	make([]sql.NullInt32, len(items)),
		BribeID: 				make([]sql.NullInt32, len(items)),
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