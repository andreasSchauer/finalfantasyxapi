package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Overdrive struct {
	ID          int32
	ODCommandID *int32
	CharClassID *int32
	Ability
	Description        string             `json:"description"`
	Effect             string             `json:"effect"`
	Topmenu            *string            `json:"topmenu"`
	OverdriveCommand   *string            `json:"overdrive_command"`
	User               string             `json:"user"`
	UnlockCondition    *string            `json:"unlock_condition"`
	CountdownInSec     *int32             `json:"countdown_in_sec"`
	Cursor             *string            `json:"cursor"`
	OverdriveAbilities []AbilityReference `json:"overdrive_abilities"`
}

func (o Overdrive) ToHashFields() []any {
	return []any{
		h.DerefOrNil(o.ODCommandID),
		h.DerefOrNil(o.CharClassID),
		o.Name,
		h.DerefOrNil(o.Version),
		o.Description,
		o.Effect,
		h.DerefOrNil(o.Topmenu),
		h.ObjPtrToID(o.Attributes),
		h.DerefOrNil(o.UnlockCondition),
		h.DerefOrNil(o.CountdownInSec),
		h.DerefOrNil(o.Cursor),
	}
}

func (o Overdrive) GetID() int32 {
	return o.ID
}

func (o Overdrive) Error() string {
	return fmt.Sprintf("overdrive %s, version %v", o.Name, h.DerefOrNil(o.Version))
}

func (l *Lookup) seedOverdrives(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/overdrives.json"

	var overdrives []Overdrive
	err := loadJSONFile(string(srcPath), &overdrives)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, overdrive := range overdrives {
			overdrive.Attributes, err = seedObjPtrAssignFK(qtx, overdrive.Attributes, l.seedAbilityAttributes)
			if err != nil {
				return h.GetErr(overdrive.Error(), err)
			}

			dbOverdrive, err := qtx.CreateOverdrive(context.Background(), database.CreateOverdriveParams{
				DataHash:        generateDataHash(overdrive),
				Name:            overdrive.Name,
				Version:         h.GetNullInt32(overdrive.Version),
				Description:     overdrive.Description,
				Effect:          overdrive.Effect,
				Topmenu:         h.NullTopmenuType(overdrive.Topmenu),
				AttributesID:    h.ObjPtrToInt32ID(overdrive.Attributes),
				UnlockCondition: h.GetNullString(overdrive.UnlockCondition),
				CountdownInSec:  h.GetNullInt32(overdrive.CountdownInSec),
				Cursor:          h.NullTargetType(overdrive.Cursor),
			})
			if err != nil {
				return h.GetErr(overdrive.Error(), err, "couldn't create overdrive")
			}

			overdrive.ID = dbOverdrive.ID
			key := CreateLookupKey(overdrive.Ability)
			l.Overdrives[key] = overdrive
		}

		return nil
	})
}

func (l *Lookup) seedOverdrivesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/overdrives.json"

	var overdrives []Overdrive
	err := loadJSONFile(string(srcPath), &overdrives)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonOverdrive := range overdrives {
			overdrive, err := GetResource(jsonOverdrive.Ability, l.Overdrives)
			if err != nil {
				return err
			}

			overdrive.ODCommandID, err = assignFKPtr(overdrive.OverdriveCommand, l.OverdriveCommands)
			if err != nil {
				return h.GetErr(overdrive.Error(), err)
			}

			overdrive.CharClassID, err = assignFKPtr(&overdrive.User, l.CharClasses)
			if err != nil {
				return h.GetErr(overdrive.Error(), err)
			}

			err = qtx.UpdateOverdrive(context.Background(), database.UpdateOverdriveParams{
				DataHash:         generateDataHash(overdrive),
				OdCommandID:      h.GetNullInt32(overdrive.ODCommandID),
				CharacterClassID: h.GetNullInt32(overdrive.CharClassID),
				ID:               overdrive.ID,
			})
			if err != nil {
				return h.GetErr(overdrive.Error(), err, "couldn't update overdrive")
			}

			err = l.seedOverdriveJunctions(qtx, overdrive)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (l *Lookup) seedOverdriveJunctions(qtx *database.Queries, overdrive Overdrive) error {
	for _, abilityRef := range overdrive.OverdriveAbilities {
		junction, err := createJunction(overdrive, abilityRef, l.OverdriveAbilities)
		if err != nil {
			return h.GetErr(overdrive.Error(), err)
		}

		err = qtx.CreateOverdrivesOverdriveAbilitiesJunction(context.Background(), database.CreateOverdrivesOverdriveAbilitiesJunctionParams{
			DataHash:           generateDataHash(junction),
			OverdriveID:        junction.ParentID,
			OverdriveAbilityID: junction.ChildID,
		})
		if err != nil {
			subjects := h.JoinSubjects(overdrive.Error(), abilityRef.Error())
			return h.GetErr(subjects, err, "couldn't junction overdrive ability")
		}

		if overdrive.UnlockCondition == nil {
			err := l.seedDefaultOverdrive(qtx, overdrive, abilityRef)
			if err != nil {
				return h.GetErr(overdrive.Error(), err)
			}
		}
	}

	return nil
}

func (l *Lookup) seedDefaultOverdrive(qtx *database.Queries, overdrive Overdrive, abilityRef AbilityReference) error {
	class, err := GetResource(overdrive.User, l.CharClasses)
	if err != nil {
		return err
	}

	junction, err := createJunction(class, abilityRef, l.OverdriveAbilities)
	if err != nil {
		return h.GetErr(abilityRef.Error(), err)
	}

	err = qtx.CreateDefaultOverdriveAbility(context.Background(), database.CreateDefaultOverdriveAbilityParams{
		DataHash:  generateDataHash(junction),
		ClassID:   junction.ParentID,
		AbilityID: junction.ChildID,
	})
	if err != nil {
		return h.GetErr(abilityRef.Error(), err, "couldn't create default overdrive ability")
	}

	return nil
}
