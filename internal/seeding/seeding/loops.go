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
		seedhandlers.Loop1AgilityTiers,
		seedhandlers.Loop1Elements,
		seedhandlers.Loop1OverdriveModes,
		seedhandlers.Loop1Modifiers,
		seedhandlers.Loop1PlayerUnits,
		seedhandlers.Loop1CharacterClasses,
		seedhandlers.Loop1BlitzballPositions,
		seedhandlers.Loop1Topmenus,
		seedhandlers.Loop1AbilityAttributes,
		seedhandlers.Loop1MasterItems,
		seedhandlers.Loop1CreatedNodes,
		seedhandlers.Loop1Locations,
		seedhandlers.Loop1BackgroundMusic,
		seedhandlers.Loop1SongCredits,
		seedhandlers.Loop1Accuracies,
		seedhandlers.Loop1InflictedDelays,
		seedhandlers.Loop1Monsters,
		seedhandlers.Loop1MonsterSelections,
		seedhandlers.Loop1EquipmentSlotsChances,
		seedhandlers.Loop1Quests,
		seedhandlers.Loop1SphereGrids,
		seedhandlers.Loop1ClassDamages,
	})
}

func seedLoop2(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 2")()
	return seedLoop(qtx, ctx, []seedFunc{
		seedhandlers.Loop2AgilitySubtiers,
		seedhandlers.Loop2UpdateElements,
		seedhandlers.Loop2ElementalResists,
		seedhandlers.Loop2Submenus,
		seedhandlers.Loop2Abilities,
		seedhandlers.Loop2Items,
		seedhandlers.Loop2KeyItems,
		seedhandlers.Loop2ItemAmounts,
		seedhandlers.Loop2Sublocations,
		seedhandlers.Loop2Songs,
		seedhandlers.Loop2Damages,
		seedhandlers.Loop2ModifierChanges,
		seedhandlers.Loop2MonsterAmounts,
		seedhandlers.Loop2MonsterEquipment,
		seedhandlers.Loop2AlteredStates,
	})
}

func seedLoop3(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 3")()
	return seedLoop(qtx, ctx, []seedFunc{
		seedhandlers.Loop3Properties,
		seedhandlers.Loop3StatusConditions,
		seedhandlers.Loop3QuestCompletions,
		seedhandlers.Loop3AeonCommands,
		seedhandlers.Loop3OverdriveCommands,
		seedhandlers.Loop3MiscAbilities,
		seedhandlers.Loop3EnemyAbilities,
		seedhandlers.Loop3OverdriveAbilities,
		seedhandlers.Loop3TriggerCommands,
		seedhandlers.Loop3ItemAbilities,
		seedhandlers.Loop3Spheres,
		seedhandlers.Loop3Primers,
		seedhandlers.Loop3Areas,
		seedhandlers.Loop3PossibleItems,
		seedhandlers.Loop3ShopItems,
		seedhandlers.Loop3BattleInteractions,
		seedhandlers.Loop3FormationBossSongs,
		seedhandlers.Loop3MonsterItems,
		seedhandlers.Loop3MonsterAbilities,
		seedhandlers.Loop3MonsterEquipmentSlots,
	})
}

func seedLoop4(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 4")()
	return seedLoop(qtx, ctx, []seedFunc{
		seedhandlers.Loop4Stats,
		seedhandlers.Loop4StatusResists,
		seedhandlers.Loop4Characters,
		seedhandlers.Loop4Aeons,
		seedhandlers.Loop4BlitzballItems,
		seedhandlers.Loop4CompletionAreas,
		seedhandlers.Loop4Overdrives,
		seedhandlers.Loop4TargetableNodes,
		seedhandlers.Loop4Treasures,
		seedhandlers.Loop4Shops,
		seedhandlers.Loop4AreaConnections,
		seedhandlers.Loop4FMVs,
		seedhandlers.Loop4Cues,
		seedhandlers.Loop4InflictedStatusses,
		seedhandlers.Loop4FormationData,
		seedhandlers.Loop4EncounterAreas,
		seedhandlers.Loop4FormationTriggerCommands,
	})
}

func seedLoop5(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 5")()
	return seedLoop(qtx, ctx, []seedFunc{
		seedhandlers.Loop5BaseStats,
		seedhandlers.Loop5OdModeActions,
		seedhandlers.Loop5Sidequests,
		seedhandlers.Loop5PlayerAbilities,
		seedhandlers.Loop5RonsoRages,
		seedhandlers.Loop5Mixes,
		seedhandlers.Loop5CelestialWeapons,
		seedhandlers.Loop5AutoAbilities,
		seedhandlers.Loop5EquipmentTables,
		seedhandlers.Loop5EquipmentNames,
		seedhandlers.Loop5MonsterFormations,
		seedhandlers.Loop5StatChanges,
		seedhandlers.Loop5Alts,
	})
}

