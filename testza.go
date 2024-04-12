package assert

import (
	"github.com/pterm/pterm"
	"sync"
)

var infoPrinter = pterm.
	DefaultSection.
	WithStyle(pterm.NewStyle(pterm.FgMagenta)).
	WithLevel(2).
	WithBottomPadding(0).
	WithTopPadding(0)

var secondary = pterm.LightCyan

var initSync sync.Mutex
