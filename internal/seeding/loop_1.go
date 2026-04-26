package seeding

import (
	"context"
	"database/sql"
	"fmt"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) seedLoop1(qtx *database.Queries, ctx context.Context) error {
	return l.seedLoop(qtx, ctx, []func(*database.Queries, context.Context) error{
		l.loop1SeedAgilityTiers,
		l.loop1SeedElements,
		l.loop1SeedOverdriveModes,
		l.loop1SeedProperties,
		l.loop1SeedModifiers,
		l.loop1SeedPlayerUnits,
		l.loop1SeedCharacterClasses,
		l.loop1SeedBlitzballPositions,
		l.loop1SeedTopmenus,
		l.loop1SeedAbilityAttributes,
		l.loop1SeedMasterItems,
		l.loop1SeedCreatedNodes,
		l.loop1SeedLocations,
		l.loop1SeedBackgroundMusic,
		l.loop1SeedSongCredits,
		l.loop1SeedAbilityAccuracies,
		l.loop1SeedInflictedDelays,
		l.loop1SeedMonsters,
		l.loop1SeedMonsterSelections,
		l.loop1SeedEquipmentSlotsChances,
	})
}

func (l *Lookup) seedLoop2(qtx *database.Queries, ctx context.Context) error {
	return l.seedLoop(qtx, ctx, []func(*database.Queries, context.Context) error{
		l.loop2SeedAgilitySubtiers,
		l.loop2UpdateElements,
		l.loop2SeedElementalResists,
		l.loop2SeedSubmenus,
		l.loop2SeedAbilities,
		l.loop2SeedItems,
		l.loop2SeedKeyItems,
		l.loop2SeedItemAmounts,
	})
}

func (l *Lookup) loop2SeedItemAmounts(qtx *database.Queries, ctx context.Context) error {
	itemAmtsJson, err := l.getItemAmounts()
	if err != nil {
		return err
	}
	itemAmts := dedupeRows(itemAmtsJson, l.Hashes)

	params := database.CreateItemAmountBulkParams{
		DataHash:     make([]string, len(itemAmts)),
		MasterItemID: make([]int32, len(itemAmts)),
		Amount:       make([]int32, len(itemAmts)),
	}

	for i, ia := range itemAmts {
		params.DataHash[i] = generateDataHash(ia)
		params.MasterItemID[i] = ia.MasterItem.ID
		params.Amount[i] = ia.Amount
	}

	dbRows, err := qtx.CreateItemAmountBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create item amounts: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getItemAmounts() ([]ItemAmount, error) {
	itemAmounts := []ItemAmount{}
	var err error

	for i := range l.json.autoAbilities {
		autoAbility := &l.json.autoAbilities[i]

		if autoAbility.RequiredItem != nil {
			itemAmt, err := l.prepareItemAmt(autoAbility.RequiredItem)
			if err != nil {
				return nil, err
			}

			itemAmounts = append(itemAmounts, *itemAmt)
		}
	}

	for i := range l.json.blitzballPositions {
		position := &l.json.blitzballPositions[i]

		for j := range position.Items {
			itemAmt, err := l.prepareItemAmt(&position.Items[j].ItemAmount)
			if err != nil {
				return nil, err
			}

			itemAmounts = append(itemAmounts, *itemAmt)
		}
	}

	for i := range l.json.monsters {
		monster := l.json.monsters[i]

		if monster.Items != nil {
			items := monster.Items
			sc := items.StealCommon
			sr := items.StealRare
			dc := items.DropCommon
			dr := items.DropRare
			sdc := items.SecondaryDropCommon
			sdr := items.SecondaryDropRare
			br := items.Bribe

			if sc != nil {
				sc, err = l.prepareItemAmt(sc)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *sc)
			}

			if sr != nil {
				sr, err = l.prepareItemAmt(sr)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *sr)
			}

			if dc != nil {
				dc, err = l.prepareItemAmt(dc)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *dc)
			}

			if dr != nil {
				dr, err = l.prepareItemAmt(dr)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *dr)
			}

			if sdc != nil {
				sdc, err = l.prepareItemAmt(sdc)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *sdc)
			}

			if sdr != nil {
				sdr, err = l.prepareItemAmt(sdr)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *sdr)
			}

			if br != nil {
				br, err = l.prepareItemAmt(br)
				if err != nil {
					return nil, err
				}
				itemAmounts = append(itemAmounts, *br)
			}

			for j := range items.OtherItems {
				itemAmt, err := l.prepareItemAmt(&items.OtherItems[j].ItemAmount)
				if err != nil {
					return nil, err
				}

				itemAmounts = append(itemAmounts, *itemAmt)
			}
		}
	}

	for i := range l.json.playerAbilities {
		playerAbility := &l.json.playerAbilities[i]

		if playerAbility.AeonLearnItem != nil {
			itemAmt, err := l.prepareItemAmt(playerAbility.AeonLearnItem)
			if err != nil {
				return nil, err
			}

			itemAmounts = append(itemAmounts, *itemAmt)
		}
	}

	for i := range l.json.sidequests {
		sidequest := &l.json.sidequests[i]

		if sidequest.Completion != nil {
			itemAmt, err := l.prepareItemAmt(&sidequest.Completion.Reward)
			if err != nil {
				return nil, err
			}

			itemAmounts = append(itemAmounts, *itemAmt)
		}

		for j := range sidequest.Subquests {
			subquest := sidequest.Subquests[j]

			if subquest.Completion != nil {
				itemAmt, err := l.prepareItemAmt(&subquest.Completion.Reward)
				if err != nil {
					return nil, err
				}

				itemAmounts = append(itemAmounts, *itemAmt)
			}
		}
	}

	for i := range l.json.treasureLists {
		treasureList := &l.json.treasureLists[i]
		for j := range treasureList.Treasures {
			treasure := &treasureList.Treasures[j]
			
			for j := range treasure.Items {
				itemAmt, err := l.prepareItemAmt(&treasure.Items[j])
				if err != nil {
					return nil, err
				}
				
				itemAmounts = append(itemAmounts, *itemAmt)
			}
		}
	}

	return itemAmounts, nil
}

