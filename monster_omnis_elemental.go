package main

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


func (cfg *Config) getOmnisElemResists(r *http.Request, mon Monster) ([]ElementalResist, error) {
	countToAffinity := map[int]string{
		0: "neutral",
		1: "halved",
		2: "immune",
		3: "absorb",
		4: "absorb",
	}

	counts, isEmpty, err := cfg.verifyOmnisQuery(r, mon)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return mon.ElemResists, nil
	}

	for i, elemResist := range mon.ElemResists {
		element := elemResist.Element

		if element.Name == "holy" {
			continue
		}

		elementLookup, err := seeding.GetResource(element.Name, cfg.l.Elements)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get element %s", element.Name), err)
		}
		oppositeName := *elementLookup.OppositeElement

		count := counts[element.Name]
		countOpposite := counts[oppositeName]

		if countOpposite == 4 {
			newAffinity, err := cfg.changeElemResist(element, "weak")
			if err != nil {
				return nil, err
			}
			mon.ElemResists[i] = newAffinity
			continue
		}

		newAffinity, err := cfg.changeElemResist(element, countToAffinity[count])
		if err != nil {
			return nil, err
		}
		mon.ElemResists[i] = newAffinity
	}

	return mon.ElemResists, nil
}


func (cfg *Config) verifyOmnisQuery(r *http.Request, mon Monster) (map[string]int, bool, error) {
	queryParam := "omnis-elements"
	elements := map[rune]string{
		'f': "fire",
		'i': "ice",
		'l': "lightning",
		'w': "water",
	}
	counts := make(map[string]int)
	isEmpty := false

	key := seeding.LookupObject{
		Name: "seymour omnis",
	}
	omnis, err := seeding.GetResource(key, cfg.l.Monsters)
	if err != nil {
		return nil, false, newHTTPError(http.StatusInternalServerError, "couldn't get seymour omnis data", err)
	}

	query := r.URL.Query().Get(queryParam)
	if query == "" {
		isEmpty = true
		return nil, isEmpty, nil
	}

	if query != "" && mon.ID != omnis.ID {
		return nil, false, newHTTPError(http.StatusBadRequest, "Invalid usage. omnis-elements can only be used on Seymour Omnis", nil)
	}

	if len(query) != 4 {
		return nil, false, newHTTPError(http.StatusBadRequest, "Invalid input. omnis-elements must contain a combination of exactly four letters. Valid letters are 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water).", nil)
	}

	for _, char := range query {
		element, ok := elements[char]
		if !ok {
			return nil, false, newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid letter %c for omnis-elements. Use any four letter combination of 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water).", char), nil)
		}

		counts[element]++
	}

	return counts, false, nil
}