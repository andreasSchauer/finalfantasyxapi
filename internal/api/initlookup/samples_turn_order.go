package initlookup

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/service/servicetypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

var TurnOrderParamsSamples = []handlerinput.SampleParams[servicetypes.TurnOrderParams]{
	{
		Description: "Load in all unique monsters of a formation and check the turn order.",
		Params: servicetypes.TurnOrderParams{
			FormationID: h.GetInt32Ptr(15),
		},
	},
	{
		Description: "Your party against a formation.",
		Params: servicetypes.TurnOrderParams{
			FormationID: h.GetInt32Ptr(25),
			Party: []servicetypes.TurnOrderParty{
				{
					UnitID:  1,
					Agility: 15,
				},
				{
					UnitID:  2,
					Agility: 16,
				},
				{
					UnitID:  3,
					Agility: 17,
				},
			},
		},
	},
	{
		Description: "Your party against a custom selection of monsters (preemptive strike; one monster was slowed).",
		Params: servicetypes.TurnOrderParams{
			BattleStart: string(database.BattleStartPreemptive),
			Party: []servicetypes.TurnOrderParty{
				{
					UnitID:  1,
					Agility: 25,
				},
				{
					UnitID:  2,
					Agility: 23,
				},
				{
					UnitID:  3,
					Agility: 20,
				},
			},
			Mons: []servicetypes.TurnOrderMon{
				{
					MonsterID: h.GetInt32Ptr(140),
					Status:    h.GetStrPtr(string(database.HasteStatusSlow)),
				},
				{
					MonsterID: h.GetInt32Ptr(139),
					AltState:  h.GetInt32Ptr(1),
				},
				{
					MonsterID: h.GetInt32Ptr(141),
				},
			},
		},
	},
	{
		Description: "Measure raw turn frequency of your party against two custom monsters.",
		Params: servicetypes.TurnOrderParams{
			IgnFirstTurn: true,
			Party: []servicetypes.TurnOrderParty{
				{
					UnitID:  1,
					Agility: 60,
				},
				{
					UnitID:  2,
					Agility: 80,
				},
				{
					UnitID:  3,
					Agility: 40,
				},
			},
			Mons: []servicetypes.TurnOrderMon{
				{
					MonsterID: h.GetInt32Ptr(140),
					Agility:   h.GetInt32Ptr(75),
				},
				{
					Name:    h.GetStrPtr("fast boi"),
					Agility: h.GetInt32Ptr(200),
				},
			},
		},
	},
	{
		Description: "Endgame-setup.",
		Params: servicetypes.TurnOrderParams{
			IgnFirstTurn: true,
			Party: []servicetypes.TurnOrderParty{
				{
					UnitID:  1,
					Agility: 255,
					Status:  h.GetStrPtr(string(database.HasteStatusAutoHaste)),
				},
				{
					UnitID:  2,
					Agility: 255,
					Status:  h.GetStrPtr(string(database.HasteStatusAutoHaste)),
				},
				{
					UnitID:  3,
					Agility: 255,
					Status:  h.GetStrPtr(string(database.HasteStatusAutoHaste)),
				},
			},
			Mons: []servicetypes.TurnOrderMon{
				{
					MonsterID: h.GetInt32Ptr(292),
				},
			},
		},
	},
}
