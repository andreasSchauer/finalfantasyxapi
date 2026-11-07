// Package queryparam aspdijasdpij
package queryparam

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type ParamUseType string

const (
	PuSingle ParamUseType = "single"
	PuList   ParamUseType = "list"
	PuSimple ParamUseType = "simple"
)

type QueryParam struct {
	ID                 int                 `json:"-"`
	Name               ts.QueryParamName   `json:"name"`
	Type               ts.QueryParamType   `json:"param_type"`
	Endpoint           ts.EndpointName     `json:"-"`
	Description        string              `json:"description"`
	ExampleVals        []string            `json:"-"`
	Usage              string              `json:"usage"`
	ExampleUses        []string            `json:"example_uses"`
	IsExclusive        bool                `json:"only_use_alone,omitempty"`
	ParamUse           ParamUseType        `json:"param_use_type"`
	ForSegment         *ts.SectionName     `json:"for_segment,omitempty"`
	EnumVals           []enums.EnumVal     `json:"-"`
	RequiresAll        []ts.QueryParamName `json:"required_params,omitempty"`
	RequiresOne        []ts.QueryParamName `json:"usable_with,omitempty"`
	ReplacedBy         []ts.QueryParamName `json:"replaced_by,omitempty"`
	ConflictsWith      []ts.QueryParamName `json:"conflicts_with,omitempty"`
	ReferencesInt      []ts.EndpointName   `json:"-"`
	ReferencesEnumsInt []ts.EnumName       `json:"-"`
	References         []string            `json:"references,omitempty"`
	ReferencesEnums    []string            `json:"references_enums,omitempty"`
	AllowedIDs         []int32             `json:"-"`
	AllowedResources   []string            `json:"allowed_resources,omitempty"`
	AllowedValues      []ts.QueryValue     `json:"allowed_values,omitempty"`
	AllowedIntRange    []int               `json:"allowed_int_range,omitempty"`
	AllowedResTypes    []string            `json:"allowed_res_types,omitempty"`
	DefaultVal         *int                `json:"default_value,omitempty"`
	SpecialInputs      []SpecialQueryInput `json:"special_inputs,omitempty"`
}

type SpecialQueryInput struct {
	Key ts.QuerySpecialVal `json:"key"`
	Val int                `json:"value"`
}

func (q QueryParam) IsZero() bool {
	return q.ID == 0 && q.Name == ""
}