func (l *Lookup) prepareItemAmt(ia *ItemAmount) (*ItemAmount, error) {
	var err error
	ia.MasterItem.Name = ia.ItemName

	ia.MasterItem.ID, err = assignFK(ia.MasterItem.Name, l.MasterItems)
	if err != nil {
		return nil, err
	}

	return ia, nil
}

func (l *Lookup) loop2SeedKeyItems(qtx *database.Queries, ctx context.Context) error {
	keyItemsJson, err := l.getKeyItems()
	keyItems := dedupeRows(keyItemsJson, l.Hashes)

	params := database.CreateKeyItemBulkParams{
		DataHash:     make([]string, len(keyItems)),
		MasterItemID: make([]int32, len(keyItems)),
		Category:     make([]database.KeyItemCategory, len(keyItems)),
		Description:  make([]string, len(keyItems)),
		Effect:       make([]string, len(keyItems)),
	}

	for i, ki := range keyItems {
		params.DataHash[i] = generateDataHash(ki)
		params.MasterItemID[i] = ki.MasterItem.ID
		params.Category[i] = database.KeyItemCategory(ki.Category)
		params.Description[i] = ki.Description
		params.Effect[i] = ki.Effect
	}

	dbRows, err := qtx.CreateKeyItemBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create key items: %v", err)
	}

	for i, row := range dbRows {
		keyItems[i].ID = row.ID
		l.json.keyItems[i].ID = row.ID
		l.KeyItems[keyItems[i].Name] = keyItems[i]
		l.KeyItemsID[row.ID] = keyItems[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getKeyItems() ([]KeyItem, error) {
	keyItems := []KeyItem{}
	var err error

	for i := range l.json.keyItems {
		keyItem := &l.json.keyItems[i]
		keyItem.MasterItem.ID, err = assignFK(keyItem.Name, l.MasterItems)
		if err != nil {
			return nil, err
		}
		keyItems = append(keyItems, *keyItem)
	}

	return keyItems, nil
}

func (l *Lookup) loop2SeedItems(qtx *database.Queries, ctx context.Context) error {
	itemsJson, err := l.getItems()
	items := dedupeRows(itemsJson, l.Hashes)

	params := database.CreateItemBulkParams{
		DataHash:     make([]string, len(items)),
		MasterItemID: make([]int32, len(items)),
		Description:  make([]string, len(items)),
		Effect:       make([]string, len(items)),
		Category:     make([]database.ItemCategory, len(items)),
		Usability:    make([]database.ItemUsability, len(items)),
		BasePrice:    make([]sql.NullInt32, len(items)),
		SellValue:    make([]int32, len(items)),
	}

	for i, item := range items {
		params.DataHash[i] = generateDataHash(item)
		params.MasterItemID[i] = item.MasterItem.ID
		params.Description[i] = item.Description
		params.Effect[i] = item.Effect
		params.Category[i] = database.ItemCategory(item.Category)
		params.Usability[i] = database.ItemUsability(item.Usability)
		params.BasePrice[i] = h.GetNullInt32(item.BasePrice)
		params.SellValue[i] = item.SellValue
	}

	dbRows, err := qtx.CreateItemBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create items: %v", err)
	}

	for i, row := range dbRows {
		items[i].ID = row.ID
		l.json.items[i].ID = row.ID
		l.Items[items[i].Name] = items[i]
		l.ItemsID[row.ID] = items[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getItems() ([]Item, error) {
	items := []Item{}
	var err error

	for i := range l.json.items {
		item := &l.json.items[i]
		item.MasterItem.ID, err = assignFK(item.Name, l.MasterItems)
		if err != nil {
			return nil, err
		}
		items = append(items, *item)
	}

	return items, nil
}

func (l *Lookup) loop2SeedAbilities(qtx *database.Queries, ctx context.Context) error {
	abilitiesJson := l.getAbilities()
	abilities := dedupeRows(abilitiesJson, l.Hashes)

	params := database.CreateAbilityBulkParams{
		DataHash:      make([]string, len(abilities)),
		Name:          make([]string, len(abilities)),
		Version:       make([]sql.NullInt32, len(abilities)),
		Specification: make([]sql.NullString, len(abilities)),
		AttributesID:  make([]int32, len(abilities)),
		Type:          make([]database.AbilityType, len(abilities)),
	}

	for i, a := range abilities {
		params.DataHash[i] = generateDataHash(a)
		params.Name[i] = a.Name
		params.Version[i] = h.GetNullInt32(a.Version)
		params.Specification[i] = h.GetNullString(a.Specification)
		params.AttributesID[i] = a.Attributes.ID
		params.Type[i] = a.Type
	}

	dbRows, err := qtx.CreateAbilityBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create abilities: %v", err)
	}

	for i, row := range dbRows {
		abilities[i].ID = row.ID
		key := CreateLookupKey(abilities[i])
		l.Abilities[key] = abilities[i]
		l.AbilitiesID[row.ID] = abilities[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getAbilities() []Ability {
	abilities := []Ability{}

	for i := range l.json.playerAbilities {
		ability := &l.json.playerAbilities[i]
		ability.Attributes.ID = l.Hashes[generateDataHash(ability.Attributes)]
		ability.Type = database.AbilityTypePlayerAbility
		abilities = append(abilities, ability.Ability)
	}

	for i := range l.json.overdriveAbilities {
		ability := &l.json.overdriveAbilities[i]
		ability.Attributes.ID = l.Hashes[generateDataHash(ability.Attributes)]
		ability.Type = database.AbilityTypeOverdriveAbility
		abilities = append(abilities, ability.Ability)
	}

	for i := range l.json.items {
		item := &l.json.items[i]
		if len(item.BattleInteractions) > 0 {
			item.Attributes.ID = l.Hashes[generateDataHash(item.Attributes)]
			item.Ability.Name = item.Name
			item.Ability.Type = database.AbilityTypeItemAbility
			abilities = append(abilities, item.Ability)
		}
	}

	for i := range l.json.triggerCommands {
		command := &l.json.triggerCommands[i]
		command.Attributes.ID = l.Hashes[generateDataHash(command.Attributes)]
		command.Type = database.AbilityTypeTriggerCommand
		abilities = append(abilities, command.Ability)
	}

	for i := range l.json.unspecifiedAbilities {
		ability := &l.json.unspecifiedAbilities[i]
		ability.Attributes.ID = l.Hashes[generateDataHash(ability.Attributes)]
		ability.Type = database.AbilityTypeUnspecifiedAbility
		abilities = append(abilities, ability.Ability)
	}

	for i := range l.json.enemyAbilities {
		ability := &l.json.enemyAbilities[i]
		ability.Attributes.ID = l.Hashes[generateDataHash(ability.Attributes)]
		ability.Type = database.AbilityTypeEnemyAbility
		abilities = append(abilities, ability.Ability)
	}

	return abilities
}

func (l *Lookup) loop2SeedSubmenus(qtx *database.Queries, ctx context.Context) error {
	submenusJson, err := l.getSubmenus()
	if err != nil {
		return err
	}
	submenus := dedupeRows(submenusJson, l.Hashes)

	params := database.CreateSubmenuBulkParams{
		DataHash:    make([]string, len(submenus)),
		Name:        make([]string, len(submenus)),
		TopmenuID:   make([]sql.NullInt32, len(submenus)),
		Description: make([]sql.NullString, len(submenus)),
		Effect:      make([]string, len(submenus)),
	}

	for i, s := range submenus {
		params.DataHash[i] = generateDataHash(s)
		params.Name[i] = s.Name
		params.TopmenuID[i] = h.GetNullInt32(s.TopmenuID)
		params.Description[i] = h.GetNullString(s.Description)
		params.Effect[i] = s.Effect
	}

	dbRows, err := qtx.CreateSubmenuBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create submenus: %v", err)
	}

	for i, row := range dbRows {
		submenus[i].ID = row.ID
		l.json.submenus[i].ID = row.ID
		l.Submenus[submenus[i].Name] = submenus[i]
		l.SubmenusID[row.ID] = submenus[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getSubmenus() ([]Submenu, error) {
	submenus := []Submenu{}
	var err error

	for i := range l.json.submenus {
		submenu := &l.json.submenus[i]

		if submenu.Topmenu != nil {
			submenu.TopmenuID, err = assignFKPtr(submenu.Topmenu, l.Topmenus)
			if err != nil {
				return nil, err
			}
		}
		submenus = append(submenus, *submenu)
	}

	return submenus, nil
}

func (l *Lookup) loop2SeedElementalResists(qtx *database.Queries, ctx context.Context) error {
	resistsJson, err := l.getElementalResists()
	if err != nil {
		return err
	}
	resists := dedupeRows(resistsJson, l.Hashes)

	params := database.CreateElementalResistBulkParams{
		DataHash:  make([]string, len(resists)),
		ElementID: make([]int32, len(resists)),
		Affinity:  make([]database.ElementalAffinity, len(resists)),
	}

	for i, er := range resists {
		params.DataHash[i] = generateDataHash(er)
		params.ElementID[i] = er.ElementID
		params.Affinity[i] = database.ElementalAffinity(er.Affinity)
	}

	dbRows, err := qtx.CreateElementalResistBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create elemental resists: %v", err)
	}

	for i, row := range dbRows {
		resists[i].ID = row.ID
		key := CreateLookupKey(resists[i])
		l.ElementalResists[key] = resists[i]
		l.ElementalResistsID[row.ID] = resists[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getElementalResists() ([]ElementalResist, error) {
	resists := []ElementalResist{}
	var err error

	for i := range l.json.autoAbilities {
		autoAbility := &l.json.autoAbilities[i]

		if autoAbility.AddedElemResist == nil {
			continue
		}

		autoAbility.AddedElemResist.ElementID, err = assignFK(autoAbility.AddedElemResist.Element, l.Elements)
		if err != nil {
			return nil, err
		}

		resists = append(resists, *autoAbility.AddedElemResist)
	}

	for i := range l.json.monsters {
		monster := &l.json.monsters[i]

		for j := range monster.ElemResists {
			resist := &monster.ElemResists[j]

			resist.ElementID, err = assignFK(resist.Element, l.Elements)
			if err != nil {
				return nil, err
			}

			resists = append(resists, *resist)
		}

		for j := range monster.AlteredStates {
			stateResists, err := l.getAltStateElemResists(&monster.AlteredStates[j])
			if err != nil {
				return nil, err
			}

			resists = append(resists, stateResists...)
		}
	}

	for i := range l.json.statusConditions {
		condition := l.json.statusConditions[i]

		if condition.AddedElemResist == nil {
			continue
		}

		condition.AddedElemResist.ElementID, err = assignFK(condition.AddedElemResist.Element, l.Elements)
		if err != nil {
			return nil, err
		}

		resists = append(resists, *condition.AddedElemResist)
	}

	return resists, nil
}

func (l *Lookup) getAltStateElemResists(state *AlteredState) ([]ElementalResist, error) {
	elemResists := []ElementalResist{}
	var err error

	for i := range state.Changes {
		change := &state.Changes[i]

		if change.ElemResists == nil {
			continue
		}

		for j := range change.ElemResists {
			resist := &change.ElemResists[j]

			resist.ElementID, err = assignFK(resist.Element, l.Elements)
			if err != nil {
				return nil, err
			}
			elemResists = append(elemResists, *resist)
		}
	}

	return elemResists, nil
}

func (l *Lookup) loop2UpdateElements(qtx *database.Queries, ctx context.Context) error {
	elementsJson, err := l.getUpdatedElements()
	if err != nil {
		return err
	}
	elements := dedupeRows(elementsJson, l.Hashes)

	params := database.UpdateElementBulkParams{
		ID:                make([]int32, len(elements)),
		DataHash:          make([]string, len(elements)),
		OppositeElementID: make([]sql.NullInt32, len(elements)),
	}

	for i, e := range elements {
		params.ID[i] = e.ID
		params.DataHash[i] = generateDataHash(e)
		params.OppositeElementID[i] = h.GetNullInt32(e.OppositeElementID)
	}

	dbRows, err := qtx.UpdateElementBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't update elements: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getUpdatedElements() ([]Element, error) {
	elements := []Element{}
	var err error

	for i := range l.json.elements {
		element := &l.json.elements[i]

		if element.OppositeElement == nil {
			continue
		}

		delete(l.Hashes, generateDataHash(element))

		element.OppositeElementID, err = assignFKPtr(element.OppositeElement, l.Elements)
		if err != nil {
			return nil, err
		}

		elements = append(elements, *element)
	}

	return elements, nil
}

func (l *Lookup) loop2SeedAgilitySubtiers(qtx *database.Queries, ctx context.Context) error {
	subtiersJson := l.getAgilitySubtiers()
	subtiers := dedupeRows(subtiersJson, l.Hashes)

	params := database.CreateAgilitySubtierBulkParams{
		DataHash:        make([]string, len(subtiers)),
		AgilityTierID:   make([]int32, len(subtiers)),
		MinAgility:      make([]int32, len(subtiers)),
		MaxAgility:      make([]int32, len(subtiers)),
		CharacterMinIcv: make([]sql.NullInt32, len(subtiers)),
	}

	for i, s := range subtiers {
		params.DataHash[i] = generateDataHash(s)
		params.AgilityTierID[i] = s.AgilityTierID
		params.MinAgility[i] = s.MinAgility
		params.MaxAgility[i] = s.MaxAgility
		params.CharacterMinIcv[i] = h.GetNullInt32(s.CharacterMinICV)
	}

	dbRows, err := qtx.CreateAgilitySubtierBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create agility subtiers: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getAgilitySubtiers() []AgilitySubtier {
	subtiers := []AgilitySubtier{}

	for i := range l.json.agilityTiers {
		tier := &l.json.agilityTiers[i]

		for j := range tier.CharacterMinICVs {
			subtier := &tier.CharacterMinICVs[j]
			subtier.AgilityTierID = tier.ID
			subtiers = append(subtiers, *subtier)
		}
	}

	return subtiers
}

func (l *Lookup) loop1SeedEquipmentSlotsChances(qtx *database.Queries, ctx context.Context) error {
	chancesJson := l.getEquipmentSlotsChances()
	chances := dedupeRows(chancesJson, l.Hashes)

	params := database.CreateEquipmentSlotsChanceBulkParams{
		DataHash: make([]string, len(chances)),
		Amount:   make([]int32, len(chances)),
		Chance:   make([]int32, len(chances)),
	}

	for i, c := range chances {
		params.DataHash[i] = generateDataHash(c)
		params.Amount[i] = c.Amount
		params.Chance[i] = c.Chance
	}

	dbRows, err := qtx.CreateEquipmentSlotsChanceBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create equipment slot chances: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getEquipmentSlotsChances() []EquipmentSlotsChance {
	chances := []EquipmentSlotsChance{}

	for _, m := range l.json.monsters {
		if m.Equipment != nil {
			chances = slices.Concat(chances, m.Equipment.AbilitySlots.Chances)
			chances = slices.Concat(chances, m.Equipment.AttachedAbilities.Chances)
		}
	}

	return chances
}

func (l *Lookup) loop1SeedMonsterSelections(qtx *database.Queries, ctx context.Context) error {
	selectionsJson := l.getMonsterSelections()
	selections := dedupeRows(selectionsJson, l.Hashes)

	dataHashes := make([]string, len(selections))

	for i, s := range selections {
		dataHashes[i] = generateDataHash(s)
	}

	dbRows, err := qtx.CreateMonsterSelectionBulk(ctx, dataHashes)
	if err != nil {
		return fmt.Errorf("couldn't create monster selections: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getMonsterSelections() []MonsterSelection {
	selections := []MonsterSelection{}

	for _, mf := range l.json.monsterFormations {
		selections = append(selections, mf.MonsterSelection)
	}

	return selections
}

func (l *Lookup) loop1SeedMonsters(qtx *database.Queries, ctx context.Context) error {
	err := l.completeMonstersElements()
	if err != nil {
		return err
	}
	monstersJson := l.json.monsters
	monsters := dedupeRows(monstersJson, l.Hashes)

	params := database.CreateMonsterBulkParams{
		DataHash:             make([]string, len(monsters)),
		Name:                 make([]string, len(monsters)),
		Version:              make([]sql.NullInt32, len(monsters)),
		Specification:        make([]sql.NullString, len(monsters)),
		Notes:                make([]sql.NullString, len(monsters)),
		Species:              make([]database.MonsterSpecies, len(monsters)),
		Availability:         make([]database.AvailabilityType, len(monsters)),
		IsRepeatable:         make([]bool, len(monsters)),
		CanBeCaptured:        make([]bool, len(monsters)),
		AreaConquestLocation: make([]database.NullMaCreationArea, len(monsters)),
		Category:             make([]database.MonsterCategory, len(monsters)),
		CtbIconType:          make([]database.CtbIconType, len(monsters)),
		HasOverdrive:         make([]bool, len(monsters)),
		IsUnderwater:         make([]bool, len(monsters)),
		IsZombie:             make([]bool, len(monsters)),
		Distance:             make([]int32, len(monsters)),
		Ap:                   make([]int32, len(monsters)),
		ApOverkill:           make([]int32, len(monsters)),
		OverkillDamage:       make([]int32, len(monsters)),
		Gil:                  make([]int32, len(monsters)),
		StealGil:             make([]sql.NullInt32, len(monsters)),
		DoomCountdown:        make([]sql.NullInt32, len(monsters)),
		PoisonRate:           make([]sql.NullFloat64, len(monsters)),
		ThreatenChance:       make([]sql.NullInt32, len(monsters)),
		ZanmatoLevel:         make([]int32, len(monsters)),
		MonsterArenaPrice:    make([]sql.NullInt32, len(monsters)),
		SensorText:           make([]sql.NullString, len(monsters)),
		ScanText:             make([]sql.NullString, len(monsters)),
	}

	for i, m := range monsters {
		params.DataHash[i] = generateDataHash(m)
		params.Name[i] = m.Name
		params.Version[i] = h.GetNullInt32(m.Version)
		params.Specification[i] = h.GetNullString(m.Specification)
		params.Notes[i] = h.GetNullString(m.Notes)
		params.Species[i] = database.MonsterSpecies(m.Species)
		params.Availability[i] = database.AvailabilityType(m.Availability)
		params.IsRepeatable[i] = m.IsRepeatable
		params.CanBeCaptured[i] = m.CanBeCaptured
		params.AreaConquestLocation[i] = database.ToNullMaCreationArea(m.AreaConquestLocation)
		params.Category[i] = database.MonsterCategory(m.Category)
		params.CtbIconType[i] = database.CtbIconType(m.CTBIconType)
		params.HasOverdrive[i] = m.HasOverdrive
		params.IsUnderwater[i] = m.IsUnderwater
		params.IsZombie[i] = m.IsZombie
		params.Distance[i] = m.Distance
		params.Ap[i] = m.AP
		params.ApOverkill[i] = m.APOverkill
		params.OverkillDamage[i] = m.OverkillDamage
		params.Gil[i] = m.Gil
		params.StealGil[i] = h.GetNullInt32(m.StealGil)
		params.DoomCountdown[i] = h.GetNullInt32(m.DoomCountdown)
		params.PoisonRate[i] = h.GetNullFloat64(m.PoisonRate)
		params.ThreatenChance[i] = h.GetNullInt32(m.ThreatenChance)
		params.ZanmatoLevel[i] = m.ZanmatoLevel
		params.MonsterArenaPrice[i] = h.GetNullInt32(m.MonsterArenaPrice)
		params.SensorText[i] = h.GetNullString(m.SensorText)
		params.ScanText[i] = h.GetNullString(m.ScanText)
	}

	dbRows, err := qtx.CreateMonsterBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create monsters: %v", err)
	}

	for i, row := range dbRows {
		monsters[i].ID = row.ID
		l.json.monsters[i].ID = row.ID
		key := CreateLookupKey(monsters[i])
		l.Monsters[key] = monsters[i]
		l.MonstersID[row.ID] = monsters[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) completeMonstersElements() error {
	elements := []string{"fire", "lightning", "water", "ice", "holy"}

	for i, mon := range l.json.monsters {
		elemResistLookup := make(map[string]string)

		for _, elemResist := range mon.ElemResists {
			elemResistLookup[elemResist.Element] = elemResist.Affinity
		}

		resists := mon.ElemResists
		for _, element := range elements {
			_, found := elemResistLookup[element]
			if !found {
				elemResist := ElementalResist{
					Element:  element,
					Affinity: "neutral",
				}
				resists = append(resists, elemResist)
			}
		}
		l.json.monsters[i].ElemResists = resists
	}

	return nil
}

func (l *Lookup) loop1SeedInflictedDelays(qtx *database.Queries, ctx context.Context) error {
	delaysJson := l.getInflictedDelays()
	delays := dedupeRows(delaysJson, l.Hashes)

	params := database.CreateInflictedDelayBulkParams{
		DataHash:       make([]string, len(delays)),
		Condition:      make([]sql.NullString, len(delays)),
		CtbAttackType:  make([]database.CtbAttackType, len(delays)),
		DelayType:      make([]database.DelayType, len(delays)),
		DamageConstant: make([]int32, len(delays)),
	}

	for i, d := range delays {
		params.DataHash[i] = generateDataHash(d)
		params.Condition[i] = h.GetNullString(d.Condition)
		params.CtbAttackType[i] = database.CtbAttackType(d.CTBAttackType)
		params.DelayType[i] = database.DelayType(d.DelayType)
		params.DamageConstant[i] = d.DamageConstant
	}

	dbRows, err := qtx.CreateInflictedDelayBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create inflicted delays: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getInflictedDelays() []InflictedDelay {
	delays := []InflictedDelay{}

	for _, ability := range l.json.enemyAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range l.json.items {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range l.json.overdriveAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range l.json.playerAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range l.json.triggerCommands {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, ability := range l.json.unspecifiedAbilities {
		for _, bi := range ability.BattleInteractions {
			if bi.InflictedDelay != nil {
				delays = append(delays, *bi.InflictedDelay)
			}
		}
	}

	for _, status := range l.json.statusConditions {
		if status.CtbOnInfliction != nil {
			delays = append(delays, *status.CtbOnInfliction)
		}
	}

	return delays
}

func (l *Lookup) loop1SeedAbilityAccuracies(qtx *database.Queries, ctx context.Context) error {
	accuraciesJson := l.getAbilityAccuracies()
	accuracies := dedupeRows(accuraciesJson, l.Hashes)

	params := database.CreateAbilityAccuracyBulkParams{
		DataHash:    make([]string, len(accuracies)),
		AccSource:   make([]database.AccSourceType, len(accuracies)),
		HitChance:   make([]sql.NullInt32, len(accuracies)),
		AccModifier: make([]sql.NullFloat64, len(accuracies)),
	}

	for i, a := range accuracies {
		params.DataHash[i] = generateDataHash(a)
		params.AccSource[i] = database.AccSourceType(a.AccSource)
		params.HitChance[i] = h.GetNullInt32(a.HitChance)
		params.AccModifier[i] = h.GetNullFloat64(a.AccModifier)
	}

	dbRows, err := qtx.CreateAbilityAccuracyBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create ability accuracies: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getAbilityAccuracies() []Accuracy {
	accuracies := []Accuracy{}

	for _, ability := range l.json.enemyAbilities {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range l.json.items {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range l.json.overdriveAbilities {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range l.json.playerAbilities {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range l.json.triggerCommands {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, ability := range l.json.unspecifiedAbilities {
		for _, bi := range ability.BattleInteractions {
			accuracies = append(accuracies, bi.Accuracy)
		}
	}

	for _, aeon := range l.json.aeons {
		if aeon.PhysAtkAccuracy != nil {
			accuracies = append(accuracies, *aeon.PhysAtkAccuracy)
		}
	}

	return accuracies
}

func (l *Lookup) loop1SeedSongCredits(qtx *database.Queries, ctx context.Context) error {
	creditsJson := l.getSongCredits()
	credits := dedupeRows(creditsJson, l.Hashes)

	params := database.CreateSongCreditBulkParams{
		DataHash:  make([]string, len(credits)),
		Composer:  make([]database.NullComposer, len(credits)),
		Arranger:  make([]database.NullArranger, len(credits)),
		Performer: make([]sql.NullString, len(credits)),
		Lyricist:  make([]sql.NullString, len(credits)),
	}

	for i, c := range credits {
		params.DataHash[i] = generateDataHash(c)
		params.Composer[i] = database.ToNullComposer(c.Composer)
		params.Arranger[i] = database.ToNullArranger(c.Arranger)
		params.Performer[i] = h.GetNullString(c.Performer)
		params.Lyricist[i] = h.GetNullString(c.Lyricist)
	}

	dbRows, err := qtx.CreateSongCreditBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create song credits: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getSongCredits() []SongCredits {
	credits := []SongCredits{}

	for _, song := range l.json.songs {
		if song.Credits != nil {
			credits = append(credits, *song.Credits)
		}
	}

	return credits
}

func (l *Lookup) loop1SeedBackgroundMusic(qtx *database.Queries, ctx context.Context) error {
	bmJson := l.getBackgroundMusic()
	bm := dedupeRows(bmJson, l.Hashes)

	params := database.CreateBackgroundMusicBulkParams{
		DataHash:               make([]string, len(bm)),
		Condition:              make([]sql.NullString, len(bm)),
		ReplacesEncounterMusic: make([]bool, len(bm)),
	}

	for i, music := range bm {
		params.DataHash[i] = generateDataHash(music)
		params.Condition[i] = h.GetNullString(music.Condition)
		params.ReplacesEncounterMusic[i] = music.ReplacesEncounterMusic
	}

	dbRows, err := qtx.CreateBackgroundMusicBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create background music: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getBackgroundMusic() []BackgroundMusic {
	bgMusic := []BackgroundMusic{}

	for _, song := range l.json.songs {
		bgMusic = slices.Concat(bgMusic, song.BackgroundMusic)
	}

	return bgMusic
}

func (l *Lookup) loop1SeedLocations(qtx *database.Queries, ctx context.Context) error {
	locationsJson := l.json.locations
	locations := dedupeRows(locationsJson, l.Hashes)

	params := database.CreateLocationBulkParams{
		DataHash: make([]string, len(locations)),
		Name:     make([]string, len(locations)),
	}

	for i, mi := range locations {
		params.DataHash[i] = generateDataHash(mi)
		params.Name[i] = mi.Name
	}

	dbRows, err := qtx.CreateLocationBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create locations: %v", err)
	}

	for i, row := range dbRows {
		locations[i].ID = row.ID
		l.json.locations[i].ID = row.ID
		l.Locations[locations[i].Name] = locations[i]
		l.LocationsID[row.ID] = locations[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) loop1SeedCreatedNodes(qtx *database.Queries, ctx context.Context) error {
	nodesJson := l.getCreatedNodes()
	nodes := dedupeRows(nodesJson, l.Hashes)

	params := database.CreateCreatedNodeBulkParams{
		DataHash: make([]string, len(nodes)),
		Node:     make([]database.NodeType, len(nodes)),
		Value:    make([]int32, len(nodes)),
	}

	for i, mi := range nodes {
		params.DataHash[i] = generateDataHash(mi)
		params.Node[i] = database.NodeType(mi.Node)
		params.Value[i] = mi.Value
	}

	dbRows, err := qtx.CreateCreatedNodeBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create created nodes: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getCreatedNodes() []CreatedNode {
	createdNodes := []CreatedNode{}

	for _, sphere := range l.json.spheres {
		if sphere.CreatedNode != nil {
			createdNodes = append(createdNodes, *sphere.CreatedNode)
		}
	}

	return createdNodes
}

func (l *Lookup) loop1SeedMasterItems(qtx *database.Queries, ctx context.Context) error {
	itemsJson := l.getMasterItems()
	items := dedupeRows(itemsJson, l.Hashes)

	params := database.CreateMasterItemBulkParams{
		DataHash: make([]string, len(items)),
		Name:     make([]string, len(items)),
		Type:     make([]database.ItemType, len(items)),
	}

	for i, mi := range items {
		params.DataHash[i] = generateDataHash(mi)
		params.Name[i] = mi.Name
		params.Type[i] = mi.Type
	}

	dbRows, err := qtx.CreateMasterItemBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create master items: %v", err)
	}

	for i, row := range dbRows {
		items[i].ID = row.ID
		key := CreateLookupKey(items[i])
		l.MasterItems[key] = items[i]
		l.MasterItemsID[row.ID] = items[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getMasterItems() []MasterItem {
	masterItems := []MasterItem{}

	for _, i := range l.json.items {
		i.MasterItem.Type = database.ItemTypeItem
		masterItems = append(masterItems, i.MasterItem)
	}

	for _, i := range l.json.keyItems {
		i.MasterItem.Type = database.ItemTypeKeyItem
		masterItems = append(masterItems, i.MasterItem)
	}

	return masterItems
}

func (l *Lookup) loop1SeedAbilityAttributes(qtx *database.Queries, ctx context.Context) error {
	attributesJson := l.getAbilityAttributes()
	attributes := dedupeRows(attributesJson, l.Hashes)

	params := database.CreateAbilityAttributesBulkParams{
		DataHash:         make([]string, len(attributes)),
		Rank:             make([]sql.NullInt32, len(attributes)),
		AppearsInHelpBar: make([]bool, len(attributes)),
		CanCopycat:       make([]bool, len(attributes)),
	}

	for i, a := range attributes {
		params.DataHash[i] = generateDataHash(a)
		params.Rank[i] = h.GetNullInt32(a.Rank)
		params.AppearsInHelpBar[i] = a.AppearsInHelpBar
		params.CanCopycat[i] = a.CanCopycat
	}

	dbRows, err := qtx.CreateAbilityAttributesBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create ability attributes: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getAbilityAttributes() []Attributes {
	attributes := []Attributes{}

	for _, ability := range l.json.enemyAbilities {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range l.json.items {
		if len(ability.BattleInteractions) > 0 {
			attributes = append(attributes, ability.Attributes)
		}
	}

	for _, ability := range l.json.overdrives {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range l.json.overdriveAbilities {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range l.json.playerAbilities {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range l.json.triggerCommands {
		attributes = append(attributes, ability.Attributes)
	}

	for _, ability := range l.json.unspecifiedAbilities {
		attributes = append(attributes, ability.Attributes)
	}

	return attributes
}

func (l *Lookup) loop1SeedTopmenus(qtx *database.Queries, ctx context.Context) error {
	topmenusJson := l.json.topmenus
	topmenus := dedupeRows(topmenusJson, l.Hashes)

	params := database.CreateTopmenuBulkParams{
		DataHash: make([]string, len(topmenus)),
		Name:     make([]string, len(topmenus)),
	}

	for i, c := range topmenus {
		params.DataHash[i] = generateDataHash(c)
		params.Name[i] = c.Name
	}

	dbRows, err := qtx.CreateTopmenuBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create topmenu: %v", err)
	}

	for i, row := range dbRows {
		topmenus[i].ID = row.ID
		l.json.topmenus[i].ID = row.ID
		l.Topmenus[topmenus[i].Name] = topmenus[i]
		l.TopmenusID[row.ID] = topmenus[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) loop1SeedBlitzballPositions(qtx *database.Queries, ctx context.Context) error {
	positionsJson := l.json.blitzballPositions
	positions := dedupeRows(positionsJson, l.Hashes)

	params := database.CreateBlitzballPositionBulkParams{
		DataHash: make([]string, len(positions)),
		Category: make([]database.BlitzballTournamentCategory, len(positions)),
		Slot:     make([]database.BlitzballPositionSlot, len(positions)),
	}

	for i, p := range positions {
		params.DataHash[i] = generateDataHash(p)
		params.Category[i] = database.BlitzballTournamentCategory(p.Category)
		params.Slot[i] = database.BlitzballPositionSlot(p.Slot)
	}

	dbRows, err := qtx.CreateBlitzballPositionBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create blitzball position: %v", err)
	}

	for i, row := range dbRows {
		positions[i].ID = row.ID
		l.json.blitzballPositions[i].ID = row.ID
		key := CreateLookupKey(positions[i])
		l.Positions[key] = positions[i]
		l.PositionsID[row.ID] = positions[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) loop1SeedCharacterClasses(qtx *database.Queries, ctx context.Context) error {
	classesJson := l.json.characterClasses
	classes := dedupeRows(classesJson, l.Hashes)

	params := database.CreateCharacterClassBulkParams{
		DataHash: make([]string, len(classes)),
		Name:     make([]string, len(classes)),
		Category: make([]database.CharacterClassCategory, len(classes)),
	}

	for i, c := range classes {
		params.DataHash[i] = generateDataHash(c)
		params.Name[i] = c.Name
		params.Category[i] = database.CharacterClassCategory(c.Category)
	}

	dbRows, err := qtx.CreateCharacterClassBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create character classes: %v", err)
	}

	for i, row := range dbRows {
		classes[i].ID = row.ID
		l.json.characterClasses[i].ID = row.ID
		l.CharClasses[classes[i].Name] = classes[i]
		l.CharClassesID[row.ID] = classes[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) loop1SeedPlayerUnits(qtx *database.Queries, ctx context.Context) error {
	unitsJson := l.getPlayerUnits()
	units := dedupeRows(unitsJson, l.Hashes)

	params := database.CreatePlayerUnitBulkParams{
		DataHash: make([]string, len(units)),
		Name:     make([]string, len(units)),
		Type:     make([]database.UnitType, len(units)),
	}

	for i, m := range units {
		params.DataHash[i] = generateDataHash(m)
		params.Name[i] = m.Name
		params.Type[i] = m.Type
	}

	dbRows, err := qtx.CreatePlayerUnitBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create player units: %v", err)
	}

	for i, row := range dbRows {
		units[i].ID = row.ID
		key := CreateLookupKey(units[i])
		l.PlayerUnits[key] = units[i]
		l.PlayerUnitsID[row.ID] = units[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) getPlayerUnits() []PlayerUnit {
	playerUnits := []PlayerUnit{}

	for _, c := range l.json.characters {
		c.PlayerUnit.Type = database.UnitTypeCharacter
		playerUnits = append(playerUnits, c.PlayerUnit)
	}

	for _, a := range l.json.aeons {
		a.PlayerUnit.Type = database.UnitTypeAeon
		playerUnits = append(playerUnits, a.PlayerUnit)
	}

	return playerUnits
}

func (l *Lookup) loop1SeedModifiers(qtx *database.Queries, ctx context.Context) error {
	modifiersJson := l.json.modifiers
	modifiers := dedupeRows(modifiersJson, l.Hashes)

	params := database.CreateModifierBulkParams{
		DataHash:     make([]string, len(modifiers)),
		Name:         make([]string, len(modifiers)),
		Effect:       make([]string, len(modifiers)),
		Category:     make([]database.ModifierCategory, len(modifiers)),
		DefaultValue: make([]sql.NullFloat64, len(modifiers)),
	}

	for i, m := range modifiers {
		params.DataHash[i] = generateDataHash(m)
		params.Name[i] = m.Name
		params.Effect[i] = m.Effect
		params.Category[i] = database.ModifierCategory(m.Category)
		params.DefaultValue[i] = h.GetNullFloat64(m.DefaultValue)
	}

	dbRows, err := qtx.CreateModifierBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create modifiers: %v", err)
	}

	for i, row := range dbRows {
		modifiers[i].ID = row.ID
		l.json.modifiers[i].ID = row.ID
		l.Modifiers[modifiers[i].Name] = modifiers[i]
		l.ModifiersID[row.ID] = modifiers[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) loop1SeedProperties(qtx *database.Queries, ctx context.Context) error {
	propertiesJson := l.json.properties
	properties := dedupeRows(propertiesJson, l.Hashes)

	params := database.CreatePropertyBulkParams{
		DataHash:       make([]string, len(properties)),
		Name:           make([]string, len(properties)),
		Effect:         make([]string, len(properties)),
		NullifyArmored: make([]database.NullNullifyArmored, len(properties)),
	}

	for i, p := range properties {
		params.DataHash[i] = generateDataHash(p)
		params.Name[i] = p.Name
		params.Effect[i] = p.Effect
		params.NullifyArmored[i] = database.ToNullNullifyArmored(p.NullifyArmored)
	}

	dbRows, err := qtx.CreatePropertyBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create properties: %v", err)
	}

	for i, row := range dbRows {
		properties[i].ID = row.ID
		l.json.properties[i].ID = row.ID
		l.Properties[properties[i].Name] = properties[i]
		l.PropertiesID[row.ID] = properties[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) loop1SeedOverdriveModes(qtx *database.Queries, ctx context.Context) error {
	modesJson := l.json.overdriveModes
	modes := dedupeRows(modesJson, l.Hashes)

	params := database.CreateOverdriveModeBulkParams{
		DataHash:    make([]string, len(modes)),
		Name:        make([]string, len(modes)),
		Description: make([]string, len(modes)),
		Effect:      make([]string, len(modes)),
		Type:        make([]database.OverdriveModeType, len(modes)),
		FillRate:    make([]sql.NullFloat64, len(modes)),
	}

	for i, m := range modes {
		params.DataHash[i] = generateDataHash(m)
		params.Name[i] = m.Name
		params.Description[i] = m.Description
		params.Effect[i] = m.Effect
		params.Type[i] = database.OverdriveModeType(m.Type)
		params.FillRate[i] = h.GetNullFloat64(m.FillRate)
	}

	dbRows, err := qtx.CreateOverdriveModeBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create overdrive modes: %v", err)
	}

	for i, row := range dbRows {
		modes[i].ID = row.ID
		l.json.overdriveModes[i].ID = row.ID
		l.OverdriveModes[modes[i].Name] = modes[i]
		l.OverdriveModesID[row.ID] = modes[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) loop1SeedElements(qtx *database.Queries, ctx context.Context) error {
	elementsJson := l.json.elements
	elements := dedupeRows(elementsJson, l.Hashes)

	params := database.CreateElementBulkParams{
		DataHash: make([]string, len(elements)),
		Name:     make([]string, len(elements)),
	}

	for i, e := range elements {
		params.DataHash[i] = generateDataHash(e)
		params.Name[i] = e.Name
	}

	dbRows, err := qtx.CreateElementBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create elements: %v", err)
	}

	for i, row := range dbRows {
		elements[i].ID = row.ID
		l.json.elements[i].ID = row.ID
		l.Elements[elements[i].Name] = elements[i]
		l.ElementsID[row.ID] = elements[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) loop1SeedAgilityTiers(qtx *database.Queries, ctx context.Context) error {
	agilityTiersJson := l.json.agilityTiers
	agilityTiers := dedupeRows(agilityTiersJson, l.Hashes)

	params := database.CreateAgilityTierBulkParams{
		DataHash:        make([]string, len(agilityTiers)),
		MinAgility:      make([]int32, len(agilityTiers)),
		MaxAgility:      make([]int32, len(agilityTiers)),
		TickSpeed:       make([]int32, len(agilityTiers)),
		MonsterMinIcv:   make([]sql.NullInt32, len(agilityTiers)),
		MonsterMaxIcv:   make([]sql.NullInt32, len(agilityTiers)),
		CharacterMaxIcv: make([]sql.NullInt32, len(agilityTiers)),
	}

	for i, a := range agilityTiers {
		params.DataHash[i] = generateDataHash(a)
		params.MinAgility[i] = a.MinAgility
		params.MaxAgility[i] = a.MaxAgility
		params.TickSpeed[i] = a.TickSpeed
		params.MonsterMinIcv[i] = h.GetNullInt32(a.MonsterMinICV)
		params.MonsterMaxIcv[i] = h.GetNullInt32(a.MonsterMaxICV)
		params.CharacterMaxIcv[i] = h.GetNullInt32(a.CharacterMaxICV)
	}

	dbRows, err := qtx.CreateAgilityTierBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create agility tiers: %v", err)
	}

	for i, row := range dbRows {
		agilityTiers[i].ID = row.ID
		l.json.agilityTiers[i].ID = row.ID
		l.AgilityTiersID[row.ID] = agilityTiers[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}
