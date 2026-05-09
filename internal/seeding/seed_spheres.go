package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func (l *Lookup) loop3SeedSpheres(qtx *database.Queries, ctx context.Context) error {
	spheres, err := l.extractSpheres()
	if err != nil {
		return err
	}

	params := database.CreateSphereBulkParams{
		DataHash:              make([]string, len(spheres)),
		ItemID:                make([]int32, len(spheres)),
		SphereGridDescription: make([]string, len(spheres)),
		SphereColor:           make([]database.SphereColor, len(spheres)),
		SphereEffect:          make([]database.SphereEffect, len(spheres)),
		TargetNodePosition:    make([]database.NodePosition, len(spheres)),
		TargetNodeState:       make([]database.NullNodeState, len(spheres)),
		CreatedNodeID:         make([]sql.NullInt32, len(spheres)),
	}

	for i, s := range spheres {
		params.DataHash[i] = generateDataHash(s)
		params.ItemID[i] = s.ItemID
		params.SphereGridDescription[i] = s.SgDescription
		params.SphereColor[i] = database.SphereColor(s.SphereColor)
		params.SphereEffect[i] = database.SphereEffect(s.SphereEffect)
		params.TargetNodePosition[i] = database.NodePosition(s.TargetNodePosition)
		params.TargetNodeState[i] = database.ToNullNodeState(s.TargetNodeState)
		params.CreatedNodeID[i] = h.ObjPtrToNullInt32ID(s.CreatedNode)
	}

	dbRows, err := qtx.CreateSphereBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create spheres: %v", err)
	}

	for i, row := range dbRows {
		spheres[i].ID = row.ID
		l.json.spheres[i].ID = row.ID
		l.Spheres[Key(spheres[i])] = spheres[i]
		l.SpheresID[row.ID] = spheres[i]
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractSpheres() ([]Sphere, error) {
	spheres := []Sphere{}
	var err error

	for i := range l.json.spheres {
		sphere := &l.json.spheres[i]

		sphere.ItemID, err = assignFK(sphere.Name, l.Items)
		if err != nil {
			return nil, err
		}

		if sphere.CreatedNode != nil {
			sphere.CreatedNode.ID, err = l.GetHashID(*sphere.CreatedNode)
			if err != nil {
				return nil, err
			}
		}

		spheres = append(spheres, *sphere)
	}

	return dedupeRows(spheres, l.Hashes), nil
}

func (l *Lookup) loop1SeedCreatedNodes(qtx *database.Queries, ctx context.Context) error {
	nodes := l.extractCreatedNodes()

	params := database.CreateCreatedNodeBulkParams{
		DataHash: make([]string, len(nodes)),
		Node:     make([]database.NodeType, len(nodes)),
		Value:    make([]int32, len(nodes)),
	}

	for i, mi := range nodes {
		params.DataHash[i] = generateDataHash(mi)
		params.Node[i] = database.NodeType(mi.Node)
		params.Value[i] = mi.Value
	}

	dbRows, err := qtx.CreateCreatedNodeBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create created nodes: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractCreatedNodes() []CreatedNode {
	createdNodes := []CreatedNode{}

	for _, sphere := range l.json.spheres {
		if sphere.CreatedNode != nil {
			createdNodes = append(createdNodes, *sphere.CreatedNode)
		}
	}

	return dedupeRows(createdNodes, l.Hashes)
}

func (l *Lookup) loop4SeedTargetableNodes(qtx *database.Queries, ctx context.Context) error {
	nodes := l.extractTargetableNodes()

	params := database.CreateSphereTargetableNodeBulkParams{
		DataHash: make([]string, len(nodes)),
		SphereID: make([]int32, len(nodes)),
		Node:     make([]database.NodeType, len(nodes)),
	}

	for i, s := range nodes {
		params.DataHash[i] = generateDataHash(s)
		params.SphereID[i] = s.SphereID
		params.Node[i] = database.NodeType(s.Node)
	}

	dbRows, err := qtx.CreateSphereTargetableNodeBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create targetable nodes: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractTargetableNodes() []TargetableNode {
	nodes := []TargetableNode{}

	for i := range l.json.spheres {
		sphere := &l.json.spheres[i]

		for j := range sphere.TargetableNodes {
			nodeString := sphere.TargetableNodes[j]

			node := TargetableNode{
				SphereID: sphere.ID,
				Node:     nodeString,
			}

			nodes = append(nodes, node)
		}
	}

	return dedupeRows(nodes, l.Hashes)
}
