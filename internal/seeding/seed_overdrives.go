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
	TopmenuID   *int32
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
		fmt.Sprintf("%T", o),
		h.DerefOrNil(o.ODCommandID),
		h.DerefOrNil(o.CharClassID),
		o.Name,
		h.DerefOrNil(o.Version),
		o.Description,
		o.Effect,
		h.DerefOrNil(o.TopmenuID),
		o.Attributes,
		h.DerefOrNil(o.UnlockCondition),
		h.DerefOrNil(o.CountdownInSec),
		h.DerefOrNil(o.Cursor),
	}
}

func (o Overdrive) GetID() int32 {
	return o.ID
}

func (o Overdrive) Error() string {
	return fmt.Sprintf("overdrive '%s'", h.NameToString(o.Name, o.Version, o.Specification))
}

func (o Overdrive) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:      o.ID,
		Name:    o.Name,
		Version: o.Version,
	}
}

type RonsoRage struct {
	ID int32
	Overdrive
}

func (r RonsoRage) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", r),
		r.Overdrive.ID,
	}
}

func (r RonsoRage) GetID() int32 {
	return r.ID
}

func (r RonsoRage) Error() string {
	return fmt.Sprintf("ronso rage %s", r.Overdrive.Name)
}

func (r RonsoRage) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   r.ID,
		Name: r.Name,
	}
}

func (l *Lookup) seedOverdrives(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/overdrives.json"

	var overdrives []Overdrive
	err := loadJSONFile(string(srcPath), &overdrives)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, overdrive := range overdrives {
			overdrive.Attributes, err = seedObjAssignID(qtx, overdrive.Attributes, l.seedAbilityAttributes)
			if err != nil {
				return h.NewErr(overdrive.Error(), err)
			}

			dbOverdrive, err := qtx.CreateOverdrive(context.Background(), database.CreateOverdriveParams{
				DataHash:        generateDataHash(overdrive),
				Name:            overdrive.Name,
				Version:         h.GetNullInt32(overdrive.Version),
				Description:     overdrive.Description,
				Effect:          overdrive.Effect,
				AttributesID:    overdrive.Attributes.ID,
				UnlockCondition: h.GetNullString(overdrive.UnlockCondition),
				CountdownInSec:  h.GetNullInt32(overdrive.CountdownInSec),
				Cursor:          database.ToNullTargetType(overdrive.Cursor),
			})
			if err != nil {
				return h.NewErr(overdrive.Error(), err, "couldn't create overdrive")
			}

			overdrive.ID = dbOverdrive.ID

			if overdrive.User == "kimahri" {
				err = l.seedRonsoRage(qtx, overdrive)
				if err != nil {
					return err
				}
			}

			lookupObj := LookupObject{
				Name:    overdrive.Name,
				Version: overdrive.Version,
			}

			key := Key(lookupObj)
			l.Overdrives[key] = overdrive
			l.OverdrivesID[overdrive.ID] = overdrive
		}

		return nil
	})
}

func (l *Lookup) seedRonsoRage(qtx *database.Queries, overdrive Overdrive) error {
	ronsoRage := RonsoRage{
		Overdrive: overdrive,
	}

	dbRonsoRage, err := qtx.CreateRonsoRage(context.Background(), database.CreateRonsoRageParams{
		DataHash:    generateDataHash(ronsoRage),
		OverdriveID: ronsoRage.Overdrive.ID,
	})
	if err != nil {
		return h.NewErr(ronsoRage.Error(), err, "couldn't create ronso rage")
	}

	ronsoRage.ID = dbRonsoRage.ID
	l.RonsoRages[ronsoRage.Overdrive.Name] = ronsoRage
	l.RonsoRagesID[ronsoRage.ID] = ronsoRage

	return nil
}

