package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (i HandlerInput[T, R, A]) DBGetRelation(ctx context.Context, filterItem seedtypes.Lookupable, dbQuery DBQueryIntOne) (A, error) {
	var zeroType A

	dbID, err := dbQuery(ctx, filterItem.GetID())
	if err != nil {
		return zeroType, i.ErrorDB(filterItem, err)
	}

	res := i.idToResFunc(i, dbID)
	return res, nil
}

// DBGetRelations gets relationship resources of item. handlerInput = endpoint of fetched resources
func (i HandlerInput[T, R, A]) DBGetRelations(ctx context.Context, filterItem seedtypes.Lookupable, dbQuery DBQueryIntMany) ([]A, error) {
	dbIds, err := dbQuery(ctx, filterItem.GetID())
	if err != nil {
		return nil, i.ErrorDB(filterItem, err)
	}

	resources := i.ToAPIResources(dbIds)
	return resources, nil
}

func (i HandlerInput[T, R, A]) DBGetRelationPtr(ctx context.Context, item seedtypes.Lookupable, dbQuery DBQueryIntOne) (*A, error) {
	dbID, err := dbQuery(ctx, item.GetID())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get %s of %s", i.resTypeSingle, item), err)
	}

	res := i.idToResFunc(i, dbID)
	return &res, nil
}


// queries db for junctionPairs and converts them into a []Junction

func (i HandlerInput[T, R, A]) GetDBJunctions[row any](ctx context.Context, ids []int32, childResType ts.ResTypeSingle, dbQuery func(context.Context, []int32) ([]row, error), converter func(row) sectiontypes.Junction) ([]sectiontypes.Junction, error) {
	dbJunctions, err := dbQuery(ctx, ids)
	if err != nil {
		return nil, i.ErrorDBPairs(childResType, err)
	}

	junctions := []sectiontypes.Junction{}

	for _, dbJunction := range dbJunctions {
		junctions = append(junctions, converter(dbJunction))
	}

	return junctions, nil
}