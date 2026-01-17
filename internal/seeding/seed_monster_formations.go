package seeding

import (
	"context"
	"database/sql"
	"fmt"
	"slices"
	"sort"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type EncounterLocation struct {
	ID           int32
	Version      *int32       `json:"version"`
	LocationArea LocationArea `json:"location_area"`
	AreaID       int32
	Notes        *string            `json:"notes"`
	Formations   []MonsterFormation `json:"formations"`
}

func (el EncounterLocation) ToHashFields() []any {
	return []any{
		h.DerefOrNil(el.Version),
		el.AreaID,
		h.DerefOrNil(el.Notes),
	}
}

func (el EncounterLocation) ToKeyFields() []any {
	return []any{
		h.DerefOrNil(el.Version),
		CreateLookupKey(el.LocationArea),
	}
}

func (el EncounterLocation) GetID() int32 {
	return el.ID
}

func (el EncounterLocation) Error() string {
	return fmt.Sprintf("encounter location with version: %v, %s", h.DerefOrNil(el.Version), el.LocationArea)
}

type MonsterFormation struct {
	ID                  int32
	EncounterLocationID int32
	Monsters            []MonsterAmount    `json:"monsters"`
	Category            string             `json:"category"`
	IsForcedAmbush      bool               `json:"is_forced_ambush"`
	CanEscape           bool               `json:"can_escape"`
	BossMusic           *FormationBossSong `json:"boss_music"`
	Notes               *string            `json:"notes"`
	TriggerCommands     []AbilityReference `json:"trigger_commands"`
}

func (mf MonsterFormation) ToHashFields() []any {
	ownFields := []any{
		mf.Category,
		mf.IsForcedAmbush,
		mf.CanEscape,
		h.ObjPtrToID(mf.BossMusic),
		h.DerefOrNil(mf.Notes),
	}

	monsters := mf.Monsters

	sort.SliceStable(monsters, func(i, j int) bool { return monsters[i].MonsterID < monsters[j].MonsterID })
	monsterKeys := []any{}

	for _, mon := range mf.Monsters {
		key := combineFields(mon.ToFormationHashFields())
		monsterKeys = append(monsterKeys, key)
	}

	return slices.Concat(ownFields, monsterKeys)
}

func (mf MonsterFormation) GetID() int32 {
	return mf.ID
}

func (mf MonsterFormation) Error() string {
	return fmt.Sprintf("monster formation with location id: %d, category: %s, forced ambush: %t, can escape: %t, boss music id: %v, notes: %v", mf.EncounterLocationID, mf.Category, mf.IsForcedAmbush, mf.CanEscape, h.ObjPtrToID(mf.BossMusic), h.DerefOrNil(mf.Notes))
}

func (mf MonsterFormation) GetResParamsUnnamed() h.ResParamsUnnamed {
	return h.ResParamsUnnamed{
		ID: mf.ID,
	}
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

func (s FormationBossSong) Error() string {
	return fmt.Sprintf("formation boss song %s, celebrate victory: %t", s.Song, s.CelebrateVictory)
}

func (l *Lookup) seedEncounterLocations(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monster_formations.json"

	var encounterLocations []EncounterLocation
	err := loadJSONFile(string(srcPath), &encounterLocations)
	if err != nil {
		return err
	}
	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, encounterLocation := range encounterLocations {
			var err error

			locationArea := encounterLocation.LocationArea
			encounterLocation.AreaID, err = assignFK(locationArea, l.Areas)
			if err != nil {
				return h.NewErr(encounterLocation.Error(), err)
			}

			dbEncounterLocation, err := qtx.CreateEncounterLocation(context.Background(), database.CreateEncounterLocationParams{
				DataHash: generateDataHash(encounterLocation),
				Version:  h.GetNullInt32(encounterLocation.Version),
				AreaID:   encounterLocation.AreaID,
				Notes:    h.GetNullString(encounterLocation.Notes),
			})
			if err != nil {
				return h.NewErr(encounterLocation.Error(), err, "couldn't create monster encounter location")
			}

			encounterLocation.ID = dbEncounterLocation.ID
			key := CreateLookupKey(encounterLocation)
			l.EncounterLocations[key] = encounterLocation
		}
		return nil
	})
}

