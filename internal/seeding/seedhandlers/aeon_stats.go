package seedhandlers

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func CompleteAeonStats() error {
	for i := range sh.RawJSON.AeonStats {
		as := &sh.RawJSON.AeonStats[i]
		err := sh.AssignIDs(as.AVals)
		if err != nil {
			return err
		}

		err = sh.AssignIDs(as.BVals)
		if err != nil {
			return err
		}

		err = completeAeonXVals(as.XVals)
		if err != nil {
			return err
		}

		aeon := lookup.L.Aeons[as.Name]
		as.AeonID = aeon.ID
		aeon.BaseStats = *as
		lookup.L.Aeons[seedtypes.Key(aeon)] = aeon
		lookup.L.AeonsID[as.AeonID] = aeon
		sh.RawJSON.Aeons[i].BaseStats = *as
	}

	return nil
}

func completeAeonXVals(xVals []seedtypes.XVal) error {
	for i := range xVals {
		xVal := &xVals[i]
		err := sh.AssignIDs(xVal.BaseStats)
		if err != nil {
			return err
		}
	}

	return nil
}
