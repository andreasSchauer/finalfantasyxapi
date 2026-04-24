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
		
	})
}

func (l *Lookup) loop1SeedEquipmentSlotsChances(qtx *database.Queries, ctx context.Context) error {
	chancesJson := l.getEquipmentSlotsChances()
	chances := dedupeRows(chancesJson, l.Hashes)

	params := database.CreateEquipmentSlotsChanceBulkParams{
		DataHash: 	make([]string, len(chances)),
		Amount: 	make([]int32, len(chances)),
		Chance: 	make([]int32, len(chances)),
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

	for i, row := range dbRows {
		chances[i].ID = row.ID
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

	for i, row := range dbRows {
		selections[i].ID = row.ID
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
	monstersJson := l.json.monsters
	monsters := dedupeRows(monstersJson, l.Hashes)

	params := database.CreateMonsterBulkParams{
		DataHash: 				make([]string, len(monsters)),
		Name: 					make([]string, len(monsters)),
		Version: 				make([]sql.NullInt32, len(monsters)),
		Specification: 			make([]sql.NullString, len(monsters)),
		Notes: 					make([]sql.NullString, len(monsters)),
		Species: 				make([]database.MonsterSpecies, len(monsters)),
		Availability: 			make([]database.AvailabilityType, len(monsters)),
		IsRepeatable: 			make([]bool, len(monsters)),
		CanBeCaptured: 			make([]bool, len(monsters)),
		AreaConquestLocation: 	make([]database.NullMaCreationArea, len(monsters)),
		Category: 				make([]database.MonsterCategory, len(monsters)),
		CtbIconType: 			make([]database.CtbIconType, len(monsters)),
		HasOverdrive: 			make([]bool, len(monsters)),
		IsUnderwater: 			make([]bool, len(monsters)),
		IsZombie: 				make([]bool, len(monsters)),
		Distance: 				make([]int32, len(monsters)),
		Ap: 					make([]int32, len(monsters)),
		ApOverkill: 			make([]int32, len(monsters)),
		OverkillDamage: 		make([]int32, len(monsters)),
		Gil: 					make([]int32, len(monsters)),
		StealGil: 				make([]sql.NullInt32, len(monsters)),
		DoomCountdown: 			make([]sql.NullInt32, len(monsters)),
		PoisonRate: 			make([]sql.NullFloat64, len(monsters)),
		ThreatenChance: 		make([]sql.NullInt32, len(monsters)),
		ZanmatoLevel: 			make([]int32, len(monsters)),
		MonsterArenaPrice: 		make([]sql.NullInt32, len(monsters)),
		SensorText:				make([]sql.NullString, len(monsters)),
		ScanText: 				make([]sql.NullString, len(monsters)),
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
		key := CreateLookupKey(monsters[i])
		l.Monsters[key] = monsters[i]
		l.MonstersID[row.ID] = monsters[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) loop1SeedInflictedDelays(qtx *database.Queries, ctx context.Context) error {
	delaysJson := l.getInflictedDelays()
	delays := dedupeRows(delaysJson, l.Hashes)

	params := database.CreateInflictedDelayBulkParams{
		DataHash: 		make([]string, len(delays)),
		Condition: 		make([]sql.NullString, len(delays)),
		CtbAttackType: 	make([]database.CtbAttackType, len(delays)),
		DelayType: 		make([]database.DelayType, len(delays)),
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

	for i, row := range dbRows {
		delays[i].ID = row.ID
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
		DataHash: 		make([]string, len(accuracies)),
		AccSource: 		make([]database.AccSourceType, len(accuracies)),
		HitChance: 		make([]sql.NullInt32, len(accuracies)),
		AccModifier: 	make([]sql.NullFloat64, len(accuracies)),
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

	for i, row := range dbRows {
		accuracies[i].ID = row.ID
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
		DataHash: 	make([]string, len(credits)),
		Composer: 	make([]database.NullComposer, len(credits)),
		Arranger: 	make([]database.NullArranger, len(credits)),
		Performer: 	make([]sql.NullString, len(credits)),
		Lyricist: 	make([]sql.NullString, len(credits)),
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

	for i, row := range dbRows {
		credits[i].ID = row.ID
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
		DataHash: 				make([]string, len(bm)),
		Condition: 				make([]sql.NullString, len(bm)),
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

	for i, row := range dbRows {
		bm[i].ID = row.ID
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

