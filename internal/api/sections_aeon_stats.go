package api

import (
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AeonStatSub struct {
	ID		int32		`json:"id"`
	URL		string		`json:"url"`
	Name	string		`json:"name"`
	AVals	StatTable	`json:"a_vals"`
	BVals	StatTable	`json:"b_vals"`
	XVals	[]Xval		`json:"x_vals"`
}

func (a AeonStatSub) GetURL() string {
	return a.URL
}

type Xval struct {
	Battles	int32		`json:"battles"`
	Stats	StatTable	`json:"stats"`
}

func convertXVal(cfg *Config, x seeding.XVal) Xval {
	return Xval{
		Battles: 	x.Battles,
		Stats: 		newStatTable(cfg, x.BaseStats),
	}
}

type StatTable struct {
	HP				int32	`json:"hp"`
	MP				int32	`json:"mp"`
	Strength		int32	`json:"strength"`
	Defense			int32	`json:"defense"`
	Magic			int32	`json:"magic"`
	MagicDefense	int32	`json:"magic_defense"`
	Agility			int32	`json:"agility"`
	Luck			*int32	`json:"luck,omitempty"`
	Evasion			int32	`json:"evasion"`
	Accuracy		int32	`json:"accuracy"`
}

func newStatTable(cfg *Config, baseStats []seeding.BaseStat) StatTable {
	stats := namesToResourceAmounts(cfg, cfg.e.stats, baseStats, newBaseStat)
	statMap := getResourceAmountMap(stats)
	luckPtr := h.GetInt32Ptr(statMap["luck"])

	if *luckPtr == 0 {
		luckPtr = nil
	}

	return StatTable{
		HP: 			statMap["hp"],
		MP: 			statMap["mp"],
		Strength: 		statMap["strength"],
		Defense: 		statMap["defense"],
		Magic: 			statMap["magic"],
		MagicDefense: 	statMap["magic defense"],
		Agility: 		statMap["agility"],
		Luck: 			luckPtr,
		Evasion: 		statMap["evasion"],
		Accuracy: 		statMap["accuracy"],
	}
}

func createAeonStatSub(cfg *Config, _ *http.Request, id int32) (SubResource, error) {
	i := cfg.e.aeons
	aeon, _ := seeding.GetResourceByID(id, cfg.l.AeonsID)

	aeonStatSub := AeonStatSub{
		ID: 	aeon.ID,
		URL: 	createResourceURL(cfg, i.endpoint, id),
		Name: 	aeon.Name,
		AVals: 	newStatTable(cfg, aeon.BaseStats.AVals),
		BVals: 	newStatTable(cfg, aeon.BaseStats.BVals),
		XVals: 	convertObjSlice(cfg, aeon.BaseStats.XVals, convertXVal),
	}

	return aeonStatSub, nil
}