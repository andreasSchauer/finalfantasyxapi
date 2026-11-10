package seeding

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func completeAeonStats() error {
	for i := range rawJSON.aeonStats {
		as := &rawJSON.aeonStats[i]
		err := assignIDs(as.AVals)
		if err != nil {
			return err
		}

		err = assignIDs(as.BVals)
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
		rawJSON.aeons[i].BaseStats = *as
	}

	return nil
}

func completeAeonXVals(xVals []seedtypes.XVal) error {
	for i := range xVals {
		xVal := &xVals[i]
		err := assignIDs(xVal.BaseStats)
		if err != nil {
			return err
		}
	}

	return nil
}
