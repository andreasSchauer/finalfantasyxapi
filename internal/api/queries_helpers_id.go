package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// validates id-queryParam and checks emptiness.
func parseIdQuery(r *http.Request, queryParam QueryParam, maxID int) (int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return 0, err
	}

	id, err := parseQueryID(query, queryParam, maxID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// validates id-queryParam and checks emptiness. accepts "none" as input.
func parseIdQueryNul(r *http.Request, queryParam QueryParam, maxID int) (*int32, error) {
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	err = checkNoneQuery(query)
	if err != nil {
		return nil, nil
	}

	id, err := parseQueryID(query, queryParam, maxID)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

// checks if query ID is valid and within range. if it's invalid, it will return an httpError. if you want to soft-check the ID and do name checks afterwards, use checkQueryID()
func parseQueryID(idStr string, queryParam QueryParam, maxID int) (int32, error) {
	id, err := checkQueryID(idStr, queryParam, maxID)
	if errors.Is(err, errNotAnID) {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid id '%s' used for parameter '%s'.", idStr, queryParam.Name), err)
	}
	if err != nil {
		return 0, err
	}

	return int32(id), nil
}

// checks if query ID is valid and within range. if it's invalid, it will return errNotAnID, providing the possibility to do name-based checks. for a hard-check with errors, use parseQueryID()
func checkQueryID(idStr string, queryParam QueryParam, maxID int) (int32, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errNotAnID
	}

	if id > maxID || id <= 0 {
		return 0, newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided id '%d' used for parameter '%s' is out of range. max id: %d.", id, queryParam.Name, maxID), err)
	}

	return int32(id), nil
}
