package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


type EnumType[E, N any] struct {
	name         string
	isEndpoint   bool
	lookup       map[string]EnumAPIResource
	convFunc     func(string) E
	nullConvFunc func(*string) N
	getNullEnum  func(*E) N
	aliasses	 map[string][]E
}



// TypeLookup holds all the enum types for the application that are either used as endpoint or query param
type TypeLookup struct {
	AbilityType                 EnumType[database.AbilityType, any]
	UnitType                    EnumType[database.UnitType, any]
	ItemType                    EnumType[database.ItemType, any]
	QuestType                   EnumType[database.QuestType, any]

	AaActivationCondition		EnumType[database.AaActivationCondition, any]
	AreaConnectionType          EnumType[database.AreaConnectionType, any]
	ArenaCreationCategory       EnumType[database.MaCreationCategory, database.NullMaCreationCategory]
	Arranger                    EnumType[database.Arranger, database.NullArranger]
	AutoAbilityCategory			EnumType[database.AutoAbilityCategory, any]
	AvailabilityType            EnumType[database.AvailabilityType, any]
	BlitzballTournamentCategory EnumType[database.BlitzballTournamentCategory, any]
	CelestialFormula			EnumType[database.CelestialFormula, any]
	CharacterClassCategory      EnumType[database.CharacterClassCategory, any]
	CounterType					EnumType[database.CounterType, database.NullCounterType]
	Composer                    EnumType[database.Composer, database.NullComposer]
	CreationArea                EnumType[database.MaCreationArea, database.NullMaCreationArea]
	CTBIconType                 EnumType[database.CtbIconType, any]
	EquipClass					EnumType[database.EquipClass, any]
	EquipType					EnumType[database.EquipType, any]
	ItemCategory                EnumType[database.ItemCategory, any]
	KeyItemCategory				EnumType[database.KeyItemCategory, any]
	LootType                    EnumType[database.LootType, any]
	MixCategory					EnumType[database.MixCategory, any]
	MonsterCategory             EnumType[database.MonsterCategory, any]
	MonsterFormationCategory    EnumType[database.MonsterFormationCategory, any]
	MonsterSpecies              EnumType[database.MonsterSpecies, any]
	NodePosition				EnumType[database.NodePosition, any]
	NodeState					EnumType[database.NodeState, database.NullNodeState]
	OverdriveModeType           EnumType[database.OverdriveModeType, any]
	PlayerAbilityCategory       EnumType[database.PlayerAbilityCategory, any]
	ShopCategory                EnumType[database.ShopCategory, any]
	ShopType                    EnumType[database.ShopType, database.NullShopType]
	SphereColor					EnumType[database.SphereColor, any]
	TreasureType                EnumType[database.TreasureType, any]

	AccSourceType     EnumType[database.AccSourceType, any]
	AttackType        EnumType[database.AttackType, any]
	BreakDmgLimitType EnumType[database.BreakDmgLmtType, database.NullBreakDmgLmtType]
	CalculationType   EnumType[database.CalculationType, any]
	CriticalType      EnumType[database.CriticalType, database.NullCriticalType]
	CtbAttackType     EnumType[database.CtbAttackType, any]
	DamageFormula     EnumType[database.DamageFormula, any]
	DamageType        EnumType[database.DamageType, any]
	DelayType         EnumType[database.DelayType, any]
	DurationType      EnumType[database.DurationType, any]
	TargetType        EnumType[database.TargetType, database.NullTargetType]
}

