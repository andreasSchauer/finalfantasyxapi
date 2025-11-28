package seeding

import (
	"fmt"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Lookupable interface {
	ToKeyFields() []any
	error
}

func createLookupKey(l Lookupable) string {
	fields := l.ToKeyFields()
	return combineFields(fields)
}


type Lookup struct {
	currentAbility     Ability           // currentAbility and currentBI are
	currentBI          BattleInteraction // used for seeding of ability damage
	currentME          MonsterEquipment  // used for some monster equipment junctions
	abilities          map[string]Ability
	enemyAbilities     map[string]EnemyAbility
	overdriveAbilities map[string]OverdriveAbility
	playerAbilities    map[string]PlayerAbility
	triggerCommands    map[string]TriggerCommand
	aeons              map[string]Aeon
	aeonCommands       map[string]AeonCommand
	affinities         map[string]Affinity
	areas              map[string]Area
	autoAbilities      map[string]AutoAbility
	celestialWeapons   map[string]CelestialWeapon
	characters         map[string]Character
	charClasses        map[string]CharacterClass
	elements           map[string]Element
	equipmentNames     map[string]EquipmentName
	equipmentTables    map[string]EquipmentTable
	encounterLocations map[string]EncounterLocation
	items              map[string]Item
	keyItems           map[string]KeyItem
	masterItems        map[string]MasterItem
	mixes              map[string]Mix
	modifiers          map[string]Modifier
	monsters           map[string]Monster
	overdriveModes     map[string]OverdriveMode
	overdriveCommands  map[string]OverdriveCommand
	overdrives         map[string]Overdrive
	positions          map[string]BlitzballPosition
	properties         map[string]Property
	quests             map[string]Quest
	sidequests         map[string]Sidequest
	subquests          map[string]Subquest
	shops              map[string]Shop
	songs              map[string]Song
	stats              map[string]Stat
	statusConditions   map[string]StatusCondition
	submenus           map[string]Submenu
	treasures          map[string]Treasure
}


func lookupInit() Lookup {
	return Lookup{
		abilities:          make(map[string]Ability),
		enemyAbilities:     make(map[string]EnemyAbility),
		overdriveAbilities: make(map[string]OverdriveAbility),
		playerAbilities:    make(map[string]PlayerAbility),
		triggerCommands:    make(map[string]TriggerCommand),
		aeons:              make(map[string]Aeon),
		aeonCommands:       make(map[string]AeonCommand),
		affinities:         make(map[string]Affinity),
		areas:              make(map[string]Area),
		autoAbilities:      make(map[string]AutoAbility),
		celestialWeapons:   make(map[string]CelestialWeapon),
		characters:         make(map[string]Character),
		charClasses:        make(map[string]CharacterClass),
		elements:           make(map[string]Element),
		equipmentNames:     make(map[string]EquipmentName),
		equipmentTables:    make(map[string]EquipmentTable),
		encounterLocations: make(map[string]EncounterLocation),
		items:              make(map[string]Item),
		keyItems:           make(map[string]KeyItem),
		masterItems:        make(map[string]MasterItem),
		mixes:              make(map[string]Mix),
		modifiers:          make(map[string]Modifier),
		monsters:           make(map[string]Monster),
		overdriveModes:     make(map[string]OverdriveMode),
		overdriveCommands:  make(map[string]OverdriveCommand),
		overdrives:         make(map[string]Overdrive),
		positions:          make(map[string]BlitzballPosition),
		properties:         make(map[string]Property),
		quests:             make(map[string]Quest),
		sidequests:         make(map[string]Sidequest),
		subquests:          make(map[string]Subquest),
		shops:              make(map[string]Shop),
		songs:              make(map[string]Song),
		stats:              make(map[string]Stat),
		statusConditions:   make(map[string]StatusCondition),
		submenus:           make(map[string]Submenu),
		treasures:          make(map[string]Treasure),
	}
}


func getResource[T any, K any](key K, lookup map[string]T) (T, error) {
	switch k := any(key).(type) {
    case string:
        return getResourceStrKey(k, lookup)
    case Lookupable:
		return getResourceObjKey(k, lookup)
    default:
        var zeroType T
        return zeroType, fmt.Errorf("key must be either string or Lookupable, got %T", key)
    }
}


func getResourceStrKey[T any](key string, lookup map[string]T) (T, error) {
	resource, found := lookup[key]
	if !found {
		var zeroType T
		return zeroType, h.GetErr(key, fmt.Errorf("couldn't find %s", getTypeName[T]()))
	}

	return resource, nil
}


func getResourceObjKey[T any](obj Lookupable, lookup map[string]T) (T, error) {
	key := createLookupKey(obj)

	resource, err := getResource(key, lookup)
	if err != nil {
		var zeroType T
		return zeroType, h.GetErr(obj.Error(), fmt.Errorf("couldn't find %s", getTypeName[T]()))
	}

	return resource, nil
}


func getTypeName[T any]() string {
	var zeroType T
	typeString := fmt.Sprintf("%T", zeroType)
	typeOnly := strings.Split(typeString, ".")

	return typeOnly[len(typeOnly)-1]
}