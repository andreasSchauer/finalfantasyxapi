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

func (o Overdrive) ToKeyFields() []any {
	return []any{
		o.Name,
		h.DerefOrNil(o.Version),
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
		l.Overdrives[Key(overdrives[i])] = overdrives[i]
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
	return typedRefsToResources(o.OverdriveAbilities, l.OverdriveAbilities)
}

func (l *Lookup) seedJuncOverdrivesOverdriveAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "overdrives + overdrive abilities"
	jParams, err := processJunctions(l, desc, l.json.overdrives, l.getOverdriveOverdriveAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateOverdrivesOverdriveAbilitiesJunctionBulk(ctx, database.CreateOverdrivesOverdriveAbilitiesJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		OverdriveID:        jParams.ParentIDs,
		OverdriveAbilityID: jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncDefaultOverdriveAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "default overdrive + overdrive abilities"
	params := database.CreateDefaultOverdriveAbilityJunctionBulkParams{
		DataHash: 	make([]string, 0),
		ClassID: 	make([]int32, 0),
		AbilityID: 	make([]int32, 0),
	}

	for _, overdrive := range l.json.overdrives {
		if overdrive.UnlockCondition != nil {
			continue
		}

		class, err := GetResource(overdrive.User, l.CharClasses)
		if err != nil {
			return err
		}

		for _, ref := range overdrive.OverdriveAbilities {
			oa, err := GetResource(ref.Untyped(), l.OverdriveAbilities)
			if err != nil {
				return err
			}

			j := StdJunction{
				ParentID: 	class.ID,
				ChildID: 	oa.ID,
			}
			dataHash := generateJunctionHash(j, desc)

			params.DataHash = append(params.DataHash, dataHash)
			params.ClassID = append(params.ClassID, class.ID)
			params.AbilityID = append(params.AbilityID, oa.ID)
		}
	}

	return qtx.CreateDefaultOverdriveAbilityJunctionBulk(ctx, params)
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
