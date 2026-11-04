package enums

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (T *Enums) initAbilityType() {
	enumDescription := "States the type of an ability in cases, when its general endpoint is used."

	typeSlice := []EnumVal{
		{
			Name:        string(database.AbilityTypePlayerAbility),
			Description: "Abilities that can either be learned via the sphere grid and/or are learned by aeons.",
		},
		{
			Name:        string(database.AbilityTypeOverdriveAbility),
			Description: "Abilities that are accessed by using an overdrive.",
		},
		{
			Name:        string(database.AbilityTypeItemAbility),
			Description: "Abilities that are accessed by using the item of the same name.",
		},
		{
			Name:        string(database.AbilityTypeTriggerCommand),
			Description: "Abilities that are only available in specific boss fights.",
		},
		{
			Name:        string(database.AbilityTypeMiscAbility),
			Description: "Abilities that don't fit the other categories. Most of these are accessible from the start of the game.",
		},
		{
			Name:        string(database.AbilityTypeEnemyAbility),
			Description: "Abilities that are used by monsters.",
		},
	}

	T.AbilityType = EnumType[database.AbilityType, any]{
		Name:         ts.EnumNameAbilityType,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.AbilityType { return database.AbilityType(s) },
		NullConvFunc: nil,
		GetNullEnum:  nil,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameAbilityType)] = EnumResponse{
		Name:               ts.EnumNameAbilityType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAbilities, ts.EpPlayerAbilities, ts.EpOverdriveAbilities, ts.EpItemAbilities, ts.EpTriggerCommands, ts.EpMiscAbilities, ts.EpEnemyAbilities, ts.EpAeonCommands, ts.EpCharacterClasses, ts.EpTopmenus, ts.EpSubmenus, ts.EpMonsters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initUnitType() {
	enumDescription := "States the type of a player unit in cases, when its general endpoint is used."

	typeSlice := []EnumVal{
		{
			Name:        string(database.UnitTypeCharacter),
			Description: "",
		},
		{
			Name:        string(database.UnitTypeAeon),
			Description: "",
		},
	}

	T.UnitType = EnumType[database.UnitType, any]{
		Name:         ts.EnumNameUnitType,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.UnitType { return database.UnitType(s) },
		NullConvFunc: nil,
		GetNullEnum:  nil,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameUnitType)] = EnumResponse{
		Name:               ts.EnumNameUnitType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAeons, ts.EpCharacterClasses, ts.EpCharacters, ts.EpPlayerUnits},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initItemType() {
	enumDescription := "States the type of an item in cases, when its general endpoint is used."

	typeSlice := []EnumVal{
		{
			Name:        string(database.ItemTypeItem),
			Description: "",
		},
		{
			Name:        string(database.ItemTypeKeyItem),
			Description: "",
		},
	}

	T.ItemType = EnumType[database.ItemType, any]{
		Name:         ts.EnumNameItemType,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.ItemType { return database.ItemType(s) },
		NullConvFunc: nil,
		GetNullEnum:  nil,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameItemType)] = EnumResponse{
		Name:               ts.EnumNameItemType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpArenaCreations, ts.EpAutoAbilities, ts.EpBlitzballPrizes, ts.EpItems, ts.EpKeyItems, ts.EpAllItems, ts.EpMonsters, ts.EpQuests, ts.EpSidequests, ts.EpSubquests, ts.EpTreasures, ts.EpShops},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initQuestType() {
	enumDescription := "States the type of a quest in cases, when its general endpoint is used."

	typeSlice := []EnumVal{
		{
			Name:        string(database.QuestTypeSidequest),
			Description: "",
		},
		{
			Name:        string(database.QuestTypeSubquest),
			Description: "",
		},
	}

	T.QuestType = EnumType[database.QuestType, any]{
		Name:         ts.EnumNameQuestType,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.QuestType { return database.QuestType(s) },
		NullConvFunc: nil,
		GetNullEnum:  nil,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameQuestType)] = EnumResponse{
		Name:               ts.EnumNameQuestType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpQuests, ts.EpSidequests, ts.EpSubquests},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initAaActivationCondition() {
	enumDescription := "Determines, when the effects of an auto-ability are active."

	typeSlice := []EnumVal{
		{
			Name:        string(database.AaActivationConditionAlways),
			Description: "The auto-ability is always active in-battle.",
		},
		{
			Name:        string(database.AaActivationConditionActiveParty),
			Description: "The auto-ability is only active in-battle, while the wearer is in the active party.",
		},
		{
			Name:        string(database.AaActivationConditionHpCritical),
			Description: "The auto-ability activates in-battle, while the wearer is in hp-critical condition.",
		},
		{
			Name:        string(database.AaActivationConditionOutsideBattle),
			Description: "The auto-ability's effects apply outside of battle.",
		},
	}

	T.AaActivationCondition = EnumType[database.AaActivationCondition, any]{
		Name:     ts.EnumNameAaActivationCondition,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.AaActivationCondition { return database.AaActivationCondition(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameAaActivationCondition)] = EnumResponse{
		Name:               ts.EnumNameAaActivationCondition,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAutoAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initAreaConnectionType() {
	enumDescription := "Determines, how two areas are connected with each other."

	typeSlice := []EnumVal{
		{
			Name:        string(database.AreaConnectionTypeBothDirections),
			Description: "The edges of two areas are directly connected with each other, and you can freely move between those areas.",
		},
		{
			Name:        string(database.AreaConnectionTypeOneDirection),
			Description: "The edges of two areas are directly connected with each other, but you can only move from area A to area B, and not vice versa.",
		},
		{
			Name:        string(database.AreaConnectionTypeWarp),
			Description: "A connection of two areas that doesn't require crossing their edges. Most of the time, their edges are not directly connected, but you can reach area B through other means. That might be due to a teleporter (like in Gagazet), or due to a story-based warp.",
		},
	}

	T.AreaConnectionType = EnumType[database.AreaConnectionType, any]{
		Name:     ts.EnumNameAreaConnectionType,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.AreaConnectionType { return database.AreaConnectionType(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameAreaConnectionType)] = EnumResponse{
		Name:               ts.EnumNameAreaConnectionType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAreas},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initArenaCreationCategory() {
	enumDescription := "The three categories of monster creations in the arena."

	typeSlice := []EnumVal{
		{
			Name: string(database.MaCreationCategoryArea),
		},
		{
			Name: string(database.MaCreationCategorySpecies),
		},
		{
			Name: string(database.MaCreationCategoryOriginal),
		},
	}

	T.ArenaCreationCategory = EnumType[database.MaCreationCategory, database.NullMaCreationCategory]{
		Name:         ts.EnumNameArenaCreationCategory,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.MaCreationCategory { return database.MaCreationCategory(s) },
		NullConvFunc: database.ToNullMaCreationCategory,
		GetNullEnum:  database.GetNullMaCreationCategory,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameArenaCreationCategory)] = EnumResponse{
		Name:               ts.EnumNameArenaCreationCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpArenaCreations},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initArmorType() {
	enumDescription := "The eight armor types, each one being associated with a character."

	typeSlice := []EnumVal{
		{
			Name: string(database.ArmorTypeShield),
		},
		{
			Name: string(database.ArmorTypeRing),
		},
		{
			Name: string(database.ArmorTypeArmguard),
		},
		{
			Name: string(database.ArmorTypeBangle),
		},
		{
			Name: string(database.ArmorTypeArmlet),
		},
		{
			Name: string(database.ArmorTypeBracer),
		},
		{
			Name: string(database.ArmorTypeTarge),
		},
		{
			Name: string(database.ArmorTypeSeymourArmor),
		},
	}

	T.ArmorType = EnumType[database.ArmorType, any]{
		Name:     ts.EnumNameArmorType,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.ArmorType { return database.ArmorType(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameArmorType)] = EnumResponse{
		Name:               ts.EnumNameArmorType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpCharacters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initArranger() {
	enumDescription := "The four people who arranged songs for this game."

	typeSlice := []EnumVal{
		{
			Name: string(database.ArrangerNobuouematsu),
		},
		{
			Name: string(database.ArrangerJunyanakano),
		},
		{
			Name: string(database.ArrangerMasashihamauzu),
		},
		{
			Name: string(database.ArrangerShirohamaguchi),
		},
	}

	T.Arranger = EnumType[database.Arranger, database.NullArranger]{
		Name:         ts.EnumNameArranger,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.Arranger { return database.Arranger(s) },
		NullConvFunc: database.ToNullArranger,
		GetNullEnum:  database.GetNullArranger,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameArranger)] = EnumResponse{
		Name:               ts.EnumNameArranger,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpSongs},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initAutoAbilityCategory() {
	enumDescription := "Divides auto-abilities into smaller categories."

	typeSlice := []EnumVal{
		{
			Name:        string(database.AutoAbilityCategoryStatX),
			Description: "Auto-abilities that increase stats or modify formulae related to that stat.",
		},
		{
			Name:        string(database.AutoAbilityCategoryElementalStrike),
			Description: "Auto-abilities that grant elemental properties to the user's attack and skills.",
		},
		{
			Name:        string(database.AutoAbilityCategoryElementalProtection),
			Description: "Auto-abilities that grant protection from elements.",
		},
		{
			Name:        string(database.AutoAbilityCategoryStatusInfliction),
			Description: "Auto-abilities that grant the chance of inflicting a status to the user's attack and skills.",
		},
		{
			Name:        string(database.AutoAbilityCategoryStatusProtection),
			Description: "Auto-abilities that grant protection from status conditions.",
		},
		{
			Name:        string(database.AutoAbilityCategoryAutoCure),
			Description: "Auto-abilities that let the user use restorative items automatically.",
		},
		{
			Name:        string(database.AutoAbilityCategoryAutoStatus),
			Description: "Auto-abilities that grant a positive status to the user at all times.",
		},
		{
			Name:        string(database.AutoAbilityCategorySosStatus),
			Description: "Auto-abilities that grant a positive status to the user, if they are in hp-critical condition.",
		},
		{
			Name:        string(database.AutoAbilityCategoryCounter),
			Description: "Auto-abilities that let the user perform a counterattack, if a certain condition is met.",
		},
		{
			Name:        string(database.AutoAbilityCategoryApOverdrive),
			Description: "Auto-abilities that modify the user's overdrive charge rate or ap gain.",
		},
		{
			Name:        string(database.AutoAbilityCategoryBreakLimit),
			Description: "Auto-abilities that raise the upper limit of the user's stats or damage.",
		},
		{
			Name:        string(database.AutoAbilityCategoryOther),
			Description: "Auto-abilities that don't match the other categories.",
		},
	}

	T.AutoAbilityCategory = EnumType[database.AutoAbilityCategory, any]{
		Name:     ts.EnumNameAutoAbilityCategory,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.AutoAbilityCategory { return database.AutoAbilityCategory(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameAutoAbilityCategory)] = EnumResponse{
		Name:               ts.EnumNameAutoAbilityCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAutoAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initAvailabilityType() {
	enumDescription := "Defines, when in the game a resource is available. The four values arise from two conditions: pre- vs. post-airship and story-only vs. always available. If both conditions are false, the resource is available, as soon as the point in the game is reached where it is obtainable ('always'). One of the two conditions being true can be seen as equally hard to get ('pre-story' and 'post'). Which one is easier to reach, depends on where the player is in the game. Both conditions being true equates to the hardest level of obtainability (post-story)."

	typeSlice := []EnumVal{
		{
			Name:        string(database.AvailabilityTypeAlways),
			Description: "The resource becomes and will always stay available once its location is reached in the story. Sometimes other conditions have to be met to unlock the resource. This only includes resources that are available before acquiring the airship.",
		},
		{
			Name:        string(database.AvailabilityTypePreStory),
			Description: "The resource is only available during certain events of the story, before acquiring the airship. This includes scripted fights, tutorials, and locations that are only accessible during those events.",
		},
		{
			Name:        string(database.AvailabilityTypePost),
			Description: "The resource becomes available, or is only accessible after acquiring the airship. Note that in order to get access to the resources that are located inside Sin, you have to do the boss rush against it.",
		},
		{
			Name:        string(database.AvailabilityTypePostStory),
			Description: "The resource is only available during the events of the story that happen after acquiring the airship. These resources are either post-airship story bosses or past the point of no return.",
		},
		{
			Name:        string(database.AvailabilityTypePostGame),
			Description: "The resource is available in the post-game, meaning it either was already available before acquiring the airship, or it becomes available after acquiring the airship. This excludes story-specific resources. This value is an alias for the values 'always' and 'post'.",
		},
		{
			Name:        string(database.AvailabilityTypeStory),
			Description: "The resource is only available during the events of the story and thus is missable. This value is an alias for the values 'pre-story' and 'post-story'.",
		},
		{
			Name:        string(database.AvailabilityTypePreAirship),
			Description: "The resource is available before acquiring the airship. This value is an alias for the values 'always' and 'pre-story'. If only this value is used in an 'availability' query, the 'pre_airship' param gets set to 'true' automatically.",
		},
		{
			Name:        string(database.AvailabilityTypePostAirship),
			Description: "The resource is only available after acquiring the airship. This value is an alias for the values 'post' and 'post-story'.",
		},
	}

	T.AvailabilityType = EnumType[database.AvailabilityType, any]{
		Name:     ts.EnumNameAvailabilityType,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.AvailabilityType { return database.AvailabilityType(s) },
		Aliasses: map[string][]database.AvailabilityType{
			string(database.AvailabilityTypePostGame): {
				database.AvailabilityTypeAlways,
				database.AvailabilityTypePost,
			},

			string(database.AvailabilityTypeStory): {
				database.AvailabilityTypePreStory,
				database.AvailabilityTypePostStory,
			},

			string(database.AvailabilityTypePreAirship): {
				database.AvailabilityTypeAlways,
				database.AvailabilityTypePreStory,
			},

			string(database.AvailabilityTypePostAirship): {
				database.AvailabilityTypePost,
				database.AvailabilityTypePostStory,
			},
		},
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameAvailabilityType)] = EnumResponse{
		Name:               ts.EnumNameAvailabilityType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAutoAbilities, ts.EpEquipment, ts.EpItems, ts.EpKeyItems, ts.EpLocations, ts.EpSublocations, ts.EpAreas, ts.EpAllItems, ts.EpMonsterFormations, ts.EpMonsters, ts.EpPrimers, ts.EpQuests, ts.EpSidequests, ts.EpSubquests, ts.EpShops, ts.EpSpheres, ts.EpTreasures},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initBgReplacementType() {
	enumDescription := "Describes the condition until which alternative background music replaces the regular background music of a particular set of areas."

	typeSlice := []EnumVal{
		{
			Name:        string(database.BgReplacementTypeUntilTrigger),
			Description: "The background music continues playing until a certain story- or event-trigger occurs.",
		},
		{
			Name:        string(database.BgReplacementTypeUntilZoneChange),
			Description: "The background music continues playing until the current area is left and a zone-change occurs.",
		},
	}

	T.BgReplacementType = EnumType[database.BgReplacementType, database.NullBgReplacementType]{
		Name:         ts.EnumNameBgReplacementType,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.BgReplacementType { return database.BgReplacementType(s) },
		NullConvFunc: database.ToNullBgReplacementType,
		GetNullEnum:  database.GetNullBgReplacementType,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameBgReplacementType)] = EnumResponse{
		Name:               ts.EnumNameBgReplacementType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpSongs},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initBlitzballPositionSlot() {
	enumDescription := "The four end results after a blitzball tournament or league that result in a prize being given."

	typeSlice := []EnumVal{
		{
			Name: string(database.BlitzballPositionSlot1st),
		},
		{
			Name: string(database.BlitzballPositionSlot2nd),
		},
		{
			Name: string(database.BlitzballPositionSlot3rd),
		},
		{
			Name: string(database.BlitzballPositionSlotTopScorer),
		},
	}

	T.BlitzballPositionSlot = EnumType[database.BlitzballPositionSlot, any]{
		Name:     ts.EnumNameBlitzballPositionSlot,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.BlitzballPositionSlot { return database.BlitzballPositionSlot(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameBlitzballPositionSlot)] = EnumResponse{
		Name:               ts.EnumNameBlitzballPositionSlot,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpBlitzballPrizes},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initBlitzballTournamentCategory() {
	enumDescription := "The two different categories of blitzball competitions."

	typeSlice := []EnumVal{
		{
			Name: string(database.BlitzballTournamentCategoryLeague),
		},
		{
			Name: string(database.BlitzballTournamentCategoryTournament),
		},
	}

	T.BlitzballTournamentCategory = EnumType[database.BlitzballTournamentCategory, any]{
		Name:     ts.EnumNameBlitzballTournamentCategory,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.BlitzballTournamentCategory { return database.BlitzballTournamentCategory(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameBlitzballTournamentCategory)] = EnumResponse{
		Name:               ts.EnumNameBlitzballTournamentCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpBlitzballPrizes},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initCelestialFormula() {
	enumDescription := "The three formulae that are used for the damage calculation of a celestial weapon's skills and attack."

	typeSlice := []EnumVal{
		{
			Name:        string(database.CelestialFormulaHpHigh),
			Description: "The celestial weapon deals more damage, the higher the user's hp are.",
		},
		{
			Name:        string(database.CelestialFormulaHpLow),
			Description: "The celestial weapon deals more damage, the lower the user's hp are.",
		},
		{
			Name:        string(database.CelestialFormulaMpHigh),
			Description: "The celestial weapon deals more damage, the higher the user's mp are.",
		},
	}

	T.CelestialFormula = EnumType[database.CelestialFormula, any]{
		Name:     ts.EnumNameCelestialFormula,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.CelestialFormula { return database.CelestialFormula(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameCelestialFormula)] = EnumResponse{
		Name:               ts.EnumNameCelestialFormula,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpCelestialWeapons},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initCharacterClassCategory() {
	enumDescription := "Divides character-classes into smaller categories."

	typeSlice := []EnumVal{
		{
			Name: string(database.CharacterClassCategoryGroup),
		},
		{
			Name: string(database.CharacterClassCategoryCharacter),
		},
		{
			Name: string(database.CharacterClassCategoryAeon),
		},
	}

	T.CharacterClassCategory = EnumType[database.CharacterClassCategory, any]{
		Name:     ts.EnumNameCharacterClassCategory,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.CharacterClassCategory { return database.CharacterClassCategory(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameCharacterClassCategory)] = EnumResponse{
		Name:               ts.EnumNameCharacterClassCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpCharacterClasses},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initComposer() {
	enumDescription := "The three people who composed songs for this game."

	typeSlice := []EnumVal{
		{
			Name: string(database.ComposerNobuouematsu),
		},
		{
			Name: string(database.ComposerJunyanakano),
		},
		{
			Name: string(database.ComposerMasashihamauzu),
		},
	}

	T.Composer = EnumType[database.Composer, database.NullComposer]{
		Name:         ts.EnumNameComposer,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.Composer { return database.Composer(s) },
		NullConvFunc: database.ToNullComposer,
		GetNullEnum:  database.GetNullComposer,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameComposer)] = EnumResponse{
		Name:               ts.EnumNameComposer,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpSongs},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initCounterType() {
	enumDescription := "Describes the two ways, in which a counter attack can be triggered."

	typeSlice := []EnumVal{
		{
			Name:        string(database.CounterTypePhysical),
			Description: "The user counters when being targeted by a physical attack.",
		},
		{
			Name:        string(database.CounterTypeMagical),
			Description: "The user counters when being targeted by a magical attack.",
		},
	}

	T.CounterType = EnumType[database.CounterType, database.NullCounterType]{
		Name:         ts.EnumNameCounterType,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.CounterType { return database.CounterType(s) },
		NullConvFunc: database.ToNullCounterType,
		GetNullEnum:  database.GetNullCounterType,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameCounterType)] = EnumResponse{
		Name:               ts.EnumNameCounterType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAutoAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initCTBIconType() {
	enumDescription := "Describes the icon that can be assigned to a monster in the CTB window."

	typeSlice := []EnumVal{
		{
			Name:        string(database.CtbIconTypeMonster),
			Description: "Used for regular monsters",
		},
		{
			Name:        string(database.CtbIconTypeBoss),
			Description: "Used for bosses",
		},
		{
			Name:        string(database.CtbIconTypeBossNumbered),
			Description: "Used for multiple bosses, or subparts of a boss",
		},
		{
			Name:        string(database.CtbIconTypeSummon),
			Description: "Used for aeons, except dark aeons",
		},
		{
			Name:        string(database.CtbIconTypeCid),
			Description: "Used for Cid during the Evrae fight",
		},
	}

	T.CTBIconType = EnumType[database.CtbIconType, any]{
		Name:     ts.EnumNameCTBIconType,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.CtbIconType { return database.CtbIconType(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameCTBIconType)] = EnumResponse{
		Name:               ts.EnumNameCTBIconType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpMonsters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initCreationArea() {
	enumDescription := "The locations in which monsters can be captured for the area conquest. While monsters can appear in different locations, each monster is only assigned one location for the area conquest. Some creation areas encompass more than one location, like Djose (includes Moonflow) or Mount Gagazet (includes Zanarkand Ruins)."

	typeSlice := []EnumVal{
		{
			Name: string(database.MaCreationAreaBesaid),
		},
		{
			Name: string(database.MaCreationAreaKilika),
		},
		{
			Name: string(database.MaCreationAreaMiihenHighroad),
		},
		{
			Name: string(database.MaCreationAreaMushroomRockRoad),
		},
		{
			Name: string(database.MaCreationAreaDjose),
		},
		{
			Name: string(database.MaCreationAreaThunderPlains),
		},
		{
			Name: string(database.MaCreationAreaMacalania),
		},
		{
			Name: string(database.MaCreationAreaBikanel),
		},
		{
			Name: string(database.MaCreationAreaCalmLands),
		},
		{
			Name: string(database.MaCreationAreaCavernOfTheStolenFayth),
		},
		{
			Name: string(database.MaCreationAreaMountGagazet),
		},
		{
			Name: string(database.MaCreationAreaSin),
		},
		{
			Name: string(database.MaCreationAreaOmegaRuins),
		},
	}

	T.CreationArea = EnumType[database.MaCreationArea, database.NullMaCreationArea]{
		Name:         ts.EnumNameCreationArea,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.MaCreationArea { return database.MaCreationArea(s) },
		NullConvFunc: database.ToNullMaCreationArea,
		GetNullEnum:  database.GetNullMaCreationArea,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameCreationArea)] = EnumResponse{
		Name:               ts.EnumNameCreationArea,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpMonsters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initCreationsUnlockedCategory() {
	enumDescription := "Is used to determine, which of the other two conquests is required to have a certain amount of monsters to unlock an original creation in the monster arena."

	typeSlice := []EnumVal{
		{
			Name: string(database.CreationsUnlockedCategoryArea),
		},
		{
			Name: string(database.CreationsUnlockedCategorySpecies),
		},
	}

	T.CreationsUnlockedCategory = EnumType[database.CreationsUnlockedCategory, database.NullCreationsUnlockedCategory]{
		Name:         ts.EnumNameCreationsUnlockedCategory,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.CreationsUnlockedCategory { return database.CreationsUnlockedCategory(s) },
		NullConvFunc: database.ToNullCreationsUnlockedCategory,
		GetNullEnum:  database.GetNullCreationsUnlockedCategory,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameCreationsUnlockedCategory)] = EnumResponse{
		Name:               ts.EnumNameCreationsUnlockedCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpArenaCreations},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initElementalAffinity() {
	enumDescription := "Determines how an element affects the damage a target takes."

	typeSlice := []EnumVal{
		{
			Name:        string(database.ElementalAffinityNeutral),
			Description: "The element doesn't affect the damage the target takes.",
		},
		{
			Name:        string(database.ElementalAffinityWeak),
			Description: "The damage the target takes from attacks bearing this element is multiplied by 1.5.",
		},
		{
			Name:        string(database.ElementalAffinityHalved),
			Description: "The damage the target takes from attacks bearing this element is halved.",
		},
		{
			Name:        string(database.ElementalAffinityImmune),
			Description: "The target takes 0 damage from attacks bearing this element.",
		},
		{
			Name:        string(database.ElementalAffinityAbsorb),
			Description: "Instead of taking damage, the target is healed by the same amount, if it is hit by an attack bearing this element.",
		},
		{
			Name:        string(database.ElementalAffinityVaries),
			Description: "The target's relationship to this element is not fixed and is dependent on other factors.",
		},
	}

	T.ElementalAffinity = EnumType[database.ElementalAffinity, any]{
		Name:     ts.EnumNameElementalAffinity,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.ElementalAffinity { return database.ElementalAffinity(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameElementalAffinity)] = EnumResponse{
		Name:               ts.EnumNameElementalAffinity,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAutoAbilities, ts.EpMonsters, ts.EpStatusConditions},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initEquipClass() {
	enumDescription := "Determines, if an equipment piece can be customized, or if its auto-abilities depend on different factors. Non-standard equipment is one-of-a-kind."

	typeSlice := []EnumVal{
		{
			Name:        string(database.EquipClassStandard),
			Description: "A standard, customizable equipment piece.",
		},
		{
			Name:        string(database.EquipClassUnique),
			Description: "The equipment piece is one-of-a-kind and its auto-abilities can only be modified by progressing through the story.",
		},
		{
			Name:        string(database.EquipClassCelestialWeapon),
			Description: "The equipment piece is a celestial weapon and its auto-abilities can only be modified by upgrading it with its equivalent crest and sigil.",
		},
	}

	T.EquipClass = EnumType[database.EquipClass, any]{
		Name:     ts.EnumNameEquipClass,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.EquipClass { return database.EquipClass(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameEquipClass)] = EnumResponse{
		Name:               ts.EnumNameEquipClass,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpEquipment, ts.EpEquipmentTables},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initEquipType() {
	enumDescription := "The two types of equipment."

	typeSlice := []EnumVal{
		{
			Name: string(database.EquipTypeWeapon),
		},
		{
			Name: string(database.EquipTypeArmor),
		},
	}

	T.EquipType = EnumType[database.EquipType, any]{
		Name:     ts.EnumNameEquipType,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.EquipType { return database.EquipType(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameEquipType)] = EnumResponse{
		Name:               ts.EnumNameEquipType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAutoAbilities, ts.EpEquipment, ts.EpEquipmentTables},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initItemCategory() {
	enumDescription := "Divides items into smaller categories."

	typeSlice := []EnumVal{
		{
			Name:        string(database.ItemCategoryHealing),
			Description: "Items that are used for recovery of HP and MP, or for curing negative status ailments.",
		},
		{
			Name:        string(database.ItemCategoryOffensive),
			Description: "Items that deal damage to other enemies or inflict status ailments.",
		},
		{
			Name:        string(database.ItemCategorySupport),
			Description: "Items that grant positive statusses or other supportive effects.",
		},
		{
			Name:        string(database.ItemCategorySphere),
			Description: "Items that can only be used within the sphere grid.",
		},
		{
			Name:        string(database.ItemCategoryDistiller),
			Description: "Items that cause enemies to drop spheres.",
		},
		{
			Name:        string(database.ItemCategoryOther),
			Description: "Uncategorized items, that are mostly used for mixes.",
		},
	}

	T.ItemCategory = EnumType[database.ItemCategory, any]{
		Name:     ts.EnumNameItemCategory,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.ItemCategory { return database.ItemCategory(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameItemCategory)] = EnumResponse{
		Name:               ts.EnumNameItemCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpItems, ts.EpItemAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initItemUsability() {
	enumDescription := "Determines, when and where an item can be used."

	typeSlice := []EnumVal{
		{
			Name:        string(database.ItemUsabilityAlways),
			Description: "The item can be used in battle, as well as outside of battle, in a menu.",
		},
		{
			Name:        string(database.ItemUsabilityInBattle),
			Description: "This item can only be used in battle.",
		},
		{
			Name:        string(database.ItemUsabilityOutsideBattle),
			Description: "This item can only be used outside of battle, in a menu.",
		},
		{
			Name:        string(database.ItemUsabilityUnusable),
			Description: "This item can't be used directly. It can only be used for customization and mixing.",
		},
	}

	T.ItemUsability = EnumType[database.ItemUsability, any]{
		Name:     ts.EnumNameItemUsability,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.ItemUsability { return database.ItemUsability(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameItemUsability)] = EnumResponse{
		Name:               ts.EnumNameItemUsability,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpItems},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initKeyItemCategory() {
	enumDescription := "Divides key-items into smaller categories."

	typeSlice := []EnumVal{
		{
			Name:        string(database.KeyItemCategoryStory),
			Description: "Key-items that are obtained during the course of the story.",
		},
		{
			Name:        string(database.KeyItemCategoryCelestial),
			Description: "Key-items that are related to the celestial weapons.",
		},
		{
			Name:        string(database.KeyItemCategoryPrimer),
			Description: "Key-items that are Al Bhed Primers.",
		},
		{
			Name:        string(database.KeyItemCategoryJechtSphere),
			Description: "Key-items that are Jecht Spheres.",
		},
		{
			Name:        string(database.KeyItemCategoryOther),
			Description: "Key-items that don't fit the other categories.",
		},
	}

	T.KeyItemCategory = EnumType[database.KeyItemCategory, any]{
		Name:     ts.EnumNameKeyItemCategory,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.KeyItemCategory { return database.KeyItemCategory(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameKeyItemCategory)] = EnumResponse{
		Name:               ts.EnumNameKeyItemCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpKeyItems},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initLootType() {

	enumDescription := "Determines the type of loot a treasure contains."

	typeSlice := []EnumVal{
		{
			Name: string(database.LootTypeItem),
		},
		{
			Name: string(database.LootTypeEquipment),
		},
		{
			Name: string(database.LootTypeGil),
		},
	}

	T.LootType = EnumType[database.LootType, any]{
		Name:     ts.EnumNameLootType,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.LootType { return database.LootType(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameLootType)] = EnumResponse{
		Name:               ts.EnumNameLootType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpTreasures},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initMixCategory() {
	enumDescription := "Divides mixes into smaller categories."

	typeSlice := []EnumVal{
		{
			Name:        string(database.MixCategoryRecovery),
			Description: "Mixes that heal the party.",
		},
		{
			Name:        string(database.MixCategoryBuffs),
			Description: "Mixes that grant positive status effects to the party.",
		},
		{
			Name:        string(database.MixCategoryHpMp),
			Description: "Mixes that double the party's HP or MP.",
		},
		{
			Name:        string(database.MixCategoryOverdriveSpeed),
			Description: "Mixes that multiply the charge speed of the party's overdrive gauges.",
		},
		{
			Name:        string(database.MixCategoryCriticalHits),
			Description: "Mixes that double the party's critical hit rate.",
		},
		{
			Name:        string(database.MixCategory9999Damage),
			Description: "Mixes that set the party's minimum amount of damage dealt to 9999.",
		},
		{
			Name:        string(database.MixCategoryFireElemental),
			Description: "Mixes that deal fire-elemental damage.",
		},
		{
			Name:        string(database.MixCategoryLightningElemental),
			Description: "Mixes that deal lightning-elemental damage.",
		},
		{
			Name:        string(database.MixCategoryWaterElemental),
			Description: "Mixes that deal water-elemental damage.",
		},
		{
			Name:        string(database.MixCategoryIceElemental),
			Description: "Mixes that deal ice-elemental damage.",
		},
		{
			Name:        string(database.MixCategoryNonElemental),
			Description: "Mixes that deal non-elemental damage.",
		},
		{
			Name:        string(database.MixCategoryGravityBased),
			Description: "Mixes that deal percentage-damage.",
		},
	}

	T.MixCategory = EnumType[database.MixCategory, any]{
		Name:     ts.EnumNameMixCategory,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.MixCategory { return database.MixCategory(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameMixCategory)] = EnumResponse{
		Name:               ts.EnumNameMixCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpMixes},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initModifierCategory() {
	enumDescription := "Divides modifiers into smaller categories."

	typeSlice := []EnumVal{
		{
			Name:        string(database.ModifierCategoryFixedValue),
			Description: "The modifier is a fixed value that can be changed to another value.",
		},
		{
			Name:        string(database.ModifierCategoryDynamicValue),
			Description: "The modifier is a value that depends on the context of the battle (current-hp, overdrive-gauge, initial-counter-value), or that can't be defined by a single value (mp-cost).",
		},
		{
			Name:        string(database.ModifierCategoryFactor),
			Description: "The modifier acts as a multiplier to its target.",
		},
		{
			Name:        string(database.ModifierCategoryPercentage),
			Description: "The modifier is a percentage from 0 to 100 in most cases. Meaning, these can be divided by 100 to convert them into factors. If a scenario for 255% is given in the modifier's description (255% being the game's equivalent for 'always'), then that value is also possible.",
		},
	}

	T.ModifierCategory = EnumType[database.ModifierCategory, any]{
		Name:     ts.EnumNameModifierCategory,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.ModifierCategory { return database.ModifierCategory(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameModifierCategory)] = EnumResponse{
		Name:               ts.EnumNameModifierCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpModifiers},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initMonsterFormationCategory() {
	enumDescription := "Divides monster-formations into smaller categories."

	typeSlice := []EnumVal{
		{
			Name:        string(database.MonsterFormationCategoryRandomEncounter),
			Description: "A typical random encounter which can effectively be triggered indefinitely.",
		},
		{
			Name:        string(database.MonsterFormationCategoryBossFight),
			Description: "A boss encounter. Can only be triggered once, usually during the events of the story.",
		},
		{
			Name:        string(database.MonsterFormationCategoryStoryFight),
			Description: "A story-based, non-boss-encounter. Is triggered during the events of the story. Usually once, unless stated otherwise.",
		},
		{
			Name:        string(database.MonsterFormationCategoryStaticEncounter),
			Description: "An encounter that is triggered by interacting with the enemy in the overworld. You can flee from these encounters. This only applies to Lord Ochu in Kilika, the Sandragoras in Bikanel and both Dark Ixion fights.",
		},
		{
			Name:        string(database.MonsterFormationCategoryTutorial),
			Description: "A unique tutorial fight. Can only be triggered once.",
		},
		{
			Name:        string(database.MonsterFormationCategoryOnDemandFight),
			Description: "An encounter that can be triggered indefinitely via the Monster Arena.",
		},
	}

	T.MonsterFormationCategory = EnumType[database.MonsterFormationCategory, any]{
		Name:     ts.EnumNameMonsterFormationCategory,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.MonsterFormationCategory { return database.MonsterFormationCategory(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameMonsterFormationCategory)] = EnumResponse{
		Name:               ts.EnumNameMonsterFormationCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpMonsterFormations},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initMonsterSpecies() {
	enumDescription := "Determines the species of a monster. Some of these are relevant for the species conquest. Monsters of the same species tend to use similar attacks and attack-patterns."

	typeSlice := []EnumVal{
		{
			Name: string(database.MonsterSpeciesAdamantoise),
		},
		{
			Name: string(database.MonsterSpeciesAeon),
		},
		{
			Name: string(database.MonsterSpeciesArmor),
		},
		{
			Name: string(database.MonsterSpeciesBasilisk),
		},
		{
			Name: string(database.MonsterSpeciesBlade),
		},
		{
			Name: string(database.MonsterSpeciesBehemoth),
		},
		{
			Name: string(database.MonsterSpeciesBird),
		},
		{
			Name: string(database.MonsterSpeciesBomb),
		},
		{
			Name: string(database.MonsterSpeciesCactuar),
		},
		{
			Name: string(database.MonsterSpeciesCephalopod),
		},
		{
			Name: string(database.MonsterSpeciesChest),
		},
		{
			Name: string(database.MonsterSpeciesChimera),
		},
		{
			Name: string(database.MonsterSpeciesCoeurl),
		},
		{
			Name: string(database.MonsterSpeciesDefender),
		},
		{
			Name: string(database.MonsterSpeciesDinofish),
		},
		{
			Name: string(database.MonsterSpeciesDoomstone),
		},
		{
			Name: string(database.MonsterSpeciesDrake),
		},
		{
			Name: string(database.MonsterSpeciesEater),
		},
		{
			Name: string(database.MonsterSpeciesElemental),
		},
		{
			Name: string(database.MonsterSpeciesEvilEye),
		},
		{
			Name: string(database.MonsterSpeciesFlan),
		},
		{
			Name: string(database.MonsterSpeciesFungus),
		},
		{
			Name: string(database.MonsterSpeciesGel),
		},
		{
			Name: string(database.MonsterSpeciesGeo),
		},
		{
			Name: string(database.MonsterSpeciesHaizhe),
		},
		{
			Name: string(database.MonsterSpeciesHelm),
		},
		{
			Name: string(database.MonsterSpeciesHermit),
		},
		{
			Name: string(database.MonsterSpeciesHumanoid),
		},
		{
			Name: string(database.MonsterSpeciesImp),
		},
		{
			Name: string(database.MonsterSpeciesIronGiant),
		},
		{
			Name: string(database.MonsterSpeciesLarva),
		},
		{
			Name: string(database.MonsterSpeciesLupine),
		},
		{
			Name: string(database.MonsterSpeciesMachina),
		},
		{
			Name: string(database.MonsterSpeciesMalboro),
		},
		{
			Name: string(database.MonsterSpeciesMech),
		},
		{
			Name: string(database.MonsterSpeciesMimic),
		},
		{
			Name: string(database.MonsterSpeciesOchu),
		},
		{
			Name: string(database.MonsterSpeciesOgre),
		},
		{
			Name: string(database.MonsterSpeciesPhantom),
		},
		{
			Name: string(database.MonsterSpeciesPiranha),
		},
		{
			Name: string(database.MonsterSpeciesPlant),
		},
		{
			Name: string(database.MonsterSpeciesReptile),
		},
		{
			Name: string(database.MonsterSpeciesRoc),
		},
		{
			Name: string(database.MonsterSpeciesRuminant),
		},
		{
			Name: string(database.MonsterSpeciesSacredBeast),
		},
		{
			Name: string(database.MonsterSpeciesSahagin),
		},
		{
			Name: string(database.MonsterSpeciesSin),
		},
		{
			Name: string(database.MonsterSpeciesSinspawn),
		},
		{
			Name: string(database.MonsterSpeciesSpellspinner),
		},
		{
			Name: string(database.MonsterSpeciesSpiritBeast),
		},
		{
			Name: string(database.MonsterSpeciesTonberry),
		},
		{
			Name: string(database.MonsterSpeciesUnspecified),
		},
		{
			Name: string(database.MonsterSpeciesWasp),
		},
		{
			Name: string(database.MonsterSpeciesWeapon),
		},
		{
			Name: string(database.MonsterSpeciesWorm),
		},
		{
			Name: string(database.MonsterSpeciesWyrm),
		},
	}

	T.MonsterSpecies = EnumType[database.MonsterSpecies, any]{
		Name:     ts.EnumNameMonsterSpecies,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.MonsterSpecies { return database.MonsterSpecies(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameMonsterSpecies)] = EnumResponse{
		Name:               ts.EnumNameMonsterSpecies,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpMonsters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initMonsterCategory() {
	enumDescription := "Divides monsters into smaller categories."

	typeSlice := []EnumVal{
		{
			Name: string(database.MonsterCategoryMonster),
		},
		{
			Name: string(database.MonsterCategoryBoss),
		},
		{
			Name: string(database.MonsterCategorySummon),
		},
	}

	T.MonsterCategory = EnumType[database.MonsterCategory, any]{
		Name:     ts.EnumNameMonsterCategory,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.MonsterCategory { return database.MonsterCategory(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameMonsterCategory)] = EnumResponse{
		Name:               ts.EnumNameMonsterCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpMonsters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initMusicUseCase() {
	enumDescription := "Defines special use cases for songs that don't fall into the other listed categories, like the chocobo theme or the main menu theme."

	typeSlice := []EnumVal{
		{
			Name: string(database.MusicUseCaseBlitzballGame),
		},
		{
			Name: string(database.MusicUseCaseBlitzballMenu),
		},
		{
			Name: string(database.MusicUseCaseBossBattleDefault),
		},
		{
			Name: string(database.MusicUseCaseChocobo),
		},
		{
			Name: string(database.MusicUseCaseGameOver),
		},
		{
			Name: string(database.MusicUseCaseMainMenu),
		},
		{
			Name: string(database.MusicUseCaseRandomEncounterDefault),
		},
		{
			Name: string(database.MusicUseCaseVictory),
		},
	}

	T.MusicUseCase = EnumType[database.MusicUseCase, database.NullMusicUseCase]{
		Name:         ts.EnumNameMusicUseCase,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.MusicUseCase { return database.MusicUseCase(s) },
		NullConvFunc: database.ToNullMusicUseCase,
		GetNullEnum:  database.GetNullMusicUseCase,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameMusicUseCase)] = EnumResponse{
		Name:               ts.EnumNameMusicUseCase,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpSongs},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initNodePosition() {
	enumDescription := "Defines which nodes a sphere can target."

	typeSlice := []EnumVal{
		{
			Name:        string(database.NodePositionNeighboring),
			Description: "The sphere can target neighboring nodes, or the node the selected character is currently positioned.",
		},
		{
			Name:        string(database.NodePositionAllyPosition),
			Description: "The sphere can only target nodes, where another character is currently positioned.",
		},
		{
			Name:        string(database.NodePositionAny),
			Description: "The sphere can target any node that it is able to.",
		},
	}

	T.NodePosition = EnumType[database.NodePosition, any]{
		Name:     ts.EnumNameNodePosition,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.NodePosition { return database.NodePosition(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameNodePosition)] = EnumResponse{
		Name:               ts.EnumNameNodePosition,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpSpheres},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initNodeState() {
	enumDescription := "Defines the state a node is in, in relation to all currently available characters on the sphere grid."

	typeSlice := []EnumVal{
		{
			Name:        string(database.NodeStateActiveSelf),
			Description: "The node has been activated by the selected character.",
		},
		{
			Name:        string(database.NodeStateActiveAlly),
			Description: "The node hasn't been activated by the selected character, but by another character.",
		},
		{
			Name:        string(database.NodeStateActiveAny),
			Description: "The node has been activated by at least one character.",
		},
		{
			Name:        string(database.NodeStateInactive),
			Description: "The node hasn't been activated by the selected character.",
		},
		{
			Name:        string(database.NodeStateAny),
			Description: "The node's activation state doesn't matter for this resource.",
		},
	}

	T.NodeState = EnumType[database.NodeState, database.NullNodeState]{
		Name:         ts.EnumNameNodeState,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.NodeState { return database.NodeState(s) },
		NullConvFunc: database.ToNullNodeState,
		GetNullEnum:  database.GetNullNodeState,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameNodeState)] = EnumResponse{
		Name:               ts.EnumNameNodeState,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpSpheres},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initNodeType() {
	enumDescription := "Defines every node on the sphere grid."

	typeSlice := []EnumVal{
		{
			Name: string(database.NodeTypeHp),
		},
		{
			Name: string(database.NodeTypeMp),
		},
		{
			Name: string(database.NodeTypeStrength),
		},
		{
			Name: string(database.NodeTypeDefense),
		},
		{
			Name: string(database.NodeTypeMagic),
		},
		{
			Name: string(database.NodeTypeMagicDefense),
		},
		{
			Name: string(database.NodeTypeAgility),
		},
		{
			Name: string(database.NodeTypeLuck),
		},
		{
			Name: string(database.NodeTypeEvasion),
		},
		{
			Name: string(database.NodeTypeAccuracy),
		},
		{
			Name: string(database.NodeTypeSkill),
		},
		{
			Name: string(database.NodeTypeSpecial),
		},
		{
			Name: string(database.NodeTypeWhtMagic),
		},
		{
			Name: string(database.NodeTypeBlkMagic),
		},
		{
			Name: string(database.NodeTypeLv1Lock),
		},
		{
			Name: string(database.NodeTypeLv2Lock),
		},
		{
			Name: string(database.NodeTypeLv3Lock),
		},
		{
			Name: string(database.NodeTypeLv4Lock),
		},
		{
			Name: string(database.NodeTypeEmpty),
		},
	}

	T.NodeType = EnumType[database.NodeType, any]{
		Name:     ts.EnumNameNodeType,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.NodeType { return database.NodeType(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameNodeType)] = EnumResponse{
		Name:               ts.EnumNameNodeType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpSpheres},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initNullifyArmored() {
	enumDescription := "In some cases, a monster's 'armored' property is nullified. This either happens, if it is the target of an attack by a user with 'piercing', or if suffers under the 'armor break' status (bearer)."

	typeSlice := []EnumVal{
		{
			Name: string(database.NullifyArmoredTarget),
		},
		{
			Name: string(database.NullifyArmoredBearer),
		},
	}

	T.NullifyArmored = EnumType[database.NullifyArmored, database.NullNullifyArmored]{
		Name:         ts.EnumNameNullifyArmored,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.NullifyArmored { return database.NullifyArmored(s) },
		NullConvFunc: database.ToNullNullifyArmored,
		GetNullEnum:  database.GetNullNullifyArmored,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameNullifyArmored)] = EnumResponse{
		Name:               ts.EnumNameNullifyArmored,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpProperties, ts.EpStatusConditions},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initOverdriveModeType() {
	enumDescription := "Defines how an overdrive-mode fills up the overdrive gauge of its user."

	typeSlice := []EnumVal{
		{
			Name:        string(database.OverdriveModeTypeFormula),
			Description: "The fill-amount of the overdrive gauge is determined by a formula.",
		},
		{
			Name:        string(database.OverdriveModeTypePerAction),
			Description: "The overdrive gauge fills by a fixed amount every time the specified action is performed.",
		},
	}

	T.OverdriveModeType = EnumType[database.OverdriveModeType, any]{
		Name:     ts.EnumNameOverdriveModeType,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.OverdriveModeType { return database.OverdriveModeType(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameOverdriveModeType)] = EnumResponse{
		Name:               ts.EnumNameOverdriveModeType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpOverdriveModes},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initPlayerAbilityCategory() {
	enumDescription := "Divides player-abilities into smaller categories."

	typeSlice := []EnumVal{
		{
			Name: string(database.PlayerAbilityCategorySkill),
		},
		{
			Name: string(database.PlayerAbilityCategorySpecial),
		},
		{
			Name: string(database.PlayerAbilityCategoryWhiteMagic),
		},
		{
			Name: string(database.PlayerAbilityCategoryBlackMagic),
		},
		{
			Name: string(database.PlayerAbilityCategoryAeon),
		},
	}

	T.PlayerAbilityCategory = EnumType[database.PlayerAbilityCategory, any]{
		Name:     ts.EnumNamePlayerAbilityCategory,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.PlayerAbilityCategory { return database.PlayerAbilityCategory(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNamePlayerAbilityCategory)] = EnumResponse{
		Name:               ts.EnumNamePlayerAbilityCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpPlayerAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initShopCategory() {
	enumDescription := "Divides shops into smaller categories."

	typeSlice := []EnumVal{
		{
			Name: string(database.ShopCategoryStandard),
		},
		{
			Name: string(database.ShopCategoryOaka),
		},
		{
			Name: string(database.ShopCategoryTravelAgency),
		},
		{
			Name: string(database.ShopCategoryWantz),
		},
	}

	T.ShopCategory = EnumType[database.ShopCategory, any]{
		Name:     ts.EnumNameShopCategory,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.ShopCategory { return database.ShopCategory(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameShopCategory)] = EnumResponse{
		Name:               ts.EnumNameShopCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpShops},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initSphereColor() {
	enumDescription := "Defines all the colors the icon of a sphere can have in the menu. Additionally, these serve as the categories of spheres."

	typeSlice := []EnumVal{
		{
			Name:        string(database.SphereColorRed),
			Description: "Spheres that activate adjacent stat- and ability-nodes.",
		},
		{
			Name:        string(database.SphereColorYellow),
			Description: "Spheres that activate remote nodes that, in most cases, have been activated by another character.",
		},
		{
			Name:        string(database.SphereColorBlack),
			Description: "Key spheres that remove sphere locks.",
		},
		{
			Name:        string(database.SphereColorPurple),
			Description: "Spheres that transform empty nodes into stat-nodes.",
		},
		{
			Name:        string(database.SphereColorBlue),
			Description: "Reserved for the Clear Sphere, which transforms a stat-node into an empty node.",
		},
		{
			Name:        string(database.SphereColorWhite),
			Description: "Spheres that move the user to a remote place on the grid.",
		},
	}

	T.SphereColor = EnumType[database.SphereColor, any]{
		Name:     ts.EnumNameSphereColor,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.SphereColor { return database.SphereColor(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameSphereColor)] = EnumResponse{
		Name:               ts.EnumNameSphereColor,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpSpheres},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initSphereEffect() {
	enumDescription := "Defines the effect, a sphere has on the target node, if used successfully."

	typeSlice := []EnumVal{
		{
			Name: string(database.SphereEffectActivation),
		},
		{
			Name: string(database.SphereEffectRemoval),
		},
		{
			Name: string(database.SphereEffectCreation),
		},
		{
			Name: string(database.SphereEffectTeleportation),
		},
	}

	T.SphereEffect = EnumType[database.SphereEffect, any]{
		Name:     ts.EnumNameSphereEffect,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.SphereEffect { return database.SphereEffect(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameSphereEffect)] = EnumResponse{
		Name:               ts.EnumNameSphereEffect,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpSpheres},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initSphereGridType() {
	enumDescription := "The two types of sphere grid. Can only be chosen once, at the start of the game."

	typeSlice := []EnumVal{
		{
			Name: string(database.SphereGridTypeStandard),
		},
		{
			Name: string(database.SphereGridTypeExpert),
		},
	}

	T.SphereGridType = EnumType[database.SphereGridType, any]{
		Name:     ts.EnumNameSphereGridType,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.SphereGridType { return database.SphereGridType(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameSphereGridType)] = EnumResponse{
		Name:               ts.EnumNameSphereGridType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpCharacters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initStatusConditionCategory() {
	enumDescription := "Divides status-conditions into smaller categories."

	typeSlice := []EnumVal{
		{
			Name: string(database.StatusConditionCategoryNegative),
		},
		{
			Name: string(database.StatusConditionCategoryPositive),
		},
		{
			Name: string(database.StatusConditionCategoryOther),
		},
	}

	T.StatusConditionCategory = EnumType[database.StatusConditionCategory, any]{
		Name:     ts.EnumNameStatusConditionCategory,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.StatusConditionCategory { return database.StatusConditionCategory(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameStatusConditionCategory)] = EnumResponse{
		Name:               ts.EnumNameStatusConditionCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpStatusConditions},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initTreasureType() {
	enumDescription := "Defines the manner of obtaining a treasure."

	typeSlice := []EnumVal{
		{
			Name:        string(database.TreasureTypeChest),
			Description: "The treasure is found in a chest.",
		},
		{
			Name:        string(database.TreasureTypeGift),
			Description: "The treasure is a gift from an NPC, received by talking to them.",
		},
		{
			Name:        string(database.TreasureTypeObject),
			Description: "The treasure is found when interacting with an in-game object. Most of the time, the treasure is the object itself (Jecht Spheres, Al Bhed Primers), other times it's not (some Celestial Weapons).",
		},
	}

	T.TreasureType = EnumType[database.TreasureType, any]{
		Name:     ts.EnumNameTreasureType,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.TreasureType { return database.TreasureType(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameTreasureType)] = EnumResponse{
		Name:               ts.EnumNameTreasureType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpTreasures},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initWeaponType() {
	enumDescription := "The eight weapon types, each one being associated with a character."

	typeSlice := []EnumVal{
		{
			Name: string(database.WeaponTypeSword),
		},
		{
			Name: string(database.WeaponTypeStaff),
		},
		{
			Name: string(database.WeaponTypeBlitzball),
		},
		{
			Name: string(database.WeaponTypeDoll),
		},
		{
			Name: string(database.WeaponTypeSpear),
		},
		{
			Name: string(database.WeaponTypeBlade),
		},
		{
			Name: string(database.WeaponTypeClaw),
		},
		{
			Name: string(database.WeaponTypeSeymourStaff),
		},
	}

	T.WeaponType = EnumType[database.WeaponType, any]{
		Name:     ts.EnumNameWeaponType,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.WeaponType { return database.WeaponType(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameWeaponType)] = EnumResponse{
		Name:               ts.EnumNameWeaponType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpCharacters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initAccSourceType() {
	enumDescription := "Describes, how the accuracy of an action is calculated."

	typeSlice := []EnumVal{
		{
			Name:        string(database.AccSourceTypeAccuracy),
			Description: "The accuracy of the ability is calculated via the user's accuracy stat.",
		},
		{
			Name:        string(database.AccSourceTypeRate),
			Description: "The ability has its own accuracy.",
		},
	}

	T.AccSourceType = EnumType[database.AccSourceType, any]{
		Name:     ts.EnumNameAccSourceType,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.AccSourceType { return database.AccSourceType(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameAccSourceType)] = EnumResponse{
		Name:               ts.EnumNameAccSourceType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAbilities, ts.EpPlayerAbilities, ts.EpOverdriveAbilities, ts.EpItemAbilities, ts.EpTriggerCommands, ts.EpMiscAbilities, ts.EpEnemyAbilities, ts.EpAeons},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initAttackType() {
	enumDescription := "Defines the type of effect an action will have on the targets' HP/MP after damage calculation, in a normal situation (no elemental absorbtion, or 'Zombie' status)."

	typeSlice := []EnumVal{
		{
			Name:        string(database.AttackTypeAttack),
			Description: "The action deals damage.",
		},
		{
			Name:        string(database.AttackTypeHeal),
			Description: "The action restores HP/MP.",
		},
		{
			Name:        string(database.AttackTypeAbsorb),
			Description: "The action deals damage and restores the same amount to its user.",
		},
	}

	T.AttackType = EnumType[database.AttackType, any]{
		Name:     ts.EnumNameAttackType,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.AttackType { return database.AttackType(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameAttackType)] = EnumResponse{
		Name:               ts.EnumNameAttackType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAbilities, ts.EpPlayerAbilities, ts.EpOverdriveAbilities, ts.EpItemAbilities, ts.EpTriggerCommands, ts.EpMiscAbilities, ts.EpEnemyAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initBreakDmgLimitType() {
	enumDescription := "Defines, whether an ability breaks the damage limit by itself, or if it needs the auto-ability to be able to."

	typeSlice := []EnumVal{
		{
			Name:        string(database.BreakDmgLmtTypeAlways),
			Description: "The ability always breaks the damage limit.",
		},
		{
			Name:        string(database.BreakDmgLmtTypeAutoAbility),
			Description: "The ability can only break the damage limit, if the user has the auto-ability 'Break Damage Limit' equipped.",
		},
	}

	T.BreakDmgLimitType = EnumType[database.BreakDmgLmtType, database.NullBreakDmgLmtType]{
		Name:         ts.EnumNameBreakDmgLimitType,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.BreakDmgLmtType { return database.BreakDmgLmtType(s) },
		NullConvFunc: database.ToNullBreakDmgLmtType,
		GetNullEnum:  database.GetNullBreakDmgLmtType,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameBreakDmgLimitType)] = EnumResponse{
		Name:               ts.EnumNameBreakDmgLimitType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAbilities, ts.EpPlayerAbilities, ts.EpOverdriveAbilities, ts.EpItemAbilities, ts.EpTriggerCommands, ts.EpMiscAbilities, ts.EpEnemyAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initCalculationType() {
	enumDescription := "Determines, how the value of a stat/modifier change is applied to its target stat/modifier, when it is activated."

	typeSlice := []EnumVal{
		{
			Name:        string(database.CalculationTypeAddedPercentage),
			Description: "The given value is added (or subtracted, if negative) to a final percentage-based factor which is applied at the end of the calculation. Example: If the value is 3 (like with Auto-Ability 'Strength +3%'), then the result of the calculation will be multiplied by 1.03.",
		},
		{
			Name:        string(database.CalculationTypeAddedValue),
			Description: "The given value is added directly to the destination. This type is either used directly on stats or on factors within the calculation and is most prominently seen on abilities like 'Cheer' and its equivalents.",
		},
		{
			Name:        string(database.CalculationTypeMultiply),
			Description: "The result of the calculation will be multiplied by the given value. Values with calculation type 'multiply' can stack on the same destination. Example: If Rikku uses 'Hot Spurs' (overdrive-charge x1.5) and then 'Eccentrick' (overdrive-charge x2), the gauge will charge 3 times as fast.",
		},
		{
			Name:        string(database.CalculationTypeMultiplyHighest),
			Description: "The result of the calculation will be multiplied by the given value. If more than one modification with calculation type 'multiply-highest' reach the same destination, only the highest factor is applied. Example: Auto-Abilities 'Double AP' and 'Triple AP' both use 'multiply-highest'. Factor 3 of 'Triple AP' will override factor 2 of 'Double AP', since it's higher.",
		},
		{
			Name:        string(database.CalculationTypeSetValue),
			Description: "The destination becomes the given value. Example: Auto-Ability 'One MP Cost' sets the MP cost every spell to 1.",
		},
	}

	T.CalculationType = EnumType[database.CalculationType, any]{
		Name:     ts.EnumNameCalculationType,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.CalculationType { return database.CalculationType(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameCalculationType)] = EnumResponse{
		Name:               ts.EnumNameCalculationType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAbilities, ts.EpPlayerAbilities, ts.EpOverdriveAbilities, ts.EpItemAbilities, ts.EpTriggerCommands, ts.EpMiscAbilities, ts.EpEnemyAbilities, ts.EpAutoAbilities, ts.EpProperties, ts.EpStatusConditions},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initCriticalType() {
	enumDescription := "Determines, which formula an ability uses to calculate critical hits."

	typeSlice := []EnumVal{
		{
			Name:        string(database.CriticalTypeCrit),
			Description: "The ability can crit, but doesn't get any further bonusses.",
		},
		{
			Name:        string(database.CriticalTypeCritEquip),
			Description: "The ability can crit and the critical plus values of the user's equipment are added toward the critical hit chance.",
		},
		{
			Name:        string(database.CriticalTypeCritAbility),
			Description: "The ability can crit and the critical plus value of the used ability is added toward the critical hit chance.",
		},
	}

	T.CriticalType = EnumType[database.CriticalType, database.NullCriticalType]{
		Name:         ts.EnumNameCriticalType,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.CriticalType { return database.CriticalType(s) },
		NullConvFunc: database.ToNullCriticalType,
		GetNullEnum:  database.GetNullCriticalType,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameCriticalType)] = EnumResponse{
		Name:               ts.EnumNameCriticalType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAbilities, ts.EpPlayerAbilities, ts.EpOverdriveAbilities, ts.EpItemAbilities, ts.EpTriggerCommands, ts.EpMiscAbilities, ts.EpEnemyAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initDamageFormula() {
	enumDescription := "Determines the damage formula an action uses."

	typeSlice := []EnumVal{
		{
			Name:        string(database.DamageFormulaStrVsDef),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaStrIgnDef),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaMagVsMdf),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaMagIgnMdf),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaPercentageCurrent),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaPercentageMax),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaHealing),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaCtbCurrent),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaSpecialNoVar),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaSpecialVar),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaSpecialMagic),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaSpecialGil),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaSpecialKills),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaSpecial9999),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaFixed9999),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaUserMaxHp),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaSwallowedA),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaSwallowedB),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaClstlHpHigh),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaClstlHpLow),
			Description: "",
		},
		{
			Name:        string(database.DamageFormulaClstlMpHigh),
			Description: "",
		},
	}

	T.DamageFormula = EnumType[database.DamageFormula, any]{
		Name:     ts.EnumNameDamageFormula,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.DamageFormula { return database.DamageFormula(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameDamageFormula)] = EnumResponse{
		Name:               ts.EnumNameDamageFormula,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAbilities, ts.EpPlayerAbilities, ts.EpOverdriveAbilities, ts.EpItemAbilities, ts.EpTriggerCommands, ts.EpMiscAbilities, ts.EpEnemyAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initDamageType() {
	enumDescription := "Determines, which type of damage an ability deals, and in turn, which status-conditions, properties, abilities, and auto-abilities can influence it."

	typeSlice := []EnumVal{
		{
			Name:        string(database.DamageTypePhysical),
			Description: "The damage can be reduced by 'Protect', 'Defend' 'Power Break', 'Sentinel', 'Shield', and 'Cheer', as well as 'Defense +X%' Auto-Abilities. It can be increased by 'Strength +X%' Auto-Abilities.",
		},
		{
			Name:        string(database.DamageTypeMagical),
			Description: "The damage can be reduced by 'Shell', 'Magic Break', 'Shield', and 'Focus', as well as 'Magic Def +X%' Auto-Abilities. It can be increased by 'Magic +X%' Auto-Abilities.",
		},
		{
			Name:        string(database.DamageTypeSpecial),
			Description: "The damage can only be reduced by 'Shield'.",
		},
	}

	T.DamageType = EnumType[database.DamageType, any]{
		Name:     ts.EnumNameDamageType,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.DamageType { return database.DamageType(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameDamageType)] = EnumResponse{
		Name:               ts.EnumNameDamageType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAbilities, ts.EpPlayerAbilities, ts.EpOverdriveAbilities, ts.EpItemAbilities, ts.EpTriggerCommands, ts.EpMiscAbilities, ts.EpEnemyAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initDelayStrength() {
	enumDescription := "Defines the strength of tickspeed-based delay."

	typeSlice := []EnumVal{
		{
			Name:        string(database.DelayStrengthStrong),
			Description: "The target suffers delay equal to a regular turn (tickspeed x3).",
		},
		{
			Name:        string(database.DelayStrengthWeak),
			Description: "The target suffers delay equal to half a regular turn (tickspeed x1.5).",
		},
	}

	T.DelayStrength = EnumType[database.DelayStrength, any]{
		Name:     ts.EnumNameDelayStrength,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.DelayStrength { return database.DelayStrength(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameDelayStrength)] = EnumResponse{
		Name:               ts.EnumNameDelayStrength,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAbilities, ts.EpPlayerAbilities, ts.EpOverdriveAbilities, ts.EpItemAbilities, ts.EpTriggerCommands, ts.EpMiscAbilities, ts.EpEnemyAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initDurationType() {
	enumDescription := "Determines, how long an inflicted status-condition is active on the target unit."

	typeSlice := []EnumVal{
		{
			Name:        string(database.DurationTypeTurns),
			Description: "The status condition wears off after a set amount of turns.",
		},
		{
			Name:        string(database.DurationTypeInflictorNextTurn),
			Description: "The status condition wears off on the inflictor's next turn. This is only used for 'Threaten'.",
		},
		{
			Name:        string(database.DurationTypeBlocks),
			Description: "The status condition is present as long as it has blocks left. Used only for 'Nul-' status conditions.",
		},
		{
			Name:        string(database.DurationTypeEndless),
			Description: "The status condition won't wear off. It is present until it is removed.",
		},
		{
			Name:        string(database.DurationTypeInstant),
			Description: "The status condition wears off instantly. Most commonly seen on 'Death' and 'Life', but there are exceptions like Sinspawn Gui and Ultima Buster gaining 'Defend' while blocking, or Penance's Arms gaining 'Haste' while taking an action.",
		},
		{
			Name:        string(database.DurationTypeAuto),
			Description: "The status condition is present forever and can't be removed. Only used on Biran Ronso's 'Mighty Guard'.",
		},
	}

	T.DurationType = EnumType[database.DurationType, any]{
		Name:     ts.EnumNameDurationType,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.DurationType { return database.DurationType(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameDurationType)] = EnumResponse{
		Name:               ts.EnumNameDurationType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAbilities, ts.EpPlayerAbilities, ts.EpOverdriveAbilities, ts.EpItemAbilities, ts.EpTriggerCommands, ts.EpMiscAbilities, ts.EpEnemyAbilities, ts.EpAutoAbilities, ts.EpMonsters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initSpecialActionType() {
	enumDescription := "Defines special actions an ability can trigger."

	typeSlice := []EnumVal{
		{
			Name: string(database.SpecialActionTypeBribe),
		},
		{
			Name: string(database.SpecialActionTypeStealGil),
		},
		{
			Name: string(database.SpecialActionTypeStealItem),
		},
		{
			Name: string(database.SpecialActionTypeTransferOverdrive),
		},
		{
			Name: string(database.SpecialActionTypeCopycat),
		},
	}

	T.SpecialActionType = EnumType[database.SpecialActionType, database.NullSpecialActionType]{
		Name:         ts.EnumNameSpecialActionType,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.SpecialActionType { return database.SpecialActionType(s) },
		NullConvFunc: database.ToNullSpecialActionType,
		GetNullEnum:  database.GetNullSpecialActionType,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameSpecialActionType)] = EnumResponse{
		Name:               ts.EnumNameSpecialActionType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAbilities, ts.EpPlayerAbilities, ts.EpOverdriveAbilities, ts.EpItemAbilities, ts.EpTriggerCommands, ts.EpMiscAbilities, ts.EpEnemyAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initTargetType() {
	enumDescription := "Defines, which units an ability can target. This type is used for abilities, as well as the cursors that can be seen, when selecting a target. However, the two don't necessarily need to match (with Rikku's 'Mix' Overdrive being the prime example)."

	typeSlice := []EnumVal{
		{
			Name:        string(database.TargetTypeSelf),
			Description: "The action targets its user.",
		},
		{
			Name:        string(database.TargetTypeSelfAllEnemies),
			Description: "The action targets its user and all units of the user's opposing party.",
		},
		{
			Name:        string(database.TargetTypeSingleAlly),
			Description: "The action targets one unit of the user's party.",
		},
		{
			Name:        string(database.TargetTypeSingleEnemy),
			Description: "The action targets one unit of the user's opposing party.",
		},
		{
			Name:        string(database.TargetTypeSingleTarget),
			Description: "The action targets the selected unit.",
		},
		{
			Name:        string(database.TargetTypeRandomAlly),
			Description: "The action targets a random unit of the user's party.",
		},
		{
			Name:        string(database.TargetTypeRandomEnemy),
			Description: "The action targets a random unit of the user's opposing party.",
		},
		{
			Name:        string(database.TargetTypeAllAllies),
			Description: "The action targets all units of the user's party.",
		},
		{
			Name:        string(database.TargetTypeAllEnemies),
			Description: "The action targets all units of the user's opposing party.",
		},
		{
			Name:        string(database.TargetTypeTargetParty),
			Description: "The action targets all units of the selected party.",
		},
		{
			Name:        string(database.TargetTypeNTargets),
			Description: "The action targets N amount of units (N is stated via the ability's hit_amount). The action can also target KO'd characters and inanimate objects. Only Seymour's and Seymour Natus' multi-spells and Spectral Keeper's counter attack, as well as its glyph mine activation use this target type.",
		},
		{
			Name:        string(database.TargetTypeEveryone),
			Description: "The action targets every unit on the field.",
		},
		{
			Name:        string(database.TargetTypeEveryoneElse),
			Description: "The action targets every unit on the field except its user.",
		},
	}

	T.TargetType = EnumType[database.TargetType, database.NullTargetType]{
		Name:         ts.EnumNameTargetType,
		Lookup:       enumSliceToMap(typeSlice),
		ConvFunc:     func(s string) database.TargetType { return database.TargetType(s) },
		NullConvFunc: database.ToNullTargetType,
		GetNullEnum:  database.GetNullTargetType,
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameTargetType)] = EnumResponse{
		Name:               ts.EnumNameTargetType,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAbilities, ts.EpPlayerAbilities, ts.EpOverdriveAbilities, ts.EpItemAbilities, ts.EpTriggerCommands, ts.EpMiscAbilities, ts.EpEnemyAbilities, ts.EpAeonCommands, ts.EpOverdrives},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initBattleStart() {
	enumDescription := fmt.Sprintf("Determines  for the /%s endpoint.", ts.EpTurnOrder)

	typeSlice := []EnumVal{
		{
			Name:        string(database.BattleStartNormal),
			Description: "",
		},
		{
			Name:        string(database.BattleStartPreemptive),
			Description: "",
		},
		{
			Name:        string(database.BattleStartAmbush),
			Description: "",
		},
	}

	T.BattleStart = EnumType[database.BattleStart, any]{
		Name:     ts.EnumNameBattleStart,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.BattleStart { return database.BattleStart(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameBattleStart)] = EnumResponse{
		Name:               ts.EnumNameBattleStart,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpTurnOrder},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initHasteStatus() {
	enumDescription := fmt.Sprintf("Determines  for the /%s endpoint.", ts.EpTurnOrder)

	typeSlice := []EnumVal{
		{
			Name:        string(database.HasteStatusHaste),
			Description: "",
		},
		{
			Name:        string(database.HasteStatusSlow),
			Description: "",
		},
		{
			Name:        string(database.HasteStatusAutoHaste),
			Description: "",
		},
	}

	T.HasteStatus = EnumType[database.HasteStatus, any]{
		Name:     ts.EnumNameTranslationDirection,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.HasteStatus { return database.HasteStatus(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameHasteStatus)] = EnumResponse{
		Name:               ts.EnumNameHasteStatus,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpTurnOrder},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initTranslationDirection() {
	enumDescription := fmt.Sprintf("Determines the direction of translation for the /%s endpoint.", ts.EpAlBhed)

	typeSlice := []EnumVal{
		{
			Name:        string(database.TranslationDirectionToAlBhed),
			Description: "The given text is translated from English to Al Bhed.",
		},
		{
			Name:        string(database.TranslationDirectionToEnglish),
			Description: "The given text is translated from Al Bhed to English.",
		},
	}

	T.TranslationDirection = EnumType[database.TranslationDirection, any]{
		Name:     ts.EnumNameTranslationDirection,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.TranslationDirection { return database.TranslationDirection(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameTranslationDirection)] = EnumResponse{
		Name:               ts.EnumNameTranslationDirection,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpAlBhed},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (T *Enums) initTurnOrderRNG() {
	enumDescription := fmt.Sprintf("Determines the RNG for the /%s endpoint.", ts.EpTurnOrder)

	typeSlice := []EnumVal{
		{
			Name:        string(database.TurnOrderRngBest),
			Description: "",
		},
		{
			Name:        string(database.TurnOrderRngWorst),
			Description: "",
		},
		{
			Name:        string(database.TurnOrderRngMedian),
			Description: "",
		},
	}

	T.TurnOrderRNG = EnumType[database.TurnOrderRng, any]{
		Name:     ts.EnumNameTurnOrderRNG,
		Lookup:   enumSliceToMap(typeSlice),
		ConvFunc: func(s string) database.TurnOrderRng { return database.TurnOrderRng(s) },
	}

	T.Lookup[ts.GetEnumKey(ts.EnumNameTurnOrderRNG)] = EnumResponse{
		Name:               ts.EnumNameTurnOrderRNG,
		Description:        enumDescription,
		UsedByEndpointsInt: []ts.EndpointName{ts.EpTurnOrder},
		Values:             getEnumValIDs(typeSlice),
	}
}
