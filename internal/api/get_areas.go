package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getArea(r *http.Request, i handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList], id int32) (Area, error) {
	area, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Area{}, err
	}

	connections, err := getAreaConnectedAreas(cfg, area)
	if err != nil {
		return Area{}, err
	}

	rel, err := getAreaRelationships(cfg, r, area)
	if err != nil {
		return Area{}, err
	}

	response := Area{
		ID:                area.ID,
		Name:              area.Name,
		Version:           area.Version,
		Specification:     area.Specification,
		DisplayName:       getAreaDisplayName(area),
		ParentLocation:    nameToNamedAPIResource(cfg, cfg.e.locations, area.Sublocation.Location.Name, nil),
		ParentSublocation: nameToNamedAPIResource(cfg, cfg.e.sublocations, area.Sublocation.Name, nil),
		StoryOnly:         area.StoryOnly,
		HasSaveSphere:     area.HasSaveSphere,
		AirshipDropOff:    area.AirshipDropOff,
		HasCompSphere:     area.HasCompilationSphere,
		CanRideChocobo:    area.CanRideChocobo,
		ConnectedAreas:    connections,
		LocRel:            rel,
	}

	return response, nil
}

func (cfg *Config) retrieveAreas(r *http.Request, i handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList]) (AreaApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return AreaApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[AreaAPIResource]{
		frl(idQuery(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetLocationAreaIDs)),
		frl(idQuery(cfg, r, i, resources, "sublocation", len(cfg.l.Sublocations), cfg.db.GetSublocationAreaIDs)),
		frl(idQueryWrapper(cfg, r, i, resources, "item", len(cfg.l.Items), getAreasByItem)),
		frl(idQueryWrapper(cfg, r, i, resources, "key_item", len(cfg.l.KeyItems), getAreasByKeyItem)),
		frl(boolQuery(cfg, r, i, resources, "story_based", cfg.db.GetAreaIDsStoryOnly)),
		frl(boolQuery(cfg, r, i, resources, "save_sphere", cfg.db.GetAreaIDsWithSaveSphere)),
		frl(boolQuery(cfg, r, i, resources, "comp_sphere", cfg.db.GetAreaIDsWithCompSphere)),
		frl(boolQuery(cfg, r, i, resources, "airship", cfg.db.GetAreaIDsWithDropOff)),
		frl(boolQuery(cfg, r, i, resources, "chocobo", cfg.db.GetAreaIDsChocobo)),
		frl(boolQuery2(cfg, r, i, resources, "characters", cfg.db.GetAreaIDsWithCharacters)),
		frl(boolQuery2(cfg, r, i, resources, "aeons", cfg.db.GetAreaIDsWithAeons)),
		frl(boolQuery2(cfg, r, i, resources, "monsters", cfg.db.GetAreaIDsWithMonsters)),
		frl(boolQuery2(cfg, r, i, resources, "boss_fights", cfg.db.GetAreaIDsWithBosses)),
		frl(boolQuery2(cfg, r, i, resources, "shops", cfg.db.GetAreaIDsWithShops)),
		frl(boolQuery2(cfg, r, i, resources, "treasures", cfg.db.GetAreaIDsWithTreasures)),
		frl(boolQuery2(cfg, r, i, resources, "sidequests", cfg.db.GetAreaIDsWithSidequests)),
		frl(boolQuery2(cfg, r, i, resources, "fmvs", cfg.db.GetAreaIDsWithFMVs)),
	})
}
