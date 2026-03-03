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

func (o UnspecifiedAbility) ToHashFields() []any {
	return []any{
		o.Ability.ID,
		o.Description,
		o.Effect,
		h.DerefOrNil(o.TopmenuID),
		h.DerefOrNil(o.Cursor),
		h.DerefOrNil(o.SubmenuID),
		h.DerefOrNil(o.OpenSubmenuID),
	}
}

func (o UnspecifiedAbility) ToKeyFields() []any {
	return []any{
		o.Ability.Name,
		h.DerefOrNil(o.Ability.Version),
	}
}

func (o UnspecifiedAbility) GetID() int32 {
	return o.ID
}

func (o UnspecifiedAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        o.Name,
		Version:     o.Version,
		AbilityType: string(database.AbilityTypeUnspecifiedAbility),
	}
}

func (o UnspecifiedAbility) Error() string {
	return fmt.Sprintf("generic ability %s, version %v", o.Name, h.DerefOrNil(o.Version))
}

func (o UnspecifiedAbility) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:            o.ID,
		Name:          o.Name,
		Version:       o.Version,
		Specification: o.Specification,
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
				Cursor:      h.NullTargetType(UnspecifiedAbility.Cursor),
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
