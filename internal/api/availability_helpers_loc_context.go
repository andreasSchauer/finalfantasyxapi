package api

import (
	"database/sql"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type locContextParams struct {
	AvlType string
	ID      sql.NullInt32
	Type    sql.NullString
}

func getLocContextParams[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) (locContextParams, error) {
	avlType := AvlTypeSelf
	var locContextID *int32
	var locContextType string

	locID, err := getQueryIdPtr(r, cfg.e.locations, "location", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		locContextID = locID
		locContextType = string(ViewSourceTypeLocation)
	}

	subLocID, err := getQueryIdPtr(r, cfg.e.sublocations, "sublocation", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		locContextID = subLocID
		locContextType = string(ViewSourceTypeSublocation)
	}

	areaID, err := getQueryIdPtr(r, cfg.e.areas, "area", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext2
		locContextID = areaID
		locContextType = string(ViewSourceTypeArea)
	}

	var locCtxTypePtr *string

	if locContextType != "" {
		locCtxTypePtr = &locContextType
	}

	params := locContextParams{
		AvlType: string(avlType),
		ID:      h.GetNullInt32(locContextID),
		Type:    h.GetNullString(locCtxTypePtr),
	}

	return params, nil
}
