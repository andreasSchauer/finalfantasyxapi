package seeding

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/junctions"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhandlers"
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
		seedhandlers.Loop1SeedAgilityTiers,
		seedhandlers.Loop1SeedElements,
		seedhandlers.Loop1SeedOverdriveModes,
		seedhandlers.Loop1SeedModifiers,
		seedhandlers.Loop1SeedPlayerUnits,
		seedhandlers.Loop1SeedCharacterClasses,
		seedhandlers.Loop1SeedBlitzballPositions,
		seedhandlers.Loop1SeedTopmenus,
		seedhandlers.Loop1SeedAbilityAttributes,
		seedhandlers.Loop1SeedMasterItems,
		seedhandlers.Loop1SeedCreatedNodes,
		seedhandlers.Loop1SeedLocations,
		seedhandlers.Loop1SeedBackgroundMusic,
		seedhandlers.Loop1SeedSongCredits,
		seedhandlers.Loop1SeedAccuracies,
		seedhandlers.Loop1SeedInflictedDelays,
		seedhandlers.Loop1SeedMonsters,
		seedhandlers.Loop1SeedMonsterSelections,
		seedhandlers.Loop1SeedEquipmentSlotsChances,
		seedhandlers.Loop1SeedQuests,
		seedhandlers.Loop1SeedSphereGrids,
		seedhandlers.Loop1SeedClassDamages,
	})
}

func seedLoop2(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 2")()
	return seedLoop(qtx, ctx, []seedFunc{
		seedhandlers.Loop2SeedAgilitySubtiers,
		seedhandlers.Loop2UpdateElements,
		seedhandlers.Loop2SeedElementalResists,
		seedhandlers.Loop2SeedSubmenus,
		seedhandlers.Loop2SeedAbilities,
		seedhandlers.Loop2SeedItems,
		seedhandlers.Loop2SeedKeyItems,
		seedhandlers.Loop2SeedItemAmounts,
		seedhandlers.Loop2SeedSublocations,
		seedhandlers.Loop2SeedSongs,
		seedhandlers.Loop2SeedDamages,
		seedhandlers.Loop2SeedModifierChanges,
		seedhandlers.Loop2SeedMonsterAmounts,
		seedhandlers.Loop2SeedMonsterEquipment,
		seedhandlers.Loop2SeedAlteredStates,
	})
}

func seedLoop3(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 3")()
	return seedLoop(qtx, ctx, []seedFunc{
		seedhandlers.Loop3SeedProperties,
		seedhandlers.Loop3SeedStatusConditions,
		seedhandlers.Loop3SeedQuestCompletions,
		seedhandlers.Loop3SeedAeonCommands,
		seedhandlers.Loop3SeedOverdriveCommands,
		seedhandlers.Loop3SeedMiscAbilities,
		seedhandlers.Loop3SeedEnemyAbilities,
		seedhandlers.Loop3SeedOverdriveAbilities,
		seedhandlers.Loop3SeedTriggerCommands,
		seedhandlers.Loop3SeedItemAbilities,
		seedhandlers.Loop3SeedSpheres,
		seedhandlers.Loop3SeedPrimers,
		seedhandlers.Loop3SeedAreas,
		seedhandlers.Loop3SeedPossibleItems,
		seedhandlers.Loop3SeedShopItems,
		seedhandlers.Loop3SeedBattleInteractions,
		seedhandlers.Loop3SeedFormationBossSongs,
		seedhandlers.Loop3SeedMonsterItems,
		seedhandlers.Loop3SeedMonsterAbilities,
		seedhandlers.Loop3SeedMonsterEquipmentSlots,
	})
}

func seedLoop4(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 4")()
	return seedLoop(qtx, ctx, []seedFunc{
		seedhandlers.Loop4SeedStats,
		seedhandlers.Loop4SeedStatusResists,
		seedhandlers.Loop4SeedCharacters,
		seedhandlers.Loop4SeedAeons,
		seedhandlers.Loop4SeedBlitzballItems,
		seedhandlers.Loop4SeedCompletionAreas,
		seedhandlers.Loop4SeedOverdrives,
		seedhandlers.Loop4SeedTargetableNodes,
		seedhandlers.Loop4SeedTreasures,
		seedhandlers.Loop4SeedShops,
		seedhandlers.Loop4SeedAreaConnections,
		seedhandlers.Loop4SeedFMVs,
		seedhandlers.Loop4SeedCues,
		seedhandlers.Loop4SeedInflictedStatusses,
		seedhandlers.Loop4SeedFormationData,
		seedhandlers.Loop4SeedEncounterAreas,
		seedhandlers.Loop4SeedFormationTriggerCommands,
	})
}

func seedLoop5(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 5")()
	return seedLoop(qtx, ctx, []seedFunc{
		seedhandlers.Loop5SeedBaseStats,
		seedhandlers.Loop5SeedOdModeActions,
		seedhandlers.Loop5SeedSidequests,
		seedhandlers.Loop5SeedPlayerAbilities,
		seedhandlers.Loop5SeedRonsoRages,
		seedhandlers.Loop5SeedMixes,
		seedhandlers.Loop5SeedCelestialWeapons,
		seedhandlers.Loop5SeedAutoAbilities,
		seedhandlers.Loop5SeedEquipmentTables,
		seedhandlers.Loop5SeedEquipmentNames,
		seedhandlers.Loop5SeedMonsterFormations,
		seedhandlers.Loop5SeedStatChanges,
		seedhandlers.Loop5SeedAlts,
	})
}

