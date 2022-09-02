package gocolors

package utils

import (
	"fmt"
)

func RGB(r, g, b int) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

// Styles
const (
	Bold = "\033[1m"
	Dim = "\033[2m"
	Italic = "\033[3m"
	Underline = "\033[4m"
	Overline = "\033[53m"
	Inverse = "\033[7m"
	Hidden = "\033[8m"
	Strikethrough = "\033[9m"
)

// Colors
const (
	Black = "\033[30m"
	Red = "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Blue = "\033[34m"
	Magenta = "\033[35m"
	Cyan = "\033[36m"
	White = "\033[37m"
)

// Bright colors
const (
	BlackBright = "\033[90m"
	RedBright = "\033[91m"
	GreenBright = "\033[92m"
	YellowBright = "\033[93m"
	BlueBright = "\033[94m"
	MagentaBright = "\033[95m"
	CyanBright = "\033[96m"
	WhiteBright = "\033[97m"
)

// Background colors
const (
	BgBlack = "\033[40m"
	BgRed = "\033[41m"
	BgGreen = "\033[42m"
	BgYellow = "\033[43m"
	BgBlue = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan = "\033[46m"
	BgWhite = "\033[47m"
)

// Bright background colors
const (
	BgBlackBright = "\033[100m"
	BgRedBright = "\033[101m"
	BgGreenBright = "\033[102m"
	BgYellowBright = "\033[103m"
	BgBlueBright = "\033[104m"
	BgMagentaBright = "\033[105m"
	BgCyanBright = "\033[106m"
	BgWhiteBright = "\033[107m"
)

// Reset all colors
const (
	Reset = "\033[0m"
)

func Colorize(color, text string) string {
	return color + text + Reset
}
