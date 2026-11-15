package seedhandlers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop3Spheres(qtx *database.Queries, ctx context.Context) error {
	spheres, err := extractSpheres()
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
		params.DataHash[i] = sh.GenerateDataHash(s)
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
		sh.RawJSON.Spheres[i].ID = row.ID
		lookup.L.Spheres[seedtypes.Key(spheres[i])] = spheres[i]
		lookup.L.SpheresID[row.ID] = spheres[i]
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractSpheres() ([]seedtypes.Sphere, error) {
	spheres := []seedtypes.Sphere{}
	var err error

	for i := range sh.RawJSON.Spheres {
		sphere := &sh.RawJSON.Spheres[i]

		sphere.ItemID, err = sh.AssignFK(sphere.Name, lookup.L.Items)
		if err != nil {
			return nil, err
		}

		if sphere.CreatedNode != nil {
			sphere.CreatedNode.ID, err = sh.GetHashID(*sphere.CreatedNode)
			if err != nil {
				return nil, err
			}
		}

		spheres = append(spheres, *sphere)
	}

	return sh.DedupeRows(spheres, sh.Hashes), nil
}

func Loop1CreatedNodes(qtx *database.Queries, ctx context.Context) error {
	nodes := extractCreatedNodes()

	params := database.CreateCreatedNodeBulkParams{
		DataHash: make([]string, len(nodes)),
		Node:     make([]database.NodeType, len(nodes)),
		Value:    make([]int32, len(nodes)),
	}

	for i, mi := range nodes {
		params.DataHash[i] = sh.GenerateDataHash(mi)
		params.Node[i] = database.NodeType(mi.Node)
		params.Value[i] = mi.Value
	}

	dbRows, err := qtx.CreateCreatedNodeBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create created nodes: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractCreatedNodes() []seedtypes.CreatedNode {
	createdNodes := []seedtypes.CreatedNode{}

	for _, sphere := range sh.RawJSON.Spheres {
		if sphere.CreatedNode != nil {
			createdNodes = append(createdNodes, *sphere.CreatedNode)
		}
	}

	return sh.DedupeRows(createdNodes, sh.Hashes)
}

func Loop4TargetableNodes(qtx *database.Queries, ctx context.Context) error {
	nodes := extractTargetableNodes()

	params := database.CreateSphereTargetableNodeBulkParams{
		DataHash: make([]string, len(nodes)),
		SphereID: make([]int32, len(nodes)),
		Node:     make([]database.NodeType, len(nodes)),
	}

	for i, s := range nodes {
		params.DataHash[i] = sh.GenerateDataHash(s)
		params.SphereID[i] = s.SphereID
		params.Node[i] = database.NodeType(s.Node)
	}

	dbRows, err := qtx.CreateSphereTargetableNodeBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create targetable nodes: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractTargetableNodes() []seedtypes.TargetableNode {
	nodes := []seedtypes.TargetableNode{}

	for i := range sh.RawJSON.Spheres {
		sphere := &sh.RawJSON.Spheres[i]

		for j := range sphere.TargetableNodes {
			nodeString := sphere.TargetableNodes[j]

			node := seedtypes.TargetableNode{
				SphereID: sphere.ID,
				Node:     nodeString,
			}

			nodes = append(nodes, node)
		}
	}

	return sh.DedupeRows(nodes, sh.Hashes)
}
