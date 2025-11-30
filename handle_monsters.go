package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// species and ctb icon type need to be named api resources, because they have a type endpoint
type Monster struct {
	ID                   int32              `json:"id"`
	Name                 string             `json:"name"`
	Version              *int32             `json:"version,omitempty"`
	Specification        *string            `json:"specification,omitempty"`
	Notes                *string            `json:"notes,omitempty"`
	Species              string             `json:"species"`
	IsStoryBased         bool               `json:"is_story_based"`
	CanBeCaptured        bool               `json:"can_be_captured"`
	AreaConquestLocation *string            `json:"area_conquest_location,omitempty"`
	CTBIconType          string             `json:"ctb_icon_type"`
	HasOverdrive         bool               `json:"has_overdrive"`
	IsUnderwater         bool               `json:"is_underwater"`
	IsZombie             bool               `json:"is_zombie"`
	Distance             int32              `json:"distance"`
	//Properties           []NamedAPIResource `json:"properties"`
	//AutoAbilities        []NamedAPIResource `json:"auto_abilities"`
	AP                   int32              `json:"ap"`
	APOverkill           int32              `json:"ap_overkill"`
	OverkillDamage       int32              `json:"overkill_damage"`
	Gil                  int32              `json:"gil"`
	StealGil             *int32             `json:"steal_gil"`
	//RonsoRages           []NamedAPIResource `json:"ronso_rages"`
	DoomCountdown        *int32             `json:"doom_countdown"`
	PoisonRate           *float32           `json:"poison_rate"`
	ThreatenChance       *int32             `json:"threaten_chance"`
	ZanmatoLevel         int32              `json:"zanmato_level"`
	MonsterArenaPrice    *int32             `json:"monster_arena_price,omitempty"`
	SensorText           *string            `json:"sensor_text"`
	ScanText             *string            `json:"scan_text"`
	//BaseStats            []BaseStat         `json:"base_stats"`
	Items                *MonsterItems      `json:"items"`
	Equipment            *MonsterEquipment  `json:"equipment"`
	//ElemResists          []ElementalResist `json:"elem_resists"`
	//StatusImmunities []NamedAPIResource `json:"status_immunities"`
	//StatusResists        []StatusResist    `json:"status_resists"`
	//AlteredStates        []AlteredState    `json:"altered_states"`
	//Abilities            []MonsterAbility  `json:"abilities"`
}

type BaseStat struct {
	Stat  NamedAPIResource `json:"stat"`
	Value int32
}


type MonsterEquipment struct {
	DropChance        int32                 `json:"drop_chance"`
	Power             int32                 `json:"power"`
	CriticalPlus      int32                 `json:"critical_plus"`
	AbilitySlots      MonsterEquipmentSlots `json:"ability_slots"`
	AttachedAbilities MonsterEquipmentSlots `json:"attached_abilities"`
	WeaponAbilities   []EquipmentDrop       `json:"weapon_abilities"`
	ArmorAbilities    []EquipmentDrop       `json:"armor_abilities"`
}

func (me MonsterEquipment) IsZero() bool {
	return me.DropChance == 0
}


type MonsterEquipmentSlots struct {
	MinAmount int32                  `json:"min_amount"`
	MaxAmount int32                  `json:"max_amount"`
	Chances   []EquipmentSlotsChance `json:"chances"`
}

type EquipmentSlotsChance struct {
	Amount int32 `json:"amount"`
	Chance int32 `json:"chance"`
}


type EquipmentDrop struct {
	AutoAbility  NamedAPIResource   	`json:"auto_ability"`
	ForcedChars  []NamedAPIResource 		`json:"forced_characters"`
	IsForced    bool     				`json:"is_forced"`
	Probability *int32   				`json:"probability,omitempty"`
}

func (cfg *apiConfig) handleMonsters(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/monsters/")
	segments := strings.Split(path, "/")

	if path == "" {
		cfg.handleMonstersRetrieve(w, r)
		return
	}

	switch len(segments) {
	case 1:
		// /api/monsters/{name or id}
		segment := segments[0]

		input, err := parseSingleSegmentResource(segment, cfg.l.Monsters)
		if handleHTTPError(w, err) {
			return
		}

		cfg.handleMonsterGet(w, r, input)
		return
	case 2:
		// /api/monsters/{name}/{version}

		name := segments[0]
		versionStr := segments[1]

		input, err := parseNameVersionResource(name, versionStr, cfg.l.Monsters)
		if handleHTTPError(w, err) {
			return
		}

		cfg.handleMonsterGet(w, r, input)
		return
	default:
		respondWithError(w, http.StatusBadRequest, `Wrong format. Usage: /api/monsters/{name or id}, or /api/monsters/{name}/{version}`, nil)
		return
	}
}

