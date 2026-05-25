package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getMultipleAPIResources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], name string) (L, error) {
	var zeroType L

	dbIDs, err := i.getMultipleQuery(r.Context(), name)
	if err != nil {
		return zeroType, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get multiple %s with name '%s'.", i.resourceType, name), err)
	}

	return idsToAPIResourceList(cfg, r, i, dbIDs)
}

func retrieveAPIResources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) ([]A, error) {
	dbIDs, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return nil, err
	}

	return idsToAPIResources(cfg, i, dbIDs), nil
}

func filterAPIResources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], resources []A, filteredLists []filteredResList[A]) (L, error) {
	var zeroType L
	filteredRes := resources

	for _, filtered := range filteredLists {
		if filtered.err != nil {
			return zeroType, filtered.err
		}
		filteredRes = getSharedResources(filteredRes, filtered.resources)
	}

	if i.avlFunc != nil {
		var err error
		filteredRes, err = i.avlFunc(cfg, r, filteredRes)
		if err != nil {
			return zeroType, err
		}
	}

	flip, err := parseBooleanQuery(r, i.queryLookup["flip"])
	if errIsNotEmptyQuery(err) {
		return zeroType, err
	}

	if flip {
		filteredRes = removeResources(resources, filteredRes)
	}

	resourceList, err := i.resToListFunc(cfg, r, filteredRes)
	if err != nil {
		return zeroType, err
	}

	return resourceList, nil
}

func filterAvlMonsters(cfg *Config, r *http.Request, resources []NamedAPIResource) ([]NamedAPIResource, error) {
	i := cfg.e.monsters

	availabilities, err := parseEnumListQuery(cfg, r, cfg.e.availabilityType.endpoint, i.queryLookup["availability"], cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}
	if errors.Is(err, errEmptyQuery) {
		return resources, nil
	}

	params := database.FilterMonsterIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		AvlType:      string(AvlTypeSelf),
		Availability: availabilities,
	}

	locID, err := getQueryIdPtr(r, cfg.e.locations, "location", i.queryLookup)
	if err == nil {
		params.AvlType = string(AvlTypeContext)
		params.LocContextID = h.GetNullInt32(locID)
		ct := string(ViewSourceTypeLocation)
		params.LocContextType = h.GetNullString(&ct)
	}

	subLocID, err := getQueryIdPtr(r, cfg.e.sublocations, "sublocation", i.queryLookup)
	if err == nil {
		params.AvlType = string(AvlTypeContext)
		params.LocContextID = h.GetNullInt32(subLocID)
		ct := string(ViewSourceTypeSublocation)
		params.LocContextType = h.GetNullString(&ct)
	}

	areaID, err := getQueryIdPtr(r, cfg.e.areas, "area", i.queryLookup)
	if err == nil {
		params.AvlType = string(AvlTypeArea)
		params.LocContextID = h.GetNullInt32(areaID)
		ct := string(ViewSourceTypeArea)
		params.LocContextType = h.GetNullString(&ct)
	}

	dbIDs, err := cfg.db.FilterMonsterIDsByAvailability(r.Context(), params)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}

func filterAvlMonsterFormations(cfg *Config, r *http.Request, resources []UnnamedAPIResource) ([]UnnamedAPIResource, error) {
	i := cfg.e.monsterFormations

	availabilities, err := parseEnumListQuery(cfg, r, cfg.e.availabilityType.endpoint, i.queryLookup["availability"], cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}
	if errors.Is(err, errEmptyQuery) {
		return resources, nil
	}

	params := database.FilterMonsterFormationIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		AvlType:      string(AvlTypeSelf),
		Availability: availabilities,
	}

	locID, err := getQueryIdPtr(r, cfg.e.locations, "location", i.queryLookup)
	if err == nil {
		params.AvlType = string(AvlTypeContext)
		params.LocContextID = h.GetNullInt32(locID)
		ct := string(ViewSourceTypeLocation)
		params.LocContextType = h.GetNullString(&ct)
	}

	subLocID, err := getQueryIdPtr(r, cfg.e.sublocations, "sublocation", i.queryLookup)
	if err == nil {
		params.AvlType = string(AvlTypeContext)
		params.LocContextID = h.GetNullInt32(subLocID)
		ct := string(ViewSourceTypeSublocation)
		params.LocContextType = h.GetNullString(&ct)
	}

	areaID, err := getQueryIdPtr(r, cfg.e.areas, "area", i.queryLookup)
	if err == nil {
		params.AvlType = string(AvlTypeArea)
		params.LocContextID = h.GetNullInt32(areaID)
		ct := string(ViewSourceTypeArea)
		params.LocContextType = h.GetNullString(&ct)
	}

	dbIDs, err := cfg.db.FilterMonsterFormationIDsByAvailability(r.Context(), params)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}

