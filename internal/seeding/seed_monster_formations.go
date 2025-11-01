package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type FormationLocation struct {
	ID           int32
	Version      *int32       `json:"version"`
	LocationArea LocationArea `json:"location_area"`
	AreaID       int32
	Notes        *string            `json:"notes"`
	Formations   []MonsterFormation `json:"formations"`
}

func (fl FormationLocation) ToHashFields() []any {
	return []any{
		derefOrNil(fl.Version),
		fl.AreaID,
		derefOrNil(fl.Notes),
	}
}

func (fl FormationLocation) ToKeyFields() []any {
	return []any{
		derefOrNil(fl.Version),
		createLookupKey(fl.LocationArea),
	}
}

func (fl FormationLocation) GetID() int32 {
	return fl.ID
}

type MonsterFormation struct {
	ID                  int32
	FormationLocationID int32
	Monsters            []MonsterAmount    `json:"monsters"`
	Category            string             `json:"category"`
	IsForcedAmbush      bool               `json:"is_forced_ambush"`
	CanEscape           bool               `json:"can_escape"`
	BossMusic           *FormationBossSong `json:"boss_music"`
	Notes               *string            `json:"notes"`
	TriggerCommands     []AbilityReference `json:"trigger_commands"`
}

func (mf MonsterFormation) ToHashFields() []any {
	return []any{
		mf.FormationLocationID,
		mf.Category,
		mf.IsForcedAmbush,
		mf.CanEscape,
		ObjPtrToHashID(mf.BossMusic),
		derefOrNil(mf.Notes),
	}
}

func (mf MonsterFormation) GetID() int32 {
	return mf.ID
}


type FormationBossSong struct {
	ID               int32
	SongID           int32
	Song             string `json:"music"`
	CelebrateVictory bool   `json:"celebrate_victory"`
}

func (s FormationBossSong) ToHashFields() []any {
	return []any{
		s.SongID,
		s.CelebrateVictory,
	}
}

func (s FormationBossSong) GetID() int32 {
	return s.ID
}

func (l *lookup) seedFormationLocations(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monster_formations.json"

	var formationLocations []FormationLocation
	err := loadJSONFile(string(srcPath), &formationLocations)
	if err != nil {
		return err
	}
	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, formationLocation := range formationLocations {
			var err error

			locationArea := formationLocation.LocationArea
			formationLocation.AreaID, err = assignFK(locationArea, l.getArea)
			if err != nil {
				return fmt.Errorf("monster formations: %v", err)
			}

			dbFormationLocation, err := qtx.CreateFormationLocation(context.Background(), database.CreateFormationLocationParams{
				DataHash: generateDataHash(formationLocation),
				Version:  getNullInt32(formationLocation.Version),
				AreaID:   formationLocation.AreaID,
				Notes:    getNullString(formationLocation.Notes),
			})
			if err != nil {
				return fmt.Errorf("couldn't create monster formation list: %s - version: %d: %v", createLookupKey(locationArea), derefOrNil(formationLocation.Version), err)
			}

			formationLocation.ID = dbFormationLocation.ID
			key := createLookupKey(formationLocation)
			l.formationLocations[key] = formationLocation
		}
		return nil
	})
}

func (l *lookup) seedMonsterFormationsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monster_formations.json"

	var formationLocations []FormationLocation
	err := loadJSONFile(string(srcPath), &formationLocations)
	if err != nil {
		return err
	}
	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonFormationLocation := range formationLocations {
			key := createLookupKey(jsonFormationLocation)
			formationLocation, err := l.getFormationLocation(key)
			if err != nil {
				return err
			}

			for _, monsterFormation := range formationLocation.Formations {
				var err error
				monsterFormation.FormationLocationID, err = assignFK(key, l.getFormationLocation)
				if err != nil {
					return err
				}

				junction, err := createJunctionSeed(qtx, formationLocation, monsterFormation, l.seedMonsterFormation)
				if err != nil {
					return fmt.Errorf("%s: %v", createLookupKey(formationLocation), err)
				}

				err = qtx.CreateEncounterLocationFormationsJunction(context.Background(), database.CreateEncounterLocationFormationsJunctionParams{
					DataHash:            generateDataHash(junction),
					EncounterLocationID: junction.ParentID,
					MonsterFormationID:  junction.ChildID,
				})
				if err != nil {
					return fmt.Errorf("%s: %v", createLookupKey(formationLocation), err)
				}
			}

		}

		return nil
	})
}

