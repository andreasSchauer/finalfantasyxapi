package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)


func (test Test) compareExpNameVer(exp NameVer, gotID int32, gotName string, gotVer *int32) {
	test.t.Helper()
	test.compare("id", exp.id, gotID)
	test.compare("name", exp.name, gotName)
	test.compare("version", exp.version, gotVer)
}

func (test Test) compareExpUnique(exp Unique, gotID int32, gotName string) {
	test.t.Helper()
	test.compare("id", exp.id, gotID)
	test.compare("name", exp.name, gotName)
}

func (test Test) compareExpIDOnly(exp IDOnly, gotID int32) {
	test.t.Helper()
	test.compare("id", exp.id, gotID)
}

func (test Test) compareListParams(exp, got apiresource.ListParams) {
	test.t.Helper()
	test.compare("count", exp.Count, got.Count)
	test.compPageURL("previous", exp.Previous, got.Previous)
	test.compPageURL("next", exp.Next, got.Next)
}
