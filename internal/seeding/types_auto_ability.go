package seeding

import (
	"fmt"

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

func (a AutoAbility) ToKeyFields() []any {
	return []any{
		a.Name,
	}
}

func (a AutoAbility) GetID() int32 {
	return a.ID
}

func (a AutoAbility) Error() string {
	return fmt.Sprintf("auto ability %s", a.Name)
}

func (a AutoAbility) GetResParamsNamed() ResParamsNamed {
	return ResParamsNamed{
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
