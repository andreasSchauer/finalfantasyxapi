package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Lookupable interface {
	ToKeyFields() []any
}

func createLookupKey(l Lookupable) string {
	fields := l.ToKeyFields()
	return combineFields(fields)
}

type lookup struct {
	abilities			map[string]Ability
	enemyAbilities		map[string]EnemyAbility
	itemAbilities		map[string]ItemAbility
	overdriveAbilities	map[string]OverdriveAbility
	playerAbilities		map[string]PlayerAbility
	triggerCommands		map[string]TriggerCommand
	aeons				map[string]Aeon
	affinities			map[string]Affinity
	areas 				map[string]Area
	autoAbilities		map[string]AutoAbility
	celestialWeapons	map[string]CelestialWeapon
	characters			map[string]Character
	charClasses			map[string]CharacterClass
	elements			map[string]Element
	equipmentTables		map[string]EquipmentTable
	items         		map[string]Item
	keyItems      		map[string]KeyItem
	masterItems			map[string]MasterItem
	mixes				map[string]Mix
	modifiers			map[string]Modifier
	overdriveModes		map[string]OverdriveMode
	overdrives			map[string]Overdrive
	positions			map[string]BlitzballPosition
	properties			map[string]Property
	quests				map[string]Quest
	sidequests			map[string]Sidequest
	subquests			map[string]Subquest
	songs         		map[string]Song
	stats         		map[string]Stat
	statusConditions	map[string]StatusCondition
	submenus			map[string]Submenu
}


func lookupInit() lookup {
	return lookup{
		abilities:			make(map[string]Ability),
		enemyAbilities:		make(map[string]EnemyAbility),
		itemAbilities: 		make(map[string]ItemAbility),
		overdriveAbilities: make(map[string]OverdriveAbility),
		playerAbilities: 	make(map[string]PlayerAbility),
		triggerCommands: 	make(map[string]TriggerCommand),
		aeons: 				make(map[string]Aeon),
		affinities: 		make(map[string]Affinity),
		areas: 				make(map[string]Area),
		autoAbilities:		make(map[string]AutoAbility),
		celestialWeapons: 	make(map[string]CelestialWeapon),
		characters: 		make(map[string]Character),
		charClasses: 		make(map[string]CharacterClass),
		elements: 			make(map[string]Element),
		equipmentTables:	make(map[string]EquipmentTable),
		items:         		make(map[string]Item),
		keyItems:      		make(map[string]KeyItem),
		masterItems: 		make(map[string]MasterItem),
		mixes:				make(map[string]Mix),
		modifiers: 			make(map[string]Modifier),
		overdriveModes:		make(map[string]OverdriveMode),
		overdrives: 		make(map[string]Overdrive),
		positions: 			make(map[string]BlitzballPosition),
		properties: 		make(map[string]Property),
		quests: 			make(map[string]Quest),
		sidequests: 		make(map[string]Sidequest),
		subquests: 			make(map[string]Subquest),
		songs:         		make(map[string]Song),
		stats:         		make(map[string]Stat),
		statusConditions: 	make(map[string]StatusCondition),
		submenus: 			make(map[string]Submenu),
	}
}


// Lookupable for composite keys, string for simple name keys


func (l *lookup) getAbility(abilityReference AbilityReference) (Ability, error) {
	key := createLookupKey(abilityReference)

	ability, found := l.abilities[key]
	if !found {
		return Ability{}, fmt.Errorf("couldn't find Ability %s - %d - %s", ability.Name, derefOrNil(ability.Version), ability.Type)
	}

	return ability, nil
}


func (l *lookup) getEnemyAbility(abilityReference AbilityReference) (EnemyAbility, error) {
	key := createLookupKey(abilityReference)

	ability, found := l.enemyAbilities[key]
	if !found {
		return EnemyAbility{}, fmt.Errorf("couldn't find Enemy Ability %s - %d", ability.Name, derefOrNil(ability.Version))
	}

	return ability, nil
}


func (l *lookup) getItemAbility(abilityReference AbilityReference) (ItemAbility, error) {
	key := createLookupKey(abilityReference)

	ability, found := l.itemAbilities[key]
	if !found {
		return ItemAbility{}, fmt.Errorf("couldn't find Item Ability %s - %d", ability.Name, derefOrNil(ability.Version))
	}

	return ability, nil
}


