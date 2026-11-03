package seeding

type AeonStat struct {
	AeonID int32	  `json:"aeon_id"`
	Name   string     `json:"name"`
	AVals  []BaseStat `json:"a_vals"`
	BVals  []BaseStat `json:"b_vals"`
	XVals  []XVal     `json:"x_vals"`
}

type XVal struct {
	Battles   int32      `json:"battles"`
	BaseStats []BaseStat `json:"base_stats"`
}