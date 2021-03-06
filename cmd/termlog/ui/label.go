package ui

import termbox "github.com/nsf/termbox-go"

type Label struct {
	xPos, yPos int
	width      int
	text       string
	bg         termbox.Attribute
	fg         termbox.Attribute
}

func NewLabel(x, y int, text string) *Label {
	return &Label{
		xPos:  x,
		yPos:  y,
		width: len(text),
		bg:    termbox.ColorDefault,
		fg:    termbox.ColorWhite,
		text:  text,
	}
}
func (l *Label) SetController(c Controller) {

}
func (l *Label) Redraw() {
	Clear(l.xPos, l.yPos, l.xPos+l.width, l.yPos, l.fg, l.bg)
	DrawText(l.xPos, l.yPos, l.text, l.fg, l.bg)
}
func (l *Label) SetText(text string) {
	l.text = text
}

func (l *Label) GetText() string {
	return l.text
}
