package seeding

import (
	"errors"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Lookupable interface {
	ToKeyFields() []any
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

func (l *Lookup) getAbility(abilityRef AbilityReference) (Ability, error) {
	key := createLookupKey(abilityRef)

	ability, found := l.abilities[key]
	if !found {
		return Ability{}, h.GetErr(abilityRef.Error(), errors.New("couldn't find ability"))
	}

	return ability, nil
}

func (l *Lookup) getEnemyAbility(abilityRef AbilityReference) (EnemyAbility, error) {
	key := createLookupKey(abilityRef)

	ability, found := l.enemyAbilities[key]
	if !found {
		return EnemyAbility{}, h.GetErr(abilityRef.Error(), errors.New("couldn't find enemy ability"))
	}

	return ability, nil
}

func (l *Lookup) getOverdriveAbility(abilityRef AbilityReference) (OverdriveAbility, error) {
	key := createLookupKey(abilityRef)

	ability, found := l.overdriveAbilities[key]
	if !found {
		return OverdriveAbility{}, h.GetErr(abilityRef.Error(), errors.New("couldn't find overdrive ability"))
	}

	return ability, nil
}

func (l *Lookup) getPlayerAbility(abilityRef AbilityReference) (PlayerAbility, error) {
	key := createLookupKey(abilityRef)

	ability, found := l.playerAbilities[key]
	if !found {
		return PlayerAbility{}, h.GetErr(abilityRef.Error(), errors.New("couldn't find player ability"))
	}

	return ability, nil
}

func (l *Lookup) getTriggerCommand(abilityRef AbilityReference) (TriggerCommand, error) {
	key := createLookupKey(abilityRef)

	command, found := l.triggerCommands[key]
	if !found {
		return TriggerCommand{}, h.GetErr(abilityRef.Error(), errors.New("couldn't find trigger command"))
	}

	return command, nil
}

func (l *Lookup) getAeon(key string) (Aeon, error) {
	playerUnit := PlayerUnit{
		Name: key,
		Type: database.UnitTypeAeon,
	}
	lookupKey := createLookupKey(playerUnit)

	aeon, found := l.aeons[lookupKey]
	if !found {
		return Aeon{}, h.GetErr(key, errors.New("couldn't find aeon"))
	}

	return aeon, nil
}

func (l *Lookup) getAeonCommand(key string) (AeonCommand, error) {
	command, found := l.aeonCommands[key]
	if !found {
		return AeonCommand{}, h.GetErr(key, errors.New("couldn't find aeon command"))
	}

	return command, nil
}

func (l *Lookup) getAffinity(key string) (Affinity, error) {
	affinity, found := l.affinities[key]
	if !found {
		return Affinity{}, h.GetErr(key, errors.New("couldn't find affinity"))
	}

	return affinity, nil
}

func (l *Lookup) getArea(locationArea LocationArea) (Area, error) {
	key := createLookupKey(locationArea)

	area, found := l.areas[key]
	if !found {
		return Area{}, h.GetErr(locationArea.Error(), errors.New("couldn't find location area"))
	}

	return area, nil
}

func (l *Lookup) getAutoAbility(key string) (AutoAbility, error) {
	autoAbility, found := l.autoAbilities[key]
	if !found {
		return AutoAbility{}, h.GetErr(key, errors.New("couldn't find auto-ability"))
	}

	return autoAbility, nil
}

func (l *Lookup) getCelestialWeapon(key string) (CelestialWeapon, error) {
	weapon, found := l.celestialWeapons[key]
	if !found {
		return CelestialWeapon{}, h.GetErr(key, errors.New("couldn't find celestial weapon"))
	}

	return weapon, nil
}

func (l *Lookup) getCharacter(key string) (Character, error) {
	playerUnit := PlayerUnit{
		Name: key,
		Type: database.UnitTypeCharacter,
	}
	lookupKey := createLookupKey(playerUnit)

	character, found := l.characters[lookupKey]
	if !found {
		return Character{}, h.GetErr(key, errors.New("couldn't find couldn't find character"))
	}

	return character, nil
}

func (l *Lookup) getCharacterClass(key string) (CharacterClass, error) {
	class, found := l.charClasses[key]
	if !found {
		return CharacterClass{}, h.GetErr(key, errors.New("couldn't find character class"))
	}

	return class, nil
}

func (l *Lookup) getElement(key string) (Element, error) {
	element, found := l.elements[key]
	if !found {
		return Element{}, h.GetErr(key, errors.New("couldn't find element"))
	}

	return element, nil
}

func (l *Lookup) getEquipmentName(key string) (EquipmentName, error) {
	equipment, found := l.equipmentNames[key]
	if !found {
		return EquipmentName{}, h.GetErr(key, errors.New("couldn't find equipment name"))
	}

	return equipment, nil
}

func (l *Lookup) getEquipmentTable(key string) (EquipmentTable, error) {
	equipmentTable, found := l.equipmentTables[key]
	if !found {
		return EquipmentTable{}, h.GetErr(key, errors.New("couldn't find equipment table"))
	}

	return equipmentTable, nil
}

func (l *Lookup) getEncounterLocation(key string) (EncounterLocation, error) {
	encounterLocation, found := l.encounterLocations[key]
	if !found {
		return EncounterLocation{}, h.GetErr(key, errors.New("couldn't find encounter location"))
	}

	return encounterLocation, nil
}

func (l *Lookup) getItem(key string) (Item, error) {
	masterItem := MasterItem{
		Name: key,
		Type: database.ItemTypeItem,
	}
	lookupKey := createLookupKey(masterItem)

	item, found := l.items[lookupKey]
	if !found {
		return Item{}, h.GetErr(key, errors.New("couldn't find item"))
	}

	return item, nil
}

func (l *Lookup) getKeyItem(key string) (KeyItem, error) {
	masterItem := MasterItem{
		Name: key,
		Type: database.ItemTypeKeyItem,
	}
	lookupKey := createLookupKey(masterItem)

	keyItem, found := l.keyItems[lookupKey]
	if !found {
		return KeyItem{}, h.GetErr(key, errors.New("couldn't find key item"))
	}

	return keyItem, nil
}

func (l *Lookup) getMasterItem(key string) (MasterItem, error) {
	masterItem, found := l.masterItems[key]
	if !found {
		return MasterItem{}, h.GetErr(key, errors.New("couldn't find master item"))
	}

	return masterItem, nil
}

func (l *Lookup) getMix(key string) (Mix, error) {
	mix, found := l.mixes[key]
	if !found {
		return Mix{}, h.GetErr(key, errors.New("couldn't find mix"))
	}

	return mix, nil
}

func (l *Lookup) getModifier(key string) (Modifier, error) {
	modifier, found := l.modifiers[key]
	if !found {
		return Modifier{}, h.GetErr(key, errors.New("couldn't find modifier"))
	}

	return modifier, nil
}

func (l *Lookup) getMonster(key string) (Monster, error) {
	monster, found := l.monsters[key]
	if !found {
		return Monster{}, h.GetErr(key, errors.New("couldn't find monster"))
	}

	return monster, nil
}

func (l *Lookup) getOverdriveCommand(key string) (OverdriveCommand, error) {
	command, found := l.overdriveCommands[key]
	if !found {
		return OverdriveCommand{}, h.GetErr(key, errors.New("couldn't find overdrive command"))
	}

	return command, nil
}

func (l *Lookup) getOverdrive(ability Ability) (Overdrive, error) {
	key := createLookupKey(ability)

	overdrive, found := l.overdrives[key]
	if !found {
		return Overdrive{}, h.GetErr(key, errors.New("couldn't find overdrive"))
	}

	return overdrive, nil
}

func (l *Lookup) getOverdriveMode(key string) (OverdriveMode, error) {
	mode, found := l.overdriveModes[key]
	if !found {
		return OverdriveMode{}, h.GetErr(key, errors.New("couldn't find overdrive mode"))
	}

	return mode, nil
}

func (l *Lookup) getPosition(key string) (BlitzballPosition, error) {
	position, found := l.positions[key]
	if !found {
		return BlitzballPosition{}, h.GetErr(key, errors.New("couldn't find blitzball position"))
	}

	return position, nil
}

func (l *Lookup) getProperty(key string) (Property, error) {
	property, found := l.properties[key]
	if !found {
		return Property{}, h.GetErr(key, errors.New("couldn't find property"))
	}

	return property, nil
}

func (l *Lookup) getQuest(quest Quest) (Quest, error) {
	key := createLookupKey(quest)

	quest, found := l.quests[key]
	if !found {
		return Quest{}, h.GetErr(quest.Error(), errors.New("couldn't find quest"))
	}

	return quest, nil
}

func (l *Lookup) getSidequest(key string) (Sidequest, error) {
	quest := Quest{
		Name: key,
		Type: database.QuestTypeSidequest,
	}
	lookupKey := createLookupKey(quest)

	sidequest, found := l.sidequests[lookupKey]
	if !found {
		return Sidequest{}, h.GetErr(key, errors.New("couldn't find sidequest"))
	}

	return sidequest, nil
}

func (l *Lookup) getSubquest(key string) (Subquest, error) {
	quest := Quest{
		Name: key,
		Type: database.QuestTypeSubquest,
	}
	lookupKey := createLookupKey(quest)

	subquest, found := l.subquests[lookupKey]
	if !found {
		return Subquest{}, h.GetErr(key, errors.New("couldn't find subquest"))
	}

	return subquest, nil
}

func (l *Lookup) getShop(key string) (Shop, error) {
	shop, found := l.shops[key]
	if !found {
		return Shop{}, h.GetErr(key, errors.New("couldn't find shop"))
	}

	return shop, nil
}

func (l *Lookup) getSong(key string) (Song, error) {
	song, found := l.songs[key]
	if !found {
		return Song{}, h.GetErr(key, errors.New("couldn't find song"))
	}

	return song, nil
}

func (l *Lookup) getStat(key string) (Stat, error) {
	stat, found := l.stats[key]
	if !found {
		return Stat{}, h.GetErr(key, errors.New("couldn't find stat"))
	}

	return stat, nil
}

func (l *Lookup) getStatusCondition(key string) (StatusCondition, error) {
	condition, found := l.statusConditions[key]
	if !found {
		return StatusCondition{}, h.GetErr(key, errors.New("couldn't find status condition"))
	}

	return condition, nil
}

func (l *Lookup) getSubmenu(key string) (Submenu, error) {
	submenu, found := l.submenus[key]
	if !found {
		return Submenu{}, h.GetErr(key, errors.New("couldn't find submenu"))
	}

	return submenu, nil
}

func (l *Lookup) getTreasure(key string) (Treasure, error) {
	treasure, found := l.treasures[key]
	if !found {
		return Treasure{}, h.GetErr(key, errors.New("couldn't find treasure"))
	}

	return treasure, nil
}
