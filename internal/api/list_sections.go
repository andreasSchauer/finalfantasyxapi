package api

import (
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type SectionList struct {
	apiresource.ListParams
	Results []string `json:"results"`
}

func (l SectionList) GetListParams() apiresource.ListParams {
	return l.ListParams
}

func getSectionList(r *http.Request, sectionMap map[ts.SectionName]handlerinput.Subsection) (SectionList, error) {
	sectionNames := []string{}

	for section := range sectionMap {
		sectionNames = append(sectionNames, string(section))
	}
	slices.Sort(sectionNames)

	listParams, sections, err := handlerinput.CreatePaginatedList(r, sectionNames)
	if err != nil {
		return SectionList{}, err
	}

	list := SectionList{
		ListParams: listParams,
		Results:    sections,
	}

	return list, nil
}
