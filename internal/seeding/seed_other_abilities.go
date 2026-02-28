package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type OtherAbility struct {
	ID int32
	Ability
	TopmenuID		   *int32
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

func (o OtherAbility) ToHashFields() []any {
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

func (o OtherAbility) ToKeyFields() []any {
	return []any{
		o.Ability.Name,
		h.DerefOrNil(o.Ability.Version),
	}
}

func (o OtherAbility) GetID() int32 {
	return o.ID
}

func (o OtherAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        o.Name,
		Version:     o.Version,
		AbilityType: string(database.AbilityTypeOtherAbility),
	}
}

func (o OtherAbility) Error() string {
	return fmt.Sprintf("generic ability %s, version %v", o.Name, h.DerefOrNil(o.Version))
}

func (o OtherAbility) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:            o.ID,
		Name:          o.Name,
		Version:       o.Version,
		Specification: o.Specification,
	}
}

func (l *Lookup) seedotherAbilities(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/other_abilities.json"

	var otherAbilities []OtherAbility

	err := loadJSONFile(string(srcPath), &otherAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, OtherAbility := range otherAbilities {
			var err error
			OtherAbility.Type = database.AbilityTypeOtherAbility

			OtherAbility.Ability, err = seedObjAssignID(qtx, OtherAbility.Ability, l.seedAbility)
			if err != nil {
				return h.NewErr(OtherAbility.Error(), err)
			}

			dbOtherAbility, err := qtx.CreateOtherAbility(context.Background(), database.CreateOtherAbilityParams{
				DataHash:    generateDataHash(OtherAbility),
				AbilityID:   OtherAbility.Ability.ID,
				Description: OtherAbility.Description,
				Effect:      OtherAbility.Effect,
				Cursor:      h.NullTargetType(OtherAbility.Cursor),
			})
			if err != nil {
				return h.NewErr(OtherAbility.Error(), err, "couldn't create generic ability")
			}

			OtherAbility.ID = dbOtherAbility.ID
			key := CreateLookupKey(OtherAbility)
			l.OtherAbilities[key] = OtherAbility
			l.OtherAbilitiesID[OtherAbility.ID] = OtherAbility
		}
		return nil
	})
}

func (l *Lookup) seedotherAbilitiesRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/other_abilities.json"

	var otherAbilities []OtherAbility

	err := loadJSONFile(string(srcPath), &otherAbilities)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonAbility := range otherAbilities {
			abilityRef := jsonAbility.GetAbilityRef()

			otherAbility, err := GetResource(abilityRef.Untyped(), l.OtherAbilities)
			if err != nil {
				return err
			}

			err = l.seedOtherAbilityFKs(qtx, otherAbility)
			if err != nil {
				return h.NewErr(otherAbility.Error(), err)
			}

			err = l.seedOtherAbilityLearnedBy(qtx, otherAbility)
			if err != nil {
				return h.NewErr(otherAbility.Error(), err)
			}

			l.currentAbility = otherAbility.Ability

			err = l.seedBattleInteractions(qtx, l.currentAbility, otherAbility.BattleInteractions)
			if err != nil {
				return h.NewErr(otherAbility.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) seedOtherAbilityFKs(qtx *database.Queries, ability OtherAbility) error {
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

	err = qtx.UpdateOtherAbility(context.Background(), database.UpdateOtherAbilityParams{
		DataHash:      generateDataHash(ability),
		TopmenuID: 	   h.GetNullInt32(ability.TopmenuID),
		SubmenuID:     h.GetNullInt32(ability.SubmenuID),
		OpenSubmenuID: h.GetNullInt32(ability.OpenSubmenuID),
		ID:            ability.ID,
	})
	if err != nil {
		return h.NewErr("", err, "couldn't update generic ability")
	}

	return nil
}

func (l *Lookup) seedOtherAbilityLearnedBy(qtx *database.Queries, ability OtherAbility) error {
	for _, charClass := range ability.LearnedBy {
		junction, err := createJunction(ability, charClass, l.CharClasses)
		if err != nil {
			return err
		}

		err = qtx.CreateotherAbilitiesLearnedByJunction(context.Background(), database.CreateotherAbilitiesLearnedByJunctionParams{
			DataHash:         generateDataHash(junction),
			OtherAbilityID:   junction.ParentID,
			CharacterClassID: junction.ChildID,
		})
		if err != nil {
			return h.NewErr(charClass, err, "couldn't junction 'learned by' class")
		}
	}

	return nil
}