func (l *Lookup) seedMonsterFormationsRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monster_formations.json"

	var encounterLocations []EncounterLocation
	err := loadJSONFile(string(srcPath), &encounterLocations)
	if err != nil {
		return err
	}
	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonEncounterLocation := range encounterLocations {
			key := CreateLookupKey(jsonEncounterLocation)
			encounterLocation, err := GetResource(key, l.EncounterLocations)
			if err != nil {
				return err
			}

			for _, monsterFormation := range encounterLocation.Formations {
				var err error
				monsterFormation.EncounterLocationID, err = assignFK(key, l.EncounterLocations)
				if err != nil {
					return h.NewErr(monsterFormation.Error(), err)
				}

				junction, err := createJunctionSeed(qtx, encounterLocation, monsterFormation, l.seedMonsterFormation)
				if err != nil {
					return h.NewErr(encounterLocation.Error(), err)
				}

				err = qtx.CreateEncounterLocationFormationsJunction(context.Background(), database.CreateEncounterLocationFormationsJunctionParams{
					DataHash:            generateDataHash(junction),
					EncounterLocationID: junction.ParentID,
					MonsterFormationID:  junction.ChildID,
				})
				if err != nil {
					subjects := h.JoinSubjects(encounterLocation.Error(), monsterFormation.Error())
					return h.NewErr(subjects, err, "couldn't junction encounter location with monster formation")
				}
			}

		}

		return nil
	})
}

func (l *Lookup) seedMonsterFormation(qtx *database.Queries, formation MonsterFormation) (MonsterFormation, error) {
	var err error

	formation.BossMusic, err = seedObjPtrAssignFK(qtx, formation.BossMusic, l.seedFormationBossSong)
	if err != nil {
		return MonsterFormation{}, h.NewErr(formation.Error(), err)
	}

	dbFormation, err := qtx.CreateMonsterFormation(context.Background(), database.CreateMonsterFormationParams{
		DataHash:       generateDataHash(formation),
		Category:       database.MonsterFormationCategory(formation.Category),
		IsForcedAmbush: formation.IsForcedAmbush,
		CanEscape:      formation.CanEscape,
		BossSongID:     h.ObjPtrToNullInt32ID(formation.BossMusic),
		Notes:          h.GetNullString(formation.Notes),
	})
	if err != nil {
		return MonsterFormation{}, h.NewErr(formation.Error(), err, "couldn't create monster formation")
	}
	formation.ID = dbFormation.ID
	l.MonsterFormationsID[formation.ID] = formation

	err = l.seedFormationMonsterAmounts(qtx, formation)
	if err != nil {
		return MonsterFormation{}, h.NewErr(formation.Error(), err)
	}

	err = l.seedFormationTriggerCommands(qtx, formation)
	if err != nil {
		return MonsterFormation{}, h.NewErr(formation.Error(), err)
	}

	return formation, nil
}

func (l *Lookup) seedFormationBossSong(qtx *database.Queries, bossSong FormationBossSong) (FormationBossSong, error) {
	var err error

	bossSong.SongID, err = assignFK(bossSong.Song, l.Songs)
	if err != nil {
		return FormationBossSong{}, h.NewErr(bossSong.Error(), err)
	}

	dbBossSong, err := qtx.CreateFormationBossSong(context.Background(), database.CreateFormationBossSongParams{
		DataHash:         generateDataHash(bossSong),
		SongID:           bossSong.SongID,
		CelebrateVictory: bossSong.CelebrateVictory,
	})
	if err != nil {
		return FormationBossSong{}, h.NewErr(bossSong.Error(), err, "couldn't create formation boss song")
	}
	bossSong.ID = dbBossSong.ID

	return bossSong, nil
}

func (l *Lookup) seedFormationMonsterAmounts(qtx *database.Queries, formation MonsterFormation) error {
	for _, monsterAmount := range formation.Monsters {
		var err error
		key := CreateLookupKey(monsterAmount)
		monsterAmount.MonsterID, err = assignFK(key, l.Monsters)
		if err != nil {
			return err
		}

		junction, err := createJunctionSeed(qtx, formation, monsterAmount, l.seedMonsterAmount)
		if err != nil {
			return err
		}

		err = qtx.CreateMonsterFormationsMonstersJunction(context.Background(), database.CreateMonsterFormationsMonstersJunctionParams{
			DataHash:           generateDataHash(junction),
			MonsterFormationID: junction.ParentID,
			MonsterAmountID:    junction.ChildID,
		})
		if err != nil {
			return h.NewErr(monsterAmount.Error(), err, "couldn't junction monster amount")
		}
	}

	return nil
}

func (l *Lookup) seedFormationTriggerCommands(qtx *database.Queries, formation MonsterFormation) error {
	for _, abilityRef := range formation.TriggerCommands {
		junction, err := createJunction(formation, abilityRef.Untyped(), l.TriggerCommands)
		if err != nil {
			return err
		}

		err = qtx.CreateMonsterFormationsTriggerCommandsJunction(context.Background(), database.CreateMonsterFormationsTriggerCommandsJunctionParams{
			DataHash:           generateDataHash(junction),
			MonsterFormationID: junction.ParentID,
			TriggerCommandID:   junction.ChildID,
		})
		if err != nil {
			return h.NewErr(abilityRef.Error(), err, "couldn't junction with trigger command")
		}
	}

	return nil
}
