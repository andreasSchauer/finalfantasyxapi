package api

import (
	"database/sql"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func nullBool(r *http.Request, queryName typestrings.QueryParamName, queryLookup map[typestrings.QueryParamName]QueryParam) (sql.NullBool, error) {
	queryParam := queryLookup[queryName]

	b, err := queryParam.parseBool(r)
	if err != nil {
		return sql.NullBool{}, err
	}

	return h.GetNullBool(&b), nil
}

func nullInt32(r *http.Request, queryName typestrings.QueryParamName, queryLookup map[typestrings.QueryParamName]QueryParam) (sql.NullInt32, error) {
	queryParam := queryLookup[queryName]

	integer, err := queryParam.parseInt(r)
	if err != nil {
		return sql.NullInt32{}, err
	}

	integer32 := int32(integer)
	return h.GetNullInt32(&integer32), nil
}

func nullValue(r *http.Request, queryName typestrings.QueryParamName, queryLookup map[typestrings.QueryParamName]QueryParam) (sql.NullString, error) {
	queryParam := queryLookup[queryName]

	value, err := queryParam.parseValue(r)
	if err != nil {
		return sql.NullString{}, err
	}

	return h.GetNullString(&value), nil
}

func nullID[T seedtypes.Lookupable, R any, A APIResource](r *http.Request, i handlerInput[T, R, A], queryName typestrings.QueryParamName, queryLookup map[typestrings.QueryParamName]QueryParam) (sql.NullInt32, error) {
	queryParam := queryLookup[queryName]

	id, err := queryParam.parseID(r, i.objLookup)
	if err != nil {
		return sql.NullInt32{}, err
	}

	return h.GetNullInt32(&id), nil
}

func nullNameID[T seedtypes.Lookupable, R any, A APIResource](r *http.Request, i handlerInput[T, R, A], queryName typestrings.QueryParamName, queryLookup map[typestrings.QueryParamName]QueryParam) (sql.NullInt32, error) {
	queryParam := queryLookup[queryName]

	id, err := queryParam.parseNameID(r, i.resTypeSingle, i.objLookup)
	if err != nil {
		return sql.NullInt32{}, err
	}

	return h.GetNullInt32(&id), nil
}

func nullEnum[E, N any](r *http.Request, queryName typestrings.QueryParamName, et enums.EnumType[E, N], queryLookup map[typestrings.QueryParamName]QueryParam) (sql.NullString, error) {
	queryParam := queryLookup[queryName]

	enumRes, err := queryParam.parseEnum(r, et)
	if err != nil {
		return sql.NullString{}, err
	}

	return h.GetNullString(&enumRes.Name), nil
}
