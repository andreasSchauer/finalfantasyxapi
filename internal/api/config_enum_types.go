package api

import (
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
	Name				EnumName		`json:"name"`
	Description			string			`json:"description"`
	UsedByEndpointsInt	[]EndpointName	`json:"-"`
	UsedByEndpoints		[]string		`json:"used_by_endpoints"`
	Values				[]EnumVal		`json:"values"`
}

func endpointsToURLs(cfg *Config, source []EndpointName) []string {
	urls := make([]string, 0, len(source))

	for _, ep := range source {
		url := createListURL(cfg, ep)
		urls = append(urls, url)
	}

	return urls
}


func (t *Types) initAbilityType() {
	enumDescription := ""
	
	typeSlice := []EnumVal{
		{
			Name:        string(database.AbilityTypePlayerAbility),
			Description: "Abilities that can either be learned via the sphere grid or are learned by aeons.",
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
		Name: 				enumNameAbilityType,
		Description: 		enumDescription,
		UsedByEndpointsInt: 	[]EndpointName{},
		Values: 			getEnumValIDs(typeSlice),
	}
}

func (t *Types) initUnitType() {
	enumDescription := ""
	
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
		Name: 			enumNameUnitType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initItemType() {
	enumDescription := ""
	
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
		Name: 			enumNameItemType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initQuestType() {
	enumDescription := ""
	
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
		Name: 			enumNameQuestType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initAaActivationCondition() {
	enumDescription := ""
	
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
		Name: 			enumNameAaActivationCondition,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initAlterationType() {
	enumDescription := ""
	
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
		Name: 			enumNameAlterationType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initAreaConnectionType() {
	enumDescription := ""
	
	typeSlice := []EnumVal{
		{
			Name:        string(database.AreaConnectionTypeBothDirections),
			Description: "The edges of two areas are directly connected with each other, and you can freely zone between those areas.",
		},
		{
			Name:        string(database.AreaConnectionTypeOneDirection),
			Description: "The edges of two areas are directly connected with each other, but you can only zone from area A to area B, and not vice versa.",
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
		Name: 			enumNameAreaConnectionType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initArenaCreationCategory() {
	enumDescription := ""
	
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
		Name: 			enumNameArenaCreationCategory,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initArmorType() {
	enumDescription := ""
	
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
		Name: 			enumNameArmorType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initArranger() {
	enumDescription := ""
	
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
		Name: 			enumNameArranger,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initAutoAbilityCategory() {
	enumDescription := ""
	
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
		Name: 			enumNameAutoAbilityCategory,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initAvailabilityType() {
	enumDescription := ""
	
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
			Description: "The resource is only available during the events of the story that happen after acquiring the airship. These resources are either post-airship story bosses or resources that are past the point of no return.",
		},
		{
			Name:        string(database.AvailabilityTypePostGame),
			Description: "The resource is available in the post-game, meaning it either was already available before acquiring the airship, or it becomes available after acquiring the airship. This excludes story-specific resources. This value is essentially a combination of 'always' and 'post'.",
		},
		{
			Name:        string(database.AvailabilityTypeStory),
			Description: "The resource is only available during the events of the story and thus is missable. This value is essentially a combination of 'pre-story' and 'post-story'.",
		},
		{
			Name:        string(database.AvailabilityTypePreAirship),
			Description: "The resource is available before acquiring the airship. This value is essentially a combination of 'always' and 'pre-story'.",
		},
		{
			Name:        string(database.AvailabilityTypePostAirship),
			Description: "The resource is only available after acquiring the airship. This value is essentially a combination of 'post' and 'post-story'.",
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
		Name: 			enumNameAvailabilityType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initBgReplacementType() {
	enumDescription := ""
	
	typeSlice := []EnumVal{
		{
			Name: string(database.BgReplacementTypeUntilTrigger),
		},
		{
			Name: string(database.BgReplacementTypeUntilZoneChange),
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
		Name: 			enumNameBgReplacementType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initBlitzballPositionSlot() {
	enumDescription := ""
	
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
		Name: 			enumNameBlitzballPositionSlot,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initBlitzballTournamentCategory() {
	enumDescription := ""
	
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
		Name: 			enumNameBlitzballTournamentCategory,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initCelestialFormula() {
	enumDescription := ""
	
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
		Name: 			enumNameCelestialFormula,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initCharacterClassCategory() {
	enumDescription := ""
	
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
		Name: 			enumNameCharacterClassCategory,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initComposer() {
	enumDescription := ""
	
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
		Name: 			enumNameComposer,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initCounterType() {
	enumDescription := ""
	
	typeSlice := []EnumVal{
		{
			Name:        string(database.CounterTypePhysical),
			Description: "The user counters when being hit by a physical attack.",
		},
		{
			Name:        string(database.CounterTypeMagical),
			Description: "The user counters when being hit by a magical attack.",
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
		Name: 			enumNameCounterType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initCTBIconType() {
	enumDescription := ""
	
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
		Name: 			enumNameCTBIconType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initCreationArea() {
	enumDescription := ""
	
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
		Name: 			enumNameCreationArea,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initCreationsUnlockedCategory() {
	enumDescription := ""
	
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
		Name: 			enumNameCreationsUnlockedCategory,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initElementalAffinity() {
	enumDescription := ""
	
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
		Name: 			enumNameElementalAffinity,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initEquipClass() {
	enumDescription := ""
	
	typeSlice := []EnumVal{
		{
			Name:        string(database.EquipClassStandard),
			Description: "A standard, customizable equipment piece.",
		},
		{
			Name:        string(database.EquipClassUnique),
			Description: "The equipment piece is one of a kind and its auto-abilities can only be modified by progressing through the story.",
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
		Name: 			enumNameEquipClass,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initEquipType() {
	enumDescription := ""
	
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
		Name: 			enumNameEquipType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initItemCategory() {
	enumDescription := ""
	
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
		Name: 			enumNameItemCategory,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initKeyItemCategory() {
	enumDescription := ""
	
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
		Name: 			enumNameKeyItemCategory,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initLootType() {

	enumDescription := ""
	
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
		Name: 			enumNameLootType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initMixCategory() {
	enumDescription := ""
	
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
		Name: 			enumNameMixCategory,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initModifierCategory() {
	enumDescription := ""
	
	typeSlice := []EnumVal{
		{
			Name:        string(database.ModifierCategoryFixedValue),
			Description: "",
		},
		{
			Name:        string(database.ModifierCategoryDynamicValue),
			Description: "",
		},
		{
			Name:        string(database.ModifierCategoryFactor),
			Description: "",
		},
		{
			Name:        string(database.ModifierCategoryPercentage),
			Description: "",
		},
	}

	t.ModifierCategory = EnumType[database.ModifierCategory, any]{
		name:     enumNameModifierCategory,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.ModifierCategory { return database.ModifierCategory(s) },
	}

	t.Lookup[getEnumKey(enumNameModifierCategory)] = EnumResponse{
		Name: 			enumNameModifierCategory,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initMonsterFormationCategory() {
	enumDescription := ""
	
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
		Name: 			enumNameMonsterFormationCategory,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initMonsterSpecies() {
	enumDescription := ""
	
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
		Name: 			enumNameMonsterSpecies,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initMonsterCategory() {
	enumDescription := ""
	
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
		Name: 			enumNameMonsterCategory,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initMusicUseCase() {
	enumDescription := ""
	
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
		Name: 			enumNameMusicUseCase,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initNodePosition() {
	enumDescription := ""
	
	typeSlice := []EnumVal{
		{
			Name:        string(database.NodePositionNeighboring),
			Description: "The sphere can target neighboring nodes, or the node the selected character is currently positioned.",
		},
		{
			Name:        string(database.NodePositionAllyPosition),
			Description: "The sphere can only target nodes, another character is currently positioned.",
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
		Name: 			enumNameNodePosition,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initNodeState() {
	enumDescription := ""
	
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
		Name: 			enumNameNodeState,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initNodeType() {
	enumDescription := ""
	
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
		Name: 			enumNameNodeType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initNullifyArmored() {
	enumDescription := ""
	
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
		Name: 			enumNameNullifyArmored,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initOverdriveModeType() {
	enumDescription := ""
	
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
		Name: 			enumNameOverdriveModeType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initPlayerAbilityCategory() {
	enumDescription := ""
	
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
		Name: 			enumNamePlayerAbilityCategory,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initShopCategory() {
	enumDescription := ""
	
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
		Name: 			enumNameShopCategory,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initShopType() {
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
		Name: 			enumNameShopType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initSpecialActionType() {
	enumDescription := ""
	
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
	}

	t.SpecialActionType = EnumType[database.SpecialActionType, database.NullSpecialActionType]{
		name:         enumNameSpecialActionType,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.SpecialActionType { return database.SpecialActionType(s) },
		nullConvFunc: database.ToNullSpecialActionType,
		getNullEnum:  database.GetNullSpecialActionType,
	}

	t.Lookup[getEnumKey(enumNameSpecialActionType)] = EnumResponse{
		Name: 			enumNameSpecialActionType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initSphereColor() {
	enumDescription := ""
	
	typeSlice := []EnumVal{
		{
			Name: string(database.SphereColorRed),
		},
		{
			Name: string(database.SphereColorYellow),
		},
		{
			Name: string(database.SphereColorBlack),
		},
		{
			Name: string(database.SphereColorPurple),
		},
		{
			Name: string(database.SphereColorBlue),
		},
		{
			Name: string(database.SphereColorWhite),
		},
	}

	t.SphereColor = EnumType[database.SphereColor, any]{
		name:     enumNameSphereColor,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.SphereColor { return database.SphereColor(s) },
	}

	t.Lookup[getEnumKey(enumNameSphereColor)] = EnumResponse{
		Name: 			enumNameSphereColor,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initSphereEffect() {
	enumDescription := ""
	
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
		Name: 			enumNameSphereEffect,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initSphereGridType() {
	enumDescription := ""
	
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
		Name: 			enumNameSphereGridType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initStatusConditionCategory() {
	enumDescription := ""
	
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
		Name: 			enumNameStatusConditionCategory,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initTreasureType() {
	enumDescription := ""
	
	typeSlice := []EnumVal{
		{
			Name:        string(database.TreasureTypeChest),
			Description: "The treasure is found in a chest.",
		},
		{
			Name:        string(database.TreasureTypeGift),
			Description: "The treasure is a gift from an NPC.",
		},
		{
			Name:        string(database.TreasureTypeObject),
			Description: "The treasure is found by interacting with an in-game object. Most of the time, the treasure is the object itself (Jecht Spheres, Al Bhed Primers), other times it's not.",
		},
	}

	t.TreasureType = EnumType[database.TreasureType, any]{
		name:     enumNameTreasureType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.TreasureType { return database.TreasureType(s) },
	}

	t.Lookup[getEnumKey(enumNameTreasureType)] = EnumResponse{
		Name: 			enumNameTreasureType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initWeaponType() {
	enumDescription := ""
	
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
		Name: 			enumNameWeaponType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initAccSourceType() {
	enumDescription := ""
	
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
		Name: 			enumNameAccSourceType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initAttackType() {
	enumDescription := ""
	
	typeSlice := []EnumVal{
		{
			Name: string(database.AttackTypeAttack),
		},
		{
			Name: string(database.AttackTypeHeal),
		},
		{
			Name: string(database.AttackTypeAbsorb),
		},
	}

	t.AttackType = EnumType[database.AttackType, any]{
		name:     enumNameAttackType,
		lookup:   enumSliceToMap(typeSlice),
		convFunc: func(s string) database.AttackType { return database.AttackType(s) },
	}

	t.Lookup[getEnumKey(enumNameAttackType)] = EnumResponse{
		Name: 			enumNameAttackType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initBreakDmgLimitType() {
	enumDescription := ""
	
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
		Name: 			enumNameBreakDmgLimitType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initCalculationType() {
	enumDescription := ""
	
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
		Name: 			enumNameCalculationType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initCriticalType() {
	enumDescription := ""
	
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
		Name: 			enumNameCriticalType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initCtbAttackType() {
	enumDescription := ""
	
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
		Name: 			enumNameCtbAttackType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initDamageFormula() {
	enumDescription := ""
	
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
		Name: 			enumNameDamageFormula,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initDamageType() {
	enumDescription := ""
	
	typeSlice := []EnumVal{
		{
			Name:        string(database.DamageTypePhysical),
			Description: "The damage can be reduced by 'Protect', 'Defend' 'Power Break', 'Sentinel', 'Shield', and 'Cheer', as well as 'Defense +X%' Auto-Abilities.",
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
		Name: 			enumNameDamageType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initDelayType() {
	enumDescription := ""
	
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
		Name: 			enumNameDelayType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initDurationType() {
	enumDescription := ""
	
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
		Name: 			enumNameDurationType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}

func (t *Types) initTargetType() {
	enumDescription := ""
	
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
		Name: 			enumNameTargetType,
		Description: 	enumDescription,
		UsedByEndpointsInt: 		[]EndpointName{},
		Values: 		getEnumValIDs(typeSlice),
	}
}
