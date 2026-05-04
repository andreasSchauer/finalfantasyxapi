package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
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

func (l *Lookup) loop1SeedMonsters(qtx *database.Queries, ctx context.Context) error {
	err := l.completeMonstersElements()
	if err != nil {
		return err
	}
	monsters := dedupeRows(l.json.monsters, l.Hashes)

	params := database.CreateMonsterBulkParams{
		DataHash:             make([]string, len(monsters)),
		Name:                 make([]string, len(monsters)),
		Version:              make([]sql.NullInt32, len(monsters)),
		Specification:        make([]sql.NullString, len(monsters)),
		Notes:                make([]sql.NullString, len(monsters)),
		Species:              make([]database.MonsterSpecies, len(monsters)),
		Availability:         make([]database.AvailabilityType, len(monsters)),
		IsRepeatable:         make([]bool, len(monsters)),
		CanBeCaptured:        make([]bool, len(monsters)),
		AreaConquestLocation: make([]database.NullMaCreationArea, len(monsters)),
		Category:             make([]database.MonsterCategory, len(monsters)),
		CtbIconType:          make([]database.CtbIconType, len(monsters)),
		HasOverdrive:         make([]bool, len(monsters)),
		IsUnderwater:         make([]bool, len(monsters)),
		IsZombie:             make([]bool, len(monsters)),
		Distance:             make([]int32, len(monsters)),
		Ap:                   make([]int32, len(monsters)),
		ApOverkill:           make([]int32, len(monsters)),
		OverkillDamage:       make([]int32, len(monsters)),
		Gil:                  make([]int32, len(monsters)),
		StealGil:             make([]sql.NullInt32, len(monsters)),
		DoomCountdown:        make([]sql.NullInt32, len(monsters)),
		PoisonRate:           make([]sql.NullFloat64, len(monsters)),
		ThreatenChance:       make([]sql.NullInt32, len(monsters)),
		ZanmatoLevel:         make([]int32, len(monsters)),
		MonsterArenaPrice:    make([]sql.NullInt32, len(monsters)),
		SensorText:           make([]sql.NullString, len(monsters)),
		ScanText:             make([]sql.NullString, len(monsters)),
	}

	for i, m := range monsters {
		params.DataHash[i] = generateDataHash(m)
		params.Name[i] = m.Name
		params.Version[i] = h.GetNullInt32(m.Version)
		params.Specification[i] = h.GetNullString(m.Specification)
		params.Notes[i] = h.GetNullString(m.Notes)
		params.Species[i] = database.MonsterSpecies(m.Species)
		params.Availability[i] = database.AvailabilityType(m.Availability)
		params.IsRepeatable[i] = m.IsRepeatable
		params.CanBeCaptured[i] = m.CanBeCaptured
		params.AreaConquestLocation[i] = database.ToNullMaCreationArea(m.AreaConquestLocation)
		params.Category[i] = database.MonsterCategory(m.Category)
		params.CtbIconType[i] = database.CtbIconType(m.CTBIconType)
		params.HasOverdrive[i] = m.HasOverdrive
		params.IsUnderwater[i] = m.IsUnderwater
		params.IsZombie[i] = m.IsZombie
		params.Distance[i] = m.Distance
		params.Ap[i] = m.AP
		params.ApOverkill[i] = m.APOverkill
		params.OverkillDamage[i] = m.OverkillDamage
		params.Gil[i] = m.Gil
		params.StealGil[i] = h.GetNullInt32(m.StealGil)
		params.DoomCountdown[i] = h.GetNullInt32(m.DoomCountdown)
		params.PoisonRate[i] = h.GetNullFloat64(m.PoisonRate)
		params.ThreatenChance[i] = h.GetNullInt32(m.ThreatenChance)
		params.ZanmatoLevel[i] = m.ZanmatoLevel
		params.MonsterArenaPrice[i] = h.GetNullInt32(m.MonsterArenaPrice)
		params.SensorText[i] = h.GetNullString(m.SensorText)
		params.ScanText[i] = h.GetNullString(m.ScanText)
	}

	dbRows, err := qtx.CreateMonsterBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create monsters: %v", err)
	}

	for i, row := range dbRows {
		monsters[i].ID = row.ID
		l.json.monsters[i].ID = row.ID
		key := Key(monsters[i])
		l.Monsters[key] = monsters[i]
		l.MonstersID[row.ID] = monsters[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) completeMonsters() error {
	for i := range l.json.monsters {
		mon := &l.json.monsters[i]

		err := assignIDs(l, mon.ElemResists)
		if err != nil {
			return err
		}

		err = assignIDs(l, mon.StatusResists)
		if err != nil {
			return err
		}

		err = assignIDs(l, mon.BaseStats)
		if err != nil {
			return err
		}

		err = assignIDs(l, mon.Abilities)
		if err != nil {
			return err
		}

		err = l.completeMonsterItems(mon.Items)
		if err != nil {
			return err
		}

		err = l.completeMonsterEquipment(mon.Equipment)
		if err != nil {
			return err
		}

		err = l.completeAlteredStates(mon.AlteredStates)
		if err != nil {
			return err
		}

		l.Monsters[Key(*mon)] = *mon
		l.MonstersID[mon.ID] = *mon
	}

	return nil
}

func (l *Lookup) completeMonstersElements() error {
	elements := []string{"fire", "lightning", "water", "ice", "holy"}

	for i, mon := range l.json.monsters {
		elemResistLookup := make(map[string]string)

		for _, elemResist := range mon.ElemResists {
			elemResistLookup[elemResist.Element] = elemResist.Affinity
		}

		resists := mon.ElemResists
		for _, element := range elements {
			_, found := elemResistLookup[element]
			if !found {
				elemResist := ElementalResist{
					Element:  element,
					Affinity: "neutral",
				}
				resists = append(resists, elemResist)
			}
		}
		l.json.monsters[i].ElemResists = resists
	}

	return nil
}

func (l *Lookup) getMonsterMonsterAbilities(m Monster) ([]MonsterAbility, error) {
	return m.Abilities, nil
}

func (l *Lookup) getMonsterAutoAbilities(m Monster) ([]AutoAbility, error) {
	return getResources(m.AutoAbilities, l.AutoAbilities)
}

func (l *Lookup) getMonsterBaseStats(m Monster) ([]BaseStat, error) {
	return m.BaseStats, nil
}

func (l *Lookup) getMonsterElementalResists(m Monster) ([]ElementalResist, error) {
	return m.ElemResists, nil
}

func (l *Lookup) getMonsterStatusImmunities(m Monster) ([]StatusCondition, error) {
	return getResources(m.StatusImmunities, l.StatusConditions)
}

func (l *Lookup) getMonsterProperties(m Monster) ([]Property, error) {
	return getResources(m.Properties, l.Properties)
}

func (l *Lookup) getMonsterRonsoRages(m Monster) ([]RonsoRage, error) {
	return getResources(m.RonsoRages, l.RonsoRages)
}

func (l *Lookup) getMonsterStatusResists(m Monster) ([]StatusResist, error) {
	return m.StatusResists, nil
}

func (l *Lookup) seedJuncMonstersMonsterAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + monster abilities"
	jParams, err := processJunctions(l, desc, l.json.monsters, l.getMonsterMonsterAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersAbilitiesJunctionBulk(ctx, database.CreateMonstersAbilitiesJunctionBulkParams{
		DataHash:         jParams.DataHashes,
		MonsterID:        jParams.ParentIDs,
		MonsterAbilityID: jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncMonstersAutoAbilities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + auto-abilities"
	jParams, err := processJunctions(l, desc, l.json.monsters, l.getMonsterAutoAbilities)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersAutoAbilitiesJunctionBulk(ctx, database.CreateMonstersAutoAbilitiesJunctionBulkParams{
		DataHash:      jParams.DataHashes,
		MonsterID:     jParams.ParentIDs,
		AutoAbilityID: jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncMonstersBaseStats(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + base stats"
	jParams, err := processJunctions(l, desc, l.json.monsters, l.getMonsterBaseStats)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersBaseStatsJunctionBulk(ctx, database.CreateMonstersBaseStatsJunctionBulkParams{
		DataHash:   jParams.DataHashes,
		MonsterID:  jParams.ParentIDs,
		BaseStatID: jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncMonstersElementalResists(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + elemental resists"
	jParams, err := processJunctions(l, desc, l.json.monsters, l.getMonsterElementalResists)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersElemResistsJunctionBulk(ctx, database.CreateMonstersElemResistsJunctionBulkParams{
		DataHash:     jParams.DataHashes,
		MonsterID:    jParams.ParentIDs,
		ElemResistID: jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncMonstersStatusImmunities(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + status immunities"
	jParams, err := processJunctions(l, desc, l.json.monsters, l.getMonsterStatusImmunities)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersImmunitiesJunctionBulk(ctx, database.CreateMonstersImmunitiesJunctionBulkParams{
		DataHash:          jParams.DataHashes,
		MonsterID:         jParams.ParentIDs,
		StatusConditionID: jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncMonstersProperties(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + properties"
	jParams, err := processJunctions(l, desc, l.json.monsters, l.getMonsterProperties)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersPropertiesJunctionBulk(ctx, database.CreateMonstersPropertiesJunctionBulkParams{
		DataHash:   jParams.DataHashes,
		MonsterID:  jParams.ParentIDs,
		PropertyID: jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncMonstersRonsoRages(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + ronso rages"
	jParams, err := processJunctions(l, desc, l.json.monsters, l.getMonsterRonsoRages)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersRonsoRagesJunctionBulk(ctx, database.CreateMonstersRonsoRagesJunctionBulkParams{
		DataHash:    jParams.DataHashes,
		MonsterID:   jParams.ParentIDs,
		RonsoRageID: jParams.ChildIDs,
	})
}

func (l *Lookup) seedJuncMonstersStatusResists(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monsters + status resists"
	jParams, err := processJunctions(l, desc, l.json.monsters, l.getMonsterStatusResists)
	if err != nil {
		return err
	}

	return qtx.CreateMonstersStatusResistsJunctionBulk(ctx, database.CreateMonstersStatusResistsJunctionBulkParams{
		DataHash:       jParams.DataHashes,
		MonsterID:      jParams.ParentIDs,
		StatusResistID: jParams.ChildIDs,
	})
}
