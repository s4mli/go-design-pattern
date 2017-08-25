package decorator

type display interface {
	getColumns() int
	getRows() int
	getRowText(row int) string
	Show(display display) string // WTF
}

/*** default implementation of the interface ***/
type defaultDisplay struct {
	display
}

func (dd *defaultDisplay) Show(display display) string {
	str := ""
	for i := 0; i < display.getRows(); i++ {
		str += display.getRowText(i)
	}
	return str
}

/*** string display ***/
type StringDisplay struct {
	*defaultDisplay
	str string
}

func (sd *StringDisplay) getColumns() int {
	return len(sd.str)
}

func (sd *StringDisplay) getRows() int {
	return 1
}

func (sd *StringDisplay) getRowText(row int) string {
	if row == 0 {
		return sd.str
	} else {
		return ""
	}
}

func NewStringDisplay(str string) *StringDisplay {
	return &StringDisplay{str: str}
}

/*** real decorator starts here ***/
type border struct {
	*defaultDisplay
	display display
} // and we need this stupid struct to aggregate a display

type SideBorder struct {
	*border
	borderChar string
}

func (sb *SideBorder) getColumns() int {
	return len(sb.borderChar)*2 + sb.display.getColumns()
}

func (sb *SideBorder) getRows() int {
	return sb.display.getRows()
}

func (sb *SideBorder) getRowText(row int) string {
	return sb.borderChar + sb.display.getRowText(row) + sb.borderChar
}

func NewSideBorder(display display, borderChar string) *SideBorder {
	return &SideBorder{&border{display: display}, borderChar}
}
