package api

import (
	"errors"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type filteredResList[T HasAPIResource] struct {
	resources []T
	err       error
}

func frl[T HasAPIResource](res []T, err error) filteredResList[T] {
	return filteredResList[T]{
		resources: res,
		err:       err,
	}
}

// a query filter that can't really be generalized. this one simply checks, if it's empty and then calls the wrapperFn
func basicQueryWrapper[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, wrapperFn func(*Config, *http.Request, string, QueryType) ([]int32, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return inputRes, nil
	}

	dbIDs, err := wrapperFn(cfg, r, query, queryParam)
	if errors.Is(err, errQueryRedirect) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}

// query uses an id of another resource type to filter resources
func nameOrIdQuery[T, P h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName, pResType string, pLookup map[string]P, dbQuery DbQueryIntMany) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	id, err := parseNameOrIdQuery(r, queryParam, pResType, pLookup)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), id)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

// query uses a list of names or ids as database input to filter for resources. alternatively, "none" can be used as input instead.
func nameIdListQueryNul[T, P h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName, pResType string, pLookup map[string]P, dbQuery DbQueryIntList) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	queryIDs, err := parseNameIdListQueryNullable(r, queryParam, pResType, pLookup)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), queryIDs)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

// query uses an id of another resource type to filter resources
func idQuery[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, maxID int, dbQuery DbQueryIntMany) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	id, err := parseIDOnlyQuery(r, queryParam, maxID)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), id)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

// like idOnlyQuery, but with more specialized logic in between (wrapperFn)
func idQueryWrapper[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, maxID int, wrapperFn func(*Config, *http.Request, int32) ([]A, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	id, err := parseIDOnlyQuery(r, queryParam, maxID)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	resources, err := wrapperFn(cfg, r, id)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

func idQueryNul[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, maxID int, dbQuery DbQueryNullIntMany) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	idPtr, err := parseIDOnlyQueryNul(r, queryParam, maxID)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIds, err := dbQuery(r.Context(), h.GetNullInt32(idPtr))
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIds)
	return resources, nil
}

// query uses a list of ids as database input to filter for resources
func idListQuery[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, maxID int, dbQuery DbQueryIntList) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	queryIDs, err := parseIdListQueryNoDupes(r, queryParam, maxID)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), queryIDs)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

// like idListQuery, but with more specialized logic in between (wrapperFn)
func idListQueryWrapper[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, maxID int, wrapperFn func(*Config, *http.Request, []int32) ([]int32, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	queryIDs, err := parseIdListQueryNoDupes(r, queryParam, maxID)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := wrapperFn(cfg, r, queryIDs)
	if err != nil {
		return nil, err
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

func boolQueryWrapper[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, wrapperFn func(*Config, *http.Request, bool) ([]int32, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := wrapperFn(cfg, r, b)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

// db query searches for resources with matching boolean db column value
func boolQuery[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, dbQuery DbQueryBool) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), b)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

// db query accumulates all resources that fulfill a certain condition (mostly if it has resources of a specific type). a false boolean flips these results
func boolQuery2[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, dbQuery DbQueryNoInput) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context())
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	if !b {
		resources = removeResources(inputRes, resources)
	}

	return resources, nil
}

// db query accumulates all resources that fulfill a certain condition (mostly if it has resources of a specific type). a true boolean flips these results
func boolQuery3[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, dbQuery DbQueryNoInput) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	b, err := parseBooleanQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context())
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	if b {
		resources = removeResources(inputRes, resources)
	}

	return resources, nil
}

// query uses an enum type (id or string possible) that needs to be checked for validity and then returns all resources matching that type
func enumQuery[T h.HasID, R any, A APIResource, L APIResourceList, E, N any](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], et EnumType[E, N], inputRes []A, queryName string, dbQuery DbQueryEnum[E]) ([]A, error) {
	queryParam := i.queryLookup[queryName]
	enum, err := parseEnumQuery(r, i.endpoint, queryParam, et)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	typedStr := et.convFunc(enum.Name)

	dbIDs, err := dbQuery(r.Context(), typedStr)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}

func enumListQuery[T h.HasID, R any, A APIResource, L APIResourceList, E, N any](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], et EnumType[E, N], inputRes []A, queryName string, dbQuery DbQueryEnumList[E]) ([]A, error) {
	queryParam := i.queryLookup[queryName]
	enums, err := parseEnumListQuery(r, i.endpoint, queryParam, et)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), enums)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}

// like enum query, but with more specialized logic in between (wrapperFn). For example, if types are grouped together (ctbIconType)
func enumQueryWrapper[T h.HasID, R any, A APIResource, L APIResourceList, E, N any](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], et EnumType[E, N], inputRes []A, queryName string, wrapperFn func(*Config, *http.Request, E) ([]int32, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]
	enum, err := parseEnumQuery(r, i.endpoint, queryParam, et)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	typedStr := et.convFunc(enum.Name)
	dbIDs, err := wrapperFn(cfg, r, typedStr)
	if err != nil {
		return nil, err
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}

// query uses an integer value as input.
func intQuery[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, dbQuery DbQueryIntMany) ([]A, error) {
	queryParam := i.queryLookup[queryName]
	integer, err := parseIntQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), int32(integer))
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

// query uses an integer value as input.
func intQueryWrapper[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, wrapperFn func(*Config, *http.Request, int32) ([]int32, error)) ([]A, error) {
	queryParam := i.queryLookup[queryName]
	integer, err := parseIntQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := wrapperFn(cfg, r, int32(integer))
	if errors.Is(err, errQueryRedirect) {
		return inputRes, nil
	}
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

func intListQuery[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], inputRes []A, queryName string, dbQuery DbQueryIntList) ([]A, error) {
	queryParam := i.queryLookup[queryName]

	ints, err := parseIntListQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return inputRes, nil
	}
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), ints)
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParam, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}
