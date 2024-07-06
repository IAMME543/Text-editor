package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type CustomEntry struct {
	widget.Entry
}

func NewCustomEntry() *CustomEntry {
	entry := &CustomEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *CustomEntry) TypedShortcut(s fyne.Shortcut) {
	if shortcut, ok := s.(*desktop.CustomShortcut); ok {
		if shortcut.KeyName == fyne.KeyBackspace && shortcut.Modifier == fyne.KeyModifierControl {
			e.deleteWordBeforeCursor()
			return
		}
	}
}

func (e *CustomEntry) deleteWordBeforeCursor() {
	pos := e.CursorColumn
	if pos == 0 {
		return
	}
	text := e.Text[:pos]
	lastSpace := strings.LastIndex(text, " ")
	if lastSpace == -1 {
		e.SetText(e.Text[pos:])
		e.CursorColumn = 0
	} else {
		e.SetText(e.Text[:lastSpace] + e.Text[pos:])
		e.CursorColumn = lastSpace
	}
	e.Refresh()
}