func (cfg *Config) TypeLookupInit() {
	cfg.t = &TypeLookup{}

	cfg.t.initAbilityType()
	cfg.t.initUnitType()
	cfg.t.initItemType()
	cfg.t.initQuestType()

	cfg.t.initAaActivationCondition()
	cfg.t.initAreaConnectionType()
	cfg.t.initArenaCreationCategory()
	cfg.t.initArranger()
	cfg.t.initAutoAbilityCategory()
	cfg.t.initAvailabilityType()
	cfg.t.initBlitzballTournamentCategory()
	cfg.t.initCelestialFormula()
	cfg.t.initCharacterClassCategory()
	cfg.t.initComposer()
	cfg.t.initCounterType()
	cfg.t.initCTBIconType()
	cfg.t.initCreationArea()
	cfg.t.initEquipClass()
	cfg.t.initEquipType()
	cfg.t.initItemCategory()
	cfg.t.initKeyItemCategory()
	cfg.t.initLootType()
	cfg.t.initMixCategory()
	cfg.t.initMonsterCategory()
	cfg.t.initMonsterFormationCategory()
	cfg.t.initMonsterSpecies()
	cfg.t.initNodePosition()
	cfg.t.initNodeState()
	cfg.t.initOverdriveModeType()
	cfg.t.initPlayerAbilityCategory()
	cfg.t.initShopCategory()
	cfg.t.initShopType()
	cfg.t.initSphereColor()
	cfg.t.initTreasureType()

	cfg.t.initAccSourceType()
	cfg.t.initAttackType()
	cfg.t.initBreakDmgLimitType()
	cfg.t.initCalculationType()
	cfg.t.initCriticalType()
	cfg.t.initCtbAttackType()
	cfg.t.initDamageFormula()
	cfg.t.initDamageType()
	cfg.t.initDelayType()
	cfg.t.initDurationType()
	cfg.t.initTargetType()
}



func (t *TypeLookup) initAbilityType() {
	typeSlice := []EnumAPIResource{
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
			Name:        string(database.AbilityTypeUnspecifiedAbility),
			Description: "Abilities that don't fit the other categories. Most of these are accessible from the start of the game.",
		},
		{
			Name:        string(database.AbilityTypeEnemyAbility),
			Description: "Abilities that are used by monsters.",
		},
	}

	t.AbilityType = EnumType[database.AbilityType, any]{
		name:         "ability type",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.AbilityType { return database.AbilityType(s) },
		nullConvFunc: nil,
		getNullEnum:  nil,
	}
}

func (t *TypeLookup) initUnitType() {
	typeSlice := []EnumAPIResource{
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
		name:         "unit type",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.UnitType { return database.UnitType(s) },
		nullConvFunc: nil,
		getNullEnum:  nil,
	}
}

func (t *TypeLookup) initItemType() {
	typeSlice := []EnumAPIResource{
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
		name:         "item type",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.ItemType { return database.ItemType(s) },
		nullConvFunc: nil,
		getNullEnum:  nil,
	}
}

func (t *TypeLookup) initQuestType() {
	typeSlice := []EnumAPIResource{
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
		name:         "quest type",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.QuestType { return database.QuestType(s) },
		nullConvFunc: nil,
		getNullEnum:  nil,
	}
}

func (t *TypeLookup) initAaActivationCondition() {
	typeSlice := []EnumAPIResource{
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
		name:         "auto ability activation condition",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.AaActivationCondition { return database.AaActivationCondition(s) },
	}
}

func (t *TypeLookup) initAreaConnectionType() {
	typeSlice := []EnumAPIResource{
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
		name:         "area connection type",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.AreaConnectionType { return database.AreaConnectionType(s) },
	}
}

func (t *TypeLookup) initArenaCreationCategory() {
	typeSlice := []EnumAPIResource{
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
		name:         "arena creation category",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.MaCreationCategory { return database.MaCreationCategory(s) },
		nullConvFunc: h.NullMaCreationCategory,
		getNullEnum:  h.GetNullMaCreationCategory,
	}
}

func (t *TypeLookup) initArranger() {
	typeSlice := []EnumAPIResource{
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
		name:         "arranger",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.Arranger { return database.Arranger(s) },
		nullConvFunc: h.NullArranger,
		getNullEnum:  h.GetNullArranger,
	}
}

func (t *TypeLookup) initAutoAbilityCategory() {
	typeSlice := []EnumAPIResource{
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
		name:         "auto-ability category",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.AutoAbilityCategory { return database.AutoAbilityCategory(s) },
	}
}

