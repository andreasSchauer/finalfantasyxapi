package seeding

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Monster struct {
	ID                   int32
	Name                 string            `json:"name"`
	Version              *int32            `json:"version"`
	Specification        *string           `json:"specification"`
	Notes                *string           `json:"notes"`
	Species              string            `json:"species"`
	IsStoryBased         bool              `json:"is_story_based"`
	CanBeCaptured        bool              `json:"can_be_captured"`
	AreaConquestLocation *string           `json:"area_conquest_location"`
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
	SensorText           string            `json:"sensor_text"`
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
		m.Name,
		derefOrNil(m.Version),
		derefOrNil(m.Specification),
		derefOrNil(m.Notes),
		m.Species,
		m.IsStoryBased,
		m.CanBeCaptured,
		derefOrNil(m.AreaConquestLocation),
		m.CTBIconType,
		m.HasOverdrive,
		m.IsUnderwater,
		m.IsZombie,
		m.Distance,
		m.AP,
		m.APOverkill,
		m.OverkillDamage,
		m.Gil,
		derefOrNil(m.StealGil),
		derefOrNil(m.DoomCountdown),
		derefOrNil(m.PoisonRate),
		derefOrNil(m.ThreatenChance),
		m.ZanmatoLevel,
		derefOrNil(m.MonsterArenaPrice),
		m.SensorText,
		derefOrNil(m.ScanText),
	}
}

func (m Monster) ToKeyFields() []any {
	return []any{
		m.Name,
		derefOrNil(m.Version),
	}
}

func (m Monster) GetID() int32 {
	return m.ID
}

func (m Monster) Error() string {
	return fmt.Sprintf("monster %s, version: %v", m.Name, derefOrNil(m.Version))
}

func (l *Lookup) seedMonsters(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monsters.json"

	var monsters []Monster
	err := loadJSONFile(string(srcPath), &monsters)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, monster := range monsters {
			dbMonster, err := qtx.CreateMonster(context.Background(), database.CreateMonsterParams{
				DataHash:             generateDataHash(monster),
				Name:                 monster.Name,
				Version:              getNullInt32(monster.Version),
				Specification:        getNullString(monster.Specification),
				Notes:                getNullString(monster.Notes),
				Species:              database.MonsterSpecies(monster.Species),
				IsStoryBased:         monster.IsStoryBased,
				CanBeCaptured:        monster.CanBeCaptured,
				AreaConquestLocation: nullMaCreationArea(monster.AreaConquestLocation),
				CtbIconType:          database.CtbIconType(monster.CTBIconType),
				HasOverdrive:         monster.HasOverdrive,
				IsUnderwater:         monster.IsUnderwater,
				IsZombie:             monster.IsZombie,
				Distance:             monster.Distance,
				Ap:                   monster.AP,
				ApOverkill:           monster.APOverkill,
				OverkillDamage:       monster.OverkillDamage,
				Gil:                  monster.Gil,
				StealGil:             getNullInt32(monster.StealGil),
				DoomCountdown:        getNullInt32(monster.DoomCountdown),
				PoisonRate:           getNullFloat64(monster.PoisonRate),
				ThreatenChance:       getNullInt32(monster.ThreatenChance),
				ZanmatoLevel:         monster.ZanmatoLevel,
				MonsterArenaPrice:    getNullInt32(monster.MonsterArenaPrice),
				SensorText:           monster.SensorText,
				ScanText:             getNullString(monster.ScanText),
			})
			if err != nil {
				return getErr(monster.Error(), err, "couldn't create monster")
			}

			monster.ID = dbMonster.ID
			key := createLookupKey(monster)
			l.monsters[key] = monster
		}

		return nil
	})
}

