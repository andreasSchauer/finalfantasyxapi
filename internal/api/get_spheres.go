package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getSphere(r *http.Request, i handlerInput[seedtypes.Sphere, Sphere, NamedAPIResource], id int32) (Sphere, error) {
	sphere, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Sphere{}, err
	}

	rel, err := getSphereRelationships(s, r, sphere)
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
		CreatedNode:        convertObjPtr(s, sphere.CreatedNode, convertCreatedNode),
		Monsters:           rel.Monsters,
		Treasures:          rel.Treasures,
		Shops:              rel.Shops,
		Quests:             rel.Quests,
		BlitzballPrizes:    rel.BlitzballPrizes,
	}

	return response, nil
}

func (s *Server) retrieveSpheres(r *http.Request, i handlerInput[seedtypes.Sphere, Sphere, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryEnumList(s, r, enums.T.SphereColor, ids, ts.QpnColor, s.Cfg.DB.GetSphereIDsByColor),
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, s.Cfg.DB.GetSphereIDsByLocation),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, s.Cfg.DB.GetSphereIDsBySublocation),
		i.queryID(r, ids, ts.QpnArea, lookup.L.Areas, s.Cfg.DB.GetSphereIDsByArea),
		i.queryValueList(s, r, ids, ts.QpnMethods, s.Cfg.DB.GetSphereIDsByMethods),
	})
}
