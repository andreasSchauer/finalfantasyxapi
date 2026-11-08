package servicetypes

import h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"

type DropChanceResponse struct {
	DropChanceFinBlow     float64               `json:"drop_chance_fin_blow"`
	DropChanceNoFinBlow   float64               `json:"drop_chance_no_fin_blow"`
	WantedEquipment       WantedEquipment       `json:"wanted_equipment"`
	EquipmentMatchChances EquipmentMatchChances `json:"equipment_match_chances"`
	CharacterChances      CharacterChances      `json:"character_chances"`
	MonsterChances        MonsterDropChances    `json:"monster_chances"`
}

func (r DropChanceResponse) Percent() DropChanceResponse {
	return DropChanceResponse{
		DropChanceFinBlow:     h.DecimalToPercent(r.DropChanceFinBlow),
		DropChanceNoFinBlow:   h.DecimalToPercent(r.DropChanceNoFinBlow),
		WantedEquipment:       r.WantedEquipment,
		EquipmentMatchChances: r.EquipmentMatchChances.Percent(),
		CharacterChances:      r.CharacterChances.Percent(),
		MonsterChances:        r.MonsterChances.Percent(),
	}
}

func (r DropChanceResponse) Round(n int32) DropChanceResponse {
	return DropChanceResponse{
		DropChanceFinBlow:     h.DecimalToPercent(r.DropChanceFinBlow),
		DropChanceNoFinBlow:   h.DecimalToPercent(r.DropChanceNoFinBlow),
		WantedEquipment:       r.WantedEquipment,
		EquipmentMatchChances: r.EquipmentMatchChances.Round(n),
		CharacterChances:      r.CharacterChances.Round(n),
		MonsterChances:        r.MonsterChances.Round(n),
	}
}


type WantedEquipment struct {
	MinSlots         int32    `json:"min_slots"`
	MaxSlots         int32    `json:"max_slots"`
	ReqAbilities     []string `json:"req_abilities"`
	ReqEmptySlots    *int32   `json:"req_empty_slots"`
	MinOptionalSlots int32    `json:"min_optional_slots"`
	MaxOptionalSlots int32    `json:"max_optional_slots"`
	OptionalSlotsAre string   `json:"optional_slots_are"`
}

type OptionalSlotDefinition string

const (
	OptionalSlotAllEmpty       OptionalSlotDefinition = "all-empty"
	OptionalSlotAbilityOrEmpty OptionalSlotDefinition = "ability-or-empty"
)

type EquipmentMatchChances struct {
	FinBlow   float64 `json:"fin_blow"`
	NoFinBlow float64 `json:"no_fin_blow"`
}

func (c EquipmentMatchChances) Percent() EquipmentMatchChances {
	c.FinBlow = h.DecimalToPercent(c.FinBlow)
	c.NoFinBlow = h.DecimalToPercent(c.NoFinBlow)

	return c
}

func (c EquipmentMatchChances) Round(n int32) EquipmentMatchChances {
	c.FinBlow = h.FloatRound(c.FinBlow, n)
	c.NoFinBlow = h.FloatRound(c.NoFinBlow, n)

	return c
}


type CharacterChances struct {
	EligibleChars []string             `json:"eligible_chars"`
	AnyFinBlow    float64              `json:"any_fin_blow"`
	AnyNoFinBlow  float64              `json:"any_no_fin_blow"`
	SpecificChar  *SpecificCharChances `json:"specific_character,omitempty"`
}

func (c CharacterChances) Percent() CharacterChances {
	c.AnyFinBlow = h.DecimalToPercent(c.AnyFinBlow)
	c.AnyNoFinBlow = h.DecimalToPercent(c.AnyNoFinBlow)

	if c.SpecificChar != nil {
		newObj := c.SpecificChar.Percent()
		c.SpecificChar = &newObj
	}

	return c
}

func (c CharacterChances) Round(n int32) CharacterChances {
	c.AnyFinBlow = h.FloatRound(c.AnyFinBlow, n)
	c.AnyNoFinBlow = h.FloatRound(c.AnyNoFinBlow, n)

	if c.SpecificChar != nil {
		newObj := c.SpecificChar.Round(n)
		c.SpecificChar = &newObj
	}

	return c
}

type SpecificCharChances struct {
	IsEligible bool    `json:"is_eligible"`
	FinBlow    float64 `json:"fin_blow"`
	NoFinBlow  float64 `json:"no_fin_blow"`
}

func (c SpecificCharChances) Percent() SpecificCharChances {
	c.FinBlow = h.DecimalToPercent(c.FinBlow)
	c.NoFinBlow = h.DecimalToPercent(c.NoFinBlow)

	return c
}

func (c SpecificCharChances) Round(n int32) SpecificCharChances {
	c.FinBlow = h.FloatRound(c.FinBlow, n)
	c.NoFinBlow = h.FloatRound(c.NoFinBlow, n)

	return c
}

type MonsterDropChances struct {
	MonsterID           int32                `json:"monster_id"`
	MonsterName         string               `json:"monster_name"`
	EquipmentDropChance float64              `json:"equipment_drop_chance"`
	SlotsChances        []ChanceAmount       `json:"slots_chances"`
	ShotsChances        []ChanceAmount       `json:"shots_chances"`
	AutoAbilities       DroppedAutoAbilities `json:"auto_abilities"`
}

func (c MonsterDropChances) Percent() MonsterDropChances {
	c.EquipmentDropChance = h.DecimalToPercent(c.EquipmentDropChance)
	c.SlotsChances = h.SlicePercent(c.SlotsChances)
	c.ShotsChances = h.SlicePercent(c.ShotsChances)
	c.AutoAbilities = c.AutoAbilities.Percent()

	return c
}

func (c MonsterDropChances) Round(n int32) MonsterDropChances {
	c.EquipmentDropChance = h.FloatRound(c.EquipmentDropChance, n)
	c.SlotsChances = h.SliceRound(c.SlotsChances, n)
	c.ShotsChances = h.SliceRound(c.ShotsChances, n)
	c.AutoAbilities = c.AutoAbilities.Round(n)

	return c
}


type ChanceAmount struct {
	Amount int32   `json:"amount"`
	Chance float64 `json:"chance"`
}

func (c ChanceAmount) Percent() ChanceAmount {
	c.Chance = h.DecimalToPercent(c.Chance)

	return c
}

func (c ChanceAmount) Round(n int32) ChanceAmount {
	c.Chance = h.FloatRound(c.Chance, n)

	return c
}


type DroppedAutoAbilities struct {
	Forced  *EqAbilitySimple  `json:"forced"`
	Dropped []EqAbilityChance `json:"dropped"`
}

func (d DroppedAutoAbilities) Percent() DroppedAutoAbilities {
	d.Dropped = h.SlicePercent(d.Dropped)

	return d
}

func (d DroppedAutoAbilities) Round(n int32) DroppedAutoAbilities {
	d.Dropped = h.SliceRound(d.Dropped, n)

	return d
}


type EqAbilitySimple struct {
	ID         int32    `json:"id"`
	Name       string   `json:"name"`
	Characters []string `json:"characters"`
}

type EqAbilityChance struct {
	EqAbilitySimple
	Chance float64 `json:"chance"`
}

func (a EqAbilityChance) Percent() EqAbilityChance {
	a.Chance = h.DecimalToPercent(a.Chance)

	return a
}

func (a EqAbilityChance) Round(n int32) EqAbilityChance {
	a.Chance = h.FloatRound(a.Chance, n)

	return a
}