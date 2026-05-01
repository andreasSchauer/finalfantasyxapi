package seeding

import (
	"context"
	"database/sql"
	"fmt"

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

func (l *Lookup) seedAeons(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/aeons.json"

	var aeons []Aeon
	err := loadJSONFile(string(srcPath), &aeons)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, aeon := range aeons {
			var err error
			aeon.Type = database.UnitTypeAeon

			aeon.PlayerUnit, err = seedObjAssignID(qtx, aeon.PlayerUnit, l.seedPlayerUnit)
			if err != nil {
				return h.NewErr(aeon.Error(), err)
			}

			dbAeon, err := qtx.CreateAeon(context.Background(), database.CreateAeonParams{
				DataHash:              generateDataHash(aeon),
				UnitID:                aeon.PlayerUnit.ID,
				UnlockCondition:       aeon.UnlockCondition,
				IsOptional:            aeon.IsOptional,
				BattlesToRegenerate:   aeon.BattlesToRegenerate,
				PhysAtkDamageConstant: h.GetNullInt32(aeon.PhysAtkDmgConstant),
				PhysAtkRange:          h.GetNullInt32(aeon.PhysAtkRange),
				PhysAtkShatterRate:    h.GetNullInt32(aeon.PhysAtkShatterRate),
			})
			if err != nil {
				return h.NewErr(aeon.Error(), err, "couldn't create aeon")
			}

			aeon.ID = dbAeon.ID
			l.Aeons[aeon.Name] = aeon
			l.AeonsID[aeon.ID] = aeon
		}
		
		return nil
	})
}

func (l *Lookup) seedAeonsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/aeons.json"

	var aeons []Aeon
	err := loadJSONFile(string(srcPath), &aeons)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAeon := range aeons {
			aeon, err := GetResource(jsonAeon.Name, l.Aeons)
			if err != nil {
				return err
			}

			aeon.AreaID, err = assignFKPtr(&aeon.LocationArea, l.Areas)
			if err != nil {
				return h.NewErr(aeon.Error(), err)
			}

			aeon.PhysAtkAccuracy, err = seedObjPtrAssignFK(qtx, aeon.PhysAtkAccuracy, l.seedAccuracy)
			if err != nil {
				return h.NewErr(aeon.Error(), err)
			}

			err = qtx.UpdateAeon(context.Background(), database.UpdateAeonParams{
				DataHash:   generateDataHash(aeon),
				AreaID:     h.GetNullInt32(aeon.AreaID),
				AccuracyID: h.ObjPtrToNullInt32ID(aeon.PhysAtkAccuracy),
				ID:         aeon.ID,
			})
			if err != nil {
				return h.NewErr(aeon.Error(), err, "couldn't update aeon")
			}

			err = l.seedAeonEquipmentRelationships(qtx, aeon, string(database.EquipTypeWeapon), aeon.Weapon)
			if err != nil {
				return h.NewErr(aeon.Error(), err)
			}

			err = l.seedAeonEquipmentRelationships(qtx, aeon, string(database.EquipTypeArmor), aeon.Armor)
			if err != nil {
				return h.NewErr(aeon.Error(), err)
			}
		}
		return nil
	})
}

func (l *Lookup) seedAeonEquipmentRelationships(qtx *database.Queries, aeon Aeon, equipType string, abilityList []AeonEquipment) error {
	for _, entry := range abilityList {
		var err error
		entry.EquipType = equipType

		junction, err := createJunctionSeed(qtx, aeon, entry, l.seedAeonEquipment)
		if err != nil {
			return err
		}

		err = qtx.CreateAeonsWeaponArmorJunction(context.Background(), database.CreateAeonsWeaponArmorJunctionParams{
			DataHash:        generateDataHash(junction),
			AeonID:          junction.ParentID,
			AeonEquipmentID: junction.ChildID,
		})
		if err != nil {
			return h.NewErr(entry.Error(), err, "couldn't junction aeon equipment")
		}
	}

	return nil
}

func (l *Lookup) seedAeonEquipment(qtx *database.Queries, aeonEquipment AeonEquipment) (AeonEquipment, error) {
	var err error

	aeonEquipment.AutoAbilityID, err = assignFK(aeonEquipment.AutoAbility, l.AutoAbilities)
	if err != nil {
		return AeonEquipment{}, h.NewErr(aeonEquipment.Error(), err)
	}

	dbAeonEquipment, err := qtx.CreateAeonEquipment(context.Background(), database.CreateAeonEquipmentParams{
		DataHash:      generateDataHash(aeonEquipment),
		AutoAbilityID: aeonEquipment.AutoAbilityID,
		CelestialWpn:  aeonEquipment.CelestialWeapon,
		EquipType:     database.EquipType(aeonEquipment.EquipType),
	})
	if err != nil {
		return AeonEquipment{}, h.NewErr(aeonEquipment.Error(), err, "couldn't create aeon equipment")
	}
	aeonEquipment.ID = dbAeonEquipment.ID

	return aeonEquipment, nil
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