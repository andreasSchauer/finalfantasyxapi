package seeding

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/pressly/goose/v3"
)

type seeder struct {
	name     string
	seedFunc func(*database.Queries, *sql.DB) error
	relFunc  func(*database.Queries, *sql.DB) error
}

func SeedDatabase(db *database.Queries, dbConn *sql.DB) (*Lookup, error) {
	const migrationsDir = "./sql/schema/"

	start := time.Now()

	err := setupDB(dbConn, migrationsDir)
	if err != nil {
		return nil, fmt.Errorf("couldn't setup database: %v", err)
	}

	l := lookupInit()
	seeders := l.getSeeders()

	err = handleDBFunctions(db, dbConn, seeders)
	if err != nil {
		return nil, err
	}

	totalDuration := time.Since(start)
	fmt.Printf("database seeding took %.3f seconds\n\n", totalDuration.Seconds())

	return &l, nil
}

func handleDBFunctions(db *database.Queries, dbConn *sql.DB, seeders []seeder) error {
	seedingStart := time.Now()
	fmt.Printf("\ninitial seeding...\n\n")

	for _, seeder := range seeders {
		err := handleDBFunction(db, dbConn, seeder.seedFunc, seeder.name)
		if err != nil {
			return h.NewErr(seeder.name, err)
		}
	}

	seedingDuration := time.Since(seedingStart)

	relationshipsStart := time.Now()
	fmt.Printf("\nseeding relationships...\n\n")

	for _, seeder := range seeders {
		err := handleDBFunction(db, dbConn, seeder.relFunc, seeder.name)
		if err != nil {
			return h.NewErr(seeder.name, err)
		}
	}

	relationshipsDuration := time.Since(relationshipsStart)

	fmt.Printf("\n\ninitial seeding took %.3f seconds\n", seedingDuration.Seconds())

	fmt.Printf("seeding relationships took %.3f seconds\n", relationshipsDuration.Seconds())

	return nil
}

func handleDBFunction(db *database.Queries, dbConn *sql.DB, function func(*database.Queries, *sql.DB) error, name string) error {
	if function == nil {
		return nil
	}

	start := time.Now()

	err := function(db, dbConn)
	if err != nil {
		return err
	}

	duration := time.Since(start)
	fmt.Printf("%.3f seconds for %s\n", duration.Seconds(), name)

	return nil
}

func setupDB(dbConn *sql.DB, migrationsDir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return err
	}

	err = goose.DownTo(dbConn, migrationsDir, 0)
	if err != nil {
		return err
	}

	err = goose.Up(dbConn, migrationsDir)
	if err != nil {
		return err
	}

	return nil
}

