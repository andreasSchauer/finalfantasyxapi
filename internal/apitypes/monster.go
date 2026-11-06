package apitypes

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Monster struct {
	ID                   int32                            `json:"id"`
	Name                 string                           `json:"name"`
	Version              *int32                           `json:"version,omitempty"`
	Specification        *string                          `json:"specification,omitempty"`
	AppliedState         *AppliedState                    `json:"applied_state,omitempty"`
	AgilityParameters    *AgilityParams                   `json:"agility_parameters"`
	Notes                *string                          `json:"notes,omitempty"`
	Species              string                           `json:"species"`
	Availability         string                           `json:"availability"`
	IsRepeatable         bool                             `json:"is_repeatable"`
	CanBeCaptured        bool                             `json:"can_be_captured"`
	AreaConquestLocation *string                          `json:"area_conquest_location,omitempty"`
	Category             string                           `json:"category"`
	CTBIconType          string                           `json:"ctb_icon_type"`
	HasOverdrive         bool                             `json:"has_overdrive"`
	IsUnderwater         bool                             `json:"is_underwater"`
	IsZombie             bool                             `json:"is_zombie"`
	Distance             int32                            `json:"distance"`
	Properties           []apiresource.NamedAPIResource   `json:"properties"`
	AutoAbilities        []apiresource.NamedAPIResource   `json:"auto_abilities"`
	AP                   int32                            `json:"ap"`
	APOverkill           int32                            `json:"ap_overkill"`
	OverkillDamage       int32                            `json:"overkill_damage"`
	Gil                  int32                            `json:"gil"`
	StealGil             *int32                           `json:"steal_gil"`
	RonsoRages           []apiresource.NamedAPIResource   `json:"ronso_rages"`
	DoomCountdown        *int32                           `json:"doom_countdown"`
	PoisonRate           *float32                         `json:"poison_rate"`
	PoisonDamage         *int32                           `json:"poison_damage,omitempty"`
	ThreatenChance       *int32                           `json:"threaten_chance"`
	ZanmatoLevel         int32                            `json:"zanmato_level"`
	MonsterArenaPrice    *int32                           `json:"monster_arena_price,omitempty"`
	SensorText           *string                          `json:"sensor_text"`
	ScanText             *string                          `json:"scan_text"`
	Areas                []apiresource.AreaAPIResource    `json:"areas"`
	Formations           []apiresource.UnnamedAPIResource `json:"monster_formations"`
	BaseStats            []BaseStat                       `json:"base_stats"`
	RegenAmounts         []RegenAmount                    `json:"regen_amounts"`
	Items                *MonsterItems                    `json:"items"`
	BribeChances         []BribeChance                    `json:"bribe_chances,omitempty"`
	Equipment            *MonsterEquipment                `json:"equipment"`
	ElemResists          []ElementalResist                `json:"elem_resists"`
	StatusImmunities     []apiresource.NamedAPIResource   `json:"status_immunities"`
	StatusResists        []StatusResist                   `json:"status_resists"`
	Abilities            []MonsterAbility                 `json:"abilities"`
	AlteredStates        []AlteredState                   `json:"altered_states"`
}

func (m Monster) String() string {
	return fmt.Sprintf("monster '%s'", h.NameToString(m.Name, m.Version, m.Specification))
}

type MonsterAbility struct {
	Ability  apiresource.TypedAPIResource `json:"ability"`
	IsForced bool                         `json:"is_forced"`
	IsUnused bool                         `json:"is_unused"`
}

func (ma MonsterAbility) GetAPIResource() apiresource.APIResource {
	return ma.Ability
}

type BribeChance struct {
	Gil    int32 `json:"gil"`
	Chance int32 `json:"chance"`
}

type AppliedState struct {
	Condition     string           `json:"condition"`
	IsTemporary   bool             `json:"is_temporary"`
	AppliedStatus *InflictedStatus `json:"applied_status,omitempty"`
}

func (as AppliedState) IsZero() bool {
	return as.Condition == ""
}