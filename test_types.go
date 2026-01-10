package main

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

type expList struct {
	count    int
	previous *string
	next     *string
	results  []int32
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
	expLocBased
}

type expLocBased struct {
	sidequest      *int32
	connectedAreas []int32
	characters     []int32
	aeons          []int32
	shops          []int32
	treasures      []int32
	monsters       []int32
	formations     []int32
	bgMusic        []int32
	cuesMusic      []int32
	fmvsMusic      []int32
	bossMusic      []int32
	fmvs           []int32
}

type expOverdriveModes struct {
	description   string
	effect        string
	modeType      int32
	fillRate      *float32
	actionsAmount map[string]int32
}

type expMonsters struct {
	appliedState		*testAppliedState
	agility				*AgilityParams
	species				int32
	ctbIconType			int32
	distance			int32
	properties			[]int32
	autoAbilities		[]int32
	ronsoRages			[]int32
	locations			[]int32
	formations			[]int32
	baseStats			map[string]int32
	items 				*testItems
	bribeChances		[]BribeChance
	equipment 			*testEquipment
	elemResists			[]testElemResist
	statusImmunities	[]int32
	statusResists		map[string]int32
	alteredStates		[]string
	abilities			[]string
}

type testAppliedState struct {
	condition		string
	isTemporary		bool
	appliedStatus	*int32
}

type testElemResist struct {
	element 	int32
	affinity 	int32
}

type testItems struct {
	itemDropChance	int32
	items			map[string]*int32
	otherItems		[]int32
}

type testEquipment struct {
	abilitySlots 		MonsterEquipmentSlots
	attachedAbilities	MonsterEquipmentSlots
	weaponAbilities		[]int32
	armorAbilities		[]int32
}

