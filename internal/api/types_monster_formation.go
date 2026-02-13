package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"


type MonsterFormation struct {
	ID              int32                     `json:"id"`
	Category        string                    `json:"category"`
	IsForcedAmbush  bool                      `json:"is_forced_ambush"`
	CanEscape       bool                      `json:"can_escape"`
	Notes           *string                   `json:"notes,omitempty"`
	BossMusic       *FormationBossSong        `json:"boss_music,omitempty"`
	Monsters        []MonsterAmount           `json:"monsters"`
	Areas           []EncounterArea           `json:"areas"`
	TriggerCommands []FormationTriggerCommand `json:"trigger_commands"`
}



type EncounterArea struct {
	Area          AreaAPIResource `json:"area"`
	Specification *string         `json:"specification"`
}

func (ec EncounterArea) GetAPIResource() APIResource {
	return ec.Area
}

func convertEncounterArea(cfg *Config, el seeding.EncounterArea) EncounterArea {
	return EncounterArea{
		Area:          locAreaToAreaAPIResource(cfg, cfg.e.areas, el.LocationArea),
		Specification: el.Specification,
	}
}



type FormationTriggerCommand struct {
	Ability   NamedAPIResource   `json:"ability"`
	Condition *string            `json:"condition"`
	UseAmount *int32             `json:"use_amount"`
	Users     []NamedAPIResource `json:"users"`
}

func (tc FormationTriggerCommand) GetAPIResource() APIResource {
	return tc.Ability
}

func convertFormationTriggerCommand(cfg *Config, tc seeding.FormationTriggerCommand) FormationTriggerCommand {
	return FormationTriggerCommand{
		Ability:   nameToNamedAPIResource(cfg, cfg.e.triggerCommands, tc.Name, tc.Version),
		Condition: tc.Condition,
		UseAmount: tc.UseAmount,
		Users:     namesToNamedAPIResources(cfg, cfg.e.characterClasses, tc.Users),
	}
}



type FormationBossSong struct {
	Song             NamedAPIResource `json:"song"`
	CelebrateVictory bool             `json:"celebrate_victory"`
}

func (bs FormationBossSong) GetAPIResource() APIResource {
	return bs.Song
}

func convertFormationBossSong(cfg *Config, bossMusic seeding.FormationBossSong) FormationBossSong {
	return FormationBossSong{
		Song:             nameToNamedAPIResource(cfg, cfg.e.songs, bossMusic.Song, nil),
		CelebrateVictory: bossMusic.CelebrateVictory,
	}
}