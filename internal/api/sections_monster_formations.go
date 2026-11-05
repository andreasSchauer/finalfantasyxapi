package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

type MonsterFormationSimple struct {
	ID             int32  `json:"id"`
	URL            string `json:"url"`
	Category       string `json:"category"`
	IsForcedAmbush bool   `json:"is_forced_ambush"`
	FormationLoot
	Monsters []string `json:"monsters"`
	Areas    []string `json:"areas"`
}

func (m MonsterFormationSimple) GetURL() string {
	return m.URL
}

func createMonsterFormationSimple(_ *http.Request, id int32, _ Subsection) (apiresource.SimpleResource, error) {
	i := E.MonsterFormations
	formation, _ := lookup.GetResourceByID(id, i.objLookupID)

	formationSub := MonsterFormationSimple{
		ID:             formation.ID,
		URL:            ah.CreateResourceURL(i.Endpoint, id),
		Category:       formation.FormationData.Category,
		IsForcedAmbush: formation.FormationData.IsForcedAmbush,
		FormationLoot:  getFormationLoot(formation),
		Monsters:       ah.ConvertObjSlice(formation.Monsters, monsterAmountString),
		Areas:          locAreaStrings(formation.EncounterAreas),
	}

	return formationSub, nil
}
