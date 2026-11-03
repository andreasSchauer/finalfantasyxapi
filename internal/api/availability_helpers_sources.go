package api

import (
	"database/sql"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type locBasedSources struct {
	SrcType         ViewSourceType
	RequiredSources []string
	ExcludedSources []string
	MonsterID       sql.NullInt32
	ItemID          sql.NullInt32
	KeyItemID       sql.NullInt32
	AutoAbilityID   sql.NullInt32
	Methods         []string
}

func (s locBasedSources) IsZero() bool {
	return reqsAreZero(s.RequiredSources, s.SrcType) &&
		s.ExcludedSources == nil &&
		h.NullInt32IsZero(s.MonsterID) &&
		h.NullInt32IsZero(s.ItemID) &&
		h.NullInt32IsZero(s.KeyItemID) &&
		s.Methods == nil
}

func reqsAreZero(reqs []string, srcType ViewSourceType) bool {
	return len(reqs) == 1 && reqs[0] == string(srcType)
}

func getLocBasedSources[T seedtypes.Lookupable, R any, A APIResource](cfg *Config, r *http.Request, i handlerInput[T, R, A], srcType ViewSourceType) (locBasedSources, error) {
	reqs := []string{string(srcType)}
	excls := []string{}

	monID, err := nullID(r, cfg.E.Monsters, QpnMonster, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}
	if !queryIsEmpty(err) {
		reqs = append(reqs, "monster-single")
	}

	itemID, err := nullID(r, cfg.E.Items, QpnItem, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}
	if !queryIsEmpty(err) {
		reqs = append(reqs, string(ViewSourceTypeItem))
	}

	keyItemID, err := nullID(r, cfg.E.KeyItems, QpnKeyItem, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}
	if !queryIsEmpty(err) {
		reqs = append(reqs, string(ViewSourceTypeKeyItem))
	}

	autoAbilityID, err := nullID(r, cfg.E.AutoAbilities, QpnAutoAbility, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}
	if !queryIsEmpty(err) {
		reqs = append(reqs, string(ViewSourceTypeAutoAbility))
	}

	methods, err := i.queryLookup[QpnMethods].parseValueList(cfg, r)
	if errExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}

	reqs, excls, err = parseBoolSources(r, i, reqs, excls, map[QueryParamName]string{
		QpnMonsters:   string(ViewSourceTypeMonster),
		QpnBossFights: string(ViewSourceTypeBoss),
		QpnShops:      string(ViewSourceTypeShop),
		QpnTreasures:  string(ViewSourceTypeTreasure),
		QpnSidequests: string(ViewSourceTypeQuest),
	})
	if err != nil {
		return locBasedSources{}, err
	}

	sources := locBasedSources{
		SrcType:         srcType,
		RequiredSources: reqs,
		ExcludedSources: h.SliceOrNil(excls),
		MonsterID:       monID,
		ItemID:          itemID,
		KeyItemID:       keyItemID,
		AutoAbilityID:   autoAbilityID,
		Methods:         h.SliceOrNil(methods),
	}

	return sources, nil
}

type shopSources struct {
	AvlType         string
	RequiredSources []string
	ExcludedSources []string
	AutoAbilityID   sql.NullInt32
	CharacterID     sql.NullInt32
	EmptySlots      []int32
}

func getShopSources[T seedtypes.Lookupable, R any, A APIResource](cfg *Config, r *http.Request, i handlerInput[T, R, A]) (shopSources, error) {
	avlType := AvlTypeSelf
	reqs := []string{}
	excls := []string{}

	autoAbilityID, err := nullID(r, cfg.E.AutoAbilities, QpnAutoAbility, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return shopSources{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		reqs = append(reqs, "equip-filter")
	}

	emptySlots, err := i.queryLookup[QpnEmptySlots].parseIntList(cfg, r)
	if errExceptEmptyQuery(err) {
		return shopSources{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		reqs = append(reqs, "equip-filter")
	}

	charID, err := nullNameID(r, cfg.E.Characters, QpnCharacter, i.queryLookup)
	if errExceptEmptyQuery(err) {
		return shopSources{}, err
	}
	if !queryIsEmpty(err) {
		avlType = AvlTypeContext
		reqs = append(reqs, "equip-filter")
	}

	reqs, excls, err = parseBoolSources(r, i, reqs, excls, map[QueryParamName]string{
		QpnItems:     string(ViewSourceTypeItem),
		QpnEquipment: string(ViewSourceTypeEquipment),
	})
	if err != nil {
		return shopSources{}, err
	}
	if len(reqs) > 0 || len(excls) > 0 {
		avlType = AvlTypeContext
	}

	sources := shopSources{
		AvlType:         string(avlType),
		RequiredSources: h.SliceOrNil(reqs),
		ExcludedSources: h.SliceOrNil(excls),
		AutoAbilityID:   autoAbilityID,
		CharacterID:     charID,
		EmptySlots:      h.SliceOrNil(emptySlots),
	}

	return sources, nil
}

func parseBoolSources[T seedtypes.Lookupable, R any, A APIResource](r *http.Request, i handlerInput[T, R, A], reqs, excls []string, sourceMap map[QueryParamName]string) ([]string, []string, error) {
	for queryName := range sourceMap {
		b, err := i.queryLookup[queryName].parseBool(r)
		if errExceptEmptyQuery(err) {
			return nil, nil, err
		}
		if !queryIsEmpty(err) {
			if b {
				reqs = append(reqs, sourceMap[queryName])
			} else {
				excls = append(excls, sourceMap[queryName])
			}
		}
	}

	return reqs, excls, nil
}
