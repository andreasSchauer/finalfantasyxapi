package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"

type Sphere struct {
	ID					int32									`json:"id"`
	Name          		string 									`json:"name"`
	Item				NamedAPIResource						`json:"item"`
	Description			string									`json:"description"`
	Effect				string									`json:"effect"`
	SgDescription		string									`json:"sphere_grid_description"`
	SphereColor			string									`json:"sphere_color"`
	SphereEffect		string									`json:"sphere_effect"`
	TargetNodePosition	string									`json:"target_node_position"`
	TargetNodeState		*string									`json:"target_node_state"`
	TargetableNodes		[]string								`json:"targetable_nodes"`
	CreatedNode			*CreatedNode							`json:"created_node,omitempty"`
	Monsters            []MonItemAmts                      		`json:"monsters"`
	Treasures          	[]ResourceAmount[UnnamedAPIResource] 	`json:"treasures"`
	Shops              	[]UnnamedAPIResource               		`json:"shops"`
	Quests             	[]ResourceAmount[QuestAPIResource]		`json:"quests"`
	BlitzballPrizes    	[]ResourceAmount[NamedAPIResource]		`json:"blitzball_prizes"`
}

type CreatedNode struct {
	Node	string	`json:"node"`
	Value	int32	`json:"value"`
}

func convertCreatedNode(cfg *Config, n seeding.CreatedNode) CreatedNode {
	return CreatedNode{
		Node: 	n.Node,
		Value: 	n.Value,
	}
}