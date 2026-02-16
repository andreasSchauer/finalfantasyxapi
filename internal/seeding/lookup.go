package seeding

import (
	"fmt"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

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
	currentAbility       Ability           // currentAbility and currentBI are
	currentBI            BattleInteraction // used for seeding of ability damage
	currentME            MonsterEquipment  // used for some monster equipment junctions
	Abilities            map[string]Ability
	AbilitiesID          map[int32]Ability
	EnemyAbilities       map[string]EnemyAbility
	EnemyAbilitiesID     map[int32]EnemyAbility
	GenericAbilities	 map[string]GenericAbility
	GenericAbilitiesID	 map[int32]GenericAbility
	OverdriveAbilities   map[string]OverdriveAbility
	OverdriveAbilitiesID map[int32]OverdriveAbility
	PlayerAbilities      map[string]PlayerAbility
	PlayerAbilitiesID    map[int32]PlayerAbility
	TriggerCommands      map[string]TriggerCommand
	TriggerCommandsID    map[int32]TriggerCommand
	Aeons                map[string]Aeon
	AeonsID              map[int32]Aeon
	AeonCommands         map[string]AeonCommand
	AeonCommandsID       map[int32]AeonCommand
	Affinities           map[string]Affinity
	AffinitiesID         map[int32]Affinity
	ArenaCreations       map[string]ArenaCreation
	ArenaCreationsID     map[int32]ArenaCreation
	Locations            map[string]Location
	LocationsID          map[int32]Location
	Sublocations         map[string]Sublocation
	SublocationsID       map[int32]Sublocation
	Areas                map[string]Area
	AreasID              map[int32]Area
	AutoAbilities        map[string]AutoAbility
	AutoAbilitiesID      map[int32]AutoAbility
	CelestialWeapons     map[string]CelestialWeapon
	CelestialWeaponsID   map[int32]CelestialWeapon
	Characters           map[string]Character
	CharactersID         map[int32]Character
	CharClasses          map[string]CharacterClass
	CharClassesID        map[int32]CharacterClass
	Elements             map[string]Element
	ElementsID           map[int32]Element
	ElementalResists     map[string]ElementalResist
	ElementalResistsID   map[int32]ElementalResist
	EquipmentNames       map[string]EquipmentName
	EquipmentNamesID     map[int32]EquipmentName
	EquipmentTables      map[string]EquipmentTable
	EquipmentTablesID    map[int32]EquipmentTable
	EncounterAreas       map[string]EncounterArea
	EncounterAreasID 	 map[int32]EncounterArea
	FMVs                 map[string]FMV
	FMVsID               map[int32]FMV
	Items                map[string]Item
	ItemsID              map[int32]Item
	KeyItems             map[string]KeyItem
	KeyItemsID           map[int32]KeyItem
	MasterItems          map[string]MasterItem
	MasterItemsID        map[int32]MasterItem
	Mixes                map[string]Mix
	MixesID              map[int32]Mix
	Modifiers            map[string]Modifier
	ModifiersID          map[int32]Modifier
	Monsters             map[string]Monster
	MonstersID           map[int32]Monster
	MonsterFormations    map[string]MonsterFormation
	MonsterFormationsID  map[int32]MonsterFormation
	OverdriveModes       map[string]OverdriveMode
	OverdriveModesID     map[int32]OverdriveMode
	OverdriveCommands    map[string]OverdriveCommand
	OverdriveCommandsID  map[int32]OverdriveCommand
	Overdrives           map[string]Overdrive
	OverdrivesID         map[int32]Overdrive
	Positions            map[string]BlitzballPosition
	PositionsID          map[int32]BlitzballPosition
	Properties           map[string]Property
	PropertiesID         map[int32]Property
	Quests               map[string]Quest
	QuestsID             map[int32]Quest
	RonsoRages           map[string]RonsoRage
	RonsoRagesID         map[int32]RonsoRage
	Sidequests           map[string]Sidequest
	SidequestsID         map[int32]Sidequest
	Subquests            map[string]Subquest
	SubquestsID          map[int32]Subquest
	Shops                map[string]Shop
	ShopsID              map[int32]Shop
	Songs                map[string]Song
	CuesID               map[int32]Cue
	BackgroundMusicID    map[int32]BackgroundMusic
	SongsID              map[int32]Song
	Stats                map[string]Stat
	StatsID              map[int32]Stat
	StatusConditions     map[string]StatusCondition
	StatusConditionsID   map[int32]StatusCondition
	Submenus             map[string]Submenu
	SubmenusID           map[int32]Submenu
	Treasures            map[string]Treasure
	TreasuresID          map[int32]Treasure
}

