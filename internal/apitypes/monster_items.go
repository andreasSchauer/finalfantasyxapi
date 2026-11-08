package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type MonsterItems struct {
	DropChance          int32                                          `json:"drop_chance"`
	DropCondition       *string                                        `json:"drop_condition,omitempty"`
	StealCommon         *apiresource.ResourceAmount[apiresource.Typed] `json:"steal_common"`
	StealRare           *apiresource.ResourceAmount[apiresource.Typed] `json:"steal_rare"`
	DropCommon          *apiresource.ResourceAmount[apiresource.Typed] `json:"drop_common"`
	DropRare            *apiresource.ResourceAmount[apiresource.Typed] `json:"drop_rare"`
	SecondaryDropCommon *apiresource.ResourceAmount[apiresource.Typed] `json:"secondary_drop_common"`
	SecondaryDropRare   *apiresource.ResourceAmount[apiresource.Typed] `json:"secondary_drop_rare"`
	Bribe               *apiresource.ResourceAmount[apiresource.Typed] `json:"bribe"`
	OtherItemsCondition *string                                        `json:"other_items_condition,omitempty"`
	OtherItems          []PossibleItem                                 `json:"other_items"`
}

func (mi MonsterItems) IsZero() bool {
	return mi.DropChance == 0 && mi.OtherItemsCondition == nil
}
