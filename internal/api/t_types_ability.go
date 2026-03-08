package api


type expAbility struct {
	testGeneral
	expNameVer
	rank				*int32
	appearsInHelpBar	bool
	canCopyCat			bool
	abilityType			int32
	typedAbility		string
	monsters			[]int32
	battleInteractions	[]expBattleInteraction
}

func (a expAbility) GetTestGeneral() testGeneral {
	return a.testGeneral
}

func compareAbilities(test test, exp expAbility, got Ability) {
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compIdApiResource(test, "ability type", test.cfg.e.abilityType.endpoint, exp.abilityType, got.Type)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)

	compPathApiResource(test, "typed ability", exp.typedAbility, got.TypedAbility)
	checkResIDsInSlice(test, "monsters", test.cfg.e.monsters.endpoint, exp.monsters, got.Monsters)
	compTestStructSlices(test, "battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type expEnemyAbility struct {
	testGeneral
	expNameVer
	rank				*int32
	appearsInHelpBar	bool
	canCopyCat			bool
	untypedAbility		int32
	monsters			[]int32
	battleInteractions	[]expBattleInteraction
}

func (a expEnemyAbility) GetTestGeneral() testGeneral {
	return a.testGeneral
}

func compareEnemyAbilities(test test, exp expEnemyAbility, got EnemyAbility) {
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)
	compIdApiResource(test, "untyped ability", test.cfg.e.abilities.endpoint, exp.untypedAbility, got.UntypedAbility)
	checkResIDsInSlice(test, "monsters", test.cfg.e.monsters.endpoint, exp.monsters, got.Monsters)
	compTestStructSlices(test, "battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type expItemAbility struct {
	testGeneral
	expUnique
	rank				*int32
	appearsInHelpBar	bool
	canCopyCat			bool
	untypedAbility		int32
	item				int32
	category			int32
	CanUseOutsideBattle	bool
	battleInteractions	[]expBattleInteraction
}

func (a expItemAbility) GetTestGeneral() testGeneral {
	return a.testGeneral
}

func compareItemAbilities(test test, exp expItemAbility, got ItemAbility) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)
	compIdApiResource(test, "untyped ability", test.cfg.e.abilities.endpoint, exp.untypedAbility, got.UntypedAbility)
	compIdApiResource(test, "item", test.cfg.e.items.endpoint, exp.item, got.Item)
	compIdApiResource(test, "category", test.cfg.e.itemCategory.endpoint, exp.category, got.Category)
	compare(test, "can use outside battle", exp.CanUseOutsideBattle, got.CanUseOutsideBattle)
	compTestStructSlices(test, "battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type expPlayerAbility struct {
	testGeneral
	expNameVer
	rank				*int32
	appearsInHelpBar	bool
	canCopyCat			bool
	untypedAbility		int32
	topmenu				*int32
	submenu				*int32
	openSubmenu			*int32
	stdChar				*int32
	expChar				*int32
	monsters			[]int32
	battleInteractions	[]expBattleInteraction
}

func (a expPlayerAbility) GetTestGeneral() testGeneral {
	return a.testGeneral
}

func comparePlayerAbilities(test test, exp expPlayerAbility, got PlayerAbility) {
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)
	compIdApiResource(test, "untyped ability", test.cfg.e.abilities.endpoint, exp.untypedAbility, got.UntypedAbility)
	compIdApiResourcePtrs(test, "topmenu", test.cfg.e.topmenus.endpoint, exp.topmenu, got.Topmenu)
	compIdApiResourcePtrs(test, "submenu", test.cfg.e.submenus.endpoint, exp.submenu, got.Submenu)
	compIdApiResourcePtrs(test, "open submenu", test.cfg.e.submenus.endpoint, exp.openSubmenu, got.OpenSubmenu)
	compIdApiResourcePtrs(test, "std sg char", test.cfg.e.characters.endpoint, exp.stdChar, got.StandardGridCharacter)
	compIdApiResourcePtrs(test, "exp sg char", test.cfg.e.characters.endpoint, exp.expChar, got.ExpertGridCharacter)
	checkResIDsInSlice(test, "monsters", test.cfg.e.monsters.endpoint, exp.monsters, got.Monsters)
	compTestStructSlices(test, "battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type expOverdriveAbility struct {
	testGeneral
	expNameVer
	rank				*int32
	untypedAbility		int32
	overdriveCommand	*int32
	overdrives			[]int32
	battleInteractions	[]expBattleInteraction
}

