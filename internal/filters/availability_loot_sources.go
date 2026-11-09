package filters

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
)

func FilterAvlShops(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := apilookups.E.Shops

	avlParams, err := i.QueryAvlGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	sources, err := getShopSources(r, i)
	if err != nil {
		return nil, err
	}

	locContext, err := getLocContextParams(r, i)
	if err != nil {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.FilterShopIDsByAvailability(r.Context(), database.FilterShopIDsByAvailabilityParams{
		Ids:             inputIDs,
		Availability:    avlParams.Availabilities,
		AvlType:         sources.AvlType,
		LocContextID:    locContext.ID,
		LocContextType:  locContext.Type,
		RequiredSources: sources.RequiredSources,
		ExcludedSources: sources.ExcludedSources,
		AutoAbilityID:   sources.AutoAbilityID,
		CharacterID:     sources.CharacterID,
		EmptySlots:      sources.EmptySlots,
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}

func FilterAvlTreasures(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := apilookups.E.Treasures

	avlParams, err := i.QueryAvlGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	dbIDs, err := config.Cfg.DB.FilterTreasureIDsByAvailability(r.Context(), database.FilterTreasureIDsByAvailabilityParams{
		Ids:          inputIDs,
		Availability: avlParams.Availabilities,
		PreAirship:   avlParams.PreAirship,
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}

func FilterAvlQuests(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := apilookups.E.Quests

	avlParams, err := i.QueryAvlRepGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	dbIDs, err := config.Cfg.DB.FilterQuestIDsByAvailability(r.Context(), database.FilterQuestIDsByAvailabilityParams{
		Ids:          inputIDs,
		Availability: avlParams.Availabilities,
		IsRepeatable: avlParams.IsRepeatable,
		PreAirship:   avlParams.PreAirship,
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}

func FilterAvlSidequests(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := apilookups.E.Sidequests

	avlParams, err := i.QueryAvlRepGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	dbIDs, err := config.Cfg.DB.FilterSidequestIDsByAvailability(r.Context(), database.FilterSidequestIDsByAvailabilityParams{
		Ids:          inputIDs,
		Availability: avlParams.Availabilities,
		PreAirship:   avlParams.PreAirship,
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}

func FilterAvlSubquests(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := apilookups.E.Subquests

	avlParams, err := i.QueryAvlRepGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	dbIDs, err := config.Cfg.DB.FilterSubquestIDsByAvailability(r.Context(), database.FilterSubquestIDsByAvailabilityParams{
		Ids:          inputIDs,
		Availability: avlParams.Availabilities,
		IsRepeatable: avlParams.IsRepeatable,
		PreAirship:   avlParams.PreAirship,
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}
