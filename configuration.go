package assert

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/chalk-ai/assert/internal"
	"github.com/klauspost/cpuid/v2"
	"github.com/pterm/pterm"
)

var randomSeed int64
var showStartupMessage = false

func init() {
	initSync.Lock()
	defer initSync.Unlock()

	// Defining flags to show up in the help message
	flag.Bool("assert.disable-color", false, "disables colored output")
	flag.Bool("assert.disable-line-numbers", false, "disables line numbers in output")
	flag.Bool("assert.disable-startup-message", false, "disable the startup message")
	flag.Int64("assert.seed", 0, "seed used for random operations")
	flag.Int("assert.diff-context-lines", 2, "sets the context line count in difference output")

	for i, arg := range os.Args {
		// Check if the argument is a flag
		if !strings.HasPrefix(arg, "--") {
			continue
		}

		// Figure out if the flag has a value
		var value string
		if i != len(os.Args)-1 {
			value = os.Args[i+1]
			if strings.HasPrefix(value, "-") {
				value = ""
			}
		}

		// Check for set flags and run the appropriate function
		switch strings.TrimPrefix(arg, "--assert.") {
		case "disable-color":
			SetColorsEnabled(false)
		case "disable-line-numbers":
			SetLineNumbersEnabled(false)
		case "disable-startup-message":
			SetShowStartupMessage(false)
		case "seed":
			seed, err := strconv.Atoi(value)
			pterm.Fatal.PrintOnError(err)
			SetRandomSeed(int64(seed))
		case "diff-context-lines":
			v, err := strconv.Atoi(value)
			pterm.Fatal.PrintOnError(err)
			SetDiffContextLines(v)
		}
	}

	go func() {
		initSync.Lock()
		defer initSync.Unlock()

		if randomSeed == 0 {
			randomSeed = time.Now().UnixNano()
			rand.Seed(randomSeed)
		}

		if showStartupMessage {
			var startupMessage string
			startupMessage += infoPrinter.WithLevel(1).Sprintln("Running tests with " + secondary("assert"))
			startupMessage += infoPrinter.Sprintfln(`Using seed "%s" for random operations`, secondary(randomSeed))
			startupMessage += infoPrinter.Sprintfln(`System info: OS=%s | arch=%s | cpu=%s | go=%s`, secondary(runtime.GOOS), secondary(runtime.GOARCH), secondary(cpuid.CPU.BrandName), secondary(runtime.Version()))
			startupMessage += fmt.Sprintln()
			pterm.Println(startupMessage)
		}
	}()
}

// SetColorsEnabled controls if assert should print colored output.
// You should use this in the init() method of the package, which contains your tests.
//
// > This setting can also be set by the command line flag --assert.disable-color.
//
// Example:
//
//	init() {
//	  assert.SetColorsEnabled(false) // Disable colored output
//	  assert.SetColorsEnabled(true)  // Enable colored output
//	}
func SetColorsEnabled(enabled bool) {
	initSync.Lock()
	defer initSync.Unlock()

	if enabled {
		pterm.EnableColor()
	} else {
		pterm.DisableColor()
	}
}

// GetColorsEnabled returns current value of ColorsEnabled setting.
// ColorsEnabled controls if assert should print colored output.
func GetColorsEnabled() bool {
	return pterm.PrintColor
}

// SetLineNumbersEnabled controls if line numbers should be printed in failing tests.
// You should use this in the init() method of the package, which contains your tests.
//
// > This setting can also be set by the command line flag --assert.disable-line-numbers.
//
// Example:
//
//	init() {
//	  assert.SetLineNumbersEnabled(false) // Disable line numbers
//	  assert.SetLineNumbersEnabled(true)  // Enable line numbers
//	}
func SetLineNumbersEnabled(enabled bool) {
	initSync.Lock()
	defer initSync.Unlock()

	internal.LineNumbersEnabled = enabled
}

// GetLineNumbersEnabled returns current value of LineNumbersEnabled setting.
// LineNumbersEnabled controls if line numbers should be printed in failing tests.
func GetLineNumbersEnabled() bool {
	return internal.LineNumbersEnabled
}

// SetRandomSeed sets the seed for the random generator used in assert.
// Using the same seed will result in the same random sequences each time and guarantee a reproducible test run.
// Use this setting, if you want a 100% deterministic test.
// You should use this in the init() method of the package, which contains your tests.
//
// > This setting can also be set by the command line flag --assert.seed.
//
// Example:
//
//	init() {
//	  assert.SetRandomSeed(1337) // Set the seed to 1337
//	  assert.SetRandomSeed(time.Now().UnixNano()) // Set the seed back to the current time (default | non-deterministic)
//	}
func SetRandomSeed(seed int64) {
	initSync.Lock()
	defer initSync.Unlock()

	randomSeed = seed
	rand.Seed(seed)
}

// GetRandomSeed returns current value of the random seed setting.
func GetRandomSeed() int64 {
	return randomSeed
}

// SetShowStartupMessage controls if the startup message should be printed.
// You should use this in the init() method of the package, which contains your tests.
//
// > This setting can also be set by the command line flag --assert.disable-startup-message.
//
// Example:
//
//	init() {
//	  assert.SetShowStartupMessage(false) // Disable the startup message
//	  assert.SetShowStartupMessage(true)  // Enable the startup message
//	}
func SetShowStartupMessage(show bool) {
	initSync.Lock()
	defer initSync.Unlock()

	showStartupMessage = show
}

// GetShowStartupMessage returns current value of showStartupMessage setting.
// showStartupMessage setting controls if the startup message should be printed.
func GetShowStartupMessage() bool {
	return showStartupMessage
}

// SetDiffContextLines controls how many lines are shown around a changed diff line.
// If set to -1 it will show full diff.
// You should use this in the init() method of the package, which contains your tests.
//
// > This setting can also be set by the command line flag --assert.diff-context-lines.
//
// Example:
//
//	init() {
//	  assert.SetDiffContextLines(-1) // Show all diff lines
//	  assert.SetDiffContextLines(3)  // Show 3 lines around every changed line
//	}
func SetDiffContextLines(lines int) {
	initSync.Lock()
	defer initSync.Unlock()

	internal.DiffContextLines = lines
}

// GetDiffContextLines returns current value of DiffContextLines setting.
// DiffContextLines setting controls how many lines are shown around a changed diff line.
// If set to -1 it will show full diff.
func GetDiffContextLines() int {
	return internal.DiffContextLines
}
