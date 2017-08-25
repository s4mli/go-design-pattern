package chain_of_responsibility

import (
	"strconv"
)

type Trouble struct {
	number int
}

func (t *Trouble) getNumber() int {
	return t.number
}

/*** interfaces ***/
type support interface {
	resolve(*Trouble) bool
	Handle(support, *Trouble) string
}

type defaultSupport struct {
	support
	id   string
	next support
}

func (ds *defaultSupport) SetNext(s support) {
	ds.next = s
}

func (ds *defaultSupport) done(t *Trouble) string {
	return "Trouble : " + strconv.Itoa(t.getNumber()) + " is resolved by " + ds.id
}

func (ds *defaultSupport) fail(t *Trouble) string {
	return "Trouble : " + strconv.Itoa(t.getNumber()) + " cannot be resolved"
}

func (ds *defaultSupport) Handle(s support, t *Trouble) string {
	// what happens if the 1st para is not self
	if s.resolve(t) {
		return ds.done(t)
	} else if ds.next != nil {
		return ds.next.Handle(ds.next, t)
	} else {
		return ds.fail(t)
	}
}

/*** ***/
type noSupport struct {
	*defaultSupport
}

func (ns *noSupport) resolve(trouble *Trouble) bool {
	return false
}

func NewNoSupport(id string) *noSupport {
	return &noSupport{&defaultSupport{id: id}}
}

type limitSupport struct {
	*defaultSupport
	limit int
}

func (ls *limitSupport) resolve(trouble *Trouble) bool {
	if trouble.getNumber() < ls.limit {
		return true
	} else {
		return false
	}
}

func NewLimitSupport(id string, limit int) *limitSupport {
	return &limitSupport{&defaultSupport{id: id}, limit}
}
