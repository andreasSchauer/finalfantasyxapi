package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type Overdrive struct {
	ID          		int32
	ODCommandID 		*int32
	CharClassID			*int32
	Ability
	Description     	string  			`json:"description"`
	Effect          	string  			`json:"effect"`
	Topmenu         	*string 			`json:"topmenu"`
	OverdriveCommand	*string				`json:"overdrive_command"`
	User				string				`json:"user"`
	UnlockCondition 	*string 			`json:"unlock_condition"`
	CountdownInSec  	*int32  			`json:"countdown_in_sec"`
	Cursor          	*string 			`json:"cursor"`
	OverdriveAbilities	[]AbilityReference	`json:"overdrive_abilities"`
}

func (o Overdrive) ToHashFields() []any {
	return []any{
		derefOrNil(o.ODCommandID),
		derefOrNil(o.CharClassID),
		o.Name,
		derefOrNil(o.Version),
		o.Description,
		o.Effect,
		derefOrNil(o.Topmenu),
		ObjPtrToHashID(o.Attributes),
		derefOrNil(o.UnlockCondition),
		derefOrNil(o.CountdownInSec),
		derefOrNil(o.Cursor),
	}
}


func (o Overdrive) GetID() int32 {
	return o.ID
}

func (l *lookup) seedOverdrives(db *database.Queries, dbConn *sql.DB) error {
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
				return fmt.Errorf("couldn't create Ability Attributes: %s-%d, type: %s: %v", overdrive.Name, derefOrNil(overdrive.Version), overdrive.Type, err)
			}

			dbOverdrive, err := qtx.CreateOverdrive(context.Background(), database.CreateOverdriveParams{
				DataHash:        generateDataHash(overdrive),
				Name:            overdrive.Name,
				Version:         getNullInt32(overdrive.Version),
				Description:     overdrive.Description,
				Effect:          overdrive.Effect,
				Topmenu:         nullTopmenuType(overdrive.Topmenu),
				AttributesID:    ObjPtrToInt32ID(overdrive.Attributes),
				UnlockCondition: getNullString(overdrive.UnlockCondition),
				CountdownInSec:  getNullInt32(overdrive.CountdownInSec),
				Cursor:          nullTargetType(overdrive.Cursor),
			})
			if err != nil {
				return fmt.Errorf("couldn't create Overdrive: %s: %v", overdrive.Name, err)
			}

			overdrive.ID = dbOverdrive.ID
			key := createLookupKey(overdrive.Ability)
			l.overdrives[key] = overdrive
		}

		return nil
	})
}


func (l *lookup) createOverdrivesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/overdrives.json"

	var overdrives []Overdrive
	err := loadJSONFile(string(srcPath), &overdrives)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonOverdrive := range overdrives {
			overdrive, err := l.getOverdrive(jsonOverdrive.Ability)
			if err != nil {
				return err
			}

			overdrive.ODCommandID, err = assignFKPtr(overdrive.OverdriveCommand, l.getOverdriveCommand)
			if err != nil {
				return err
			}

			overdrive.CharClassID, err = assignFKPtr(&overdrive.User, l.getCharacterClass)
			if err != nil {
				return err
			}

			err = qtx.UpdateOverdrive(context.Background(), database.UpdateOverdriveParams{
				DataHash:       	 generateDataHash(overdrive),
				OdCommandID: 		getNullInt32(overdrive.ODCommandID),
				CharacterClassID: 	getNullInt32(overdrive.CharClassID),
				ID: 				overdrive.ID,
			})
			if err != nil {
				return fmt.Errorf("couldn't update Overdrive: %s: %v", overdrive.Name, err)
			}
			
			err = l.createOverdriveJunctions(qtx, overdrive)
			if err != nil {
				return err
			}
		}
		return nil
	})
}


func (l *lookup) createOverdriveJunctions(qtx *database.Queries, overdrive Overdrive) error {
	for _, abilityRef := range overdrive.OverdriveAbilities {
		junction, err := createJunction(overdrive, abilityRef, l.getOverdriveAbility)
		if err != nil {
			return fmt.Errorf("couldn't create junction between overdrive %s and ability %s: %v", overdrive.Name, createLookupKey(abilityRef), err)
		}

		err = qtx.CreateOverdriveAbilityJunction(context.Background(), database.CreateOverdriveAbilityJunctionParams{
			DataHash: generateDataHash(junction),
			OverdriveID: 		junction.ParentID,
			OverdriveAbilityID: junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't create junction between overdrive %s and ability %s: %v", overdrive.Name, createLookupKey(abilityRef), err)
		}

		if overdrive.UnlockCondition == nil {
			err := l.seedDefaultOverdrive(qtx, overdrive, abilityRef)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (l *lookup) seedDefaultOverdrive(qtx *database.Queries, overdrive Overdrive, abilityRef AbilityReference) error {
	class, err := l.getCharacterClass(overdrive.User)
	if err != nil {
		return err
	}

	ability, err := l.getOverdriveAbility(abilityRef)
	if err != nil {
		return err
	}

	junction, err := createJunction(class, abilityRef, l.getOverdriveAbility)
	if err != nil {
		return err
	}

	err = qtx.CreateDefaultOverdriveAbility(context.Background(), database.CreateDefaultOverdriveAbilityParams{
		DataHash: 	generateDataHash(junction),
		ClassID: 	junction.ParentID,
		AbilityID: 	junction.ChildID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create junction between character class %s with ID %d and overdrive ability %s with ID %d: %v", overdrive.User, derefOrNil(overdrive.CharClassID), createLookupKey(abilityRef), ability.ID, err)
	}

	return nil
}