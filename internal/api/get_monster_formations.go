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
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return UnnamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[UnnamedAPIResource]{
		frl(idQuery(cfg, r, i, resources, "monster", len(cfg.l.Monsters), cfg.db.GetMonsterFormationIDsByMonster)),
		frl(idQuery(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetMonsterFormationIDsByLocation)),
		frl(idQuery(cfg, r, i, resources, "sublocation", len(cfg.l.Sublocations), cfg.db.GetMonsterFormationIDsBySublocation)),
		frl(idQuery(cfg, r, i, resources, "area", len(cfg.l.Areas), cfg.db.GetMonsterFormationIDsByArea)),
		frl(enumListQuery(cfg, r, i, cfg.t.AvailabilityType, resources, "availability", cfg.db.GetMonsterFormationIDsByAvailability)),
		frl(boolQuery2(cfg, r, i, resources, "repeatable", cfg.db.GetMonsterFormationIDsByRepeatable)),
		frl(boolQuery(cfg, r, i, resources, "ambush", cfg.db.GetMonsterFormationIDsByForcedAmbush)),
		frl(enumListQuery(cfg, r, i, cfg.t.MonsterFormationCategory, resources, "category", cfg.db.GetMonsterFormationIDsByCategory)),
	})
}
