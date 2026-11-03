package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (cfg *Config) getSphere(r *http.Request, i handlerInput[seedtypes.Sphere, Sphere, NamedAPIResource], id int32) (Sphere, error) {
	sphere, err := i.queryParamsVerifyGet(r, id)
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

func (cfg *Config) retrieveSpheres(r *http.Request, i handlerInput[seedtypes.Sphere, Sphere, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryEnumList(cfg, r, cfg.T.SphereColor, ids, QpnColor, cfg.db.GetSphereIDsByColor),
		i.queryID(r, ids, QpnLocation, cfg.L.Locations, cfg.db.GetSphereIDsByLocation),
		i.queryID(r, ids, QpnSublocation, cfg.L.Sublocations, cfg.db.GetSphereIDsBySublocation),
		i.queryID(r, ids, QpnArea, cfg.L.Areas, cfg.db.GetSphereIDsByArea),
		i.queryValueList(cfg, r, ids, QpnMethods, cfg.db.GetSphereIDsByMethods),
	})
}
