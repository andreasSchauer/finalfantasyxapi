package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getMonster(r *http.Request, i handlerInput[seedtypes.Monster, Monster, NamedAPIResource], id int32) (Monster, error) {
	monster, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Monster{}, err
	}

	rel, err := getMonsterRelationships(s, r, monster)
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
		Properties:           namesToNamedAPIResources(s, E.Properties, monster.Properties),
		AutoAbilities:        namesToNamedAPIResources(s, E.AutoAbilities, monster.AutoAbilities),
		AP:                   monster.AP,
		APOverkill:           monster.APOverkill,
		OverkillDamage:       monster.OverkillDamage,
		Gil:                  monster.Gil,
		StealGil:             monster.StealGil,
		RonsoRages:           namesToNamedAPIResources(s, E.RonsoRages, monster.RonsoRages),
		DoomCountdown:        monster.DoomCountdown,
		PoisonRate:           monster.PoisonRate,
		ThreatenChance:       monster.ThreatenChance,
		ZanmatoLevel:         monster.ZanmatoLevel,
		MonsterArenaPrice:    monster.MonsterArenaPrice,
		SensorText:           monster.SensorText,
		ScanText:             monster.ScanText,
		Areas:                rel.Areas,
		Formations:           rel.Formations,
		BaseStats:            E.Stats.toResAmtType(s, monster.BaseStats, newBaseStat),
		Items:                convertObjPtr(s, monster.Items, convertMonsterItems),
		Equipment:            convertObjPtr(s, monster.Equipment, convertMonsterEquipment),
		ElemResists:          getMonsterElemResists(s, monster.ElemResists),
		StatusImmunities:     namesToNamedAPIResources(s, E.StatusConditions, monster.StatusImmunities),
		StatusResists:        E.StatusConditions.toResAmtType(s, monster.StatusResists, newStatusResist),
		Abilities:            convertObjSlice(s, monster.Abilities, convertMonsterAbility),
		AlteredStates:        getMonsterAlteredStates(s, r, monster),
	}

	return completeMonsterResponse(s, r, response)
}

func (s *Server) retrieveMonsters(r *http.Request, i handlerInput[seedtypes.Monster, Monster, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryBasicWrapper(s, r, ids, ts.QpnElementalResists, getMonstersByElemResists),
		i.queryIDListWrapper(s, r, ids, ts.QpnStatusResists, lookup.L.StatusConditions, getMonstersByStatusResists),

		i.queryIDWrapper(s, r, ids, ts.QpnItem, lookup.L.Items, getMonstersByItem),
		i.queryID(r, ids, ts.QpnRonsoRage, lookup.L.RonsoRages, s.Cfg.DB.GetMonsterIDsByRonsoRage),
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, s.Cfg.DB.GetMonsterIDsByLocation),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, s.Cfg.DB.GetMonsterIDsBySublocation),
		i.queryID(r, ids, ts.QpnArea, lookup.L.Areas, s.Cfg.DB.GetMonsterIDsByArea),
		i.queryIDWrapper(s, r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, getMonstersByAutoAbility),

		i.queryIntList(s, r, ids, ts.QpnEmptySlots, s.Cfg.DB.GetMonsterIDsByEmptySlots),
		i.queryIntList(s, r, ids, ts.QpnDistance, s.Cfg.DB.GetMonsterIDsByDistance),

		i.queryEnumList(s, r, enums.T.MonsterCategory, ids, ts.QpnCategory, s.Cfg.DB.GetMonsterIDsByCategory),
		i.queryEnum(r, enums.T.MonsterSpecies, ids, ts.QpnSpecies, s.Cfg.DB.GetMonsterIDsBySpecies),
		i.queryEnum(r, enums.T.CreationArea, ids, ts.QpnCreationArea, ToEnumQuery(enums.T.CreationArea, s.Cfg.DB.GetMonsterIDsByMaCreationArea)),

		i.queryBool(r, ids, ts.QpnRepeatable, s.Cfg.DB.GetMonsterIDsByIsRepeatable),
		i.queryBool(r, ids, ts.QpnCapture, s.Cfg.DB.GetMonsterIDsByCanBeCaptured),
		i.queryBool(r, ids, ts.QpnHasOverdrive, s.Cfg.DB.GetMonsterIDsByHasOverdrive),
		i.queryBool(r, ids, ts.QpnUnderwater, s.Cfg.DB.GetMonsterIDsByIsUnderwater),
		i.queryBool(r, ids, ts.QpnZombie, s.Cfg.DB.GetMonsterIDsByIsZombie),
	})
}
