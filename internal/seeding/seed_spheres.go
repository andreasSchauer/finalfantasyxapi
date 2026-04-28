package seeding

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Sphere struct {
	ID 			  		int32
	ItemID     	  		int32
	Name          		string 			`json:"name"`
	SgDescription		string			`json:"sphere_grid_description"`
	SphereColor			string			`json:"sphere_color"`
	SphereEffect		string			`json:"sphere_effect"`
	TargetNodePosition	string			`json:"target_node_position"`
	TargetNodeState		*string			`json:"target_node_state"`
	TargetableNodes		[]string		`json:"targetable_nodes"`
	CreatedNode			*CreatedNode	`json:"created_node"`
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
		ID: 	s.ID,
		Name: 	s.Name,
	}
}

type CreatedNode struct {
	ID		int32
	Node	string	`json:"node"`
	Value	int32	`json:"value"`
}

func (n CreatedNode) ToHashFields() []any{
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
	SphereID	int32
	Node		string
}

func (n TargetableNode) ToHashFields() []any{
	return []any{
		fmt.Sprintf("%T", n),
		n.SphereID,
		n.Node,
	}
}

func (l *Lookup) seedSpheres(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/spheres.json"

	var spheres []Sphere
	err := loadJSONFile(string(srcPath), &spheres)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, sphere := range spheres {
			var err error

			sphere.ItemID, err = assignFK(sphere.Name, l.Items)
			if err != nil {
				return h.NewErr(sphere.Error(), err)
			}

			dbSphere, err := qtx.CreateSphere(context.Background(), database.CreateSphereParams{
				DataHash:      			generateDataHash(sphere),
				ItemID: 				sphere.ItemID,
				SphereGridDescription: 	sphere.SgDescription,
				SphereColor: 			database.SphereColor(sphere.SphereColor),
				SphereEffect: 			database.SphereEffect(sphere.SphereEffect),
				TargetNodePosition: 	database.NodePosition(sphere.TargetNodePosition),
				TargetNodeState: 		database.ToNullNodeState(sphere.TargetNodeState),
			})
			if err != nil {
				return h.NewErr(sphere.Error(), err, "couldn't create sphere")
			}

			sphere.ID = dbSphere.ID
			l.Spheres[sphere.Name] = sphere
			l.SpheresID[sphere.ID] = sphere
		}

		return nil
	})
}


func (l *Lookup) seedSpheresRelationships(db *database.Queries, dbConn *sql.DB) error {
	const srcPath = "data/spheres.json"

	var spheres []Sphere
	err := loadJSONFile(string(srcPath), &spheres)
	if err != nil {
		return err
	}

	return queryInTransaction(db, dbConn, func(qtx *database.Queries) error {
		for _, jsonSphere := range spheres {
			sphere, err := GetResource(jsonSphere.Name, l.Spheres)
			if err != nil {
				return err
			}

			sphere.CreatedNode, err = seedObjPtrAssignFK(qtx, sphere.CreatedNode, l.seedCreatedNode)
			if err != nil {
				return h.NewErr(sphere.Error(), err)
			}

			err = qtx.UpdateSphere(context.Background(), database.UpdateSphereParams{
				DataHash: 		generateDataHash(sphere),
				CreatedNodeID: 	h.ObjPtrToNullInt32ID(sphere.CreatedNode),
				ID: 			sphere.ID,
			})
			
			err = l.seedSphereTargetableNodes(qtx, sphere)
			if err != nil {
				return h.NewErr(sphere.Error(), err)
			}
		}

		return nil
	})
}

func (l *Lookup) seedCreatedNode(qtx *database.Queries, node CreatedNode) (CreatedNode, error) {
	dbNode, err := qtx.CreateCreatedNode(context.Background(), database.CreateCreatedNodeParams{
		DataHash: generateDataHash(node),
		Node: database.NodeType(node.Node),
		Value: node.Value,
	})
	if err != nil {
		return CreatedNode{}, h.NewErr(node.Error(), err)
	}

	node.ID = dbNode.ID

	return node, nil
}

func (l *Lookup) seedSphereTargetableNodes(qtx *database.Queries, sphere Sphere) error {
	for _, jsonNode := range sphere.TargetableNodes {
		node := TargetableNode{
			SphereID: 	sphere.ID,
			Node: 		jsonNode,
		}

		err := qtx.CreateSphereTargetableNode(context.Background(), database.CreateSphereTargetableNodeParams{
			DataHash: generateDataHash(node),
			SphereID:   node.SphereID,
			Node:   	database.NodeType(node.Node),
		})
		if err != nil {
			return h.NewErr(jsonNode, err, "couldn't junction targetable node")
		}
	}

	return nil
}


func (l *Lookup) loop3SeedSpheres(qtx *database.Queries, ctx context.Context) error {
	spheres, err := l.extractSpheres()
	if err != nil {
		return err
	}

	params := database.CreateSphereBulkParams{
		DataHash:   			make([]string, len(spheres)),
		ItemID: 				make([]int32, len(spheres)),
		SphereGridDescription: 	make([]string, len(spheres)),
		SphereColor: 			make([]database.SphereColor, len(spheres)),
		SphereEffect: 			make([]database.SphereEffect, len(spheres)),
		TargetNodePosition: 	make([]database.NodePosition, len(spheres)),
		TargetNodeState: 		make([]database.NullNodeState, len(spheres)),
		CreatedNodeID: 			make([]sql.NullInt32, len(spheres)),
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
		l.Spheres[spheres[i].Name] = spheres[i]
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
			sphere.CreatedNode.ID, err = l.getHashID(*sphere.CreatedNode)
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