package api

type ValidationTree map[FieldName]FieldRules

type FieldRules struct {
	Doc      FieldDoc
	Children ValidationTree
}

func compileValidationTree(fields []FieldDoc) ValidationTree {
	if len(fields) == 0 {
		return nil
	}

	tree := make(ValidationTree, len(fields))

	for _, f := range fields {
		tree[f.Field] = FieldRules{
			Doc:      f,
			Children: compileValidationTree(f.ChildProps),
		}
	}

	return tree
}
