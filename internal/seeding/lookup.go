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

func CreateLookupKey(l Lookupable) string {
	fields := l.ToKeyFields()
	return combineFields(fields)
}

type LookupObject struct {
	Name    string
	Version *int32
}

func (l LookupObject) ToKeyFields() []any {
	return []any{
		l.Name,
		h.DerefOrNil(l.Version),
	}
}

func (l LookupObject) Error() string {
	return fmt.Sprintf("lookup object name: %s, version %d", l.Name, h.DerefOrNil(l.Version))
}

type Lookup struct {
	currentAbility     Ability           // currentAbility and currentBI are
	currentBI          BattleInteraction // used for seeding of ability damage
	currentME          MonsterEquipment  // used for some monster equipment junctions
	Abilities          map[string]Ability
	EnemyAbilities     map[string]EnemyAbility
	OverdriveAbilities map[string]OverdriveAbility
	PlayerAbilities    map[string]PlayerAbility
	TriggerCommands    map[string]TriggerCommand
	Aeons              map[string]Aeon
	AeonCommands       map[string]AeonCommand
	Affinities         map[string]Affinity
	Locations		   map[string]Location
	SubLocations	   map[string]SubLocation
	Areas              map[string]Area
	AutoAbilities      map[string]AutoAbility
	CelestialWeapons   map[string]CelestialWeapon
	Characters         map[string]Character
	CharClasses        map[string]CharacterClass
	Elements           map[string]Element
	ElementalResists   map[string]ElementalResist
	EquipmentNames     map[string]EquipmentName
	EquipmentTables    map[string]EquipmentTable
	EncounterLocations map[string]EncounterLocation
	FMVs			   map[string]FMV
	Items              map[string]Item
	KeyItems           map[string]KeyItem
	MasterItems        map[string]MasterItem
	Mixes              map[string]Mix
	Modifiers          map[string]Modifier
	Monsters           map[string]Monster
	OverdriveModes     map[string]OverdriveMode
	OverdriveCommands  map[string]OverdriveCommand
	Overdrives         map[string]Overdrive
	Positions          map[string]BlitzballPosition
	Properties         map[string]Property
	Quests             map[string]Quest
	Sidequests         map[string]Sidequest
	Subquests          map[string]Subquest
	Shops              map[string]Shop
	Songs              map[string]Song
	Stats              map[string]Stat
	StatusConditions   map[string]StatusCondition
	Submenus           map[string]Submenu
	Treasures          map[string]Treasure
}

func lookupInit() Lookup {
	return Lookup{
		Abilities:          make(map[string]Ability),
		EnemyAbilities:     make(map[string]EnemyAbility),
		OverdriveAbilities: make(map[string]OverdriveAbility),
		PlayerAbilities:    make(map[string]PlayerAbility),
		TriggerCommands:    make(map[string]TriggerCommand),
		Aeons:              make(map[string]Aeon),
		AeonCommands:       make(map[string]AeonCommand),
		Affinities:         make(map[string]Affinity),
		Locations: 			make(map[string]Location),
		SubLocations: 		make(map[string]SubLocation),
		Areas:              make(map[string]Area),
		AutoAbilities:      make(map[string]AutoAbility),
		CelestialWeapons:   make(map[string]CelestialWeapon),
		Characters:         make(map[string]Character),
		CharClasses:        make(map[string]CharacterClass),
		Elements:           make(map[string]Element),
		ElementalResists: 	make(map[string]ElementalResist),
		EquipmentNames:     make(map[string]EquipmentName),
		EquipmentTables:    make(map[string]EquipmentTable),
		EncounterLocations: make(map[string]EncounterLocation),
		FMVs:				make(map[string]FMV),
		Items:              make(map[string]Item),
		KeyItems:           make(map[string]KeyItem),
		MasterItems:        make(map[string]MasterItem),
		Mixes:              make(map[string]Mix),
		Modifiers:          make(map[string]Modifier),
		Monsters:           make(map[string]Monster),
		OverdriveModes:     make(map[string]OverdriveMode),
		OverdriveCommands:  make(map[string]OverdriveCommand),
		Overdrives:         make(map[string]Overdrive),
		Positions:          make(map[string]BlitzballPosition),
		Properties:         make(map[string]Property),
		Quests:             make(map[string]Quest),
		Sidequests:         make(map[string]Sidequest),
		Subquests:          make(map[string]Subquest),
		Shops:              make(map[string]Shop),
		Songs:              make(map[string]Song),
		Stats:              make(map[string]Stat),
		StatusConditions:   make(map[string]StatusCondition),
		Submenus:           make(map[string]Submenu),
		Treasures:          make(map[string]Treasure),
	}
}

func GetResource[T any, K any](key K, lookup map[string]T) (T, error) {
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
	key := CreateLookupKey(obj)

	resource, err := GetResource(key, lookup)
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
