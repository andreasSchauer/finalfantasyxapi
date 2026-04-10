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
	Monster       		NamedAPIResource  	`json:"monster"`
	StealCommon   		int32 				`json:"steal_common,omitempty"`
	StealRare   		int32 				`json:"steal_rare,omitempty"`
	DropCommon          int32 				`json:"drop_common,omitempty"`
	DropRare          	int32 				`json:"drop_rare,omitempty"`
	SecondaryDropCommon int32 				`json:"secondary_drop_common,omitempty"`
	SecondaryDropRare 	int32 				`json:"secondary_drop_rare,omitempty"`
	Bribe         		int32             	`json:"bribe,omitempty"`
	Other		  		int32				`json:"other,omitempty"`
}


func createItemMonster(cfg *Config, itemName string, mon NamedAPIResource) MonItemAmts {
	monster, _ := seeding.GetResourceByID(mon.ID, cfg.l.MonstersID)

	return MonItemAmts{
		Monster:       			mon,
		StealCommon: 			getAmountMonItem(itemName, monster.Items.StealCommon),
		StealRare: 				getAmountMonItem(itemName, monster.Items.StealRare),
		DropCommon: 			getAmountMonItem(itemName, monster.Items.DropCommon),
		DropRare: 				getAmountMonItem(itemName, monster.Items.DropRare),
		SecondaryDropCommon: 	getAmountMonItem(itemName, monster.Items.SecondaryDropCommon),
		SecondaryDropRare: 		getAmountMonItem(itemName, monster.Items.SecondaryDropRare),
		Bribe:         			getAmountMonItem(itemName, monster.Items.Bribe),
		Other: 		   			getAmountOther(itemName, monster.Items.OtherItems),
	}
}


func getAmountMonItem(wantedItem string, ia *seeding.ItemAmount) int32 {
	if ia == nil || ia.ItemName != wantedItem {
		return 0
	}

	return ia.Amount
}

func getAmountOther(wantedItem string, otherItems []seeding.PossibleItem) int32 {
	for _, otherItem := range otherItems {
		if otherItem.ItemAmount.ItemName == wantedItem {
			return otherItem.ItemAmount.Amount
		}
	}

	return 0
}