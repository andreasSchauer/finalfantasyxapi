package seeding

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/pressly/goose/v3"
)

type seeder struct {
	name		string
	seedFunc	func(*database.Queries, *sql.DB) error
	relFunc		func(*database.Queries, *sql.DB) error
}


func SeedDatabase(db *database.Queries, dbConn *sql.DB) error {
	const migrationsDir = "./sql/schema/"

	start := time.Now()

	err := setupDB(dbConn, migrationsDir)
	if err != nil {
		return fmt.Errorf("couldn't setup database: %v", err)
	}

	l := lookupInit()
	seeders := l.getSeeders()
	
	err = handleDBFunctions(db, dbConn, seeders)
	if err != nil {
		return err
	}

	totalDuration := time.Since(start)
	fmt.Printf("database seeding took %.3f seconds\n", totalDuration.Seconds())

	return nil
}


func handleDBFunctions(db *database.Queries, dbConn *sql.DB, seeders []seeder) error {
	seedingStart := time.Now()
	fmt.Println()
	fmt.Printf("initial seeding...\n\n")
	
	for _, seeder := range seeders {
		err := handleDBFunction(db, dbConn, seeder.seedFunc, seeder.name)
		if err != nil {
			return err
		}
	}

	seedingDuration := time.Since(seedingStart)
	fmt.Printf("\ninitial seeding took %.3f seconds\n\n", seedingDuration.Seconds())

	relationshipsStart := time.Now()
	fmt.Printf("creating relationships...\n\n")

	for _, seeder := range seeders {
		err := handleDBFunction(db, dbConn, seeder.relFunc, seeder.name)
		if err != nil {
			return err
		}
	}

	relationshipsDuration := time.Since(relationshipsStart)
	fmt.Printf("\ncreating relationships took %.3f seconds\n\n", relationshipsDuration.Seconds())

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


func (l *lookup) getSeeders() []seeder {
	return []seeder{
		{
			name: 		"stats",
			seedFunc: 	l.seedStats,
			relFunc: 	l.createStatsRelationships,
		},
		{
			name: 		"elements",
			seedFunc: 	l.seedElements,
			relFunc: 	l.createElementsRelationships,
		},
		{
			name: 		"affinities",
			seedFunc: 	l.seedAffinities,
			relFunc: 	nil,
		},
		{
			name: 		"agility tiers",
			seedFunc: 	l.seedAgilityTiers,
			relFunc: 	nil,
		},
		{
			name: 		"overdrive modes",
			seedFunc: 	l.seedOverdriveModes,
			relFunc: 	l.createOverdriveModesRelationships,
		},
		{
			name: 		"status conditions",
			seedFunc: 	l.seedStatusConditions,
			relFunc: 	l.createStatusConditionsRelationships,
		},
		{
			name: 		"properties",
			seedFunc: 	l.seedProperties,
			relFunc: 	l.createPropertiesRelationships,
		},
		{
			name: 		"modifiers",
			seedFunc: 	l.seedModifiers,
			relFunc: 	nil,
		},
		{
			name: 		"characters",
			seedFunc: 	l.seedCharacters,
			relFunc: 	l.createCharactersRelationships,
		},
		{
			name: 		"aeons",
			seedFunc: 	l.seedAeons,
			relFunc: 	l.createAeonsRelationships,
		},
		{
			name: 		"default abilities",
			seedFunc: 	nil,
			relFunc: 	l.createDefaultAbilitiesRelationships,
		},
		{
			name: 		"blitzball items",
			seedFunc: 	l.seedBlitzballItems,
			relFunc: 	l.createBlitzballItemsRelationships,
		},
		{
			name: 		"sidequests",
			seedFunc: 	l.seedSidequests,
			relFunc: 	l.createSidequestsRelationships,
		},
		{
			name: 		"monster arena creations",
			seedFunc: 	l.seedMonsterArenaCreations,
			relFunc: 	nil,
		},
		{
			name: 		"monsters",
			seedFunc: 	l.seedMonsters,
			relFunc: 	nil,
		},
		{
			name: 		"aeon commands",
			seedFunc: 	l.seedAeonCommands,
			relFunc: 	l.createAeonCommandsRelationships,
		},
		{
			name: 		"submenus",
			seedFunc: 	l.seedSubmenus,
			relFunc: 	l.createSubmenusRelationships,
		},
		{
			name: 		"player abilities",
			seedFunc: 	l.seedPlayerAbilities,
			relFunc: 	l.createPlayerAbilitiesRelationships,
		},
		{
			name: 		"enemy abilities",
			seedFunc: 	l.seedEnemyAbilities,
			relFunc: 	l.createEnemyAbilitiesRelationships,
		},
		{
			name: 		"overdrive abilities",
			seedFunc: 	l.seedOverdriveAbilities,
			relFunc: 	l.createOverdriveAbilitiesRelationships,
		},
		{
			name: 		"trigger commands",
			seedFunc: 	l.seedTriggerCommands,
			relFunc: 	l.createTriggerCommandsRelationships,
		},
		{
			name: 		"overdrive commands",
			seedFunc: 	l.seedOverdriveCommands,
			relFunc: 	l.createOverdriveCommandsRelationships,
		},
		{
			name: 		"overdrives",
			seedFunc: 	l.seedOverdrives,
			relFunc: 	l.createOverdrivesRelationships,
		},
		{
			name: 		"items",
			seedFunc: 	l.seedItems,
			relFunc: 	l.createItemsRelationships,
		},
		{
			name: 		"key items",
			seedFunc: 	l.seedKeyItems,
			relFunc: 	nil,
		},
		{
			name: 		"primers",
			seedFunc: 	l.seedPrimers,
			relFunc: 	nil,
		},
		{
			name: 		"mixes",
			seedFunc: 	l.seedMixes,
			relFunc: 	l.createMixesRelationships,
		},
		{
			name: 		"celestial weapons",
			seedFunc: 	l.seedCelestialWeapons,
			relFunc: 	l.createCelestialWeaponsRelationships,
		},
		{
			name: 		"auto abilities",
			seedFunc: 	l.seedAutoAbilities,
			relFunc: 	l.createAutoAbilitiesRelationships,
		},
		{
			name: 		"equipment",
			seedFunc: 	l.seedEquipment,
			relFunc: 	l.createEquipmentRelationships,
		},
		{
			name: 		"location areas",
			seedFunc: 	l.seedLocations,
			relFunc: 	l.createAreasRelationships,
		},
		{
			name: 		"treasures",
			seedFunc: 	l.seedTreasures,
			relFunc: 	l.createTreasuresRelationships,
		},
		{
			name: 		"shops",
			seedFunc: 	l.seedShops,
			relFunc: 	l.createShopsRelationships,
		},
		{
			name: 		"monster formations",
			seedFunc: 	l.seedFormationLocations,
			relFunc: 	l.createMonsterFormationsRelationships,
		},
		{
			name: 		"songs",
			seedFunc: 	l.seedSongs,
			relFunc: 	l.createSongsRelationships,
		},
		{
			name: 		"fmvs",
			seedFunc: 	l.seedFMVs,
			relFunc: 	nil,
		},
	}
}