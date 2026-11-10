package handlerinput

import (
	"database/sql"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/queryparam"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func NullBool(r *http.Request, queryName ts.QueryParamName, queryLookup map[ts.QueryParamName]queryparam.QueryParam) (sql.NullBool, error) {
	queryParam := queryLookup[queryName]

	b, err := queryParam.ParseBool(r)
	if err != nil {
		return sql.NullBool{}, err
	}

	return h.GetNullBool(&b), nil
}

func NullInt32(r *http.Request, queryName ts.QueryParamName, queryLookup map[ts.QueryParamName]queryparam.QueryParam) (sql.NullInt32, error) {
	queryParam := queryLookup[queryName]

	integer, err := queryParam.ParseInt(r)
	if err != nil {
		return sql.NullInt32{}, err
	}

	integer32 := int32(integer)
	return h.GetNullInt32(&integer32), nil
}

func NullValue(r *http.Request, queryName ts.QueryParamName, queryLookup map[ts.QueryParamName]queryparam.QueryParam) (sql.NullString, error) {
	queryParam := queryLookup[queryName]

	value, err := queryParam.ParseValue(r)
	if err != nil {
		return sql.NullString{}, err
	}

	return h.GetNullString(&value), nil
}

func NullID[T seedtypes.Lookupable, R any, A apiresource.APIResource](r *http.Request, i Main[T, R, A], queryName ts.QueryParamName, queryLookup map[ts.QueryParamName]queryparam.QueryParam) (sql.NullInt32, error) {
	queryParam := queryLookup[queryName]

	id, err := queryParam.ParseID(r, i.ObjLookup)
	if err != nil {
		return sql.NullInt32{}, err
	}

	return h.GetNullInt32(&id), nil
}

func NullNameID[T seedtypes.Lookupable, R any, A apiresource.APIResource](r *http.Request, i Main[T, R, A], queryName ts.QueryParamName, queryLookup map[ts.QueryParamName]queryparam.QueryParam) (sql.NullInt32, error) {
	queryParam := queryLookup[queryName]

	id, err := queryParam.ParseNameID(r, i.ResTypeSingle, i.ObjLookup)
	if err != nil {
		return sql.NullInt32{}, err
	}

	return h.GetNullInt32(&id), nil
}

func NullEnum[E, N any](r *http.Request, queryName ts.QueryParamName, et enums.EnumType[E, N], queryLookup map[ts.QueryParamName]queryparam.QueryParam) (sql.NullString, error) {
	queryParam := queryLookup[queryName]

	enumRes, err := queryParam.ParseEnum(r, et)
	if err != nil {
		return sql.NullString{}, err
	}

	return h.GetNullString(&enumRes.Name), nil
}
