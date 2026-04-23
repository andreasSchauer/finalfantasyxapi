package seeding

import (
	"context"
	"database/sql"
	"errors"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

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
		return errors.New("couldn't create player units.")
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
	aeonsJson := l.json.aeons
	charsJson := l.json.characters
	playerUnits := []PlayerUnit{}

	for _, a := range aeonsJson {
		a.PlayerUnit.Type = database.UnitTypeAeon
		playerUnits = append(playerUnits, a.PlayerUnit)
	}

	for _, c := range charsJson {
		c.PlayerUnit.Type = database.UnitTypeCharacter
		playerUnits = append(playerUnits, c.PlayerUnit)
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
		return errors.New("couldn't create modifiers.")
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
		return errors.New("couldn't create properties.")
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
		return errors.New("couldn't create overdrive modes.")
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
		return errors.New("couldn't create elements.")
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
		return errors.New("couldn't create agility tiers.")
	}

	for i, row := range dbRows {
		agilityTiers[i].ID = row.ID
		l.AgilityTiersID[row.ID] = agilityTiers[i]

		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func Seed(db *database.Queries, dbConn *sql.DB) (*Lookup, error) {
	l := lookupInit()
	err := l.loadJSONFiles()
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

	checkErr(loadJSONFile("data/aeons.json", &l.json.aeons))
	checkErr(loadJSONFile("data/agility_tiers.json", &l.json.agilityTiers))
	checkErr(loadJSONFile("data/characters.json", &l.json.characters))
	checkErr(loadJSONFile("data/character_classes.json", &l.json.characterClasses))
	checkErr(loadJSONFile("data/elements.json", &l.json.elements))
	checkErr(loadJSONFile("data/modifiers.json", &l.json.modifiers))
	checkErr(loadJSONFile("data/overdrive_modes.json", &l.json.overdriveModes))
	checkErr(loadJSONFile("data/properties.json", &l.json.properties))

	return err
}

type jsonLookup struct {
	aeons				[]Aeon
	agilityTiers 		[]AgilityTier
	characters			[]Character
	characterClasses	[]CharacterClass
	elements	 		[]Element
	modifiers			[]Modifier
	properties			[]Property
	overdriveModes		[]OverdriveMode
}
