package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
)

func filterAvlMonsters(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := apilookups.E.Monsters

	avlParams, err := i.QueryAvlRepGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	locContext, err := getLocContextParams(r, i)
	if err != nil {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.FilterMonsterIDsByAvailability(r.Context(), database.FilterMonsterIDsByAvailabilityParams{
		Ids:            inputIDs,
		Availability:   avlParams.Availabilities,
		IsRepeatable:   avlParams.IsRepeatable,
		PreAirship:     avlParams.PreAirship,
		AvlType:        locContext.AvlType,
		LocContextID:   locContext.ID,
		LocContextType: locContext.Type,
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}

func filterAvlMonsterFormations(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := apilookups.E.MonsterFormations

	avlParams, err := i.QueryAvlRepGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	locContext, err := getLocContextParams(r, i)
	if err != nil {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.FilterMonsterFormationIDsByAvailability(r.Context(), database.FilterMonsterFormationIDsByAvailabilityParams{
		Ids:            inputIDs,
		Availability:   avlParams.Availabilities,
		IsRepeatable:   avlParams.IsRepeatable,
		PreAirship:     avlParams.PreAirship,
		AvlType:        locContext.AvlType,
		LocContextID:   locContext.ID,
		LocContextType: locContext.Type,
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}
