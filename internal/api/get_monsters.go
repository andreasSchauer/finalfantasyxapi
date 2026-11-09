package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/typehelpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getMonster(r *http.Request, i handlerinput.Main[seedtypes.Monster, apitypes.Monster, apiresource.Named], id int32) (apitypes.Monster, error) {
	monster, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Monster{}, err
	}

	rel, err := getMonsterRelationships(r, monster)
	if err != nil {
		return apitypes.Monster{}, err
	}

	response := apitypes.Monster{
		ID:                   monster.ID,
		Name:                 monster.Name,
		Version:              monster.Version,
		Specification:        monster.Specification,
		Notes:                monster.Notes,
		Species:              monster.Species,
		Availability:         monster.Availability,
		IsRepeatable:         monster.IsRepeatable,
		CanBeCaptured:        monster.CanBeCaptured,
		AreaConquestLocation: monster.AreaConquestLocation,
		Category:             monster.Category,
		CTBIconType:          monster.CTBIconType,
		HasOverdrive:         monster.HasOverdrive,
		IsUnderwater:         monster.IsUnderwater,
		IsZombie:             monster.IsZombie,
		Distance:             monster.Distance,
		Properties:           handlerinput.NamesToNamedAPIResources(apilookups.E.Properties, monster.Properties),
		AutoAbilities:        handlerinput.NamesToNamedAPIResources(apilookups.E.AutoAbilities, monster.AutoAbilities),
		AP:                   monster.AP,
		APOverkill:           monster.APOverkill,
		OverkillDamage:       monster.OverkillDamage,
		Gil:                  monster.Gil,
		StealGil:             monster.StealGil,
		RonsoRages:           handlerinput.NamesToNamedAPIResources(apilookups.E.RonsoRages, monster.RonsoRages),
		DoomCountdown:        monster.DoomCountdown,
		PoisonRate:           monster.PoisonRate,
		ThreatenChance:       monster.ThreatenChance,
		ZanmatoLevel:         monster.ZanmatoLevel,
		MonsterArenaPrice:    monster.MonsterArenaPrice,
		SensorText:           monster.SensorText,
		ScanText:             monster.ScanText,
		Areas:                rel.Areas,
		Formations:           rel.Formations,
		BaseStats:            apilookups.E.Stats.ToResAmtType(monster.BaseStats, th.NewBaseStat),
		Items:                ah.ConvertObjPtr(monster.Items, apiconv.ConvertMonsterItems),
		Equipment:            ah.ConvertObjPtr(monster.Equipment, apiconv.ConvertMonsterEquipment),
		ElemResists:          getMonsterElemResists(monster.ElemResists),
		StatusImmunities:     handlerinput.NamesToNamedAPIResources(apilookups.E.StatusConditions, monster.StatusImmunities),
		StatusResists:        apilookups.E.StatusConditions.ToResAmtType(monster.StatusResists, th.NewStatusResist),
		Abilities:            ah.ConvertObjSlice(monster.Abilities, apiconv.ConvertMonsterAbility),
		AlteredStates:        getMonsterAlteredStates(r, monster),
	}

	return completeMonsterResponse(r, response)
}

func retrieveMonsters(r *http.Request, i handlerinput.Main[seedtypes.Monster, apitypes.Monster, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryWrapper(r, ids, ts.QpnElementalResists, getMonstersByElemResists),
		i.QueryIDListWrapper(r, ids, ts.QpnStatusResists, lookup.L.StatusConditions, getMonstersByStatusResists),

		i.QueryIDWrapper(r, ids, ts.QpnItem, lookup.L.Items, getMonstersByItem),
		i.QueryID(r, ids, ts.QpnRonsoRage, lookup.L.RonsoRages, config.Cfg.DB.GetMonsterIDsByRonsoRage),
		i.QueryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetMonsterIDsByLocation),
		i.QueryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetMonsterIDsBySublocation),
		i.QueryID(r, ids, ts.QpnArea, lookup.L.Areas, config.Cfg.DB.GetMonsterIDsByArea),
		i.QueryIDWrapper(r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, getMonstersByAutoAbility),

		i.QueryIntList(r, ids, ts.QpnEmptySlots, config.Cfg.DB.GetMonsterIDsByEmptySlots),
		i.QueryIntList(r, ids, ts.QpnDistance, config.Cfg.DB.GetMonsterIDsByDistance),

		i.QueryEnumList(r, enums.T.MonsterCategory, ids, ts.QpnCategory, config.Cfg.DB.GetMonsterIDsByCategory),
		i.QueryEnum(r, enums.T.MonsterSpecies, ids, ts.QpnSpecies, config.Cfg.DB.GetMonsterIDsBySpecies),
		i.QueryEnum(r, enums.T.CreationArea, ids, ts.QpnCreationArea, handlerinput.ToEnumQuery(enums.T.CreationArea, config.Cfg.DB.GetMonsterIDsByMaCreationArea)),

		i.QueryBool(r, ids, ts.QpnRepeatable, config.Cfg.DB.GetMonsterIDsByIsRepeatable),
		i.QueryBool(r, ids, ts.QpnCapture, config.Cfg.DB.GetMonsterIDsByCanBeCaptured),
		i.QueryBool(r, ids, ts.QpnHasOverdrive, config.Cfg.DB.GetMonsterIDsByHasOverdrive),
		i.QueryBool(r, ids, ts.QpnUnderwater, config.Cfg.DB.GetMonsterIDsByIsUnderwater),
		i.QueryBool(r, ids, ts.QpnZombie, config.Cfg.DB.GetMonsterIDsByIsZombie),
	})
}