func seedLoop6(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 6")()
	return seedLoop(qtx, ctx, []seedFunc{
		seedhandlers.Loop6SeedAeonEquipment,
		seedhandlers.Loop6SeedSubquests,
		seedhandlers.Loop6SeedMixCombinations,
		seedhandlers.Loop6SeedAbilityPools,
		seedhandlers.Loop6SeedTreasureEquipment,
		seedhandlers.Loop6SeedShopEquipment,
		seedhandlers.Loop6SeedEquipmentDrops,
	})
}

func seedLoop7(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 7")()
	return seedLoop(qtx, ctx, []seedFunc{
		seedhandlers.Loop7SeedArenaCreations,
	})
}

func seedJunctions(qtx *database.Queries, ctx context.Context) error {
	return seedLoop(qtx, ctx, []seedFunc{
		junctions.SeedJuncOverdriveModeActions,
		junctions.SeedJuncPropertiesRelatedStats,
		junctions.SeedJuncStatusConditionModifierChanges,
		junctions.SeedJuncStatusConditionRelatedStats,
		junctions.SeedJuncStatusConditionRemovedConditions,
		junctions.SeedJuncStatusConditionStatChanges,
		junctions.SeedJuncAeonBaseStatsA,
		junctions.SeedJuncAeonBaseStatsB,
		junctions.SeedJuncAeonBaseStatsX,
		junctions.SeedJuncAeonAeonEquipment,
		junctions.SeedJuncCharacterClassesPlayerUnits,
		junctions.SeedJuncDefaultAbilities,
		junctions.SeedJuncDefaultOverdriveAbilities,
		junctions.SeedJuncCharactersBaseStats,
		junctions.SeedJuncAeonCommandsPossibleAbilities,
		junctions.SeedJuncSubmenusUsers,
		junctions.SeedJuncAbilitiesBattleInteractions,
		junctions.SeedJuncOverdriveAbilitiesRelatedStats,
		junctions.SeedJuncOverdrivesOverdriveAbilities,
		junctions.SeedJuncPlayerAbilitiesLearnedBy,
		junctions.SeedJuncPlayerAbilitiesRelatedStats,
		junctions.SeedJuncTriggerCommandsRelatedStats,
		junctions.SeedJuncMiscAbilitiesLearnedBy,
		junctions.SeedJuncItemsAvailableMenus,
		junctions.SeedJuncItemsRelatedStats,
		junctions.SeedJuncAbilityPoolsAutoAbilities,
		junctions.SeedJuncAutoAbilitiesAddedStatusResists,
		junctions.SeedJuncAutoAbilitiesAddedStatusses,
		junctions.SeedJuncAutoAbilitiesAutoItems,
		junctions.SeedJuncAutoAbilitiesLockedOutAbilities,
		junctions.SeedJuncAutoAbilitiesModifierChanges,
		junctions.SeedJuncAutoAbilitiesRelatedStats,
		junctions.SeedJuncAutoAbilitiesStatChanges,
		junctions.SeedJuncEquipmentTablesNames,
		junctions.SeedJuncEquipmentReqAutoAbilities,
		junctions.SeedJuncAreaConnectedAreas,
		junctions.SeedJuncShopEquipmentAutoAbilities,
		junctions.SeedJuncShopsShopItems,
		junctions.SeedJuncTreasureEquipmentAutoAbilities,
		junctions.SeedJuncTreasuresItems,
		junctions.SeedJuncSongsBackgroundMusic,
		junctions.SeedJuncCuesIncludedAreas,
		junctions.SeedJuncBattleInteractionsAffectedBy,
		junctions.SeedJuncBattleInteractionsCopiedStatusConditions,
		junctions.SeedJuncBattleInteractionsDamages,
		junctions.SeedJuncBattleInteractionsInflictedStatusConditions,
		junctions.SeedJuncBattleInteractionsModifierChanges,
		junctions.SeedJuncBattleInteractionsRemovedStatusConditions,
		junctions.SeedJuncBattleInteractionsStatChanges,
		junctions.SeedJuncAltsAutoAbilities,
		junctions.SeedJuncAltsBaseStats,
		junctions.SeedJuncAltsElementalResists,
		junctions.SeedJuncAltsProperties,
		junctions.SeedJuncAltsStatusImmunities,
		junctions.SeedJuncMonsterItemsOtherItems,
		junctions.SeedJuncMonsterEquipmentEquipmentDrops,
		junctions.SeedJuncMonsterEquipmentSlotChances,
		junctions.SeedJuncEquipmentDropsCharacters,
		junctions.SeedJuncFormationTriggerCommandsUsers,
		junctions.SeedJuncMonsterFormationsEncounterAreas,
		junctions.SeedJuncMonsterFormationsTriggerCommands,
		junctions.SeedJuncMonsterSelectionMonsterAmounts,
		junctions.SeedJuncMonstersMonsterAbilities,
		junctions.SeedJuncMonstersAutoAbilities,
		junctions.SeedJuncMonstersBaseStats,
		junctions.SeedJuncMonstersElementalResists,
		junctions.SeedJuncMonstersStatusImmunities,
		junctions.SeedJuncMonstersProperties,
		junctions.SeedJuncMonstersRonsoRages,
		junctions.SeedJuncMonstersStatusResists,
	})
}
