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

func (cfg *Config) getMonstersElemResist(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := "elemental-affinities"
	query := r.URL.Query().Get(queryParam)

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

		element, err := parseSingleSegmentResourceQuery("element", parts[0], queryParam, cfg.l.Elements)
		if err != nil {
			return nil, err
		}

		affinity, err := parseSingleSegmentResourceQuery("affinity", parts[1], queryParam, cfg.l.Affinities)
		if err != nil {
			return nil, err
		}

		elemResist := seeding.ElementalResist{
			ElementID:  element.ID,
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

	return resources, nil
}

func (cfg *Config) getMonstersStatusResist(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam1 := "status-resists"
	queryParam2 := "resistance"
	queryStatusses := r.URL.Query().Get(queryParam1)
	queryResistance := r.URL.Query().Get(queryParam2)
	defaultResistance := 1
	maxResistance := 254
	var resistance int

	if queryStatusses == "" {
		if queryResistance != "" {
			return nil, newHTTPError(http.StatusBadRequest, "invalid input. resistance parameter must be paired with status-resists parameter. usage: status-resists={status},{status},...&resistance={1-254 or immune}", nil)
		}
		return inputMons, nil
	}

	switch queryResistance {
	case "":
		resistance = defaultResistance
	case "immune":
		resistance = maxResistance
	default:
		var err error
		resistance, err = strconv.Atoi(queryResistance)
		if err != nil {
			return nil, newHTTPError(http.StatusBadRequest, "invalid resistance", err)
		}
	}

	if resistance > maxResistance || resistance <= 0 {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid resistance: %s. resistance must be a positive integer between %d or %d, with %d being the default value, if no resistance was provided.", queryResistance, defaultResistance, maxResistance, defaultResistance), nil)
	}

	statusses := strings.Split(queryStatusses, ",")
	var ids []int32

	for _, qStatus := range statusses {
		status, err := parseSingleSegmentResourceQuery("status-condition", qStatus, queryParam1, cfg.l.StatusConditions)
		if err != nil {
			return nil, err
		}

		ids = append(ids, status.ID)
	}

	dbMons, err := cfg.db.GetMonstersByStatusResists(r.Context(), database.GetMonstersByStatusResistsParams{
		StatusConditionIds: ids,
		MinResistance:      resistance,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by status conditions", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersItem(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryItem := "item"
	queryMethod := r.URL.Query().Get("method")

	item, itemIsEmpty, err := parseUniqueNameQuery(r, queryItem, cfg.l.Items)
	if err != nil {
		return nil, err
	}
	if itemIsEmpty {
		if queryMethod != "" {
			return nil, newHTTPError(http.StatusBadRequest, "invalid input. method parameter must be paired with item parameter. usage: item={item}&method={steal/drop/bribe/other}", nil)
		}
		return inputMons, nil
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
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid method value: %s. allowed methods: steal, drop, bribe, other.", queryMethod), err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersAutoAbility(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	// could turn this into a parseQuery helper?
	// multiple unique name/id inputs to []id
	queryParam := "auto-abilities"
	query := r.URL.Query().Get(queryParam)

	if query == "" {
		return inputMons, nil
	}

	abilities := strings.Split(query, ",")
	var ids []int32

	for _, ability := range abilities {
		autoAbility, err := parseSingleSegmentResourceQuery("auto-ability", ability, queryParam, cfg.l.AutoAbilities)
		if err != nil {
			return nil, err
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

	return resources, nil
}

func (cfg *Config) getMonstersRonsoRage(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := "ronso-rage"
	const ronsoRageOffset int32 = 35
	const ronsoRageLimit int32 = ronsoRageOffset + 12
	query := r.URL.Query().Get(queryParam)
	var ronsoRageID int32

	if query == "" {
		return inputMons, nil
	}

	id, err := strconv.Atoi(query)
	if err == nil {
		ronsoRageID = int32(id) + ronsoRageOffset
	}

	if ronsoRageID == 0 {
		ronsoRage, err := parseNameVersionResourceQuery("ronso-rage", query, "", queryParam, cfg.l.Overdrives)
		if err != nil {
			return nil, err
		}
		ronsoRageID = ronsoRage.ID
	}

	if ronsoRageID > ronsoRageLimit || ronsoRageID <= ronsoRageOffset {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided ronso rage ID %d in %s is out of range. Max ID: 12", id, queryParam), err)
	}

	dbMons, err := cfg.db.GetMonstersByRonsoRageID(r.Context(), ronsoRageID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by ronso rage", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersLocation(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := "location"
	location, isEmpty, err := parseUniqueNameQuery(r, queryParam, cfg.l.Locations)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputMons, nil
	}

	dbMons, err := cfg.db.GetLocationMonsters(r.Context(), location.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by location", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.GetLocationMonstersRow) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersSubLocation(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := "sublocation"
	sublocation, isEmpty, err := parseUniqueNameQuery(r, queryParam, cfg.l.SubLocations)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputMons, nil
	}

	dbMons, err := cfg.db.GetSublocationMonsters(r.Context(), sublocation.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by sublocation", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.GetSublocationMonstersRow) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersArea(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := "area"
	areaID, isEmpty, err := parseIDBasedQuery(r, queryParam, len(cfg.l.Areas))
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputMons, nil
	}

	dbMons, err := cfg.db.GetAreaMonsters(r.Context(), areaID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by area", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.GetAreaMonstersRow) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersDistance(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
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

	return resources, nil
}

func (cfg *Config) getMonstersStoryBased(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := "story-based"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputMons, nil
	}

	dbMons, err := cfg.db.GetMonstersByIsStoryBased(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve story-based monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersRepeatable(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := "repeatable"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputMons, nil
	}

	dbMons, err := cfg.db.GetMonstersByIsRepeatable(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve repeatable monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersCanBeCaptured(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := "capture"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputMons, nil
	}

	dbMons, err := cfg.db.GetMonstersByCanBeCaptured(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters that can be captured", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersHasOverdrive(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := "has-overdrive"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputMons, nil
	}

	dbMons, err := cfg.db.GetMonstersByHasOverdrive(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters with overdrive gauge", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersUnderwater(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := "underwater"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputMons, nil
	}

	dbMons, err := cfg.db.GetMonstersByIsUnderwater(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve underwater monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersZombie(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := "zombie"
	b, isEmpty, err := parseBooleanQuery(r, queryParam)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputMons, nil
	}

	dbMons, err := cfg.db.GetMonstersByIsZombie(r.Context(), b)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve zombie monsters", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersSpecies(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := "species"
	enum, isEmpty, err := parseTypeQuery(r, queryParam, cfg.t.MonsterSpecies)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputMons, nil
	}

	species := database.MonsterSpecies(enum.Name)

	dbMons, err := cfg.db.GetMonstersBySpecies(r.Context(), species)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by species", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersCreationArea(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := "creation-area"
	enum, isEmpty, err := parseTypeQuery(r, queryParam, cfg.t.MaCreationArea)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputMons, nil
	}

	area := h.NullMaCreationArea(&enum.Name)

	dbMons, err := cfg.db.GetMonstersByMaCreationArea(r.Context(), area)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't retrieve monsters by creation area", err)
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	return resources, nil
}

func (cfg *Config) getMonstersType(r *http.Request, inputMons []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := "type"
	enum, isEmpty, err := parseTypeQuery(r, queryParam, cfg.t.CTBIconType)
	if err != nil {
		return nil, err
	}
	if isEmpty {
		return inputMons, nil
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

	return resources, nil
}
