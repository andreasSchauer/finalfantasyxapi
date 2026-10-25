package api

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type SectionName NamedParam

const (
	SnSimple              SectionName = "simple"
	SnSections            SectionName = "sections"
	SnParameters          SectionName = "parameters"
	SnBody                SectionName = "body"
	SnDocs                SectionName = "docs"
	SnSamples             SectionName = "samples"
	SnTemplate            SectionName = "template"
	SnAbilities           SectionName = "abilities"
	SnAreas               SectionName = "areas"
	SnAutoAbilities       SectionName = "auto-abilities"
	SnConnected           SectionName = "connected"
	SnDefaultAbilities    SectionName = "default-abilities"
	SnDefaultOverdrives   SectionName = "default-overdrives"
	SnExpSgAbilities      SectionName = "exp-sg-abilities"
	SnLearnableAbilities  SectionName = "learnable-abilities"
	SnLearnableOverdrives SectionName = "learnable-overdrives"
	SnMixes               SectionName = "mixes"
	SnMonsterFormations   SectionName = "monster-formations"
	SnMonsters            SectionName = "monsters"
	SnOverdriveAbilities  SectionName = "overdrive-abilities"
	SnOverdrives          SectionName = "overdrives"
	SnShops               SectionName = "shops"
	SnSongs               SectionName = "songs"
	SnStats               SectionName = "stats"
	SnStdSgAbilities      SectionName = "std-sg-abilities"
	SnSublocations        SectionName = "sublocations"
	SnSubquests           SectionName = "subquests"
	SnTreasures           SectionName = "treasures"
)

func SnsToNamedParams(sns []SectionName) []NamedParam {
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
