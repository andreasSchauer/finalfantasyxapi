package api


type DamageParams struct {

}


type BattleEntityParams struct {
	Monster		*BattleMonsterParams	`json:"monster,omitempty"`
}

type BattleMonsterParams struct {
	ID				int32	`json:"id"`
	AltState		*int32	`json:"alt_state,omitempty"`
	StatsOverride
}

type StatsOverride struct {
	HP				*int32		`json:"hp,omitempty"`
	MP				*int32		`json:"mp,omitempty"`
	Strength		*int32		`json:"strength,omitempty"`
	Defense			*int32		`json:"defense,omitempty"`
	Magic			*int32		`json:"magic,omitempty"`
	MagicDefense	*int32		`json:"magic_defense,omitempty"`
	Agility			*int32		`json:"agility,omitempty"`
	Luck			*int32		`json:"luck,omitempty"`
	Evasion			*int32		`json:"evasion,omitempty"`
	Accuracy		*int32		`json:"accuracy,omitempty"`
}

type CustomEntityParams struct {
	PlayerUnit		*int32		`json:",omitempty"`
	StatTable

}

type StatTable struct {
	HP				int32		`json:"hp,omitempty"`
	MP				int32		`json:"mp,omitempty"`
	Strength		int32		`json:"strength,omitempty"`
	Defense			int32		`json:"defense,omitempty"`
	Magic			int32		`json:"magic,omitempty"`
	MagicDefense	int32		`json:"magic_defense,omitempty"`
	Agility			int32		`json:"agility,omitempty"`
	Luck			int32		`json:"luck,omitempty"`
	Evasion			int32		`json:"evasion,omitempty"`
	Accuracy		int32		`json:"accuracy,omitempty"`
}

type ElementTable struct {
	
}