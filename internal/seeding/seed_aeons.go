package seeding

import (
	"context"
	"database/sql"
	"fmt"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Aeon struct {
	ID int32
	PlayerUnit
	UnlockCondition     string       	`json:"unlock_condition"`
	LocationArea        LocationArea 	`json:"location_area"`
	AreaID              *int32
	Category            *string         `json:"category"`
	IsOptional          bool            `json:"is_optional"`
	BattlesToRegenerate int32           `json:"num_battles_to_regenerate"`
	Weapon              []AeonEquipment `json:"weapon"`
	Armor               []AeonEquipment `json:"armor"`
	PhysAtkDmgConstant  *int32          `json:"phys_atk_damage_constant"`
	PhysAtkRange        *int32          `json:"phys_atk_range"`
	PhysAtkShatterRate  *int32          `json:"phys_atk_shatter_rate"`
	PhysAtkAccuracy     *Accuracy       `json:"phys_atk_accuracy"`
	BaseStats			AeonStat
}

func (a Aeon) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.PlayerUnit.ID,
		a.UnlockCondition,
		h.DerefOrNil(a.AreaID),
		a.IsOptional,
		a.BattlesToRegenerate,
		h.DerefOrNil(a.PhysAtkDmgConstant),
		h.DerefOrNil(a.PhysAtkRange),
		h.DerefOrNil(a.PhysAtkShatterRate),
		h.ObjPtrToID(a.PhysAtkAccuracy),
	}
}

func (a Aeon) GetID() int32 {
	return a.ID
}

func (a Aeon) Error() string {
	return fmt.Sprintf("aeon %s", a.Name)
}

func (a Aeon) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   a.ID,
		Name: a.Name,
	}
}

type AeonEquipment struct {
	ID              int32
	AutoAbilityID   int32
	AutoAbility     string `json:"ability"`
	CelestialWeapon bool   `json:"celestial_wpn"`
	EquipType       string
}

func (a AeonEquipment) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.AutoAbilityID,
		a.CelestialWeapon,
		a.EquipType,
	}
}

func (a AeonEquipment) GetID() int32 {
	return a.ID
}

func (a *AeonEquipment) SetID(id int32) {
	a.ID = id
}

func (a AeonEquipment) Error() string {
	return fmt.Sprintf("aeon equipment with auto ability: %s, clstl_wpn: %t, equip type: %s", a.AutoAbility, a.CelestialWeapon, a.EquipType)
}


