package api

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type EnumType[E, N any] struct {
	name         EnumName
	lookup       map[string]EnumVal
	convFunc     func(string) E
	nullConvFunc func(*string) N
	getNullEnum  func(*E) N
	aliasses     map[string][]E
}

type EnumResponse struct {
	Name               EnumName       `json:"name"`
	Description        string         `json:"description"`
	UsedByEndpointsInt []EndpointName `json:"-"`
	UsedByEndpoints    []string       `json:"used_by_endpoints"`
	Values             []EnumVal      `json:"values"`
}

func endpointsToURLs(cfg *Config, source []EndpointName) []string {
	urls := make([]string, 0, len(source))

	for _, ep := range source {
		url := createListURL(cfg, ep)
		urls = append(urls, url)
	}

	slices.Sort(urls)
	return urls
}

func (t *Enums) initAbilityType() {
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

	t.AbilityType = EnumType[database.AbilityType, any]{
		name:         enumNameAbilityType,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.AbilityType { return database.AbilityType(s) },
		nullConvFunc: nil,
		getNullEnum:  nil,
	}

	t.Lookup[getEnumKey(enumNameAbilityType)] = EnumResponse{
		Name:               enumNameAbilityType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAbilities, epPlayerAbilities, epOverdriveAbilities, epItemAbilities, epTriggerCommands, epMiscAbilities, epEnemyAbilities, epAeonCommands, epCharacterClasses, epTopmenus, epSubmenus, epMonsters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initUnitType() {
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

	t.UnitType = EnumType[database.UnitType, any]{
		name:         enumNameUnitType,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.UnitType { return database.UnitType(s) },
		nullConvFunc: nil,
		getNullEnum:  nil,
	}

	t.Lookup[getEnumKey(enumNameUnitType)] = EnumResponse{
		Name:               enumNameUnitType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAeons, epCharacterClasses, epCharacters, epPlayerUnits},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initItemType() {
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

	t.ItemType = EnumType[database.ItemType, any]{
		name:         enumNameItemType,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.ItemType { return database.ItemType(s) },
		nullConvFunc: nil,
		getNullEnum:  nil,
	}

	t.Lookup[getEnumKey(enumNameItemType)] = EnumResponse{
		Name:               enumNameItemType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epArenaCreations, epAutoAbilities, epBlitzballPrizes, epItems, epKeyItems, epAllItems, epMonsters, epQuests, epSidequests, epSubquests, epTreasures, epShops},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initQuestType() {
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

	t.QuestType = EnumType[database.QuestType, any]{
		name:         enumNameQuestType,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.QuestType { return database.QuestType(s) },
		nullConvFunc: nil,
		getNullEnum:  nil,
	}

	t.Lookup[getEnumKey(enumNameQuestType)] = EnumResponse{
		Name:               enumNameQuestType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epQuests, epSidequests, epSubquests},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initAaActivationCondition() {
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

	t.AaActivationCondition = EnumType[database.AaActivationCondition, any]{
		name:     enumNameAaActivationCondition,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.AaActivationCondition { return database.AaActivationCondition(s) },
	}

	t.Lookup[getEnumKey(enumNameAaActivationCondition)] = EnumResponse{
		Name:               enumNameAaActivationCondition,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAutoAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}
/*
func (t *Enums) initAlterationType() {
	//enumDescription := ""

	typeSlice := []EnumVal{
		{
			Name: string(database.AlterationTypeChange),
		},
		{
			Name: string(database.AlterationTypeGain),
		},
		{
			Name: string(database.AlterationTypeLoss),
		},
	}

	t.AlterationType = EnumType[database.AlterationType, any]{
		name:     enumNameAlterationType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.AlterationType { return database.AlterationType(s) },
	}


	t.Lookup[getEnumKey(enumNameAlterationType)] = EnumResponse{
		Name:               enumNameAlterationType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{},
		Values:             getEnumValIDs(typeSlice),
	}

}
*/
func (t *Enums) initAreaConnectionType() {
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

	t.AreaConnectionType = EnumType[database.AreaConnectionType, any]{
		name:     enumNameAreaConnectionType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.AreaConnectionType { return database.AreaConnectionType(s) },
	}

	t.Lookup[getEnumKey(enumNameAreaConnectionType)] = EnumResponse{
		Name:               enumNameAreaConnectionType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAreas},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initArenaCreationCategory() {
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

	t.ArenaCreationCategory = EnumType[database.MaCreationCategory, database.NullMaCreationCategory]{
		name:         enumNameArenaCreationCategory,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.MaCreationCategory { return database.MaCreationCategory(s) },
		nullConvFunc: database.ToNullMaCreationCategory,
		getNullEnum:  database.GetNullMaCreationCategory,
	}

	t.Lookup[getEnumKey(enumNameArenaCreationCategory)] = EnumResponse{
		Name:               enumNameArenaCreationCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epArenaCreations},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initArmorType() {
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

	t.ArmorType = EnumType[database.ArmorType, any]{
		name:     enumNameArmorType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.ArmorType { return database.ArmorType(s) },
	}

	t.Lookup[getEnumKey(enumNameArmorType)] = EnumResponse{
		Name:               enumNameArmorType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epCharacters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initArranger() {
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

	t.Arranger = EnumType[database.Arranger, database.NullArranger]{
		name:         enumNameArranger,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.Arranger { return database.Arranger(s) },
		nullConvFunc: database.ToNullArranger,
		getNullEnum:  database.GetNullArranger,
	}

	t.Lookup[getEnumKey(enumNameArranger)] = EnumResponse{
		Name:               enumNameArranger,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epSongs},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initAutoAbilityCategory() {
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

	t.AutoAbilityCategory = EnumType[database.AutoAbilityCategory, any]{
		name:     enumNameAutoAbilityCategory,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.AutoAbilityCategory { return database.AutoAbilityCategory(s) },
	}

	t.Lookup[getEnumKey(enumNameAutoAbilityCategory)] = EnumResponse{
		Name:               enumNameAutoAbilityCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAutoAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initAvailabilityType() {
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

	t.AvailabilityType = EnumType[database.AvailabilityType, any]{
		name:     enumNameAvailabilityType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.AvailabilityType { return database.AvailabilityType(s) },
		aliasses: map[string][]database.AvailabilityType{
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

	t.Lookup[getEnumKey(enumNameAvailabilityType)] = EnumResponse{
		Name:               enumNameAvailabilityType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAutoAbilities, epEquipment, epItems, epKeyItems, epLocations, epSublocations, epAreas, epAllItems, epMonsterFormations, epMonsters, epPrimers, epQuests, epSidequests, epSubquests, epShops, epSpheres, epTreasures},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initBgReplacementType() {
	enumDescription := "Describes the condition until which alternative background music replaces the regular background music of a particular set of areas."

	typeSlice := []EnumVal{
		{
			Name: string(database.BgReplacementTypeUntilTrigger),
			Description: "The background music continues playing until a certain story- or event-trigger occurs.",
		},
		{
			Name: string(database.BgReplacementTypeUntilZoneChange),
			Description: "The background music continues playing until the current area is left and a zone-change occurs.",
		},
	}

	t.BgReplacementType = EnumType[database.BgReplacementType, database.NullBgReplacementType]{
		name:         enumNameBgReplacementType,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.BgReplacementType { return database.BgReplacementType(s) },
		nullConvFunc: database.ToNullBgReplacementType,
		getNullEnum:  database.GetNullBgReplacementType,
	}

	t.Lookup[getEnumKey(enumNameBgReplacementType)] = EnumResponse{
		Name:               enumNameBgReplacementType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epSongs},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initBlitzballPositionSlot() {
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

	t.BlitzballPositionSlot = EnumType[database.BlitzballPositionSlot, any]{
		name:     enumNameBlitzballPositionSlot,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.BlitzballPositionSlot { return database.BlitzballPositionSlot(s) },
	}

	t.Lookup[getEnumKey(enumNameBlitzballPositionSlot)] = EnumResponse{
		Name:               enumNameBlitzballPositionSlot,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epBlitzballPrizes},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initBlitzballTournamentCategory() {
	enumDescription := "The two different categories of blitzball competitions."

	typeSlice := []EnumVal{
		{
			Name: string(database.BlitzballTournamentCategoryLeague),
		},
		{
			Name: string(database.BlitzballTournamentCategoryTournament),
		},
	}

	t.BlitzballTournamentCategory = EnumType[database.BlitzballTournamentCategory, any]{
		name:     enumNameBlitzballTournamentCategory,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.BlitzballTournamentCategory { return database.BlitzballTournamentCategory(s) },
	}

	t.Lookup[getEnumKey(enumNameBlitzballTournamentCategory)] = EnumResponse{
		Name:               enumNameBlitzballTournamentCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epBlitzballPrizes},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initCelestialFormula() {
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

	t.CelestialFormula = EnumType[database.CelestialFormula, any]{
		name:     enumNameCelestialFormula,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.CelestialFormula { return database.CelestialFormula(s) },
	}

	t.Lookup[getEnumKey(enumNameCelestialFormula)] = EnumResponse{
		Name:               enumNameCelestialFormula,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epCelestialWeapons},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initCharacterClassCategory() {
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

	t.CharacterClassCategory = EnumType[database.CharacterClassCategory, any]{
		name:     enumNameCharacterClassCategory,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.CharacterClassCategory { return database.CharacterClassCategory(s) },
	}

	t.Lookup[getEnumKey(enumNameCharacterClassCategory)] = EnumResponse{
		Name:               enumNameCharacterClassCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epCharacterClasses},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initComposer() {
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

	t.Composer = EnumType[database.Composer, database.NullComposer]{
		name:         enumNameComposer,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.Composer { return database.Composer(s) },
		nullConvFunc: database.ToNullComposer,
		getNullEnum:  database.GetNullComposer,
	}

	t.Lookup[getEnumKey(enumNameComposer)] = EnumResponse{
		Name:               enumNameComposer,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epSongs},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initCounterType() {
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

	t.CounterType = EnumType[database.CounterType, database.NullCounterType]{
		name:         enumNameCounterType,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.CounterType { return database.CounterType(s) },
		nullConvFunc: database.ToNullCounterType,
		getNullEnum:  database.GetNullCounterType,
	}

	t.Lookup[getEnumKey(enumNameCounterType)] = EnumResponse{
		Name:               enumNameCounterType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAutoAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initCTBIconType() {
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

	t.CTBIconType = EnumType[database.CtbIconType, any]{
		name:     enumNameCTBIconType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.CtbIconType { return database.CtbIconType(s) },
	}

	t.Lookup[getEnumKey(enumNameCTBIconType)] = EnumResponse{
		Name:               enumNameCTBIconType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epMonsters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initCreationArea() {
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

	t.CreationArea = EnumType[database.MaCreationArea, database.NullMaCreationArea]{
		name:         enumNameCreationArea,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.MaCreationArea { return database.MaCreationArea(s) },
		nullConvFunc: database.ToNullMaCreationArea,
		getNullEnum:  database.GetNullMaCreationArea,
	}

	t.Lookup[getEnumKey(enumNameCreationArea)] = EnumResponse{
		Name:               enumNameCreationArea,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epMonsters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initCreationsUnlockedCategory() {
	enumDescription := "Is used to determine, which of the other two conquests is required to have a certain amount of monsters to unlock an original creation in the monster arena."

	typeSlice := []EnumVal{
		{
			Name: string(database.CreationsUnlockedCategoryArea),
		},
		{
			Name: string(database.CreationsUnlockedCategorySpecies),
		},
	}

	t.CreationsUnlockedCategory = EnumType[database.CreationsUnlockedCategory, database.NullCreationsUnlockedCategory]{
		name:         enumNameCreationsUnlockedCategory,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.CreationsUnlockedCategory { return database.CreationsUnlockedCategory(s) },
		nullConvFunc: database.ToNullCreationsUnlockedCategory,
		getNullEnum:  database.GetNullCreationsUnlockedCategory,
	}

	t.Lookup[getEnumKey(enumNameCreationsUnlockedCategory)] = EnumResponse{
		Name:               enumNameCreationsUnlockedCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epArenaCreations},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initElementalAffinity() {
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

	t.ElementalAffinity = EnumType[database.ElementalAffinity, any]{
		name:     enumNameElementalAffinity,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.ElementalAffinity { return database.ElementalAffinity(s) },
	}

	t.Lookup[getEnumKey(enumNameElementalAffinity)] = EnumResponse{
		Name:               enumNameElementalAffinity,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAutoAbilities, epMonsters, epStatusConditions},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initEquipClass() {
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

	t.EquipClass = EnumType[database.EquipClass, any]{
		name:     enumNameEquipClass,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.EquipClass { return database.EquipClass(s) },
	}

	t.Lookup[getEnumKey(enumNameEquipClass)] = EnumResponse{
		Name:               enumNameEquipClass,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epEquipment, epEquipmentTables},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initEquipType() {
	enumDescription := "The two types of equipment."

	typeSlice := []EnumVal{
		{
			Name: string(database.EquipTypeWeapon),
		},
		{
			Name: string(database.EquipTypeArmor),
		},
	}

	t.EquipType = EnumType[database.EquipType, any]{
		name:     enumNameEquipType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.EquipType { return database.EquipType(s) },
	}

	t.Lookup[getEnumKey(enumNameEquipType)] = EnumResponse{
		Name:               enumNameEquipType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAutoAbilities, epEquipment, epEquipmentTables},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initItemCategory() {
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

	t.ItemCategory = EnumType[database.ItemCategory, any]{
		name:     enumNameItemCategory,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.ItemCategory { return database.ItemCategory(s) },
	}

	t.Lookup[getEnumKey(enumNameItemCategory)] = EnumResponse{
		Name:               enumNameItemCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epItems, epItemAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initItemUsability() {
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

	t.ItemUsability = EnumType[database.ItemUsability, any]{
		name:     enumNameItemUsability,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.ItemUsability { return database.ItemUsability(s) },
	}

	t.Lookup[getEnumKey(enumNameItemUsability)] = EnumResponse{
		Name:               enumNameItemUsability,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epItems},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initKeyItemCategory() {
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

	t.KeyItemCategory = EnumType[database.KeyItemCategory, any]{
		name:     enumNameKeyItemCategory,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.KeyItemCategory { return database.KeyItemCategory(s) },
	}

	t.Lookup[getEnumKey(enumNameKeyItemCategory)] = EnumResponse{
		Name:               enumNameKeyItemCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epKeyItems},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initLootType() {

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

	t.LootType = EnumType[database.LootType, any]{
		name:     enumNameLootType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.LootType { return database.LootType(s) },
	}

	t.Lookup[getEnumKey(enumNameLootType)] = EnumResponse{
		Name:               enumNameLootType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epTreasures},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initMixCategory() {
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

	t.MixCategory = EnumType[database.MixCategory, any]{
		name:     enumNameMixCategory,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.MixCategory { return database.MixCategory(s) },
	}

	t.Lookup[getEnumKey(enumNameMixCategory)] = EnumResponse{
		Name:               enumNameMixCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epMixes},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initModifierCategory() {
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

	t.ModifierCategory = EnumType[database.ModifierCategory, any]{
		name:     enumNameModifierCategory,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.ModifierCategory { return database.ModifierCategory(s) },
	}

	t.Lookup[getEnumKey(enumNameModifierCategory)] = EnumResponse{
		Name:               enumNameModifierCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epModifiers},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initMonsterFormationCategory() {
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

	t.MonsterFormationCategory = EnumType[database.MonsterFormationCategory, any]{
		name:     enumNameMonsterFormationCategory,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.MonsterFormationCategory { return database.MonsterFormationCategory(s) },
	}

	t.Lookup[getEnumKey(enumNameMonsterFormationCategory)] = EnumResponse{
		Name:               enumNameMonsterFormationCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epMonsterFormations},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initMonsterSpecies() {
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

	t.MonsterSpecies = EnumType[database.MonsterSpecies, any]{
		name:     enumNameMonsterSpecies,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.MonsterSpecies { return database.MonsterSpecies(s) },
	}

	t.Lookup[getEnumKey(enumNameMonsterSpecies)] = EnumResponse{
		Name:               enumNameMonsterSpecies,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epMonsters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initMonsterCategory() {
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

	t.MonsterCategory = EnumType[database.MonsterCategory, any]{
		name:     enumNameMonsterCategory,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.MonsterCategory { return database.MonsterCategory(s) },
	}

	t.Lookup[getEnumKey(enumNameMonsterCategory)] = EnumResponse{
		Name:               enumNameMonsterCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epMonsters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initMusicUseCase() {
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

	t.MusicUseCase = EnumType[database.MusicUseCase, database.NullMusicUseCase]{
		name:         enumNameMusicUseCase,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.MusicUseCase { return database.MusicUseCase(s) },
		nullConvFunc: database.ToNullMusicUseCase,
		getNullEnum:  database.GetNullMusicUseCase,
	}

	t.Lookup[getEnumKey(enumNameMusicUseCase)] = EnumResponse{
		Name:               enumNameMusicUseCase,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epSongs},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initNodePosition() {
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

	t.NodePosition = EnumType[database.NodePosition, any]{
		name:     enumNameNodePosition,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.NodePosition { return database.NodePosition(s) },
	}

	t.Lookup[getEnumKey(enumNameNodePosition)] = EnumResponse{
		Name:               enumNameNodePosition,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epSpheres},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initNodeState() {
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

	t.NodeState = EnumType[database.NodeState, database.NullNodeState]{
		name:         enumNameNodeState,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.NodeState { return database.NodeState(s) },
		nullConvFunc: database.ToNullNodeState,
		getNullEnum:  database.GetNullNodeState,
	}

	t.Lookup[getEnumKey(enumNameNodeState)] = EnumResponse{
		Name:               enumNameNodeState,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epSpheres},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initNodeType() {
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

	t.NodeType = EnumType[database.NodeType, any]{
		name:     enumNameNodeType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.NodeType { return database.NodeType(s) },
	}

	t.Lookup[getEnumKey(enumNameNodeType)] = EnumResponse{
		Name:               enumNameNodeType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epSpheres},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initNullifyArmored() {
	enumDescription := "In some cases, a monster's 'armored' property is nullified. This either happens, if it is the target of an attack by a user with 'piercing', or if suffers under the 'armor break' status (bearer)."

	typeSlice := []EnumVal{
		{
			Name: string(database.NullifyArmoredTarget),
		},
		{
			Name: string(database.NullifyArmoredBearer),
		},
	}

	t.NullifyArmored = EnumType[database.NullifyArmored, database.NullNullifyArmored]{
		name:         enumNameNullifyArmored,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.NullifyArmored { return database.NullifyArmored(s) },
		nullConvFunc: database.ToNullNullifyArmored,
		getNullEnum:  database.GetNullNullifyArmored,
	}

	t.Lookup[getEnumKey(enumNameNullifyArmored)] = EnumResponse{
		Name:               enumNameNullifyArmored,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epProperties, epStatusConditions},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initOverdriveModeType() {
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

	t.OverdriveModeType = EnumType[database.OverdriveModeType, any]{
		name:     enumNameOverdriveModeType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.OverdriveModeType { return database.OverdriveModeType(s) },
	}

	t.Lookup[getEnumKey(enumNameOverdriveModeType)] = EnumResponse{
		Name:               enumNameOverdriveModeType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epOverdriveModes},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initPlayerAbilityCategory() {
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

	t.PlayerAbilityCategory = EnumType[database.PlayerAbilityCategory, any]{
		name:     enumNamePlayerAbilityCategory,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.PlayerAbilityCategory { return database.PlayerAbilityCategory(s) },
	}

	t.Lookup[getEnumKey(enumNamePlayerAbilityCategory)] = EnumResponse{
		Name:               enumNamePlayerAbilityCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epPlayerAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initShopCategory() {
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

	t.ShopCategory = EnumType[database.ShopCategory, any]{
		name:     enumNameShopCategory,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.ShopCategory { return database.ShopCategory(s) },
	}

	t.Lookup[getEnumKey(enumNameShopCategory)] = EnumResponse{
		Name:               enumNameShopCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epShops},
		Values:             getEnumValIDs(typeSlice),
	}
}
/*
func (t *Enums) initShopType() {
	enumDescription := ""

	typeSlice := []EnumVal{
		{
			Name: string(database.ShopTypePreAirship),
		},
		{
			Name: string(database.ShopTypePostAirship),
		},
	}

	t.ShopType = EnumType[database.ShopType, database.NullShopType]{
		name:         enumNameShopType,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.ShopType { return database.ShopType(s) },
		nullConvFunc: database.ToNullShopType,
		getNullEnum:  database.GetNullShopType,
	}

	t.Lookup[getEnumKey(enumNameShopType)] = EnumResponse{
		Name:               enumNameShopType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{},
		Values:             getEnumValIDs(typeSlice),
	}
}
*/
func (t *Enums) initSphereColor() {
	enumDescription := "Defines all the colors the icon of a sphere can have in the menu. Additionally, these serve as the categories of spheres."

	typeSlice := []EnumVal{
		{
			Name: string(database.SphereColorRed),
			Description: "Spheres that activate adjacent stat- and ability-nodes.",
		},
		{
			Name: string(database.SphereColorYellow),
			Description: "Spheres that activate remote nodes that, in most cases, have been activated by another character.",
		},
		{
			Name: string(database.SphereColorBlack),
			Description: "Key spheres that remove sphere locks.",
		},
		{
			Name: string(database.SphereColorPurple),
			Description: "Spheres that transform empty nodes into stat-nodes.",
		},
		{
			Name: string(database.SphereColorBlue),
			Description: "Reserved for the Clear Sphere, which transforms a stat-node into an empty node.",
		},
		{
			Name: string(database.SphereColorWhite),
			Description: "Spheres that move the user to a remote place on the grid.",
		},
	}

	t.SphereColor = EnumType[database.SphereColor, any]{
		name:     enumNameSphereColor,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.SphereColor { return database.SphereColor(s) },
	}

	t.Lookup[getEnumKey(enumNameSphereColor)] = EnumResponse{
		Name:               enumNameSphereColor,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epSpheres},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initSphereEffect() {
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

	t.SphereEffect = EnumType[database.SphereEffect, any]{
		name:     enumNameSphereEffect,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.SphereEffect { return database.SphereEffect(s) },
	}

	t.Lookup[getEnumKey(enumNameSphereEffect)] = EnumResponse{
		Name:               enumNameSphereEffect,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epSpheres},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initSphereGridType() {
	enumDescription := "The two types of sphere grid. Can only be chosen once, at the start of the game."

	typeSlice := []EnumVal{
		{
			Name: string(database.SphereGridTypeStandard),
		},
		{
			Name: string(database.SphereGridTypeExpert),
		},
	}

	t.SphereGridType = EnumType[database.SphereGridType, any]{
		name:     enumNameSphereGridType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.SphereGridType { return database.SphereGridType(s) },
	}

	t.Lookup[getEnumKey(enumNameSphereGridType)] = EnumResponse{
		Name:               enumNameSphereGridType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epCharacters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initStatusConditionCategory() {
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

	t.StatusConditionCategory = EnumType[database.StatusConditionCategory, any]{
		name:     enumNameStatusConditionCategory,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.StatusConditionCategory { return database.StatusConditionCategory(s) },
	}

	t.Lookup[getEnumKey(enumNameStatusConditionCategory)] = EnumResponse{
		Name:               enumNameStatusConditionCategory,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epStatusConditions},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initTreasureType() {
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

	t.TreasureType = EnumType[database.TreasureType, any]{
		name:     enumNameTreasureType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.TreasureType { return database.TreasureType(s) },
	}

	t.Lookup[getEnumKey(enumNameTreasureType)] = EnumResponse{
		Name:               enumNameTreasureType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epTreasures},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initWeaponType() {
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

	t.WeaponType = EnumType[database.WeaponType, any]{
		name:     enumNameWeaponType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.WeaponType { return database.WeaponType(s) },
	}

	t.Lookup[getEnumKey(enumNameWeaponType)] = EnumResponse{
		Name:               enumNameWeaponType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epCharacters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initAccSourceType() {
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

	t.AccSourceType = EnumType[database.AccSourceType, any]{
		name:     enumNameAccSourceType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.AccSourceType { return database.AccSourceType(s) },
	}

	t.Lookup[getEnumKey(enumNameAccSourceType)] = EnumResponse{
		Name:               enumNameAccSourceType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAbilities, epPlayerAbilities, epOverdriveAbilities, epItemAbilities, epTriggerCommands, epMiscAbilities, epEnemyAbilities, epAeons},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initAttackType() {
	enumDescription := "Defines the type of effect an action will have on the targets' HP/MP after damage calculation, in a normal situation (no elemental absorbtion, or 'Zombie' status)."

	typeSlice := []EnumVal{
		{
			Name: string(database.AttackTypeAttack),
			Description: "The action deals damage.",
		},
		{
			Name: string(database.AttackTypeHeal),
			Description: "The action restores HP/MP.",
		},
		{
			Name: string(database.AttackTypeAbsorb),
			Description: "The action deals damage and restores the same amount to its user.",
		},
	}

	t.AttackType = EnumType[database.AttackType, any]{
		name:     enumNameAttackType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.AttackType { return database.AttackType(s) },
	}

	t.Lookup[getEnumKey(enumNameAttackType)] = EnumResponse{
		Name:               enumNameAttackType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAbilities, epPlayerAbilities, epOverdriveAbilities, epItemAbilities, epTriggerCommands, epMiscAbilities, epEnemyAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initBreakDmgLimitType() {
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

	t.BreakDmgLimitType = EnumType[database.BreakDmgLmtType, database.NullBreakDmgLmtType]{
		name:         enumNameBreakDmgLimitType,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.BreakDmgLmtType { return database.BreakDmgLmtType(s) },
		nullConvFunc: database.ToNullBreakDmgLmtType,
		getNullEnum:  database.GetNullBreakDmgLmtType,
	}

	t.Lookup[getEnumKey(enumNameBreakDmgLimitType)] = EnumResponse{
		Name:               enumNameBreakDmgLimitType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAbilities, epPlayerAbilities, epOverdriveAbilities, epItemAbilities, epTriggerCommands, epMiscAbilities, epEnemyAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initCalculationType() {
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

	t.CalculationType = EnumType[database.CalculationType, any]{
		name:     enumNameCalculationType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.CalculationType { return database.CalculationType(s) },
	}

	t.Lookup[getEnumKey(enumNameCalculationType)] = EnumResponse{
		Name:               enumNameCalculationType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAbilities, epPlayerAbilities, epOverdriveAbilities, epItemAbilities, epTriggerCommands, epMiscAbilities, epEnemyAbilities, epAutoAbilities, epProperties, epStatusConditions},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initCriticalType() {
	enumDescription := "Determines, which formula an ability uses to calculate critical hits."

	typeSlice := []EnumVal{
		{
			Name:        string(database.CriticalTypeCrit),
			Description: "The ability uses the normal critical hit formula.",
		},
		{
			Name:        string(database.CriticalTypeCritweapon),
			Description: "The critical plus values of the user's equipment are added toward the critical hit chance.",
		},
		{
			Name:        string(database.CriticalTypeCritability),
			Description: "The critical plus value of the used ability is added toward the critical hit chance.",
		},
	}

	t.CriticalType = EnumType[database.CriticalType, database.NullCriticalType]{
		name:         enumNameCriticalType,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.CriticalType { return database.CriticalType(s) },
		nullConvFunc: database.ToNullCriticalType,
		getNullEnum:  database.GetNullCriticalType,
	}

	t.Lookup[getEnumKey(enumNameCriticalType)] = EnumResponse{
		Name:               enumNameCriticalType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAbilities, epPlayerAbilities, epOverdriveAbilities, epItemAbilities, epTriggerCommands, epMiscAbilities, epEnemyAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initCtbAttackType() {
	enumDescription := "Defines the type of effect an action will have on the targets' CTB. A negative effect is better known as 'Delay'."

	typeSlice := []EnumVal{
		{
			Name:        string(database.CtbAttackTypeAttack),
			Description: "The action inflicts delay and makes the target's next turn come later.",
		},
		{
			Name:        string(database.CtbAttackTypeHeal),
			Description: "The action makes the target's next turn come earlier.",
		},
	}

	t.CtbAttackType = EnumType[database.CtbAttackType, any]{
		name:     enumNameCtbAttackType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.CtbAttackType { return database.CtbAttackType(s) },
	}

	t.Lookup[getEnumKey(enumNameCtbAttackType)] = EnumResponse{
		Name:               enumNameCtbAttackType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAbilities, epPlayerAbilities, epOverdriveAbilities, epItemAbilities, epTriggerCommands, epMiscAbilities, epEnemyAbilities, epStatusConditions},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initDamageFormula() {
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
	}

	t.DamageFormula = EnumType[database.DamageFormula, any]{
		name:     enumNameDamageFormula,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.DamageFormula { return database.DamageFormula(s) },
	}

	t.Lookup[getEnumKey(enumNameDamageFormula)] = EnumResponse{
		Name:               enumNameDamageFormula,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAbilities, epPlayerAbilities, epOverdriveAbilities, epItemAbilities, epTriggerCommands, epMiscAbilities, epEnemyAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initDamageType() {
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

	t.DamageType = EnumType[database.DamageType, any]{
		name:     enumNameDamageType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.DamageType { return database.DamageType(s) },
	}

	t.Lookup[getEnumKey(enumNameDamageType)] = EnumResponse{
		Name:               enumNameDamageType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAbilities, epPlayerAbilities, epOverdriveAbilities, epItemAbilities, epTriggerCommands, epMiscAbilities, epEnemyAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initDelayType() {
	enumDescription := "Determines the type of delay an action deals."

	typeSlice := []EnumVal{
		{
			Name:        string(database.DelayTypeCtbBased),
			Description: "Delay is based on current ticks. CTB damage/heal is only applied, if 'Slow'/'Haste' is succcessful or if the status was successfully removed.",
		},
		{
			Name:        string(database.DelayTypeTickSpeedBased),
			Description: "Delay is based on tick speed. CTB damage is applied via an attack. Example: 'Delay Attack'.",
		},
	}

	t.DelayType = EnumType[database.DelayType, any]{
		name:     enumNameDelayType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.DelayType { return database.DelayType(s) },
	}

	t.Lookup[getEnumKey(enumNameDelayType)] = EnumResponse{
		Name:               enumNameDelayType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAbilities, epPlayerAbilities, epOverdriveAbilities, epItemAbilities, epTriggerCommands, epMiscAbilities, epEnemyAbilities, epStatusConditions},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initDurationType() {
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

	t.DurationType = EnumType[database.DurationType, any]{
		name:     enumNameDurationType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.DurationType { return database.DurationType(s) },
	}

	t.Lookup[getEnumKey(enumNameDurationType)] = EnumResponse{
		Name:               enumNameDurationType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAbilities, epPlayerAbilities, epOverdriveAbilities, epItemAbilities, epTriggerCommands, epMiscAbilities, epEnemyAbilities, epAutoAbilities, epMonsters},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initSpecialActionType() {
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

	t.SpecialActionType = EnumType[database.SpecialActionType, database.NullSpecialActionType]{
		name:         enumNameSpecialActionType,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.SpecialActionType { return database.SpecialActionType(s) },
		nullConvFunc: database.ToNullSpecialActionType,
		getNullEnum:  database.GetNullSpecialActionType,
	}

	t.Lookup[getEnumKey(enumNameSpecialActionType)] = EnumResponse{
		Name:               enumNameSpecialActionType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAbilities, epPlayerAbilities, epOverdriveAbilities, epItemAbilities, epTriggerCommands, epMiscAbilities, epEnemyAbilities},
		Values:             getEnumValIDs(typeSlice),
	}
}

func (t *Enums) initTargetType() {
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

	t.TargetType = EnumType[database.TargetType, database.NullTargetType]{
		name:         enumNameTargetType,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.TargetType { return database.TargetType(s) },
		nullConvFunc: database.ToNullTargetType,
		getNullEnum:  database.GetNullTargetType,
	}

	t.Lookup[getEnumKey(enumNameTargetType)] = EnumResponse{
		Name:               enumNameTargetType,
		Description:        enumDescription,
		UsedByEndpointsInt: []EndpointName{epAbilities, epPlayerAbilities, epOverdriveAbilities, epItemAbilities, epTriggerCommands, epMiscAbilities, epEnemyAbilities, epAeonCommands, epOverdrives},
		Values:             getEnumValIDs(typeSlice),
	}
}
