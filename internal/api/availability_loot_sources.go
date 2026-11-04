package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func filterAvlShops(server *Server, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.Shops

	avlParams, err := i.queryAvlGetParams(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	sources, err := getShopSources(server, r, i)
	if err != nil {
		return nil, err
	}

	locContext, err := getLocContextParams(r, i)
	if err != nil {
		return nil, err
	}

	dbIDs, err := server.Cfg.DB.FilterShopIDsByAvailability(r.Context(), database.FilterShopIDsByAvailabilityParams{
		Ids:             inputIDs,
		Availability:    avlParams.availabilities,
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
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}

func filterAvlTreasures(server *Server, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.Treasures

	avlParams, err := i.queryAvlGetParams(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	dbIDs, err := server.Cfg.DB.FilterTreasureIDsByAvailability(r.Context(), database.FilterTreasureIDsByAvailabilityParams{
		Ids:          inputIDs,
		Availability: avlParams.availabilities,
		PreAirship:   avlParams.preAirship,
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}

func filterAvlQuests(server *Server, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.Quests

	avlParams, err := i.queryAvlRepGetParams(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	dbIDs, err := server.Cfg.DB.FilterQuestIDsByAvailability(r.Context(), database.FilterQuestIDsByAvailabilityParams{
		Ids:          inputIDs,
		Availability: avlParams.availabilities,
		IsRepeatable: avlParams.isRepeatable,
		PreAirship:   avlParams.preAirship,
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}

func filterAvlSidequests(server *Server, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.Sidequests

	avlParams, err := i.queryAvlRepGetParams(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	dbIDs, err := server.Cfg.DB.FilterSidequestIDsByAvailability(r.Context(), database.FilterSidequestIDsByAvailabilityParams{
		Ids:          inputIDs,
		Availability: avlParams.availabilities,
		PreAirship:   avlParams.preAirship,
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}

func filterAvlSubquests(server *Server, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.Subquests

	avlParams, err := i.queryAvlRepGetParams(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	dbIDs, err := server.Cfg.DB.FilterSubquestIDsByAvailability(r.Context(), database.FilterSubquestIDsByAvailabilityParams{
		Ids:          inputIDs,
		Availability: avlParams.availabilities,
		IsRepeatable: avlParams.isRepeatable,
		PreAirship:   avlParams.preAirship,
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}
