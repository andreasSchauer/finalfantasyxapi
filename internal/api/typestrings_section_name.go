package api

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)


type SectionName NamedParam

const (
	snSimple					SectionName = "simple"
	snSections					SectionName = "sections"
	snParameters				SectionName = "parameters"
	snAbilities					SectionName = "abilities"
	snAreas						SectionName = "areas"
	snAutoAbilities				SectionName = "auto-abilities"
	snConnected					SectionName = "connected"
	snDefaultAbilities			SectionName = "default-abilities"
	snDefaultOverdrives			SectionName = "default-overdrives"
	snExpSgAbilities			SectionName = "exp-sg-abilities"
	snLearnableAbilities		SectionName = "learnable-abilities"
	snLearnableOverdrives		SectionName = "learnable-overdrives"
	snMixes						SectionName = "mixes"
	snMonsterFormations			SectionName = "monster-formations"
	snMonsters					SectionName = "monsters"
	snOverdriveAbilities		SectionName = "overdrive-abilities"
	snOverdrives				SectionName = "overdrives"
	snShops						SectionName = "shops"
	snSongs						SectionName = "songs"
	snStats						SectionName = "stats"
	snStdSgAbilities			SectionName = "std-sg-abilities"
	snSublocations				SectionName = "sublocations"
	snSubquests					SectionName = "subquests"
	snTreasures					SectionName = "treasures"
)

func snsToNamedParams(sns []SectionName) []NamedParam {
	if sns == nil {
		return nil
	}
	
	nps := make([]NamedParam, len(sns))

	for i, sn := range sns {
		nps[i] = NamedParam(sn)
	}

	return nps
}


func getSnPtr(s SectionName) *SectionName {
	return &s
}

func formatSectionNames(itemMap map[SectionName]Subsection) string {
	keys := []string{}

	for key := range itemMap {
		keyFormatted := fmt.Sprintf("'%s'", key)
		keys = append(keys, keyFormatted)
	}

	slices.SortStableFunc(keys, func(a, b string) int {
		return cmp.Compare(a, b)
	})

	return strings.Join(keys, ", ")
}