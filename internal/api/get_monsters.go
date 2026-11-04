package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getMonster(r *http.Request, i handlerInput[seedtypes.Monster, Monster, NamedAPIResource], id int32) (Monster, error) {
	monster, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Monster{}, err
	}

	rel, err := getMonsterRelationships(srv, r, monster)
	if err != nil {
		return Monster{}, err
	}

	response := Monster{
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
		Properties:           namesToNamedAPIResources(srv, E.Properties, monster.Properties),
		AutoAbilities:        namesToNamedAPIResources(srv, E.AutoAbilities, monster.AutoAbilities),
		AP:                   monster.AP,
		APOverkill:           monster.APOverkill,
		OverkillDamage:       monster.OverkillDamage,
		Gil:                  monster.Gil,
		StealGil:             monster.StealGil,
		RonsoRages:           namesToNamedAPIResources(srv, E.RonsoRages, monster.RonsoRages),
		DoomCountdown:        monster.DoomCountdown,
		PoisonRate:           monster.PoisonRate,
		ThreatenChance:       monster.ThreatenChance,
		ZanmatoLevel:         monster.ZanmatoLevel,
		MonsterArenaPrice:    monster.MonsterArenaPrice,
		SensorText:           monster.SensorText,
		ScanText:             monster.ScanText,
		Areas:                rel.Areas,
		Formations:           rel.Formations,
		BaseStats:            E.Stats.toResAmtType(srv, monster.BaseStats, newBaseStat),
		Items:                convertObjPtr(srv, monster.Items, convertMonsterItems),
		Equipment:            convertObjPtr(srv, monster.Equipment, convertMonsterEquipment),
		ElemResists:          getMonsterElemResists(srv, monster.ElemResists),
		StatusImmunities:     namesToNamedAPIResources(srv, E.StatusConditions, monster.StatusImmunities),
		StatusResists:        E.StatusConditions.toResAmtType(srv, monster.StatusResists, newStatusResist),
		Abilities:            convertObjSlice(srv, monster.Abilities, convertMonsterAbility),
		AlteredStates:        getMonsterAlteredStates(srv, r, monster),
	}

	return completeMonsterResponse(srv, r, response)
}

func (srv *Server) retrieveMonsters(r *http.Request, i handlerInput[seedtypes.Monster, Monster, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryBasicWrapper(srv, r, ids, ts.QpnElementalResists, getMonstersByElemResists),
		i.queryIDListWrapper(srv, r, ids, ts.QpnStatusResists, lookup.L.StatusConditions, getMonstersByStatusResists),

		i.queryIDWrapper(srv, r, ids, ts.QpnItem, lookup.L.Items, getMonstersByItem),
		i.queryID(r, ids, ts.QpnRonsoRage, lookup.L.RonsoRages, srv.Cfg.DB.GetMonsterIDsByRonsoRage),
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, srv.Cfg.DB.GetMonsterIDsByLocation),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, srv.Cfg.DB.GetMonsterIDsBySublocation),
		i.queryID(r, ids, ts.QpnArea, lookup.L.Areas, srv.Cfg.DB.GetMonsterIDsByArea),
		i.queryIDWrapper(srv, r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, getMonstersByAutoAbility),

		i.queryIntList(srv, r, ids, ts.QpnEmptySlots, srv.Cfg.DB.GetMonsterIDsByEmptySlots),
		i.queryIntList(srv, r, ids, ts.QpnDistance, srv.Cfg.DB.GetMonsterIDsByDistance),

		i.queryEnumList(srv, r, enums.T.MonsterCategory, ids, ts.QpnCategory, srv.Cfg.DB.GetMonsterIDsByCategory),
		i.queryEnum(r, enums.T.MonsterSpecies, ids, ts.QpnSpecies, srv.Cfg.DB.GetMonsterIDsBySpecies),
		i.queryEnum(r, enums.T.CreationArea, ids, ts.QpnCreationArea, ToEnumQuery(enums.T.CreationArea, srv.Cfg.DB.GetMonsterIDsByMaCreationArea)),

		i.queryBool(r, ids, ts.QpnRepeatable, srv.Cfg.DB.GetMonsterIDsByIsRepeatable),
		i.queryBool(r, ids, ts.QpnCapture, srv.Cfg.DB.GetMonsterIDsByCanBeCaptured),
		i.queryBool(r, ids, ts.QpnHasOverdrive, srv.Cfg.DB.GetMonsterIDsByHasOverdrive),
		i.queryBool(r, ids, ts.QpnUnderwater, srv.Cfg.DB.GetMonsterIDsByIsUnderwater),
		i.queryBool(r, ids, ts.QpnZombie, srv.Cfg.DB.GetMonsterIDsByIsZombie),
	})
}
