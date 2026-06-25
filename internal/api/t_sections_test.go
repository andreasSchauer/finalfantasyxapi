package api

import (
	"reflect"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


func TestAreaToLocAreaString(t *testing.T) {
	tests := []struct {
		input	int32
		exp		string
	}{
		{
			input: 	6,
			exp:	"baaj temple - submerged ruins",
		},
		{
			input: 	221,
			exp:	"zanarkand ruins",
		},
		{
			input: 	37,
			exp:	"besaid - besaid village - house - 1 (southern)",
		},
		{
			input: 	48,
			exp:	"kilika - kilika port - dock",
		},
	}

	for i, tc := range tests {
		area, _ := seeding.GetResourceByID(tc.input, testCfg.l.AreasID)
		got := areaToLocAreaString(area)

		if !reflect.DeepEqual(got, tc.exp) {
			t.Errorf("areaToLocAreaString: Testcase %d: input: %v. expected: %v, got: %v", i+1, tc.input, tc.exp, got)
		}
	}
}


func TestConvertStatusResistSimple(t *testing.T) {
	tests := []struct {
		input	seeding.StatusResist
		exp		string
	}{
		{
			input: 	seeding.StatusResist{
				StatusCondition: "darkness",
				Resistance: 50,
			},
			exp:	"darkness (50%)",
		},
		{
			input: 	seeding.StatusResist{
				StatusCondition: "silence",
				Resistance: 254,
			},
			exp:	"silence (immune)",
		},
	}

	for i, tc := range tests {
		got := convertStatusResistSimple(testCfg, tc.input)

		if !reflect.DeepEqual(got, tc.exp) {
			t.Errorf("convertStatusResistSimple: Testcase %d: input: %v. expected: %v, got: %v", i+1, tc.input, tc.exp, got)
		}
	}
}


func TestConvertAccuracySimple(t *testing.T) {
	tests := []struct {
		input	seeding.Accuracy
		exp		string
	}{
		{
			input: 	seeding.Accuracy{
				AccSource: string(database.AccSourceTypeRate),
				HitChance: h.GetInt32Ptr(80),
				AccModifier: nil,
			},
			exp:	"80% base chance of hitting",
		},
		{
			input: 	seeding.Accuracy{
				AccSource: string(database.AccSourceTypeRate),
				HitChance: h.GetInt32Ptr(255),
				AccModifier: nil,
			},
			exp:	"always hits",
		},
		{
			input: 	seeding.Accuracy{
				AccSource: string(database.AccSourceTypeAccuracy),
				HitChance: nil,
				AccModifier: h.GetFloat32Ptr(2.5),
			},
			exp:	"based on accuracy with 2.5 modifier",
		},
	}

	for i, tc := range tests {
		got := convertAccuracySimple(testCfg, tc.input)

		if !reflect.DeepEqual(got, tc.exp) {
			t.Errorf("convertStatusResistSimple: Testcase %d: input: %v. expected: %v, got: %v", i+1, tc.input, tc.exp, got)
		}
	}
}


func TestConvertInflictedDelaySimple(t *testing.T) {
	tests := []struct {
		input	seeding.InflictedDelay
		exp		string
	}{
		{
			input: 	seeding.InflictedDelay{
				DelayType: string(database.DelayTypeCtbBased),
				DamageConstant: 8,
			},
			exp:	"weak",
		},
		{
			input: 	seeding.InflictedDelay{
				DelayType: string(database.DelayTypeCtbBased),
				DamageConstant: 16,
			},
			exp:	"strong",
		},
		{
			input: 	seeding.InflictedDelay{
				DelayType: string(database.DelayTypeTickSpeedBased),
				DamageConstant: 24,
			},
			exp:	"weak",
		},
		{
			input: 	seeding.InflictedDelay{
				DelayType: string(database.DelayTypeTickSpeedBased),
				DamageConstant: 48,
			},
			exp:	"strong",
		},
	}

	for i, tc := range tests {
		got := convertInflictedDelaySimple(testCfg, tc.input)

		if !reflect.DeepEqual(got, tc.exp) {
			t.Errorf("convertStatusResistSimple: Testcase %d: input: %v. expected: %v, got: %v", i+1, tc.input, tc.exp, got)
		}
	}
}


func TestConvertInflictedStatusSimple(t *testing.T) {
	tests := []struct {
		input	seeding.InflictedStatus
		exp		string
	}{
		{
			input: 	seeding.InflictedStatus{
				StatusCondition: "darkness",
				Probability: 80,
			},
			exp:	"darkness (80%)",
		},
		{
			input: 	seeding.InflictedStatus{
				StatusCondition: "silence",
				Probability: 254,
			},
			exp:	"silence (infinite %)",
		},
		{
			input: 	seeding.InflictedStatus{
				StatusCondition: "sleep",
				Probability: 255,
			},
			exp:	"sleep (always)",
		},
	}

	for i, tc := range tests {
		got := convertInflictedStatusSimple(testCfg, tc.input)

		if !reflect.DeepEqual(got, tc.exp) {
			t.Errorf("convertStatusResistSimple: Testcase %d: input: %v. expected: %v, got: %v", i+1, tc.input, tc.exp, got)
		}
	}
}


func TestFormatChange(t *testing.T) {
	tests := []struct {
		name		string
		calcType	database.CalculationType
		val			float32
		exp			string
	}{
		{
			name:		"current-hp",
			calcType:	database.CalculationTypeAddedValue,
			val: 		5.5,
			exp:		"current-hp +5.5",
		},
		{
			name:		"current-hp",
			calcType:	database.CalculationTypeAddedValue,
			val: 		-3,
			exp:		"current-hp -3",
		},
		{
			name:		"current-hp",
			calcType:	database.CalculationTypeAddedPercentage,
			val: 		50,
			exp:		"current-hp +50%",
		},
		{
			name:		"current-hp",
			calcType:	database.CalculationTypeAddedPercentage,
			val: 		-50,
			exp:		"current-hp -50%",
		},
		{
			name:		"current-hp",
			calcType:	database.CalculationTypeMultiply,
			val: 		1.550000000,
			exp:		"current-hp x1.55",
		},
		{
			name:		"current-hp",
			calcType:	database.CalculationTypeSetValue,
			val: 		5,
			exp:		"current-hp = 5",
		},
	}

	for i, tc := range tests {
		got := formatChange(tc.name, string(tc.calcType), tc.val)

		if !reflect.DeepEqual(got, tc.exp) {
			t.Errorf("convertStatusResistSimple: Testcase %d: name: %s, calc type: %s, val: %.3f. expected: %v, got: %v", i+1, tc.name, string(tc.calcType), tc.val, tc.exp, got)
		}
	}
}


func TestConvertModChangeSimple(t *testing.T) {
	tests := []struct {
		input	seeding.ModifierChange
		exp		string
	}{
		{
			input: 	seeding.ModifierChange{
				ModifierName: "ap-gain",
				CalculationType: string(database.CalculationTypeMultiplyHighest),
				Value: 2,
			},
			exp:	"ap-gain x2",
		},
		{
			input: 	seeding.ModifierChange{
				ModifierName: "accuracy-percentage",
				CalculationType: string(database.CalculationTypeSetValue),
				Value: 10,
			},
			exp:	"accuracy-percentage = 10%",
		},
	}

	for i, tc := range tests {
		got := convertModChangeSimple(testCfg, tc.input)

		if !reflect.DeepEqual(got, tc.exp) {
			t.Errorf("convertStatusResistSimple: Testcase %d: input: %v. expected: %v, got: %v", i+1, tc.input, tc.exp, got)
		}
	}
}


func TestFoundEquipmentAbilitiesStringPtr(t *testing.T) {
	tests := []struct {
		input	seeding.TreasureEquipment
		exp		*string
	}{
		{
			input: 	seeding.TreasureEquipment{
				Abilities: []string{"firestrike", "triple overdrive"},
				EmptySlotsAmount: 0,
			},
			exp: h.GetStrPtr("firestrike, triple overdrive"),
		},
		{
			input: 	seeding.TreasureEquipment{
				Abilities: []string{"firestrike", "triple overdrive"},
				EmptySlotsAmount: 2,
			},
			exp: h.GetStrPtr("firestrike, triple overdrive, (2)"),
		},
		{
			input: 	seeding.TreasureEquipment{
				Abilities: []string{},
				EmptySlotsAmount: 4,
			},
			exp: h.GetStrPtr("(4)"),
		},
		{
			input: 	seeding.TreasureEquipment{
				Abilities: []string{},
				EmptySlotsAmount: 0,
			},
			exp: nil,
		},
	}

	for i, tc := range tests {
		got := foundEquipmentAbilitiesStringPtr(tc.input)
		match := false

		if (got == nil && tc.exp == nil) ||
		   (got != nil && tc.exp != nil && *got == *tc.exp) {
			match = true
		}

		if !match {
			t.Errorf("convertStatusResistSimple: Testcase %d: input: %v. expected: %v, got: %v", i+1, tc.input, h.DerefStringPtr(tc.exp), h.DerefStringPtr(got))
		}
	}
}



func TestPopulateTreasuresLocSimple(t *testing.T) {
	tests := []struct {
		input	[]int32
		expTreasureCount int
		expItemCount	 int
		expEquipCount	 int
		expTotalGil	 	 int32
	}{
		{
			input: 				[]int32{175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186},
			expTreasureCount: 	12,
			expItemCount: 		7,
			expEquipCount: 		3,
			expTotalGil:		7000,
		},
		{
			input: 				[]int32{13, 14},
			expTreasureCount: 	2,
			expItemCount: 		2,
			expEquipCount: 		0,
			expTotalGil:		0,
		},
		{
			input: 				[]int32{167, 168, 169, 170, 171, 172, 173, 174},
			expTreasureCount: 	8,
			expItemCount: 		6,
			expEquipCount: 		1,
			expTotalGil:		3000,
		},
	}

	for i, tc := range tests {
		got := populateTreasuresLocSimple(testCfg, tc.input)

		if !reflect.DeepEqual(got.TreasureCount, tc.expTreasureCount) {
			t.Errorf("convertStatusResistSimple: Testcase %d: input: %s. expected treasure count: %d, got: %d", i+1, h.FormatIntSlice(tc.input), tc.expTreasureCount, got.TreasureCount)
		}

		if !reflect.DeepEqual(len(got.Items), tc.expItemCount) {
			t.Errorf("convertStatusResistSimple: Testcase %d: input: %s. expected item count: %d, got: %d", i+1, h.FormatIntSlice(tc.input), tc.expItemCount, len(got.Items))
		}

		if !reflect.DeepEqual(len(got.Equipment), tc.expEquipCount) {
			t.Errorf("convertStatusResistSimple: Testcase %d: input: %s. expected equip count: %d, got: %d", i+1, h.FormatIntSlice(tc.input), tc.expEquipCount, len(got.Equipment))
		}

		if !reflect.DeepEqual(got.TotalGil, tc.expTotalGil) {
			t.Errorf("convertStatusResistSimple: Testcase %d: input: %s. expected total gil: %d, got: %d", i+1, h.FormatIntSlice(tc.input), tc.expTotalGil, got.TotalGil)
		}
	}
}



func TestGetJunctionIDs(t *testing.T) {
	tests := []struct {
		input	[]Junction
		parentIDs []int32
		expIDs [][]int32
	}{
		{
			input: []Junction{
				{
					ParentID: 15,
					ChildID: 13,
				},
				{
					ParentID: 15,
					ChildID: 14,
				},
				{
					ParentID: 36,
					ChildID: 32,
				},
				{
					ParentID: 36,
					ChildID: 33,
				},
				{
					ParentID: 36,
					ChildID: 34,
				},
				{
					ParentID: 36,
					ChildID: 35,
				},
				{
					ParentID: 36,
					ChildID: 36,
				},
				{
					ParentID: 36,
					ChildID: 37,
				},
				{
					ParentID: 143,
					ChildID: 187,
				},
				{
					ParentID: 143,
					ChildID: 188,
				},
				{
					ParentID: 143,
					ChildID: 189,
				},
			},
			parentIDs: []int32{15, 36, 143},
			expIDs: [][]int32{
				{13, 14},
				{32, 33, 34, 35, 36, 37},
				{187, 188, 189},
			},
		},
	}

	for i, tc := range tests {
		for j, id := range tc.parentIDs {
			var got []int32
			initialLen := len(tc.input)
			got, tc.input = getJunctionIDs(id, tc.input)
			combinedLen := len(got) + (len(tc.input))

			if !reflect.DeepEqual(got, tc.expIDs[j]) {
				t.Errorf("getJunctionIDs: Testcase %d: input parent id: %d. expected ids: %s, got: %s", i+1, id, h.FormatIntSlice(tc.expIDs[i]), h.FormatIntSlice(got))
			}


			if !reflect.DeepEqual(combinedLen, initialLen) {
				t.Errorf("getJunctionIDs: Testcase %d: input parent id: %d. combined length %d doesn't match with initial length: %d", i+1, id, combinedLen, initialLen)
			}
		}
	}
}