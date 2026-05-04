package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type AutoAbility struct {
	ID                  int32
	GradRecoveryStatID  *int32
	OnHitElementID      *int32
	AddedPropertyID     *int32
	CnvrsnFromModID     *int32
	CnvrsnToModID       *int32
	Name                string           `json:"name"`
	Description         *string          `json:"description"`
	Effect              string           `json:"effect"`
	Type                string           `json:"type"`
	Category            string           `json:"category"`
	RelatedStats        []string         `json:"related_stats"`
	AbilityValue        *int32           `json:"ability_value"`
	RequiredItem        *ItemAmount      `json:"required_item"`
	LockedOutAbilities  []string         `json:"locked_out_abilities"`
	ActivationCondition string           `json:"activation_condition"`
	Counter             *string          `json:"counter"`
	GradualRecovery     *string          `json:"gradual_recovery"`
	AutoItemUse         []string         `json:"auto_item_use"`
	OnHitElement        *string          `json:"on_hit_element"`
	AddedElemResist     *ElementalResist `json:"added_elem_resist"`
	OnHitStatus         *InflictedStatus `json:"on_hit_status"`
	AddedStatusResists  []StatusResist   `json:"added_status_resists"`
	AddedStatusses      []string         `json:"added_statusses"`
	AddedProperty       *string          `json:"added_property"`
	ConversionFrom      *string          `json:"conversion_from"`
	ConversionTo        *string          `json:"conversion_to"`
	StatChanges         []StatChange     `json:"stat_changes"`
	ModifierChanges     []ModifierChange `json:"modifier_changes"`
}

func (a AutoAbility) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.Name,
		h.DerefOrNil(a.Description),
		a.Effect,
		a.Type,
		a.Category,
		h.DerefOrNil(a.AbilityValue),
		h.ObjPtrToID(a.RequiredItem),
		a.ActivationCondition,
		h.DerefOrNil(a.Counter),
		h.DerefOrNil(a.GradRecoveryStatID),
		h.DerefOrNil(a.OnHitElementID),
		h.ObjPtrToID(a.AddedElemResist),
		h.ObjPtrToID(a.OnHitStatus),
		h.DerefOrNil(a.AddedPropertyID),
		h.DerefOrNil(a.CnvrsnFromModID),
		h.DerefOrNil(a.CnvrsnToModID),
	}
}

func (a AutoAbility) GetID() int32 {
	return a.ID
}

func (a AutoAbility) Error() string {
	return fmt.Sprintf("auto ability %s", a.Name)
}

func (a AutoAbility) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   a.ID,
		Name: a.Name,
	}
}

func (a AutoAbility) GetItemAmount() ItemAmount {
	itemAmtPtr := a.RequiredItem

	if itemAmtPtr == nil {
		return ItemAmount{}
	}

	return *itemAmtPtr
}

