package input

import (
	"github.com/charmbracelet/lipgloss"
	"time"
)

type Option func(i *Input)

// WithPrompt set the prompt
func WithPrompt(prompt string) Option {
	return func(i *Input) {
		i.inner.Prompt = prompt
	}
}

// WithPlaceholder setthe placeholder
func WithPlaceholder(placeholder string) Option {
	return func(i *Input) {
		i.inner.Placeholder = placeholder
	}
}

// WithBlinkSpeed setthe blink speed
func WithBlinkSpeed(blinkSpeed time.Duration) Option {
	return func(i *Input) {
		i.inner.BlinkSpeed = blinkSpeed
	}
}

// WithEchoMode sets the input behavior of the text input field.
func WithEchoMode(echoMode EchoMode) Option {
	return func(i *Input) {
		i.inner.EchoMode = echoMode
	}
}

// WithEchoCharacter setthe echo char shape
func WithEchoCharacter(echoCharacter rune) Option {
	return func(i *Input) {
		i.inner.EchoCharacter = echoCharacter
	}
}

// WithPromptStyle setthe prompt style
func WithPromptStyle(style lipgloss.Style) Option {
	return func(i *Input) {
		i.inner.PromptStyle = style
	}
}

// WithTextStyle setthe text style
func WithTextStyle(style lipgloss.Style) Option {
	return func(i *Input) {
		i.inner.TextStyle = style
	}
}

// WithBackgroundStyle setthe background style
func WithBackgroundStyle(style lipgloss.Style) Option {
	return func(i *Input) {
		i.inner.BackgroundStyle = style
	}
}

// WithPlaceholderStyle set the placeholder style
func WithPlaceholderStyle(style lipgloss.Style) Option {
	return func(i *Input) {
		i.inner.PlaceholderStyle = style
	}
}

// WithCursorStyle setthe cursor style
func WithCursorStyle(style lipgloss.Style) Option {
	return func(i *Input) {
		i.inner.CursorStyle = style
	}
}

// WithCharLimit is the maximum amount of characters this input element will
// accept. If 0 or less, there's no limit.
func WithCharLimit(charLimit int) Option {
	return func(i *Input) {
		i.inner.CharLimit = charLimit
	}
}
