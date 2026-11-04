package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func loop1SeedMonsters(qtx *database.Queries, ctx context.Context) error {
	err := completeMonstersElements()
	if err != nil {
		return err
	}
	monsters := dedupeRows(rawJSON.monsters, hashes)

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
		rawJSON.monsters[i].ID = row.ID
		lookup.L.Monsters[seedtypes.Key(monsters[i])] = monsters[i]
		lookup.L.MonstersID[row.ID] = monsters[i]
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func completeMonsters() error {
	for i := range rawJSON.monsters {
		mon := &rawJSON.monsters[i]

		err := assignIDs(mon.ElemResists)
		if err != nil {
			return err
		}

		err = assignIDs(mon.StatusResists)
		if err != nil {
			return err
		}

		err = assignIDs(mon.BaseStats)
		if err != nil {
			return err
		}

		err = assignIDs(mon.Abilities)
		if err != nil {
			return err
		}

		err = completeMonsterItems(mon.Items)
		if err != nil {
			return err
		}

		err = completeMonsterEquipment(mon.Equipment)
		if err != nil {
			return err
		}

		err = completeAlteredStates(mon.AlteredStates)
		if err != nil {
			return err
		}

		lookup.L.Monsters[seedtypes.Key(*mon)] = *mon
		lookup.L.MonstersID[mon.ID] = *mon
	}

	return nil
}

func completeMonstersElements() error {
	elements := []string{"fire", "lightning", "water", "ice", "holy"}

	for i, mon := range rawJSON.monsters {
		elemResistLookup := make(map[string]string)

		for _, elemResist := range mon.ElemResists {
			elemResistLookup[elemResist.Element] = elemResist.Affinity
		}

		resists := mon.ElemResists
		for _, element := range elements {
			_, found := elemResistLookup[element]
			if !found {
				elemResist := seedtypes.ElementalResist{
					Element:  element,
					Affinity: "neutral",
				}
				resists = append(resists, elemResist)
			}
		}
		rawJSON.monsters[i].ElemResists = resists
	}

	return nil
}
