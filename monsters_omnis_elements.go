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

	counts, err := cfg.parseOmnisQuery(r)
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

		// if the opposite element has 4 circles, seymour omnis is weak to this element
		if countOpposite == 4 {
			newAffinity, err := cfg.changeElemResist(element, "weak")
			if err != nil {
				return nil, err
			}
			mon.ElemResists[i] = newAffinity
			continue
		}

		// else, change the affinity based on the count of the circles
		newAffinity, err := cfg.changeElemResist(element, countToAffinity[count])
		if err != nil {
			return nil, err
		}
		mon.ElemResists[i] = newAffinity
	}

	return mon.ElemResists, nil
}

// verifies the input and counts the amounts of circles pointing to each element
func (cfg *Config) parseOmnisQuery(r *http.Request) (map[string]int, error) {
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


func (cfg *Config) changeElemResist(element NamedAPIResource, newAffinityName string) (ElementalResist, error) {
	newAffinity, err := seeding.GetResource(newAffinityName, cfg.l.Affinities)
	if err != nil {
		return ElementalResist{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get affinity '%s'.", newAffinityName), err)
	}

	newResist := cfg.newElemResist(
		element.ID,
		newAffinity.ID,
		element.Name,
		newAffinity.Name,
	)

	return newResist, nil
}