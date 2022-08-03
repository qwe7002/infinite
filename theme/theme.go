package theme

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/style"
)

type Theme struct {
	PromptStyle                  lipgloss.Style
	MultiSelectedHintSymbolStyle lipgloss.Style
	ChoiceTextStyle              lipgloss.Style
	CursorSymbolStyle            lipgloss.Style
	UnHintSymbolStyle            lipgloss.Style
}

var (
	DefaultTheme = Theme{
		PromptStyle:                  style.New().Foreground(color.Cyan),
		MultiSelectedHintSymbolStyle: style.New().Foreground(color.Special),
		ChoiceTextStyle:              style.New().Foreground(color.Highlight).Bold(true),
		CursorSymbolStyle:            style.New(),
		UnHintSymbolStyle:            style.New().Foreground(color.Red),
	}
)