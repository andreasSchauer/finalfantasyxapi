package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getMonsterFormation(r *http.Request, i handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedApiResourceList], id int32) (MonsterFormation, error) {
	formation, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return MonsterFormation{}, err
	}

	response := MonsterFormation{
		ID:              formation.ID,
		Category:        formation.FormationData.Category,
		Availability:    enumToNamedAPIResource(cfg, cfg.e.availabilityType.endpoint, formation.FormationData.Availability, cfg.t.AvailabilityType),
		IsForcedAmbush:  formation.FormationData.IsForcedAmbush,
		CanEscape:       formation.FormationData.CanEscape,
		Notes:           formation.FormationData.Notes,
		BossMusic:       convertObjPtr(cfg, formation.FormationData.BossMusic, convertFormationBossSong),
		Monsters:        nameAmtsToResAmts(cfg, cfg.e.monsters, formation.Monsters),
		Areas:           convertObjSlice(cfg, formation.EncounterAreas, convertEncounterArea),
		TriggerCommands: convertObjSlice(cfg, formation.TriggerCommands, convertFormationTriggerCommand),
	}

	return response, nil
}

func (cfg *Config) retrieveMonsterFormations(r *http.Request, i handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedApiResourceList]) (UnnamedApiResourceList, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return UnnamedApiResourceList{}, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(idQuery(r, i, ids, qpnMonster, cfg.l.Monsters, cfg.db.GetMonsterFormationIDsByMonster)),
		fidl(idQuery(r, i, ids, qpnLocation, cfg.l.Locations, cfg.db.GetMonsterFormationIDsByLocation)),
		fidl(idQuery(r, i, ids, qpnSublocation, cfg.l.Sublocations, cfg.db.GetMonsterFormationIDsBySublocation)),
		fidl(idQuery(r, i, ids, qpnArea, cfg.l.Areas, cfg.db.GetMonsterFormationIDsByArea)),
		fidl(boolQuery(r, i, ids, qpnAmbush, cfg.db.GetMonsterFormationIDsByForcedAmbush)),
		fidl(enumListQuery(cfg, r, i, cfg.t.MonsterFormationCategory, ids, qpnCategory, cfg.db.GetMonsterFormationIDsByCategory)),
	})
}
