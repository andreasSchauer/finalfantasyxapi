package api

import (
	"testing"
)

var benchMarkResult DropChanceResponse

func BenchmarkHandleDropChance(b *testing.B) {
	params := DropChanceParams{
		Monster:           261,
		PartyMembers:      []int32{1, 2, 3, 4, 5, 6, 7},
		AutoAbilities:     []int32{31, 30},
		OptionalAbilities: true,
	}

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		res, err := handleDropChance(testCfg, params)
		if err != nil {
			b.Fatalf("Calculation failed during benchmark: %v", err)
		}
		benchMarkResult = res
	}
}
