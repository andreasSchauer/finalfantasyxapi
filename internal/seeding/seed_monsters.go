package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Monster struct {
	//id 		int32
	//dataHash	string
	Name                 string   `json:"name"`
	Version              *int32   `json:"version"`
	Specification        *string  `json:"specification"`
	Notes                *string  `json:"notes"`
	Species              string   `json:"species"`
	IsStoryBased         bool     `json:"is_story_based"`
	CanBeCaptured        bool     `json:"can_be_captured"`
	AreaConquestLocation *string  `json:"area_conquest_location"`
	CTBIconType          string   `json:"ctb_icon_type"`
	HasOverdrive         bool     `json:"has_overdrive"`
	IsUnderwater         bool     `json:"is_underwater"`
	IsZombie             bool     `json:"is_zombie"`
	Distance             int32    `json:"distance"`
	AP                   int32    `json:"ap"`
	APOverkill           int32    `json:"ap_overkill"`
	OverkillDamage       int32    `json:"overkill_damage"`
	Gil                  int32    `json:"gil"`
	StealGil             *int32   `json:"steal_gil"`
	DoomCountdown        *int32   `json:"doom_countdown"`
	PoisonRate           *float32 `json:"poison_rate"`
	ThreatenChance       *int32   `json:"threaten_chance"`
	ZanmatoLevel         int32    `json:"zanmato_level"`
	MonsterArenaPrice    *int32   `json:"monster_arena_price"`
	SensorText           string   `json:"sensor_text"`
	ScanText             *string  `json:"scan_text"`
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

func (l *lookup) seedMonsters(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monsters.json"

	var monsters []Monster
	err := loadJSONFile(string(srcPath), &monsters)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, monster := range monsters {
			err = qtx.CreateMonster(context.Background(), database.CreateMonsterParams{
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
				return fmt.Errorf("couldn't create Monster: %s-%d: %v", monster.Name, derefOrNil(monster.Version), err)
			}
		}
		return nil
	})
}
