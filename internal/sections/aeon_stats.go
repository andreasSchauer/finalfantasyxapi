package sections

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/typehelpers"
)

func CreateAeonStatSimple(_ *http.Request, id int32, _ handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.Aeons
	aeon, _ := lookup.GetResourceByID(id, lookup.L.AeonsID)

	aeonStatSub := sectiontypes.AeonStatSimple{
		ID:    aeon.ID,
		URL:   ah.CreateResourceURL(i.Endpoint, id),
		Name:  aeon.Name,
		AVals: newAeonStatTable(aeon.BaseStats.AVals),
		BVals: newAeonStatTable(aeon.BaseStats.BVals),
		XVals: ah.ConvertObjSlice(aeon.BaseStats.XVals, convertXVal),
	}

	return aeonStatSub, nil
}

func convertXVal(x seedtypes.XVal) sectiontypes.Xval {
	return sectiontypes.Xval{
		Battles: x.Battles,
		Stats:   newAeonStatTable(x.BaseStats),
	}
}

func newAeonStatTable(baseStats []seedtypes.BaseStat) sectiontypes.AeonStatTable {
	stats := apilookups.E.Stats.ToResAmtType(baseStats, th.NewBaseStat)
	statMap := handlerinput.GetResAmtTypeMap(stats)
	luckPtr := h.GetInt32Ptr(statMap["luck"])

	if *luckPtr == 0 {
		luckPtr = nil
	}

	return sectiontypes.AeonStatTable{
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
