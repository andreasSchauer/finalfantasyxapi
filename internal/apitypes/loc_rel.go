package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type LocRel struct {
	Characters []apiresource.Named   `json:"characters"`
	Aeons      []apiresource.Named   `json:"aeons"`
	Shops      []apiresource.Unnamed `json:"shops"`
	Treasures  []apiresource.Unnamed `json:"treasures"`
	Monsters   []apiresource.Named   `json:"monsters"`
	Formations []apiresource.Unnamed `json:"monster_formations"`
	Quests     []apiresource.Quest   `json:"quests"`
	Music      *LocBasedMusic        `json:"music"`
	FMVs       []apiresource.Named   `json:"fmvs"`
}

type LocBasedMusic struct {
	BackgroundMusic []apiresource.Named `json:"background_music"`
	Cues            []apiresource.Named `json:"cues"`
	FMVs            []apiresource.Named `json:"fmvs"`
	BossMusic       []apiresource.Named `json:"boss_fights"`
}

func (m LocBasedMusic) IsZero() bool {
	return len(m.BackgroundMusic) == 0 &&
		len(m.Cues) == 0 &&
		len(m.FMVs) == 0 &&
		len(m.BossMusic) == 0
}
