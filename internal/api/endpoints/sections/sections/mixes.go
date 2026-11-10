package sections

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func CreateMixSimple(r *http.Request, id int32, _ handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.Mixes
	mix, _ := lookup.GetResourceByID(id, i.ObjLookupID)
	combinations := mix.PossibleCombinations

	prefix := fmt.Sprintf("/api/%s", apilookups.E.Items.Endpoint)
	path := strings.ToLower(r.URL.Path)
	segments := ah.GetPathSegments(path, apilookups.E.Items.Endpoint)
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
