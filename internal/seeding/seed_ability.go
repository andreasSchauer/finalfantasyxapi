package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Ability struct {
	ID            int32
	Name          string  `json:"name"`
	Version       *int32  `json:"version"`
	Specification *string `json:"specification"`
	Type          database.AbilityType
	Attributes
}

func (a Ability) ToHashFields() []any {
	return []any{
		a.Name,
		derefOrNil(a.Version),
		derefOrNil(a.Specification),
		a.Type,
		derefOrNil(a.Attributes.ID),
	}
}

func (a Ability) ToKeyFields() []any {
	return []any{
		a.Name,
		derefOrNil(a.Version),
		a.Type,
	}
}


type Attributes struct {
	ID               *int32
	Rank             *int32 `json:"rank"`
	AppearsInHelpBar bool   `json:"appears_in_help_bar"`
	CanCopycat       bool   `json:"can_copycat"`
}

func (a Attributes) ToHashFields() []any {
	return []any{
		derefOrNil(a.Rank),
		a.AppearsInHelpBar,
		a.CanCopycat,
	}
}

func (l *lookup) seedAbility(qtx *database.Queries, ability Ability) (database.Ability, error) {
	if ability.Type != database.AbilityTypeOverdriveAbility {
		dbAttributes, err := l.seedAbilityAttributes(qtx, ability)
		if err != nil {
			return database.Ability{}, err
		}

		ability.Attributes.ID = &dbAttributes.ID
	}

	dbAbility, err := qtx.CreateAbility(context.Background(), database.CreateAbilityParams{
		DataHash:      generateDataHash(ability),
		Name:          ability.Name,
		Version:       getNullInt32(ability.Version),
		Specification: getNullString(ability.Specification),
		AttributesID:  getNullInt32(ability.Attributes.ID),
		Type:          ability.Type,
	})
	if err != nil {
		return database.Ability{}, fmt.Errorf("couldn't create Ability: %s-%d, type: %s: %v", ability.Name, derefOrNil(ability.Version), ability.Type, err)
	}

	return dbAbility, nil
}

func (l *lookup) seedAbilityAttributes(qtx *database.Queries, ability Ability) (database.AbilityAttribute, error) {
	attributes := ability.Attributes

	dbAttributes, err := qtx.CreateAbilityAttributes(context.Background(), database.CreateAbilityAttributesParams{
		DataHash:         generateDataHash(attributes),
		Rank:             getNullInt32(attributes.Rank),
		AppearsInHelpBar: attributes.AppearsInHelpBar,
		CanCopycat:       attributes.CanCopycat,
	})
	if err != nil {
		return database.AbilityAttribute{}, fmt.Errorf("couldn't create Ability Attributes: %s-%d, type: %s: %v", ability.Name, *ability.Version, ability.Type, err)
	}

	return dbAttributes, nil
}
