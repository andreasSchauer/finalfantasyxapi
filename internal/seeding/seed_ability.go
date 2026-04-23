package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Ability struct {
	ID            int32
	Name          string `json:"name"`
	Version       *int32 `json:"version"`
	Type          database.AbilityType
	Specification *string `json:"specification"`
	Attributes
}

func (a Ability) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.Name,
		h.DerefOrNil(a.Version),
		h.DerefOrNil(a.Specification),
		a.Type,
		a.Attributes,
	}
}

func (a Ability) ToKeyFields() []any {
	return []any{
		a.Name,
		h.DerefOrNil(a.Version),
		a.Type,
	}
}

func (a Ability) GetID() int32 {
	return a.ID
}

func (a Ability) Error() string {
	return fmt.Sprintf("ability '%s', type %s", h.NameToString(a.Name, a.Version, a.Specification), a.Type)
}

func (a Ability) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        a.Name,
		Version:     a.Version,
		AbilityType: string(a.Type),
	}
}

func (a Ability) GetResParamsTyped() h.ResParamsTyped {
	return h.ResParamsTyped{
		ID:            a.ID,
		Name:          a.Name,
		Version:       a.Version,
		Specification: a.Specification,
		Type:          string(a.Type),
	}
}

type AbilityReference struct {
	Name        string `json:"name"`
	Version     *int32 `json:"version"`
	AbilityType string `json:"ability_type"`
}

func (a AbilityReference) ToKeyFields() []any {
	return []any{
		a.Name,
		h.DerefOrNil(a.Version),
		a.AbilityType,
	}
}

func (a AbilityReference) Error() string {
	return fmt.Sprintf("ability reference '%s', type %s", h.NameToString(a.Name, a.Version, nil), a.AbilityType)
}


func (a AbilityReference) Untyped() UntypedAbilityRef {
	return UntypedAbilityRef{
		Name:    a.Name,
		Version: a.Version,
	}
}

type UntypedAbilityRef struct {
	Name    string
	Version *int32
}

func (a UntypedAbilityRef) ToKeyFields() []any {
	return []any{
		a.Name,
		h.DerefOrNil(a.Version),
	}
}

func (a UntypedAbilityRef) Error() string {
	return fmt.Sprintf("untyped ability reference '%s'", h.NameToString(a.Name, a.Version, nil))
}

type Attributes struct {
	ID               int32
	Rank             *int32 `json:"rank"`
	AppearsInHelpBar bool   `json:"appears_in_help_bar"`
	CanCopycat       bool   `json:"can_copycat"`
}

func (a Attributes) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		h.DerefOrNil(a.Rank),
		a.AppearsInHelpBar,
		a.CanCopycat,
	}
}

func (a Attributes) GetID() int32 {
	return a.ID
}

func (a Attributes) Error() string {
	return fmt.Sprintf("ability attributes with rank: %v, help bar: %t, copycat: %t", h.PtrToString(a.Rank), a.AppearsInHelpBar, a.CanCopycat)
}

func (l *Lookup) seedAbility(qtx *database.Queries, ability Ability) (Ability, error) {
	var err error

	ability.Attributes, err = seedObjAssignID(qtx, ability.Attributes, l.seedAbilityAttributes)
	if err != nil {
		return Ability{}, h.NewErr(ability.Error(), err)
	}

	dbAbility, err := qtx.CreateAbility(context.Background(), database.CreateAbilityParams{
		DataHash:      generateDataHash(ability),
		Name:          ability.Name,
		Version:       h.GetNullInt32(ability.Version),
		Specification: h.GetNullString(ability.Specification),
		AttributesID:  ability.Attributes.ID,
		Type:          ability.Type,
	})
	if err != nil {
		return Ability{}, h.NewErr(ability.Error(), err, "couldn't create ability")
	}

	ability.ID = dbAbility.ID
	key := CreateLookupKey(ability)
	l.Abilities[key] = ability
	l.AbilitiesID[ability.ID] = ability

	return ability, nil
}

func (l *Lookup) seedAbilityAttributes(qtx *database.Queries, attributes Attributes) (Attributes, error) {
	dbAttributes, err := qtx.CreateAbilityAttributes(context.Background(), database.CreateAbilityAttributesParams{
		DataHash:         generateDataHash(attributes),
		Rank:             h.GetNullInt32(attributes.Rank),
		AppearsInHelpBar: attributes.AppearsInHelpBar,
		CanCopycat:       attributes.CanCopycat,
	})
	if err != nil {
		return Attributes{}, h.NewErr(attributes.Error(), err, "couldn't create ability attributes")
	}

	attributes.ID = dbAttributes.ID

	return attributes, nil
}
