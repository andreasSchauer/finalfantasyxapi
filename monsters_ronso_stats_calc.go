package main

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func getRonsoStats(mon Monster, kimahriStats map[string]int32) map[string]int32 {
	ronsoStats := make(map[string]int32)

	ronsoStats["hp"] = getRonsoHP(mon, kimahriStats)
	ronsoStats["strength"] = getRonsoStrength(mon, kimahriStats)
	ronsoStats["magic"] = getRonsoMagic(mon, kimahriStats)
	ronsoStats["agility"] = getRonsoAgility(mon, kimahriStats)

	return ronsoStats
}

func getRonsoHP(mon Monster, kimahriStats map[string]int32) int32 {
	kimahriStr := kimahriStats["strength"]
	kimahriMag := kimahriStats["magic"]

	v1 := float64(h.PowInt(kimahriStr, 3))
	v2 := float64(h.PowInt(kimahriMag, 3))
	v3 := (v1 + v2) / 2 * 16 / 15

	hpMod := ((int32(v3)/32)+30)*586/730 + 1

	if mon.Name == "biran ronso" {
		return int32(hpMod) * 8
	}

	if mon.Name == "yenke ronso" {
		return int32(hpMod) * 6
	}

	return 0
}

func getRonsoStrength(mon Monster, kimahriStats map[string]int32) int32 {
	kimahriHP := kimahriStats["hp"]
	strengthVals := []int32{11, 12, 13, 15, 17, 19, 21, 22, 23, 24, 25, 27}

	powerMod := min((kimahriHP-644)/200, 11)

	strength := strengthVals[powerMod]

	if mon.Name == "yenke ronso" {
		strength /= 2
	}

	return strength
}

func getRonsoMagic(mon Monster, kimahriStats map[string]int32) int32 {
	kimahriHP := kimahriStats["hp"]
	magicVals := []int32{8, 8, 9, 10, 12, 14, 16, 17, 19, 20, 21, 22}

	powerMod := min((kimahriHP-644)/200, 11)

	magic := magicVals[powerMod]

	if mon.Name == "biran ronso" {
		magic /= 2
	}

	return magic
}

func getRonsoAgility(mon Monster, kimahriStats map[string]int32) int32 {
	var agility int32
	kimahriAgility := kimahriStats["agility"]

	if mon.Name == "biran ronso" {
		agility = max(kimahriAgility-4, 1)
	}

	if mon.Name == "yenke ronso" {
		agility = max(kimahriAgility-6, 1)
	}

	return agility
}
