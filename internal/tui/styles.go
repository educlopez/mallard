package tui

import "github.com/charmbracelet/lipgloss"

// Neon-blue palette built around the isotype's brand blue #2679F3.
//
// Every color is adaptive: lipgloss picks Dark on a dark terminal background
// and Light on a light one. Dark values are bright, saturated "neon" tones;
// light values are deeper shades of the same hue. All text colors keep a
// WCAG contrast of at least 4.5:1 on black, #191724 and #1E1E2E (dark) and on
// white and #F5F5F7 (light).
var (
	ColorOverlay = lipgloss.AdaptiveColor{Dark: "#4A5272", Light: "#B4BCCF"} // panel borders
	ColorText    = lipgloss.AdaptiveColor{Dark: "#E3E8F4", Light: "#1C2233"}
	ColorSubtext = lipgloss.AdaptiveColor{Dark: "#8C93AD", Light: "#5A6378"}
	ColorRed     = lipgloss.AdaptiveColor{Dark: "#FF5C8A", Light: "#C8245A"}
	ColorGreen   = lipgloss.AdaptiveColor{Dark: "#4DFFC3", Light: "#00805A"}
	ColorYellow  = lipgloss.AdaptiveColor{Dark: "#FFD166", Light: "#9A6300"}

	// Brand blues (replace the old rubber-duck yellow ramp).
	ColorNeon   = lipgloss.AdaptiveColor{Dark: "#4DB8FF", Light: "#0A5BD3"} // title, cursor, frame
	ColorAccent = lipgloss.AdaptiveColor{Dark: "#3D8BFF", Light: "#1A56DB"} // headings, selection
	ColorGlow   = lipgloss.AdaptiveColor{Dark: "#38E8FF", Light: "#00708A"} // checkmarks
)

// Cursor is the prefix used for the currently focused item.
const Cursor = "▸ "

// Tagline returns the welcome screen tagline with the given version.
func Tagline(version string) string {
	if version == "" {
		return "mallard — multi-agent skill manager"
	}
	return "mallard " + version + " — multi-agent skill manager"
}

// Pre-built reusable styles (mirror gentle-ai's exported style API).
var (
	TitleStyle = lipgloss.NewStyle().
			Foreground(ColorNeon).
			Bold(true)

	HeadingStyle = lipgloss.NewStyle().
			Foreground(ColorAccent).
			Bold(true)

	HelpStyle = lipgloss.NewStyle().
			Foreground(ColorSubtext)

	SubtextStyle = lipgloss.NewStyle().
			Foreground(ColorSubtext)

	SelectedStyle = lipgloss.NewStyle().
			Foreground(ColorNeon).
			Bold(true)

	UnselectedStyle = lipgloss.NewStyle().
			Foreground(ColorText)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(ColorGreen)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(ColorRed)

	WarningStyle = lipgloss.NewStyle().
			Foreground(ColorYellow)

	FrameStyle = lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder()).
			BorderForeground(ColorNeon).
			Padding(1, 2)

	PanelStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(ColorOverlay).
			Padding(0, 1)
)

// Legacy lowercase aliases — kept so the non-welcome screens (agents,
// skills, install, update flow) continue to compile without churn. They
// reuse the same colors as the exported styles above.
var (
	colorAccent  = ColorAccent
	colorSuccess = ColorGlow
	colorError   = ColorRed
	colorWarning = ColorYellow
	colorMuted   = ColorSubtext
	colorTitle   = ColorNeon

	styleTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorTitle).
			MarginBottom(1)

	styleSuccess = lipgloss.NewStyle().
			Foreground(colorSuccess)

	styleError = lipgloss.NewStyle().
			Foreground(colorError)

	styleWarning = lipgloss.NewStyle().
			Foreground(colorWarning).
			Bold(true)

	styleMuted = lipgloss.NewStyle().
			Foreground(colorMuted)

	styleAccent = lipgloss.NewStyle().
			Foreground(colorAccent)

	styleSelected = lipgloss.NewStyle().
			Foreground(colorAccent).
			Bold(true)

	styleCursor = lipgloss.NewStyle().
			Foreground(ColorNeon).
			Bold(true)

	styleBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorAccent).
			Padding(1, 2).
			MarginTop(1)

	styleAgentBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorMuted).
			Padding(0, 2).
			MarginTop(1)

	styleKey = lipgloss.NewStyle().
			Foreground(colorMuted)
)

var spinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
