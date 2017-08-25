package template_method

type printer interface {
	open() string
	print() string
	close() string
}

/*** parent class - TM_Display defines template method Display ***/
type TM_Display struct{}

// Im a template method with a horrible para
func (tmd *TM_Display) Display(p printer) string {
	result := p.open()
	for i := 0; i < 5; i++ {
		result += "< " + p.print() + " >"
	}
	result += p.close()
	return result
}

/*** child class - CharDisplay with template method Display ***/
type CharDisplay struct {
	*TM_Display // magic here - get the interfaces from TM_Display
	Char        rune
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

func NewCharDisplay(c rune) *CharDisplay {
	return &CharDisplay{&TM_Display{}, c}
}

/*** child class - StringDisplay with template method Display ***/
type StringDisplay struct {
	*TM_Display // Yo
	Str         string
}

func (sd *StringDisplay) open() string {
	return sd.printLine()
}
func (sd *StringDisplay) print() string {
	return "| " + sd.Str + " |\n"
}
func (sd *StringDisplay) close() string {
	return sd.printLine()
}

func (sd *StringDisplay) printLine() string {
	line := "+-"
	for _, _ = range sd.Str {
		line += "-"
	}
	line += "-+\n"
	return line
}

func NewStringDisplay(s string) *StringDisplay {
	return &StringDisplay{&TM_Display{}, s}
}