func (l *Lookup) getSeeders() []seeder {
	return []seeder{
		{
			name:     "stats",
			seedFunc: l.seedStats,
			relFunc:  l.seedStatsRelationships,
		},
		{
			name:     "elements",
			seedFunc: l.seedElements,
			relFunc:  l.seedElementsRelationships,
		},
		{
			name:     "affinities",
			seedFunc: l.seedAffinities,
			relFunc:  nil,
		},
		{
			name:     "agility tiers",
			seedFunc: l.seedAgilityTiers,
			relFunc:  nil,
		},
		{
			name:     "overdrive modes",
			seedFunc: l.seedOverdriveModes,
			relFunc:  l.seedOverdriveModesRelationships,
		},
		{
			name:     "status conditions",
			seedFunc: l.seedStatusConditions,
			relFunc:  l.seedStatusConditionsRelationships,
		},
		{
			name:     "properties",
			seedFunc: l.seedProperties,
			relFunc:  l.seedPropertiesRelationships,
		},
		{
			name:     "modifiers",
			seedFunc: l.seedModifiers,
			relFunc:  nil,
		},
		{
			name:     "characters",
			seedFunc: l.seedCharacters,
			relFunc:  l.seedCharactersRelationships,
		},
		{
			name:     "aeons",
			seedFunc: l.seedAeons,
			relFunc:  l.seedAeonsRelationships,
		},
		{
			name:     "aeon stats",
			seedFunc: nil,
			relFunc:  l.seedAeonStats,
		},
		{
			name:     "default abilities",
			seedFunc: nil,
			relFunc:  l.seedDefaultAbilitiesRelationships,
		},
		{
			name:     "blitzball items",
			seedFunc: l.seedBlitzballItems,
			relFunc:  l.seedBlitzballItemsRelationships,
		},
		{
			name:     "sidequests",
			seedFunc: l.seedSidequests,
			relFunc:  l.seedSidequestsRelationships,
		},
		{
			name:     "monster arena creations",
			seedFunc: l.seedArenaCreations,
			relFunc:  l.seedArenaCreationsRelationships,
		},
		{
			name:     "monsters",
			seedFunc: l.seedMonsters,
			relFunc:  l.seedMonstersRelationships,
		},
		{
			name:     "aeon commands",
			seedFunc: l.seedAeonCommands,
			relFunc:  l.seedAeonCommandsRelationships,
		},
		{
			name:     "submenus",
			seedFunc: l.seedSubmenus,
			relFunc:  l.seedSubmenusRelationships,
		},
		{
			name:     "player abilities",
			seedFunc: l.seedPlayerAbilities,
			relFunc:  l.seedPlayerAbilitiesRelationships,
		},
		{
			name:     "enemy abilities",
			seedFunc: l.seedEnemyAbilities,
			relFunc:  l.seedEnemyAbilitiesRelationships,
		},
		{
			name:     "overdrive abilities",
			seedFunc: l.seedOverdriveAbilities,
			relFunc:  l.seedOverdriveAbilitiesRelationships,
		},
		{
			name:     "trigger commands",
			seedFunc: l.seedTriggerCommands,
			relFunc:  l.seedTriggerCommandsRelationships,
		},
		{
			name:     "overdrive commands",
			seedFunc: l.seedOverdriveCommands,
			relFunc:  l.seedOverdriveCommandsRelationships,
		},
		{
			name:     "overdrives",
			seedFunc: l.seedOverdrives,
			relFunc:  l.seedOverdrivesRelationships,
		},
		{
			name:     "items",
			seedFunc: l.seedItems,
			relFunc:  l.seedItemsRelationships,
		},
		{
			name:     "key items",
			seedFunc: l.seedKeyItems,
			relFunc:  nil,
		},
		{
			name:     "primers",
			seedFunc: l.seedPrimers,
			relFunc:  nil,
		},
		{
			name:     "mixes",
			seedFunc: l.seedMixes,
			relFunc:  l.seedMixesRelationships,
		},
		{
			name:     "celestial weapons",
			seedFunc: l.seedCelestialWeapons,
			relFunc:  l.seedCelestialWeaponsRelationships,
		},
		{
			name:     "auto abilities",
			seedFunc: l.seedAutoAbilities,
			relFunc:  l.seedAutoAbilitiesRelationships,
		},
		{
			name:     "equipment",
			seedFunc: l.seedEquipment,
			relFunc:  l.seedEquipmentRelationships,
		},
		{
			name:     "location areas",
			seedFunc: l.seedLocations,
			relFunc:  l.seedAreasRelationships,
		},
		{
			name:     "treasures",
			seedFunc: l.seedTreasures,
			relFunc:  l.seedTreasuresRelationships,
		},
		{
			name:     "shops",
			seedFunc: l.seedShops,
			relFunc:  l.seedShopsRelationships,
		},
		{
			name:     "songs",
			seedFunc: l.seedSongs,
			relFunc:  l.seedSongsRelationships,
		},
		{
			name:     "fmvs",
			seedFunc: l.seedFMVs,
			relFunc:  nil,
		},
		{
			name:     "monster formations",
			seedFunc: l.seedMonsterFormations,
			relFunc:  l.seedMonsterFormationsRelationships,
		},
	}
}
