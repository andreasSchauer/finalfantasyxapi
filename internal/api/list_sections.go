package api

import (
	"net/http"
	"slices"
)

type SectionList struct {
	ListParams
	Results []string `json:"results"`
}

func (l SectionList) getListParams() ListParams {
	return l.ListParams
}

func getSectionList(cfg *Config, r *http.Request, sectionMap map[SectionName]Subsection) (SectionList, error) {
	sectionNames := []string{}

	for section := range sectionMap {
		sectionNames = append(sectionNames, string(section))
	}
	slices.Sort(sectionNames)

	listParams, sections, err := createPaginatedList(cfg, r, sectionNames)
	if err != nil {
		return SectionList{}, err
	}

	list := SectionList{
		ListParams: listParams,
		Results:    sections,
	}

	return list, nil
}
