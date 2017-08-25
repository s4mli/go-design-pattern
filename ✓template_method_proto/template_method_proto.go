package template_method_proto

type DisplayTemplateFn interface {
	open() string
	print() string
	close() string
}

type Display interface {
	DisplayTemplateFn
	Display() string
}

type TemplateDisplay struct {
	openFn  func() string
	printFn func() string
	closeFn func() string
}

func (td *TemplateDisplay) open() string {
	if td.openFn != nil {
		return td.openFn()
	} else {
		return ""
	}
}

func (td *TemplateDisplay) print() string {
	if td.printFn != nil {
		return td.printFn()
	} else {
		return ""
	}
}

func (td *TemplateDisplay) close() string {
	if td.closeFn != nil {
		return td.closeFn()
	} else {
		return ""
	}
}

func (td *TemplateDisplay) Display() string {
	result := td.open()
	for i := 0; i < 5; i++ {
		result += "< " + td.print() + " >"
	}
	result += td.close()
	return result
}

func newDisplay(openFn, printFn, closeFn func() string) Display {
	return &TemplateDisplay{openFn, printFn, closeFn}
}

/*** I want to override all ***/
type CharDisplay struct {
	Char rune
}

func (cd *CharDisplay) open() string {
	return "<<"
}
func (cd *CharDisplay) print() string {
	return string(cd.Char)
}
func (cd *CharDisplay) close() string {
	return ">>"
}

func NewCharDisplay(c rune) Display {
	cdm := CharDisplay{c}
	return newDisplay(cdm.open, cdm.print, cdm.close)
}

/*** I don't want to override close ***/
type CharDisplayWithoutClose struct {
	Char rune
}

func (cdmwc *CharDisplayWithoutClose) open() string {
	return "<<"
}
func (cdmwc *CharDisplayWithoutClose) print() string {
	return string(cdmwc.Char)
}

func NewCharDisplayWithoutClose(c rune) Display {
	cdmwc := CharDisplayWithoutClose{c}
	return newDisplay(cdmwc.open, cdmwc.print, nil)
}