func (l *Lookup) seedMonstersRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monsters.json"

	var monsters []Monster
	err := loadJSONFile(string(srcPath), &monsters)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonMonster := range monsters {
			key := createLookupKey(jsonMonster)

			monster, err := l.getMonster(key)
			if err != nil {
				return err
			}

			err = l.seedMonsterJunctions(qtx, monster)
			if err != nil {
				return getErr(monster.Error(), err)
			}

			if monster.Items != nil {
				var err error
				monster.Items.MonsterID = monster.ID

				monster.Items, err = seedObjPtrAssignFK(qtx, monster.Items, l.seedMonsterItems)
				if err != nil {
					return getErr(monster.Error(), err)
				}
			}

			if monster.Equipment != nil {
				var err error
				monster.Equipment.MonsterID = monster.ID

				monster.Equipment, err = seedObjPtrAssignFK(qtx, monster.Equipment, l.seedMonsterEquipment)
				if err != nil {
					return getErr(monster.Error(), err)
				}
			}

			err = l.seedAlteredStates(qtx, monster)
			if err != nil {
				return getErr(monster.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) seedMonsterJunctions(qtx *database.Queries, monster Monster) error {
	functions := []func(*database.Queries, Monster) error{
		l.seedMonsterProperties,
		l.seedMonsterAutoAbilities,
		l.seedMonsterRonsoRages,
		l.seedMonsterBaseStats,
		l.seedMonsterElemResists,
		l.seedMonsterImmunities,
		l.seedMonsterStatusResists,
		l.seedMonsterAbilities,
	}

	for _, function := range functions {
		err := function(qtx, monster)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lookup) seedMonsterProperties(qtx *database.Queries, monster Monster) error {
	for _, propertyStr := range monster.Properties {
		junction, err := createJunction(monster, propertyStr, l.getProperty)
		if err != nil {
			return err
		}

		err = qtx.CreateMonstersPropertiesJunction(context.Background(), database.CreateMonstersPropertiesJunctionParams{
			DataHash:   generateDataHash(junction),
			MonsterID:  junction.ParentID,
			PropertyID: junction.ChildID,
		})
		if err != nil {
			return getErr(propertyStr, err, "couldn't junction property")
		}
	}

	return nil
}

func (l *Lookup) seedMonsterAutoAbilities(qtx *database.Queries, monster Monster) error {
	for _, autoAbilityStr := range monster.AutoAbilities {
		junction, err := createJunction(monster, autoAbilityStr, l.getAutoAbility)
		if err != nil {
			return err
		}

		err = qtx.CreateMonstersAutoAbilitiesJunction(context.Background(), database.CreateMonstersAutoAbilitiesJunctionParams{
			DataHash:      generateDataHash(junction),
			MonsterID:     junction.ParentID,
			AutoAbilityID: junction.ChildID,
		})
		if err != nil {
			return getErr(autoAbilityStr, err, "couldn't junction auto-ability")
		}
	}

	return nil
}

func (l *Lookup) seedMonsterRonsoRages(qtx *database.Queries, monster Monster) error {
	for _, rage := range monster.RonsoRages {
		key := Ability{
			Name: rage,
		}

		overdrive, err := l.getOverdrive(key)
		if err != nil {
			return err
		}

		if overdrive.User != "kimahri" {
			return getErr(rage, errors.New("overdrive has to be a ronso rage"))
		}

		junction, err := createJunction(monster, key, l.getOverdrive)
		if err != nil {
			return err
		}

		err = qtx.CreateMonstersRonsoRagesJunction(context.Background(), database.CreateMonstersRonsoRagesJunctionParams{
			DataHash:    generateDataHash(junction),
			MonsterID:   junction.ParentID,
			OverdriveID: junction.ChildID,
		})
		if err != nil {
			return getErr(rage, err, "couldn't junction ronso rage")
		}
	}

	return nil
}

func (l *Lookup) seedMonsterBaseStats(qtx *database.Queries, monster Monster) error {
	for _, baseStat := range monster.BaseStats {
		junction, err := createJunctionSeed(qtx, monster, baseStat, l.seedBaseStat)
		if err != nil {
			return err
		}

		err = qtx.CreateMonstersBaseStatsJunction(context.Background(), database.CreateMonstersBaseStatsJunctionParams{
			DataHash:   generateDataHash(junction),
			MonsterID:  junction.ParentID,
			BaseStatID: junction.ChildID,
		})
		if err != nil {
			return getErr(baseStat.Error(), err, "couldn't junction base stat")
		}
	}

	return nil
}

func (l *Lookup) seedMonsterElemResists(qtx *database.Queries, monster Monster) error {
	for _, elemResist := range monster.ElemResists {
		junction, err := createJunctionSeed(qtx, monster, elemResist, l.seedElementalResist)
		if err != nil {
			return err
		}

		err = qtx.CreateMonstersElemResistsJunction(context.Background(), database.CreateMonstersElemResistsJunctionParams{
			DataHash:     generateDataHash(junction),
			MonsterID:    junction.ParentID,
			ElemResistID: junction.ChildID,
		})
		if err != nil {
			return getErr(elemResist.Error(), err, "couldn't junction elemental resist")
		}
	}

	return nil
}

func (l *Lookup) seedMonsterImmunities(qtx *database.Queries, monster Monster) error {
	for _, conditionStr := range monster.StatusImmunities {
		junction, err := createJunction(monster, conditionStr, l.getStatusCondition)
		if err != nil {
			return err
		}

		err = qtx.CreateMonstersImmunitiesJunction(context.Background(), database.CreateMonstersImmunitiesJunctionParams{
			DataHash:          generateDataHash(junction),
			MonsterID:         junction.ParentID,
			StatusConditionID: junction.ChildID,
		})
		if err != nil {
			return getErr(conditionStr, err, "couldn't junction immunity")
		}
	}

	return nil
}

func (l *Lookup) seedMonsterStatusResists(qtx *database.Queries, monster Monster) error {
	for _, statusResist := range monster.StatusResists {
		junction, err := createJunctionSeed(qtx, monster, statusResist, l.seedStatusResist)
		if err != nil {
			return err
		}

		err = qtx.CreateMonstersStatusResistsJunction(context.Background(), database.CreateMonstersStatusResistsJunctionParams{
			DataHash:       generateDataHash(junction),
			MonsterID:      junction.ParentID,
			StatusResistID: junction.ChildID,
		})
		if err != nil {
			return getErr(statusResist.Error(), err, "couldn't junction status resist")
		}
	}

	return nil
}

func (l *Lookup) seedMonsterAbilities(qtx *database.Queries, monster Monster) error {
	for _, ability := range monster.Abilities {
		junction, err := createJunctionSeed(qtx, monster, ability, l.seedMonsterAbility)
		if err != nil {
			return err
		}

		err = qtx.CreateMonstersAbilitiesJunction(context.Background(), database.CreateMonstersAbilitiesJunctionParams{
			DataHash:         generateDataHash(junction),
			MonsterID:        junction.ParentID,
			MonsterAbilityID: junction.ChildID,
		})
		if err != nil {
			return getErr(ability.Error(), err, "couldn't junction monster ability")
		}
	}

	return nil
}
