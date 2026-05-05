package seeding

type AeonStat struct {
	AeonID int32
	Name   string     `json:"name"`
	AVals  []BaseStat `json:"a_vals"`
	BVals  []BaseStat `json:"b_vals"`
	XVals  []XVal     `json:"x_vals"`
}

type XVal struct {
	Battles   int32      `json:"battles"`
	BaseStats []BaseStat `json:"base_stats"`
}

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
		l.Aeons[Key(aeon)] = aeon
		l.AeonsID[as.AeonID] = aeon
		l.json.aeons[i].BaseStats = *as
	}

	return nil
}

func (l *Lookup) completeAeonXVals(xVals []XVal) error {
	for i := range xVals {
		xVal := &xVals[i]
		err := assignIDs(l, xVal.BaseStats)
		if err != nil {
			return err
		}
	}

	return nil
}
