package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
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

func (cc CharacterClass) Error() string {
	return fmt.Sprintf("character class %s", cc.Name)
}

func (l *Lookup) seedCharacterClasses(qtx *database.Queries, unit PlayerUnit) error {
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

func (l *Lookup) seedCharClassesCharacter(qtx *database.Queries, unit PlayerUnit) error {
	character, err := GetResource(unit.Name, l.Characters)
	if err != nil {
		return err
	}

	err = l.seedUnitCharClass(qtx, character.Name, character.PlayerUnit)
	if err != nil {
		return h.NewErr(character.Name, err, "character class")
	}

	if !character.StoryOnly {
		err := l.seedUnitCharClass(qtx, "characters", character.PlayerUnit)
		if err != nil {
			return h.NewErr("characters", err, "character class")
		}
	}

	return nil
}

func (l *Lookup) seedCharClassesAeon(qtx *database.Queries, unit PlayerUnit) error {
	aeon, err := GetResource(unit.Name, l.Aeons)
	if err != nil {
		return err
	}

	aeonCategory := h.StringPtrToString(aeon.Category)

	err = l.seedUnitCharClass(qtx, aeon.Name, aeon.PlayerUnit)
	if err != nil {
		return h.NewErr(aeon.Name, err, "character class")
	}

	err = l.seedUnitCharClass(qtx, "aeons", aeon.PlayerUnit)
	if err != nil {
		return h.NewErr("aeons", err, "character class")
	}

	if aeonCategory == "standard-aeons" {
		err = l.seedUnitCharClass(qtx, "standard-aeons", aeon.PlayerUnit)
		if err != nil {
			return h.NewErr("standard-aeons", err, "character class")
		}
	}

	if aeonCategory == "magus-sisters" {
		err = l.seedUnitCharClass(qtx, "magus-sisters", aeon.PlayerUnit)
		if err != nil {
			return h.NewErr("magus-sisters", err, "character class")
		}
	}

	return nil
}

func (l *Lookup) seedUnitCharClass(qtx *database.Queries, className string, unit PlayerUnit) error {
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
		return h.NewErr(unit.Error(), err, "couldn't junction player unit")
	}

	return nil
}

func (l *Lookup) seedCharacterClass(qtx *database.Queries, class CharacterClass) (CharacterClass, error) {
	dbClass, err := qtx.CreateCharacterClass(context.Background(), database.CreateCharacterClassParams{
		DataHash: generateDataHash(class),
		Name:     class.Name,
	})
	if err != nil {
		return CharacterClass{}, h.NewErr(class.Error(), err, "couldn't create character class")
	}

	class.ID = dbClass.ID
	l.CharClasses[class.Name] = class

	return class, nil
}
