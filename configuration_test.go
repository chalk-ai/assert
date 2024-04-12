package assert

import (
	"testing"
	"time"

	"github.com/chalk-ai/assert/internal"
	"github.com/pterm/pterm"
)

func TestSetColorsEnabled(t *testing.T) {
	t.Run("Disable", func(t *testing.T) {
		SetColorsEnabled(false)
		False(t, pterm.PrintColor)
		False(t, GetColorsEnabled())
	})

	t.Run("Enable", func(t *testing.T) {
		SetColorsEnabled(true)
		True(t, pterm.PrintColor)
		True(t, GetColorsEnabled())
	})
}

func TestSetLineNumbersEnabled(t *testing.T) {
	t.Run("Disable", func(t *testing.T) {
		SetLineNumbersEnabled(false)
		False(t, internal.LineNumbersEnabled)
		False(t, GetLineNumbersEnabled())
	})

	t.Run("Enable", func(t *testing.T) {
		SetLineNumbersEnabled(true)
		True(t, internal.LineNumbersEnabled)
		True(t, GetLineNumbersEnabled())
	})
}

func TestSetRandomSeed(t *testing.T) {
	SetRandomSeed(1337)
	Equal(t, int64(1337), randomSeed)
	Equal(t, int64(1337), GetRandomSeed())
	Equal(t, "4U390O49B9", FuzzStringGenerateRandom(1, 10)[0])
	un := time.Now().UnixNano()
	SetRandomSeed(un)
	Equal(t, un, randomSeed)
	Equal(t, un, GetRandomSeed())
}

func TestSetShowStartupMessage(t *testing.T) {
	t.Run("Default is false", func(t *testing.T) {
		False(t, showStartupMessage)
		False(t, GetShowStartupMessage())
	})

	t.Run("Set to false", func(t *testing.T) {
		SetShowStartupMessage(false)
		False(t, showStartupMessage)
		False(t, GetShowStartupMessage())
	})

	t.Run("Set to true", func(t *testing.T) {
		SetShowStartupMessage(true)
		True(t, showStartupMessage)
		True(t, GetShowStartupMessage())
	})
}

func TestSetEqualContextLineCount(t *testing.T) {
	t.Run("Default is 2", func(t *testing.T) {
		Equal(t, 2, internal.DiffContextLines)
		Equal(t, 2, GetDiffContextLines())
	})

	t.Run("Set to -1", func(t *testing.T) {
		SetDiffContextLines(-1)
		Equal(t, -1, internal.DiffContextLines)
		Equal(t, -1, GetDiffContextLines())
	})

	t.Run("Set to 3", func(t *testing.T) {
		SetDiffContextLines(3)
		Equal(t, 3, internal.DiffContextLines)
		Equal(t, 3, GetDiffContextLines())
	})

	SetDiffContextLines(2)
}