func (l *Lookup) loop5SeedAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	abilities, err := l.extractAutoAbilities()
	if err != nil {
		return err
	}

	params := database.CreateAutoAbilityBulkParams{
		DataHash:             make([]string, len(abilities)),
		Name:                 make([]string, len(abilities)),
		Description:          make([]sql.NullString, len(abilities)),
		Effect:               make([]string, len(abilities)),
		Type:                 make([]database.EquipType, len(abilities)),
		Category:             make([]database.AutoAbilityCategory, len(abilities)),
		AbilityValue:         make([]sql.NullInt32, len(abilities)),
		ActivationCondition:  make([]database.AaActivationCondition, len(abilities)),
		Counter:              make([]database.NullCounterType, len(abilities)),
		RequiredItemAmountID: make([]sql.NullInt32, len(abilities)),
		GradRcvryStatID:      make([]sql.NullInt32, len(abilities)),
		OnHitElementID:       make([]sql.NullInt32, len(abilities)),
		AddedElemResistID:    make([]sql.NullInt32, len(abilities)),
		OnHitStatusID:        make([]sql.NullInt32, len(abilities)),
		AddedPropertyID:      make([]sql.NullInt32, len(abilities)),
		CnvrsnFromModID:      make([]sql.NullInt32, len(abilities)),
		CnvrsnToModID:        make([]sql.NullInt32, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = generateDataHash(a)
		params.Name[i] = a.Name
		params.Description[i] = h.GetNullString(a.Description)
		params.Effect[i] = a.Effect
		params.Type[i] = database.EquipType(a.Type)
		params.Category[i] = database.AutoAbilityCategory(a.Category)
		params.AbilityValue[i] = h.GetNullInt32(a.AbilityValue)
		params.ActivationCondition[i] = database.AaActivationCondition(a.ActivationCondition)
		params.Counter[i] = database.ToNullCounterType(a.Counter)
		params.RequiredItemAmountID[i] = h.ObjPtrToNullInt32ID(a.RequiredItem)
		params.GradRcvryStatID[i] = h.GetNullInt32(a.GradRecoveryStatID)
		params.OnHitElementID[i] = h.GetNullInt32(a.OnHitElementID)
		params.AddedElemResistID[i] = h.ObjPtrToNullInt32ID(a.AddedElemResist)
		params.OnHitStatusID[i] = h.ObjPtrToNullInt32ID(a.OnHitStatus)
		params.AddedPropertyID[i] = h.GetNullInt32(a.AddedPropertyID)
		params.CnvrsnFromModID[i] = h.GetNullInt32(a.CnvrsnFromModID)
		params.CnvrsnToModID[i] = h.GetNullInt32(a.CnvrsnToModID)
	}

	dbRows, err := qtx.CreateAutoAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create auto-abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		l.json.autoAbilities[i].ID = row.ID
		l.AutoAbilities[abilities[i].Name] = abilities[i]
		l.AutoAbilitiesID[row.ID] = abilities[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractAutoAbilities() ([]AutoAbility, error) {
	abilities := []AutoAbility{}
	var err error

	for i := range l.json.autoAbilities {
		ability := &l.json.autoAbilities[i]

		ability.GradRecoveryStatID, err = assignFKPtr(ability.GradualRecovery, l.Stats)
		if err != nil {
			return nil, err
		}

		ability.OnHitElementID, err = assignFKPtr(ability.OnHitElement, l.Elements)
		if err != nil {
			return nil, err
		}

		ability.AddedPropertyID, err = assignFKPtr(ability.AddedProperty, l.Properties)
		if err != nil {
			return nil, err
		}

		ability.CnvrsnFromModID, err = assignFKPtr(ability.ConversionFrom, l.Modifiers)
		if err != nil {
			return nil, err
		}

		ability.CnvrsnToModID, err = assignFKPtr(ability.ConversionTo, l.Modifiers)
		if err != nil {
			return nil, err
		}

		if ability.RequiredItem != nil {
			ability.RequiredItem.ID, err = l.getHashID(ability.RequiredItem)
			if err != nil {
				return nil, err
			}
		}

		if ability.AddedElemResist != nil {
			ability.AddedElemResist.ID, err = l.getHashID(ability.AddedElemResist)
			if err != nil {
				return nil, err
			}
		}

		if ability.OnHitStatus != nil {
			ability.OnHitStatus.ID, err = l.getHashID(ability.OnHitStatus)
			if err != nil {
				return nil, err
			}
		}

		abilities = append(abilities, *ability)
	}

	return dedupeRows(abilities, l.Hashes), nil
}

func (l *Lookup) completeAutoAbilities() error {
	for i := range l.json.autoAbilities {
		autoAbility := &l.json.autoAbilities[i]

		err := assignIDs(l, autoAbility.AddedStatusResists)
		if err != nil {
			return err
		}

		err = assignIDs(l, autoAbility.StatChanges)
		if err != nil {
			return err
		}

		err = assignIDs(l, autoAbility.ModifierChanges)
		if err != nil {
			return err
		}

		l.AutoAbilities[autoAbility.Name] = *autoAbility
		l.AutoAbilitiesID[autoAbility.ID] = *autoAbility
	}

	return nil
}

func (l *Lookup) getAutoAbilityAddedStatusResists(a AutoAbility) ([]StatusResist, error) {
	return a.AddedStatusResists, nil
}

func (l *Lookup) getAutoAbilityAddedStatusses(a AutoAbility) ([]StatusCondition, error) {
	return getResources(a.AddedStatusses, l.StatusConditions)
}

func (l *Lookup) getAutoAbilityAutoItems(a AutoAbility) ([]Item, error) {
	return getResources(a.AutoItemUse, l.Items)
}

func (l *Lookup) getAutoAbilityLockedOutAbilities(a AutoAbility) ([]AutoAbility, error) {
	return getResources(a.LockedOutAbilities, l.AutoAbilities)
}

func (l *Lookup) getAutoAbilityModifierChanges(a AutoAbility) ([]ModifierChange, error) {
	return a.ModifierChanges, nil
}

func (l *Lookup) getAutoAbilityRelatedStats(a AutoAbility) ([]Stat, error) {
	return getResources(a.RelatedStats, l.Stats)
}

func (l *Lookup) getAutoAbilityStatChanges(a AutoAbility) ([]StatChange, error) {
	return a.StatChanges, nil
}

func (l *Lookup) seedJuncAutoAbilitiesAddedStatusResists(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + added status resists"
	jParams, err := processJunctions(l, desc, l.json.autoAbilities, l.getAutoAbilityAddedStatusResists)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesAddedStatusResistsJunctionBulk(ctx, database.CreateAutoAbilitiesAddedStatusResistsJunctionBulkParams{
		DataHash:       jParams.DataHashes,
		AutoAbilityID:  jParams.ParentIDs,
		StatusResistID: jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncAutoAbilitiesAddedStatusses(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + added statusses"
	jParams, err := processJunctions(l, desc, l.json.autoAbilities, l.getAutoAbilityAddedStatusses)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesAddedStatussesJunctionBulk(ctx, database.CreateAutoAbilitiesAddedStatussesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		AutoAbilityID:     jParams.ParentIDs,
		StatusConditionID: jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncAutoAbilitiesAutoItems(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + auto items"
	jParams, err := processJunctions(l, desc, l.json.autoAbilities, l.getAutoAbilityAutoItems)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesAutoItemJunctionBulk(ctx, database.CreateAutoAbilitiesAutoItemJunctionBulkParams{
		DataHash:      jParams.DataHashes,
		AutoAbilityID: jParams.ParentIDs,
		ItemID:        jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncAutoAbilitiesLockedOutAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + locked out"
	jParams, err := processJunctions(l, desc, l.json.autoAbilities, l.getAutoAbilityLockedOutAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesLockedOutJunctionBulk(ctx, database.CreateAutoAbilitiesLockedOutJunctionBulkParams{
		DataHash:        jParams.DataHashes,
		ParentAbilityID: jParams.ParentIDs,
		ChildAbilityID:  jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncAutoAbilitiesModifierChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + modifier changes"
	jParams, err := processJunctions(l, desc, l.json.autoAbilities, l.getAutoAbilityModifierChanges)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesModifierChangesJunctionBulk(ctx, database.CreateAutoAbilitiesModifierChangesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		AutoAbilityID:    jParams.ParentIDs,
		ModifierChangeID: jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncAutoAbilitiesRelatedStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + related stats"
	jParams, err := processJunctions(l, desc, l.json.autoAbilities, l.getAutoAbilityRelatedStats)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesRelatedStatsJunctionBulk(ctx, database.CreateAutoAbilitiesRelatedStatsJunctionBulkParams{
		DataHash:      jParams.DataHashes,
		AutoAbilityID: jParams.ParentIDs,
		StatID:        jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncAutoAbilitiesStatChanges(qtx *database.Queries, ctx context.Context) error {
	const desc string = "auto-abilities + stat changes"
	jParams, err := processJunctions(l, desc, l.json.autoAbilities, l.getAutoAbilityStatChanges)
	if err != nil {
		return err
	}

	return qtx.CreateAutoAbilitiesStatChangesJunctionBulk(ctx, database.CreateAutoAbilitiesStatChangesJunctionBulkParams{
		DataHash:      jParams.DataHashes,
		AutoAbilityID: jParams.ParentIDs,
		StatChangeID:  jParams.ChildIDs,
	})
}
