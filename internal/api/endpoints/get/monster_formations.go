package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func MonsterFormation(r *http.Request, i handlerinput.Main[seedtypes.MonsterFormation, apitypes.MonsterFormation, apiresource.Unnamed], id int32) (apitypes.MonsterFormation, error) {
	formation, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.MonsterFormation{}, err
	}

	response := apitypes.MonsterFormation{
		ID:              formation.ID,
		Category:        formation.FormationData.Category,
		Availability:    formation.FormationData.Availability,
		IsForcedAmbush:  formation.FormationData.IsForcedAmbush,
		CanEscape:       formation.FormationData.CanEscape,
		Notes:           formation.FormationData.Notes,
		FormationLoot:   FormationLoot(formation),
		BossMusic:       ah.ConvertObjPtr(formation.FormationData.BossMusic, apiconvert.FormationBossSong),
		Monsters:        apilookups.E.Monsters.ToResAmtsFromNames(formation.Monsters),
		Areas:           ah.ConvertObjSlice(formation.EncounterAreas, apiconvert.EncounterArea),
		TriggerCommands: ah.ConvertObjSlice(formation.TriggerCommands, apiconvert.FormationTriggerCommand),
	}

	return response, nil
}