func lookupInit() Lookup {
	return Lookup{
		Abilities:            make(map[string]Ability),
		AbilitiesID:          make(map[int32]Ability),
		GenericAbilities: 	  make(map[string]GenericAbility),
		GenericAbilitiesID:   make(map[int32]GenericAbility),
		EnemyAbilities:       make(map[string]EnemyAbility),
		EnemyAbilitiesID:     make(map[int32]EnemyAbility),
		OverdriveAbilities:   make(map[string]OverdriveAbility),
		OverdriveAbilitiesID: make(map[int32]OverdriveAbility),
		PlayerAbilities:      make(map[string]PlayerAbility),
		PlayerAbilitiesID:    make(map[int32]PlayerAbility),
		TriggerCommands:      make(map[string]TriggerCommand),
		TriggerCommandsID:    make(map[int32]TriggerCommand),
		Aeons:                make(map[string]Aeon),
		AeonsID:              make(map[int32]Aeon),
		AeonCommands:         make(map[string]AeonCommand),
		AeonCommandsID:       make(map[int32]AeonCommand),
		Affinities:           make(map[string]Affinity),
		AffinitiesID:         make(map[int32]Affinity),
		ArenaCreations:       make(map[string]ArenaCreation),
		ArenaCreationsID:     make(map[int32]ArenaCreation),
		Locations:            make(map[string]Location),
		LocationsID:          make(map[int32]Location),
		Sublocations:         make(map[string]Sublocation),
		SublocationsID:       make(map[int32]Sublocation),
		Areas:                make(map[string]Area),
		AreasID:              make(map[int32]Area),
		AutoAbilities:        make(map[string]AutoAbility),
		AutoAbilitiesID:      make(map[int32]AutoAbility),
		CelestialWeapons:     make(map[string]CelestialWeapon),
		CelestialWeaponsID:   make(map[int32]CelestialWeapon),
		Characters:           make(map[string]Character),
		CharactersID:         make(map[int32]Character),
		CharClasses:          make(map[string]CharacterClass),
		CharClassesID:        make(map[int32]CharacterClass),
		Elements:             make(map[string]Element),
		ElementsID:           make(map[int32]Element),
		ElementalResists:     make(map[string]ElementalResist),
		ElementalResistsID:   make(map[int32]ElementalResist),
		EquipmentNames:       make(map[string]EquipmentName),
		EquipmentNamesID:     make(map[int32]EquipmentName),
		EquipmentTables:      make(map[string]EquipmentTable),
		EquipmentTablesID:    make(map[int32]EquipmentTable),
		EncounterAreas:       make(map[string]EncounterArea),
		EncounterAreasID: make(map[int32]EncounterArea),
		FMVs:                 make(map[string]FMV),
		FMVsID:               make(map[int32]FMV),
		Items:                make(map[string]Item),
		ItemsID:              make(map[int32]Item),
		KeyItems:             make(map[string]KeyItem),
		KeyItemsID:           make(map[int32]KeyItem),
		MasterItems:          make(map[string]MasterItem),
		MasterItemsID:        make(map[int32]MasterItem),
		Mixes:                make(map[string]Mix),
		MixesID:              make(map[int32]Mix),
		Modifiers:            make(map[string]Modifier),
		ModifiersID:          make(map[int32]Modifier),
		Monsters:             make(map[string]Monster),
		MonstersID:           make(map[int32]Monster),
		MonsterFormations:    make(map[string]MonsterFormation),
		MonsterFormationsID:  make(map[int32]MonsterFormation),
		OverdriveModes:       make(map[string]OverdriveMode),
		OverdriveModesID:     make(map[int32]OverdriveMode),
		OverdriveCommands:    make(map[string]OverdriveCommand),
		OverdriveCommandsID:  make(map[int32]OverdriveCommand),
		Overdrives:           make(map[string]Overdrive),
		OverdrivesID:         make(map[int32]Overdrive),
		Positions:            make(map[string]BlitzballPosition),
		PositionsID:          make(map[int32]BlitzballPosition),
		Properties:           make(map[string]Property),
		PropertiesID:         make(map[int32]Property),
		Quests:               make(map[string]Quest),
		QuestsID:             make(map[int32]Quest),
		RonsoRages:           make(map[string]RonsoRage),
		RonsoRagesID:         make(map[int32]RonsoRage),
		Sidequests:           make(map[string]Sidequest),
		SidequestsID:         make(map[int32]Sidequest),
		Subquests:            make(map[string]Subquest),
		SubquestsID:          make(map[int32]Subquest),
		Shops:                make(map[string]Shop),
		ShopsID:              make(map[int32]Shop),
		Songs:                make(map[string]Song),
		SongsID:              make(map[int32]Song),
		CuesID:               make(map[int32]Cue),
		BackgroundMusicID:    make(map[int32]BackgroundMusic),
		Stats:                make(map[string]Stat),
		StatsID:              make(map[int32]Stat),
		StatusConditions:     make(map[string]StatusCondition),
		StatusConditionsID:   make(map[int32]StatusCondition),
		Submenus:             make(map[string]Submenu),
		SubmenusID:           make(map[int32]Submenu),
		Treasures:            make(map[string]Treasure),
		TreasuresID:          make(map[int32]Treasure),
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

// maybe make one for all id-only endpoints
// might also make one for some complex resources like abilities, since it will reduce the amount of queries needed
// I could also do that for monsters tbh
// But that is a can of worms I won't be opening now
func GetResourceByID[T h.HasID](id int32, lookup map[int32]T) (T, error) {
	resource, found := lookup[id]
	if !found {
		var zeroType T
		return zeroType, fmt.Errorf("couldn't find %s with id '%d'.", getTypeName[T](), id)
	}

	return resource, nil
}

func getResourceStrKey[T any](key string, lookup map[string]T) (T, error) {
	resource, found := lookup[key]
	if !found {
		var zeroType T
		return zeroType, h.NewErr(key, fmt.Errorf("couldn't find %s '%s'.", getTypeName[T](), key))
	}

	return resource, nil
}

func getResourceObjKey[T any](obj Lookupable, lookup map[string]T) (T, error) {
	key := CreateLookupKey(obj)

	resource, err := GetResource(key, lookup)
	if err != nil {
		var zeroType T
		return zeroType, h.NewErr(obj.Error(), fmt.Errorf("couldn't find %s.", getTypeName[T]()))
	}

	return resource, nil
}

func getTypeName[T any]() string {
	var zeroType T
	typeString := fmt.Sprintf("%T", zeroType)
	typeOnly := strings.Split(typeString, ".")

	return typeOnly[len(typeOnly)-1]
}
