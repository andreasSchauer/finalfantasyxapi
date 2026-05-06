package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop4SeedOverdrives(qtx *database.Queries, ctx context.Context) error {
	overdrives, err := l.extractOverdrives()
	if err != nil {
		return err
	}

	params := database.CreateOverdriveBulkParams{
		DataHash:         make([]string, len(overdrives)),
		Name:             make([]string, len(overdrives)),
		Version:          make([]sql.NullInt32, len(overdrives)),
		Specification:    make([]sql.NullString, len(overdrives)),
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
		params.Specification[i] = h.GetNullString(o.Specification)
		params.Description[i] = o.Description
		params.Effect[i] = o.Effect
		params.AttributesID[i] = o.Attributes.ID
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

		overdrive.Attributes.ID, err = l.getHashID(overdrive.Attributes)
		if err != nil {
			return nil, err
		}

		overdrives = append(overdrives, *overdrive)
	}

	return dedupeRows(overdrives, l.Hashes), nil
}
