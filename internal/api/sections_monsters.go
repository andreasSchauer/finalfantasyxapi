package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type MonsterSub struct {
	ID             int32                `json:"id"`
	URL            string               `json:"url"`
	Name           string               `json:"name"`
	Version        *int32               `json:"version,omitempty"`
	Specification  *string              `json:"specification,omitempty"`
	HP             int32                `json:"hp"`
	OverkillDamage int32                `json:"overkill_damage"`
	AP             int32                `json:"ap"`
	APOverkill     int32                `json:"ap_overkill"`
	Gil            int32                `json:"gil"`
	MaxBribeAmount *int32               `json:"max_bribe_amount"`
	IsStoryBased   bool                 `json:"is_story_based"`
	IsRepeatable   bool                 `json:"is_repeatable"`
	CanBeCaptured  bool                 `json:"can_be_captured"`
	RonsoRages     []string             `json:"ronso_rages"`
	Items          *MonsterItemsSub     `json:"items"`
	Equipment      *MonsterEquipmentSub `json:"equipment"`
}

func (m MonsterSub) GetURL() string {
	return m.URL
}


func createMonsterSub(cfg *Config, _ *http.Request, id int32) (SubResource, error) {
	i := cfg.e.monsters
	mon, _ := seeding.GetResourceByID(id, i.objLookupID)
	monHP := getMonsterSubHP(mon)

	monSub := MonsterSub{
		ID:             mon.ID,
		URL:            createResourceURL(cfg, i.endpoint, id),
		Name:           mon.Name,
		Version:        mon.Version,
		Specification:  mon.Specification,
		HP:             getMonsterSubHP(mon),
		OverkillDamage: mon.OverkillDamage,
		AP:             mon.AP,
		APOverkill:     mon.APOverkill,
		Gil:            mon.Gil,
		MaxBribeAmount: getMonsterSubBribeAmount(mon, monHP),
		IsStoryBased:   mon.IsStoryBased,
		IsRepeatable:   mon.IsRepeatable,
		CanBeCaptured:  mon.CanBeCaptured,
		RonsoRages:     mon.RonsoRages,
		Items:          convertObjPtr(cfg, mon.Items, convertMonsterSubItems),
		Equipment:      convertObjPtr(cfg, mon.Equipment, convertMonsterSubEquipment),
	}

	return monSub, nil
}


func getMonsterSubHP(mon seeding.Monster) int32 {
	for _, stat := range mon.BaseStats {
		if stat.StatName == "hp" {
			return stat.Value
		}
	}

	return 0
}

func getMonsterSubBribeAmount(mon seeding.Monster, hp int32) *int32 {
	if mon.Items == nil || mon.Items.Bribe == nil {
		return nil
	}

	bribeAmount := hp * 25
	return &bribeAmount
}
