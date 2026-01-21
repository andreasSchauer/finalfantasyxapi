package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func applyOmnisElements(cfg *Config, r *http.Request, mon Monster, queryName string) ([]ElementalResist, error) {
	circlesAffinity := map[int]string{
		0: "neutral",
		1: "halved",
		2: "immune",
		3: "absorb",
		4: "absorb",
	}

	circleCounts, err := parseOmnisQuery(cfg, r, queryName)
	if errors.Is(err, errEmptyQuery) {
		return mon.ElemResists, nil
	}
	if err != nil {
		return nil, err
	}

	for i, elemResist := range mon.ElemResists {
		elementName := elemResist.Element.Name

		if elementName == "holy" {
			continue
		}

		// lookup element to find out the opposite element
		elementLookup, err := seeding.GetResource(elementName, cfg.l.Elements)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get element '%s'.", elementName), err)
		}
		oppositeElementName := *elementLookup.OppositeElement

		circleCount := circleCounts[elementName]
		circleCountOpposite := circleCounts[oppositeElementName]

		// if the opposite element has 4 circles, seymour omnis is weak to this element
		if circleCountOpposite == 4 {
			mon.ElemResists[i] = newElemResist(cfg, elementName, "weak")
			continue
		}

		// else, change the affinity based on the count of the circles
		mon.ElemResists[i] = newElemResist(cfg, elementName, circlesAffinity[circleCount])
	}

	return mon.ElemResists, nil
}

// verifies the input and counts the amounts of circles pointing to each element
func parseOmnisQuery(cfg *Config, r *http.Request, queryName string) (map[string]int, error) {
	queryParam := cfg.q.monsters[queryName]
	query, err := checkEmptyQuery(r, queryParam)
	if err != nil {
		return nil, err
	}

	elements := map[rune]string{
		'f': "fire",
		'i': "ice",
		'l': "lightning",
		'w': "water",
	}
	circleCounts := make(map[string]int)

	if len(query) != 4 {
		return nil, newHTTPError(http.StatusBadRequest, "invalid input. omnis_elements must contain a combination of exactly four letters. valid letters are 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water).", nil)
	}

	for _, char := range query {
		element, ok := elements[char]
		if !ok {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid letter '%c' for omnis_elements. use any four-letter-combination of 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water).", char), nil)
		}

		circleCounts[element]++
	}

	return circleCounts, nil
}