func (cfg *apiConfig) handleMonsterGet(w http.ResponseWriter, r *http.Request, input parseResponse) {
	if input.Name != "" {
		dbMons, err := cfg.db.GetMonstersByName(r.Context(), input.Name)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Couldn't get multiple Monsters", err)
			return
		}

		resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
			return mon.ID, mon.Name, &mon.Version.Int32, &mon.Specification.String
		})

		resourceList, err := cfg.newNamedAPIResourceList(r, resources)
		if handleHTTPError(w, err) {
			return
		}

		respondWithJSON(w, http.StatusMultipleChoices, resourceList)
		return
	}

	dbMonster, err := cfg.db.GetMonster(r.Context(), input.ID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get Monster. Monster with this ID doesn't exist.", err)
		return
	}

	monsterItems, err := cfg.getMonsterItems(r, dbMonster)
	if handleHTTPError(w, err) {
		return
	}

	monsterEquipment, err := cfg.getMonsterEquipment(r, dbMonster)
	if handleHTTPError(w, err) {
		return
	}

	response := Monster{
		ID:                   	dbMonster.ID,
		Name:                 	dbMonster.Name,
		Version:              	h.NullInt32ToPtr(dbMonster.Version),
		Specification:        	h.NullStringToPtr(dbMonster.Specification),
		Notes:                	h.NullStringToPtr(dbMonster.Notes),
		Species:              	string(dbMonster.Species),
		IsStoryBased:         	dbMonster.IsStoryBased,
		CanBeCaptured:        	dbMonster.CanBeCaptured,
		AreaConquestLocation: 	h.ConvertNullMaCreationArea(dbMonster.AreaConquestLocation),
		CTBIconType:          	string(dbMonster.CtbIconType),
		HasOverdrive:         	dbMonster.HasOverdrive,
		IsUnderwater:         	dbMonster.IsUnderwater,
		IsZombie:             	dbMonster.IsZombie,
		Distance:             	anyToInt32(dbMonster.Distance),
		AP:                   	dbMonster.Ap,
		APOverkill:           	dbMonster.ApOverkill,
		OverkillDamage:       	dbMonster.OverkillDamage,
		Gil:                  	dbMonster.Gil,
		StealGil:             	h.NullInt32ToPtr(dbMonster.StealGil),
		DoomCountdown:       	anyToInt32Ptr(dbMonster.DoomCountdown),
		PoisonRate:           	anyToFloat32Ptr(dbMonster.PoisonRate),
		ThreatenChance:       	anyToInt32Ptr(dbMonster.ThreatenChance),
		ZanmatoLevel:         	anyToInt32(dbMonster.ZanmatoLevel),
		MonsterArenaPrice:    	h.NullInt32ToPtr(dbMonster.MonsterArenaPrice),
		SensorText:           	h.NullStringToPtr(dbMonster.SensorText),
		ScanText:            	h.NullStringToPtr(dbMonster.ScanText),
		Items: 				  	h.NilOrPtr(monsterItems),
		Equipment: 				h.NilOrPtr(monsterEquipment),
	}

	respondWithJSON(w, http.StatusOK, response)
}