func (t *TypeLookup) initAvailabilityType() {
	typeSlice := []EnumAPIResource{
		{
			Name:        string(database.AvailabilityTypeAlways),
			Description: "The resource is always available once its location is reached in the story.",
		},
		{
			Name:        string(database.AvailabilityTypeStory),
			Description: "The resource is only available during the events of the story.",
		},
		{
			Name:        string(database.AvailabilityTypePost),
			Description: "The resource is only available after acquiring the airship. Note that for the resources that are located inside Sin to become available, you have to do the boss rush against it.",
		},
		{
			Name:        string(database.AvailabilityTypePostStory),
			Description: "The resource is only available during the events of the story that happen after acquiring the airship.",
		},
		{
			Name:		 string(database.AvailabilityTypePostGame),
			Description: "The resource is available in the post-game, meaning it either was already available before acquiring the airship, or it becomes available after acquiring the airship. This excludes story-specific resources. This value is essentially a combination of 'always' and 'post'.",
		},
		{
			Name:		 string(database.AvailabilityTypeStoryOnly),
			Description: "The resource is only available during the events of the story. This value is essentially a combination of 'story' and 'post-story'.",
		},
	}

	t.AvailabilityType = EnumType[database.AvailabilityType, any]{
		name:         "availability type",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.AvailabilityType { return database.AvailabilityType(s) },
		aliasses: 	  map[string][]database.AvailabilityType{
			string(database.AvailabilityTypePostGame): {
				database.AvailabilityTypeAlways,
				database.AvailabilityTypePost,
			},

			string(database.AvailabilityTypeStoryOnly): {
				database.AvailabilityTypeStory,
				database.AvailabilityTypePostStory,
			},
		},
	}
}

func (t *TypeLookup) initBlitzballTournamentCategory() {
	typeSlice := []EnumAPIResource{
		{
			Name: string(database.BlitzballTournamentCategoryLeague),
		},
		{
			Name: string(database.BlitzballTournamentCategoryTournament),
		},
	}

	t.BlitzballTournamentCategory = EnumType[database.BlitzballTournamentCategory, any]{
		name:         "blitzball tournament category",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.BlitzballTournamentCategory { return database.BlitzballTournamentCategory(s) },
	}
}

func (t *TypeLookup) initCelestialFormula() {
	typeSlice := []EnumAPIResource{
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
		name:         "celestial formula",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.CelestialFormula { return database.CelestialFormula(s) },
	}
}

func (t *TypeLookup) initCharacterClassCategory() {
	typeSlice := []EnumAPIResource{
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
		name:         "character class category",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.CharacterClassCategory { return database.CharacterClassCategory(s) },
	}
}

func (t *TypeLookup) initComposer() {
	typeSlice := []EnumAPIResource{
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
		name:         "composer",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.Composer { return database.Composer(s) },
		nullConvFunc: h.NullComposer,
		getNullEnum:  h.GetNullComposer,
	}
}

func (t *TypeLookup) initCounterType() {
	typeSlice := []EnumAPIResource{
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
		name:         "counter type",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.CounterType { return database.CounterType(s) },
		nullConvFunc: h.NullCounterType,
		getNullEnum:  h.GetNullCounterType,
	}
}

func (t *TypeLookup) initCTBIconType() {
	typeSlice := []EnumAPIResource{
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
		name:         "ctb icon type",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.CtbIconType { return database.CtbIconType(s) },
	}
}

func (t *TypeLookup) initCreationArea() {
	typeSlice := []EnumAPIResource{
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
		name:         "creation area",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.MaCreationArea { return database.MaCreationArea(s) },
		nullConvFunc: h.NullMaCreationArea,
		getNullEnum:  h.GetNullMaCreationArea,
	}
}

func (t *TypeLookup) initEquipClass() {
	typeSlice := []EnumAPIResource{
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
		name:         "equip class",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.EquipClass { return database.EquipClass(s) },
	}
}

func (t *TypeLookup) initEquipType() {
	typeSlice := []EnumAPIResource{
		{
			Name:        string(database.EquipTypeWeapon),
		},
		{
			Name:        string(database.EquipTypeArmor),
		},
	}

	t.EquipType = EnumType[database.EquipType, any]{
		name:         "equip type",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.EquipType { return database.EquipType(s) },
	}
}

