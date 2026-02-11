package main

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type MonsterFormationSub struct {
	ID             int32    `json:"id"`
	URL            string   `json:"url"`
	Category       string   `json:"category"`
	IsForcedAmbush bool     `json:"is_forced_ambush"`
	Monsters       []string `json:"monsters"`
	Areas          []string `json:"areas"`
}

func (m MonsterFormationSub) GetURL() string {
	return m.URL
}

func handleMonsterFormationsSection(cfg *Config, _ *http.Request, dbIDs []int32) ([]SubResource, error) {
	i := cfg.e.monsterFormations
	formations := []MonsterFormationSub{}

	for _, formationID := range dbIDs {
		formation, _ := seeding.GetResourceByID(formationID, i.objLookupID)

		formationSub := MonsterFormationSub{
			ID:             formation.ID,
			URL:            createResourceURL(cfg, i.endpoint, formationID),
			Category:       formation.FormationData.Category,
			IsForcedAmbush: formation.FormationData.IsForcedAmbush,
			Monsters:       convertObjSlice(cfg, formation.Monsters, monsterAmountString),
			Areas:          locAreaStrings(cfg, formation.EncounterAreas),
		}

		formations = append(formations, formationSub)
	}

	return toSubResourceSlice(formations), nil
}
