package api

import (
	"net/http"
	"slices"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type SectionList struct {
	ListParams
	Results []string `json:"results"`
}

func (l SectionList) GetListParams() ListParams {
	return l.ListParams
}

func getSectionList(srv *Server, r *http.Request, sectionMap map[ts.SectionName]Subsection) (SectionList, error) {
	sectionNames := []string{}

	for section := range sectionMap {
		sectionNames = append(sectionNames, string(section))
	}
	slices.Sort(sectionNames)

	listParams, sections, err := createPaginatedList(srv, r, sectionNames)
	if err != nil {
		return SectionList{}, err
	}

	list := SectionList{
		ListParams: listParams,
		Results:    sections,
	}

	return list, nil
}
