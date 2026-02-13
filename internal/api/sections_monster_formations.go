package api

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

func createMonsterFormationSub(cfg *Config, _ *http.Request, id int32) (SubResource, error) {
	i := cfg.e.monsterFormations
	formation, _ := seeding.GetResourceByID(id, i.objLookupID)

	formationSub := MonsterFormationSub{
		ID:             formation.ID,
		URL:            createResourceURL(cfg, i.endpoint, id),
		Category:       formation.FormationData.Category,
		IsForcedAmbush: formation.FormationData.IsForcedAmbush,
		Monsters:       convertObjSlice(cfg, formation.Monsters, monsterAmountString),
		Areas:          locAreaStrings(cfg, formation.EncounterAreas),
	}

	return formationSub, nil
}
