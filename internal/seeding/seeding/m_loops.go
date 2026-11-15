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
		Loop1SeedAgilityTiers,
		Loop1SeedElements,
		Loop1SeedOverdriveModes,
		Loop1SeedModifiers,
		Loop1SeedPlayerUnits,
		Loop1SeedCharacterClasses,
		Loop1SeedBlitzballPositions,
		Loop1SeedTopmenus,
		Loop1SeedAbilityAttributes,
		Loop1SeedMasterItems,
		Loop1SeedCreatedNodes,
		Loop1SeedLocations,
		Loop1SeedBackgroundMusic,
		Loop1SeedSongCredits,
		Loop1SeedAccuracies,
		Loop1SeedInflictedDelays,
		Loop1SeedMonsters,
		Loop1SeedMonsterSelections,
		Loop1SeedEquipmentSlotsChances,
		Loop1SeedQuests,
		Loop1SeedSphereGrids,
		Loop1SeedClassDamages,
	})
}

func seedLoop2(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 2")()
	return seedLoop(qtx, ctx, []seedFunc{
		Loop2SeedAgilitySubtiers,
		loop2UpdateElements,
		Loop2SeedElementalResists,
		Loop2SeedSubmenus,
		Loop2SeedAbilities,
		Loop2SeedItems,
		Loop2SeedKeyItems,
		Loop2SeedItemAmounts,
		Loop2SeedSublocations,
		Loop2SeedSongs,
		Loop2SeedDamages,
		Loop2SeedModifierChanges,
		Loop2SeedMonsterAmounts,
		Loop2SeedMonsterEquipment,
		Loop2SeedAlteredStates,
	})
}

func seedLoop3(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 3")()
	return seedLoop(qtx, ctx, []seedFunc{
		Loop3SeedProperties,
		Loop3SeedStatusConditions,
		Loop3SeedQuestCompletions,
		Loop3SeedAeonCommands,
		Loop3SeedOverdriveCommands,
		Loop3SeedMiscAbilities,
		Loop3SeedEnemyAbilities,
		Loop3SeedOverdriveAbilities,
		Loop3SeedTriggerCommands,
		Loop3SeedItemAbilities,
		Loop3SeedSpheres,
		Loop3SeedPrimers,
		Loop3SeedAreas,
		Loop3SeedPossibleItems,
		Loop3SeedShopItems,
		Loop3SeedBattleInteractions,
		Loop3SeedFormationBossSongs,
		Loop3SeedMonsterItems,
		Loop3SeedMonsterAbilities,
		Loop3SeedMonsterEquipmentSlots,
	})
}

func seedLoop4(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 4")()
	return seedLoop(qtx, ctx, []seedFunc{
		Loop4SeedStats,
		Loop4SeedStatusResists,
		Loop4SeedCharacters,
		Loop4SeedAeons,
		Loop4SeedBlitzballItems,
		Loop4SeedCompletionAreas,
		Loop4SeedOverdrives,
		Loop4SeedTargetableNodes,
		Loop4SeedTreasures,
		Loop4SeedShops,
		Loop4SeedAreaConnections,
		Loop4SeedFMVs,
		Loop4SeedCues,
		Loop4SeedInflictedStatusses,
		Loop4SeedFormationData,
		Loop4SeedEncounterAreas,
		Loop4SeedFormationTriggerCommands,
	})
}

func seedLoop5(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 5")()
	return seedLoop(qtx, ctx, []seedFunc{
		Loop5SeedBaseStats,
		Loop5SeedOdModeActions,
		Loop5SeedSidequests,
		Loop5SeedPlayerAbilities,
		Loop5SeedRonsoRages,
		Loop5SeedMixes,
		Loop5SeedCelestialWeapons,
		Loop5SeedAutoAbilities,
		Loop5SeedEquipmentTables,
		Loop5SeedEquipmentNames,
		Loop5SeedMonsterFormations,
		Loop5SeedStatChanges,
		Loop5SeedAlts,
	})
}

func seedLoop6(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 6")()
	return seedLoop(qtx, ctx, []seedFunc{
		Loop6SeedAeonEquipment,
		Loop6SeedSubquests,
		Loop6SeedMixCombinations,
		Loop6SeedAbilityPools,
		Loop6SeedTreasureEquipment,
		Loop6SeedShopEquipment,
		Loop6SeedEquipmentDrops,
	})
}

