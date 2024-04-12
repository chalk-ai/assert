package assert_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	. "github.com/chalk-ai/assert"
)

// region FuzzUtil

func TestFuzzUtilMergeSets(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		set1 := []string{"A", "B", "C"}
		set2 := []string{"D", "E", "F"}
		expected := []string{"A", "B", "C", "D", "E", "F"}
		Equal(t, expected, FuzzUtilMergeSets(set1, set2))
	})

	t.Run("Int", func(t *testing.T) {
		set1 := []int{1, 2, 3}
		set2 := []int{4, 5, 6}
		expected := []int{1, 2, 3, 4, 5, 6}
		Equal(t, expected, FuzzUtilMergeSets(set1, set2))
	})
}

func TestFuzzUtilModifySet(t *testing.T) {
	t.Run("String Slice", func(t *testing.T) {
		slice := []string{"Hello", "World", "TeSt"}
		expected := []string{"hello", "world", "test"}
		input := FuzzUtilModifySet(slice, func(index int, value string) string {
			return strings.ToLower(value)
		})
		Equal(t, expected, input)
	})

	t.Run("Int Slice", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		expected := []int{2, 4, 6, 8, 10}
		input := FuzzUtilModifySet(slice, func(index int, value int) int {
			return value * 2
		})
		Equal(t, expected, input)
	})

	t.Run("Bool Slice", func(t *testing.T) {
		slice := []bool{true, false}
		expected := []bool{false, true}
		input := FuzzUtilModifySet(slice, func(index int, value bool) bool {
			return !value
		})
		Equal(t, expected, input)
	})
}

func TestFuzzUtilLimitSet(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("String (Limit=%d)", i), func(t *testing.T) {
			Len(t, FuzzUtilLimitSet(FuzzStringFull(), i), i)
		})

		t.Run(fmt.Sprintf("Int (Limit=%d)", i), func(t *testing.T) {
			Len(t, FuzzUtilLimitSet(FuzzIntFull(), i), i)
		})

		t.Run(fmt.Sprintf("Float64 (Limit=%d)", i), func(t *testing.T) {
			Len(t, FuzzUtilLimitSet(FuzzFloat64Full(), i), i)
		})
	}

	t.Run(fmt.Sprintf("Max bigger than len (Limit=%d)", 10), func(t *testing.T) {
		Len(t, FuzzUtilLimitSet([]string{"a", "b", "c"}, 10), 3)
	})
}

func TestFuzzUtilDistinctSet(t *testing.T) {
	Equal(t, FuzzUtilDistinctSet([]string{"A", "B", "B", "A", "C"}), []string{"A", "B", "C"})
	Equal(t, FuzzUtilDistinctSet([]int{1, 2, 2, 1, 3}), []int{1, 2, 3})
}

// endregion

// region FuzzString
func TestFuzzStringGenerateRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("Length=%d", i), func(t *testing.T) {
			Len(t, FuzzStringGenerateRandom(1, i)[0], i)
		})
	}

	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprintf("Count=%d", i), func(t *testing.T) {
			Len(t, FuzzStringGenerateRandom(i, 5), i)
		})
	}
}

func TestFuzzStringNumeric(t *testing.T) {
	FuzzUtilRunTests(t, FuzzStringNumeric(), func(t *testing.T, index int, v string) {
		f, err := strconv.ParseFloat(v, 64)
		Numeric(t, f)
		NoError(t, err)
	})
}

func TestFuzzStringFull(t *testing.T) {
	Greater(t, len(FuzzStringFull()), 0)
}

func TestFuzzStringEmpty(t *testing.T) {
	Equal(t, FuzzStringEmpty()[0], "")
	Len(t, FuzzStringEmpty(), 1)
}

func TestFuzzStringEmailAddresses(t *testing.T) {
	t.Run("Must contain @", func(t *testing.T) {
		FuzzUtilRunTests(t, FuzzStringEmailAddresses(), func(t *testing.T, index int, v string) {
			Contains(t, v, "@")
		})
	})
}

func TestFuzzStringHtmlTags(t *testing.T) {
	t.Run("Must contain < and >", func(t *testing.T) {
		FuzzUtilRunTests(t, FuzzStringHtmlTags(), func(t *testing.T, index int, v string) {
			Contains(t, v, "<")
			Contains(t, v, ">")
		})
	})
}

func TestFuzzStringLong(t *testing.T) {
	t.Run("Length fits docs", func(t *testing.T) {
		set := FuzzStringLong()
		Len(t, set[0], 25)
		Len(t, set[1], 50)
		Len(t, set[2], 100)
		Len(t, set[3], 1_000)
		Len(t, set[4], 100_000)
	})
}

func TestFuzzStringUsernames(t *testing.T) {
	Greater(t, len(FuzzStringUsernames()), 0)
}

// endregion

// region FuzzBool

func TestFuzzBoolFull(t *testing.T) {
	Equal(t, []bool{true, false}, FuzzBoolFull())
}

func TestFuzzBoolRunTests(t *testing.T) {
	Greater(t, len(FuzzBoolFull()), 0)
	FuzzUtilRunTests(t, FuzzBoolFull(), func(t *testing.T, index int, f bool) {
		NotNil(t, f)
	})
}

// endregion

// region FuzzInt
func TestFuzzIntGenerateRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Run("GenerateRandomPositive generates positive numbers only", func(t *testing.T) {
			Greater(t, FuzzIntGenerateRandomPositive(1, 100)[0], 0)
		})

		t.Run("GenerateRandomNegative generates negative numbers only", func(t *testing.T) {
			Less(t, FuzzIntGenerateRandomNegative(1, 100)[0], 0)
		})
	}
}

func TestFuzzIntFull(t *testing.T) {
	Greater(t, len(FuzzIntFull()), 0)
}

func TestFuzzIntGenerateRandomRange(t *testing.T) {
	for i := 0; i < 1000; i++ {
		generated := FuzzIntGenerateRandomRange(1, i*10, i*10+10)[0]
		GreaterOrEqual(t, generated, i*10)
		LessOrEqual(t, generated, i*10+10)
	}
}

// endregion

// region FuzzFloat64
func TestFuzzFloat64Full(t *testing.T) {
	Greater(t, len(FuzzFloat64Full()), 0)
}

func TestFuzzFloat64RunTests(t *testing.T) {
	FuzzUtilRunTests(t, FuzzFloat64Full(), func(t *testing.T, index int, f float64) {
		NotNil(t, f)
	})
}

func TestFuzzFloat64GenerateRandomNegative(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := FuzzFloat64GenerateRandomNegative(1, 0)[0]
		t.Run(fmt.Sprintf("%v", n), func(t *testing.T) {
			Less(t, n, 0)
		})
	}
}

func TestFuzzFloat64GenerateRandomPositive(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := FuzzFloat64GenerateRandomPositive(1, 0)[0]
		t.Run(fmt.Sprintf("%v", n), func(t *testing.T) {
			Greater(t, n, 0)
		})
	}
}

func TestFuzzFloat64GenerateRandomRange(t *testing.T) {
	for i := 0; i < 1000; i++ {
		generated := FuzzFloat64GenerateRandomRange(1, float64(i*10), float64(i*10+10))[0]
		GreaterOrEqual(t, generated, float64(i*10))
		LessOrEqual(t, generated, float64(i*10+10))
	}
}
