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
	items         map[string]ItemLookup
	keyItems      map[string]KeyItemLookup
	locationAreas map[string]LocationAreaLookup
	songs         map[string]SongLookup
	stats         map[string]StatLookup
}

func lookupInit() lookup {
	return lookup{
		items:         make(map[string]ItemLookup),
		keyItems:      make(map[string]KeyItemLookup),
		locationAreas: make(map[string]LocationAreaLookup),
		songs:         make(map[string]SongLookup),
		stats:         make(map[string]StatLookup),
	}

}

// whether to take a Lookupable or a string as an input depends on vibes. if it's just one field, I'll take the string. if it's a composite key, I'll make a struct.

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

func (l *lookup) getSong(songName string) (SongLookup, error) {
	songID, found := l.songs[songName]
	if !found {
		return SongLookup{}, fmt.Errorf("couldn't find Song %s", songName)
	}

	return songID, nil
}

func (l *lookup) getStat(statName string) (StatLookup, error) {
	stat, found := l.stats[statName]
	if !found {
		return StatLookup{}, fmt.Errorf("couldn't find Stat %s", statName)
	}

	return stat, nil
}
