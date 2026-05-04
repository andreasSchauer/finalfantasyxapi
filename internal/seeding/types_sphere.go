package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Sphere struct {
	ID                 int32
	ItemID             int32
	Name               string       `json:"name"`
	SgDescription      string       `json:"sphere_grid_description"`
	SphereColor        string       `json:"sphere_color"`
	SphereEffect       string       `json:"sphere_effect"`
	TargetNodePosition string       `json:"target_node_position"`
	TargetNodeState    *string      `json:"target_node_state"`
	TargetableNodes    []string     `json:"targetable_nodes"`
	CreatedNode        *CreatedNode `json:"created_node"`
}

func (s Sphere) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.ItemID,
		s.SgDescription,
		s.SphereColor,
		s.SphereEffect,
		s.TargetNodePosition,
		s.TargetNodeState,
		h.ObjPtrToID(s.CreatedNode),
	}
}

func (s Sphere) GetID() int32 {
	return s.ID
}

func (s Sphere) Error() string {
	return fmt.Sprintf("sphere %s", s.Name)
}

func (s Sphere) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   s.ID,
		Name: s.Name,
	}
}

type CreatedNode struct {
	ID    int32
	Node  string `json:"node"`
	Value int32  `json:"value"`
}

func (n CreatedNode) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", n),
		n.Node,
		n.Value,
	}
}

func (n CreatedNode) GetID() int32 {
	return n.ID
}

func (n CreatedNode) Error() string {
	return fmt.Sprintf("created node %s with value: %d", n.Node, n.Value)
}

type TargetableNode struct {
	SphereID int32
	Node     string
}

func (n TargetableNode) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", n),
		n.SphereID,
		n.Node,
	}
}
