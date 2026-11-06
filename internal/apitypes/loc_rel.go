package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"


type LocRel struct {
	Characters []apiresource.NamedAPIResource   `json:"characters"`
	Aeons      []apiresource.NamedAPIResource   `json:"aeons"`
	Shops      []apiresource.UnnamedAPIResource `json:"shops"`
	Treasures  []apiresource.UnnamedAPIResource `json:"treasures"`
	Monsters   []apiresource.NamedAPIResource   `json:"monsters"`
	Formations []apiresource.UnnamedAPIResource `json:"monster_formations"`
	Quests     []apiresource.QuestAPIResource   `json:"quests"`
	Music      *LocBasedMusic                   `json:"music"`
	FMVs       []apiresource.NamedAPIResource   `json:"fmvs"`
}

type LocBasedMusic struct {
	BackgroundMusic []apiresource.NamedAPIResource `json:"background_music"`
	Cues            []apiresource.NamedAPIResource `json:"cues"`
	FMVs            []apiresource.NamedAPIResource `json:"fmvs"`
	BossMusic       []apiresource.NamedAPIResource `json:"boss_fights"`
}

func (m LocBasedMusic) IsZero() bool {
	return len(m.BackgroundMusic) == 0 &&
		len(m.Cues) == 0 &&
		len(m.FMVs) == 0 &&
		len(m.BossMusic) == 0
}