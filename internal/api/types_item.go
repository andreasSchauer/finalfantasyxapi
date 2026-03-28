package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"

type Item struct {
	ID                 int32                              	`json:"id"`
	Name               string                             	`json:"name"`
	Description        string                             	`json:"description"`
	SgDescription      *string                            	`json:"sg_description,omitempty"`
	Effect             string                             	`json:"effect"`
	Category           NamedAPIResource                   	`json:"category"`
	Usability          string                             	`json:"usability"`
	BasePrice          *int32                             	`json:"base_price"`
	SellValue          int32                              	`json:"sell_value"`
	ItemAbility        *NamedAPIResource                  	`json:"item_ability"`
	AvailableMenus     []NamedAPIResource                 	`json:"available_menus"`
	RelatedStats       []NamedAPIResource                 	`json:"related_stats"`
	Monsters           []MonItemAmts                      	`json:"monsters"`             // GetMonsterIDsByItem
	Treasures          []ResourceAmount[UnnamedAPIResource] `json:"treasures"`            // GetTreasureIDsByItem
	Shops              []UnnamedAPIResource               	`json:"shops"`                // GetItemShopIDs
	Quests             []ResourceAmount[QuestAPIResource]	`json:"quests"`               // GetItemQuestIDs
	BlitzballPrizes    []ResourceAmount[NamedAPIResource]	`json:"blitzball_prizes"`     // GetItemBlitzballPrizeIDs
	AeonLearnAbilities []ResourceAmount[NamedAPIResource] 	`json:"aeon_learn_abilities"` // GetItemPlayerAbilityIDs
	AutoAbilities      []ResourceAmount[NamedAPIResource] 	`json:"auto_abilities"`       // GetItemAutoAbilityIDs
	Mixes              []NamedAPIResource                 	`json:"mixes"`                // GetItemMixIDs
}

// need to account for otherItems
type MonItemAmts struct {
	Monster       NamedAPIResource  `json:"monster"`
	Steal         *CommonRareAmount `json:"steal,omitempty"`
	Drop          *CommonRareAmount `json:"drop,omitempty"`
	SecondaryDrop *CommonRareAmount `json:"secondary_drop,omitempty"`
	Bribe         int32             `json:"bribe,omitempty"`
}

type CommonRareAmount struct {
	Common int32 `json:"common,omitempty"`
	Rare   int32 `json:"rare,omitempty"`
}

func (cra CommonRareAmount) IsZero() bool {
	return cra.Common == 0 && cra.Rare == 0
}

func createItemMonster(cfg *Config, item seeding.Item, mon NamedAPIResource) MonItemAmts {
	monster, _ := seeding.GetResourceByID(mon.ID, cfg.l.MonstersID)

	return MonItemAmts{
		Monster:       mon,
		Steal:         craPtr(item, monster.Items.StealCommon, monster.Items.StealRare),
		Drop:          craPtr(item, monster.Items.DropCommon, monster.Items.DropRare),
		SecondaryDrop: craPtr(item, monster.Items.SecondaryDropCommon, monster.Items.SecondaryDropRare),
		Bribe:         getAmountMonItem(item, monster.Items.Bribe),
	}
}

func craPtr(item seeding.Item, common, rare *seeding.ItemAmount) *CommonRareAmount {
	cra := CommonRareAmount{
		Common: getAmountMonItem(item, common),
		Rare:   getAmountMonItem(item, rare),
	}

	if cra.IsZero() {
		return nil
	}

	return &cra
}

func getAmountMonItem(item seeding.Item, ia *seeding.ItemAmount) int32 {
	if ia == nil || ia.ItemName != item.Name {
		return 0
	}

	return ia.Amount
}
