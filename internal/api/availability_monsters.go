package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func filterAvlMonsters(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.Monsters

	avlParams, err := i.queryAvlRepGetParams(r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	locContext, err := getLocContextParams(r, i)
	if err != nil {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.FilterMonsterIDsByAvailability(r.Context(), database.FilterMonsterIDsByAvailabilityParams{
		Ids:            inputIDs,
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship:     avlParams.preAirship,
		AvlType:        locContext.AvlType,
		LocContextID:   locContext.ID,
		LocContextType: locContext.Type,
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}

func filterAvlMonsterFormations(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.MonsterFormations

	avlParams, err := i.queryAvlRepGetParams(r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	locContext, err := getLocContextParams(r, i)
	if err != nil {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.FilterMonsterFormationIDsByAvailability(r.Context(), database.FilterMonsterFormationIDsByAvailabilityParams{
		Ids:            inputIDs,
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship:     avlParams.preAirship,
		AvlType:        locContext.AvlType,
		LocContextID:   locContext.ID,
		LocContextType: locContext.Type,
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}
