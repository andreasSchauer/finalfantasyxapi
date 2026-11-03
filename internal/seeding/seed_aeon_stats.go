package seeding

import "github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"

func (l *Lookup) completeAeonStats() error {
	for i := range l.json.aeonStats {
		as := &l.json.aeonStats[i]
		err := assignIDs(l, as.AVals)
		if err != nil {
			return err
		}

		err = assignIDs(l, as.BVals)
		if err != nil {
			return err
		}

		err = l.completeAeonXVals(as.XVals)
		if err != nil {
			return err
		}

		aeon := l.Aeons[as.Name]
		as.AeonID = aeon.ID
		aeon.BaseStats = *as
		l.Aeons[seedtypes.Key(aeon)] = aeon
		l.AeonsID[as.AeonID] = aeon
		l.json.aeons[i].BaseStats = *as
	}

	return nil
}

func (l *Lookup) completeAeonXVals(xVals []seedtypes.XVal) error {
	for i := range xVals {
		xVal := &xVals[i]
		err := assignIDs(l, xVal.BaseStats)
		if err != nil {
			return err
		}
	}

	return nil
}
