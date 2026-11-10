package handlerinput

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/queryparam"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func (i Main[T, R, A]) ErrorDBFilter(queryParam queryparam.QueryParam, err error) errs.HTTPError {
	return errs.NewHTTPErrorDBFilter(i.ResTypePlural, queryParam.Name, err)
}

func (i Main[T, R, A]) ErrorDBOne(parentItem seedtypes.Lookupable, err error) errs.HTTPError {
	return errs.NewHTTPErrorDBOne(i.ResTypeSingle, parentItem, err)
}

func (i Main[T, R, A]) ErrorDB(parentItem seedtypes.Lookupable, err error) errs.HTTPError {
	return errs.NewHTTPErrorDB(i.ResTypePlural, parentItem, err)
}

func (i Main[T, R, A]) ErrorDBPairs(parentResType ts.ResTypeSingle, err error) errs.HTTPError {
	return errs.NewHTTPErrorDBPairs(i.ResTypeSingle, parentResType, err)
}

func (i Main[T, R, A]) ErrorAvl(err error) errs.HTTPError {
	return errs.NewHTTPErrorAvailability(i.ResTypePlural, err)
}
