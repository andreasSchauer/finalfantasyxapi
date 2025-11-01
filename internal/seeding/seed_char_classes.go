package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type CharacterClass struct {
	ID   int32
	Name string
}

func (cc CharacterClass) ToHashFields() []any {
	return []any{
		cc.Name,
	}
}

func (cc CharacterClass) GetID() int32 {
	return cc.ID
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

	err = l.seedUnitCharClass(qtx, character.Name, character.PlayerUnit)
	if err != nil {
		return fmt.Errorf("%s: %v", character.Name, err)
	}

	if !character.StoryOnly {
		err := l.seedUnitCharClass(qtx, "characters", character.PlayerUnit)
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

	err = l.seedUnitCharClass(qtx, aeon.Name, aeon.PlayerUnit)
	if err != nil {
		return fmt.Errorf("%s: %v", aeon.Name, err)
	}

	err = l.seedUnitCharClass(qtx, "aeons", aeon.PlayerUnit)
	if err != nil {
		return fmt.Errorf("%s: %v", aeon.Name, err)
	}

	if aeonCategory == "standard-aeons" {
		err = l.seedUnitCharClass(qtx, "standard-aeons", aeon.PlayerUnit)
		if err != nil {
			return fmt.Errorf("%s: %v", aeon.Name, err)
		}
	}

	if aeonCategory == "magus-sisters" {
		err = l.seedUnitCharClass(qtx, "magus-sisters", aeon.PlayerUnit)
		if err != nil {
			return fmt.Errorf("%s: %v", aeon.Name, err)
		}
	}

	return nil
}

func (l *lookup) seedUnitCharClass(qtx *database.Queries, className string, unit PlayerUnit) error {
	class := CharacterClass{
		Name: className,
	}

	junction, err := createJunctionSeed(qtx, unit, class, l.seedCharacterClass)
	if err != nil {
		return err
	}

	err = qtx.CreatePlayerUnitsCharacterClassJunction(context.Background(), database.CreatePlayerUnitsCharacterClassJunctionParams{
		DataHash: generateDataHash(junction),
		UnitID:   junction.ParentID,
		ClassID:  junction.ChildID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (l *lookup) seedCharacterClass(qtx *database.Queries, class CharacterClass) (CharacterClass, error) {
	dbClass, err := qtx.CreateCharacterClass(context.Background(), database.CreateCharacterClassParams{
		DataHash: generateDataHash(class),
		Name:     class.Name,
	})
	if err != nil {
		return CharacterClass{}, fmt.Errorf("couldn't create Character Class: %s: %v", class.Name, err)
	}

	class.ID = dbClass.ID
	l.charClasses[class.Name] = class

	return class, nil
}
