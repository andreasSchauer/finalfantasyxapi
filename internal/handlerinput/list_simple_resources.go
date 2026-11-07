package handlerinput

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
)

func (i HandlerInput[T, R, A]) ToSimpleResourceList(r *http.Request, id int32, sectionName string, subsection Subsection) (apiresource.SimpleResourceList, error) {
	dbIDs, err := subsection.dbQuery(r.Context(), id)
	if err != nil {
		return apiresource.SimpleResourceList{}, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %s of %s with id '%d'", sectionName, i.resTypeSingle, id), err)
	}

	results, err := CreateSimpleResources(r, dbIDs, subsection)
	if err != nil {
		return apiresource.SimpleResourceList{}, err
	}

	listParams, shownResults, err := CreatePaginatedList(r, results)
	if err != nil {
		return apiresource.SimpleResourceList{}, err
	}

	subResList := apiresource.SimpleResourceList{
		ListParams:     listParams,
		ParentResource: i.idToResFunc(i, id),
		Results:        shownResults,
	}

	return subResList, nil
}

func CreateSimpleResources(r *http.Request, dbIDs []int32, subsection Subsection) ([]apiresource.SimpleResource, error) {
	subs := []apiresource.SimpleResource{}

	if subsection.relationsFn != nil {
		var err error
		subsection.relations, err = subsection.relationsFn(r, dbIDs)
		if err != nil {
			return nil, err
		}
	}

	for _, id := range dbIDs {
		subRes, err := subsection.createSubFn(r, id, subsection)
		if err != nil {
			return nil, err
		}

		subs = append(subs, subRes)
	}

	return subs, nil
}
