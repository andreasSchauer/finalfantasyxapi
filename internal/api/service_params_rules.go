package api

type ParamsRules map[FieldName]FieldRules

type FieldRules struct {
	Doc      FieldDoc
	Children ParamsRules
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
