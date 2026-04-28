package seeding

import (
	"context"
	"database/sql"
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


func (l *Lookup) loop2SeedAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := l.extractAbilities()
	if err != nil {
		return err
	}

	params := database.CreateAbilityBulkParams{
		DataHash:      make([]string, len(abilities)),
		Name:          make([]string, len(abilities)),
		Version:       make([]sql.NullInt32, len(abilities)),
		Specification: make([]sql.NullString, len(abilities)),
		AttributesID:  make([]int32, len(abilities)),
		Type:          make([]database.AbilityType, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = generateDataHash(a)
		params.Name[i] = a.Name
		params.Version[i] = h.GetNullInt32(a.Version)
		params.Specification[i] = h.GetNullString(a.Specification)
		params.AttributesID[i] = a.Attributes.ID
		params.Type[i] = a.Type
	}

	dbRows, err := qtx.CreateAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		key := CreateLookupKey(abilities[i])
		l.Abilities[key] = abilities[i]
		l.AbilitiesID[row.ID] = abilities[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAbilities() ([]Ability, error) {
	abilities := []Ability{}
	var err error

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]
		ability.Attributes.ID, err = l.getHashID(ability.Attributes)
		if err != nil {
			return nil, err
		}

		ability.Type = database.AbilityTypePlayerAbility
		abilities = append(abilities, ability.Ability)
	}

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]
		ability.Attributes.ID, err = l.getHashID(ability.Attributes)
		if err != nil {
			return nil, err
		}

		ability.Type = database.AbilityTypeOverdriveAbility
		abilities = append(abilities, ability.Ability)
	}

	for i := range l.json.items {
		item := &l.json.items[i]
		if len(item.BattleInteractions) > 0 {
			item.Attributes.ID, err = l.getHashID(item.Attributes)
			if err != nil {
				return nil, err
			}
			item.Ability.Name = item.Name
			item.Ability.Type = database.AbilityTypeItemAbility
			abilities = append(abilities, item.Ability)
		}
	}

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]
		command.Attributes.ID, err = l.getHashID(command.Attributes)
		if err != nil {
			return nil, err
		}

		command.Type = database.AbilityTypeTriggerCommand
		abilities = append(abilities, command.Ability)
	}

	for i := range l.json.unspecifiedAbilities {
		ability := &l.json.unspecifiedAbilities[i]
		ability.Attributes.ID, err = l.getHashID(ability.Attributes)
		if err != nil {
			return nil, err
		}

		ability.Type = database.AbilityTypeUnspecifiedAbility
		abilities = append(abilities, ability.Ability)
	}

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]
		ability.Attributes.ID, err = l.getHashID(ability.Attributes)
		if err != nil {
			return nil, err
		}

		ability.Type = database.AbilityTypeEnemyAbility
		abilities = append(abilities, ability.Ability)
	}

	return dedupeRows(abilities, l.Hashes), nil
}




func (l *Lookup) loop1SeedAbilityAttributes(qtx *database.Queries, ctx context.Context) error {
	attributes := l.extractAbilityAttributes()

	params := database.CreateAbilityAttributesBulkParams{
		DataHash:         make([]string, len(attributes)),
		Rank:             make([]sql.NullInt32, len(attributes)),
		AppearsInHelpBar: make([]bool, len(attributes)),
		CanCopycat:       make([]bool, len(attributes)),
	}

	for i, a := range attributes {
		params.DataHash[i] = generateDataHash(a)
		params.Rank[i] = h.GetNullInt32(a.Rank)
		params.AppearsInHelpBar[i] = a.AppearsInHelpBar
		params.CanCopycat[i] = a.CanCopycat
	}

	dbRows, err := qtx.CreateAbilityAttributesBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create ability attributes: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAbilityAttributes() []Attributes {
	attributes := []Attributes{}

	for _, ability := range l.json.enemyAbilities {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range l.json.items {
		if len(ability.BattleInteractions) > 0 {
			attributes = append(attributes, ability.Attributes)
		}
	}

	for _, ability := range l.json.overdrives {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range l.json.overdriveAbilities {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range l.json.playerAbilities {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range l.json.triggerCommands {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range l.json.unspecifiedAbilities {
		attributes = append(attributes, ability.Attributes)
	}

	return dedupeRows(attributes, l.Hashes)
}

