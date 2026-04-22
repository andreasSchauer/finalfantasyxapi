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

func (l *Lookup) seedunspecifiedAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/unspecified_abilities.json"

	var unspecifiedAbilities []UnspecifiedAbility

	err := loadJSONFile(string(srcPath), &unspecifiedAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, UnspecifiedAbility := range unspecifiedAbilities {
			var err error
			UnspecifiedAbility.Type = database.AbilityTypeUnspecifiedAbility

			UnspecifiedAbility.Ability, err = seedObjAssignID(qtx, UnspecifiedAbility.Ability, l.seedAbility)
			if err != nil {
				return h.NewErr(UnspecifiedAbility.Error(), err)
			}

			dbUnspecifiedAbility, err := qtx.CreateUnspecifiedAbility(context.Background(), database.CreateUnspecifiedAbilityParams{
				DataHash:    generateDataHash(UnspecifiedAbility),
				AbilityID:   UnspecifiedAbility.Ability.ID,
				Description: UnspecifiedAbility.Description,
				Effect:      UnspecifiedAbility.Effect,
				Cursor:      database.ToNullTargetType(UnspecifiedAbility.Cursor),
			})
			if err != nil {
				return h.NewErr(UnspecifiedAbility.Error(), err, "couldn't create generic ability")
			}

			UnspecifiedAbility.ID = dbUnspecifiedAbility.ID
			key := CreateLookupKey(UnspecifiedAbility)
			l.UnspecifiedAbilities[key] = UnspecifiedAbility
			l.UnspecifiedAbilitiesID[UnspecifiedAbility.ID] = UnspecifiedAbility
		}
		return nil
	})
}

func (l *Lookup) seedunspecifiedAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/unspecified_abilities.json"

	var unspecifiedAbilities []UnspecifiedAbility

	err := loadJSONFile(string(srcPath), &unspecifiedAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAbility := range unspecifiedAbilities {
			abilityRef := jsonAbility.GetAbilityRef()

			unspecifiedAbility, err := GetResource(abilityRef.Untyped(), l.UnspecifiedAbilities)
			if err != nil {
				return err
			}

			err = l.seedUnspecifiedAbilityFKs(qtx, unspecifiedAbility)
			if err != nil {
				return h.NewErr(unspecifiedAbility.Error(), err)
			}

			err = l.seedUnspecifiedAbilityLearnedBy(qtx, unspecifiedAbility)
			if err != nil {
				return h.NewErr(unspecifiedAbility.Error(), err)
			}

			l.currentAbility = unspecifiedAbility.Ability

			err = l.seedBattleInteractions(qtx, l.currentAbility, unspecifiedAbility.BattleInteractions)
			if err != nil {
				return h.NewErr(unspecifiedAbility.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) seedUnspecifiedAbilityFKs(qtx *database.Queries, ability UnspecifiedAbility) error {
	var err error

	ability.TopmenuID, err = assignFKPtr(ability.Topmenu, l.Topmenus)
	if err != nil {
		return err
	}

	ability.SubmenuID, err = assignFKPtr(ability.Submenu, l.Submenus)
	if err != nil {
		return err
	}

	ability.OpenSubmenuID, err = assignFKPtr(ability.OpenSubmenu, l.Submenus)
	if err != nil {
		return err
	}

	err = qtx.UpdateUnspecifiedAbility(context.Background(), database.UpdateUnspecifiedAbilityParams{
		DataHash:      generateDataHash(ability),
		TopmenuID:     h.GetNullInt32(ability.TopmenuID),
		SubmenuID:     h.GetNullInt32(ability.SubmenuID),
		OpenSubmenuID: h.GetNullInt32(ability.OpenSubmenuID),
		ID:            ability.ID,
	})
	if err != nil {
		return h.NewErr("", err, "couldn't update generic ability")
	}

	return nil
}

func (l *Lookup) seedUnspecifiedAbilityLearnedBy(qtx *database.Queries, ability UnspecifiedAbility) error {
	for _, charClass := range ability.LearnedBy {
		junction, err := createJunction(ability, charClass, l.CharClasses)
		if err != nil {
			return err
		}

		err = qtx.CreateunspecifiedAbilitiesLearnedByJunction(context.Background(), database.CreateunspecifiedAbilitiesLearnedByJunctionParams{
			DataHash:             generateDataHash(junction),
			UnspecifiedAbilityID: junction.ParentID,
			CharacterClassID:     junction.ChildID,
		})
		if err != nil {
			return h.NewErr(charClass, err, "couldn't junction 'learned by' class")
		}
	}

	return nil
}
