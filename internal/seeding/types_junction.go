package seeding

import "slices"

type Junction interface {
	ToHashFieldsJ(string) []any
	Hashable
}

type StdJunction struct {
	ParentID int32
	ChildID  int32
}

func (j StdJunction) ToHashFields() []any {
	return []any{
		j.ParentID,
		j.ChildID,
	}
}

func (j StdJunction) ToHashFieldsJ(name string) []any {
	return slices.Concat([]any{name}, j.ToHashFields())
}

type ThreeWayJunction struct {
	GrandparentID int32
	StdJunction
}

func (j ThreeWayJunction) ToHashFields() []any {
	return []any{
		j.GrandparentID,
		j.ParentID,
		j.ChildID,
	}
}

func (j ThreeWayJunction) ToHashFieldsJ(name string) []any {
	return slices.Concat([]any{name}, j.ToHashFields())
}

type FourWayJunction struct {
	GreatGrandparentID int32
	ThreeWayJunction
}

func (j FourWayJunction) ToHashFields() []any {
	return []any{
		j.GreatGrandparentID,
		j.GrandparentID,
		j.ParentID,
		j.ChildID,
	}
}

func (j FourWayJunction) ToHashFieldsJ(name string) []any {
	return slices.Concat([]any{name}, j.ToHashFields())
}
