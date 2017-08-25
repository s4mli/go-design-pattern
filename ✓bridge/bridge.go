package bridge

import "strconv"

type DisplayImp interface {
	rawOpen(int) string
	rawClose(int) string
	rawPrint(int) string
}

/*** String imp ***/
type StringDisplayImp struct {
}

func (sdi *StringDisplayImp) rawOpen(what int) string {
	return sdi.printLine(what)
}

func (sdi *StringDisplayImp) rawPrint(what int) string {
	return "|" + strconv.Itoa(what) + "|\n"
}

func (sdi *StringDisplayImp) rawClose(what int) string {
	return sdi.printLine(what)
}

func (sdi *StringDisplayImp) printLine(what int) string {
	str := "+"
	for _, _ = range strconv.Itoa(what) {
		str += "-"
	}
	str += "+\n"
	return str
}

/*** Integer imp ***/
type IntegerDisplayImp struct {
}

func (idi *IntegerDisplayImp) rawOpen(what int) string {
	return idi.printLine(what)
}

func (idi *IntegerDisplayImp) rawPrint(what int) string {
	return "|" + strconv.Itoa(what) + "|\n"
}

func (idi *IntegerDisplayImp) rawClose(what int) string {
	return idi.printLine(what)
}

func (idi *IntegerDisplayImp) printLine(what int) string {
	str := "+"
	for i := 0; i < what; i++ {
		str += "-"
	}
	str += "+\n"
	return str
}

/*** Default display using imp ***/
type DisplayHelper struct {
	imp  DisplayImp
	what int
}

func (dh *DisplayHelper) open() string {
	return dh.imp.rawOpen(dh.what)
}

func (dh *DisplayHelper) print() string {
	return dh.imp.rawPrint(dh.what)
}

func (dh *DisplayHelper) close() string {
	return dh.imp.rawClose(dh.what)
}

func (dh *DisplayHelper) Display() string {
	return dh.open() + dh.print() + dh.close()
}

/*** Count display ***/
type CountDisplay struct {
	*DisplayHelper
}

func (cd *CountDisplay) MultiDisplay(num int) string {
	str := cd.open()
	for i := 0; i < num; i++ {
		str += cd.print()
	}
	str += cd.close()
	return str
}

func NewDefaultDisplayWithImp(what int, imp DisplayImp) *DisplayHelper {
	return &DisplayHelper{imp, what}
}

func NewCountDisplayWithImp(what int, imp DisplayImp) *CountDisplay {
	return &CountDisplay{NewDefaultDisplayWithImp(what, imp)}
}
