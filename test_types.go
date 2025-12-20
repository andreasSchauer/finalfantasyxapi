package main

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
	next     *string
	previous *string
	results  []string
}

type expAreas struct {
	parentLocation    string
	parentSublocation string
	expLocBased
}

type expLocBased struct {
	sidequest      *string
	connectedAreas []string
	characters     []string
	aeons          []string
	shops          []string
	treasures      []string
	monsters       []string
	formations     []string
	bgMusic        []string
	cuesMusic      []string
	fmvsMusic      []string
	bossMusic      []string
	fmvs           []string
}

type expOverdriveModes struct {
	description   string
	effect        string
	modeType      string
	fillRate      *float32
	actionsAmount map[string]int32
}

type expMonsters struct {
	appliedState		*testAppliedState
	agility				*AgilityParams
	species				string
	ctbIconType			string
	distance			int32
	properties			[]string
	autoAbilities		[]string
	ronsoRages			[]string
	locations			[]string
	formations			[]string
	baseStats			map[string]int32
	items 				*testItems
	equipment 			*testEquipment
	elemResists			[]testElemResist
	statusImmunities	[]string
	statusResists		map[string]int32
	alteredStates		[]string
	abilities			[]string
}

type testAppliedState struct {
	condition		string
	isTemporary		bool
	appliedStatus	*string
}

type testElemResist struct {
	element 	string
	affinity 	string
}

type testItems struct {
	itemDropChance	int32
	items			map[string]*string
	otherItems		[]string
}

type testEquipment struct {
	abilitySlots 		MonsterEquipmentSlots
	attachedAbilities	MonsterEquipmentSlots
	weaponAbilities		[]string
	armorAbilities		[]string
}

