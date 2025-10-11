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
	aeons				map[string]AeonLookup
	characters			map[string]CharacterLookup
	charClasses			map[string]CharClassLookup
	elements			map[string]ElementLookup
	items         		map[string]ItemLookup
	keyItems      		map[string]KeyItemLookup
	locationAreas 		map[string]LocationAreaLookup
	modifiers			map[string]ModifierLookup
	overdriveModes		map[string]OverdriveModeLookup
	overdrives			map[string]OverdriveLookup
	properties			map[string]PropertyLookup
	songs         		map[string]SongLookup
	stats         		map[string]StatLookup
	statusConditions	map[string]StatusConditionLookup
}

func lookupInit() lookup {
	return lookup{
		aeons: 				make(map[string]AeonLookup),
		characters: 		make(map[string]CharacterLookup),
		charClasses: 		make(map[string]CharClassLookup),
		elements: 			make(map[string]ElementLookup),
		items:         		make(map[string]ItemLookup),
		keyItems:      		make(map[string]KeyItemLookup),
		locationAreas: 		make(map[string]LocationAreaLookup),
		modifiers: 			make(map[string]ModifierLookup),
		overdriveModes:		make(map[string]OverdriveModeLookup),
		overdrives: 		make(map[string]OverdriveLookup),
		properties: 		make(map[string]PropertyLookup),
		songs:         		make(map[string]SongLookup),
		stats:         		make(map[string]StatLookup),
		statusConditions: 	make(map[string]StatusConditionLookup),
	}

}


// Lookupable for composite keys, string for simple name keys

func (l *lookup) getAeon(aeonName string) (AeonLookup, error) {
	playerUnit := PlayerUnit{
		Name: aeonName,
		Type: database.UnitTypeAeon,
	}
	key := createLookupKey(playerUnit)

	aeon, found := l.aeons[key]
	if !found {
		return AeonLookup{}, fmt.Errorf("couldn't find Aeon %s", aeonName)
	}

	return aeon, nil
}


func (l *lookup) getCharacter(charName string) (CharacterLookup, error) {
	playerUnit := PlayerUnit{
		Name: charName,
		Type: database.UnitTypeCharacter,
	}
	key := createLookupKey(playerUnit)

	character, found := l.characters[key]
	if !found {
		return CharacterLookup{}, fmt.Errorf("couldn't find Character %s", charName)
	}

	return character, nil
}


func (l *lookup) getCharacterClass(className string) (CharClassLookup, error) {
	class, found := l.charClasses[className]
	if !found {
		return CharClassLookup{}, fmt.Errorf("couldn't find Character Class %s", className)
	}

	return class, nil
}


func (l *lookup) getElement(elementName string) (ElementLookup, error) {
	element, found := l.elements[elementName]
	if !found {
		return ElementLookup{}, fmt.Errorf("couldn't find Element %s", elementName)
	}

	return element, nil
}


func (l *lookup) getArea(locationArea LocationArea) (LocationAreaLookup, error) {
	key := createLookupKey(locationArea)

	area, found := l.locationAreas[key]
	if !found {
		return LocationAreaLookup{}, fmt.Errorf("couldn't find location area: %s - %s - %s - %d", locationArea.Location, locationArea.SubLocation, locationArea.Area, derefOrNil(locationArea.Version))
	}

	return area, nil
}


func (l *lookup) getItem(itemName string) (ItemLookup, error) {
	masterItem := MasterItem{
		Name: itemName,
		Type: database.ItemTypeItem,
	}
	key := createLookupKey(masterItem)

	item, found := l.items[key]
	if !found {
		return ItemLookup{}, fmt.Errorf("couldn't find Item %s", itemName)
	}

	return item, nil
}


func (l *lookup) getKeyItem(itemName string) (KeyItemLookup, error) {
	masterItem := MasterItem{
		Name: itemName,
		Type: database.ItemTypeKeyItem,
	}
	key := createLookupKey(masterItem)

	keyItem, found := l.keyItems[key]
	if !found {
		return KeyItemLookup{}, fmt.Errorf("couldn't find Key Item %s", itemName)
	}

	return keyItem, nil
}


func (l *lookup) getModifier(modifierName string) (ModifierLookup, error) {
	modifier, found := l.modifiers[modifierName]
	if !found {
		return ModifierLookup{}, fmt.Errorf("couldn't find Modifier %s", modifierName)
	}

	return modifier, nil
}


func (l *lookup) getOverdrive(ability Ability) (OverdriveLookup, error) {
	key := createLookupKey(ability)

	overdrive, found := l.overdrives[key]
	if !found {
		return OverdriveLookup{}, fmt.Errorf("couldn't find overdrive: %s - %d", ability.Name, derefOrNil(ability.Version))
	}

	return overdrive, nil
}


func (l *lookup) getOverdriveMode(modeName string) (OverdriveModeLookup, error) {
	mode, found := l.overdriveModes[modeName]
	if !found {
		return OverdriveModeLookup{}, fmt.Errorf("couldn't find Overdrive Mode %s", modeName)
	}

	return mode, nil
}


func (l *lookup) getProperty(propertyName string) (PropertyLookup, error) {
	property, found := l.properties[propertyName]
	if !found {
		return PropertyLookup{}, fmt.Errorf("couldn't find Property %s", propertyName)
	}

	return property, nil
}


func (l *lookup) getSong(songName string) (SongLookup, error) {
	song, found := l.songs[songName]
	if !found {
		return SongLookup{}, fmt.Errorf("couldn't find Song %s", songName)
	}

	return song, nil
}


func (l *lookup) getStat(statName string) (StatLookup, error) {
	stat, found := l.stats[statName]
	if !found {
		return StatLookup{}, fmt.Errorf("couldn't find Stat %s", statName)
	}

	return stat, nil
}


func (l *lookup) getStatusCondition(conditionName string) (StatusConditionLookup, error) {
	condition, found := l.statusConditions[conditionName]
	if !found {
		return StatusConditionLookup{}, fmt.Errorf("couldn't find Status Condition %s", conditionName)
	}

	return condition, nil
}
