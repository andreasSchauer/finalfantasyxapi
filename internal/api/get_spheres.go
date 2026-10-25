package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getSphere(r *http.Request, i handlerInput[seeding.Sphere, Sphere, NamedAPIResource, NamedApiResourceList], id int32) (Sphere, error) {
	sphere, err := verifyParamsAndGet(r, i, id)
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

func (cfg *Config) retrieveSpheres(r *http.Request, i handlerInput[seeding.Sphere, Sphere, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		enumListQuery(cfg, r, i, cfg.T.SphereColor, ids, QpnColor, cfg.db.GetSphereIDsByColor),
		idQuery(r, i, ids, QpnLocation, cfg.L.Locations, cfg.db.GetSphereIDsByLocation),
		idQuery(r, i, ids, QpnSublocation, cfg.L.Sublocations, cfg.db.GetSphereIDsBySublocation),
		idQuery(r, i, ids, QpnArea, cfg.L.Areas, cfg.db.GetSphereIDsByArea),
		valueListQuery(cfg, r, i, ids, QpnMethods, cfg.db.GetSphereIDsByMethods),
	})
}
