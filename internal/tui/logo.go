package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// logoLines is the mallard isotype (mallard-docs/public/logo-mark.svg) in
// braille, 24 columns wide. Regenerate it with scripts/logo-braille.py
// rather than editing the dots by hand, so it keeps matching the brand mark.
var logoLines = []string{
	"⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣴⣾⣿⣿⣷⣦⡀⠀⠀⠀⠀⠀⠀",
	"⠀⠀⠀⠀⠀⠀⠀⠀⠀⣴⣿⣿⣿⣿⠉⠙⣿⣿⡄⠀⠀⠀⠀⠀",
	"⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⣿⣿⣷⣾⣿⡿⠁⢠⣤⣀⣀⣀",
	"⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⣿⣿⣿⡛⠿⠟⠋⠀⠰⠿⠿⠿⠟⠋",
	"⣶⣄⠀⠀⠀⠀⠀⠀⠀⠙⣿⣿⣿⣿⣦⣄⣀⠀⠀⠀⠀⠀⠀⠀",
	"⢹⣿⣷⣤⣀⠀⠀⠀⠀⠀⠈⠙⢿⣿⣿⣿⣿⣿⣦⡀⠀⠀⠀⠀",
	"⠀⠻⣿⣿⣿⣿⣷⣶⣶⣶⣶⣤⠀⠹⣿⣿⣿⣿⣿⣿⣆⠀⠀⠀",
	"⣤⡀⠈⠛⠿⣿⣿⣿⣿⠿⣿⣿⡇⠀⣿⣿⣿⣯⠀⢻⣿⡄⠀⠀",
	"⠘⢿⣷⣦⣄⣀⣀⣀⣀⣤⣾⣿⠇⢀⣿⣿⣿⣿⠀⢸⣿⡇⠀⠀",
	"⠀⠈⠙⠿⣿⣿⣿⣿⣿⡿⠟⠁⢀⣾⣿⣿⣿⠇⠀⣾⣿⠃⠀⠀",
	"⠀⠀⠀⠀⢀⡀⠉⠉⢀⣀⣤⣶⣿⣿⣿⡿⠋⢀⣼⣿⠏⠀⠀⠀",
	"⠀⠀⠀⠀⠀⠙⠻⠿⣿⣿⣿⡿⠿⠛⠉⢀⣴⣿⠟⠁⠀⠀⠀⠀",
	"⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠠⠴⠞⠛⠉⠀⠀⠀⠀⠀⠀⠀",
}

// duckGradient is the top-to-bottom gradient for the logo, built around the
// isotype's brand blue #2679F3. Dark terminals get a neon ramp that fades
// from a bright head into the brand blue; light terminals get deeper blues.
var duckGradient = []lipgloss.AdaptiveColor{
	{Dark: "#8FD3FF", Light: "#1A5FD6"}, // band 1 — head highlight
	{Dark: "#5CB8FF", Light: "#154FBF"}, // band 2 — head/bill
	{Dark: "#3D9BFF", Light: "#1244A6"}, // band 3 — neck
	{Dark: "#2679F3", Light: "#0F398C"}, // band 4 — body, brand blue
	{Dark: "#1F63D9", Light: "#0B2E73"}, // band 5 — tail/shadow
}

// RenderLogo returns the braille duck logo with a top-to-bottom gradient.
func RenderLogo() string {
	total := len(logoLines)
	if total == 0 {
		return ""
	}

	bands := len(duckGradient)
	var b strings.Builder

	for i, line := range logoLines {
		bandIdx := (i * bands) / total
		if bandIdx >= bands {
			bandIdx = bands - 1
		}
		style := lipgloss.NewStyle().Foreground(duckGradient[bandIdx]).Bold(true)
		b.WriteString(style.Render(line))
		if i < total-1 {
			b.WriteByte('\n')
		}
	}

	return b.String()
}
