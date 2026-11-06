package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func getAlterations(alts []apitypes.Alt) apitypes.Alterations {
	alterations := apitypes.Alterations{}

	for _, alt := range alts {
		switch alt.AlterationType {
		case database.AlterationTypeChange:
			change := apitypes.AltStateChange{
				Distance:    alt.Distance,
				BaseStats:   alt.BaseStats,
				ElemResists: alt.ElemResists,
			}
			alterations.Change = &change

		case database.AlterationTypeGain:
			gain := apitypes.AltStateGain{
				Properties:       alt.Properties,
				AutoAbilities:    alt.AutoAbilities,
				StatusImmunities: alt.StatusImmunities,
				StatusResists:    alt.StatusResists,
				Status:           alt.AddedStatus,
			}
			alterations.Gain = &gain

		case database.AlterationTypeLoss:
			loss := apitypes.AltStateLoss{
				Properties:       alt.Properties,
				AutoAbilities:    alt.AutoAbilities,
				StatusImmunities: alt.StatusImmunities,
				Status:           alt.RemovedStatus,
			}
			alterations.Loss = &loss
		}
	}

	return alterations
}
