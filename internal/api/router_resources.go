package api


import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)


type routerResources interface {
	getSegments() []string
	list()
	single()
	subsections()
	formatError()
}

func route(r routerResources) {
	switch len(r.getSegments()) {
	case 0:
		r.list()
		return

	case 1:
		r.single()
		return

	case 2:
		r.subsections()
		return

	default:
		r.formatError()
		return
	}
}


type routerIDOnly[T seedtypes.Lookupable, R any, A apiresource.APIResource] struct {
	routerParams
	i handlerinput.Main[T, R, A]
}

func newRouterIDOnly[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.Main[T, R, A]) routerIDOnly[T, R, A] {
	i.Usage = []string{"/{id}"}

	return routerIDOnly[T, R, A]{
		routerParams: newRouterParams(w, r, i.Endpoint),
		i:            i,
	}
}

func (r routerIDOnly[T, R, A]) route() {
	route(r)
}

func (r routerIDOnly[T, R, A]) getSegments() []string {
	return r.segments
}

func (r routerIDOnly[T, R, A]) list() {
	handleList(r.w, r.r, r.i)
}

func (r routerIDOnly[T, R, A]) single() {
	handleIDOnly(r.w, r.r, r.i, r.segments[0])
}

func (r routerIDOnly[T, R, A]) subsections() {
	handleSubsections(r.w, r.r, r.i, r.segments[0], r.segments[1])
}

func (r routerIDOnly[T, R, A]) formatError() {
	formatError(r.w, r.i)
}



type routerNameID[T seedtypes.Lookupable, R any, A apiresource.APIResource] struct {
	routerParams
	i handlerinput.Main[T, R, A]
}

func newRouterNameID[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.Main[T, R, A]) routerNameID[T, R, A] {
	i.Usage = []string{"/{id}", "/{name}"}

	return routerNameID[T, R, A]{
		routerParams: newRouterParams(w, r, i.Endpoint),
		i:            i,
	}
}

func (r routerNameID[T, R, A]) route() {
	route(r)
}

func (r routerNameID[T, R, A]) getSegments() []string {
	return r.segments
}

func (r routerNameID[T, R, A]) list() {
	handleList(r.w, r.r, r.i)
}

func (r routerNameID[T, R, A]) single() {
	handleNameID(r.w, r.r, r.i, r.segments[0])
}

func (r routerNameID[T, R, A]) subsections() {
	handleSubsections(r.w, r.r, r.i, r.segments[0], r.segments[1])
}

func (r routerNameID[T, R, A]) formatError() {
	formatError(r.w, r.i)
}






type routerNameVersion[T seedtypes.Lookupable, R any, A apiresource.APIResource] struct {
	routerParams
	i handlerinput.Main[T, R, A]
}

func newRouterNameVersion[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.Main[T, R, A]) routerNameVersion[T, R, A] {
	i.Usage = []string{"/{id}", "/{name}", "/{name}/{version}"}

	return routerNameVersion[T, R, A]{
		routerParams: newRouterParams(w, r, i.Endpoint),
		i:            i,
	}
}

func (r routerNameVersion[T, R, A]) route() {
	route(r)
}

func (r routerNameVersion[T, R, A]) getSegments() []string {
	return r.segments
}

func (r routerNameVersion[T, R, A]) list() {
	handleList(r.w, r.r, r.i)
}

func (r routerNameVersion[T, R, A]) single() {
	handleNameID(r.w, r.r, r.i, r.segments[0])
}

func (r routerNameVersion[T, R, A]) subsections() {
	handleSubOrNameVer(r.w, r.r, r.i, r.segments[0], r.segments[1])
}

func (r routerNameVersion[T, R, A]) formatError() {
	formatError(r.w, r.i)
}