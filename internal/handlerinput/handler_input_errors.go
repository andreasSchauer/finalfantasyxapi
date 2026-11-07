package handlerinput

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (i HandlerInput[T, R, A]) ErrorDBFilter(queryParam queryparam.QueryParam, err error) errs.HTTPError {
	return errs.NewHTTPErrorDBFilter(i.resTypePlural, queryParam.Name, err)
}

func (i HandlerInput[T, R, A]) ErrorDBOne(parentItem seedtypes.Lookupable, err error) errs.HTTPError {
	return errs.NewHTTPErrorDBOne(i.resTypeSingle, parentItem, err)
}

func (i HandlerInput[T, R, A]) ErrorDB(parentItem seedtypes.Lookupable, err error) errs.HTTPError {
	return errs.NewHTTPErrorDB(i.resTypePlural, parentItem, err)
}

func (i HandlerInput[T, R, A]) ErrorDBPairs(parentResType ts.ResTypeSingle, err error) errs.HTTPError {
	return errs.NewHTTPErrorDBPairs(i.resTypeSingle, parentResType, err)
}

func (i HandlerInput[T, R, A]) ErrorAvl(err error) errs.HTTPError {
	return errs.NewHTTPErrorAvailability(i.resTypePlural, err)
}
