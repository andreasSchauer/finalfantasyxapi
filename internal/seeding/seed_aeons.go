package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Aeon struct {
	ID int32
	PlayerUnit
	UnlockCondition     string       `json:"unlock_condition"`
	LocationArea        LocationArea `json:"location_area"`
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
}

func (a Aeon) ToHashFields() []any {
	return []any{
		a.PlayerUnit.ID,
		a.UnlockCondition,
		derefOrNil(a.AreaID),
		a.IsOptional,
		a.BattlesToRegenerate,
		derefOrNil(a.PhysAtkDmgConstant),
		derefOrNil(a.PhysAtkRange),
		derefOrNil(a.PhysAtkShatterRate),
		ObjPtrToID(a.PhysAtkAccuracy),
	}
}

func (a Aeon) GetID() int32 {
	return a.ID
}

func (a Aeon) Error() string {
	return fmt.Sprintf("aeon %s", a.Name)
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
		a.AutoAbilityID,
		a.CelestialWeapon,
		a.EquipType,
	}
}

func (a AeonEquipment) GetID() int32 {
	return a.ID
}

func (a AeonEquipment) Error() string {
	return fmt.Sprintf("aeon equipment with auto ability: %s, clstl_wpn: %t, equip type: %s", a.AutoAbility, a.CelestialWeapon, a.EquipType)
}

func (l *Lookup) seedAeons(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/aeons.json"

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
				return getErr(aeon.Error(), err)
			}

			dbAeon, err := qtx.CreateAeon(context.Background(), database.CreateAeonParams{
				DataHash:              generateDataHash(aeon),
				UnitID:                aeon.PlayerUnit.ID,
				UnlockCondition:       aeon.UnlockCondition,
				IsOptional:            aeon.IsOptional,
				BattlesToRegenerate:   aeon.BattlesToRegenerate,
				PhysAtkDamageConstant: getNullInt32(aeon.PhysAtkDmgConstant),
				PhysAtkRange:          getNullInt32(aeon.PhysAtkRange),
				PhysAtkShatterRate:    getNullInt32(aeon.PhysAtkShatterRate),
			})
			if err != nil {
				return getErr(aeon.Error(), err, "couldn't create aeon")
			}

			aeon.ID = dbAeon.ID
			key := createLookupKey(aeon.PlayerUnit)
			l.aeons[key] = aeon

			err = l.seedCharacterClasses(qtx, aeon.PlayerUnit)
			if err != nil {
				return getErr(aeon.Error(), err)
			}
		}
		return nil
	})
}

func (l *Lookup) seedAeonsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/aeons.json"

	var aeons []Aeon
	err := loadJSONFile(string(srcPath), &aeons)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAeon := range aeons {
			aeon, err := l.getAeon(jsonAeon.Name)
			if err != nil {
				return err
			}

			aeon.AreaID, err = assignFKPtr(&aeon.LocationArea, l.getArea)
			if err != nil {
				return getErr(aeon.Error(), err)
			}

			aeon.PhysAtkAccuracy, err = seedObjPtrAssignFK(qtx, aeon.PhysAtkAccuracy, l.seedAccuracy)
			if err != nil {
				return getErr(aeon.Error(), err)
			}

			err = qtx.UpdateAeon(context.Background(), database.UpdateAeonParams{
				DataHash:   generateDataHash(aeon),
				AreaID:     getNullInt32(aeon.AreaID),
				AccuracyID: ObjPtrToNullInt32ID(aeon.PhysAtkAccuracy),
				ID:         aeon.ID,
			})
			if err != nil {
				return getErr(aeon.Error(), err, "couldn't update aeon")
			}

			err = l.seedAeonEquipmentRelationships(qtx, aeon, string(database.EquipTypeWeapon), aeon.Weapon)
			if err != nil {
				return getErr(aeon.Error(), err)
			}

			err = l.seedAeonEquipmentRelationships(qtx, aeon, string(database.EquipTypeArmor), aeon.Armor)
			if err != nil {
				return getErr(aeon.Error(), err)
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
			return getErr(entry.Error(), err, "couldn't junction aeon equipment")
		}
	}

	return nil
}

func (l *Lookup) seedAeonEquipment(qtx *database.Queries, aeonEquipment AeonEquipment) (AeonEquipment, error) {
	var err error

	aeonEquipment.AutoAbilityID, err = assignFK(aeonEquipment.AutoAbility, l.getAutoAbility)
	if err != nil {
		return AeonEquipment{}, getErr(aeonEquipment.Error(), err)
	}

	dbAeonEquipment, err := qtx.CreateAeonEquipment(context.Background(), database.CreateAeonEquipmentParams{
		DataHash:      generateDataHash(aeonEquipment),
		AutoAbilityID: aeonEquipment.AutoAbilityID,
		CelestialWpn:  aeonEquipment.CelestialWeapon,
		EquipType:     database.EquipType(aeonEquipment.EquipType),
	})
	if err != nil {
		return AeonEquipment{}, getErr(aeonEquipment.Error(), err, "couldn't create aeon equipment")
	}
	aeonEquipment.ID = dbAeonEquipment.ID

	return aeonEquipment, nil
}
