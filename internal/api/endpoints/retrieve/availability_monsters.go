package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
)

func FilterAvlMonsters(r *http.Request, inputIDs []int32) ([]int32, error) {
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
		AvlType:        locContext.avlType,
		LocContextID:   locContext.id,
		LocContextType: locContext.locCtxType,
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}

func FilterAvlMonsterFormations(r *http.Request, inputIDs []int32) ([]int32, error) {
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
		AvlType:        locContext.avlType,
		LocContextID:   locContext.id,
		LocContextType: locContext.locCtxType,
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}
