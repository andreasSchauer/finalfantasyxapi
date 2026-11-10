package get

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/typehelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

func applyOmnisElements(r *http.Request, mon apitypes.Monster, queryName ts.QueryParamName) ([]apitypes.ElementalResist, error) {
	circlesAffinity := map[int]string{
		0: string(database.ElementalAffinityNeutral),
		1: string(database.ElementalAffinityHalved),
		2: string(database.ElementalAffinityImmune),
		3: string(database.ElementalAffinityAbsorb),
		4: string(database.ElementalAffinityAbsorb),
	}

	circleCounts, err := parseOmnisQuery(r, queryName)
	if errs.QueryIsEmpty(err) {
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
		elementLookup, err := lookup.GetResource(elementName, lookup.L.Elements)
		if err != nil {
			return nil, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get element '%s'", elementName), err)
		}
		oppositeElementName := *elementLookup.OppositeElement // safe, because holy continues

		circleCount := circleCounts[elementName]
		circleCountOpposite := circleCounts[oppositeElementName]

		// if the opposite element has 4 circles, seymour omnis is weak to this element
		if circleCountOpposite == 4 {
			mon.ElemResists[i] = th.NewElemResist(elementName, string(database.ElementalAffinityWeak))
			continue
		}

		// else, change the affinity based on the count of the circles
		mon.ElemResists[i] = th.NewElemResist(elementName, circlesAffinity[circleCount])
	}

	return mon.ElemResists, nil
}

// verifies the input and counts the amounts of circles pointing to each element
func parseOmnisQuery(r *http.Request, queryName ts.QueryParamName) (map[string]int, error) {
	queryParam := apilookups.Q.Monsters[queryName]
	query, err := queryParam.CheckEmpty(r)
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
		return nil, errs.NewHTTPError(http.StatusBadRequest, "invalid input. omnis_elements must contain a combination of exactly four letters. valid letters are 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water)", nil)
	}

	for _, char := range query {
		element, ok := elements[char]
		if !ok {
			return nil, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid letter '%c' for omnis_elements. use any four-letter-combination of 'f' (fire), 'i' (ice), 'l' (lightning), 'w' (water)", char), nil)
		}

		circleCounts[element]++
	}

	return circleCounts, nil
}