func (l *lookup) getOverdriveAbility(abilityReference AbilityReference) (OverdriveAbility, error) {
	key := createLookupKey(abilityReference)

	ability, found := l.overdriveAbilities[key]
	if !found {
		return OverdriveAbility{}, fmt.Errorf("couldn't find Overdrive Ability %s - %d", ability.Name, derefOrNil(ability.Version))
	}

	return ability, nil
}


func (l *lookup) getPlayerAbility(abilityReference AbilityReference) (PlayerAbility, error) {
	key := createLookupKey(abilityReference)

	ability, found := l.playerAbilities[key]
	if !found {
		return PlayerAbility{}, fmt.Errorf("couldn't find Player Ability %s - %d", ability.Name, derefOrNil(ability.Version))
	}

	return ability, nil
}


func (l *lookup) getTriggerCommand(abilityReference AbilityReference) (TriggerCommand, error) {
	key := createLookupKey(abilityReference)

	command, found := l.triggerCommands[key]
	if !found {
		return TriggerCommand{}, fmt.Errorf("couldn't find Trigger Command %s - %d", command.Name, derefOrNil(command.Version))
	}

	return command, nil
}


func (l *lookup) getAeon(aeonName string) (Aeon, error) {
	playerUnit := PlayerUnit{
		Name: aeonName,
		Type: database.UnitTypeAeon,
	}
	key := createLookupKey(playerUnit)

	aeon, found := l.aeons[key]
	if !found {
		return Aeon{}, fmt.Errorf("couldn't find Aeon %s", aeonName)
	}

	return aeon, nil
}


func (l *lookup) getAffinity(affinityName string) (Affinity, error) {
	affinity, found := l.affinities[affinityName]
	if !found {
		return Affinity{}, fmt.Errorf("couldn't find Affinity %s", affinityName)
	}

	return affinity, nil
}


func (l *lookup) getArea(locationArea LocationArea) (Area, error) {
	key := createLookupKey(locationArea)

	area, found := l.areas[key]
	if !found {
		return Area{}, fmt.Errorf("couldn't find location area: %s - %s - %s - %d", locationArea.Location, locationArea.SubLocation, locationArea.Area, derefOrNil(locationArea.Version))
	}

	return area, nil
}


func (l *lookup) getAutoAbility(abilityName string) (AutoAbility, error) {
	autoAbility, found := l.autoAbilities[abilityName]
	if !found {
		return AutoAbility{}, fmt.Errorf("couldn't find Auto Ability %s", abilityName)
	}

	return autoAbility, nil
}


func (l *lookup) getCelestialWeapon(weaponName string) (CelestialWeapon, error) {
	weapon, found := l.celestialWeapons[weaponName]
	if !found {
		return CelestialWeapon{}, fmt.Errorf("couldn't find Celestial Weapon %s", weaponName)
	}

	return weapon, nil
}


func (l *lookup) getCharacter(charName string) (Character, error) {
	playerUnit := PlayerUnit{
		Name: charName,
		Type: database.UnitTypeCharacter,
	}
	key := createLookupKey(playerUnit)

	character, found := l.characters[key]
	if !found {
		return Character{}, fmt.Errorf("couldn't find Character %s", charName)
	}

	return character, nil
}


func (l *lookup) getCharacterClass(className string) (CharacterClass, error) {
	class, found := l.charClasses[className]
	if !found {
		return CharacterClass{}, fmt.Errorf("couldn't find Character Class %s", className)
	}

	return class, nil
}


func (l *lookup) getElement(elementName string) (Element, error) {
	element, found := l.elements[elementName]
	if !found {
		return Element{}, fmt.Errorf("couldn't find Element %s", elementName)
	}

	return element, nil
}


func (l *lookup) getEquipmentTable(tableKey string) (EquipmentTable, error) {
	equipmentTable, found := l.equipmentTables[tableKey]
	if !found {
		return EquipmentTable{}, fmt.Errorf("couldn't find Equipment Table %s", tableKey)
	}

	return equipmentTable, nil
}


func (l *lookup) getItem(itemName string) (Item, error) {
	masterItem := MasterItem{
		Name: itemName,
		Type: database.ItemTypeItem,
	}
	key := createLookupKey(masterItem)

	item, found := l.items[key]
	if !found {
		return Item{}, fmt.Errorf("couldn't find Item %s", itemName)
	}

	return item, nil
}


