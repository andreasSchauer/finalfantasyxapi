package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconv"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func GetMonsterFormation(r *http.Request, i handlerinput.Main[seedtypes.MonsterFormation, apitypes.MonsterFormation, apiresource.Unnamed], id int32) (apitypes.MonsterFormation, error) {
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
		FormationLoot:   GetFormationLoot(formation),
		BossMusic:       ah.ConvertObjPtr(formation.FormationData.BossMusic, apiconv.ConvertFormationBossSong),
		Monsters:        apilookups.E.Monsters.ToResAmtsFromNames(formation.Monsters),
		Areas:           ah.ConvertObjSlice(formation.EncounterAreas, apiconv.ConvertEncounterArea),
		TriggerCommands: ah.ConvertObjSlice(formation.TriggerCommands, apiconv.ConvertFormationTriggerCommand),
	}

	return response, nil
}
