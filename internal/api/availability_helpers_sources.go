package api

import (
	"database/sql"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type locBasedSources struct {
	SrcType         ts.ViewSourceType
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

func reqsAreZero(reqs []string, srcType ts.ViewSourceType) bool {
	return len(reqs) == 1 && reqs[0] == string(srcType)
}

func getLocBasedSources[T seedtypes.Lookupable, R any, A apiresource.APIResource](r *http.Request, i handlerinput.Main[T, R, A], srcType ts.ViewSourceType) (locBasedSources, error) {
	reqs := []string{string(srcType)}
	excls := []string{}

	monID, err := handlerinput.NullID(r, apilookups.E.Monsters, ts.QpnMonster, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}
	if !errs.QueryIsEmpty(err) {
		reqs = append(reqs, "monster-single")
	}

	itemID, err := handlerinput.NullID(r, apilookups.E.Items, ts.QpnItem, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}
	if !errs.QueryIsEmpty(err) {
		reqs = append(reqs, string(ts.VstItem))
	}

	keyItemID, err := handlerinput.NullID(r, apilookups.E.KeyItems, ts.QpnKeyItem, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}
	if !errs.QueryIsEmpty(err) {
		reqs = append(reqs, string(ts.VstKeyItem))
	}

	autoAbilityID, err := handlerinput.NullID(r, apilookups.E.AutoAbilities, ts.QpnAutoAbility, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}
	if !errs.QueryIsEmpty(err) {
		reqs = append(reqs, string(ts.VstAutoAbility))
	}

	methods, err := i.QueryLookup[ts.QpnMethods].ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return locBasedSources{}, err
	}

	reqs, excls, err = parseBoolSources(r, i, reqs, excls, map[ts.QueryParamName]string{
		ts.QpnMonsters:   string(ts.VstMonster),
		ts.QpnBossFights: string(ts.VstBoss),
		ts.QpnShops:      string(ts.VstShop),
		ts.QpnTreasures:  string(ts.VstTreasure),
		ts.QpnSidequests: string(ts.VstQuest),
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

func getShopSources[T seedtypes.Lookupable, R any, A apiresource.APIResource](r *http.Request, i handlerinput.Main[T, R, A]) (shopSources, error) {
	avlType := ts.AvlTypeSelf
	reqs := []string{}
	excls := []string{}

	autoAbilityID, err := handlerinput.NullID(r, apilookups.E.AutoAbilities, ts.QpnAutoAbility, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return shopSources{}, err
	}
	if !errs.QueryIsEmpty(err) {
		avlType = ts.AvlTypeContext
		reqs = append(reqs, "equip-filter")
	}

	emptySlots, err := i.QueryLookup[ts.QpnEmptySlots].ParseIntList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return shopSources{}, err
	}
	if !errs.QueryIsEmpty(err) {
		avlType = ts.AvlTypeContext
		reqs = append(reqs, "equip-filter")
	}

	charID, err := handlerinput.NullNameID(r, apilookups.E.Characters, ts.QpnCharacter, i.QueryLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return shopSources{}, err
	}
	if !errs.QueryIsEmpty(err) {
		avlType = ts.AvlTypeContext
		reqs = append(reqs, "equip-filter")
	}

	reqs, excls, err = parseBoolSources(r, i, reqs, excls, map[ts.QueryParamName]string{
		ts.QpnItems:     string(ts.VstItem),
		ts.QpnEquipment: string(ts.VstEquipment),
	})
	if err != nil {
		return shopSources{}, err
	}
	if len(reqs) > 0 || len(excls) > 0 {
		avlType = ts.AvlTypeContext
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

func parseBoolSources[T seedtypes.Lookupable, R any, A apiresource.APIResource](r *http.Request, i handlerinput.Main[T, R, A], reqs, excls []string, sourceMap map[ts.QueryParamName]string) ([]string, []string, error) {
	for queryName := range sourceMap {
		b, err := i.QueryLookup[queryName].ParseBool(r)
		if errs.ErrExceptEmptyQuery(err) {
			return nil, nil, err
		}
		if !errs.QueryIsEmpty(err) {
			if b {
				reqs = append(reqs, sourceMap[queryName])
			} else {
				excls = append(excls, sourceMap[queryName])
			}
		}
	}

	return reqs, excls, nil
}