func (cfg *apiConfig) getMonsterEquipment(r *http.Request, mon database.Monster) (MonsterEquipment, error) {
	dbEquipment, err := cfg.db.GetMonsterEquipment(r.Context(), mon.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return MonsterEquipment{}, nil
		}
		return MonsterEquipment{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't retrieve Equipment of Monster %s Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	abilitySlots, attachedAbilities, err := cfg.getMonsterEquipmentSlots(r, mon, dbEquipment)
	if err != nil {
		return MonsterEquipment{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't retrieve Equipment Slots of Monster %s Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}
	
	weaponAbilities, err := cfg.getEquipmentDrops(r, mon, dbEquipment, database.EquipTypeWeapon)
	if err != nil {
		return MonsterEquipment{}, err
	}
	
	armorAbilities, err := cfg.getEquipmentDrops(r, mon, dbEquipment, database.EquipTypeArmor)
	if err != nil {
		return MonsterEquipment{}, err
	}

	return MonsterEquipment{
		DropChance: 		anyToInt32(dbEquipment.DropChance),
		Power: 				anyToInt32(dbEquipment.Power),
		CriticalPlus: 		dbEquipment.CriticalPlus,
		AbilitySlots: 		abilitySlots,
		AttachedAbilities: 	attachedAbilities,
		WeaponAbilities: 	weaponAbilities,
		ArmorAbilities: 	armorAbilities,
	}, nil
}


func (cfg *apiConfig) getEquipmentDrops(r *http.Request, mon database.Monster, equipment database.MonsterEquipment, equipType database.EquipType) ([]EquipmentDrop, error) {
	dbDrops, err := cfg.db.GetMonsterEquipmentAbilities(r.Context(), database.GetMonsterEquipmentAbilitiesParams{
		MonsterEquipmentID: equipment.ID,
		Type: 				equipType,	
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't retrieve dropped %s abilities of Monster %s Version %d", string(equipType), mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	drops := []EquipmentDrop{}

	for _, dbDrop := range dbDrops {
		forcedChars, err := cfg.getEquipmentDropForcedChars(r, mon, equipment, dbDrop)
		if err != nil {
			return nil, err
		}
		autoAbility := cfg.newNamedAPIResourceSimple("auto-abilities", dbDrop.AutoAbilityID.Int32, dbDrop.AutoAbility.String)

		drop := EquipmentDrop{
			AutoAbility: 	autoAbility,
			ForcedChars: 	forcedChars,
			IsForced: 		dbDrop.IsForced.Bool,
			Probability: 	anyToInt32Ptr(dbDrop.Probability),
		}

		drops = append(drops, drop)
	}

	return drops, nil
}


func (cfg *apiConfig) getEquipmentDropForcedChars(r *http.Request, mon database.Monster, equipment database.MonsterEquipment, drop database.GetMonsterEquipmentAbilitiesRow) ([]NamedAPIResource, error) {
	dbChars, err := cfg.db.GetEquipmentDropCharacters(r.Context(), database.GetEquipmentDropCharactersParams{
		MonsterEquipmentID: equipment.ID,
		EquipmentDropID: 	drop.ID.Int32,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't retrieve characters of auto ability %s dropped by monster %s version %d", drop.AutoAbility.String, mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	characters := createNamedAPIResourcesSimple(cfg, dbChars, "characters", func(char database.GetEquipmentDropCharactersRow) (int32, string) {
		return h.NullInt32ToVal(char.CharacterID), h.NullStringToVal(char.CharacterName)
	})

	return characters, nil
}


func (cfg *apiConfig) getMonsterEquipmentSlots(r *http.Request, mon database.Monster, equipment database.MonsterEquipment) (MonsterEquipmentSlots, MonsterEquipmentSlots, error) {
	dbEquipmentSlots, err := cfg.db.GetMonsterEquipmentSlots(r.Context(), equipment.ID)
	if err != nil {
		return MonsterEquipmentSlots{}, MonsterEquipmentSlots{},newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't retrieve Equipment Slots of Monster %s Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	dbAbilitySlots := dbEquipmentSlots[0]
	dbAttachedAbilities := dbEquipmentSlots[1]

	abilitySlots, err := cfg.assembleMonsterEquipmentSlots(r, mon, equipment, dbAbilitySlots)
	if err != nil {
		return MonsterEquipmentSlots{}, MonsterEquipmentSlots{}, err
	}

	attachedAbilities, err := cfg.assembleMonsterEquipmentSlots(r, mon, equipment, dbAttachedAbilities)
	if err != nil {
		return MonsterEquipmentSlots{}, MonsterEquipmentSlots{}, err
	}

	return abilitySlots, attachedAbilities, nil
}


func (cfg *apiConfig) assembleMonsterEquipmentSlots(r *http.Request,mon database.Monster, equipment database.MonsterEquipment, dbEquipmentSlots database.MonsterEquipmentSlot) (MonsterEquipmentSlots, error) {
	equipmentSlotsChances, err := cfg.getMonsterEquipmentSlotsChances(r, mon, equipment, dbEquipmentSlots)
	if err != nil {
		return MonsterEquipmentSlots{}, err
	}

	equipmentSlots := MonsterEquipmentSlots{
		MinAmount: 	anyToInt32(dbEquipmentSlots.MinAmount),
		MaxAmount: 	anyToInt32(dbEquipmentSlots.MaxAmount),
		Chances: 	equipmentSlotsChances,
	}

	return equipmentSlots, nil
}


func (cfg *apiConfig) getMonsterEquipmentSlotsChances(r *http.Request, mon database.Monster, equipment database.MonsterEquipment, slots database.MonsterEquipmentSlot) ([]EquipmentSlotsChance, error) {
	dbSlotsChances, err := cfg.db.GetMonsterEquipmentSlotsChances(r.Context(), database.GetMonsterEquipmentSlotsChancesParams{
		MonsterEquipmentID: equipment.ID,
		EquipmentSlotsID: 	slots.ID,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't assemble Equipment Slots of Monster %s Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	chances := []EquipmentSlotsChance{}

	for _, dbChance := range dbSlotsChances {
		chance := EquipmentSlotsChance{
			Amount: anyToInt32(dbChance.Amount),
			Chance: anyToInt32(dbChance.Chance),
		}

		chances = append(chances, chance)
	}

	return chances, nil
}


func (cfg *apiConfig) handleMonstersRetrieve(w http.ResponseWriter, r *http.Request) {
	dbMons, err := cfg.db.GetMonsters(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve monsters", err)
		return
	}

	resources := createNamedAPIResources(cfg, dbMons, "monsters", func(mon database.Monster) (int32, string, *int32, *string) {
		return mon.ID, mon.Name, h.NullInt32ToPtr(mon.Version), h.NullStringToPtr(mon.Specification)
	})

	resourceList, err := cfg.newNamedAPIResourceList(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resourceList)
}
