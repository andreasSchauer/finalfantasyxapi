package handlerinput

import (
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

type SectionList struct {
	apiresource.ListParams
	Results []string `json:"results"`
}

func (l SectionList) GetListParams() apiresource.ListParams {
	return l.ListParams
}

func GetSectionList(r *http.Request, sectionMap map[ts.SectionName]Subsection) (SectionList, error) {
	sectionNames := []string{}

	for section := range sectionMap {
		sectionNames = append(sectionNames, string(section))
	}
	slices.Sort(sectionNames)

	listParams, sections, err := createPaginatedList(r, sectionNames)
	if err != nil {
		return SectionList{}, err
	}

	list := SectionList{
		ListParams: listParams,
		Results:    sections,
	}

	return list, nil
}
