package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Monster struct {
	ID                   int32
	Name                 string            `json:"name"`
	Version              *int32            `json:"version"`
	Specification        *string           `json:"specification"`
	Notes                *string           `json:"notes"`
	Species              string            `json:"species"`
	Availability         string            `json:"availability"`
	IsRepeatable         bool              `json:"is_repeatable"`
	CanBeCaptured        bool              `json:"can_be_captured"`
	AreaConquestLocation *string           `json:"area_conquest_location"`
	Category             string            `json:"category"`
	CTBIconType          string            `json:"ctb_icon_type"`
	HasOverdrive         bool              `json:"has_overdrive"`
	IsUnderwater         bool              `json:"is_underwater"`
	IsZombie             bool              `json:"is_zombie"`
	Distance             int32             `json:"distance"`
	Properties           []string          `json:"properties"`
	AutoAbilities        []string          `json:"auto_abilities"`
	AP                   int32             `json:"ap"`
	APOverkill           int32             `json:"ap_overkill"`
	OverkillDamage       int32             `json:"overkill_damage"`
	Gil                  int32             `json:"gil"`
	StealGil             *int32            `json:"steal_gil"`
	RonsoRages           []string          `json:"ronso_rages"`
	DoomCountdown        *int32            `json:"doom_countdown"`
	PoisonRate           *float32          `json:"poison_rate"`
	ThreatenChance       *int32            `json:"threaten_chance"`
	ZanmatoLevel         int32             `json:"zanmato_level"`
	MonsterArenaPrice    *int32            `json:"monster_arena_price"`
	SensorText           *string           `json:"sensor_text"`
	ScanText             *string           `json:"scan_text"`
	BaseStats            []BaseStat        `json:"base_stats"`
	Items                *MonsterItems     `json:"items"`
	Equipment            *MonsterEquipment `json:"equipment"`
	ElemResists          []ElementalResist `json:"elem_resists"`
	StatusImmunities     []string          `json:"status_immunities"`
	StatusResists        []StatusResist    `json:"status_resists"`
	AlteredStates        []AlteredState    `json:"altered_states"`
	Abilities            []MonsterAbility  `json:"abilities"`
}

func (m Monster) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", m),
		m.Name,
		h.DerefOrNil(m.Version),
		h.DerefOrNil(m.Specification),
		h.DerefOrNil(m.Notes),
		m.Species,
		m.Availability,
		m.IsRepeatable,
		m.CanBeCaptured,
		h.DerefOrNil(m.AreaConquestLocation),
		m.CTBIconType,
		m.HasOverdrive,
		m.IsUnderwater,
		m.IsZombie,
		m.Distance,
		m.AP,
		m.APOverkill,
		m.OverkillDamage,
		m.Gil,
		h.DerefOrNil(m.StealGil),
		h.DerefOrNil(m.DoomCountdown),
		h.DerefOrNil(m.PoisonRate),
		h.DerefOrNil(m.ThreatenChance),
		m.ZanmatoLevel,
		h.DerefOrNil(m.MonsterArenaPrice),
		h.DerefOrNil(m.SensorText),
		h.DerefOrNil(m.ScanText),
	}
}

func (m Monster) ToKeyFields() []any {
	return []any{
		m.Name,
		h.DerefOrNil(m.Version),
	}
}

func (m Monster) GetID() int32 {
	return m.ID
}

func (m Monster) Error() string {
	return h.NameToString(m.Name, m.Version, m.Specification)
}

func (m Monster) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:            m.ID,
		Name:          m.Name,
		Version:       m.Version,
		Specification: m.Specification,
	}
}
