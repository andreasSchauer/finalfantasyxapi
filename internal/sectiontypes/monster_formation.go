package sectiontypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"

type MonsterFormationSimple struct {
	ID             int32  `json:"id"`
	URL            string `json:"url"`
	Category       string `json:"category"`
	IsForcedAmbush bool   `json:"is_forced_ambush"`
	apitypes.FormationLoot
	Monsters []string `json:"monsters"`
	Areas    []string `json:"areas"`
}

func (m MonsterFormationSimple) GetURL() string {
	return m.URL
}
