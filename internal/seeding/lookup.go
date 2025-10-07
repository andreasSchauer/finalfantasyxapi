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
	locationAreaToID map[string]int32
	itemNameToID map[string]int32
	keyItemNameToID map[string]int32
	songNameToID map[string]int32
}


func lookupInit() lookup {
	return lookup{
		locationAreaToID: make(map[string]int32),
		itemNameToID: make(map[string]int32),
		keyItemNameToID: make(map[string]int32),
		songNameToID: make(map[string]int32),
	}

}


func (l *lookup) getAreaID(locationArea LocationArea) (int32, error) {
	key := createLookupKey(locationArea)
	locationAreaID, found := l.locationAreaToID[key]
	if !found {
		return 0, fmt.Errorf("couldn't find location area: %s - %s - %s - %d", locationArea.Location, locationArea.SubLocation, locationArea.Area, derefOrNil(locationArea.Version))
	}

	return locationAreaID, nil
}


func (l *lookup) getItemID(itemName string) (int32, error) {
	masterItem := MasterItem{
		Name: itemName,
		Type: database.ItemTypeItem,
	}
	key := createLookupKey(masterItem)

	itemID, found := l.itemNameToID[key]
	if !found {
		return 0, fmt.Errorf("couldn't find Item %s", itemName)
	}

	return itemID, nil
}


func (l *lookup) getKeyItemID(itemName string) (int32, error) {
	masterItem := MasterItem{
		Name: itemName,
		Type: database.ItemTypeKeyItem,
	}
	key := createLookupKey(masterItem)

	itemID, found := l.keyItemNameToID[key]
	if !found {
		return 0, fmt.Errorf("couldn't find Key Item %s", itemName)
	}

	return itemID, nil
}


func (l *lookup) getSongID(songName string) (int32, error) {
	song := Song{
		Name: songName,
	}
	key := createLookupKey(song)

	songID, found := l.songNameToID[key]
	if !found {
		return 0, fmt.Errorf("couldn't find Song %s", songName)
	}

	return songID, nil
}