package api

type Item struct {
	ID                    	int32					`json:"id"`
	Name                  	string					`json:"name"`
	Description           	string					`json:"description"`
	SphereGridDescription	*string					`json:"sphere_grid_description,omitempty"`
	Effect                	string					`json:"effect"`
	Category              	NamedAPIResource		`json:"category"`
	Usability				string					`json:"usability"`
	BasePrice				*int32					`json:"base_price"`
	SellValue				int32					`json:"sell_value"`
	ItemAbility				*NamedAPIResource		`json:"item_ability"`
	AvailableMenus			[]NamedAPIResource		`json:"available_menus"`
	RelatedStats          	[]NamedAPIResource		`json:"related_stats"`
	Monsters				[]NamedAPIResource		`json:"monsters"`
	Treasures				[]UnnamedAPIResource	`json:"treasures"`
	Shops					[]UnnamedAPIResource	`json:"shops"`
	Sidequests				[]NamedAPIResource		`json:"sidequests"`
	BlitzballPrizes			[]UnnamedAPIResource	`json:"blitzball_prizes"`
	AeonLearnAbilities		[]AbilityAmount			`json:"aeon_learn_abilities"`
	AutoAbilities			[]AutoAbilityAmount		`json:"auto_abilities"`
	Mixes					[]NamedAPIResource		`json:"mixes"`
}


type AbilityAmount struct {
	Ability NamedAPIResource	`json:"ability"`
	Amount	int32				`json:"amount"`
}

func (aa AbilityAmount) IsZero() bool {
	return aa.Ability.Name == ""
}

func (aa AbilityAmount) GetAPIResource() APIResource {
	return aa.Ability
}

func (aa AbilityAmount) GetName() string {
	return aa.Ability.Name
}

func (aa AbilityAmount) GetVersion() *int32 {
	return nil
}

func (aa AbilityAmount) GetVal() int32 {
	return aa.Amount
}

func newAbilityAmount(res NamedAPIResource, amount int32) AbilityAmount {
	return AbilityAmount{
		Ability:   	res,
		Amount: 	amount,
	}
}


type AutoAbilityAmount struct {
	AutoAbility NamedAPIResource	`json:"ability"`
	Amount		int32				`json:"amount"`
}

func (aa AutoAbilityAmount) IsZero() bool {
	return aa.AutoAbility.Name == ""
}

func (aa AutoAbilityAmount) GetAPIResource() APIResource {
	return aa.AutoAbility
}

func (aa AutoAbilityAmount) GetName() string {
	return aa.AutoAbility.Name
}

func (aa AutoAbilityAmount) GetVersion() *int32 {
	return nil
}

func (aa AutoAbilityAmount) GetVal() int32 {
	return aa.Amount
}

func newAutoAbilityAmount(res NamedAPIResource, amount int32) AutoAbilityAmount {
	return AutoAbilityAmount{
		AutoAbility:   	res,
		Amount: 		amount,
	}
}