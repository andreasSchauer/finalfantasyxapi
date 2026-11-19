package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
	"golang.org/x/sync/errgroup"
)

func Property(r *http.Request, i handlerinput.Main[seedtypes.Property, apitypes.Property, apiresource.Named], id int32) (apitypes.Property, error) {
	property, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Property{}, err
	}
	
	response := apitypes.Property{
		ID:             property.ID,
		Name:           property.Name,
		Effect:         property.Effect,
		NullifyArmored: property.NullifyArmored,
		RelatedStats:   handlerinput.NamesToNamedAPIResources(apilookups.E.Stats, property.RelatedStats),
		ModifierChange: ah.ConvertObjPtr(property.ModifierChange, apiconvert.ModifierChange),
	}

	err = getPropertyRelationships(r, property, &response)
	if err != nil {
		return apitypes.Property{}, err
	}

	return response, nil
}

func getPropertyRelationships(r *http.Request, property seedtypes.Property, response *apitypes.Property) error {
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		response.AutoAbilities, err = apilookups.E.AutoAbilities.DBGetRelations(ctx, property, handlerinput.ToIntManyNull(config.Cfg.DB.GetPropertyAutoAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		response.Monsters, err = apilookups.E.Monsters.DBGetRelations(ctx, property, config.Cfg.DB.GetPropertyMonsterIDs)
		return err
	})

	return g.Wait()
}
