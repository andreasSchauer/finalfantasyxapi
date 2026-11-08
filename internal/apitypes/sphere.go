package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type Sphere struct {
	ID                 int32                                             `json:"id"`
	Name               string                                            `json:"name"`
	Item               apiresource.Named                                 `json:"item"`
	Description        string                                            `json:"description"`
	Effect             string                                            `json:"effect"`
	SgDescription      string                                            `json:"sphere_grid_description"`
	SphereColor        string                                            `json:"sphere_color"`
	SphereEffect       string                                            `json:"sphere_effect"`
	TargetNodePosition string                                            `json:"target_node_position"`
	TargetNodeState    *string                                           `json:"target_node_state"`
	TargetableNodes    []string                                          `json:"targetable_nodes"`
	CreatedNode        *CreatedNode                                      `json:"created_node,omitempty"`
	Monsters           []MonItemAmts                                     `json:"monsters"`
	Treasures          []apiresource.ResourceAmount[apiresource.Unnamed] `json:"treasures"`
	Shops              []apiresource.Unnamed                             `json:"shops"`
	Quests             []apiresource.ResourceAmount[apiresource.Quest]   `json:"quests"`
	BlitzballPrizes    []apiresource.ResourceAmount[apiresource.Named]   `json:"blitzball_prizes"`
}

type CreatedNode struct {
	Node  string `json:"node"`
	Value int32  `json:"value"`
}
