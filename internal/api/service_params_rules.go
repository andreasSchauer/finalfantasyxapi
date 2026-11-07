package api

import ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"

type ParamsRules map[ts.FieldName]FieldRules

type FieldRules struct {
	Doc      FieldDoc
	Children ParamsRules
}

func (d ParamsDoc) CompileRules() ParamsRules {
	return compileParamsRules(d.Fields)
}

func compileParamsRules(fields []FieldDoc) ParamsRules {
	if len(fields) == 0 {
		return nil
	}

	paramsRules := make(ParamsRules, len(fields))

	for _, f := range fields {
		paramsRules[f.Field] = FieldRules{
			Doc:      f,
			Children: compileParamsRules(f.ChildProps),
		}
	}

	return paramsRules
}
