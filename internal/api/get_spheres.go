package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getSphere(r *http.Request, i handlerInput[seeding.Sphere, Sphere, NamedAPIResource, NamedApiResourceList], id int32) (Sphere, error) {
	sphere, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Sphere{}, err
	}

	rel, err := getSphereRelationships(cfg, r, sphere)
	if err != nil {
		return Sphere{}, err
	}

	response := Sphere{
		ID:                 sphere.ID,
		Name:               sphere.Name,
		Item:               rel.Item,
		Description:        rel.Description,
		Effect:             rel.Effect,
		SgDescription:      sphere.SgDescription,
		SphereColor:        sphere.SphereColor,
		SphereEffect:       sphere.SphereEffect,
		TargetNodePosition: sphere.TargetNodePosition,
		TargetNodeState:    sphere.TargetNodeState,
		TargetableNodes:    sphere.TargetableNodes,
		CreatedNode:        convertObjPtr(cfg, sphere.CreatedNode, convertCreatedNode),
		Monsters:           rel.Monsters,
		Treasures:          rel.Treasures,
		Shops:              rel.Shops,
		Quests:             rel.Quests,
		BlitzballPrizes:    rel.BlitzballPrizes,
	}

	return response, nil
}

func (cfg *Config) retrieveSpheres(r *http.Request, i handlerInput[seeding.Sphere, Sphere, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumListQuery(cfg, r, i, cfg.t.SphereColor, ids, "color", cfg.db.GetSphereIDsByColor)),
		fidl(idQuery(r, i, ids, "location", cfg.l.Locations, cfg.db.GetSphereIDsByLocation)),
		fidl(idQuery(r, i, ids, "sublocation", cfg.l.Sublocations, cfg.db.GetSphereIDsBySublocation)),
		fidl(idQuery(r, i, ids, "area", cfg.l.Areas, cfg.db.GetSphereIDsByArea)),
		fidl(valueListQuery(cfg, r, i, ids, "methods", cfg.db.GetSphereIDsByMethods)),
	})
}
