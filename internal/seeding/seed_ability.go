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
	*Attributes
}

func (a Ability) ToHashFields() []any {
	return []any{
		a.Name,
		derefOrNil(a.Version),
		derefOrNil(a.Specification),
		a.Type,
		ObjPtrToHashID(a.Attributes),
	}
}

func (a Ability) ToKeyFields() []any {
	return []any{
		a.Name,
		derefOrNil(a.Version),
		a.Type,
	}
}

func (a Ability) GetID() int32 {
	return a.ID
}

type Attributes struct {
	ID               int32
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

func (a Attributes) GetID() int32 {
	return a.ID
}

func (l *lookup) seedAbility(qtx *database.Queries, ability Ability) (Ability, error) {
	if ability.Type != database.AbilityTypeOverdriveAbility {
		var err error

		ability.Attributes, err = seedObjPtrAssignFK(qtx, ability.Attributes, l.seedAbilityAttributes)
		if err != nil {
			return Ability{}, fmt.Errorf("couldn't create Ability Attributes: %s-%d, type: %s: %v", ability.Name, derefOrNil(ability.Version), ability.Type, err)
		}
	}

	dbAbility, err := qtx.CreateAbility(context.Background(), database.CreateAbilityParams{
		DataHash:      generateDataHash(ability),
		Name:          ability.Name,
		Version:       getNullInt32(ability.Version),
		Specification: getNullString(ability.Specification),
		AttributesID:  ObjPtrToNullInt32ID(ability.Attributes),
		Type:          ability.Type,
	})
	if err != nil {
		return Ability{}, fmt.Errorf("couldn't create Ability: %s-%d, type: %s: %v", ability.Name, derefOrNil(ability.Version), ability.Type, err)
	}

	ability.ID = dbAbility.ID

	return ability, nil
}

func (l *lookup) seedAbilityAttributes(qtx *database.Queries, attributes Attributes) (Attributes, error) {
	dbAttributes, err := qtx.CreateAbilityAttributes(context.Background(), database.CreateAbilityAttributesParams{
		DataHash:         generateDataHash(attributes),
		Rank:             getNullInt32(attributes.Rank),
		AppearsInHelpBar: attributes.AppearsInHelpBar,
		CanCopycat:       attributes.CanCopycat,
	})
	if err != nil {
		return Attributes{}, err
	}

	attributes.ID = dbAttributes.ID

	return attributes, nil
}