func (l *Lookup) seedOverdrivesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/overdrives.json"

	var overdrives []Overdrive
	err := loadJSONFile(string(srcPath), &overdrives)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonOverdrive := range overdrives {
			key := LookupObject{
				Name:    jsonOverdrive.Name,
				Version: jsonOverdrive.Version,
			}
			overdrive, err := GetResource(key, l.Overdrives)
			if err != nil {
				return err
			}

			overdrive.TopmenuID, err = assignFKPtr(overdrive.Topmenu, l.Topmenus)
			if err != nil {
				return h.NewErr(overdrive.Error(), err)
			}

			overdrive.ODCommandID, err = assignFKPtr(overdrive.OverdriveCommand, l.OverdriveCommands)
			if err != nil {
				return h.NewErr(overdrive.Error(), err)
			}

			overdrive.CharClassID, err = assignFKPtr(&overdrive.User, l.CharClasses)
			if err != nil {
				return h.NewErr(overdrive.Error(), err)
			}

			err = qtx.UpdateOverdrive(context.Background(), database.UpdateOverdriveParams{
				DataHash:         generateDataHash(overdrive),
				TopmenuID:        h.GetNullInt32(overdrive.TopmenuID),
				OdCommandID:      h.GetNullInt32(overdrive.ODCommandID),
				CharacterClassID: h.GetNullInt32(overdrive.CharClassID),
				ID:               overdrive.ID,
			})
			if err != nil {
				return h.NewErr(overdrive.Error(), err, "couldn't update overdrive")
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
		junction, err := createJunction(overdrive, abilityRef.Untyped(), l.OverdriveAbilities)
		if err != nil {
			return h.NewErr(overdrive.Error(), err)
		}

		err = qtx.CreateOverdrivesOverdriveAbilitiesJunction(context.Background(), database.CreateOverdrivesOverdriveAbilitiesJunctionParams{
			DataHash:           generateDataHash(junction),
			OverdriveID:        junction.ParentID,
			OverdriveAbilityID: junction.ChildID,
		})
		if err != nil {
			subjects := h.JoinErrSubjects(overdrive.Error(), abilityRef.Error())
			return h.NewErr(subjects, err, "couldn't junction overdrive ability")
		}

		if overdrive.UnlockCondition == nil {
			err := l.seedDefaultOverdrive(qtx, overdrive, abilityRef)
			if err != nil {
				return h.NewErr(overdrive.Error(), err)
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

	junction, err := createJunction(class, abilityRef.Untyped(), l.OverdriveAbilities)
	if err != nil {
		return h.NewErr(abilityRef.Error(), err)
	}

	err = qtx.CreateDefaultOverdriveAbilityJunction(context.Background(), database.CreateDefaultOverdriveAbilityJunctionParams{
		DataHash:  generateDataHash(junction),
		ClassID:   junction.ParentID,
		AbilityID: junction.ChildID,
	})
	if err != nil {
		return h.NewErr(abilityRef.Error(), err, "couldn't create default overdrive ability")
	}

	return nil
}

func (l *Lookup) loop4SeedOverdrives(qtx *database.Queries, ctx context.Context) error {
	overdrives, err := l.extractOverdrives()
	if err != nil {
		return err
	}

	params := database.CreateOverdriveBulkParams{
		DataHash:         make([]string, len(overdrives)),
		Name:             make([]string, len(overdrives)),
		Version:          make([]sql.NullInt32, len(overdrives)),
		Description:      make([]string, len(overdrives)),
		Effect:           make([]string, len(overdrives)),
		AttributesID:     make([]int32, len(overdrives)),
		UnlockCondition:  make([]sql.NullString, len(overdrives)),
		CountdownInSec:   make([]sql.NullInt32, len(overdrives)),
		Cursor:           make([]database.NullTargetType, len(overdrives)),
		TopmenuID:        make([]sql.NullInt32, len(overdrives)),
		CharacterClassID: make([]sql.NullInt32, len(overdrives)),
		OdCommandID:      make([]sql.NullInt32, len(overdrives)),
	}

	for i, o := range overdrives {
		params.DataHash[i] = generateDataHash(o)
		params.Name[i] = o.Name
		params.Version[i] = h.GetNullInt32(o.Version)
		params.Description[i] = o.Description
		params.Effect[i] = o.Effect
		params.AttributesID[i] = o.Ability.Attributes.ID
		params.UnlockCondition[i] = h.GetNullString(o.UnlockCondition)
		params.CountdownInSec[i] = h.GetNullInt32(o.CountdownInSec)
		params.Cursor[i] = database.ToNullTargetType(o.Cursor)
		params.TopmenuID[i] = h.GetNullInt32(o.TopmenuID)
		params.CharacterClassID[i] = h.GetNullInt32(o.CharClassID)
		params.OdCommandID[i] = h.GetNullInt32(o.ODCommandID)
	}

	dbRows, err := qtx.CreateOverdriveBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create overdrives: %v", err)
	}

	for i, row := range dbRows {
		overdrives[i].ID = row.ID
		l.json.overdrives[i].ID = row.ID
		l.Overdrives[overdrives[i].Name] = overdrives[i]
		l.OverdrivesID[row.ID] = overdrives[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractOverdrives() ([]Overdrive, error) {
	overdrives := []Overdrive{}
	var err error

	for i := range l.json.overdrives {
		overdrive := &l.json.overdrives[i]

		overdrive.ODCommandID, err = assignFKPtr(overdrive.OverdriveCommand, l.OverdriveCommands)
		if err != nil {
			return nil, err
		}

		overdrive.CharClassID, err = assignFKPtr(&overdrive.User, l.CharClasses)
		if err != nil {
			return nil, err
		}

		overdrive.TopmenuID, err = assignFKPtr(overdrive.Topmenu, l.Topmenus)
		if err != nil {
			return nil, err
		}

		overdrive.Ability.Attributes.ID, err = l.getHashID(overdrive.Ability.Attributes)
		if err != nil {
			return nil, err
		}

		overdrives = append(overdrives, *overdrive)
	}

	return dedupeRows(overdrives, l.Hashes), nil
}

func (l *Lookup) getOverdriveOverdriveAbilities(o Overdrive) ([]OverdriveAbility, error) {
	return typedAbilityRefsToObjects(o.OverdriveAbilities, l.OverdriveAbilities)
}

func (l *Lookup) seedJuncOverdrivesOverdriveAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "overdrives + overdrive abilities"
	jParams, err := processJunctions(l, desc, l.json.overdrives, l.getOverdriveOverdriveAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateOverdrivesOverdriveAbilitiesJunctionBulk(ctx, database.CreateOverdrivesOverdriveAbilitiesJunctionBulkParams{
		DataHash:   		jParams.DataHashes,
		OverdriveID:  		jParams.ParentIDs,
		OverdriveAbilityID: jParams.ChildIDs,
	})
}

func (l *Lookup) loop5SeedRonsoRages(qtx *database.Queries, ctx context.Context) error {
	rages := l.extractRonsoRages()

	params := database.CreateRonsoRageBulkParams{
		DataHash:    make([]string, len(rages)),
		OverdriveID: make([]int32, len(rages)),
	}

	for i, r := range rages {
		params.DataHash[i] = generateDataHash(r)
		params.OverdriveID[i] = r.Overdrive.ID
	}

	dbRows, err := qtx.CreateRonsoRageBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create ronso rages: %v", err)
	}

	for i, row := range dbRows {
		rages[i].ID = row.ID
		l.RonsoRages[rages[i].Name] = rages[i]
		l.RonsoRagesID[row.ID] = rages[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractRonsoRages() []RonsoRage {
	rages := []RonsoRage{}

	for _, overdrive := range l.json.overdrives {
		if overdrive.User != "kimahri" {
			continue
		}

		rage := RonsoRage{
			ID:        0,
			Overdrive: overdrive,
		}

		rages = append(rages, rage)
	}

	return dedupeRows(rages, l.Hashes)
}