func filterAvlShops(cfg *Config, r *http.Request, resources []UnnamedAPIResource) ([]UnnamedAPIResource, error) {
	i := cfg.e.shops

	availabilities, err := parseEnumListQuery(cfg, r, cfg.e.availabilityType.endpoint, i.queryLookup["availability"], cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}
	if errors.Is(err, errEmptyQuery) {
		return resources, nil
	}

	params := database.FilterShopIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		AvlType:      string(AvlTypeSelf),
		Availability: availabilities,
		SubTypes:     []string{},
	}

	_, err = parseIdQuery(r, i.queryLookup["auto_ability"], len(cfg.l.AutoAbilities))
	if err == nil {
		params.AvlType = string(AvlTypeContext)
	}

	_, err = parseIntListQuery(cfg, r, i.queryLookup["empty_slots"])
	if err == nil {
		params.AvlType = string(AvlTypeContext)
	}

	_, err = parseBooleanQuery(r, i.queryLookup["items"])
	if err == nil {
		params.AvlType = string(AvlTypeContext)
		params.SubTypes = append(params.SubTypes, "item")
	}

	_, err = parseBooleanQuery(r, i.queryLookup["equipment"])
	if err == nil {
		params.AvlType = string(AvlTypeContext)
		params.SubTypes = append(params.SubTypes, "equip")
	}

	params.SubTypes = h.SliceOrNil(params.SubTypes)

	dbIDs, err := cfg.db.FilterShopIDsByAvailability(r.Context(), params)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}

func filterAvlTreasures(cfg *Config, r *http.Request, resources []UnnamedAPIResource) ([]UnnamedAPIResource, error) {
	i := cfg.e.treasures

	availabilities, err := parseEnumListQuery(cfg, r, cfg.e.availabilityType.endpoint, i.queryLookup["availability"], cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}
	if errors.Is(err, errEmptyQuery) {
		return resources, nil
	}

	params := database.FilterTreasureIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		Availability: availabilities,
	}

	dbIDs, err := cfg.db.FilterTreasureIDsByAvailability(r.Context(), params)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}

func filterAvlQuests(cfg *Config, r *http.Request, resources []QuestAPIResource) ([]QuestAPIResource, error) {
	i := cfg.e.quests

	availabilities, err := parseEnumListQuery(cfg, r, cfg.e.availabilityType.endpoint, i.queryLookup["availability"], cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}
	if errors.Is(err, errEmptyQuery) {
		return resources, nil
	}

	params := database.FilterQuestIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		Availability: availabilities,
	}

	dbIDs, err := cfg.db.FilterQuestIDsByAvailability(r.Context(), params)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}

func filterAvlSidequests(cfg *Config, r *http.Request, resources []QuestAPIResource) ([]QuestAPIResource, error) {
	i := cfg.e.sidequests

	availabilities, err := parseEnumListQuery(cfg, r, cfg.e.availabilityType.endpoint, i.queryLookup["availability"], cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}
	if errors.Is(err, errEmptyQuery) {
		return resources, nil
	}

	params := database.FilterSidequestIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		Availability: availabilities,
	}

	dbIDs, err := cfg.db.FilterSidequestIDsByAvailability(r.Context(), params)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}

func filterAvlSubquests(cfg *Config, r *http.Request, resources []QuestAPIResource) ([]QuestAPIResource, error) {
	i := cfg.e.subquests

	availabilities, err := parseEnumListQuery(cfg, r, cfg.e.availabilityType.endpoint, i.queryLookup["availability"], cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}
	if errors.Is(err, errEmptyQuery) {
		return resources, nil
	}

	params := database.FilterSubquestIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		Availability: availabilities,
	}

	dbIDs, err := cfg.db.FilterSubquestIDsByAvailability(r.Context(), params)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}

