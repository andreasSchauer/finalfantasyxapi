package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type MonsterFormation struct {
	ID             int32   `json:"id"`
	Category       string  `json:"category"`
	Availability   string  `json:"availability"`
	IsForcedAmbush bool    `json:"is_forced_ambush"`
	CanEscape      bool    `json:"can_escape"`
	Notes          *string `json:"notes,omitempty"`
	FormationLoot
	BossMusic       *FormationBossSong                                         `json:"boss_music,omitempty"`
	Monsters        []apiresource.ResourceAmount[apiresource.NamedAPIResource] `json:"monsters"`
	Areas           []EncounterArea                                            `json:"areas"`
	TriggerCommands []FormationTriggerCommand                                  `json:"trigger_commands"`
}

type FormationLoot struct {
	AP         int32 `json:"ap"`
	ApOverkill int32 `json:"ap_overkill"`
	Gil        int32 `json:"gil"`
}

type EncounterArea struct {
	Area          apiresource.AreaAPIResource `json:"area"`
	Specification *string                     `json:"specification"`
	Availability  string                      `json:"availability"`
}

func (ec EncounterArea) GetAPIResource() apiresource.APIResource {
	return ec.Area
}

type FormationTriggerCommand struct {
	Ability   apiresource.NamedAPIResource   `json:"ability"`
	Condition *string                        `json:"condition"`
	UseAmount *int32                         `json:"use_amount"`
	Users     []apiresource.NamedAPIResource `json:"users"`
}

func (tc FormationTriggerCommand) GetAPIResource() apiresource.APIResource {
	return tc.Ability
}

type FormationBossSong struct {
	Song             apiresource.NamedAPIResource `json:"song"`
	CelebrateVictory bool                         `json:"celebrate_victory"`
}

func (bs FormationBossSong) GetAPIResource() apiresource.APIResource {
	return bs.Song
}