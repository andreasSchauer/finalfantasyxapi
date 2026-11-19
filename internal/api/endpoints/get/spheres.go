package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Sphere(r *http.Request, i handlerinput.Main[seedtypes.Sphere, apitypes.Sphere, apiresource.Named], id int32) (apitypes.Sphere, error) {
	sphere, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Sphere{}, err
	}
	
	response := apitypes.Sphere{
		ID:                 sphere.ID,
		Name:               sphere.Name,
		SgDescription:      sphere.SgDescription,
		SphereColor:        sphere.SphereColor,
		SphereEffect:       sphere.SphereEffect,
		TargetNodePosition: sphere.TargetNodePosition,
		TargetNodeState:    sphere.TargetNodeState,
		TargetableNodes:    sphere.TargetableNodes,
		CreatedNode:        ah.ConvertObjPtr(sphere.CreatedNode, apiconvert.CreatedNode),
	}

	err = getSphereRelationships(r, sphere, &response)
	if err != nil {
		return apitypes.Sphere{}, err
	}

	return response, nil
}

func getSphereRelationships(r *http.Request, sphere seedtypes.Sphere, response *apitypes.Sphere) error {
	item, _ := lookup.GetResourceByID(sphere.ItemID, lookup.L.ItemsID)

	availabilityParams, err := apilookups.E.Spheres.QueryAvlRelGetParams(r, item.ID)
	if err != nil {
		return err
	}

	var itemRel apitypes.Item
	err = runItemRelQueries(r, item, &itemRel, availabilityParams)
	if err != nil {
		return err
	}

	response.Item = handlerinput.NameToNamedAPIResource(apilookups.E.Items, item.Name, nil)
	response.Description = item.Description
	response.Effect = item.Effect
	response.Monsters = itemRel.Monsters
	response.Treasures = itemRel.Treasures
	response.Shops = itemRel.Shops
	response.Quests = itemRel.Quests
	response.BlitzballPrizes = itemRel.BlitzballPrizes

	return nil
}
