package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (i handlerInput[T, R, A]) dbGetRelation(server *Server, ctx context.Context, filterItem seedtypes.Lookupable, dbQuery DBQueryIntOne) (A, error) {
	var zeroType A

	dbID, err := dbQuery(ctx, filterItem.GetID())
	if err != nil {
		return zeroType, newHTTPErrorDB(i.resTypePlural, filterItem, err)
	}

	res := i.idToResFunc(server, i, dbID)
	return res, nil
}

// get relationship resources of item. handlerInput = endpoint of fetched resources
func (i handlerInput[T, R, A]) dbGetRelations(server *Server, ctx context.Context, filterItem seedtypes.Lookupable, dbQuery DBQueryIntMany) ([]A, error) {
	dbIds, err := dbQuery(ctx, filterItem.GetID())
	if err != nil {
		return nil, newHTTPErrorDB(i.resTypePlural, filterItem, err)
	}

	resources := i.toResources(server, dbIds)
	return resources, nil
}

func (i handlerInput[T, R, A]) dbGetRelationPtr(server *Server, ctx context.Context, item seedtypes.Lookupable, dbQuery DBQueryIntOne) (*A, error) {
	dbID, err := dbQuery(ctx, item.GetID())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get %s of %s", i.resTypeSingle, item), err)
	}

	res := i.idToResFunc(server, i, dbID)
	return &res, nil
}
