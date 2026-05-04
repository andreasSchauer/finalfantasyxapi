package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type seedFunc func(*database.Queries, context.Context) error

func (l *Lookup) seedLoop(qtx *database.Queries, ctx context.Context, fns []seedFunc) error {
	for _, fn := range fns {
		err := fn(qtx, ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *Lookup) seedLoop1(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 1")()
	return l.seedLoop(qtx, ctx, []seedFunc{
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
	defer h.MeasureTime("- seed loop 2")()
	return l.seedLoop(qtx, ctx, []seedFunc{
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
	defer h.MeasureTime("- seed loop 3")()
	return l.seedLoop(qtx, ctx, []seedFunc{
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
	defer h.MeasureTime("- seed loop 4")()
	return l.seedLoop(qtx, ctx, []seedFunc{
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

func (l *Lookup) seedLoop5(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 5")()
	return l.seedLoop(qtx, ctx, []seedFunc{
		l.loop5SeedBaseStats,
		l.loop5SeedOdModeActions,
		l.loop5SeedSidequests,
		l.loop5SeedPlayerAbilities,
		l.loop5SeedRonsoRages,
		l.loop5SeedMixes,
		l.loop5SeedCelestialWeapons,
		l.loop5SeedAutoAbilities,
		l.loop5SeedEquipmentTables,
		l.loop5SeedEquipmentNames,
		l.loop5SeedMonsterFormations,
		l.loop5SeedAbilityDamages,
		l.loop5SeedStatChanges,
		l.loop5SeedAltStateChanges,
	})
}

func (l *Lookup) seedLoop6(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 6")()
	return l.seedLoop(qtx, ctx, []seedFunc{
		l.loop6SeedAeonEquipment,
		l.loop6SeedSubquests,
		l.loop6SeedMixCombinations,
		l.loop6SeedAbilityPools,
		l.loop6SeedTreasureEquipment,
		l.loop6SeedShopEquipment,
		l.loop6SeedEquipmentDrops,
	})
}

func (l *Lookup) seedLoop7(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 7")()
	return l.seedLoop(qtx, ctx, []seedFunc{
		l.loop7SeedArenaCreations,
	})
}

func (l *Lookup) seedJunctions(qtx *database.Queries, ctx context.Context) error {
	return l.seedLoop(qtx, ctx, []seedFunc{
		l.seedJuncOverdriveModeActions,
		l.seedJuncPropertiesModifierChanges,
		l.seedJuncPropertiesRelatedStats,
		l.seedJuncStatusConditionModifierChanges,
		l.seedJuncStatusConditionRelatedStats,
		l.seedJuncStatusConditionRemovedConditions,
		l.seedJuncStatusConditionStatChanges,
		l.seedJuncAeonBaseStatsA,
		l.seedJuncAeonBaseStatsB,
		l.seedJuncAeonAeonEquipment,
		l.seedJuncCharacterClassesPlayerUnits,
		l.seedJuncCharactersBaseStats,
		l.seedJuncSubmenusUsers,
		l.seedJuncAbilitiesBattleInteractions,
		l.seedJuncOverdriveAbilitiesRelatedStats,
		l.seedJuncOverdrivesOverdriveAbilities,
		l.seedJuncPlayerAbilitiesLearnedBy,
		l.seedJuncPlayerAbilitiesRelatedStats,
		l.seedJuncTriggerCommandsRelatedStats,
		l.seedJuncUnspecifiedAbilitiesLearnedBy,
		l.seedJuncItemsAvailableMenus,
		l.seedJuncItemsRelatedStats,
		l.seedJuncAbilityPoolsAutoAbilities,
		l.seedJuncAutoAbilitiesAddedStatusResists,
		l.seedJuncAutoAbilitiesAddedStatusses,
		l.seedJuncAutoAbilitiesAutoItems,
		l.seedJuncAutoAbilitiesLockedOutAbilities,
		l.seedJuncAutoAbilitiesModifierChanges,
		l.seedJuncAutoAbilitiesRelatedStats,
		l.seedJuncAutoAbilitiesStatChanges,
		l.seedJuncAreaConnectedAreas,
		l.seedJuncShopEquipmentAutoAbilities,
		l.seedJuncTreasureEquipmentAutoAbilities,
		l.seedJuncTreasuresItems,
		l.seedJuncCuesIncludedAreas,
		l.seedJuncBattleInteractionsAffectedBy,
		l.seedJuncBattleInteractionsCopiedStatusConditions,
		l.seedJuncBattleInteractionsDamages,
		l.seedJuncBattleInteractionsInflictedStatusConditions,
		l.seedJuncBattleInteractionsModifierChanges,
		l.seedJuncBattleInteractionsRemovedStatusConditions,
		l.seedJuncBattleInteractionsStatChanges,
		l.seedJuncAltStateChangesAutoAbilities,
		l.seedJuncAltStateChangesBaseStats,
		l.seedJuncAltStateChangesElementalResists,
		l.seedJuncAltStateChangesProperties,
		l.seedJuncAltStateChangesStatusImmunities,
		l.seedJuncMonsterItemsOtherItems,
		l.seedJuncMonsterEquipmentEquipmentDrops,
		l.seedJuncFormationTriggerCommandsUsers,
		l.seedJuncMonsterFormationsEncounterAreas,
		l.seedJuncMonsterFormationsTriggerCommands,
		l.seedJuncMonsterSelectionMonsterAmounts,
		l.seedJuncMonstersMonsterAbilities,
		l.seedJuncMonstersAutoAbilities,
		l.seedJuncMonstersBaseStats,
		l.seedJuncMonstersElementalResists,
		l.seedJuncMonstersStatusImmunities,
		l.seedJuncMonstersProperties,
		l.seedJuncMonstersRonsoRages,
		l.seedJuncMonstersStatusResists,
	})
}