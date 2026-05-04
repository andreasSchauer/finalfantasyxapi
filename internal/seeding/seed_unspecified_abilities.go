package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type UnspecifiedAbility struct {
	ID int32
	Ability
	TopmenuID          *int32
	SubmenuID          *int32
	OpenSubmenuID      *int32
	Description        string              `json:"description"`
	Effect             string              `json:"effect"`
	RelatedStats       []string            `json:"related_stats"`
	Topmenu            *string             `json:"topmenu"`
	Submenu            *string             `json:"submenu"`
	OpenSubmenu        *string             `json:"open_submenu"`
	LearnedBy          []string            `json:"learned_by"`
	Cursor             *string             `json:"cursor"`
	BattleInteractions []BattleInteraction `json:"battle_interactions"`
}

func (u UnspecifiedAbility) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", u),
		u.Ability.ID,
		u.Description,
		u.Effect,
		h.DerefOrNil(u.TopmenuID),
		h.DerefOrNil(u.Cursor),
		h.DerefOrNil(u.SubmenuID),
		h.DerefOrNil(u.OpenSubmenuID),
	}
}

func (u UnspecifiedAbility) ToKeyFields() []any {
	return []any{
		u.Ability.Name,
		h.DerefOrNil(u.Ability.Version),
	}
}

func (u UnspecifiedAbility) GetID() int32 {
	return u.ID
}

func (u UnspecifiedAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        u.Name,
		Version:     u.Version,
		AbilityType: string(database.AbilityTypeUnspecifiedAbility),
	}
}

func (u UnspecifiedAbility) Error() string {
	return fmt.Sprintf("unspecified ability '%s'", h.NameToString(u.Name, u.Version, u.Specification))
}

func (u UnspecifiedAbility) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:            u.ID,
		Name:          u.Name,
		Version:       u.Version,
		Specification: u.Specification,
	}
}

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
		key := Key(abilities[i])
		l.UnspecifiedAbilities[key] = abilities[i]
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
