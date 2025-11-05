package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
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
		m.MonsterID,
		m.DropChance,
		derefOrNil(m.DropCondition),
		derefOrNil(m.OtherItemsCondition),
		ObjPtrToID(m.StealCommon),
		ObjPtrToID(m.StealRare),
		ObjPtrToID(m.DropCommon),
		ObjPtrToID(m.DropRare),
		ObjPtrToID(m.SecondaryDropCommon),
		ObjPtrToID(m.SecondaryDropRare),
		ObjPtrToID(m.Bribe),
	}
}

func (m MonsterItems) GetID() int32 {
	return m.ID
}

func (m MonsterItems) Error() string {
	return fmt.Sprintf("monster items of monster with id %d", m.MonsterID)
}

func (l *lookup) seedMonsterItems(qtx *database.Queries, monsterItems MonsterItems) (MonsterItems, error) {
	var err error

	monsterItems.StealCommon, err = seedObjPtrAssignFK(qtx, monsterItems.StealCommon, l.seedItemAmount)
	if err != nil {
		return MonsterItems{}, getErr(monsterItems.Error(), err, "steal common")
	}

	monsterItems.StealRare, err = seedObjPtrAssignFK(qtx, monsterItems.StealRare, l.seedItemAmount)
	if err != nil {
		return MonsterItems{}, getErr(monsterItems.Error(), err, "steal rare")
	}

	monsterItems.DropCommon, err = seedObjPtrAssignFK(qtx, monsterItems.DropCommon, l.seedItemAmount)
	if err != nil {
		return MonsterItems{}, getErr(monsterItems.Error(), err, "drop common")
	}

	monsterItems.DropRare, err = seedObjPtrAssignFK(qtx, monsterItems.DropRare, l.seedItemAmount)
	if err != nil {
		return MonsterItems{}, getErr(monsterItems.Error(), err, "drop rare")
	}

	monsterItems.SecondaryDropCommon, err = seedObjPtrAssignFK(qtx, monsterItems.SecondaryDropCommon, l.seedItemAmount)
	if err != nil {
		return MonsterItems{}, getErr(monsterItems.Error(), err, "secondary drop common")
	}

	monsterItems.SecondaryDropRare, err = seedObjPtrAssignFK(qtx, monsterItems.SecondaryDropRare, l.seedItemAmount)
	if err != nil {
		return MonsterItems{}, getErr(monsterItems.Error(), err, "secondary drop rare")
	}

	monsterItems.Bribe, err = seedObjPtrAssignFK(qtx, monsterItems.Bribe, l.seedItemAmount)
	if err != nil {
		return MonsterItems{}, getErr(monsterItems.Error(), err, "bribe")
	}

	dbMonsterItems, err := qtx.CreateMonsterItem(context.Background(), database.CreateMonsterItemParams{
		DataHash:              generateDataHash(monsterItems),
		MonsterID:             monsterItems.MonsterID,
		DropChance:            monsterItems.DropChance,
		DropCondition:         getNullString(monsterItems.DropCondition),
		OtherItemsCondition:   getNullString(monsterItems.OtherItemsCondition),
		StealCommonID:         ObjPtrToNullInt32ID(monsterItems.StealCommon),
		StealRareID:           ObjPtrToNullInt32ID(monsterItems.StealRare),
		DropCommonID:          ObjPtrToNullInt32ID(monsterItems.DropCommon),
		DropRareID:            ObjPtrToNullInt32ID(monsterItems.DropRare),
		SecondaryDropCommonID: ObjPtrToNullInt32ID(monsterItems.SecondaryDropCommon),
		SecondaryDropRareID:   ObjPtrToNullInt32ID(monsterItems.SecondaryDropRare),
		BribeID:               ObjPtrToNullInt32ID(monsterItems.Bribe),
	})
	if err != nil {
		return MonsterItems{}, getErr(monsterItems.Error(), err, "couldn't create monster items")
	}

	monsterItems.ID = dbMonsterItems.ID

	err = l.seedMonsterOtherItems(qtx, monsterItems)
	if err != nil {
		return MonsterItems{}, getErr(monsterItems.Error(), err)
	}

	return monsterItems, nil
}

func (l *lookup) seedMonsterOtherItems(qtx *database.Queries, monsterItems MonsterItems) error {
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
			return getErr(posItem.Error(), err, "couldn't junction other item")
		}
	}

	return nil
}
