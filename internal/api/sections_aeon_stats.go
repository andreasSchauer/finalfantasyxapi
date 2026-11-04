package api

import (
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type AeonStatSimple struct {
	ID    int32         `json:"id"`
	URL   string        `json:"url"`
	Name  string        `json:"name"`
	AVals AeonStatTable `json:"a_vals"`
	BVals AeonStatTable `json:"b_vals"`
	XVals []Xval        `json:"x_vals"`
}

func (a AeonStatSimple) GetURL() string {
	return a.URL
}

type Xval struct {
	Battles int32         `json:"battles"`
	Stats   AeonStatTable `json:"stats"`
}

func convertXVal(srv *Server, x seedtypes.XVal) Xval {
	return Xval{
		Battles: x.Battles,
		Stats:   newAeonStatTable(srv, x.BaseStats),
	}
}

type AeonStatTable struct {
	HP           int32  `json:"hp"`
	MP           int32  `json:"mp"`
	Strength     int32  `json:"strength"`
	Defense      int32  `json:"defense"`
	Magic        int32  `json:"magic"`
	MagicDefense int32  `json:"magic_defense"`
	Agility      int32  `json:"agility"`
	Luck         *int32 `json:"luck,omitempty"`
	Evasion      int32  `json:"evasion"`
	Accuracy     int32  `json:"accuracy"`
}

func newAeonStatTable(srv *Server, baseStats []seedtypes.BaseStat) AeonStatTable {
	stats := E.Stats.toResAmtType(srv, baseStats, newBaseStat)
	statMap := GetResAmtTypeMap(stats)
	luckPtr := h.GetInt32Ptr(statMap["luck"])

	if *luckPtr == 0 {
		luckPtr = nil
	}

	return AeonStatTable{
		HP:           statMap["hp"],
		MP:           statMap["mp"],
		Strength:     statMap["strength"],
		Defense:      statMap["defense"],
		Magic:        statMap["magic"],
		MagicDefense: statMap["magic defense"],
		Agility:      statMap["agility"],
		Luck:         luckPtr,
		Evasion:      statMap["evasion"],
		Accuracy:     statMap["accuracy"],
	}
}

func createAeonStatSimple(srv *Server, _ *http.Request, id int32, _ Subsection) (SimpleResource, error) {
	i := E.Aeons
	aeon, _ := lookup.GetResourceByID(id, lookup.L.AeonsID)

	aeonStatSub := AeonStatSimple{
		ID:    aeon.ID,
		URL:   CreateResourceURL(srv, i.Endpoint, id),
		Name:  aeon.Name,
		AVals: newAeonStatTable(srv, aeon.BaseStats.AVals),
		BVals: newAeonStatTable(srv, aeon.BaseStats.BVals),
		XVals: convertObjSlice(srv, aeon.BaseStats.XVals, convertXVal),
	}

	return aeonStatSub, nil
}
