package sections

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/get"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
)

func CreateMonsterFormationSimple(_ *http.Request, id int32, _ handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.MonsterFormations
	formation, _ := lookup.GetResourceByID(id, i.ObjLookupID)

	formationSub := sectiontypes.MonsterFormationSimple{
		ID:             formation.ID,
		URL:            ah.CreateResourceURL(i.Endpoint, id),
		Category:       formation.FormationData.Category,
		IsForcedAmbush: formation.FormationData.IsForcedAmbush,
		FormationLoot:  get.GetFormationLoot(formation),
		Monsters:       ah.ConvertObjSlice(formation.Monsters, monsterAmountString),
		Areas:          locAreaStrings(formation.EncounterAreas),
	}

	return formationSub, nil
}
