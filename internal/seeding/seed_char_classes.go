package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type CharacterClass struct {
	Name string
}

func (cc CharacterClass) ToHashFields() []any {
	return []any{
		cc.Name,
	}
}

type CharClassLookup struct {
	CharacterClass
	ID int32
}

func (l *lookup) seedCharacterClasses(qtx *database.Queries, unit PlayerUnit) error {
	if unit.Type == database.UnitTypeCharacter {
		err := l.seedCharClassesCharacter(qtx, unit)
		if err != nil {
			return err
		}
	}

	if unit.Type == database.UnitTypeAeon {
		err := l.seedCharClassesAeon(qtx, unit)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *lookup) seedCharClassesCharacter(qtx *database.Queries, unit PlayerUnit) error {
	character, err := l.getCharacter(unit.Name)
	if err != nil {
		return err
	}

	err = l.seedCharacterClass(qtx, character.Name, character.UnitID)
	if err != nil {
		return fmt.Errorf("%s: %v", character.Name, err)
	}

	if !character.StoryOnly {
		err := l.seedCharacterClass(qtx, "characters", character.UnitID)
		if err != nil {
			return fmt.Errorf("%s: %v", character.Name, err)
		}
	}

	return nil
}

func (l *lookup) seedCharClassesAeon(qtx *database.Queries, unit PlayerUnit) error {
	aeon, err := l.getAeon(unit.Name)
	if err != nil {
		return err
	}

	aeonCategory := stringPtrToString(aeon.Category)

	err = l.seedCharacterClass(qtx, aeon.Name, aeon.UnitID)
	if err != nil {
		return fmt.Errorf("%s: %v", aeon.Name, err)
	}

	err = l.seedCharacterClass(qtx, "aeons", aeon.UnitID)
	if err != nil {
		return fmt.Errorf("%s: %v", aeon.Name, err)
	}

	if aeonCategory == "standard-aeons" {
		err := l.seedCharacterClass(qtx, "standard-aeons", aeon.UnitID)
		if err != nil {
			return fmt.Errorf("%s: %v", aeon.Name, err)
		}
	}

	if aeonCategory == "magus-sisters" {
		err := l.seedCharacterClass(qtx, "magus-sisters", aeon.UnitID)
		if err != nil {
			return fmt.Errorf("%s: %v", aeon.Name, err)
		}
	}

	return nil
}

func (l *lookup) seedCharacterClass(qtx *database.Queries, className string, unitID int32) error {
	class := CharacterClass{
		Name: className,
	}

	dbClass, err := qtx.CreateCharacterClass(context.Background(), database.CreateCharacterClassParams{
		DataHash: generateDataHash(class),
		Name:     class.Name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create Character Class: %s: %v", class.Name, err)
	}

	l.charClasses[className] = CharClassLookup{
		CharacterClass: class,
		ID:             dbClass.ID,
	}

	err = l.seedUnitCharClassJunction(qtx, unitID, dbClass.ID)
	if err != nil {
		return fmt.Errorf("couldn't create junction with Character Class: %s: %v", className, err)
	}

	return nil
}


func (l *lookup) seedUnitCharClassJunction(qtx *database.Queries, unitID int32, classID int32) error {
	junction := Junction{
		ParentID: 	unitID,
		ChildID:  	classID,
	}

	err := qtx.CreateUnitsCharClassesJunction(context.Background(), database.CreateUnitsCharClassesJunctionParams{
		DataHash: generateDataHash(junction),
		UnitID:   junction.ParentID,
		ClassID:  junction.ChildID,
	})
	if err != nil {
		return err
	}

	return nil
}
