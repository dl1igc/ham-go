package ui

import (
	"unicode"

	termbox "github.com/nsf/termbox-go"
	"github.com/tzneal/ham-go/cmd/termlog/input"
)

type ComboBox struct {
	xPos, yPos int

	selected   int
	width      int
	focused    bool
	items      []string
	controller Controller
}

func NewComboBox(xPos, yPos int) *ComboBox {
	cb := &ComboBox{
		xPos:  xPos,
		yPos:  yPos,
		width: 4,
	}
	return cb
}

func (c *ComboBox) Redraw() {
	fg := c.controller.Theme().ComboBoxFg
	bg := c.controller.Theme().ComboBoxBg
	Clear(c.xPos, c.yPos, c.xPos+c.width-1, c.yPos, fg, bg)
	if c.selected >= 0 && c.selected < len(c.items) {
		text := c.items[c.selected]
		DrawText(c.xPos, c.yPos, text, fg, bg)
		termbox.SetCell(c.xPos+c.width, c.yPos, '∇', fg, bg)
		if c.focused {
			termbox.SetCursor(c.xPos+c.width, c.yPos)
		}
	}
}
func (c *ComboBox) Focus(b bool) {
	c.focused = b
}

func (c *ComboBox) SetController(cn Controller) {
	c.controller = cn
}
func (c *ComboBox) AddItem(t string) {
	if len(t) > c.width {
		c.width = len(t) + 1
	}
	c.items = append(c.items, t)
}

func (c *ComboBox) HandleEvent(key input.Key) {
	switch key {
	case input.KeyTab:
		c.controller.FocusNext()
	case input.KeyArrowDown:
		c.selected++
	case input.KeyArrowUp:
		c.selected--
	default:
		if (key > 'a' && key < 'z') || (key > 'A' && key < 'Z') || (key >= '0' && key <= '9') {
			for i := c.selected + 1; i != c.selected; i++ {
				i = i % len(c.items)
				if len(c.items[i]) == 0 {
					continue
				}
				if unicode.ToUpper(rune(c.items[i][0])) == unicode.ToUpper(rune(key)) {
					c.selected = i
					break
				}
			}
		}
	}
	if c.selected < 0 {
		c.selected = len(c.items) - 1
	} else if c.selected >= len(c.items) {
		c.selected = 0
	}
}

func (c *ComboBox) SetSelected(text string) {
	for i, v := range c.items {
		if v == text {
			c.selected = i
		}
	}
}

func (c *ComboBox) Value() string {
	if c.selected < 0 || c.selected > len(c.items) {
		return ""
	}
	return c.items[c.selected]
}

func (c *ComboBox) Width() int {
	return c.width
}
