package logger

import "github.com/fatih/color"

type LogColor *color.Color

var (
	DebugColor LogColor = color.New(color.FgMagenta)
	InfoColor  LogColor = color.New(color.FgBlue)
	WarnColor  LogColor = color.New(color.FgYellow)
	ErrorColor LogColor = color.New(color.FgRed)
	FatalColor LogColor = color.New(color.FgRed, color.Bold)
)
