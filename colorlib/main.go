package colorlib

import (
	"fmt"
	"github.com/fatih/color"
)

type RGB struct {
	R int
	G int
	B int
}

type LogLine struct {
	Text        string
	BgRGB       RGB
	FontRGB     RGB
	WithNewLine bool
	WithPadding bool
	Padding     int
}

func SprintlnP(text string, bgRGB RGB, padding int) *LogLine {
	return &LogLine{
		Text:        text,
		BgRGB:       bgRGB,
		WithNewLine: true,
		WithPadding: true,
		Padding:     padding,
	}
}

func Sprintln(text string, bgRGB RGB) *LogLine {
	return &LogLine{
		Text:        text,
		BgRGB:       bgRGB,
		WithNewLine: true,
		WithPadding: false,
	}
}
func PrintP(text string, bgRGB RGB, padding int) *LogLine {
	return &LogLine{
		Text:        text,
		BgRGB:       bgRGB,
		WithNewLine: false,
		WithPadding: true,
		Padding:     padding,
	}
}

func Print(text string, bgRGB RGB) *LogLine {
	return &LogLine{
		Text:        text,
		BgRGB:       bgRGB,
		WithNewLine: false,
		WithPadding: false,
	}
}

func (l *LogLine) GetString() string {

	if l.FontRGB == (RGB{}) {
		l.FontRGB = RGB{255, 255, 255}
	}

	str := color.RGB(l.FontRGB.R, l.FontRGB.G, l.FontRGB.B).AddBgRGB(l.BgRGB.R, l.BgRGB.G, l.BgRGB.B).Sprintf(l.Text)
	if l.WithPadding {
		str = PadWithOverride(str, l.Padding, len(l.Text))
	}

	if l.WithNewLine {
		str += "\n"
	}
	return str
}

func Pad(s string, width int) string {
	// width is the full width of the screen
	// we want s to be centered in the screen
	// so we need to calculate the padding on the left
	padding := (width - len(s)) / 2
	return fmt.Sprintf("%*s%s", padding, " ", s)
}

func PadWithOverride(s string, width int, lenOverride int) string {
	// width is the full width of the screen
	// we want s to be centered in the screen
	// so we need to calculate the padding on the left
	padding := (width - lenOverride) / 2
	return fmt.Sprintf("%*s%s", padding, " ", s)
}
