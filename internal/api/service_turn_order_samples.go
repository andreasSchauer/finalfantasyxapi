package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func getTurnOrderParamsSamples() []SampleParams[TurnOrderParams] {
	return []SampleParams[TurnOrderParams]{
		{
			Description: "Check the turns of one party member.",
			Params: TurnOrderParams{
				Party: []turnOrderParty{
					{
						UnitID: 1,
						Agility: 34,
					},
				},
			},
		},
		{
			Description: "Your party against a formation.",
			Params: TurnOrderParams{
				FormationID: h.GetInt32Ptr(25),
				Party: []turnOrderParty{
					{
						UnitID: 1,
						Agility: 15,
					},
					{
						UnitID: 2,
						Agility: 16,
					},
					{
						UnitID: 3,
						Agility: 17,
					},
				},
			},
		},
		{
			Description: "Your party against a custom selection of monsters (preemptive strike; one monster was slowed).",
			Params: TurnOrderParams{
				BattleStart: string(database.BattleStartPreemptive),
				Party: []turnOrderParty{
					{
						UnitID: 1,
						Agility: 25,
					},
					{
						UnitID: 2,
						Agility: 23,
					},
					{
						UnitID: 3,
						Agility: 20,
					},
				},
				Mons: []turnOrderMon{
					{
						MonsterID: h.GetInt32Ptr(140),
						Status: h.GetStrPtr(string(database.HasteStatusSlow)),
					},
					{
						MonsterID: h.GetInt32Ptr(139),
						AltState: h.GetInt32Ptr(1),
					},
					{
						MonsterID: h.GetInt32Ptr(141),
					},
				},
			},
		},
		{
			Description: "Measure raw turn frequency of your party against two custom monsters.",
			Params: TurnOrderParams{
				IgnFirstTurn: true,
				Party: []turnOrderParty{
					{
						UnitID: 1,
						Agility: 160,
					},
					{
						UnitID: 2,
						Agility: 180,
					},
					{
						UnitID: 3,
						Agility: 120,
					},
				},
				Mons: []turnOrderMon{
					{
						MonsterID: h.GetInt32Ptr(140),
						Agility: h.GetInt32Ptr(135),
					},
					{
						Name: h.GetStrPtr("fast boi"),
						Agility: h.GetInt32Ptr(200),
					},
				},
			},
		},
		{
			Description: "Endgame-setup.",
			Params: TurnOrderParams{
				IgnFirstTurn: true,
				Party: []turnOrderParty{
					{
						UnitID: 1,
						Agility: 255,
						Status: h.GetStrPtr(string(database.HasteStatusAutoHaste)),
					},
					{
						UnitID: 2,
						Agility: 255,
						Status: h.GetStrPtr(string(database.HasteStatusAutoHaste)),
					},
					{
						UnitID: 3,
						Agility: 255,
						Status: h.GetStrPtr(string(database.HasteStatusAutoHaste)),
					},
				},
				Mons: []turnOrderMon{
					{
						MonsterID: h.GetInt32Ptr(292),
					},
				},
			},
		},
	}
}