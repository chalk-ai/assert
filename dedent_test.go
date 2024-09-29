package assert

import (
	"fmt"
	"testing"
)

type dedentTest struct {
	text, expect string
}

func TestDedentNoMargin(t *testing.T) {
	t.Parallel()
	for _, text := range []string{
		// No lines indented
		"Hello there.\nHow are you?\nOh good, I'm glad.",
		// Similar with a blank line
		"Hello there.\n\nBoo!",
		// Some lines indented, but overall margin is still zero
		"Hello there.\n  This is indented.",
		// Again, add a blank line.
		"Hello there.\n\n  Boo!\n",
	} {
		t.Run(text, func(t *testing.T) {
			t.Parallel()
			Equal(t, text, Dedent(text))
		})
	}
}

func TestDedentEven(t *testing.T) {
	t.Parallel()
	texts := []dedentTest{
		{
			// All lines indented by two spaces
			text:   "  Hello there.\n  How are ya?\n  Oh good.",
			expect: "Hello there.\nHow are ya?\nOh good.",
		},
		{
			// Same, with blank lines
			text:   "  Hello there.\n\n  How are ya?\n  Oh good.\n",
			expect: "Hello there.\n\nHow are ya?\nOh good.\n",
		},
		{
			// Now indent one of the blank lines
			text:   "  Hello there.\n  \n  How are ya?\n  Oh good.\n",
			expect: "Hello there.\n\nHow are ya?\nOh good.\n",
		},
	}

	for _, text := range texts {
		Equal(t, text.expect, Dedent(text.text))
	}
}

func TestDedentUneven(t *testing.T) {
	t.Parallel()
	texts := []dedentTest{
		{
			// Lines indented unevenly
			text: `
			def foo():
				while 1:
					return foo
			`,
			expect: `
def foo():
	while 1:
		return foo
`,
		},
		{
			// Uneven indentation with a blank line
			text:   "  Foo\n    Bar\n\n   Baz\n",
			expect: "Foo\n  Bar\n\n Baz\n",
		},
		{
			// Uneven indentation with a whitespace-only line
			text:   "  Foo\n    Bar\n \n   Baz\n",
			expect: "Foo\n  Bar\n\n Baz\n",
		},
	}

	for _, text := range texts {
		Equal(t, text.expect, Dedent(text.text))
	}
}

// Dedent() should not mangle internal tabs.
func TestDedentPreserveInternalTabs(t *testing.T) {
	t.Parallel()
	text := "  hello\tthere\n  how are\tyou?"
	expect := "hello\tthere\nhow are\tyou?"
	Equal(t, expect, Dedent(text))

	// Make sure that it preserves tabs when it's not making any changes at all
	Equal(t, expect, Dedent(expect))
}

// Dedent() should not mangle tabs in the margin (i.e. tabs and spaces both
// count as margin, but are *not* considered equivalent).
func TestDedentPreserveMarginTabs(t *testing.T) {
	t.Parallel()
	for _, text := range []string{
		"  hello there\n\thow are you?",
		// Same effect even if we have 8 spaces
		"        hello there\n\thow are you?",
	} {
		Equal(t, text, Dedent(text))
	}

	texts2 := []dedentTest{
		{
			// Dedent() only removes whitespace that can be uniformly removed!
			text:   "\thello there\n\thow are you?",
			expect: "hello there\nhow are you?",
		},
		{
			text:   "  \thello there\n  \thow are you?",
			expect: "hello there\nhow are you?",
		},
		{
			text:   "  \t  hello there\n  \t  how are you?",
			expect: "hello there\nhow are you?",
		},
		{
			text:   "  \thello there\n  \t  how are you?",
			expect: "hello there\n  how are you?",
		},
	}

	for _, text := range texts2 {
		Equal(t, text.expect, Dedent(text.text))
	}
}

func ExampleDedent() {
	s := `
		Lorem ipsum dolor sit amet,
		consectetur adipiscing elit.
		Curabitur justo tellus, facilisis nec efficitur dictum,
		fermentum vitae ligula. Sed eu convallis sapien.`
	fmt.Println(Dedent(s))
	fmt.Println("-------------")
	fmt.Println(s)
	// Output:
	// Lorem ipsum dolor sit amet,
	// consectetur adipiscing elit.
	// Curabitur justo tellus, facilisis nec efficitur dictum,
	// fermentum vitae ligula. Sed eu convallis sapien.
	// -------------
	//
	//		Lorem ipsum dolor sit amet,
	//		consectetur adipiscing elit.
	//		Curabitur justo tellus, facilisis nec efficitur dictum,
	//		fermentum vitae ligula. Sed eu convallis sapien.
}
