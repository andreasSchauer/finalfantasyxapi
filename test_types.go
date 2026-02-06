package main

import "net/http"

/*
fields that need to be explicitly ignored with dontCheck:
- standard types and pointers (through compare),
- direct apiResource references or pointers
- structs and pointers to structs that are borrowed from the result
- basically anything that isn't a slice or a map and doesn't have nil checks in the test function body


fields that can be explicitly ignored with dontCheck:
- any field that is referenced explicitly as part of dont check in the function body


fields that are implicitly ignored by leaving them blank:
- slices of api resource references
- resourceAmount map references
- slices, structs, pointers to structs of some kind that have nil checks in the function body
*/

type testGeneral struct {
	requestURL     string
	expectedStatus int
	expectedErr    string
	dontCheck      map[string]bool
	expLengths     map[string]int
	httpHandler    func(http.ResponseWriter, *http.Request)
}

type expNameVer struct {
	id      int32
	name    string
	version *int32
}

type expUnique struct {
	id   int32
	name string
}

type expIdOnly struct {
	id int32
}

type expList struct {
	count          int
	previous       *string
	next           *string
	parentResource *string
	results        []int32
}

type expListParams struct {
	count    int
	previous *string
	next     *string
	results  []string
}

type expAreas struct {
	parentLocation    int32
	parentSublocation int32
	connectedAreas    []int32
	expLocRel
}

type expLocations struct {
	connectedLocations []int32
	sublocations       []int32
	expLocRel
}

type expSublocations struct {
	parentLocation        int32
	connectedSublocations []int32
	areas                 []int32
	expLocRel
}

type expLocRel struct {
	sidequests []int32
	characters []int32
	aeons      []int32
	shops      []int32
	treasures  []int32
	monsters   []int32
	formations []int32
	bgMusic    []int32
	cuesMusic  []int32
	fmvsMusic  []int32
	bossMusic  []int32
	fmvs       []int32
}

type expMonsterFormations struct {
	category		string
	isForcedAmbush	bool
	canEscape		bool
	bossMusic		*int32
	monsters		map[string]int32
	areas			[]int32
	triggerCommands	[]testFormationTC
}

type testFormationTC struct {
	Ability	int32
	Users	[]int32
}

func compareFormationTCs(test test, exp testFormationTC, got FormationTriggerCommand) {
	tcEndpoint := test.cfg.e.triggerCommands.endpoint
	charClassesEndpoint := test.cfg.e.characterClasses.endpoint

	compAPIResourcesFromID(test, "tc ability", tcEndpoint, exp.Ability, got.Ability)
	testResourceList(test, newResListTestFromIDs("tc users", charClassesEndpoint, exp.Users, got.Users))
}

type expOverdriveModes struct {
	description   string
	effect        string
	modeType      int32
	fillRate      *float32
	actionsAmount map[string]int32
}

type expMonsters struct {
	appliedState     *testAppliedState
	agility          *AgilityParams
	species          int32
	ctbIconType      int32
	distance         int32
	properties       []int32
	autoAbilities    []int32
	ronsoRages       []int32
	areas            []int32
	formations       []int32
	baseStats        map[string]int32
	items            *testMonItems
	bribeChances     []BribeChance
	equipment        *testMonEquipment
	elemResists      []testElemResist
	statusImmunities []int32
	statusResists    map[string]int32
	defaultState     *testDefaultState
	abilities        []string
}

type testAppliedState struct {
	condition     string
	isTemporary   bool
	appliedStatus *int32
}

type testDefaultState struct {
	IsTemporary bool                 `json:"is_temporary"`
	Changes     []testAltStateChange `json:"changes"`
}

type testAltStateChange struct {
	AlterationType   string
	Distance         *int32
	Properties       []int32
	AutoAbilities    []int32
	BaseStats        map[string]int32
	ElemResists      []testElemResist
	StatusImmunities []int32
	StatusResists    map[string]int32
	AddedStatus      *InflictedStatus
	RemovedStatus    *int32
}

type testElemResist struct {
	element  int32
	affinity int32
}

func compareMonsterElemResists(test test, exp testElemResist, got ElementalResist) {
	elemEndpoint := test.cfg.e.elements.endpoint
	affinityEndpoint := test.cfg.e.affinities.endpoint

	compAPIResourcesFromID(test, "elements", elemEndpoint, exp.element, got.Element)
	compAPIResourcesFromID(test, "affinities", affinityEndpoint, exp.affinity, got.Affinity)
}

type testMonItems struct {
	itemDropChance int32
	items          map[string]*int32
	otherItems     []int32
}

type testMonEquipment struct {
	abilitySlots      MonsterEquipmentSlots
	attachedAbilities MonsterEquipmentSlots
	weaponAbilities   []int32
	armorAbilities    []int32
}