func (a expOverdriveAbility) GetTestGeneral() testGeneral {
	return a.testGeneral
}

func compareOverdriveAbilities(test test, exp expOverdriveAbility, got OverdriveAbility) {
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "rank", exp.rank, got.Rank)
	compIdApiResource(test, "untyped ability", test.cfg.e.abilities.endpoint, exp.untypedAbility, got.UntypedAbility)
	compIdApiResourcePtrs(test, "overdrive command", test.cfg.e.overdriveCommands.endpoint, exp.overdriveCommand, got.OverdriveCommand)
	checkResIDsInSlice(test, "overdrives", test.cfg.e.overdrives.endpoint, exp.overdrives, got.Overdrives)
	compTestStructSlices(test, "battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type expUnspecifiedAbility struct {
	testGeneral
	expNameVer
	rank				*int32
	appearsInHelpBar	bool
	canCopyCat			bool
	untypedAbility		int32
	topmenu				*int32
	submenu				*int32
	openSubmenu			*int32
	battleInteractions	[]expBattleInteraction
}

func (a expUnspecifiedAbility) GetTestGeneral() testGeneral {
	return a.testGeneral
}

func compareUnspecifiedAbilities(test test, exp expUnspecifiedAbility, got UnspecifiedAbility) {
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)
	compIdApiResource(test, "untyped ability", test.cfg.e.abilities.endpoint, exp.untypedAbility, got.UntypedAbility)
	compIdApiResourcePtrs(test, "topmenu", test.cfg.e.topmenus.endpoint, exp.topmenu, got.Topmenu)
	compIdApiResourcePtrs(test, "submenu", test.cfg.e.submenus.endpoint, exp.submenu, got.Submenu)
	compIdApiResourcePtrs(test, "open submenu", test.cfg.e.submenus.endpoint, exp.openSubmenu, got.OpenSubmenu)
	compTestStructSlices(test, "battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}

type expTriggerCommand struct {
	testGeneral
	expNameVer
	rank				*int32
	appearsInHelpBar	bool
	canCopyCat			bool
	untypedAbility		int32
	topmenu				*int32
	usedBy				[]int32
	monsterFormations	[]int32
	battleInteractions	[]expBattleInteraction
}

func (a expTriggerCommand) GetTestGeneral() testGeneral {
	return a.testGeneral
}

func compareTriggerCommands(test test, exp expTriggerCommand, got TriggerCommand) {
	compareExpNameVer(test, exp.expNameVer, got.ID, got.Name, got.Version)
	compare(test, "rank", exp.rank, got.Rank)
	compare(test, "appears in help bar", exp.appearsInHelpBar, got.AppearsInHelpBar)
	compare(test, "can copycat", exp.canCopyCat, got.CanCopycat)
	compIdApiResource(test, "untyped ability", test.cfg.e.abilities.endpoint, exp.untypedAbility, got.UntypedAbility)
	compIdApiResourcePtrs(test, "topmenu", test.cfg.e.topmenus.endpoint, exp.topmenu, got.Topmenu)
	checkResIDsInSlice(test, "used by", test.cfg.e.characterClasses.endpoint, exp.usedBy, got.UsedBy)
	checkResIDsInSlice(test, "monster formations", test.cfg.e.monsterFormations.endpoint, exp.monsterFormations, got.MonsterFormations)
	compTestStructSlices(test, "battle interactions", exp.battleInteractions, got.BattleInteractions, compareBattleInteractions)
}