package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type EquipmentDrop struct {
	ID            int32
	AutoAbilityID int32
	Ability       string   `json:"ability"`
	Characters    []string `json:"characters"`
	IsForced      bool     `json:"is_forced"`
	Probability   *int32   `json:"probability"`
	Type          database.EquipType
}

func (e EquipmentDrop) ToHashFields() []any {
	return []any{
		e.AutoAbilityID,
		e.IsForced,
		h.DerefOrNil(e.Probability),
		e.Type,
	}
}

func (e EquipmentDrop) GetID() int32 {
	return e.ID
}

func (e EquipmentDrop) Error() string {
	return fmt.Sprintf("equipment drop with auto-ability id: %d, type: %s, is forced: %t, probability: %v", e.AutoAbilityID, e.Type, e.IsForced, h.DerefOrNil(e.Probability))
}

func (l *Lookup) seedEquipmentDrops(qtx *database.Queries, monsterEquipment MonsterEquipment, drops []EquipmentDrop, equipType database.EquipType) error {
	for _, drop := range drops {
		var err error
		drop.Type = equipType

		junction, err := createJunctionSeed(qtx, monsterEquipment, drop, l.seedEquipmentDrop)
		if err != nil {
			return err
		}

		err = qtx.CreateMonsterEquipmentAbilitiesJunction(context.Background(), database.CreateMonsterEquipmentAbilitiesJunctionParams{
			DataHash:           generateDataHash(junction),
			MonsterEquipmentID: junction.ParentID,
			EquipmentDropID:    junction.ChildID,
		})
		if err != nil {
			return h.GetErr(drop.Error(), err, "couldn't junction equipment drop")
		}
	}

	return nil
}

func (l *Lookup) seedEquipmentDrop(qtx *database.Queries, drop EquipmentDrop) (EquipmentDrop, error) {
	var err error

	drop.AutoAbilityID, err = assignFK(drop.Ability, l.autoAbilities)
	if err != nil {
		return EquipmentDrop{}, h.GetErr(drop.Error(), err)
	}

	dbEquipmentDrop, err := qtx.CreateEquipmentDrop(context.Background(), database.CreateEquipmentDropParams{
		DataHash:      generateDataHash(drop),
		AutoAbilityID: drop.AutoAbilityID,
		IsForced:      drop.IsForced,
		Probability:   h.GetNullInt32(drop.Probability),
		Type:          drop.Type,
	})
	if err != nil {
		return EquipmentDrop{}, h.GetErr(drop.Error(), err, "couldn't create equipment drop")
	}

	drop.ID = dbEquipmentDrop.ID

	err = l.seedEquipmentDropCharacters(qtx, drop)
	if err != nil {
		return EquipmentDrop{}, h.GetErr(drop.Error(), err)
	}

	return drop, nil
}

func (l *Lookup) seedEquipmentDropCharacters(qtx *database.Queries, drop EquipmentDrop) error {
	monsterEquipment := l.currentME

	for _, character := range drop.Characters {
		threeWay, err := createThreeWayJunction(monsterEquipment, drop, character, l.characters)
		if err != nil {
			return err
		}

		err = qtx.CreateEquipmentDropsCharactersJunction(context.Background(), database.CreateEquipmentDropsCharactersJunctionParams{
			DataHash:           generateDataHash(threeWay),
			MonsterEquipmentID: threeWay.GrandparentID,
			EquipmentDropID:    threeWay.ParentID,
			CharacterID:        threeWay.ChildID,
		})
		if err != nil {
			return h.GetErr(character, err, "couldn't junction character")
		}
	}

	return nil
}