func seedLoop7(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 7")()
	return seedLoop(qtx, ctx, []seedFunc{
		Loop7SeedArenaCreations,
	})
}

func seedJunctions(qtx *database.Queries, ctx context.Context) error {
	return seedLoop(qtx, ctx, []seedFunc{
		SeedJuncOverdriveModeActions,
		SeedJuncPropertiesRelatedStats,
		SeedJuncStatusConditionModifierChanges,
		SeedJuncStatusConditionRelatedStats,
		SeedJuncStatusConditionRemovedConditions,
		SeedJuncStatusConditionStatChanges,
		SeedJuncAeonBaseStatsA,
		SeedJuncAeonBaseStatsB,
		SeedJuncAeonBaseStatsX,
		SeedJuncAeonAeonEquipment,
		SeedJuncCharacterClassesPlayerUnits,
		SeedJuncDefaultAbilities,
		SeedJuncDefaultOverdriveAbilities,
		SeedJuncCharactersBaseStats,
		SeedJuncAeonCommandsPossibleAbilities,
		SeedJuncSubmenusUsers,
		SeedJuncAbilitiesBattleInteractions,
		SeedJuncOverdriveAbilitiesRelatedStats,
		SeedJuncOverdrivesOverdriveAbilities,
		SeedJuncPlayerAbilitiesLearnedBy,
		SeedJuncPlayerAbilitiesRelatedStats,
		SeedJuncTriggerCommandsRelatedStats,
		SeedJuncMiscAbilitiesLearnedBy,
		SeedJuncItemsAvailableMenus,
		SeedJuncItemsRelatedStats,
		SeedJuncAbilityPoolsAutoAbilities,
		SeedJuncAutoAbilitiesAddedStatusResists,
		SeedJuncAutoAbilitiesAddedStatusses,
		SeedJuncAutoAbilitiesAutoItems,
		SeedJuncAutoAbilitiesLockedOutAbilities,
		SeedJuncAutoAbilitiesModifierChanges,
		SeedJuncAutoAbilitiesRelatedStats,
		SeedJuncAutoAbilitiesStatChanges,
		SeedJuncEquipmentTablesNames,
		SeedJuncEquipmentReqAutoAbilities,
		SeedJuncAreaConnectedAreas,
		SeedJuncShopEquipmentAutoAbilities,
		SeedJuncShopsShopItems,
		SeedJuncTreasureEquipmentAutoAbilities,
		SeedJuncTreasuresItems,
		SeedJuncSongsBackgroundMusic,
		SeedJuncCuesIncludedAreas,
		SeedJuncBattleInteractionsAffectedBy,
		SeedJuncBattleInteractionsCopiedStatusConditions,
		SeedJuncBattleInteractionsDamages,
		SeedJuncBattleInteractionsInflictedStatusConditions,
		SeedJuncBattleInteractionsModifierChanges,
		SeedJuncBattleInteractionsRemovedStatusConditions,
		SeedJuncBattleInteractionsStatChanges,
		SeedJuncAltsAutoAbilities,
		SeedJuncAltsBaseStats,
		SeedJuncAltsElementalResists,
		SeedJuncAltsProperties,
		SeedJuncAltsStatusImmunities,
		SeedJuncMonsterItemsOtherItems,
		SeedJuncMonsterEquipmentEquipmentDrops,
		SeedJuncMonsterEquipmentSlotChances,
		SeedJuncEquipmentDropsCharacters,
		SeedJuncFormationTriggerCommandsUsers,
		SeedJuncMonsterFormationsEncounterAreas,
		SeedJuncMonsterFormationsTriggerCommands,
		SeedJuncMonsterSelectionMonsterAmounts,
		SeedJuncMonstersMonsterAbilities,
		SeedJuncMonstersAutoAbilities,
		SeedJuncMonstersBaseStats,
		SeedJuncMonstersElementalResists,
		SeedJuncMonstersStatusImmunities,
		SeedJuncMonstersProperties,
		SeedJuncMonstersRonsoRages,
		SeedJuncMonstersStatusResists,
	})
}
