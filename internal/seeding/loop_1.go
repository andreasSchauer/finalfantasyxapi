package seeding

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

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
		if ability.Attributes != nil {
			attributes = append(attributes, *ability.Attributes)
		}
	}
	
	for _, ability := range l.json.items {
		if ability.Attributes != nil && len(ability.BattleInteractions) > 0 {
			attributes = append(attributes, *ability.Attributes)
		}
	}

	for _, ability := range l.json.overdrives {
		if ability.Attributes != nil {
			attributes = append(attributes, *ability.Attributes)
		}
	}
	
	for _, ability := range l.json.playerAbilities {
		if ability.Attributes != nil {
			attributes = append(attributes, *ability.Attributes)
		}
	}

	for _, ability := range l.json.triggerCommands {
		if ability.Attributes != nil {
			attributes = append(attributes, *ability.Attributes)
		}
	}
	
	for _, ability := range l.json.unspecifiedAbilities {
		if ability.Attributes != nil {
			attributes = append(attributes, *ability.Attributes)
		}
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

func Seed(db *database.Queries, dbConn *sql.DB) (*Lookup, error) {
	const migrationsDir = "sql/schema/"
	fullPath, err := h.GetAbsoluteFilepath(migrationsDir)
	if err != nil {
		return nil, err
	}

	start := time.Now()

	err = setupDB(dbConn, fullPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't setup database: %v", err)
	}
	
	l := lookupInit()
	err = l.loadJSONFiles()
	if err != nil {
		return nil, err
	}

	err = queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		err = l.seedLoop1(qtx, context.Background())
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	
	totalDuration := time.Since(start)
	fmt.Printf("database seeding took %.3f seconds\n\n", totalDuration.Seconds())

	return l, nil
}

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
	})
}

func (l *Lookup) seedLoop(qtx *database.Queries, ctx context.Context, fns []func(*database.Queries, context.Context) error) error {
	for _, fn := range fns {
		err := fn(qtx, ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func dedupeRows[T Hashable](rows []T, hashes map[string]int32) []T {
    seen := make(map[string]bool)
    ordered := []T{}

    for _, row := range rows {
        hash := generateDataHash(row)
		_, ok := hashes[hash]
		if ok || seen[hash] {
			continue
		}

		seen[hash] = true
		ordered = append(ordered, row)
    }
    return ordered
}



func (l *Lookup) loadJSONFiles() error {
	l.json = jsonLookup{}
	var err error

	checkErr := func(e error) {
		if err != nil {
			return
		}

		err = e
	}

	checkErr(loadJSONFile("data/aeon_commands.json", &l.json.aeonCommands))
	checkErr(loadJSONFile("data/aeon_stats.json", &l.json.aeonStats))
	checkErr(loadJSONFile("data/aeons.json", &l.json.aeons))
	checkErr(loadJSONFile("data/agility_tiers.json", &l.json.agilityTiers))
	checkErr(loadJSONFile("data/auto_abilities.json", &l.json.autoAbilities))
	checkErr(loadJSONFile("data/blitzball_items.json", &l.json.blitzballPositions))
	checkErr(loadJSONFile("data/celestial_weapons.json", &l.json.celestialWeapons))
	checkErr(loadJSONFile("data/character_classes.json", &l.json.characterClasses))
	checkErr(loadJSONFile("data/characters.json", &l.json.characters))
	checkErr(loadJSONFile("data/default_abilities.json", &l.json.defaultAbilities))
	checkErr(loadJSONFile("data/elements.json", &l.json.elements))
	checkErr(loadJSONFile("data/enemy_abilities.json", &l.json.enemyAbilities))
	checkErr(loadJSONFile("data/equipment.json", &l.json.equipment))
	checkErr(loadJSONFile("data/fmvs.json", &l.json.fmvs))
	checkErr(loadJSONFile("data/items.json", &l.json.items))
	checkErr(loadJSONFile("data/key_items.json", &l.json.keyItems))
	checkErr(loadJSONFile("data/locations.json", &l.json.locations))
	checkErr(loadJSONFile("data/mixes.json", &l.json.mixes))
	checkErr(loadJSONFile("data/modifiers.json", &l.json.modifiers))
	checkErr(loadJSONFile("data/monster_arena_creations.json", &l.json.monsterArenaCreations))
	checkErr(loadJSONFile("data/monster_formations.json", &l.json.monsterFormations))
	checkErr(loadJSONFile("data/monsters.json", &l.json.monsters))
	checkErr(loadJSONFile("data/overdrive_abilities.json", &l.json.overdriveAbilities))
	checkErr(loadJSONFile("data/overdrive_commands.json", &l.json.overdriveCommands))
	checkErr(loadJSONFile("data/overdrive_modes.json", &l.json.overdriveModes))
	checkErr(loadJSONFile("data/overdrives.json", &l.json.overdrives))
	checkErr(loadJSONFile("data/player_abilities.json", &l.json.playerAbilities))
	checkErr(loadJSONFile("data/primers.json", &l.json.primers))
	checkErr(loadJSONFile("data/properties.json", &l.json.properties))
	checkErr(loadJSONFile("data/shops.json", &l.json.shops))
	checkErr(loadJSONFile("data/sidequests.json", &l.json.sidequests))
	checkErr(loadJSONFile("data/songs.json", &l.json.songs))
	checkErr(loadJSONFile("data/spheres.json", &l.json.spheres))
	checkErr(loadJSONFile("data/stats.json", &l.json.stats))
	checkErr(loadJSONFile("data/status_conditions.json", &l.json.statusConditions))
	checkErr(loadJSONFile("data/submenus.json", &l.json.submenus))
	checkErr(loadJSONFile("data/topmenus.json", &l.json.topmenus))
	checkErr(loadJSONFile("data/treasures.json", &l.json.treasures))
	checkErr(loadJSONFile("data/trigger_commands.json", &l.json.triggerCommands))
	checkErr(loadJSONFile("data/unspecified_abilities.json", &l.json.unspecifiedAbilities))

	return err
}

type jsonLookup struct {
	aeonCommands			[]AeonCommand
	aeonStats				[]AeonStat
	aeons					[]Aeon
	agilityTiers 			[]AgilityTier
	autoAbilities			[]AutoAbility
	blitzballPositions		[]BlitzballPosition
	celestialWeapons		[]CelestialWeapon
	characterClasses		[]CharacterClass
	characters				[]Character
	defaultAbilities		[]DefaultAbilitiesEntry
	elements	 			[]Element
	enemyAbilities			[]EnemyAbility
	equipment				[]EquipmentTable
	fmvs					[]FMV
	items					[]Item
	keyItems				[]KeyItem
	locations				[]Location
	mixes					[]Mix
	modifiers				[]Modifier
	monsterArenaCreations	[]ArenaCreation
	monsterFormations		[]MonsterFormation
	monsters				[]Monster
	overdriveAbilities		[]OverdriveAbility
	overdriveCommands		[]OverdriveCommand
	overdriveModes			[]OverdriveMode
	overdrives				[]Overdrive
	playerAbilities			[]PlayerAbility
	primers					[]Primer
	properties				[]Property
	shops					[]Shop
	sidequests				[]Sidequest
	songs					[]Song
	spheres					[]Sphere
	stats					[]Stat
	statusConditions		[]StatusCondition
	submenus				[]Submenu
	topmenus				[]Topmenu
	treasures				[]Treasure
	triggerCommands			[]TriggerCommand
	unspecifiedAbilities	[]UnspecifiedAbility
}
