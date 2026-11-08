package handlerinput

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
)

func (i Main[T, R, A]) ToSimpleResourceList(r *http.Request, id int32, sectionName string, subsection Subsection) (apiresource.SimpleResourceList, error) {
	dbIDs, err := subsection.DBQuery(r.Context(), id)
	if err != nil {
		return apiresource.SimpleResourceList{}, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %s of %s with id '%d'", sectionName, i.ResTypeSingle, id), err)
	}

	subResList, err := CreateSimpleResourceList(r, dbIDs, subsection)
	if err != nil {
		return apiresource.SimpleResourceList{}, err
	}

	subResList.ParentResource = i.IDToResFunc(i, id)

	return subResList, nil
}

func CreateSimpleResourceList(r *http.Request, dbIDs []int32, subsection Subsection) (apiresource.SimpleResourceList, error) {
	results, err := createSimpleResources(r, dbIDs, subsection)
	if err != nil {
		return apiresource.SimpleResourceList{}, err
	}

	listParams, shownResults, err := createPaginatedList(r, results)
	if err != nil {
		return apiresource.SimpleResourceList{}, err
	}

	subResList := apiresource.SimpleResourceList{
		ListParams: listParams,
		Results:    shownResults,
	}

	return subResList, nil
}

func createSimpleResources(r *http.Request, dbIDs []int32, subsection Subsection) ([]apiresource.SimpleResource, error) {
	subs := []apiresource.SimpleResource{}

	if subsection.RelationsFn != nil {
		var err error
		subsection.Relations, err = subsection.RelationsFn(r, dbIDs)
		if err != nil {
			return nil, err
		}
	}

	for _, id := range dbIDs {
		subRes, err := subsection.CreateSubFn(r, id, subsection)
		if err != nil {
			return nil, err
		}

		subs = append(subs, subRes)
	}

	return subs, nil
}
