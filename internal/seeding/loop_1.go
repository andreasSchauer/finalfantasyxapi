package seeding

import (
	"context"
	"database/sql"
	"fmt"

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
	})
}

func (l *Lookup) loop1SeedLocations(qtx *database.Queries, ctx context.Context) error {
	locationsJson := l.json.locations
	locations := dedupeRows(locationsJson, l.Hashes)

	params := database.CreateLocationBulkParams{
		DataHash: 	make([]string, len(locations)),
		Name: 		make([]string, len(locations)),
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
		DataHash: 	make([]string, len(nodes)),
		Node: 		make([]database.NodeType, len(nodes)),
		Value: 		make([]int32, len(nodes)),
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

	for i, row := range dbRows {
		nodes[i].ID = row.ID
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
		DataHash: 	make([]string, len(items)),
		Name: 		make([]string, len(items)),
		Type: 		make([]database.ItemType, len(items)),
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
		DataHash: 			make([]string, len(attributes)),
		Rank:				make([]sql.NullInt32, len(attributes)),
		AppearsInHelpBar: 	make([]bool, len(attributes)),
		CanCopycat:			make([]bool, len(attributes)),
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

	for i, row := range dbRows {
		attributes[i].ID = row.ID
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
		DataHash: 	make([]string, len(topmenus)),
		Name: 		make([]string, len(topmenus)),
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
		DataHash: 	make([]string, len(positions)),
		Category: 	make([]database.BlitzballTournamentCategory, len(positions)),
		Slot: 		make([]database.BlitzballPositionSlot, len(positions)),
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
		DataHash: 	make([]string, len(classes)),
		Name: 		make([]string, len(classes)),
		Category: 	make([]database.CharacterClassCategory, len(classes)),
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
		DataHash: 	make([]string, len(units)),
		Name: 		make([]string, len(units)),
		Type: 		make([]database.UnitType, len(units)),
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
		DataHash: 		make([]string, len(modifiers)),
		Name: 			make([]string, len(modifiers)),
		Effect: 		make([]string, len(modifiers)),
		Category: 		make([]database.ModifierCategory, len(modifiers)),
		DefaultValue: 	make([]sql.NullFloat64, len(modifiers)),
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
		DataHash: 		make([]string, len(properties)),
		Name: 			make([]string, len(properties)),
		Effect: 		make([]string, len(properties)),
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
		DataHash: 		make([]string, len(modes)),
		Name: 			make([]string, len(modes)),
		Description: 	make([]string, len(modes)),
		Effect: 		make([]string, len(modes)),
		Type: 			make([]database.OverdriveModeType, len(modes)),
		FillRate: 		make([]sql.NullFloat64, len(modes)),
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
		DataHash: 	make([]string, len(elements)),
		Name: 		make([]string, len(elements)),
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
		DataHash: 			make([]string, len(agilityTiers)),
		MinAgility: 		make([]int32, len(agilityTiers)),
		MaxAgility: 		make([]int32, len(agilityTiers)),
		TickSpeed: 			make([]int32, len(agilityTiers)),
		MonsterMinIcv: 		make([]sql.NullInt32, len(agilityTiers)),
		MonsterMaxIcv: 		make([]sql.NullInt32, len(agilityTiers)),
		CharacterMaxIcv: 	make([]sql.NullInt32, len(agilityTiers)),
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
		l.AgilityTiersID[row.ID] = agilityTiers[i]

		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

