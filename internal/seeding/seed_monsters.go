package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Monster struct {
	ID					 	int32
	Name                 	string   			`json:"name"`
	Version              	*int32   			`json:"version"`
	Specification        	*string  			`json:"specification"`
	Notes                	*string  			`json:"notes"`
	Species              	string   			`json:"species"`
	IsStoryBased         	bool     			`json:"is_story_based"`
	CanBeCaptured        	bool     			`json:"can_be_captured"`
	AreaConquestLocation 	*string  			`json:"area_conquest_location"`
	CTBIconType          	string   			`json:"ctb_icon_type"`
	HasOverdrive         	bool     			`json:"has_overdrive"`
	IsUnderwater         	bool     			`json:"is_underwater"`
	IsZombie             	bool     			`json:"is_zombie"`
	Distance             	int32    			`json:"distance"`
	Properties				[]string			`json:"properties"`
	AutoAbilities			[]string			`json:"auto_abilities"`
	AP                   	int32    			`json:"ap"`
	APOverkill           	int32    			`json:"ap_overkill"`
	OverkillDamage       	int32    			`json:"overkill_damage"`
	Gil                  	int32    			`json:"gil"`
	StealGil             	*int32   			`json:"steal_gil"`
	RonsoRange				[]string			`json:"ronso_rage"`
	DoomCountdown       	*int32   			`json:"doom_countdown"`
	PoisonRate           	*float32 			`json:"poison_rate"`
	ThreatenChance       	*int32   			`json:"threaten_chance"`
	ZanmatoLevel         	int32    			`json:"zanmato_level"`
	MonsterArenaPrice    	*int32   			`json:"monster_arena_price"`
	SensorText           	string   			`json:"sensor_text"`
	ScanText             	*string  			`json:"scan_text"`
	Stats					[]BaseStat			`json:"stats"`
	Items					*MonsterItems		`json:"items"`
	Equipment				*MonsterEquipment	`json:"equipment"`
	ElemResists				[]ElementalResist	`json:"elem_resists"`
	StatusImmunities		[]string			`json:"status_immunities"`
	StatusResistances		[]StatusResist		`json:"status_resistances"`
	AlteredStates			[]AlteredState		`json:"altered_states"`
	Abilities				[]MonsterAbility	`json:"abilities"`
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


type MonsterItems struct {
	DropChance			int32				`json:"drop_chance"`
	DropCondition		*string				`json:"drop_condition"`
	OtherItems			*MonsterOtherItems	`json:"other_items"`
	StealCommon			*ItemAmount			`json:"steal_common"`
	StealRare			*ItemAmount			`json:"steal_rare"`
	DropCommon			*ItemAmount			`json:"drop_common"`
	DropRare			*ItemAmount			`json:"drop_rare"`
	SecondaryDropCommon	*ItemAmount			`json:"secondary_drop_common"`
	SecondaryDropRare	*ItemAmount			`json:"secondary_drop_rare"`
	Bribe				*ItemAmount			`json:"bribe"`
}


type MonsterOtherItems struct {
	Condition		string			`json:"condition"`
	Items			[]PossibleItem	`json:"items"`
}


type MonsterEquipment struct {
	DropChance			int32					`json:"drop_chance"`
	Power				int32					`json:"power"`
	CriticalPlus		int32					`json:"critical_plus"`
	AbilitySlots		MonsterEquipmentSlots	`json:"ability_slots"`
	AttachedAbilities	MonsterEquipmentSlots	`json:"attached_abilities"`
	WeaponAbilities		[]EquipmentDrop			`json:"weapon_abilities"`
	ArmorAbilities		[]EquipmentDrop			`json:"armor_abilities"`
}


type MonsterEquipmentSlots struct {
	MinAmount		int32					`json:"min_amount"`
	MaxAmount		int32					`json:"max_amount"`
	Chances			[]EquipmentSlotsChance	`json:"chances"`
}


type EquipmentSlotsChance struct {
	Amount		int32	`json:"amount"`
	Chance		int32	`json:"chance"`
}

type EquipmentDrop struct {
	Ability			string		`json:"ability"`
	Characters		[]string	`json:"characters"`
	IsForced		bool		`json:"is_forced"`
	Probability		*int32		`json:"probability"`
}

type AlteredState struct {
	Condition		string				`json:"condition"`
	IsTemporary		bool				`json:"is_temporary"`
	Changes			[]AltStateChange 	`json:"changes"`
}

type AltStateChange struct {
	AlterationType		string				`json:"alteration_type"`
	Properties			*[]string			`json:"properties"`
	AutoAbilities		*[]string			`json:"auto_abilities"`
	Distance			*int32				`json:"distance"`
	Stats				*[]BaseStat			`json:"stats"`
	ElemResists			*[]ElementalResist	`json:"elem_resists"`
	StatusImmunities	*[]string			`json:"status_immunities"`
	AddedStatusses		*[]InflictedStatus	`json:"added_statusses"`
}

type MonsterAbility struct {
	AbilityReference
	IsForced			bool	`json:"is_forced"`
	IsUnused			bool	`json:"is_unused"`
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



func (l *lookup) seedMonstersRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "./data/monsters.json"

	var monsters []Monster
	err := loadJSONFile(string(srcPath), &monsters)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {


		return nil
	})
}