package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetEnemyAbility(t *testing.T) {
	t.Parallel()
	tests := []EnemyAbility{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/enemy-abilities/444",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "enemy ability with provided id '444' doesn't exist. max id: 443",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/enemy-abilities/56",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":            1,
					"battle interactions": 1,
					"battle interactions 0 removed status conditions": 10,
				},
			},
			NameVer:          newExpNameVer(56, "attack", 40),
			rank:             h.GetInt32Ptr(1),
			appearsInHelpBar: false,
			canCopyCat:       false,
			untypedAbility:   453,
			monsters:         []int32{297},
			battleInteractions: []BattleInteraction{
				{
					target:         "single-target",
					abilityRange:   h.GetInt32Ptr(1),
					hitAmount:      1,
					shatterRate:    0,
					basedOnUserAtk: false,
					darkable:       false,
					silenceable:    false,
					reflectable:    false,
					accuracy: Accuracy{
						accSource:   "rate",
						hitChance:   h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage: &Damage{
						attackType: database.AttackTypeAttack,
						damageType: database.DamageTypePhysical,
						hp: &ClassDamage{
							damageFormula:  database.DamageFormulaStrVsDef,
							damageConstant: 16,
						},
						critical:        h.GetStrPtr("crit-ability"),
						criticalPlusVal: h.GetInt32Ptr(25),
						isPiercing:      false,
						breakDmgLmt:     h.GetStrPtr("always"),
						element:         nil,
					},
					inflictedDelay:            nil,
					inflictedStatusConditions: []int32{},
					removedStatusConditions:   []int32{21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
					copiedStatusConditions:    []int32{},
					statChanges:               []StatChange{},
					modifierChanges:           []ModChange{},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/enemy-abilities/228",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":            1,
					"battle interactions": 1,
				},
			},
			NameVer:          newExpNameVer(228, "karma", 2),
			rank:             h.GetInt32Ptr(3),
			appearsInHelpBar: true,
			canCopyCat:       false,
			untypedAbility:   625,
			monsters:         []int32{265},
			battleInteractions: []BattleInteraction{
				{
					target:         "single-target",
					abilityRange:   h.GetInt32Ptr(3),
					hitAmount:      1,
					shatterRate:    0,
					basedOnUserAtk: false,
					darkable:       false,
					silenceable:    false,
					reflectable:    false,
					accuracy: Accuracy{
						accSource:   "rate",
						hitChance:   h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage: &Damage{
						attackType: database.AttackTypeAttack,
						damageType: database.DamageTypeSpecial,
						hp: &ClassDamage{
							damageFormula:  database.DamageFormulaSpecialKills,
							damageConstant: 100,
						},
						critical:        nil,
						criticalPlusVal: nil,
						isPiercing:      false,
						breakDmgLmt:     h.GetStrPtr("always"),
						element:         nil,
					},
					inflictedDelay:            nil,
					inflictedStatusConditions: []int32{},
					removedStatusConditions:   []int32{},
					copiedStatusConditions:    []int32{},
					statChanges:               []StatChange{},
					modifierChanges:           []ModChange{},
				},
			},
		},
		{
			TestGeneral: TestGeneral{
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
			NameVer:          newExpNameVer(393, "swallow", 2),
			rank:             h.GetInt32Ptr(3),
			appearsInHelpBar: false,
			canCopyCat:       false,
			untypedAbility:   790,
			monsters:         []int32{227},
			battleInteractions: []BattleInteraction{
				{
					target:         "single-target",
					abilityRange:   h.GetInt32Ptr(3),
					hitAmount:      1,
					shatterRate:    0,
					basedOnUserAtk: false,
					darkable:       false,
					silenceable:    false,
					reflectable:    false,
					accuracy: Accuracy{
						accSource:   "rate",
						hitChance:   h.GetInt32Ptr(255),
						accModifier: nil,
					},
					damage:                    nil,
					inflictedDelay:            nil,
					inflictedStatusConditions: []int32{},
					removedStatusConditions:   []int32{4, 10, 12, 14, 16, 17, 18, 19, 20},
					copiedStatusConditions:    []int32{4, 12, 14, 16, 20, 21, 27, 28, 24, 25},
					statChanges:               []StatChange{},
					modifierChanges:           []ModChange{},
				},
			},
		},
	}

	testSingleResources(t, tests, "GetEnemyAbility", testSrv.HandleEnemyAbilities)
}

func TestGetMultipleEnemyAbilities(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/enemy-abilities/attack?limit=max",
				ExpectedStatus: http.StatusMultipleChoices,
			},
			Count:   43,
			Results: []int32{17, 22, 24, 32, 37, 39, 45, 54, 55, 59},
		},
		{
			TestGeneral: TestGeneral{
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
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/enemy-abilities?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    443,
			Previous: nil,
			Next:     nil,
			Results:  []int32{1, 223, 18, 45, 177, 386, 443},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/enemy-abilities?target_type=single-target&rank=3&damage_type=magical&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:    21,
			Previous: nil,
			Next:     nil,
			Results:  []int32{12, 111, 150, 180, 257, 379, 435},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/enemy-abilities?damage_formula=11",
				ExpectedStatus: http.StatusOK,
			},
			Count:    16,
			Previous: nil,
			Next:     nil,
			Results:  []int32{126, 133, 155, 247, 402},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/enemy-abilities?monster=244",
				ExpectedStatus: http.StatusOK,
			},
			Count:    4,
			Previous: nil,
			Next:     nil,
			Results:  []int32{46, 227, 427, 429},
		},
		{
			TestGeneral: TestGeneral{
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
