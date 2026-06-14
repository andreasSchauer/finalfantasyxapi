package api

type AvlType string

const (
	AvlTypeSelf     AvlType = "self"
	AvlTypeContext  AvlType = "context"
	AvlTypeContext2 AvlType = "context-2"
)

type ViewSourceType string

const (
	ViewSourceTypeMonster          ViewSourceType = "monster"
	ViewSourceTypeBoss             ViewSourceType = "boss"
	ViewSourceTypeMonsterFormation ViewSourceType = "monster-formation"
	ViewSourceTypeLocation         ViewSourceType = "location"
	ViewSourceTypeSublocation      ViewSourceType = "sublocation"
	ViewSourceTypeArea             ViewSourceType = "area"
	ViewSourceTypeTreasure         ViewSourceType = "treasure"
	ViewSourceTypeShop             ViewSourceType = "shop"
	ViewSourceTypeQuest            ViewSourceType = "quest"
	ViewSourceTypeBlitzball        ViewSourceType = "blitzball"
	ViewSourceTypeItem             ViewSourceType = "item"
	ViewSourceTypeKeyItem          ViewSourceType = "key-item"
	ViewSourceTypeEquipment        ViewSourceType = "equip"
)
