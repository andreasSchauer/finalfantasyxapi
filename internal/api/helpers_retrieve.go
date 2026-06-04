package api

import (
	"database/sql"
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
	if errExceptEmptyQuery(err) {
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

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterMonsterIDsByAvailability(r.Context(), database.FilterMonsterIDsByAvailabilityParams{
		Ids:            resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
		AvlType:        locContext.AvlType,
		LocContextID:   locContext.ID,
		LocContextType: locContext.Type,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlMonsterFormations(cfg *Config, r *http.Request, resources []UnnamedAPIResource) ([]UnnamedAPIResource, error) {
	i := cfg.e.monsterFormations

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterMonsterFormationIDsByAvailability(r.Context(), database.FilterMonsterFormationIDsByAvailabilityParams{
		Ids:          	resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
		AvlType:      	locContext.AvlType,
		LocContextID: 	locContext.ID,
		LocContextType: locContext.Type,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlMasterItems(cfg *Config, r *http.Request, resources []TypedAPIResource) ([]TypedAPIResource, error) {
	i := cfg.e.allItems

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	method, err := getQueryValuePtr(r, "method", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterMasterItemIDsByAvailability(r.Context(), database.FilterMasterItemIDsByAvailabilityParams{
		Ids:          	resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
		AvlType:      	locContext.AvlType,
		LocContextID: 	locContext.ID,
		LocContextType: locContext.Type,
		Method: 		h.GetNullString(method),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlItems(cfg *Config, r *http.Request, resources []NamedAPIResource) ([]NamedAPIResource, error) {
	i := cfg.e.items

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	method, err := getQueryValuePtr(r, "method", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterItemIDsByAvailability(r.Context(), database.FilterItemIDsByAvailabilityParams{
		Ids:          	resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
		AvlType:      	locContext.AvlType,
		LocContextID: 	locContext.ID,
		LocContextType: locContext.Type,
		Method: 		h.GetNullString(method),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlKeyItems(cfg *Config, r *http.Request, resources []NamedAPIResource) ([]NamedAPIResource, error) {
	i := cfg.e.keyItems

	avlParams, err := checkAvl(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterKeyItemIDsByAvailability(r.Context(), database.FilterKeyItemIDsByAvailabilityParams{
		Ids:          	resToIDs(resources),
		Availability: 	avlParams.availabilities,
		PreAirship:		avlParams.preAirship,
		LocContextID: 	locContext.ID,
		LocContextType: locContext.Type,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlSpheres(cfg *Config, r *http.Request, resources []NamedAPIResource) ([]NamedAPIResource, error) {
	i := cfg.e.spheres

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	method, err := getQueryValuePtr(r, "method", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterSphereIDsByAvailability(r.Context(), database.FilterSphereIDsByAvailabilityParams{
		Ids:          	resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
		AvlType:      	locContext.AvlType,
		LocContextID: 	locContext.ID,
		LocContextType: locContext.Type,
		Method: 		h.GetNullString(method),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlPrimers(cfg *Config, r *http.Request, resources []NamedAPIResource) ([]NamedAPIResource, error) {
	i := cfg.e.primers

	avlParams, err := checkAvl(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	dbIDs, err := cfg.db.FilterPrimerIDsByAvailability(r.Context(), database.FilterPrimerIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		Availability: 	avlParams.availabilities,
		PreAirship:		avlParams.preAirship,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlAutoAbilities(cfg *Config, r *http.Request, resources []NamedAPIResource) ([]NamedAPIResource, error) {
	i := cfg.e.autoAbilities

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	reqItem, err := getQueryBoolPtr(r, "req_item", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	charID, err := getQueryIdPtr(r, cfg.e.characters, "character", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterAutoAbilityIDsByAvailability(r.Context(), database.FilterAutoAbilityIDsByAvailabilityParams{
		Ids:          	resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
		AvlType:      	locContext.AvlType,
		LocContextID: 	locContext.ID,
		LocContextType: locContext.Type,
		CharacterID: 	h.GetNullInt32(charID),
		ReqItem:      	h.GetNullBool(reqItem),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlShops(cfg *Config, r *http.Request, resources []UnnamedAPIResource) ([]UnnamedAPIResource, error) {
	i := cfg.e.shops

	avlParams, err := checkAvl(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	sources, err := getShopSources(cfg, r, i)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterShopIDsByAvailability(r.Context(), database.FilterShopIDsByAvailabilityParams{
		Ids:          		resToIDs(resources),
		Availability: 	avlParams.availabilities,
		PreAirship:		avlParams.preAirship,
		AvlType:      		sources.AvlType,
		RequiredSources:    sources.RequiredSources,
		ExcludedSources: 	sources.ExcludedSources,
		AutoAbilityID: 		sources.AutoAbilityID,
		CharacterID: 		sources.CharacterID,
		EmptySlots: 		sources.EmptySlots,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlTreasures(cfg *Config, r *http.Request, resources []UnnamedAPIResource) ([]UnnamedAPIResource, error) {
	i := cfg.e.treasures

	avlParams, err := checkAvl(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	dbIDs, err := cfg.db.FilterTreasureIDsByAvailability(r.Context(), database.FilterTreasureIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		Availability: 	avlParams.availabilities,
		PreAirship:		avlParams.preAirship,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlQuests(cfg *Config, r *http.Request, resources []QuestAPIResource) ([]QuestAPIResource, error) {
	i := cfg.e.quests

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	dbIDs, err := cfg.db.FilterQuestIDsByAvailability(r.Context(), database.FilterQuestIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlSidequests(cfg *Config, r *http.Request, resources []QuestAPIResource) ([]QuestAPIResource, error) {
	i := cfg.e.sidequests

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	dbIDs, err := cfg.db.FilterSidequestIDsByAvailability(r.Context(), database.FilterSidequestIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlSubquests(cfg *Config, r *http.Request, resources []QuestAPIResource) ([]QuestAPIResource, error) {
	i := cfg.e.subquests

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	dbIDs, err := cfg.db.FilterSubquestIDsByAvailability(r.Context(), database.FilterSubquestIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlAreas(cfg *Config, r *http.Request, resources []AreaAPIResource) ([]AreaAPIResource, error) {
	i := cfg.e.areas

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	sources, err := getLocBasedSources(cfg, r, i)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterAreaIDsByAvailability(r.Context(), database.FilterAreaIDsByAvailabilityParams{
		Ids:             resToIDs(resources),
		Availability:    avlParams.availabilities,
		IsRepeatable:    avlParams.isRepeatable,
		PreAirship: 	 avlParams.preAirship,
		RequiredSources: sources.RequiredSources,
		ExcludedSources: sources.ExcludedSources,
		MonsterID: 		 sources.MonsterID,
		ItemID: 		 sources.ItemID,
		KeyItemID: 		 sources.KeyItemID,
		Methods: 		 sources.Methods,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlSublocations(cfg *Config, r *http.Request, resources []NamedAPIResource) ([]NamedAPIResource, error) {
	i := cfg.e.sublocations

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	sources, err := getLocBasedSources(cfg, r, i)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterSublocationIDsByAvailability(r.Context(), database.FilterSublocationIDsByAvailabilityParams{
		Ids:             resToIDs(resources),
		Availability:    avlParams.availabilities,
		IsRepeatable:    avlParams.isRepeatable,
		PreAirship: 	 avlParams.preAirship,
		RequiredSources: sources.RequiredSources,
		ExcludedSources: sources.ExcludedSources,
		MonsterID: 		 sources.MonsterID,
		ItemID: 		 sources.ItemID,
		KeyItemID: 		 sources.KeyItemID,
		Methods: 		 sources.Methods,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlLocations(cfg *Config, r *http.Request, resources []NamedAPIResource) ([]NamedAPIResource, error) {
	i := cfg.e.locations

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	sources, err := getLocBasedSources(cfg, r, i)
	if err != nil {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterLocationIDsByAvailability(r.Context(), database.FilterLocationIDsByAvailabilityParams{
		Ids:             resToIDs(resources),
		Availability:    avlParams.availabilities,
		IsRepeatable:    avlParams.isRepeatable,
		PreAirship: 	 avlParams.preAirship,
		RequiredSources: sources.RequiredSources,
		ExcludedSources: sources.ExcludedSources,
		MonsterID: 		 sources.MonsterID,
		ItemID: 		 sources.ItemID,
		KeyItemID: 		 sources.KeyItemID,
		Methods: 		 sources.Methods,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}

type avlParams struct {
	availabilities 	[]int32
	isRepeatable	sql.NullBool
	preAirship		bool
}


func checkAvlAndRep[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) (avlParams, error) {
	availabilities, errAvl := parseEnumListQuery(cfg, r, cfg.e.availabilityType.endpoint, i.queryLookup["availability"], cfg.t.AvailabilityType)
	if errExceptEmptyQuery(errAvl) {
		return avlParams{}, errAvl
	}

	isRepeatable, errRepl := getQueryBoolPtr(r, "repeatable", i.queryLookup)
	if errExceptEmptyQuery(errRepl) {
		return avlParams{}, errRepl
	}

	if queryIsEmpty(errAvl) && queryIsEmpty(errRepl) {
		return avlParams{}, errEmptyQuery
	}

	preAirship, err := parseBooleanQuery(r, i.queryLookup["pre_airship"])
	if errExceptEmptyQuery(err) {
		return avlParams{}, err
	}
	avlRanks := avlToRanks(availabilities, preAirship)

	params := avlParams{
		availabilities: h.SliceOrNil(avlRanks),
		isRepeatable: 	h.GetNullBool(isRepeatable),
		preAirship: 	preAirship,
	}

	return params, nil
}


func checkAvl[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) (avlParams, error) {
	availabilities, err := parseEnumListQuery(cfg, r, cfg.e.availabilityType.endpoint, i.queryLookup["availability"], cfg.t.AvailabilityType)
	if errExceptEmptyQuery(err) {
		return avlParams{}, err
	}
	if queryIsEmpty(err) {
		return avlParams{}, errEmptyQuery
	}

	preAirship, err := parseBooleanQuery(r, i.queryLookup["pre_airship"])
	if errExceptEmptyQuery(err) {
		return avlParams{}, err
	}
	avlRanks := avlToRanks(availabilities, preAirship)

	params := avlParams{
		availabilities: h.SliceOrNil(avlRanks),
		preAirship: 	preAirship,
	}

	return params, nil
}


func avlToRanks(avls []database.AvailabilityType, preAirship bool) []int32 {
	rankMap := map[database.AvailabilityType]int32{
		database.AvailabilityTypeAlways: 1,
		database.AvailabilityTypePost: 2,
		database.AvailabilityTypePreStory: 3,
		database.AvailabilityTypePostStory: 4,
	}
	if preAirship {
		rankMap[database.AvailabilityTypePreStory] = 2
		rankMap[database.AvailabilityTypePost] = 3
	}

	ranks := []int32{}

	for _, avl := range avls {
		ranks = append(ranks, rankMap[avl])
	}

	return ranks
}



type locContextParams struct {
	AvlType 	string
	ID   		sql.NullInt32
	Type 		sql.NullString
}

func getLocContextParams[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) (locContextParams, error) {
	avlType := AvlTypeSelf
	var locContextID *int32
	var locContextType string

	locID, err := getQueryIdPtr(r, cfg.e.locations, "location", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		locContextID = locID
		locContextType = string(ViewSourceTypeLocation)
	}

	subLocID, err := getQueryIdPtr(r, cfg.e.sublocations, "sublocation", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		locContextID = subLocID
		locContextType = string(ViewSourceTypeSublocation)
	}

	areaID, err := getQueryIdPtr(r, cfg.e.areas, "area", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locContextParams{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeArea
		locContextID = areaID
		locContextType = string(ViewSourceTypeArea)
	}

	params := locContextParams{
		AvlType:        string(avlType),
		ID:   h.GetNullInt32(locContextID),
		Type: h.GetNullString(&locContextType),
	}

	return params, nil
}


type locBasedSources struct {
	RequiredSources []string
	ExcludedSources []string
	MonsterID       sql.NullInt32
	ItemID          sql.NullInt32
	KeyItemID       sql.NullInt32
	Methods         []string
}


func getLocBasedSources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) (locBasedSources, error) {
	reqs := []string{}
	excls := []string{}

	monID, err := getQueryIdPtr(r, cfg.e.monsters, "monster", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}
	if !queryIsEmpty(err) {
		reqs = append(reqs, "monster-single")
	}
	
	itemID, err := getQueryIdPtr(r, cfg.e.items, "item", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}
	if !queryIsEmpty(err) {
		reqs = append(reqs, string(ViewSourceTypeItem))
	}

	keyItemID, err := getQueryIdPtr(r, cfg.e.keyItems, "key_item", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}
	if !queryIsEmpty(err) {
		reqs = append(reqs, string(ViewSourceTypeKeyItem))
	}

	methods, err := parseValueListQuery(cfg, r, i.queryLookup["method"])
	if errExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}

	reqs, excls, err = parseBoolSources(r, i, reqs, excls, map[string]string{
		"monsters": 	string(ViewSourceTypeMonster),
		"boss_fights": 	string(ViewSourceTypeBoss),
		"shops": 		string(ViewSourceTypeShop),
		"treasures": 	string(ViewSourceTypeTreasure),
		"sidequests": 	string(ViewSourceTypeQuest),
	})
	if err != nil {
		return locBasedSources{}, err
	}

	sources := locBasedSources{
		RequiredSources: h.SliceOrNil(reqs),
		ExcludedSources: h.SliceOrNil(excls),
		MonsterID: 		 h.GetNullInt32(monID),
		ItemID: 		 h.GetNullInt32(itemID),
		KeyItemID: 		 h.GetNullInt32(keyItemID),
		Methods: 		 h.SliceOrNil(methods),
	}

	return sources, nil
}


type shopSources struct {
	AvlType				string
	RequiredSources 	[]string
	ExcludedSources 	[]string
	AutoAbilityID       sql.NullInt32
	CharacterID       	sql.NullInt32
	EmptySlots        	[]int32
}


func getShopSources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) (shopSources, error) {
	avlType := AvlTypeSelf
	reqs := []string{}
	excls := []string{}

	autoAbilityID, err := getQueryIdPtr(r, cfg.e.autoAbilities, "auto_ability", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return shopSources{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		reqs = append(reqs, "equip_filter")
	}

	emptySlots, err := parseIntListQuery(cfg, r, i.queryLookup["empty_slots"])
	if errExceptEmptyQuery(err) {
		return shopSources{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		reqs = append(reqs, "equip_filter")
	}

	charID, err := getQueryNameIdPtr(r, cfg.e.characters, "character", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return shopSources{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		reqs = append(reqs, "equip_filter")
	}

	reqs, excls, err = parseBoolSources(r, i, reqs, excls, map[string]string{
		"items": 		string(ViewSourceTypeItem),
		"equipment": 	string(ViewSourceTypeEquipment),
	})
	if err != nil {
		return shopSources{}, err
	}

	sources := shopSources{
		AvlType:      		string(avlType),
		RequiredSources:    h.SliceOrNil(reqs),
		ExcludedSources: 	h.SliceOrNil(excls),
		AutoAbilityID: 		h.GetNullInt32(autoAbilityID),
		CharacterID: 		h.GetNullInt32(charID),
		EmptySlots: 		h.SliceOrNil(emptySlots),
	}

	return sources, nil
}



func parseBoolSources[T seeding.Lookupable, R any, A APIResource, L APIResourceList](r *http.Request, i handlerInput[T, R, A, L], reqs, excls []string, sourceMap map[string]string) ([]string, []string, error) {
	for queryParam := range sourceMap {
		b, err := parseBooleanQuery(r, i.queryLookup[queryParam])
		if errExceptEmptyQuery(err) {
			return nil, nil, err
		}
		if !queryIsEmpty(err) {
			if b {
				reqs = append(reqs, sourceMap[queryParam])
			} else {
				excls = append(excls, sourceMap[queryParam])
			}
		}
	}

	return reqs, excls, nil
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
	ViewSourceTypeEquipment		   ViewSourceType = "equip"
)

func resToIDs[A APIResource](resources []A) []int32 {
	ids := []int32{}

	for _, res := range resources {
		ids = append(ids, res.GetID())
	}

	return ids
}