func filterAvlAreas(cfg *Config, r *http.Request, resources []AreaAPIResource) ([]AreaAPIResource, error) {
	i := cfg.e.areas

	availabilities, err := parseEnumListQuery(cfg, r, cfg.e.availabilityType.endpoint, i.queryLookup["availability"], cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}
	if errors.Is(err, errEmptyQuery) {
		return resources, nil
	}

	params := database.FilterAreaIDsByAvailabilityParams{
		Ids:          		resToIDs(resources),
		Availability: 		availabilities,
		RequiredSources: 	[]string{},
		ExcludedSources: 	[]string{},
	}

	monID, _ := getQueryIdPtr(r, cfg.e.monsters, "monster", i.queryLookup)
	if monID != nil {
		params.MonsterID = h.GetNullInt32(monID)
		sourceType := ViewSourceTypeMonster
		params.RequiredSources = append(params.RequiredSources, string(sourceType))
	}

	keyItemID, _ := getQueryIdPtr(r, cfg.e.keyItems, "key_item", i.queryLookup)
	if keyItemID != nil {
		params.KeyItemID = h.GetNullInt32(keyItemID)
		sourceType := ViewSourceTypeKeyItem
		params.RequiredSources = append(params.RequiredSources, string(sourceType))
	}

	itemID, _ := getQueryIdPtr(r, cfg.e.items, "item", i.queryLookup)
	if itemID != nil {
		params.ItemID = h.GetNullInt32(itemID)
		sourceType := ViewSourceTypeItem
		params.RequiredSources = append(params.RequiredSources, string(sourceType))
	}

	methods, err := parseValueListQuery(cfg, r, i.queryLookup["method"])
	if err == nil {
		params.Methods = methods
	}

	hasMons, err := parseBooleanQuery(r, i.queryLookup["monsters"])
	if err == nil {
		sourceType := ViewSourceTypeMonster
		if hasMons {
			params.RequiredSources = append(params.RequiredSources, string(sourceType))
		} else {
			params.ExcludedSources = append(params.ExcludedSources, string(sourceType))
		}
	}

	hasBosses, err := parseBooleanQuery(r, i.queryLookup["boss_fights"])
	if err == nil {
		sourceType := ViewSourceTypeBoss
		if hasBosses {
			params.RequiredSources = append(params.RequiredSources, string(sourceType))
		} else {
			params.ExcludedSources = append(params.ExcludedSources, string(sourceType))
		}
	}

	hasShops, err := parseBooleanQuery(r, i.queryLookup["shops"])
	if err == nil {
		sourceType := ViewSourceTypeShop
		if hasShops {
			params.RequiredSources = append(params.RequiredSources, string(sourceType))
		} else {
			params.ExcludedSources = append(params.ExcludedSources, string(sourceType))
		}
	}

	hasTreasures, err := parseBooleanQuery(r, i.queryLookup["treasures"])
	if err == nil {
		sourceType := ViewSourceTypeTreasure
		if hasTreasures {
			params.RequiredSources = append(params.RequiredSources, string(sourceType))
		} else {
			params.ExcludedSources = append(params.ExcludedSources, string(sourceType))
		}
	}

	hasQuests, err := parseBooleanQuery(r, i.queryLookup["sidequests"])
	if err == nil {
		sourceType := ViewSourceTypeQuest
		if hasQuests {
			params.RequiredSources = append(params.RequiredSources, string(sourceType))
		} else {
			params.ExcludedSources = append(params.ExcludedSources, string(sourceType))
		}
	}

	params.RequiredSources = h.SliceOrNil(params.RequiredSources)
	params.ExcludedSources = h.SliceOrNil(params.ExcludedSources)

	dbIDs, err := cfg.db.FilterAreaIDsByAvailability(r.Context(), params)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}

