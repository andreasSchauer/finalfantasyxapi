package sectiontypes

type MonsterSimple struct {
	ID             int32                   `json:"id"`
	URL            string                  `json:"url"`
	Name           string                  `json:"name"`
	Version        *int32                  `json:"version,omitempty"`
	Specification  *string                 `json:"specification,omitempty"`
	HP             string                  `json:"hp"`
	AP             string                  `json:"ap"`
	Gil            int32                   `json:"gil"`
	MaxBribeAmount *int32                  `json:"max_bribe_amount,omitempty"`
	Availability   string                  `json:"availability"`
	IsRepeatable   bool                    `json:"is_repeatable"`
	CanBeCaptured  bool                    `json:"can_be_captured"`
	RonsoRages     []string                `json:"ronso_rages,omitempty"`
	Areas          []string                `json:"areas"`
	Items          *MonsterItemsSimple     `json:"items,omitempty"`
	Equipment      *MonsterEquipmentSimple `json:"equipment,omitempty"`
}

func (m MonsterSimple) GetURL() string {
	return m.URL
}

type MonsterItemsSimple struct {
	StealCommon         *string  `json:"steal_common,omitempty"`
	StealRare           *string  `json:"steal_rare,omitempty"`
	DropCommon          *string  `json:"drop_common,omitempty"`
	DropRare            *string  `json:"drop_rare,omitempty"`
	SecondaryDropCommon *string  `json:"secondary_drop_common,omitempty"`
	SecondaryDropRare   *string  `json:"secondary_drop_rare,omitempty"`
	Bribe               *string  `json:"bribe,omitempty"`
	OtherItems          []string `json:"other_items,omitempty"`
}

type MonsterEquipmentSimple struct {
	WeaponAbilities []string `json:"weapon_abilities"`
	ArmorAbilities  []string `json:"armor_abilities"`
}
