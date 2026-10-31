package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getResDBItemOne[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, ctx context.Context, i handlerInput[T, R, A, L], filterItem seeding.Lookupable, dbQuery DBQueryIntOne) (A, error) {
	var zeroType A

	dbID, err := dbQuery(ctx, filterItem.GetID())
	if err != nil {
		return zeroType, newHTTPErrorDB(i.resTypePlural, filterItem, err)
	}

	res := i.idToResFunc(cfg, i, dbID)
	return res, nil
}

// get relationship resources of item. handlerInput = endpoint of fetched resources
func getResourcesDBItem[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, ctx context.Context, i handlerInput[T, R, A, L], filterItem seeding.Lookupable, dbQuery DBQueryIntMany) ([]A, error) {
	dbIds, err := dbQuery(ctx, filterItem.GetID())
	if err != nil {
		return nil, newHTTPErrorDB(i.resTypePlural, filterItem, err)
	}

	resources := idsToAPIResources(cfg, i, dbIds)
	return resources, nil
}

func getResPtrDB[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, ctx context.Context, i handlerInput[T, R, A, L], item seeding.Lookupable, dbQuery DBQueryIntOne) (*A, error) {
	dbID, err := dbQuery(ctx, item.GetID())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get %s of %s.", i.resTypeSingle, item), err)
	}

	res := i.idToResFunc(cfg, i, dbID)
	return &res, nil
}
