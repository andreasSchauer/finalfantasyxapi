package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type MonsterSimple struct {
	ID             int32                   `json:"id"`
	URL            string                  `json:"url"`
	Name           string                  `json:"name"`
	Version        *int32                  `json:"version,omitempty"`
	Specification  *string                 `json:"specification,omitempty"`
	HP             string                  `json:"hp"`
	AP             string                  `json:"ap"`
	Gil            int32                   `json:"gil"`
	MaxBribeAmount *int32                  `json:"max_bribe_amount,omitempty"`
	Availability   string			  	   `json:"availability"`
	IsRepeatable   bool                    `json:"is_repeatable"`
	CanBeCaptured  bool                    `json:"can_be_captured"`
	RonsoRages     []string                `json:"ronso_rages,omitempty"`
	Areas          []string                `json:"areas"`
	Items          *MonsterItemsSimple     `json:"items,omitempty"`
	Equipment      *MonsterEquipmentSimple `json:"equipment,omitempty"`
}

func (m MonsterSimple) GetURL() string {
	return m.URL
}

func createMonsterSimple(cfg *Config, r *http.Request, id int32, subsection Subsection) (SimpleResource, error) {
	i := cfg.e.monsters
	mon, _ := seeding.GetResourceByID(id, i.objLookupID)

	areaIDs := subsection.relations[id][RelationAreas]

	monSimple := MonsterSimple{
		ID:             mon.ID,
		URL:            createResourceURL(cfg, i.endpoint, id),
		Name:           mon.Name,
		Version:        mon.Version,
		Specification:  mon.Specification,
		HP:             getMonsterSimpleHP(mon),
		AP:             getMonsterSimpleAP(mon),
		Gil:            mon.Gil,
		MaxBribeAmount: getMonsterSimpleBribeAmount(mon, getMonsterHP(mon)),
		Availability:   mon.Availability,
		IsRepeatable:   mon.IsRepeatable,
		CanBeCaptured:  mon.CanBeCaptured,
		RonsoRages:     h.SliceOrNil(mon.RonsoRages),
		Areas:          convertObjSlice(cfg, areaIDs, idToLocAreaString),
		Items:          convertObjPtr(cfg, mon.Items, convertMonsterItemsSimple),
		Equipment:      convertObjPtr(cfg, mon.Equipment, convertMonsterEquipmentSimple),
	}

	return monSimple, nil
}

func getMonsterHP(mon seeding.Monster) int32 {
	for _, stat := range mon.BaseStats {
		if stat.StatName == "hp" {
			return stat.Value
		}
	}

	return 0
}

func getMonsterSimpleHP(mon seeding.Monster) string {
	hp := getMonsterHP(mon)
	return fmt.Sprintf("%d (%d)", hp, mon.OverkillDamage)
}

func getMonsterSimpleAP(mon seeding.Monster) string {
	return fmt.Sprintf("%d (%d)", mon.AP, mon.APOverkill)
}

func getMonsterSimpleBribeAmount(mon seeding.Monster, hp int32) *int32 {
	if mon.Items == nil || mon.Items.Bribe == nil {
		return nil
	}

	bribeAmount := hp * 25
	return &bribeAmount
}


func getMonsterSectionRelations(cfg *Config, r *http.Request, monIDs []int32) (map[int32]map[Relation][]int32, error) {
	i := cfg.e.monsters
	relations := make(map[int32]map[Relation][]int32)

	monsterJunctions, err := getJunctions(r, monIDs, i.resourceType, cfg.e.areas.resourceType, cfg.db.GetMonsterAreaIdPairs, juncMonsterArea)
	if err != nil {
		return nil, err
	}

	for _, monID := range monIDs {
		relationMap := make(map[Relation][]int32)

		relationMap[RelationMonsters], monsterJunctions = getJunctionIDs(monID, monsterJunctions)

		relations[monID] = relationMap
	}

	return relations, nil
}