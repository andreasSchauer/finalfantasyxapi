package main

import (
	"fmt"
)


type MonsterAbilitiesList struct {
	ListParams
	Results []NamedAPIResource
}

func (l MonsterAbilitiesList) getListParams() ListParams {
	return l.ListParams
}

func (l MonsterAbilitiesList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

func (cfg *Config) getMonsterAbilitiesMid(subsection string) (IsAPIResourceList, error) {
	fmt.Printf("this should trigger /api/monsters/{id}/%s\n", subsection)
	return MonsterAbilitiesList{}, nil
}