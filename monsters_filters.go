package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


func (cfg *apiConfig) getMonstersElemResist(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("elemental-affinities")
	
	if query == "" {
		return inputMons, nil
	}

	eaPairs := strings.Split(query, ",")
	var ids []int32

	for _, pair := range eaPairs {
		parts := strings.Split(pair, "-")
		if len(parts) != 2 {
			return nil, newHTTPError(http.StatusBadRequest, "invalid input. usage: elemental-affinities={element}-{affinity},{element}-{affinity}", nil)
		}

		element, err := parseSingleSegmentResource(parts[0], cfg.l.Elements)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid element: %s. element doesn't exist. use /api/elements to see existing elements.", parts[0]), err)
		}

		affinity, err := parseSingleSegmentResource(parts[1], cfg.l.Affinities)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid affinity: %s. affinity doesn't exist. use /api/affinities to see existing affinities.", parts[1]), err)
		}

		elemResist := seeding.ElementalResist{
			ElementID: element.ID,
			AffinityID: affinity.ID,
		}

		elemResistLookup, err := seeding.GetResource(elemResist, cfg.l.ElementalResists)
		if err != nil {
			return []NamedAPIResource{}, nil
		}

		ids = append(ids, elemResistLookup.ID)
	}

	dbMons, err := cfg.db.GetMonstersByElemResistIDs(r.Context(), ids)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by elemental affinities", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersStatusResist(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryStatusses := r.URL.Query().Get("status-resists")
	queryResistance := r.URL.Query().Get("resistance")
	defaultResistance := 1
	var resistance int

	if queryStatusses == "" && queryResistance != "" {
		return nil, newHTTPError(http.StatusBadRequest, "invalid input. resistance parameter must be paired with status-resists parameter. usage: status-resists={status},{status},...&resistance={int from 1-254}", nil)
	}
	
	if queryStatusses == "" {
		return inputMons, nil
	}
	
	switch queryResistance {
	case "":
		resistance = defaultResistance
	case "immune":
		resistance = 254
	default:
		var err error
		resistance, err = strconv.Atoi(queryResistance)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, "invalid resistance", err)
		}
	}
	
	if resistance > 254 || resistance <= 0 {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid resistance. resistance must be a positive integer between 1 or 254, with %d being the default value, if no resistance was provided.", defaultResistance), nil)
	}
	
	statusses := strings.Split(queryStatusses, ",")
	var ids []int32
	
	for _, qStatus := range statusses {
		status, err := parseSingleSegmentResource(qStatus, cfg.l.StatusConditions)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid status condition: %s. status condition doesn't exist. use /api/status-conditions to see existing status conditions.", qStatus), err)
		}

		ids = append(ids, status.ID)
	}

	dbMons, err := cfg.db.GetMonstersByStatusResists(r.Context(), database.GetMonstersByStatusResistsParams{
		StatusConditionIds: ids,
		MinResistance: 		resistance,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by status conditions", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersItem(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryItem := r.URL.Query().Get("item")
	queryMethod := r.URL.Query().Get("method")

	if queryItem == "" && queryMethod != "" {
		return nil, newHTTPError(http.StatusBadRequest, "invalid input. method parameter must be paired with item parameter. usage: item={item},&method={steal/drop/bribe/other}", nil)
	}
	
	if queryItem == "" {
		return inputMons, nil
	}

	item, err := parseSingleSegmentResource(queryItem, cfg.l.MasterItems)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid item: %s. item doesn't exist. use /api/items to see existing items.", queryItem), err)
	}

	var dbMons []database.Monster

	switch queryMethod {
	case "":
		dbMons, err = cfg.db.GetMonstersByItem(r.Context(), item.ID)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by item", err)
		}
	case "steal":
		dbMons, err = cfg.db.GetMonstersByItemSteal(r.Context(), item.ID)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by steal item", err)
		}
	case "drop":
		dbMons, err = cfg.db.GetMonstersByItemDrop(r.Context(), item.ID)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by drop item", err)
		}
	case "bribe":
		dbMons, err = cfg.db.GetMonstersByItemBribe(r.Context(), item.ID)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by bribe item", err)
		}
	case "other":
		dbMons, err = cfg.db.GetMonstersByItemOther(r.Context(), item.ID)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by other items", err)
		}
	default:
		return nil, newHTTPError(http.StatusBadRequest, "invalid method. allowed methods: steal, drop, bribe, other.", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersAutoAbility(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("auto-abilities")
	
	if query == "" {
		return inputMons, nil
	}
	
	abilities := strings.Split(query, ",")
	var ids []int32

	for _, ability := range abilities {
		autoAbility, err := parseSingleSegmentResource(ability, cfg.l.AutoAbilities)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid auto ability: %s. auto ability doesn't exist. use /api/auto-abilities to see existing auto-abilities.", ability), err)
		}

		ids = append(ids, autoAbility.ID)
	}

	dbMons, err := cfg.db.GetMonstersByAutoAbilityIDs(r.Context(), ids)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by auto ability", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersRonsoRage(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	const ronsoRageOffset int = 35
	query := r.URL.Query().Get("ronso-rage")
	var ronsoRageID int32
	
	if query == "" {
		return inputMons, nil
	}

	id, err := strconv.Atoi(query)
	if err == nil {
		id += ronsoRageOffset
		ronsoRageID = int32(id)
	}

	if ronsoRageID == 0 {
		ronsoRage, err := parseNameVersionResource(query, "", cfg.l.Overdrives)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid ronso rage: %s. ronso rage doesn't exist. use /api/ronso-rages to see existing ronso-rages.", query), err)
		}
		ronsoRageID = ronsoRage.ID
	}

	if int(ronsoRageID) > ronsoRageOffset + 12 || int(ronsoRageID) <= ronsoRageOffset {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid ronso rage id: %d. provided id must be between 1 and 12", ronsoRageID-35), err)
	}

	dbMons, err := cfg.db.GetMonstersByRonsoRageID(r.Context(), ronsoRageID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by ronso rage", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersLocation(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("location")
	
	if query == "" {
		return inputMons, nil
	}

	location, err := parseSingleSegmentResource(query, cfg.l.Locations)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid location: %s. location doesn't exist. use /api/locations to see existing locations", query), err)
	}

	dbMons, err := cfg.db.GetLocationMonsters(r.Context(), location.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by location", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.GetLocationMonstersRow) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersSubLocation(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("sublocation")
	
	if query == "" {
		return inputMons, nil
	}

	sublocation, err := parseSingleSegmentResource(query, cfg.l.SubLocations)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid sublocation: %s. sublocation doesn't exist. use /api/sublocations to see existing sublocations.", query), err)
	}

	dbMons, err := cfg.db.GetSublocationMonsters(r.Context(), sublocation.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by sublocation", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.GetSublocationMonstersRow) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersArea(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryArea := r.URL.Query().Get("area")
	areaCount, _ := cfg.db.GetAreaCount(r.Context())
	
	if queryArea == "" {
		return inputMons, nil
	}

	area_id, err := strconv.Atoi(queryArea)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "invalid area id", err)
	}

	if area_id < 0 || area_id > int(areaCount) {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid area: %d. area doesn't exist. use /api/areas to see existing sublocations.", area_id), err)
	}

	dbMons, err := cfg.db.GetAreaMonsters(r.Context(), int32(area_id))
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by area", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.GetAreaMonstersRow) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersDistance(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("distance")

	if query == "" {
		return inputMons, nil
	}

	distance, err := strconv.Atoi(query)
	if err != nil || distance > 4 || distance <= 0 {
		return nil, newHTTPError(http.StatusBadRequest, "invalid value. distance must be an integer from 1 to 4.", err)
	}

	dbMons, err := cfg.db.GetMonstersByDistance(r.Context(), distance)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by distance", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersStoryBased(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("is-story-based")

	if query == "" {
		return inputMons, nil
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "invalid value. usage: is-story-based={boolean}", err)
	}

	dbMons, err := cfg.db.GetMonstersByIsStoryBased(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve story-based monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersRepeatable(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("is-repeatable")

	if query == "" {
		return inputMons, nil
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "invalid value. usage: is-repeatable={boolean}", err)
	}

	dbMons, err := cfg.db.GetMonstersByIsRepeatable(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve repeatable monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersCanBeCaptured(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("can-capture")

	if query == "" {
		return inputMons, nil
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "invalid value. usage: can-capture={boolean}", err)
	}

	dbMons, err := cfg.db.GetMonstersByCanBeCaptured(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters that can be captured", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersHasOverdrive(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("has-overdrive")

	if query == "" {
		return inputMons, nil
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "invalid value. usage: has-overdrive={boolean}", err)
	}

	dbMons, err := cfg.db.GetMonstersByHasOverdrive(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters with overdrive gauge", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersUnderwater(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("is-underwater")

	if query == "" {
		return inputMons, nil
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "invalid value. usage: is-underwater={boolean}", err)
	}

	dbMons, err := cfg.db.GetMonstersByIsUnderwater(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve underwater monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersZombie(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("is-zombie")

	if query == "" {
		return inputMons, nil
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "invalid value. usage: is-zombie={boolean}", err)
	}

	dbMons, err := cfg.db.GetMonstersByIsZombie(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve zombie monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersSpecies(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("species")

	if query == "" {
		return inputMons, nil
	}

	enum, err := GetEnumType(query, cfg.t.MonsterSpecies)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value: %s, use /api/monster-species to see valid values", query), err)
	}

	species := database.MonsterSpecies(enum.Name)

	dbMons, err := cfg.db.GetMonstersBySpecies(r.Context(), species)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by species", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}



func (cfg *apiConfig) getMonstersCreationArea(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("creation-area")

	if query == "" {
		return inputMons, nil
	}

	enum, err := GetEnumType(query, cfg.t.MaCreationArea)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value: %s, use /api/creation-areas to see valid values", query), err)
	}

	area := h.NullMaCreationArea(&enum.Name)

	dbMons, err := cfg.db.GetMonstersByMaCreationArea(r.Context(), area)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by creation area", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}


func (cfg *apiConfig) getMonstersType(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("type")

	if query == "" {
		return inputMons, nil
	}

	enum, err := GetEnumType(query, cfg.t.CTBIconType)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value: %s, use /api/ctb-icon-type to see valid values. 'boss' and 'boss-numbered' will both yield all boss monsters.", query), err)
	}

	species := database.CtbIconType(enum.Name)
	var dbMons []database.Monster

	switch species {
	case database.CtbIconTypeBoss, database.CtbIconTypeBossNumbered:
		dbMons, err = cfg.db.GetMonstersByCTBIconTypeBoss(r.Context())
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by type", err)
		}
	default:
		dbMons, err = cfg.db.GetMonstersByCTBIconType(r.Context(), species)
		if err != nil {
			return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by type", err)
		}
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	sharedResources := getSharedResources(inputMons, resources)

	return sharedResources, nil
}