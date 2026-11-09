package typestrings

type AvlType string

const (
	AvlTypeSelf     AvlType = "self"
	AvlTypeContext  AvlType = "context"
	AvlTypeContext2 AvlType = "context-2"
)

type ViewSourceType string

const (
	VstMonster          ViewSourceType = "monster"
	VstBoss             ViewSourceType = "boss"
	VstMonsterFormation ViewSourceType = "monster-formation"
	VstLocation         ViewSourceType = "location"
	VstSublocation      ViewSourceType = "sublocation"
	VstArea             ViewSourceType = "area"
	VstTreasure         ViewSourceType = "treasure"
	VstShop             ViewSourceType = "shop"
	VstQuest            ViewSourceType = "quest"
	VstBlitzball        ViewSourceType = "blitzball"
	VstItem             ViewSourceType = "item"
	VstKeyItem          ViewSourceType = "key-item"
	VstAutoAbility      ViewSourceType = "auto-ability"
	VstEquipment        ViewSourceType = "equip"
)
