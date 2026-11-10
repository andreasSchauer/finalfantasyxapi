package sections

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/get"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
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
