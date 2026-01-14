package main

import (
	"fmt"
)

type AreaTreasuresList struct {
	ListParams
	Results []NamedAPIResource
}

func (l AreaTreasuresList) getListParams() ListParams {
	return l.ListParams
}

func (l AreaTreasuresList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

type AreaShopsList struct {
	ListParams
	Results []UnnamedAPIResource
}

func (l AreaShopsList) getListParams() ListParams {
	return l.ListParams
}

func (l AreaShopsList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

type AreaMonstersList struct {
	ListParams
	Results []UnnamedAPIResource
}

func (l AreaMonstersList) getListParams() ListParams {
	return l.ListParams
}

func (l AreaMonstersList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

type AreaConnectionsList struct {
	ListParams
	Results []LocationAPIResource
}

func (l AreaConnectionsList) getListParams() ListParams {
	return l.ListParams
}

func (l AreaConnectionsList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

type AreaFormationsList struct {
	ListParams
	Results []LocationAPIResource
}

func (l AreaFormationsList) getListParams() ListParams {
	return l.ListParams
}

func (l AreaFormationsList) getResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

func (cfg *Config) getAreaTreasuresMid(subsection string) (APIResourceList, error) {
	fmt.Printf("this should trigger /api/areas/{id}/%s\n", subsection)
	return AreaTreasuresList{}, nil
}

func (cfg *Config) getAreaShopsMid(subsection string) (APIResourceList, error) {
	fmt.Printf("this should trigger /api/areas/{id}/%s\n", subsection)
	return AreaShopsList{}, nil
}

func (cfg *Config) getAreaMonstersMid(subsection string) (APIResourceList, error) {
	fmt.Printf("this should trigger /api/areas/{id}/%s\n", subsection)
	return AreaMonstersList{}, nil
}

func (cfg *Config) getAreaConnectionsMid(subsection string) (APIResourceList, error) {
	fmt.Printf("this should trigger /api/areas/{id}/%s\n", subsection)
	return AreaConnectionsList{}, nil
}

func (cfg *Config) getAreaFormationsMid(subsection string) (APIResourceList, error) {
	fmt.Printf("this should trigger /api/areas/{id}/%s\n", subsection)
	return AreaConnectionsList{}, nil
}
