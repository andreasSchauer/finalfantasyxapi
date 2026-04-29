package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
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
		l.loop1SeedAccuracies,
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
		l.loop2SeedSublocations,
		l.loop2SeedSongs,
		l.loop2SeedDamages,
		l.loop2SeedModifierChanges,
		l.loop2SeedMonsterAmounts,
		l.loop2SeedMonsterEquipments,
		l.loop2SeedAlteredStates,
	})
}

func (l *Lookup) seedLoop3(qtx *database.Queries, ctx context.Context) error {
	return l.seedLoop(qtx, ctx, []func(*database.Queries, context.Context) error{
		l.loop3SeedStatusConditions,
		l.loop3SeedQuestCompletions,
		l.loop3SeedAeonCommands,
		l.loop3SeedOverdriveCommands,
		l.loop3SeedUnspecifiedAbilities,
		l.loop3SeedEnemyAbilities,
		l.loop3SeedOverdriveAbilities,
		l.loop3SeedTriggerCommands,
		l.loop3SeedItemAbilities,
		l.loop3SeedSpheres,
		l.loop3SeedPrimers,
		l.loop3SeedAreas,
		l.loop3SeedPossibleItems,
		l.loop3SeedShopItems,
		l.loop3SeedBattleInteractions,
		l.loop3SeedFormationBossSongs,
		l.loop3SeedMonsterItems,
		l.loop3SeedMonsterAbilities,
		l.loop3SeedMonsterEquipmentSlots,
	})
}

func (l *Lookup) seedLoop4(qtx *database.Queries, ctx context.Context) error {
	return l.seedLoop(qtx, ctx, []func(*database.Queries, context.Context) error{
		l.loop4SeedStats,
		l.loop4SeedStatusResists,
		l.loop4SeedCharacters,
		l.loop4SeedAeons,
		l.loop4SeedQuests,
		l.loop4SeedBlitzballItems,
		l.loop4SeedCompletionAreas,
		l.loop4SeedOverdrives,
		l.loop4SeedTargetableNodes,
		l.loop4SeedTreasures,
		l.loop4SeedShops,
		l.loop4SeedAreaConnections,
		l.loop4SeedFMVs,
		l.loop4SeedCues,
		l.loop4SeedInflictedStatusses,
		l.loop4SeedFormationData,
		l.loop4SeedEncounterAreas,
		l.loop4SeedFormationTriggerCommands,
	})
}