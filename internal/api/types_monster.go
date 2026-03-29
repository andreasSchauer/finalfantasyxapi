package api

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type Monster struct {
	ID                   int32                `json:"id"`
	Name                 string               `json:"name"`
	Version              *int32               `json:"version,omitempty"`
	Specification        *string              `json:"specification,omitempty"`
	AppliedState         *AppliedState        `json:"applied_state,omitempty"`
	AgilityParameters    *AgilityParams       `json:"agility_parameters"`
	Notes                *string              `json:"notes,omitempty"`
	Species              NamedAPIResource     `json:"species"`
	Availability	     NamedAPIResource	  `json:"availability"`
	IsRepeatable         bool                 `json:"is_repeatable"`
	CanBeCaptured        bool                 `json:"can_be_captured"`
	AreaConquestLocation *string              `json:"area_conquest_location,omitempty"`
	Category             NamedAPIResource     `json:"category"`
	CTBIconType          string               `json:"ctb_icon_type"`
	HasOverdrive         bool                 `json:"has_overdrive"`
	IsUnderwater         bool                 `json:"is_underwater"`
	IsZombie             bool                 `json:"is_zombie"`
	Distance             int32                `json:"distance"`
	Properties           []NamedAPIResource   `json:"properties"`
	AutoAbilities        []NamedAPIResource   `json:"auto_abilities"`
	AP                   int32                `json:"ap"`
	APOverkill           int32                `json:"ap_overkill"`
	OverkillDamage       int32                `json:"overkill_damage"`
	Gil                  int32                `json:"gil"`
	StealGil             *int32               `json:"steal_gil"`
	RonsoRages           []NamedAPIResource   `json:"ronso_rages"`
	DoomCountdown        *int32               `json:"doom_countdown"`
	PoisonRate           *float32             `json:"poison_rate"`
	PoisonDamage         *int32               `json:"poison_damage,omitempty"`
	ThreatenChance       *int32               `json:"threaten_chance"`
	ZanmatoLevel         int32                `json:"zanmato_level"`
	MonsterArenaPrice    *int32               `json:"monster_arena_price,omitempty"`
	SensorText           *string              `json:"sensor_text"`
	ScanText             *string              `json:"scan_text"`
	Areas                []AreaAPIResource    `json:"areas"`
	Formations           []UnnamedAPIResource `json:"monster_formations"`
	BaseStats            []BaseStat           `json:"base_stats"`
	Items                *MonsterItems        `json:"items"`
	BribeChances         []BribeChance        `json:"bribe_chances,omitempty"`
	Equipment            *MonsterEquipment    `json:"equipment"`
	ElemResists          []ElementalResist    `json:"elem_resists"`
	StatusImmunities     []NamedAPIResource   `json:"status_immunities"`
	StatusResists        []StatusResist       `json:"status_resists"`
	Abilities            []MonsterAbility     `json:"abilities"`
	AlteredStates        []AlteredState       `json:"altered_states"`
}

func (m Monster) Error() string {
	return fmt.Sprintf("monster '%s'", h.NameToString(m.Name, m.Version, m.Specification))
}

type MonsterAbility struct {
	Ability  TypedAPIResource `json:"ability"`
	IsForced bool             `json:"is_forced"`
	IsUnused bool             `json:"is_unused"`
}

func convertMonsterAbility(cfg *Config, ability seeding.MonsterAbility) MonsterAbility {
	return MonsterAbility{
		Ability:  keyToTypedAPIResource(cfg, cfg.e.abilities, ability.AbilityReference),
		IsForced: ability.IsForced,
		IsUnused: ability.IsUnused,
	}
}

func (ma MonsterAbility) GetAPIResource() APIResource {
	return ma.Ability
}

type BribeChance struct {
	Gil    int32 `json:"gil"`
	Chance int32 `json:"chance"`
}
