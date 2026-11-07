package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)


func createMixSimple(r *http.Request, id int32, _ Subsection) (apiresource.SimpleResource, error) {
	i := E.Mixes
	mix, _ := lookup.GetResourceByID(id, i.objLookupID)
	combinations := mix.PossibleCombinations

	prefix := fmt.Sprintf("/api/%s", E.Items.Endpoint)
	path := strings.ToLower(r.URL.Path)
	segments := getPathSegments(path, E.Items.Endpoint)
	isItemsSection := len(segments) == 2 && strings.HasPrefix(path, prefix)

	if isItemsSection {
		itemID, _ := strconv.Atoi(segments[0])
		item, _ := lookup.GetResourceByID(int32(itemID), lookup.L.ItemsID)

		combinations = h.Filter(combinations, func(c seedtypes.MixCombination) bool {
			return item.Name == c.FirstItem || item.Name == c.SecondItem
		})
	}

	mixSimple := sectiontypes.MixSimple{
		ID:           mix.ID,
		Name:         mix.Name,
		URL:          ah.CreateResourceURL(i.Endpoint, id),
		Category:     mix.Category,
		Combinations: ah.ConvertObjSlice(combinations, convertMixCombinationSimple),
	}

	return mixSimple, nil
}

func convertMixCombinationSimple(mc seedtypes.MixCombination) string {
	return fmt.Sprintf("%s + %s", mc.FirstItem, mc.SecondItem)
}
