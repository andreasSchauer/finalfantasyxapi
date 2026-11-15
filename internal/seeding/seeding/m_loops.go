package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

type seedFunc func(*database.Queries, context.Context) error

func seedLoop(qtx *database.Queries, ctx context.Context, fns []seedFunc) error {
	for _, fn := range fns {
		err := fn(qtx, ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func seedLoop1(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 1")()
	return seedLoop(qtx, ctx, []seedFunc{
		loop1SeedAgilityTiers,
		loop1SeedElements,
		loop1SeedOverdriveModes,
		loop1SeedModifiers,
		loop1SeedPlayerUnits,
		loop1SeedCharacterClasses,
		loop1SeedBlitzballPositions,
		loop1SeedTopmenus,
		loop1SeedAbilityAttributes,
		loop1SeedMasterItems,
		loop1SeedCreatedNodes,
		loop1SeedLocations,
		loop1SeedBackgroundMusic,
		loop1SeedSongCredits,
		loop1SeedAccuracies,
		loop1SeedInflictedDelays,
		loop1SeedMonsters,
		loop1SeedMonsterSelections,
		loop1SeedEquipmentSlotsChances,
		loop1SeedQuests,
		loop1SeedSphereGrids,
		loop1SeedClassDamages,
	})
}

func seedLoop2(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 2")()
	return seedLoop(qtx, ctx, []seedFunc{
		loop2SeedAgilitySubtiers,
		loop2UpdateElements,
		loop2SeedElementalResists,
		loop2SeedSubmenus,
		loop2SeedAbilities,
		loop2SeedItems,
		loop2SeedKeyItems,
		loop2SeedItemAmounts,
		loop2SeedSublocations,
		loop2SeedSongs,
		loop2SeedDamages,
		loop2SeedModifierChanges,
		loop2SeedMonsterAmounts,
		loop2SeedMonsterEquipment,
		loop2SeedAlteredStates,
	})
}

func seedLoop3(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 3")()
	return seedLoop(qtx, ctx, []seedFunc{
		loop3SeedProperties,
		loop3SeedStatusConditions,
		loop3SeedQuestCompletions,
		loop3SeedAeonCommands,
		loop3SeedOverdriveCommands,
		loop3SeedMiscAbilities,
		loop3SeedEnemyAbilities,
		loop3SeedOverdriveAbilities,
		loop3SeedTriggerCommands,
		loop3SeedItemAbilities,
		loop3SeedSpheres,
		loop3SeedPrimers,
		loop3SeedAreas,
		loop3SeedPossibleItems,
		loop3SeedShopItems,
		loop3SeedBattleInteractions,
		loop3SeedFormationBossSongs,
		loop3SeedMonsterItems,
		loop3SeedMonsterAbilities,
		loop3SeedMonsterEquipmentSlots,
	})
}

func seedLoop4(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 4")()
	return seedLoop(qtx, ctx, []seedFunc{
		loop4SeedStats,
		loop4SeedStatusResists,
		loop4SeedCharacters,
		loop4SeedAeons,
		loop4SeedBlitzballItems,
		loop4SeedCompletionAreas,
		loop4SeedOverdrives,
		loop4SeedTargetableNodes,
		loop4SeedTreasures,
		loop4SeedShops,
		loop4SeedAreaConnections,
		loop4SeedFMVs,
		loop4SeedCues,
		loop4SeedInflictedStatusses,
		loop4SeedFormationData,
		loop4SeedEncounterAreas,
		loop4SeedFormationTriggerCommands,
	})
}

func seedLoop5(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 5")()
	return seedLoop(qtx, ctx, []seedFunc{
		loop5SeedBaseStats,
		loop5SeedOdModeActions,
		loop5SeedSidequests,
		loop5SeedPlayerAbilities,
		loop5SeedRonsoRages,
		loop5SeedMixes,
		loop5SeedCelestialWeapons,
		loop5SeedAutoAbilities,
		loop5SeedEquipmentTables,
		loop5SeedEquipmentNames,
		loop5SeedMonsterFormations,
		loop5SeedStatChanges,
		loop5SeedAlts,
	})
}

func seedLoop6(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 6")()
	return seedLoop(qtx, ctx, []seedFunc{
		loop6SeedAeonEquipment,
		loop6SeedSubquests,
		loop6SeedMixCombinations,
		loop6SeedAbilityPools,
		loop6SeedTreasureEquipment,
		loop6SeedShopEquipment,
		loop6SeedEquipmentDrops,
	})
}

func seedLoop7(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 7")()
	return seedLoop(qtx, ctx, []seedFunc{
		loop7SeedArenaCreations,
	})
}

func seedJunctions(qtx *database.Queries, ctx context.Context) error {
	return seedLoop(qtx, ctx, []seedFunc{
		seedJuncOverdriveModeActions,
		seedJuncPropertiesRelatedStats,
		seedJuncStatusConditionModifierChanges,
		seedJuncStatusConditionRelatedStats,
		seedJuncStatusConditionRemovedConditions,
		seedJuncStatusConditionStatChanges,
		seedJuncAeonBaseStatsA,
		seedJuncAeonBaseStatsB,
		seedJuncAeonBaseStatsX,
		seedJuncAeonAeonEquipment,
		seedJuncCharacterClassesPlayerUnits,
		seedJuncDefaultAbilities,
		seedJuncDefaultOverdriveAbilities,
		seedJuncCharactersBaseStats,
		seedJuncAeonCommandsPossibleAbilities,
		seedJuncSubmenusUsers,
		seedJuncAbilitiesBattleInteractions,
		seedJuncOverdriveAbilitiesRelatedStats,
		seedJuncOverdrivesOverdriveAbilities,
		seedJuncPlayerAbilitiesLearnedBy,
		seedJuncPlayerAbilitiesRelatedStats,
		seedJuncTriggerCommandsRelatedStats,
		seedJuncMiscAbilitiesLearnedBy,
		seedJuncItemsAvailableMenus,
		seedJuncItemsRelatedStats,
		seedJuncAbilityPoolsAutoAbilities,
		seedJuncAutoAbilitiesAddedStatusResists,
		seedJuncAutoAbilitiesAddedStatusses,
		seedJuncAutoAbilitiesAutoItems,
		seedJuncAutoAbilitiesLockedOutAbilities,
		seedJuncAutoAbilitiesModifierChanges,
		seedJuncAutoAbilitiesRelatedStats,
		seedJuncAutoAbilitiesStatChanges,
		seedJuncEquipmentTablesNames,
		seedJuncEquipmentReqAutoAbilities,
		seedJuncAreaConnectedAreas,
		seedJuncShopEquipmentAutoAbilities,
		seedJuncShopsShopItems,
		seedJuncTreasureEquipmentAutoAbilities,
		seedJuncTreasuresItems,
		seedJuncSongsBackgroundMusic,
		seedJuncCuesIncludedAreas,
		seedJuncBattleInteractionsAffectedBy,
		seedJuncBattleInteractionsCopiedStatusConditions,
		seedJuncBattleInteractionsDamages,
		seedJuncBattleInteractionsInflictedStatusConditions,
		seedJuncBattleInteractionsModifierChanges,
		seedJuncBattleInteractionsRemovedStatusConditions,
		seedJuncBattleInteractionsStatChanges,
		seedJuncAltsAutoAbilities,
		seedJuncAltsBaseStats,
		seedJuncAltsElementalResists,
		seedJuncAltsProperties,
		seedJuncAltsStatusImmunities,
		seedJuncMonsterItemsOtherItems,
		seedJuncMonsterEquipmentEquipmentDrops,
		seedJuncMonsterEquipmentSlotChances,
		seedJuncEquipmentDropsCharacters,
		seedJuncFormationTriggerCommandsUsers,
		seedJuncMonsterFormationsEncounterAreas,
		seedJuncMonsterFormationsTriggerCommands,
		seedJuncMonsterSelectionMonsterAmounts,
		seedJuncMonstersMonsterAbilities,
		seedJuncMonstersAutoAbilities,
		seedJuncMonstersBaseStats,
		seedJuncMonstersElementalResists,
		seedJuncMonstersStatusImmunities,
		seedJuncMonstersProperties,
		seedJuncMonstersRonsoRages,
		seedJuncMonstersStatusResists,
	})
}
