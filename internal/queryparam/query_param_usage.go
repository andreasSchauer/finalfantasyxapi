package queryparam

import (
	"fmt"
	"strings"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (q *QueryParam) AssignUsage() {
	s := fmt.Sprintf("?%s=", q.Name)

	switch q.Type {
	case ts.QptBool:
		q.Usage = s + "{bool}"
		q.ExampleUses = []string{s + "true", s + "false"}
		return

	case ts.QptEnum:
		e := q.EnumVals[0].Name
		q.Usage = s + "{value|id}"
		q.ExampleUses = []string{s + "1", s + e}
		return

	case ts.QptEnumList:
		e1 := q.EnumVals[0].Name
		e2 := q.EnumVals[1].Name
		q.Usage = s + "{value|id},..."
		q.ExampleUses = []string{s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2)}
		return

	case ts.QptID:
		q.Usage = s + "{id}"
		q.ExampleUses = []string{s + "1"}
		return

	case ts.QptIDNul:
		q.Usage = s + "{id|'none'}"
		q.ExampleUses = []string{s + "1", s + "none"}
		return

	case ts.QptIDList:
		q.Usage = s + "{id},..."
		q.ExampleUses = []string{s + "1", s + "1,2"}
		return

	case ts.QptInt:
		q.Usage = s + "{int}"
		q.ExampleUses = []string{s + "1"}
		return

	case ts.QptIntList:
		q.Usage = s + "{int},...|{int}-{int}"
		q.ExampleUses = []string{s + "1", s + "1,2", s + "1-3", s + "1,2-4"}
		return

	case ts.QptNameID:
		e := q.ExampleVals[0]
		q.Usage = s + "{name|id}"
		q.ExampleUses = []string{s + "1", s + e}
		return

	case ts.QptNameIDList:
		e1 := q.ExampleVals[0]
		e2 := q.ExampleVals[1]
		q.Usage = s + "{name|id},..."
		q.ExampleUses = []string{s + "1", s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2)}
		return

	case ts.QptNameIDListNul:
		e1 := q.ExampleVals[0]
		e2 := q.ExampleVals[1]
		q.Usage = s + "{name|id},...|{'none'}"
		q.ExampleUses = []string{s + "1", s + "1,2", s + fmt.Sprintf("%s,%s", e1, e2), s + "none"}
		return

	case ts.QptValue:
		e1 := string(q.AllowedValues[0])
		e2 := string(q.AllowedValues[1])
		q.Usage = s + "{val}"
		q.ExampleUses = []string{s + e1, s + e2}
		return

	case ts.QptValueList:
		e1 := string(q.AllowedValues[0])
		e2 := string(q.AllowedValues[1])
		q.Usage = s + "{val}"
		q.ExampleUses = []string{s + e1, s + fmt.Sprintf("%s,%s", e1, e2)}
		return

	case ts.QptStat:
		q.Usage = s + "{stat}:{int},..."
		return
	}

	if q.SpecialInputs != nil {
		for _, input := range q.SpecialInputs {
			usageTrimmed := strings.TrimSuffix(q.Usage, "}")
			q.Usage = fmt.Sprintf("%s|'%s'}", usageTrimmed, input.Key)
			q.ExampleUses = append(q.ExampleUses, s+string(input.Key))
		}
	}
}

func exampleUse(queryName ts.QueryParamName, value string) string {
	return fmt.Sprintf("?%s=%s", queryName, value)
}
