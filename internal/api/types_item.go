package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"

type Item struct {
	ID                 int32                              	`json:"id"`
	Name               string                             	`json:"name"`
	UntypedItem		   TypedAPIResource						`json:"untyped_item"`
	Category           NamedAPIResource                   	`json:"category"`
	Description        string                             	`json:"description"`
	SgDescription      *string                            	`json:"sg_description,omitempty"`
	Effect             string                             	`json:"effect"`
	Usability          string                             	`json:"usability"`
	BasePrice          *int32                             	`json:"base_price"`
	SellValue          int32                              	`json:"sell_value"`
	ItemAbility        *NamedAPIResource                  	`json:"item_ability,omitempty"`
	Sphere			   *NamedAPIResource					`json:"sphere,omitempty"`
	AvailableMenus     []NamedAPIResource                 	`json:"available_menus"`
	RelatedStats       []NamedAPIResource                 	`json:"related_stats"`
	Monsters           []MonItemAmts                      	`json:"monsters"`
	Treasures          []ResourceAmount[UnnamedAPIResource] `json:"treasures"`
	Shops              []UnnamedAPIResource               	`json:"shops"`
	Quests             []ResourceAmount[QuestAPIResource]	`json:"quests"`
	BlitzballPrizes    []ResourceAmount[NamedAPIResource]	`json:"blitzball_prizes"`
	AeonLearnAbilities []ResourceAmount[NamedAPIResource] 	`json:"aeon_learn_abilities"`
	AutoAbilities      []ResourceAmount[NamedAPIResource] 	`json:"auto_abilities"`
	Mixes              []NamedAPIResource                 	`json:"mixes"`
}


type MonItemAmts struct {
	Monster       NamedAPIResource  `json:"monster"`
	Steal         *CommonRareAmount `json:"steal,omitempty"`
	Drop          *CommonRareAmount `json:"drop,omitempty"`
	SecondaryDrop *CommonRareAmount `json:"secondary_drop,omitempty"`
	Bribe         int32             `json:"bribe,omitempty"`
	Other		  int32				`json:"other,omitempty"`
}

type CommonRareAmount struct {
	Common int32 `json:"common,omitempty"`
	Rare   int32 `json:"rare,omitempty"`
}

func (cra CommonRareAmount) IsZero() bool {
	return cra.Common == 0 && cra.Rare == 0
}

func createItemMonster(cfg *Config, itemName string, mon NamedAPIResource) MonItemAmts {
	monster, _ := seeding.GetResourceByID(mon.ID, cfg.l.MonstersID)

	return MonItemAmts{
		Monster:       mon,
		Steal:         craPtr(itemName, monster.Items.StealCommon, monster.Items.StealRare),
		Drop:          craPtr(itemName, monster.Items.DropCommon, monster.Items.DropRare),
		SecondaryDrop: craPtr(itemName, monster.Items.SecondaryDropCommon, monster.Items.SecondaryDropRare),
		Bribe:         getAmountMonItem(itemName, monster.Items.Bribe),
		Other: 		   getAmountOther(itemName, monster.Items.OtherItems),
	}
}

func craPtr(itemName string, common, rare *seeding.ItemAmount) *CommonRareAmount {
	cra := CommonRareAmount{
		Common: getAmountMonItem(itemName, common),
		Rare:   getAmountMonItem(itemName, rare),
	}

	if cra.IsZero() {
		return nil
	}

	return &cra
}

func getAmountMonItem(itemName string, ia *seeding.ItemAmount) int32 {
	if ia == nil || ia.ItemName != itemName {
		return 0
	}

	return ia.Amount
}

func getAmountOther(itemName string, otherItems []seeding.PossibleItem) int32 {
	for _, otherItem := range otherItems {
		if otherItem.ItemAmount.ItemName == itemName {
			return otherItem.ItemAmount.Amount
		}
	}

	return 0
}