package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Character struct {
	ID int32
	PlayerUnit
	LocationArea       LocationArea `json:"location_area"`
	AreaID             *int32
	IsStoryBased       bool       	`json:"is_story_based"`
	WeaponType         string     	`json:"weapon_type"`
	ArmorType          string     	`json:"armor_type"`
	PhysAtkRange       int32      	`json:"physical_attack_range"`
	CanFightUnderwater bool       	`json:"can_fight_underwater"`
	BaseStats          []BaseStat 	`json:"base_stats"`
	StdSphereGrid	   *SphereGrid	`json:"std_sphere_grid"`
	ExpSphereGrid	   *SphereGrid	`json:"exp_sphere_grid"`
}

func (c Character) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", c),
		c.PlayerUnit.ID,
		h.DerefOrNil(c.AreaID),
		c.IsStoryBased,
		c.WeaponType,
		c.ArmorType,
		c.PhysAtkRange,
		c.CanFightUnderwater,
		h.ObjPtrToID(c.StdSphereGrid),
		h.ObjPtrToID(c.ExpSphereGrid),
	}
}

func (c Character) ToKeyFields() []any {
	return []any{
		c.Name,
	}
}

func (c Character) GetID() int32 {
	return c.ID
}

func (c Character) GetResParamsNamed() ResParamsNamed {
	return ResParamsNamed{
		ID:   c.ID,
		Name: c.Name,
	}
}

func (c Character) Error() string {
	return fmt.Sprintf("character %s", c.Name)
}

type SphereGrid struct {
	ID			 int32
	Type		 database.SphereGridType
	StatTable
	Lv1Locks	 int32	`json:"lv_1_locks"`
	Lv2Locks	 int32	`json:"lv_2_locks"`
	Lv3Locks	 int32	`json:"lv_3_locks"`
	Lv4Locks	 int32	`json:"lv_4_locks"`
	EmptyNodes	 int32	`json:"empty_nodes"`
}

func (s SphereGrid) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.Type,
		s.HP,
		s.MP,
		s.Strength,
		s.Defense,
		s.Magic,
		s.MagicDefense,
		s.Agility,
		s.Luck,
		s.Evasion,
		s.Accuracy,
		s.Lv1Locks,
		s.Lv2Locks,
		s.Lv3Locks,
		s.Lv4Locks,
		s.EmptyNodes,
	}
}

func (s SphereGrid) GetID() int32 {
	return s.ID
}

func (s *SphereGrid) SetID(id int32) {
	s.ID = id
}

func (s SphereGrid) Error() string {
	return fmt.Sprintf("sphere grid with type: %s, hp: %d, mp: %d, strength: %d, defense: %d, magic: %d, magic defense: %d, agility: %d, luck: %d, evasion: %d, accuracy: %d, lv 1 locks: %d, lv 2 locks: %d, lv 3 locks: %d, lv 4 locks: %d, empty nodes: %d", s.Type, s.HP, s.MP, s.Strength, s.Defense, s.Magic, s.MagicDefense, s.Agility, s.Luck, s.Evasion, s.Accuracy, s.Lv1Locks, s.Lv2Locks, s.Lv3Locks, s.Lv4Locks, s.EmptyNodes)
}


type StatTable struct {
	HP           int32  `json:"hp"`
	MP           int32  `json:"mp"`
	Strength     int32  `json:"strength"`
	Defense      int32  `json:"defense"`
	Magic        int32  `json:"magic"`
	MagicDefense int32  `json:"magic_defense"`
	Agility      int32  `json:"agility"`
	Luck         int32  `json:"luck"`
	Evasion      int32  `json:"evasion"`
	Accuracy     int32  `json:"accuracy"`
}

func (s StatTable) ToMap() map[string]int32 {
	return map[string]int32{
		"hp": 				s.HP,
		"mp": 				s.MP,
		"strength": 		s.Strength,
		"defense": 			s.Defense,
		"magic": 			s.Magic,
		"magic defense": 	s.MagicDefense,
		"agility": 			s.Agility,
		"luck": 			s.Luck,
		"evasion": 			s.Evasion,
		"accuracy": 		s.Accuracy,
	}
}