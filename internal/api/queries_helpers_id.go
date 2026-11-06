package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

// ParseID validates id-queryParam and checks emptiness.
func (q QueryParam) ParseID[F seedtypes.Lookupable](r *http.Request, fLookup map[string]F) (int32, error) {
	query, err := q.CheckEmpty(r)
	if err != nil {
		return 0, err
	}

	id, err := q.ParseIDVal(query, len(fLookup))
	if err != nil {
		return 0, err
	}

	return id, nil
}

// ParseIDPtr validates id-queryParam and checks emptiness. accepts "none" as input.
func (q QueryParam) ParseIDPtr[F seedtypes.Lookupable](r *http.Request, fLookup map[string]F) (*int32, error) {
	query, err := q.CheckEmpty(r)
	if err != nil {
		return nil, err
	}

	err = checkNoneQuery(query)
	if err != nil {
		return nil, nil
	}

	id, err := q.ParseIDVal(query, len(fLookup))
	if err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIDVal checks if query ID is valid and within range. if it's invalid, it will return an httpError. if you want to soft-check the ID and do name checks afterwards, use checkQueryID()
func (q QueryParam) ParseIDVal(idStr string, maxID int) (int32, error) {
	id, err := q.CheckID(idStr, maxID)
	if errors.Is(err, errs.ErrNotAnID) {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid id '%s' used for parameter '%s'", idStr, q.Name), err)
	}
	if err != nil {
		return 0, err
	}

	return int32(id), nil
}

// CheckID checks if query ID is valid and within range. if it's invalid, it will return errNotAnID, providing the possibility to do name-based checks. for a hard-check with errors, use parseQueryID()
func (q QueryParam) CheckID(idStr string, maxID int) (int32, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errs.ErrNotAnID
	}

	if id > maxID || id <= 0 {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided id '%d' used for parameter '%s' is out of range. max id: %d", id, q.Name, maxID), err)
	}

	return int32(id), nil
}
