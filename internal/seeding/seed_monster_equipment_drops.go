package seeding

import (
	"context"
	"database/sql"
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
		fmt.Sprintf("%T", e),
		e.AutoAbilityID,
		e.IsForced,
		h.DerefOrNil(e.Probability),
		e.Type,
	}
}

func (e EquipmentDrop) GetID() int32 {
	return e.ID
}

func (e *EquipmentDrop) SetID(id int32) {
	e.ID = id
}

func (e EquipmentDrop) Error() string {
	return fmt.Sprintf("equipment drop with auto-ability id: %d, type: %s, is forced: %t, probability: %v", e.AutoAbilityID, e.Type, e.IsForced, h.PtrToString(e.Probability))
}

func (l *Lookup) loop6SeedEquipmentDrops(qtx *database.Queries, ctx context.Context) error {
	drops, err := l.extractEquipmentDrops()
	if err != nil {
		return err
	}

	params := database.CreateEquipmentDropBulkParams{
		DataHash:      make([]string, len(drops)),
		AutoAbilityID: make([]int32, len(drops)),
		IsForced:      make([]bool, len(drops)),
		Probability:   make([]sql.NullInt32, len(drops)),
		Type:          make([]database.EquipType, len(drops)),
	}

	for i, d := range drops {
		params.DataHash[i] = generateDataHash(d)
		params.AutoAbilityID[i] = d.AutoAbilityID
		params.IsForced[i] = d.IsForced
		params.Probability[i] = h.GetNullInt32(d.Probability)
		params.Type[i] = d.Type
	}

	dbRows, err := qtx.CreateEquipmentDropBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create equipment drops: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractEquipmentDrops() ([]EquipmentDrop, error) {
	drops := []EquipmentDrop{}
	var err error

	for i := range l.json.monsters {
		mon := &l.json.monsters[i]

		if mon.Equipment == nil {
			continue
		}

		for j := range mon.Equipment.WeaponAbilities {
			drop := &mon.Equipment.WeaponAbilities[j]
			drop.Type = database.EquipTypeWeapon

			drop.AutoAbilityID, err = assignFK(drop.Ability, l.AutoAbilities)
			if err != nil {
				return nil, err
			}

			drops = append(drops, *drop)
		}

		for j := range mon.Equipment.ArmorAbilities {
			drop := &mon.Equipment.ArmorAbilities[j]
			drop.Type = database.EquipTypeArmor

			drop.AutoAbilityID, err = assignFK(drop.Ability, l.AutoAbilities)
			if err != nil {
				return nil, err
			}

			drops = append(drops, *drop)
		}
	}

	return dedupeRows(drops, l.Hashes), nil
}

func (l *Lookup) getEquipmentDropCharacters(ed EquipmentDrop) ([]Character, error) {
	return getResources(ed.Characters, l.Characters)
}

func (l *Lookup) seedJuncEquipmentDropsCharacters(qtx *database.Queries, ctx context.Context) error {
	const desc string = "equipment drops + characters"
	jParams, err := processThreewayJunctions(l, desc, l.getMonsterEquipments(), l.getMonsterEquipmentEquipmentDrops, l.getEquipmentDropCharacters)
	if err != nil {
		return err
	}

	return qtx.CreateEquipmentDropsCharactersJunctionBulk(ctx, database.CreateEquipmentDropsCharactersJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		MonsterEquipmentID: jParams.GrandParentIDs,
		EquipmentDropID:    jParams.ParentIDs,
		CharacterID:        jParams.ChildIDs,
	})
}
