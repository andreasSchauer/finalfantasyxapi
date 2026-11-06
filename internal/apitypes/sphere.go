package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type Sphere struct {
	ID                 int32                                                        `json:"id"`
	Name               string                                                       `json:"name"`
	Item               apiresource.NamedAPIResource                                 `json:"item"`
	Description        string                                                       `json:"description"`
	Effect             string                                                       `json:"effect"`
	SgDescription      string                                                       `json:"sphere_grid_description"`
	SphereColor        string                                                       `json:"sphere_color"`
	SphereEffect       string                                                       `json:"sphere_effect"`
	TargetNodePosition string                                                       `json:"target_node_position"`
	TargetNodeState    *string                                                      `json:"target_node_state"`
	TargetableNodes    []string                                                     `json:"targetable_nodes"`
	CreatedNode        *CreatedNode                                                 `json:"created_node,omitempty"`
	Monsters           []MonItemAmts                                                `json:"monsters"`
	Treasures          []apiresource.ResourceAmount[apiresource.UnnamedAPIResource] `json:"treasures"`
	Shops              []apiresource.UnnamedAPIResource                             `json:"shops"`
	Quests             []apiresource.ResourceAmount[apiresource.QuestAPIResource]   `json:"quests"`
	BlitzballPrizes    []apiresource.ResourceAmount[apiresource.NamedAPIResource]   `json:"blitzball_prizes"`
}

type CreatedNode struct {
	Node  string `json:"node"`
	Value int32  `json:"value"`
}