func (l *Lookup) loop4SeedAeons(qtx *database.Queries, ctx context.Context) error {
	aeons, err := l.extractAeons()
	if err != nil {
		return err
	}

	params := database.CreateAeonBulkParams{
		DataHash:   			make([]string, len(aeons)),
		UnitID: 				make([]int32, len(aeons)),
		UnlockCondition: 		make([]string, len(aeons)),
		IsOptional: 			make([]bool, len(aeons)),
		BattlesToRegenerate: 	make([]int32, len(aeons)),
		PhysAtkDamageConstant: 	make([]sql.NullInt32, len(aeons)),
		PhysAtkRange: 			make([]sql.NullInt32, len(aeons)),
		PhysAtkShatterRate: 	make([]sql.NullInt32, len(aeons)),
		AreaID: 				make([]sql.NullInt32, len(aeons)),
		AccuracyID: 			make([]sql.NullInt32, len(aeons)),
	}

	for i, a := range aeons {
		params.DataHash[i] = generateDataHash(a)
		params.UnitID[i] = a.PlayerUnit.ID
		params.UnlockCondition[i] = a.UnlockCondition
		params.IsOptional[i] = a.IsOptional
		params.BattlesToRegenerate[i] = a.BattlesToRegenerate
		params.PhysAtkDamageConstant[i] = h.GetNullInt32(a.PhysAtkDmgConstant)
		params.PhysAtkRange[i] = h.GetNullInt32(a.PhysAtkRange)
		params.PhysAtkShatterRate[i] = h.GetNullInt32(a.PhysAtkShatterRate)
		params.AreaID[i] = h.GetNullInt32(a.AreaID)
		params.AccuracyID[i] = h.ObjPtrToNullInt32ID(a.PhysAtkAccuracy)
	}

	dbRows, err := qtx.CreateAeonBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create aeons: %v", err)
	}

	for i, row := range dbRows {
		aeons[i].ID = row.ID
		l.json.aeons[i].ID = row.ID
		l.Aeons[aeons[i].Name] = aeons[i]
		l.AeonsID[row.ID] = aeons[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAeons() ([]Aeon, error) {
	aeons := []Aeon{}
	var err error

	for i := range l.json.aeons {
		aeon := &l.json.aeons[i]

		aeon.PlayerUnit.ID, err = l.getHashID(aeon.PlayerUnit)
		if err != nil {
			return nil, err
		}

		aeon.AreaID, err = assignFKPtr(&aeon.LocationArea, l.Areas)
		if err != nil {
			return nil, err
		}

		if aeon.PhysAtkAccuracy != nil {
			aeon.PhysAtkAccuracy.ID, err = l.getHashID(aeon.PhysAtkAccuracy)
			if err != nil {
				return nil, err
			}
		}

		aeons = append(aeons, *aeon)
	}

	return dedupeRows(aeons, l.Hashes), nil
}

func (l *Lookup) loop6SeedAeonEquipment(qtx *database.Queries, ctx context.Context) error {
	equipment, err := l.extractAeonEquipment()
	if err != nil {
		return err
	}

	params := database.CreateAeonEquipmentBulkParams{
		DataHash:   	make([]string, len(equipment)),
		AutoAbilityID: 	make([]int32, len(equipment)),
		CelestialWpn: 	make([]bool, len(equipment)),
		EquipType: 		make([]database.EquipType, len(equipment)),
	}

	for i, e := range equipment {
		params.DataHash[i] = generateDataHash(e)
		params.AutoAbilityID[i] = e.AutoAbilityID
		params.CelestialWpn[i] = e.CelestialWeapon
		params.EquipType[i] = database.EquipType(e.EquipType)
	}

	dbRows, err := qtx.CreateAeonEquipmentBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create aeon equipment: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAeonEquipment() ([]AeonEquipment, error) {
	equipment := []AeonEquipment{}
	var err error

	for i := range l.json.aeons {
		aeon := &l.json.aeons[i]

		for j := range aeon.Weapon {
			ae := &aeon.Weapon[j]
			ae.EquipType = string(database.EquipTypeWeapon)

			ae.AutoAbilityID, err = assignFK(ae.AutoAbility, l.AutoAbilities)
			if err != nil {
				return nil, err
			}

			equipment = append(equipment, *ae)
		}
		
		for j := range aeon.Armor {
			ae := &aeon.Armor[j]
			ae.EquipType = string(database.EquipTypeArmor)

			ae.AutoAbilityID, err = assignFK(ae.AutoAbility, l.AutoAbilities)
			if err != nil {
				return nil, err
			}

			equipment = append(equipment, *ae)
		}
	}

	return dedupeRows(equipment, l.Hashes), nil
}

func (l *Lookup) completeAeons() error {
	for i := range l.json.aeons {
		aeon := &l.json.aeons[i]
		err := assignIDs(l, aeon.Weapon)
		if err != nil {
			return err
		}

		err = assignIDs(l, aeon.Armor)
		if err != nil {
			return err
		}

		l.Aeons[aeon.Name] = *aeon
		l.AeonsID[aeon.ID] = *aeon
	}

	return nil
}

func (l *Lookup) getAeonAeonEquipment(a Aeon) ([]AeonEquipment, error) {
	return slices.Concat(a.Weapon, a.Armor), nil
}

func (l *Lookup) seedJuncAeonAeonEquipment(qtx *database.Queries, ctx context.Context) error {
	const desc string = "aeons + aeon equipment"
	jParams, err := processJunctions(l, desc, l.json.aeons, l.getAeonAeonEquipment)
	if err != nil {
		return err
	}

	return qtx.CreateAeonsWeaponArmorJunctionBulk(ctx, database.CreateAeonsWeaponArmorJunctionBulkParams{
		DataHash:       	jParams.DataHashes,
		AeonID: 			jParams.ParentIDs,
		AeonEquipmentID:  	jParams.ChildIDs,
	})
}