package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) applyOmnisElemResists(r *http.Request, mon Monster) ([]ElementalResist, error) {
	countToAffinity := map[int]string{
		0: "neutral",
		1: "halved",
		2: "immune",
		3: "absorb",
		4: "absorb",
	}

	counts, err := cfg.verifyOmnisQuery(r)
	if errors.Is(err, errEmptyQuery) {
		return mon.ElemResists, nil
	}
	if err != nil {
		return nil, err
	}

	for i, elemResist := range mon.ElemResists {
		element := elemResist.Element

		if element.Name == "holy" {
			continue
		}

		elementLookup, err := seeding.GetResource(element.Name, cfg.l.Elements)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get element '%s'.", element.Name), err)
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

func (cfg *Config) verifyOmnisQuery(r *http.Request) (map[string]int, error) {
	queryParam := cfg.q.monsters["omnis-elements"]
	elements := map[rune]string{
		'f': "fire",
		'i': "ice",
		'l': "lightning",
		'w': "water",
	}
	counts := make(map[string]int)

	query := r.URL.Query().Get(queryParam.Name)
	if query == "" {
		return nil, errEmptyQuery
	}

	if len(query) != 4 {
		return nil, newHTTPError(http.StatusBadRequest, "invalid input. omnis-elements must contain a combination of exactly four letters. valid letters are 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water).", nil)
	}

	queryLower := strings.ToLower(query)

	for _, char := range queryLower {
		element, ok := elements[char]
		if !ok {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid letter '%c' for omnis-elements. use any four-letter-combination of 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water).", char), nil)
		}

		counts[element]++
	}

	return counts, nil
}
