package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type Item struct {
	ID                 int32                                                        `json:"id"`
	Name               string                                                       `json:"name"`
	UntypedItem        apiresource.TypedAPIResource                                 `json:"untyped_item"`
	Category           string                                                       `json:"category"`
	Description        string                                                       `json:"description"`
	SgDescription      *string                                                      `json:"sg_description,omitempty"`
	Effect             string                                                       `json:"effect"`
	Usability          string                                                       `json:"usability"`
	BasePrice          *int32                                                       `json:"base_price"`
	SellValue          int32                                                        `json:"sell_value"`
	ItemAbility        *apiresource.NamedAPIResource                                `json:"item_ability,omitempty"`
	Sphere             *apiresource.NamedAPIResource                                `json:"sphere,omitempty"`
	AvailableMenus     []apiresource.NamedAPIResource                               `json:"available_menus"`
	RelatedStats       []apiresource.NamedAPIResource                               `json:"related_stats"`
	Monsters           []MonItemAmts                                                `json:"monsters"`
	Treasures          []apiresource.ResourceAmount[apiresource.UnnamedAPIResource] `json:"treasures"`
	Shops              []apiresource.UnnamedAPIResource                             `json:"shops"`
	Quests             []apiresource.ResourceAmount[apiresource.QuestAPIResource]   `json:"quests"`
	BlitzballPrizes    []apiresource.ResourceAmount[apiresource.NamedAPIResource]   `json:"blitzball_prizes"`
	AeonLearnAbilities []apiresource.ResourceAmount[apiresource.NamedAPIResource]   `json:"aeon_learn_abilities"`
	AutoAbilities      []apiresource.ResourceAmount[apiresource.NamedAPIResource]   `json:"auto_abilities"`
	Mixes              []apiresource.NamedAPIResource                               `json:"mixes"`
}

type MonItemAmts struct {
	Monster             apiresource.NamedAPIResource `json:"monster"`
	StealCommon         int32                        `json:"steal_common,omitempty"`
	StealRare           int32                        `json:"steal_rare,omitempty"`
	DropCommon          int32                        `json:"drop_common,omitempty"`
	DropRare            int32                        `json:"drop_rare,omitempty"`
	SecondaryDropCommon int32                        `json:"secondary_drop_common,omitempty"`
	SecondaryDropRare   int32                        `json:"secondary_drop_rare,omitempty"`
	Bribe               int32                        `json:"bribe,omitempty"`
	Other               int32                        `json:"other,omitempty"`
}