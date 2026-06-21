package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getArea(r *http.Request, i handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList], id int32) (Area, error) {
	area, err := verifyParamsAndGet(cfg, r, i, id)
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
		Availability:      enumToNamedAPIResource(cfg, cfg.e.availabilityType.endpoint, area.Availability, cfg.t.AvailabilityType),
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
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return AreaApiResourceList{}, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(idQuery(r, i, ids, "location", cfg.l.Locations, cfg.db.GetLocationAreaIDs)),
		fidl(idQuery(r, i, ids, "sublocation", cfg.l.Sublocations, cfg.db.GetSublocationAreaIDs)),
		fidl(idQuery(r, i, ids, "monster", cfg.l.Monsters, cfg.db.GetAreaIDsWithMonster)),
		fidl(idQueryWrapper(cfg, r, i, ids, "item", cfg.l.Items, getAreasByItem)),
		fidl(idQuery(r, i, ids, "key_item", cfg.l.KeyItems, cfg.db.GetAreaIDsWithKeyItem)),
		fidl(idQuery(r, i, ids, "monster", cfg.l.Monsters, cfg.db.GetAreasByMonster)),
		fidl(idQuery(r, i, ids, "auto_ability", cfg.l.AutoAbilities, cfg.db.GetAreaIDsWithAutoAbility)),
		fidl(boolQuery(r, i, ids, "save_sphere", cfg.db.GetAreaIDsWithSaveSphere)),
		fidl(boolQuery(r, i, ids, "comp_sphere", cfg.db.GetAreaIDsWithCompSphere)),
		fidl(boolQuery(r, i, ids, "airship", cfg.db.GetAreaIDsWithDropOff)),
		fidl(boolQuery(r, i, ids, "chocobo", cfg.db.GetAreaIDsChocobo)),
		fidl(boolQuery2(r, i, ids, "characters", cfg.db.GetAreaIDsWithCharacters)),
		fidl(boolQuery2(r, i, ids, "aeons", cfg.db.GetAreaIDsWithAeons)),
		fidl(boolQuery2(r, i, ids, "monsters", cfg.db.GetAreaIDsWithMonsters)),
		fidl(boolQuery2(r, i, ids, "boss_fights", cfg.db.GetAreaIDsWithBosses)),
		fidl(boolQuery2(r, i, ids, "shops", cfg.db.GetAreaIDsWithShops)),
		fidl(boolQuery2(r, i, ids, "treasures", cfg.db.GetAreaIDsWithTreasures)),
		fidl(boolQuery2(r, i, ids, "sidequests", cfg.db.GetAreaIDsWithSidequests)),
		fidl(boolQuery2(r, i, ids, "fmvs", cfg.db.GetAreaIDsWithFMVs)),
	})
}
