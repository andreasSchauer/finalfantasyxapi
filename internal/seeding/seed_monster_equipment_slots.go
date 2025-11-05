package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type MonsterEquipmentSlots struct {
	ID					int32
	MonsterEquipmentID	int32
	MinAmount			int32						`json:"min_amount"`
	MaxAmount			int32						`json:"max_amount"`
	Chances				[]EquipmentSlotsChance		`json:"chances"`
	Type				database.EquipmentSlotsType
}

func (m MonsterEquipmentSlots) ToHashFields() []any {
	return []any{
		m.MonsterEquipmentID,
		m.MinAmount,
		m.MaxAmount,
		m.Type,
	}
}

func (m MonsterEquipmentSlots) GetID() int32 {
	return m.ID
}

func (m MonsterEquipmentSlots) Error() string {
	return fmt.Sprintf("monster equipment slots with monster equipment id: %d, type: %s, min amount: %d, max amount: %d", m.ID, m.Type, m.MinAmount, m.MaxAmount)
}



type EquipmentSlotsChance struct {
	ID			int32
	Amount		int32	`json:"amount"`
	Chance		int32	`json:"chance"`
}

func (e EquipmentSlotsChance) ToHashFields() []any {
	return []any{
		e.Amount,
		e.Chance,
	}
}

func (e EquipmentSlotsChance) GetID() int32 {
	return e.ID
}

func (e EquipmentSlotsChance) Error() string {
	return fmt.Sprintf("equipment slots chance with amount: %d, chance: %d", e.Amount, e.Chance)
}



func (l *lookup) seedMonsterEquipmentSlotsWrapper(qtx *database.Queries, monsterEquipment MonsterEquipment, mes MonsterEquipmentSlots, slotsType database.EquipmentSlotsType) error {
	var err error

	mes.MonsterEquipmentID = monsterEquipment.ID
	mes.Type = slotsType
	
	mes, err = seedObjAssignID(qtx, mes, l.seedMonsterEquipmentSlots)
	if err != nil {
		return getErr(monsterEquipment.Error(), err)
	}

	return nil
}


func (l *lookup) seedMonsterEquipmentSlots(qtx *database.Queries, mes MonsterEquipmentSlots) (MonsterEquipmentSlots, error) {
	dbMes, err := qtx.CreateMonsterEquipmentSlots(context.Background(), database.CreateMonsterEquipmentSlotsParams{
		DataHash: 			generateDataHash(mes),
		MonsterEquipmentID: mes.MonsterEquipmentID,
		MinAmount: 			mes.MinAmount,
		MaxAmount: 			mes.MaxAmount,
		Type: 				mes.Type,
	})
	if err != nil {
		return MonsterEquipmentSlots{}, getErr(mes.Error(), err, "couldn't create monster equipment slots")
	}
	mes.ID = dbMes.ID

	err = l.seedMonsterEquipmentSlotsChances(qtx, mes)
	if err != nil {
		return MonsterEquipmentSlots{}, getErr(mes.Error(), err)
	}

	return mes, nil
}


func (l *lookup) seedMonsterEquipmentSlotsChances(qtx *database.Queries, mes MonsterEquipmentSlots) error {
	monsterEquipment := l.currentME

	for _, chance := range mes.Chances {
		threeWay, err := createThreeWayJunctionSeed(qtx, monsterEquipment, mes, chance, l.seedEquipmentSlotsChance)
		if err != nil {
			return err
		}

		err = qtx.CreateMonsterEquipmentSlotsChancesJunction(context.Background(), database.CreateMonsterEquipmentSlotsChancesJunctionParams{
			DataHash: 			generateDataHash(threeWay),
			MonsterEquipmentID: threeWay.GrandparentID,
			EquipmentSlotsID: 	threeWay.ParentID,
			SlotsChanceID: 		threeWay.ChildID,
		})
		if err != nil {
			return getErr(chance.Error(), err, "couldn't junction chance")
		}
	}

	return nil
}


func (l *lookup) seedEquipmentSlotsChance(qtx *database.Queries, esc EquipmentSlotsChance) (EquipmentSlotsChance, error) {
	dbEsc, err := qtx.CreateEquipmentSlotsChance(context.Background(), database.CreateEquipmentSlotsChanceParams{
		DataHash: 	generateDataHash(esc),
		Amount: 	esc.Amount,
		Chance: 	esc.Chance,
	})
	if err != nil {
		return EquipmentSlotsChance{}, getErr(esc.Error(), err, "couldn't create equipment slots chance")
	}

	esc.ID = dbEsc.ID

	return esc, nil
}