func filterAvlSublocations(cfg *Config, r *http.Request, resources []NamedAPIResource) ([]NamedAPIResource, error) {
	i := cfg.e.sublocations

	availabilities, err := parseEnumListQuery(cfg, r, cfg.e.availabilityType.endpoint, i.queryLookup["availability"], cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}
	if errors.Is(err, errEmptyQuery) {
		return resources, nil
	}

	params := database.FilterSublocationIDsByAvailabilityParams{
		Ids:          		resToIDs(resources),
		Availability: 		availabilities,
		RequiredSources: 	[]string{},
		ExcludedSources: 	[]string{},
	}

	monID, _ := getQueryIdPtr(r, cfg.e.monsters, "monster", i.queryLookup)
	if monID != nil {
		params.MonsterID = h.GetNullInt32(monID)
		sourceType := ViewSourceTypeMonster
		params.RequiredSources = append(params.RequiredSources, string(sourceType))
	}

	keyItemID, _ := getQueryIdPtr(r, cfg.e.keyItems, "key_item", i.queryLookup)
	if keyItemID != nil {
		params.KeyItemID = h.GetNullInt32(keyItemID)
		sourceType := ViewSourceTypeKeyItem
		params.RequiredSources = append(params.RequiredSources, string(sourceType))
	}

	itemID, _ := getQueryIdPtr(r, cfg.e.items, "item", i.queryLookup)
	if itemID != nil {
		params.ItemID = h.GetNullInt32(itemID)
		sourceType := ViewSourceTypeItem
		params.RequiredSources = append(params.RequiredSources, string(sourceType))
	}

	methods, err := parseValueListQuery(cfg, r, i.queryLookup["method"])
	if err == nil {
		params.Methods = methods
	}

	hasMons, err := parseBooleanQuery(r, i.queryLookup["monsters"])
	if err == nil {
		sourceType := ViewSourceTypeMonster
		if hasMons {
			params.RequiredSources = append(params.RequiredSources, string(sourceType))
		} else {
			params.ExcludedSources = append(params.ExcludedSources, string(sourceType))
		}
	}

	hasBosses, err := parseBooleanQuery(r, i.queryLookup["boss_fights"])
	if err == nil {
		sourceType := ViewSourceTypeBoss
		if hasBosses {
			params.RequiredSources = append(params.RequiredSources, string(sourceType))
		} else {
			params.ExcludedSources = append(params.ExcludedSources, string(sourceType))
		}
	}

	hasShops, err := parseBooleanQuery(r, i.queryLookup["shops"])
	if err == nil {
		sourceType := ViewSourceTypeShop
		if hasShops {
			params.RequiredSources = append(params.RequiredSources, string(sourceType))
		} else {
			params.ExcludedSources = append(params.ExcludedSources, string(sourceType))
		}
	}

	hasTreasures, err := parseBooleanQuery(r, i.queryLookup["treasures"])
	if err == nil {
		sourceType := ViewSourceTypeTreasure
		if hasTreasures {
			params.RequiredSources = append(params.RequiredSources, string(sourceType))
		} else {
			params.ExcludedSources = append(params.ExcludedSources, string(sourceType))
		}
	}

	hasQuests, err := parseBooleanQuery(r, i.queryLookup["sidequests"])
	if err == nil {
		sourceType := ViewSourceTypeQuest
		if hasQuests {
			params.RequiredSources = append(params.RequiredSources, string(sourceType))
		} else {
			params.ExcludedSources = append(params.ExcludedSources, string(sourceType))
		}
	}

	params.RequiredSources = h.SliceOrNil(params.RequiredSources)
	params.ExcludedSources = h.SliceOrNil(params.ExcludedSources)

	dbIDs, err := cfg.db.FilterSublocationIDsByAvailability(r.Context(), params)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}