func (l *lookup) seedMonsterFormation(qtx *database.Queries, formation MonsterFormation) (MonsterFormation, error) {
	var err error

	formation.BossMusic, err = seedObjPtrAssignFK(qtx, formation.BossMusic, l.seedFormationBossSong)
	if err != nil {
		return MonsterFormation{}, err
	}

	dbFormation, err := qtx.CreateMonsterFormation(context.Background(), database.CreateMonsterFormationParams{
		DataHash:            generateDataHash(formation),
		EncounterLocationID: formation.FormationLocationID,
		Category:            database.MonsterFormationCategory(formation.Category),
		IsForcedAmbush:      formation.IsForcedAmbush,
		CanEscape:           formation.CanEscape,
		BossSongID:          ObjPtrToNullInt32ID(formation.BossMusic),
		Notes:               getNullString(formation.Notes),
	})
	if err != nil {
		return MonsterFormation{}, fmt.Errorf("couldn't create monster formation: %v", err)
	}
	formation.ID = dbFormation.ID

	err = l.seedFormationMonsterAmounts(qtx, formation)
	if err != nil {
		return MonsterFormation{}, err
	}

	err = l.seedFormationTriggerCommands(qtx, formation)
	if err != nil {
		return MonsterFormation{}, err
	}

	return formation, nil
}


func (l *lookup) seedFormationBossSong(qtx *database.Queries, bossSong FormationBossSong) (FormationBossSong, error) {
	var err error

	bossSong.SongID, err = assignFK(bossSong.Song, l.getSong)
	if err != nil {
		return FormationBossSong{}, err
	}

	dbBossSong, err := qtx.CreateFormationBossSong(context.Background(), database.CreateFormationBossSongParams{
		DataHash:         generateDataHash(bossSong),
		SongID:           bossSong.SongID,
		CelebrateVictory: bossSong.CelebrateVictory,
	})
	if err != nil {
		return FormationBossSong{}, fmt.Errorf("couldn't create formation boss song: %v", err)
	}
	bossSong.ID = dbBossSong.ID

	return bossSong, nil
}


func (l *lookup) seedFormationMonsterAmounts(qtx *database.Queries, formation MonsterFormation) error {
	for _, monsterAmount := range formation.Monsters {
		var err error
		key := createLookupKey(monsterAmount)
		monsterAmount.MonsterID, err = assignFK(key, l.getMonster)
		if err != nil {
			return err
		}

		junction, err := createJunctionSeed(qtx, formation, monsterAmount, l.seedMonsterAmount)
		if err != nil {
			return fmt.Errorf("couldn't create junction with Monster Amount: %s: %v", key, err)
		}

		err = qtx.CreateMonsterFormationsMonstersJunction(context.Background(), database.CreateMonsterFormationsMonstersJunctionParams{
			DataHash:           generateDataHash(junction),
			MonsterFormationID: junction.ParentID,
			MonsterAmountID:    junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't seed junction with Monster Amount: %s: %v", key, err)
		}
	}

	return nil
}


func (l *lookup) seedFormationTriggerCommands(qtx *database.Queries, formation MonsterFormation) error {
	for _, abilityRef := range formation.TriggerCommands {
		junction, err := createJunction(formation, abilityRef, l.getTriggerCommand)
		if err != nil {
			return fmt.Errorf("couldn't create junction with trigger command %s: %v", createLookupKey(abilityRef), err)
		}

		err = qtx.CreateMonsterFormationsTriggerCommandsJunction(context.Background(), database.CreateMonsterFormationsTriggerCommandsJunctionParams{
			DataHash:           generateDataHash(junction),
			MonsterFormationID: junction.ParentID,
			TriggerCommandID:   junction.ChildID,
		})
		if err != nil {
			return fmt.Errorf("couldn't seed junction with trigger command %s: %v", createLookupKey(abilityRef), err)
		}
	}

	return nil
}