func seedLoop6(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 6")()
	return seedLoop(qtx, ctx, []seedFunc{
		seedhandlers.Loop6AeonEquipment,
		seedhandlers.Loop6Subquests,
		seedhandlers.Loop6MixCombinations,
		seedhandlers.Loop6AbilityPools,
		seedhandlers.Loop6TreasureEquipment,
		seedhandlers.Loop6ShopEquipment,
		seedhandlers.Loop6EquipmentDrops,
	})
}

func seedLoop7(qtx *database.Queries, ctx context.Context) error {
	defer h.MeasureTime("- seed loop 7")()
	return seedLoop(qtx, ctx, []seedFunc{
		seedhandlers.Loop7ArenaCreations,
	})
}

func seedJunctions(qtx *database.Queries, ctx context.Context) error {
	return seedLoop(qtx, ctx, []seedFunc{
		junctions.OverdriveModeActions,
		junctions.PropertiesRelatedStats,
		junctions.StatusConditionModifierChanges,
		junctions.StatusConditionRelatedStats,
		junctions.StatusConditionRemovedConditions,
		junctions.StatusConditionStatChanges,
		junctions.AeonBaseStatsA,
		junctions.AeonBaseStatsB,
		junctions.AeonBaseStatsX,
		junctions.AeonAeonEquipment,
		junctions.CharacterClassesPlayerUnits,
		junctions.DefaultAbilities,
		junctions.DefaultOverdriveAbilities,
		junctions.CharactersBaseStats,
		junctions.AeonCommandsPossibleAbilities,
		junctions.SubmenusUsers,
		junctions.AbilitiesBattleInteractions,
		junctions.OverdriveAbilitiesRelatedStats,
		junctions.OverdrivesOverdriveAbilities,
		junctions.PlayerAbilitiesLearnedBy,
		junctions.PlayerAbilitiesRelatedStats,
		junctions.TriggerCommandsRelatedStats,
		junctions.MiscAbilitiesLearnedBy,
		junctions.ItemsAvailableMenus,
		junctions.ItemsRelatedStats,
		junctions.AbilityPoolsAutoAbilities,
		junctions.AutoAbilitiesAddedStatusResists,
		junctions.AutoAbilitiesAddedStatusses,
		junctions.AutoAbilitiesAutoItems,
		junctions.AutoAbilitiesLockedOutAbilities,
		junctions.AutoAbilitiesModifierChanges,
		junctions.AutoAbilitiesRelatedStats,
		junctions.AutoAbilitiesStatChanges,
		junctions.EquipmentTablesNames,
		junctions.EquipmentReqAutoAbilities,
		junctions.AreaConnectedAreas,
		junctions.ShopEquipmentAutoAbilities,
		junctions.ShopsShopItems,
		junctions.TreasureEquipmentAutoAbilities,
		junctions.TreasuresItems,
		junctions.SongsBackgroundMusic,
		junctions.CuesIncludedAreas,
		junctions.BattleInteractionsAffectedBy,
		junctions.BattleInteractionsCopiedStatusConditions,
		junctions.BattleInteractionsDamages,
		junctions.BattleInteractionsInflictedStatusConditions,
		junctions.BattleInteractionsModifierChanges,
		junctions.BattleInteractionsRemovedStatusConditions,
		junctions.BattleInteractionsStatChanges,
		junctions.AltsAutoAbilities,
		junctions.AltsBaseStats,
		junctions.AltsElementalResists,
		junctions.AltsProperties,
		junctions.AltsStatusImmunities,
		junctions.MonsterItemsOtherItems,
		junctions.MonsterEquipmentEquipmentDrops,
		junctions.MonsterEquipmentSlotChances,
		junctions.EquipmentDropsCharacters,
		junctions.FormationTriggerCommandsUsers,
		junctions.MonsterFormationsEncounterAreas,
		junctions.MonsterFormationsTriggerCommands,
		junctions.MonsterSelectionMonsterAmounts,
		junctions.MonstersMonsterAbilities,
		junctions.MonstersAutoAbilities,
		junctions.MonstersBaseStats,
		junctions.MonstersElementalResists,
		junctions.MonstersStatusImmunities,
		junctions.MonstersProperties,
		junctions.MonstersRonsoRages,
		junctions.MonstersStatusResists,
	})
}