func filterAvlLocations(cfg *Config, r *http.Request, resources []NamedAPIResource) ([]NamedAPIResource, error) {
	i := cfg.e.locations

	availabilities, err := parseEnumListQuery(cfg, r, cfg.e.availabilityType.endpoint, i.queryLookup["availability"], cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}
	if errors.Is(err, errEmptyQuery) {
		return resources, nil
	}

	params := database.FilterLocationIDsByAvailabilityParams{
		Ids:          		resToIDs(resources),
		Availability: 		availabilities,
		RequiredSources: 	[]string{},
		ExcludedSources: 	[]string{},
	}

	monID, _ := getQueryIdPtr(r, cfg.e.monsters, "monster", i.queryLookup)
	if monID != nil {
		params.MonsterID = h.GetNullInt32(monID)
		sourceType := ViewSourceTypeMonster
		params.RequiredSources = append(params.RequiredSources, string(sourceType))
	}

	keyItemID, _ := getQueryIdPtr(r, cfg.e.keyItems, "key_item", i.queryLookup)
	if keyItemID != nil {
		params.KeyItemID = h.GetNullInt32(keyItemID)
		sourceType := ViewSourceTypeKeyItem
		params.RequiredSources = append(params.RequiredSources, string(sourceType))
	}

	itemID, _ := getQueryIdPtr(r, cfg.e.items, "item", i.queryLookup)
	if itemID != nil {
		params.ItemID = h.GetNullInt32(itemID)
		sourceType := ViewSourceTypeItem
		params.RequiredSources = append(params.RequiredSources, string(sourceType))
	}

	methods, err := parseValueListQuery(cfg, r, i.queryLookup["method"])
	if err == nil {
		params.Methods = methods
	}

	hasMons, err := parseBooleanQuery(r, i.queryLookup["monsters"])
	if err == nil {
		sourceType := ViewSourceTypeMonster
		if hasMons {
			params.RequiredSources = append(params.RequiredSources, string(sourceType))
		} else {
			params.ExcludedSources = append(params.ExcludedSources, string(sourceType))
		}
	}

	hasBosses, err := parseBooleanQuery(r, i.queryLookup["boss_fights"])
	if err == nil {
		sourceType := ViewSourceTypeBoss
		if hasBosses {
			params.RequiredSources = append(params.RequiredSources, string(sourceType))
		} else {
			params.ExcludedSources = append(params.ExcludedSources, string(sourceType))
		}
	}

	hasShops, err := parseBooleanQuery(r, i.queryLookup["shops"])
	if err == nil {
		sourceType := ViewSourceTypeShop
		if hasShops {
			params.RequiredSources = append(params.RequiredSources, string(sourceType))
		} else {
			params.ExcludedSources = append(params.ExcludedSources, string(sourceType))
		}
	}

	hasTreasures, err := parseBooleanQuery(r, i.queryLookup["treasures"])
	if err == nil {
		sourceType := ViewSourceTypeTreasure
		if hasTreasures {
			params.RequiredSources = append(params.RequiredSources, string(sourceType))
		} else {
			params.ExcludedSources = append(params.ExcludedSources, string(sourceType))
		}
	}

	hasQuests, err := parseBooleanQuery(r, i.queryLookup["sidequests"])
	if err == nil {
		sourceType := ViewSourceTypeQuest
		if hasQuests {
			params.RequiredSources = append(params.RequiredSources, string(sourceType))
		} else {
			params.ExcludedSources = append(params.ExcludedSources, string(sourceType))
		}
	}

	params.RequiredSources = h.SliceOrNil(params.RequiredSources)
	params.ExcludedSources = h.SliceOrNil(params.ExcludedSources)

	dbIDs, err := cfg.db.FilterLocationIDsByAvailability(r.Context(), params)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}

type AvlType string

const (
	AvlTypeSelf    AvlType = "self"
	AvlTypeContext AvlType = "context"
	AvlTypeArea    AvlType = "area"
)

type ViewSourceType string

const (
	ViewSourceTypeMonster          ViewSourceType = "monster"
	ViewSourceTypeBoss             ViewSourceType = "boss"
	ViewSourceTypeMonsterFormation ViewSourceType = "monster-formation"
	ViewSourceTypeLocation         ViewSourceType = "location"
	ViewSourceTypeSublocation      ViewSourceType = "sublocation"
	ViewSourceTypeArea             ViewSourceType = "area"
	ViewSourceTypeTreasure         ViewSourceType = "treasure"
	ViewSourceTypeShop             ViewSourceType = "shop"
	ViewSourceTypeQuest            ViewSourceType = "quest"
	ViewSourceTypeBlitzball        ViewSourceType = "blitzball"
	ViewSourceTypeItem             ViewSourceType = "item"
	ViewSourceTypeKeyItem          ViewSourceType = "key-item"
)

func resToIDs[A APIResource](resources []A) []int32 {
	ids := []int32{}

	for _, res := range resources {
		ids = append(ids, res.GetID())
	}

	return ids
}
