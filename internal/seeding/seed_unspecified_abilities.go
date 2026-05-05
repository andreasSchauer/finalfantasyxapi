package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop3SeedUnspecifiedAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := l.extractUnspecifiedAbilities()
	if err != nil {
		return err
	}

	params := database.CreateUnspecifiedAbilityBulkParams{
		DataHash:      make([]string, len(abilities)),
		AbilityID:     make([]int32, len(abilities)),
		Description:   make([]string, len(abilities)),
		Effect:        make([]string, len(abilities)),
		Cursor:        make([]database.NullTargetType, len(abilities)),
		TopmenuID:     make([]sql.NullInt32, len(abilities)),
		SubmenuID:     make([]sql.NullInt32, len(abilities)),
		OpenSubmenuID: make([]sql.NullInt32, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = generateDataHash(a)
		params.AbilityID[i] = a.Ability.ID
		params.Description[i] = a.Description
		params.Effect[i] = a.Effect
		params.Cursor[i] = database.ToNullTargetType(a.Cursor)
		params.TopmenuID[i] = h.GetNullInt32(a.TopmenuID)
		params.SubmenuID[i] = h.GetNullInt32(a.SubmenuID)
		params.OpenSubmenuID[i] = h.GetNullInt32(a.OpenSubmenuID)
	}

	dbRows, err := qtx.CreateUnspecifiedAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create unspecified abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		l.json.unspecifiedAbilities[i].ID = row.ID
		l.UnspecifiedAbilities[Key(abilities[i])] = abilities[i]
		l.UnspecifiedAbilitiesID[row.ID] = abilities[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractUnspecifiedAbilities() ([]UnspecifiedAbility, error) {
	abilities := []UnspecifiedAbility{}
	var err error

	for i := range l.json.unspecifiedAbilities {
		ability := &l.json.unspecifiedAbilities[i]

		ability.Ability.ID, err = l.getHashID(ability.Ability)
		if err != nil {
			return nil, err
		}

		ability.TopmenuID, err = assignFKPtr(ability.Topmenu, l.Topmenus)
		if err != nil {
			return nil, err
		}

		ability.SubmenuID, err = assignFKPtr(ability.Submenu, l.Submenus)
		if err != nil {
			return nil, err
		}

		ability.OpenSubmenuID, err = assignFKPtr(ability.OpenSubmenu, l.Submenus)
		if err != nil {
			return nil, err
		}

		abilities = append(abilities, *ability)
	}

	return dedupeRows(abilities, l.Hashes), nil
}

func (l *Lookup) completeUnspecifiedAbilities() error {
	for i := range l.json.unspecifiedAbilities {
		ability := &l.json.unspecifiedAbilities[i]

		err := l.completeBattleInteractions(ability.BattleInteractions)
		if err != nil {
			return err
		}

		l.UnspecifiedAbilities[Key(ability)] = *ability
		l.UnspecifiedAbilitiesID[ability.ID] = *ability
	}

	return nil
}

func (l *Lookup) getUnspecifiedAbilityLearnedBy(ua UnspecifiedAbility) ([]CharacterClass, error) {
	return getResources(ua.LearnedBy, l.CharClasses)
}

func (l *Lookup) seedJuncUnspecifiedAbilitiesLearnedBy(qtx *database.Queries, ctx context.Context) error {
	const desc string = "unspecified abilities + learned by"
	jParams, err := processJunctions(l, desc, l.json.unspecifiedAbilities, l.getUnspecifiedAbilityLearnedBy)
	if err != nil {
		return err
	}

	return qtx.CreateUnspecifiedAbilitiesLearnedByJunctionBulk(ctx, database.CreateUnspecifiedAbilitiesLearnedByJunctionBulkParams{
		DataHash:             jParams.DataHashes,
		UnspecifiedAbilityID: jParams.ParentIDs,
		CharacterClassID:     jParams.ChildIDs,
	})
}
