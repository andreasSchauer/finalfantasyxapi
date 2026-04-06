package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type MixSimple struct {
	ID            	int32       `json:"id"`
	Name		 	string		`json:"name"`
	URL           	string      `json:"url"`
	Category		string		`json:"category"`
	Combinations	[]string	`json:"combinations"`
}

func (m MixSimple) GetURL() string {
	return m.URL
}

func createMixSimple(cfg *Config, r *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.mixes
	mix, _ := seeding.GetResourceByID(id, i.objLookupID)
	combinations := mix.PossibleCombinations
	
	prefix := fmt.Sprintf("/api/%s", cfg.e.items.endpoint)
	path := strings.ToLower(r.URL.Path)
	segments := getPathSegments(path, cfg.e.items.endpoint)
	isItemsSection := len(segments) == 2 && strings.HasPrefix(path, prefix)

	if isItemsSection {
		itemID, _ := strconv.Atoi(segments[0])
		item, _ := seeding.GetResourceByID(int32(itemID), cfg.l.ItemsID)

		combinations = h.Filter(combinations, func(c seeding.MixCombination) bool {
			return item.Name == c.FirstItem || item.Name == c.SecondItem
		})
	}

	mixSimple := MixSimple{
		ID: 			mix.ID,
		Name: 			mix.Name,
		URL: 			createResourceURL(cfg, i.endpoint, id),
		Category: 		mix.Category,
		Combinations: 	convertObjSlice(cfg, combinations, convertMixCombinationSimple),
	}

	return mixSimple, nil
}

func convertMixCombinationSimple(cfg *Config, mc seeding.MixCombination) string {
	return fmt.Sprintf("%s + %s", mc.FirstItem, mc.SecondItem)
}