func (l *lookup) getKeyItem(itemName string) (KeyItem, error) {
	masterItem := MasterItem{
		Name: itemName,
		Type: database.ItemTypeKeyItem,
	}
	key := createLookupKey(masterItem)

	keyItem, found := l.keyItems[key]
	if !found {
		return KeyItem{}, fmt.Errorf("couldn't find Key Item %s", itemName)
	}

	return keyItem, nil
}


func (l *lookup) getMasterItem(itemName string) (MasterItem, error) {
	masterItem, found := l.masterItems[itemName]
	if !found {
		return MasterItem{}, fmt.Errorf("couldn't find Master Item %s", itemName)
	}

	return masterItem, nil
}


func (l *lookup) getMix(mixName string) (Mix, error) {
	mix, found := l.mixes[mixName]
	if !found {
		return Mix{}, fmt.Errorf("couldn't find Mix %s", mixName)
	}

	return mix, nil
}


func (l *lookup) getModifier(modifierName string) (Modifier, error) {
	modifier, found := l.modifiers[modifierName]
	if !found {
		return Modifier{}, fmt.Errorf("couldn't find Modifier %s", modifierName)
	}

	return modifier, nil
}


func (l *lookup) getOverdrive(ability Ability) (Overdrive, error) {
	key := createLookupKey(ability)

	overdrive, found := l.overdrives[key]
	if !found {
		return Overdrive{}, fmt.Errorf("couldn't find overdrive: %s - %d", ability.Name, derefOrNil(ability.Version))
	}

	return overdrive, nil
}


func (l *lookup) getOverdriveMode(modeName string) (OverdriveMode, error) {
	mode, found := l.overdriveModes[modeName]
	if !found {
		return OverdriveMode{}, fmt.Errorf("couldn't find Overdrive Mode %s", modeName)
	}

	return mode, nil
}


func (l *lookup) getPosition(positionKey string) (BlitzballPosition, error) {
	position, found := l.positions[positionKey]
	if !found {
		return BlitzballPosition{}, fmt.Errorf("couldn't find Blitzball Position %s", positionKey)
	}

	return position, nil
}


func (l *lookup) getProperty(propertyName string) (Property, error) {
	property, found := l.properties[propertyName]
	if !found {
		return Property{}, fmt.Errorf("couldn't find Property %s", propertyName)
	}

	return property, nil
}


func (l *lookup) getSidequest(questName string) (Sidequest, error) {
	quest := Quest{
		Name: questName,
		Type: database.QuestTypeSidequest,
	}
	key := createLookupKey(quest)

	sidequest, found := l.sidequests[key]
	if !found {
		return Sidequest{}, fmt.Errorf("couldn't find Sidequest %s", questName)
	}

	return sidequest, nil
}


func (l *lookup) getQuest(quest Quest) (Quest, error) {
	key := createLookupKey(quest)

	quest, found := l.quests[key]
	if !found {
		return Quest{}, fmt.Errorf("couldn't find Quest %s", quest.Name)
	}

	return quest, nil
}


func (l *lookup) getSubquest(questName string) (Subquest, error) {
	quest := Quest{
		Name: questName,
		Type: database.QuestTypeSubquest,
	}
	key := createLookupKey(quest)

	subquest, found := l.subquests[key]
	if !found {
		return Subquest{}, fmt.Errorf("couldn't find Subquest %s", questName)
	}

	return subquest, nil
}


func (l *lookup) getSong(songName string) (Song, error) {
	song, found := l.songs[songName]
	if !found {
		return Song{}, fmt.Errorf("couldn't find Song %s", songName)
	}

	return song, nil
}


func (l *lookup) getStat(statName string) (Stat, error) {
	stat, found := l.stats[statName]
	if !found {
		return Stat{}, fmt.Errorf("couldn't find Stat %s", statName)
	}

	return stat, nil
}


func (l *lookup) getStatusCondition(conditionName string) (StatusCondition, error) {
	condition, found := l.statusConditions[conditionName]
	if !found {
		return StatusCondition{}, fmt.Errorf("couldn't find Status Condition %s", conditionName)
	}

	return condition, nil
}


func (l *lookup) getSubmenu(menuName string) (Submenu, error) {
	submenu, found := l.submenus[menuName]
	if !found {
		return Submenu{}, fmt.Errorf("couldn't find Submenu %s", menuName)
	}

	return submenu, nil
}