func (t *TypeLookup) initItemCategory() {
	typeSlice := []EnumAPIResource{
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
		name:         "item category",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.ItemCategory { return database.ItemCategory(s) },
	}
}

func (t *TypeLookup) initKeyItemCategory() {
	typeSlice := []EnumAPIResource{
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
		name:         "key-item category",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.KeyItemCategory { return database.KeyItemCategory(s) },
	}
}

func (t *TypeLookup) initLootType() {

	typeSlice := []EnumAPIResource{
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
		name:         "loot type",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.LootType { return database.LootType(s) },
	}
}

func (t *TypeLookup) initMixCategory() {
	typeSlice := []EnumAPIResource{
		{
			Name:        string(database.MixCategoryRecovery),
			Description: "Mixes that heal the party.",
		},
		{
			Name:        string(database.MixCategoryPositiveStatus),
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
		name:         "mix category",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.MixCategory { return database.MixCategory(s) },
	}
}

func (t *TypeLookup) initMonsterFormationCategory() {
	typeSlice := []EnumAPIResource{
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
		name:         "monster-formation category",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.MonsterFormationCategory { return database.MonsterFormationCategory(s) },
	}
}

func (t *TypeLookup) initMonsterSpecies() {
	typeSlice := []EnumAPIResource{
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
		name:         "monster species",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.MonsterSpecies { return database.MonsterSpecies(s) },
	}
}

func (t *TypeLookup) initMonsterCategory() {
	typeSlice := []EnumAPIResource{
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
		name:         "monster category",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.MonsterCategory { return database.MonsterCategory(s) },
	}
}

func (t *TypeLookup) initNodePosition() {
	typeSlice := []EnumAPIResource{
		{
			Name: string(database.NodePositionNeighboring),
			Description: "The sphere can target neighboring nodes, or the node the selected character is currently positioned.",
		},
		{
			Name: string(database.NodePositionAllyPosition),
			Description: "The sphere can only target nodes, another character is currently positioned.",
		},
		{
			Name: string(database.NodePositionAny),
			Description: "The sphere can target any node that it is able to.",
		},
	}

	t.NodePosition = EnumType[database.NodePosition, any]{
		name:         "node position",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.NodePosition { return database.NodePosition(s) },
	}
}

func (t *TypeLookup) initNodeState() {
	typeSlice := []EnumAPIResource{
		{
			Name: string(database.NodeStateActiveSelf),
			Description: "The node has been activated by the selected character.",
		},
		{
			Name: string(database.NodeStateActiveAlly),
			Description: "The node hasn't been activated by the selected character, but by another character.",
		},
		{
			Name: string(database.NodeStateActiveAny),
			Description: "The node has been activated by at least one character.",
		},
		{
			Name: string(database.NodeStateInactive),
			Description: "The node hasn't been activated by the selected character.",
		},
		{
			Name: string(database.NodeStateAny),
			Description: "The node's activation state doesn't matter for this resource.",
		},
	}

	t.NodeState = EnumType[database.NodeState, database.NullNodeState]{
		name:         "node state",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.NodeState { return database.NodeState(s) },
		nullConvFunc: h.NullNodeState,
		getNullEnum:  h.GetNullNodeState,
	}
}

func (t *TypeLookup) initOverdriveModeType() {
	typeSlice := []EnumAPIResource{
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
		name:         "overdrive mode type",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.OverdriveModeType { return database.OverdriveModeType(s) },
	}
}

func (t *TypeLookup) initPlayerAbilityCategory() {
	typeSlice := []EnumAPIResource{
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
		name:         "player ability category",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.PlayerAbilityCategory { return database.PlayerAbilityCategory(s) },
	}
}

func (t *TypeLookup) initShopCategory() {
	typeSlice := []EnumAPIResource{
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
		name:         "shop category",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.ShopCategory { return database.ShopCategory(s) },
	}
}

func (t *TypeLookup) initShopType() {
	typeSlice := []EnumAPIResource{
		{
			Name: string(database.ShopTypePreAirship),
		},
		{
			Name: string(database.ShopTypePostAirship),
		},
	}

	t.ShopType = EnumType[database.ShopType, database.NullShopType]{
		name:         "shop type",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.ShopType { return database.ShopType(s) },
		nullConvFunc: h.NullShopType,
		getNullEnum:  h.GetNullShopType,
	}
}

func (t *TypeLookup) initSphereColor() {
	typeSlice := []EnumAPIResource{
		{
			Name:        string(database.SphereColorRed),
		},
		{
			Name:        string(database.SphereColorYellow),
		},
		{
			Name:        string(database.SphereColorBlack),
		},
		{
			Name:        string(database.SphereColorPurple),
		},
		{
			Name:        string(database.SphereColorBlue),
		},
		{
			Name:        string(database.SphereColorWhite),
		},
	}

	t.SphereColor = EnumType[database.SphereColor, any]{
		name:         "sphere color",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.SphereColor { return database.SphereColor(s) },
	}
}

func (t *TypeLookup) initTreasureType() {
	typeSlice := []EnumAPIResource{
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
		name:         "treasure type",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.TreasureType { return database.TreasureType(s) },
	}
}






func (t *TypeLookup) initAccSourceType() {
	typeSlice := []EnumAPIResource{
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
		name:         "accuracy source type",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.AccSourceType { return database.AccSourceType(s) },
	}
}

func (t *TypeLookup) initAttackType() {
	typeSlice := []EnumAPIResource{
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
		name:         "attack type",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.AttackType { return database.AttackType(s) },
	}
}

func (t *TypeLookup) initBreakDmgLimitType() {
	typeSlice := []EnumAPIResource{
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
		name:         "break damage limit type",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.BreakDmgLmtType { return database.BreakDmgLmtType(s) },
		nullConvFunc: h.NullBreakDmgLmtType,
		getNullEnum:  h.GetNullBreakDmgLmtType,
	}
}

func (t *TypeLookup) initCalculationType() {
	typeSlice := []EnumAPIResource{
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
		name:         "calculation type",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.CalculationType { return database.CalculationType(s) },
	}
}

func (t *TypeLookup) initCriticalType() {
	typeSlice := []EnumAPIResource{
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
		name:         "critical type",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.CriticalType { return database.CriticalType(s) },
		nullConvFunc: h.NullCriticalType,
		getNullEnum:  h.GetNullCriticalType,
	}
}

func (t *TypeLookup) initCtbAttackType() {
	typeSlice := []EnumAPIResource{
		{
			Name: string(database.CtbAttackTypeAttack),
		},
		{
			Name: string(database.CtbAttackTypeHeal),
		},
	}

	t.CtbAttackType = EnumType[database.CtbAttackType, any]{
		name:         "ctb attack type",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.CtbAttackType { return database.CtbAttackType(s) },
	}
}

func (t *TypeLookup) initDamageFormula() {
	typeSlice := []EnumAPIResource{
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
		name:         "damage formula",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.DamageFormula { return database.DamageFormula(s) },
	}
}

func (t *TypeLookup) initDamageType() {
	typeSlice := []EnumAPIResource{
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
		name:         "damage type",
		isEndpoint:   true,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.DamageType { return database.DamageType(s) },
	}
}

func (t *TypeLookup) initDelayType() {
	typeSlice := []EnumAPIResource{
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
		name:         "delay type",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.DelayType { return database.DelayType(s) },
	}
}

func (t *TypeLookup) initDurationType() {
	typeSlice := []EnumAPIResource{
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
		name:         "duration type",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.DurationType { return database.DurationType(s) },
	}
}

func (t *TypeLookup) initTargetType() {
	typeSlice := []EnumAPIResource{
		{
			Name:        string(database.TargetTypeSelf),
			Description: "The action targets its user.",
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
	}

	t.TargetType = EnumType[database.TargetType, database.NullTargetType]{
		name:         "target type",
		isEndpoint:   false,
		lookup:       enumSliceToMap(typeSlice),
		convFunc:     func(s string) database.TargetType { return database.TargetType(s) },
		nullConvFunc: h.NullTargetType,
		getNullEnum:  h.GetNullTargetType,
	}
}
