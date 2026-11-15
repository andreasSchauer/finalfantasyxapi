package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetEnemyAbility(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.EnemyAbility{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/enemy-abilities/444",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "enemy ability with provided id '444' doesn't exist. max id: 443",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/enemy-abilities/56",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":            1,
					"battle interactions": 1,
					"battle interactions 0 removed status conditions": 10,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(56, "attack", 40),
			Rank:             h.GetInt32Ptr(1),
			AppearsInHelpBar: false,
			CanCopyCat:       false,
			UntypedAbility:   453,
			Monsters:         []int32{297},
			BattleInteractions: []apitesttypes.BattleInteraction{
				{
					Target:         "single-target",
					AbilityRange:   h.GetInt32Ptr(1),
					HitAmount:      1,
					ShatterRate:    0,
					BasedOnUserAtk: false,
					Darkable:       false,
					Silenceable:    false,
					Reflectable:    false,
					Accuracy: apitesttypes.Accuracy{
						AccSource:   "rate",
						HitChance:   h.GetInt32Ptr(255),
						AccModifier: nil,
					},
					Damage: &apitesttypes.Damage{
						AttackType: database.AttackTypeAttack,
						DamageType: database.DamageTypePhysical,
						Hp: &apitesttypes.ClassDamage{
							DamageFormula:  database.DamageFormulaStrVsDef,
							DamageConstant: 16,
						},
						Critical:        h.GetStrPtr("crit-ability"),
						CriticalPlusVal: h.GetInt32Ptr(25),
						IsPiercing:      false,
						BreakDmgLmt:     h.GetStrPtr("always"),
						Element:         nil,
					},
					InflictedDelay:            nil,
					InflictedStatusConditions: []int32{},
					RemovedStatusConditions:   []int32{21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
					CopiedStatusConditions:    []int32{},
					StatChanges:               []apitesttypes.StatChange{},
					ModifierChanges:           []apitesttypes.ModChange{},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/enemy-abilities/228",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":            1,
					"battle interactions": 1,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(228, "karma", 2),
			Rank:             h.GetInt32Ptr(3),
			AppearsInHelpBar: true,
			CanCopyCat:       false,
			UntypedAbility:   625,
			Monsters:         []int32{265},
			BattleInteractions: []apitesttypes.BattleInteraction{
				{
					Target:         "single-target",
					AbilityRange:   h.GetInt32Ptr(3),
					HitAmount:      1,
					ShatterRate:    0,
					BasedOnUserAtk: false,
					Darkable:       false,
					Silenceable:    false,
					Reflectable:    false,
					Accuracy: apitesttypes.Accuracy{
						AccSource:   "rate",
						HitChance:   h.GetInt32Ptr(255),
						AccModifier: nil,
					},
					Damage: &apitesttypes.Damage{
						AttackType: database.AttackTypeAttack,
						DamageType: database.DamageTypeSpecial,
						Hp: &apitesttypes.ClassDamage{
							DamageFormula:  database.DamageFormulaSpecialKills,
							DamageConstant: 100,
						},
						Critical:        nil,
						CriticalPlusVal: nil,
						IsPiercing:      false,
						BreakDmgLmt:     h.GetStrPtr("always"),
						Element:         nil,
					},
					InflictedDelay:            nil,
					InflictedStatusConditions: []int32{},
					RemovedStatusConditions:   []int32{},
					CopiedStatusConditions:    []int32{},
					StatChanges:               []apitesttypes.StatChange{},
					ModifierChanges:           []apitesttypes.ModChange{},
				},
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/enemy-abilities/393",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":            1,
					"battle interactions": 1,
					"battle interactions 0 removed status conditions": 9,
					"battle interactions 0 copied status conditions":  18,
				},
			},
			NameVer:          apitesttypes.NewExpNameVer(393, "swallow", 2),
			Rank:             h.GetInt32Ptr(3),
			AppearsInHelpBar: false,
			CanCopyCat:       false,
			UntypedAbility:   790,
			Monsters:         []int32{227},
			BattleInteractions: []apitesttypes.BattleInteraction{
				{
					Target:         "single-target",
					AbilityRange:   h.GetInt32Ptr(3),
					HitAmount:      1,
					ShatterRate:    0,
					BasedOnUserAtk: false,
					Darkable:       false,
					Silenceable:    false,
					Reflectable:    false,
					Accuracy: apitesttypes.Accuracy{
						AccSource:   "rate",
						HitChance:   h.GetInt32Ptr(255),
						AccModifier: nil,
					},
					Damage:                    nil,
					InflictedDelay:            nil,
					InflictedStatusConditions: []int32{},
					RemovedStatusConditions:   []int32{4, 10, 12, 14, 16, 17, 18, 19, 20},
					CopiedStatusConditions:    []int32{4, 12, 14, 16, 20, 21, 27, 28, 24, 25},
					StatChanges:               []apitesttypes.StatChange{},
					ModifierChanges:           []apitesttypes.ModChange{},
				},
			},
		},
	}

	testSingleResources(t, tests, "GetEnemyAbility", testSrv.HandleEnemyAbilities)
}

func TestGetMultipleEnemyAbilities(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/enemy-abilities/attack?limit=max",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   43,
			Results: []int32{17, 22, 24, 32, 37, 39, 45, 54, 55, 59},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/enemy-abilities/breath",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   6,
			Results: []int32{88, 89, 90, 91, 92, 93},
		},
	}

	testIDList(t, tests, apilookups.E.EnemyAbilities.Endpoint, "GetMultipleEnemyAbilities", testSrv.HandleEnemyAbilities, compareAPIResourceLists[apiresource.Named])
}

func TestRetrieveEnemyAbilities(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/enemy-abilities?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    443,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 223, 18, 45, 177, 386, 443},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/enemy-abilities?target_type=single-target&rank=3&damage_type=magical&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    21,
			Previous: nil,
			Next:     nil,
			Results:  []int32{12, 111, 150, 180, 257, 379, 435},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/enemy-abilities?damage_formula=11",
				ExpectedStatus: http.StatusOK,
			},
			Count:    16,
			Previous: nil,
			Next:     nil,
			Results:  []int32{126, 133, 155, 247, 402},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/enemy-abilities?monster=244",
				ExpectedStatus: http.StatusOK,
			},
			Count:    4,
			Previous: nil,
			Next:     nil,
			Results:  []int32{46, 227, 427, 429},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/enemy-abilities?bdl=true&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    101,
			Previous: nil,
			Next:     nil,
			Results:  []int32{2, 50, 69, 133, 196, 228, 277, 322, 406, 428, 443},
		},
	}

	testIDList(t, tests, apilookups.E.EnemyAbilities.Endpoint, "RetrieveEnemyAbilities", testSrv.HandleEnemyAbilities, compareAPIResourceLists[apiresource.Named])
}
