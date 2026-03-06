package api

import (
	"fmt"
	"net/http"

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
	IsStoryBased   bool                    `json:"is_story_based"`
	IsRepeatable   bool                    `json:"is_repeatable"`
	CanBeCaptured  bool                    `json:"can_be_captured"`
	RonsoRages     []string                `json:"ronso_rages,omitempty"`
	Items          *MonsterItemsSimple     `json:"items,omitempty"`
	Equipment      *MonsterEquipmentSimple `json:"equipment,omitempty"`
}

func (m MonsterSimple) GetURL() string {
	return m.URL
}

func createMonsterSimple(cfg *Config, _ *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.monsters
	mon, _ := seeding.GetResourceByID(id, i.objLookupID)
	monHP := getMonsterHP(mon)

	monSimple := MonsterSimple{
		ID:             mon.ID,
		URL:            createResourceURL(cfg, i.endpoint, id),
		Name:           mon.Name,
		Version:        mon.Version,
		Specification:  mon.Specification,
		HP:             getMonsterSimpleHP(mon),
		AP:             getMonsterSimpleAP(mon),
		Gil:            mon.Gil,
		MaxBribeAmount: getMonsterSimpleBribeAmount(mon, monHP),
		IsStoryBased:   mon.IsStoryBased,
		IsRepeatable:   mon.IsRepeatable,
		CanBeCaptured:  mon.CanBeCaptured,
		RonsoRages:     sliceOrNil(mon.RonsoRages),
		Items:          convertObjPtr(cfg, mon.Items, convertMonsterItemsSimple),
		Equipment:      convertObjPtr(cfg, mon.Equipment, convertMonsterEquipmentSimple),
	}

	return monSimple, nil
}

func getMonsterHP(mon seeding.Monster) int32 {
	for _, stat := range mon.BaseStats {
		if stat.StatName == "hp" {
			return  stat.Value
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
