package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type Stats struct {
	HP				BaseStat	`json:"hp"`
	MP				BaseStat	`json:"mp"`
	Strength		BaseStat	`json:"strength"`
	Defense			BaseStat	`json:"defense"`
	Magic			BaseStat	`json:"magic"`
	MagicDefense	BaseStat	`json:"magic_defense"`
	Agility			BaseStat	`json:"agility"`
	Luck			BaseStat	`json:"luck"`
	Evasion			BaseStat	`json:"evasion"`
	Accuracy		BaseStat	`json:"accuracy"`
}

func createStats(bs []BaseStat) Stats {
	statMap := getResourceMap(bs)
	return Stats{
		HP: 			statMap[1],
		MP: 			statMap[2],
		Strength: 		statMap[3],
		Defense: 		statMap[4],
		Magic: 			statMap[5],
		MagicDefense: 	statMap[6],
		Agility: 		statMap[7],
		Luck: 			statMap[8],
		Evasion: 		statMap[9],
		Accuracy: 		statMap[10],
	}
}

func statsToBaseStats(s Stats) []BaseStat {
	return []BaseStat{
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
	}
}


type BaseStat struct {
	Stat  NamedAPIResource `json:"stat"`
	Value int32            `json:"value"`
}

func (bs BaseStat) GetAPIResource() APIResource {
	return bs.Stat
}

func newBaseStat(res NamedAPIResource, value int32) BaseStat {
	return BaseStat{
		Stat:  res,
		Value: value,
	}
}

func (bs BaseStat) GetName() string {
	return bs.Stat.Name
}

func (bs BaseStat) GetVersion() *int32 {
	return nil
}

func (bs BaseStat) GetVal() int32 {
	return bs.Value
}

func (bs BaseStat) ToResAmount() ResourceAmount[NamedAPIResource] {
	return ResourceAmount[NamedAPIResource]{
		Resource: bs.Stat,
		Amount:   bs.Value,
	}
}

func getBaseStat(cfg *Config, stat string, baseStats []BaseStat) BaseStat {
	statLookup, _ := seeding.GetResource(stat, cfg.l.Stats)
	statMap := getResourceMap(baseStats)

	return statMap[statLookup.ID]
}

func replaceBaseStats(baseStats []BaseStat, statMap map[string]int32) []BaseStat {
	for i, baseStat := range baseStats {
		newVal, ok := statMap[baseStat.Stat.Name]
		if ok {
			baseStats[i].Value = newVal
		}
	}

	return baseStats
}


func addToBaseStats(baseStats []BaseStat, statMap map[string]int32) []BaseStat {
	for i, baseStat := range baseStats {
		toAdd, ok := statMap[baseStat.Stat.Name]
		if ok {
			baseStats[i].Value += toAdd
		}
	}

	return baseStats
}