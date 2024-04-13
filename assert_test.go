package assert_test

import (
	"errors"
	"fmt"
	"go/types"
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/pterm/pterm"

	. "github.com/chalk-ai/assert"
)

type testMock struct {
	ErrorCalled  bool
	ErrorMessage string
}

func (m *testMock) fail(msg ...any) {
	m.ErrorCalled = true
	m.ErrorMessage = fmt.Sprint(msg...)
}

func (m *testMock) Error(args ...any) {
	m.fail(args...)
}

// Errorf is a mock of testing.T.
func (m *testMock) Errorf(format string, args ...any) {
	m.fail(fmt.Sprintf(format, args...))
}

// Fail is a mock of testing.T.
func (m *testMock) Fail() {
	m.fail()
}

// FailNow is a mock of testing.T.
func (m *testMock) FailNow() {
	m.fail()
}

// Fatal is a mock of testing.T.
func (m *testMock) Fatal(args ...any) {
	m.fail(args...)
}

// Fatalf is a mock of testing.T.
func (m *testMock) Fatalf(format string, args ...any) {
	m.fail(fmt.Sprintf(format, args...))
}

type assertionTestStruct struct {
	Name string
	Age  int
	Meta assertionTestStructNested
}

type assertionTestStructNested struct {
	ID    int
	Admin bool
}

func randomString() string {
	return FuzzStringGenerateRandom(1, rand.Intn(10))[0]
}

func generateStruct() (ret assertionTestStruct) {
	ret.Name = randomString()
	ret.Age = rand.Intn(40)
	ret.Meta.ID = rand.Intn(10_000)
	ret.Meta.Admin = rand.Intn(1) == 1

	return
}

func testEqual(t *testing.T, expected, actual any) {
	t.Helper()

	t.Run("Equal", func(t *testing.T) {
		t.Helper()

		Equal(t, expected, actual)
	})

	t.Run("EqualValues", func(t *testing.T) {
		t.Helper()

		EqualValues(t, expected, actual)
	})
}

func TestAssertKindOf(t *testing.T) {
	tests := []struct {
		kind  reflect.Kind
		value any
	}{
		{kind: reflect.String, value: "Hello, World!"},
		{kind: reflect.Int, value: 1337},
		{kind: reflect.Bool, value: true},
		{kind: reflect.Bool, value: false},
		{kind: reflect.Map, value: make(map[string]int)},
		{kind: reflect.Func, value: func() {}},
		{kind: reflect.Struct, value: testing.T{}},
		{kind: reflect.Slice, value: []string{}},
		{kind: reflect.Slice, value: []int{}},
		{kind: reflect.Slice, value: []float64{}},
		{kind: reflect.Slice, value: []float32{}},
		{kind: reflect.Float64, value: 13.37},
		{kind: reflect.Float32, value: float32(13.37)},
	}
	for _, test := range tests {
		t.Run(test.kind.String(), func(t *testing.T) {
			KindOf(t, test.kind, test.value)
		})
	}
}

func TestAssertKindOf_fails(t *testing.T) {
	tests := []struct {
		kind  reflect.Kind
		value any
	}{
		{kind: reflect.Int, value: "Hello, World!"},
		{kind: reflect.Bool, value: 1337},
		{kind: reflect.Float64, value: true},
		{kind: reflect.Float32, value: false},
		{kind: reflect.Struct, value: make(map[string]int)},
		{kind: reflect.Map, value: func() {}},
		{kind: reflect.Slice, value: testing.T{}},
		{kind: reflect.Struct, value: []string{}},
		{kind: reflect.Array, value: []int{}},
		{kind: reflect.Chan, value: []float64{}},
		{kind: reflect.Complex64, value: []float32{}},
		{kind: reflect.Complex128, value: 13.37},
		{kind: reflect.Float64, value: float32(13.37)},
	}
	for _, test := range tests {
		t.Run(test.kind.String(), func(t *testing.T) {
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				KindOf(t, test.kind, test.value)
			})
		})
	}
}

func TestAssertNotKindOf(t *testing.T) {
	tests := []struct {
		kind  reflect.Kind
		value any
	}{
		{kind: reflect.Int, value: "Hello, World!"},
		{kind: reflect.Bool, value: 1337},
		{kind: reflect.Float64, value: true},
		{kind: reflect.Float32, value: false},
		{kind: reflect.Struct, value: make(map[string]int)},
		{kind: reflect.Map, value: func() {}},
		{kind: reflect.Slice, value: testing.T{}},
		{kind: reflect.Struct, value: []string{}},
		{kind: reflect.Array, value: []int{}},
		{kind: reflect.Chan, value: []float64{}},
		{kind: reflect.Complex64, value: []float32{}},
		{kind: reflect.Complex128, value: 13.37},
		{kind: reflect.Float64, value: float32(13.37)},
	}
	for _, test := range tests {
		t.Run(test.kind.String(), func(t *testing.T) {
			NotKindOf(t, test.kind, test.value)
		})
	}
}

func TestAssertNotKindOf_fails(t *testing.T) {
	tests := []struct {
		kind  reflect.Kind
		value any
	}{
		{kind: reflect.String, value: "Hello, World!"},
		{kind: reflect.Int, value: 1337},
		{kind: reflect.Bool, value: true},
		{kind: reflect.Bool, value: false},
		{kind: reflect.Map, value: make(map[string]int)},
		{kind: reflect.Func, value: func() {}},
		{kind: reflect.Struct, value: testing.T{}},
		{kind: reflect.Slice, value: []string{}},
		{kind: reflect.Slice, value: []int{}},
		{kind: reflect.Slice, value: []float64{}},
		{kind: reflect.Slice, value: []float32{}},
		{kind: reflect.Float64, value: 13.37},
		{kind: reflect.Float32, value: float32(13.37)},
	}
	for _, test := range tests {
		t.Run(test.kind.String(), func(t *testing.T) {
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				NotKindOf(t, test.kind, test.value)
			})
		})
	}
}

func TestAssertNumeric(t *testing.T) {
	var numbers []any

	for i := 0; i < 10; i++ {
		numbers = append(numbers,
			i,
			int8(i),
			int16(i),
			int32(i),
			int64(i),
			uint(i),
			uint8(i),
			uint16(i),
			uint32(i),
			uint64(i),
			float32(i)+0.25,
			float32(i)+0.5,
			float64(i)+0.25,
			float64(i)+0.5,
			complex64(complex(float64(i), 2)),
			complex(float64(i), 2),
		)
	}

	for _, number := range numbers {
		t.Run(pterm.Sprintf("Type=%s;Value=%#v", reflect.TypeOf(number).Kind().String(), number), func(t *testing.T) {
			Numeric(t, number)
		})
	}
}

func TestAssertNumeric_fails(t *testing.T) {
	noNumbers := []any{"Hello, World!", true, false}
	for _, number := range noNumbers {
		t.Run(pterm.Sprintf("Type=%s;Value=%#v", reflect.TypeOf(number).Kind().String(), number), func(t *testing.T) {
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				Numeric(t, number)
			})
		})
	}
}

func TestAssertNotNumeric_fails(t *testing.T) {
	var numbers []any

	for i := 0; i < 10; i++ {
		numbers = append(numbers,
			i,
			int8(i),
			int16(i),
			int32(i),
			int64(i),
			uint(i),
			uint8(i),
			uint16(i),
			uint32(i),
			uint64(i),
			float32(i)+0.25,
			float32(i)+0.5,
			float64(i)+0.25,
			float64(i)+0.5,
			complex64(complex(float64(i), 2)),
			complex(float64(i), 2),
		)
	}

	for _, number := range numbers {
		t.Run(pterm.Sprintf("Type=%s;Value=%#v", reflect.TypeOf(number).Kind().String(), number), func(t *testing.T) {
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				NotNumeric(t, number)
			})
		})
	}
}

func TestAssertZero(t *testing.T) {
	var zeroValues []any
	zeroValues = append(zeroValues, 0, "", false, nil)

	for i, value := range zeroValues {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			Zero(t, value)
		})
	}
}

func TestAssertZero_fails(t *testing.T) {
	var nonZeroValues []any
	nonZeroValues = append(nonZeroValues, true, "asd", 123, 1.5, 'a')

	for i, value := range nonZeroValues {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				Zero(t, value)
			})
		})
	}
}

func TestAssertNotZero(t *testing.T) {
	var nonZeroValues []any
	nonZeroValues = append(nonZeroValues, true, "asd", 123, 1.5, 'a')

	for i, value := range nonZeroValues {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			NotZero(t, value)
		})
	}
}

func TestAssertNotZero_fails(t *testing.T) {
	var zeroValues []any
	zeroValues = append(zeroValues, 0, "", false, nil)

	for i, value := range zeroValues {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				NotZero(t, value)
			})
		})
	}
}

func TestAssertEqual(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		FuzzUtilRunTests(t, FuzzStringFull(), func(t *testing.T, index int, str string) {
			t.Helper()

			testEqual(t, str, str)
		})
	})

	t.Run("Nil and Nil are equal", func(t *testing.T) {
		Equal(t, nil, nil)
	})

	t.Run("Equal pointers to same struct", func(t *testing.T) {
		s := generateStruct()
		Equal(t, &s, &s)
	})

	t.Run("Equal dereference to same pointer struct", func(t *testing.T) {
		s := generateStruct()
		s2 := &s
		Equal(t, *s2, *s2)
	})

	t.Run("Structs", func(t *testing.T) {
		// Test ten random structs for equality.
		for i := 0; i < 10; i++ {
			s := generateStruct()
			testEqual(t, s, s)
		}
	})
}

func TestAssertEqual_fails(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		FuzzUtilRunTests(t, FuzzStringFull(), func(t *testing.T, index int, str string) {
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				Equal(t, str, str+" addon")
			})
		})
	})

	t.Run("Structs", func(t *testing.T) {
		// Test ten random structs for equality.
		for i := 0; i < 10; i++ {
			s := generateStruct()
			s2 := s
			s2.Name += " addon"
			t.Run("NotEqual", func(t *testing.T) {
				TestFails(t, func(t TestingPackageWithFailFunctions) {
					Equal(t, s, s2)
				})
			})
			t.Run("NotEqualValues", func(t *testing.T) {
				TestFails(t, func(t TestingPackageWithFailFunctions) {
					EqualValues(t, s, s2)
				})
			})
		}
	})
}

func TestAssertNotEqual(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		FuzzUtilRunTests(t, FuzzStringFull(), func(t *testing.T, index int, str string) {
			NotEqual(t, str, str+" addon")
		})
	})

	t.Run("Structs", func(t *testing.T) {
		// Test ten random structs for equality.
		for i := 0; i < 10; i++ {
			s := generateStruct()
			s2 := s
			s2.Name += " addon"
			t.Run("NotEqual", func(t *testing.T) {
				NotEqual(t, s, s2)
			})
			t.Run("NotEqualValues", func(t *testing.T) {
				NotEqualValues(t, s, s2)
			})
		}
	})
}

func TestAssertNotEqual_fails(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		FuzzUtilRunTests(t, FuzzStringFull(), func(t *testing.T, index int, str string) {
			t.Helper()

			TestFails(t, func(t TestingPackageWithFailFunctions) {
				NotEqual(t, str, str)
			})
		})
	})

	t.Run("Nil and Nil are equal", func(t *testing.T) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			NotEqual(t, nil, nil)
		})
	})

	t.Run("Equal pointers to same struct", func(t *testing.T) {
		s := generateStruct()
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			NotEqual(t, &s, &s)
		})
	})

	t.Run("Equal dereference to same pointer struct", func(t *testing.T) {
		s := generateStruct()
		s2 := &s
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			NotEqual(t, *s2, *s2)
		})
	})

	t.Run("Structs", func(t *testing.T) {
		// Test ten random structs for equality.
		for i := 0; i < 10; i++ {
			s := generateStruct()
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				NotEqual(t, s, s)
			})
		}
	})
}

func TestAssertEqualValues(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		FuzzUtilRunTests(t, FuzzStringFull(), func(t *testing.T, index int, str string) {
			EqualValues(t, str, str)
		})
	})
}

func TestAssertEqualValues_fails(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		FuzzUtilRunTests(t, FuzzStringFull(), func(t *testing.T, index int, str string) {
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				EqualValues(t, str, str+" addon")
			})
		})
	})
}

func TestAssertNotEqualValues(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		FuzzUtilRunTests(t, FuzzStringFull(), func(t *testing.T, index int, str string) {
			NotEqualValues(t, str, str+" addon")
		})
	})
}

func TestAssertNotEqualValues_fails(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		FuzzUtilRunTests(t, FuzzStringFull(), func(t *testing.T, index int, str string) {
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				NotEqualValues(t, str, str)
			})
		})
	})
}

func TestAssertTrue(t *testing.T) {
	True(t, true)
}

func TestAssertTrue_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		True(t, false)
	})
}

func TestAssertFalse(t *testing.T) {
	False(t, false)
}

func TestAssertFalse_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		False(t, true)
	})
}

func TestAssertImplements(t *testing.T) {
	t.Run("ConstImplementsStringer", func(t *testing.T) {
		Implements(t, (*fmt.Stringer)(nil), new(types.Const))
	})
}

func TestAssertImplements_fails(t *testing.T) {
	t.Run("assertionTestStructNotImplementsFmtStringer", func(t *testing.T) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			Implements(t, (*fmt.Stringer)(nil), new(assertionTestStruct))
		})
	})
}

func TestAssertNotImplements(t *testing.T) {
	t.Run("assertionTestStructNotImplementsFmtStringer", func(t *testing.T) {
		NotImplements(t, (*fmt.Stringer)(nil), new(assertionTestStruct))
	})
}

func TestAssertNotImplements_fails(t *testing.T) {
	t.Run("ConstImplementsStringer", func(t *testing.T) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			NotImplements(t, (*fmt.Stringer)(nil), new(types.Const))
		})
	})
}

func TestAssertContains(t *testing.T) {
	tests := []struct {
		name     string
		obj      interface{}
		contains interface{}
	}{
		{name: "String Slice", obj: []string{"Hello", "World", "!"}, contains: "World"},
		{name: "Int Slice", obj: []int{1, 2, 3, 4, 5, 6, 7, 8}, contains: 4},
		{name: "String", obj: "Hello, World!", contains: "World"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Contains(t, test.obj, test.contains)
		})
	}
}

func TestAssertContains_fails(t *testing.T) {
	tests := []struct {
		name     string
		obj      interface{}
		contains interface{}
	}{
		{name: "String Slice", obj: []string{"Hello", "World", "!"}, contains: "asdasdasd"},
		{name: "Int Slice", obj: []int{1, 2, 3, 4, 5, 6, 7, 8}, contains: 1337},
		{name: "String", obj: "Hello, World!", contains: "asdasdasd"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				Contains(t, test.obj, test.contains)
			})
		})
	}
}

func TestAssertNotContains(t *testing.T) {
	s := []string{"Hello", "World"}

	NotContains(t, s, "asdasd")
}

func TestAssertNotContains_fails(t *testing.T) {
	s := []string{"Hello", "World"}
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		NotContains(t, s, "World")
	})
}

func TestAssertPanics(t *testing.T) {
	Panics(t, func() {
		panic("TestPanic")
	})
}

func TestAssertPanics_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		Panics(t, func() {
			// If we do nothing here it can't panic ;)
		})
	})
}

func TestAssertNotPanics(t *testing.T) {
	NotPanics(t, func() {
		// If we do nothing here it can't panic ;)
	})
}

func TestAssertNotPanics_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		NotPanics(t, func() {
			panic("TestPanic")
		})
	})
}

func TestAssertNil(t *testing.T) {
	Nil(t, nil)
}

func TestAssertNil_pointer(t *testing.T) {
	var a *int = nil
	Nil(t, a)
}

func TestAssertNil_interface(t *testing.T) {
	var i *int = nil
	var a any = i
	Nil(t, a)
}

func TestAssertNil_pointer_struct(t *testing.T) {
	var a *assertionTestStruct
	Nil(t, a)
}

func TestAssertNil_fails(t *testing.T) {
	objectsNotNil := []any{"asd", 0, false, true, 'c', "", []int{1, 2, 3}}

	for _, v := range objectsNotNil {
		t.Run(pterm.Sprintf("Value=%#v", v), func(t *testing.T) {
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				Nil(t, v)
			})
		})
	}
}

func TestAssertNotNil(t *testing.T) {
	objectsNotNil := []any{"asd", 0, false, true, 'c', "", []int{1, 2, 3}}

	for _, v := range objectsNotNil {
		t.Run(pterm.Sprintf("Value=%#v", v), func(t *testing.T) {
			NotNil(t, v)
		})
	}
}

func TestAssertNotNil_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		NotNil(t, nil)
	})
}

func TestAssertCompletesIn(t *testing.T) {
	CompletesIn(t, 50*time.Millisecond, func() {
		time.Sleep(5 * time.Millisecond)
	})
}

func TestAssertCompletesIn_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		CompletesIn(t, 10*time.Microsecond, func() {
			time.Sleep(15 * time.Millisecond)
		})
	})
}

func TestAssertNotCompletesIn(t *testing.T) {
	NotCompletesIn(t, 10*time.Microsecond, func() {
		time.Sleep(20 * time.Millisecond)
	})
}

func TestAssertNotCompletesIn_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		NotCompletesIn(t, 50*time.Millisecond, func() {
			time.Sleep(5 * time.Millisecond)
		})
	})
}

func TestAssertNoError(t *testing.T) {
	NoError(t, nil)
}

func TestAssertNoError_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		NoError(t, errors.New("hello error"))
	})
}

func TestAssertGreater(t *testing.T) {
	Greater(t, 2, 1)
	Greater(t, 5, 4)
}

func TestAssertGreater_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		Greater(t, 1, 2)
	})

	TestFails(t, func(t TestingPackageWithFailFunctions) {
		Greater(t, 4, 5)
	})

	TestFails(t, func(t TestingPackageWithFailFunctions) {
		Greater(t, 5, 5)
	})
}

func TestAssertGreaterOrEqual(t *testing.T) {
	GreaterOrEqual(t, 2, 1)
	GreaterOrEqual(t, 5, 4)
	GreaterOrEqual(t, 5, 5)
}

func TestAssertGreaterOrEqual_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		GreaterOrEqual(t, 1, 2)
	})

	TestFails(t, func(t TestingPackageWithFailFunctions) {
		GreaterOrEqual(t, 4, 5)
	})
}

func TestAssertLess(t *testing.T) {
	Less(t, 1, 2)
	Less(t, 4, 5)
}

func TestAssertLess_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		Less(t, 2, 1)
	})

	TestFails(t, func(t TestingPackageWithFailFunctions) {
		Less(t, 5, 4)
	})

	TestFails(t, func(t TestingPackageWithFailFunctions) {
		Less(t, 5, 5)
	})
}

func TestAssertLessOrEqual(t *testing.T) {
	LessOrEqual(t, 1, 2)
	LessOrEqual(t, 4, 5)
	LessOrEqual(t, 5, 5)
}

func TestAssertLessOrEqual_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		LessOrEqual(t, 2, 1)
	})

	TestFails(t, func(t TestingPackageWithFailFunctions) {
		LessOrEqual(t, 5, 4)
	})
}

func TestAssertTestFails(t *testing.T) {
	t.Run("Wrong assertion should fail", func(t *testing.T) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			True(t, false)
		})
	})

	t.Run(".Error should make the test pass", func(t *testing.T) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			t.Error()
		})
	})

	t.Run(".Errorf should make the test pass", func(t *testing.T) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			t.Errorf("")
		})
	})

	t.Run(".Fail should make the test pass", func(t *testing.T) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			t.Fail()
		})
	})

	t.Run(".FailNow should make the test pass", func(t *testing.T) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			t.FailNow()
		})
	})

	t.Run(".Fatal should make the test pass", func(t *testing.T) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			t.Fatal()
		})
	})

	t.Run(".Fatalf should make the test pass", func(t *testing.T) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			t.Fatalf("")
		})
	})
}

func TestAssertTestFails_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			True(t, true)
		})
	})
}

func TestAssertErrorIs(t *testing.T) {
	var testErr = errors.New("hello world")
	var testErrWrapped = fmt.Errorf("test err: %w", testErr)
	ErrorIs(t, testErrWrapped, testErr)
}

func TestAssertErrorIs_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		var testErr = errors.New("hello world")
		var test2Err = errors.New("hello world 2")
		var testErrWrapped = fmt.Errorf("test err: %w", testErr)
		ErrorIs(t, testErrWrapped, test2Err)
	})
}

func TestAssertNotErrorIs(t *testing.T) {
	var testErr = errors.New("hello world")
	var test2Err = errors.New("hello world 2")
	var testErrWrapped = fmt.Errorf("test err: %w", testErr)
	NotErrorIs(t, testErrWrapped, test2Err)
}

func TestAssertNotErrorIs_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		var testErr = errors.New("hello world")
		var testErrWrapped = fmt.Errorf("test err: %w", testErr)
		NotErrorIs(t, testErrWrapped, testErr)
	})
}

func TestAssertLen(t *testing.T) {
	tests := []struct {
		object any
		expLen int
	}{
		{object: "", expLen: 0},
		{object: "a", expLen: 1},
		{object: "123", expLen: 3},
		{object: []string{"", "", "", ""}, expLen: 4},
		{object: []string{}, expLen: 0},
		{object: []int{1, 2, 3, 4, 5}, expLen: 5},
	}
	for _, test := range tests {
		t.Run(fmt.Sprint(test.object), func(t *testing.T) {
			Len(t, test.object, test.expLen)
		})
	}
}

func TestAssertLen_fails(t *testing.T) {
	tests := []struct {
		object any
		expLen int
	}{
		{object: "", expLen: 1},
		{object: "a", expLen: 2},
		{object: "123", expLen: 4},
		{object: []string{"", "", "", ""}, expLen: 5},
		{object: []string{}, expLen: 1},
		{object: []int{1, 2, 3, 4, 5}, expLen: 4},
		{object: 1, expLen: 4},
		{object: 0.1, expLen: 4},
		{object: generateStruct(), expLen: 1},
		{object: any(nil), expLen: 1},
		{object: true, expLen: 0},
		{object: uint(137), expLen: 0},
	}
	for _, test := range tests {
		t.Run(fmt.Sprint(test.object), func(t *testing.T) {
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				Len(t, test.object, test.expLen)
			})
		})
	}
}

func TestAssertLen_full(t *testing.T) {
	FuzzUtilRunTests(t, FuzzStringFull(), func(t *testing.T, index int, str string) {
		Len(t, str, len(str))
	})
}

func TestAssertIsIncreasing(t *testing.T) {
	tests := []struct {
		name   string
		object any
	}{
		{name: "int", object: []int{1, 2, 3, 4, 5, 6}},
		{name: "int2", object: []int{1, 7, 10, 11, 134, 700432}},
		{name: "int_negative", object: []int{-2, -1, 0, 1, 2, 3, 4, 5, 6}},
		{name: "int8", object: []int8{1, 2, 3, 4, 5, 6}},
		{name: "int16", object: []int16{1, 2, 3, 4, 5, 6}},
		{name: "int32", object: []int32{1, 2, 3, 4, 5, 6}},
		{name: "int64", object: []int64{1, 2, 3, 4, 5, 6}},
		{name: "uint", object: []uint{1, 2, 3, 4, 5, 6}},
		{name: "uint8", object: []uint8{1, 2, 3, 4, 5, 6}},
		{name: "uint16", object: []uint16{1, 2, 3, 4, 5, 6}},
		{name: "uint32", object: []uint32{1, 2, 3, 4, 5, 6}},
		{name: "uint64", object: []uint64{1, 2, 3, 4, 5, 6}},
		{name: "float32", object: []float32{1.0, 2.1, 3.0, 4.12345, 5.45623, 6.1}},
		{name: "float64", object: []float64{1.0, 2.1, 3.0, 4.12345, 5.45623, 6.1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Increasing(t, test.object)
		})
	}
}

func TestAssertIsIncreasing_fail(t *testing.T) {
	tests := []struct {
		name   string
		object any
	}{
		{name: "int_empty", object: []int{}},
		{name: "int_one", object: []int{4}},
		{name: "int", object: []int{4, 2, 3, 4, 5, 6}},
		{name: "int8", object: []int8{4, 2, 3, 4, 5, 6}},
		{name: "int16", object: []int16{4, 2, 3, 4, 5, 6}},
		{name: "int32", object: []int32{4, 2, 3, 4, 5, 6}},
		{name: "int64", object: []int64{4, 2, 3, 4, 5, 6}},
		{name: "uint", object: []uint{4, 2, 3, 4, 5, 6}},
		{name: "uint8", object: []uint8{4, 2, 3, 4, 5, 6}},
		{name: "uint16", object: []uint16{4, 2, 3, 4, 5, 6}},
		{name: "uint32", object: []uint32{4, 2, 3, 4, 5, 6}},
		{name: "uint64", object: []uint64{4, 2, 3, 4, 5, 6}},
		{name: "float32", object: []float32{4.0, 2.1, 3.0, 4.12345, 5.45623, 6.1}},
		{name: "float64", object: []float64{4.0, 2.1, 3.0, 4.12345, 5.45623, 6.1}},
		{name: "complex64", object: []complex64{complex64(complex(float64(4), 2)), complex64(complex(float64(2), 2)), complex64(complex(float64(4), 2))}},
		{name: "complex128", object: []complex128{complex(400.0, 2), complex(133.4, 2), complex(144.4, 2)}},
		{name: "stringSlice", object: []string{"a"}},
		{name: "string", object: "a"},
		{name: "bool", object: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				Increasing(t, test.object)
			})
		})
	}
}

func TestAssertIsDecreasing(t *testing.T) {
	tests := []struct {
		name   string
		object any
	}{
		{name: "int", object: []int{6, 5, 4, 3, 2, 1}},
		{name: "int2", object: []int{700432, 134, 11, 10, 7, 1}},
		{name: "int_negative", object: []int{6, 5, 4, 3, 2, 1, 0, -1, -2}},
		{name: "int8", object: []int8{6, 5, 4, 3, 2, 1}},
		{name: "int16", object: []int16{6, 5, 4, 3, 2, 1}},
		{name: "int32", object: []int32{6, 5, 4, 3, 2, 1}},
		{name: "int64", object: []int64{6, 5, 4, 3, 2, 1}},
		{name: "uint", object: []uint{6, 5, 4, 3, 2, 1}},
		{name: "uint8", object: []uint8{6, 5, 4, 3, 2, 1}},
		{name: "uint16", object: []uint16{6, 5, 4, 3, 2, 1}},
		{name: "uint32", object: []uint32{6, 5, 4, 3, 2, 1}},
		{name: "uint64", object: []uint64{6, 5, 4, 3, 2, 1}},
		{name: "float32", object: []float32{5.45623, 4.12345, 3.0, 2.1, 1.0}},
		{name: "float64", object: []float64{5.45623, 4.12345, 3.0, 2.1, 1.0}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Decreasing(t, test.object)
		})
	}
}

func TestAssertIsDecreasing_fail(t *testing.T) {
	tests := []struct {
		name   string
		object any
	}{
		{name: "int_empty", object: []int{}},
		{name: "int_one", object: []int{4}},
		{name: "int", object: []int{4, 2, 3, 4, 5, 6}},
		{name: "int8", object: []int8{4, 2, 3, 4, 5, 6}},
		{name: "int16", object: []int16{4, 2, 3, 4, 5, 6}},
		{name: "int32", object: []int32{4, 2, 3, 4, 5, 6}},
		{name: "int64", object: []int64{4, 2, 3, 4, 5, 6}},
		{name: "uint", object: []uint{4, 2, 3, 4, 5, 6}},
		{name: "uint8", object: []uint8{4, 2, 3, 4, 5, 6}},
		{name: "uint16", object: []uint16{4, 2, 3, 4, 5, 6}},
		{name: "uint32", object: []uint32{4, 2, 3, 4, 5, 6}},
		{name: "uint64", object: []uint64{4, 2, 3, 4, 5, 6}},
		{name: "float32", object: []float32{4.0, 2.1, 3.0, 4.12345, 5.45623, 6.1}},
		{name: "float64", object: []float64{4.0, 2.1, 3.0, 4.12345, 5.45623, 6.1}},
		{name: "complex64", object: []complex64{complex64(complex(float64(4), 2)), complex64(complex(float64(2), 2)), complex64(complex(float64(4), 2))}},
		{name: "complex128", object: []complex128{complex(400.0, 2), complex(133.4, 2), complex(144.4, 2)}},
		{name: "stringSlice", object: []string{"a"}},
		{name: "string", object: "a"},
		{name: "bool", object: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				Decreasing(t, test.object)
			})
		})
	}
}

func TestAssertSameElements(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		FuzzUtilRunTests(t, FuzzStringFull(), func(t *testing.T, index int, str string) {
			t.Helper()

			testEqual(t, str, str)
		})
	})

	t.Run("Nil and Nil are equal", func(t *testing.T) {
		var x []string
		SameElements(t, x, nil)
	})

	t.Run("Same arrays to same struct are equal", func(t *testing.T) {
		s := generateStruct()
		ss := []assertionTestStruct{s}
		SameElements(t, ss, ss)
	})

	t.Run("Same arrays to same pointer struct are equal", func(t *testing.T) {
		s := generateStruct()
		ss := []*assertionTestStruct{&s}
		SameElements(t, ss, ss)
	})

	t.Run("Structs", func(t *testing.T) {
		// Test ten random structs for equality.
		ss := make([]assertionTestStruct, 0)
		for i := 0; i < 10; i++ {
			s := generateStruct()
			ss = append(ss, s)
		}

		ss2 := ss
		SameElements(t, ss, ss2)
	})

	t.Run("Strings", func(t *testing.T) {
		s := []string{"Hello", "World"}
		ss := []string{"World", "Hello"}
		SameElements(t, s, ss)
	})

	t.Run("Integers", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5, 6, 7, 8}
		ss := []int{8, 7, 6, 5, 4, 3, 2, 1}
		SameElements(t, s, ss)
	})
}

func TestAssertSameElementsFails(t *testing.T) {
	t.Run("Structs", func(t *testing.T) {
		// Test ten random structs for equality.
		ss := make([]assertionTestStruct, 0)
		for i := 0; i < 10; i++ {
			s := generateStruct()
			ss = append(ss, s)
		}

		ss2 := ss
		ss2 = append(ss2, generateStruct())
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			SameElements(t, ss, ss2)
		})
	})

	t.Run("Strings", func(t *testing.T) {
		s := []string{"Hello", "World"}
		ss := []string{"World", "Hello", "Again"}
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			SameElements(t, s, ss)
		})
	})

	t.Run("Integers", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		ss := []int{9, 8, 7, 6}
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			SameElements(t, s, ss)
		})
	})
}

func TestAssertNotSameElements(t *testing.T) {
	t.Run("Structs", func(t *testing.T) {
		// Test ten random structs for equality.
		ss := make([]assertionTestStruct, 0)
		for i := 0; i < 10; i++ {
			s := generateStruct()
			ss = append(ss, s)
		}

		ss2 := ss
		ss2 = append(ss2, generateStruct())
		NotSameElements(t, ss, ss2)
	})

	t.Run("Strings", func(t *testing.T) {
		NotSameElements(
			t,
			[]string{"Hello", "World"},
			[]string{"World", "Hello", "Again"},
		)
	})

	t.Run("Integers", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5, 6, 7, 8}
		ss := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
		NotSameElements(t, s, ss)
	})
}

func TestAssertNotSameElementsFails(t *testing.T) {
	t.Run("Structs", func(t *testing.T) {
		// Test ten random structs for equality.
		ss := make([]assertionTestStruct, 0)
		for i := 0; i < 10; i++ {
			s := generateStruct()
			ss = append(ss, s)
		}

		ss2 := ss
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			NotSameElements(t, ss, ss2)
		})
	})

	t.Run("Strings", func(t *testing.T) {
		s := []string{"Hello", "World"}
		ss := []string{"World", "Hello"}
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			NotSameElements(t, s, ss)
		})
	})

	t.Run("Integers", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5, 6, 7, 8}
		ss := []int{8, 7, 6, 5, 4, 3, 2, 1}
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			NotSameElements(t, s, ss)
		})
	})
}

func TestAssertRegexp(t *testing.T) {
	Regexp(t, "p([a-z]+)ch", "peache")
	rxp, err := regexp.Compile("Hello, .*")
	NoError(t, err)
	Regexp(t, rxp, "Hello, World!")
}

func TestAssertRegexp_fail(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		Regexp(t, "p([a-z]+)ch", "peahe")
		rxp, err := regexp.Compile("Hello, .*")
		NoError(t, err)
		Regexp(t, rxp, "Hello World!")
	})
}

func TestAssertNotRegexp(t *testing.T) {
	NotRegexp(t, "p([a-z]+)ch", "peahe")
	rxp, err := regexp.Compile("Hello, .*")
	NoError(t, err)
	NotRegexp(t, rxp, "Hello World!")
}

func TestAssertNotRegexp_fail(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		NotRegexp(t, "p([a-z]+)ch", "peache")
		rxp, err := regexp.Compile("Hello, .*")
		NoError(t, err)
		NotRegexp(t, rxp, "Hello, World!")
	})
}

func TestAssertFileExists(t *testing.T) {
	t.Run("LICENSE", func(t *testing.T) {
		FileExists(t, "LICENSE")
	})

	t.Run("README.md", func(t *testing.T) {
		FileExists(t, "README.md")
	})

	t.Run("CODE_OF_CONDUCT.md", func(t *testing.T) {
		FileExists(t, "CODE_OF_CONDUCT.md")
	})

	t.Run("CONTRIBUTING.md", func(t *testing.T) {
		FileExists(t, "CONTRIBUTING.md")
	})

	t.Run("CHANGELOG.md", func(t *testing.T) {
		FileExists(t, "CHANGELOG.md")
	})
}

func TestAssertFileExists_fail(t *testing.T) {
	t.Run("asdasdasdasd.md", func(t *testing.T) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			FileExists(t, "asdasdasdasd.md")
		})
	})
}

func TestAssertNoFileExists(t *testing.T) {
	t.Run("asdasdasdasd.md", func(t *testing.T) {
		NoFileExists(t, "asdasdasdasd.md")
	})
}

func TestAssertNoFileExists_fail(t *testing.T) {
	t.Run("LICENSE", func(t *testing.T) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			NoFileExists(t, "LICENSE")
		})
	})

	t.Run("README.md", func(t *testing.T) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			NoFileExists(t, "README.md")
		})
	})

	t.Run("CODE_OF_CONDUCT.md", func(t *testing.T) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			NoFileExists(t, "CODE_OF_CONDUCT.md")
		})
	})

	t.Run("CONTRIBUTING.md", func(t *testing.T) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			NoFileExists(t, "CONTRIBUTING.md")
		})
	})

	t.Run("CHANGELOG.md", func(t *testing.T) {
		TestFails(t, func(t TestingPackageWithFailFunctions) {
			NoFileExists(t, "CHANGELOG.md")
		})
	})
}

func TestAssertDirExists(t *testing.T) {
	for _, dir := range []string{"ci", "internal", "testdata", os.TempDir()} {
		t.Run(dir, func(t *testing.T) {
			DirExists(t, dir)
		})
	}
}

func TestAssertDirExists_fail(t *testing.T) {
	for _, dir := range []string{"asdasdasdasd", "LICENSE", "CHANGELOG.md"} {
		t.Run(dir, func(t *testing.T) {
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				DirExists(t, dir)
			})
		})
	}
}

func TestAssertNoDirExists(t *testing.T) {
	for _, dir := range []string{"asdasdasdasd", "LICENSE", "CHANGELOG.md"} {
		t.Run(dir, func(t *testing.T) {
			NoDirExists(t, dir)
		})
	}
}

func TestAssertNoDirExists_fail(t *testing.T) {
	for _, dir := range []string{"ci", "internal", "testdata", os.TempDir()} {
		t.Run(dir, func(t *testing.T) {
			TestFails(t, func(t TestingPackageWithFailFunctions) {
				NoDirExists(t, dir)
			})
		})
	}
}

func TestAssertDirEmpty(t *testing.T) {
	tempdir, _ := ioutil.TempDir("testdata", "temp")
	DirEmpty(t, tempdir)
	defer os.RemoveAll(tempdir)
}

func TestAssertDirEmpty_fail(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		DirEmpty(t, "non-existent")
		DirEmpty(t, "internal")
	})
}

func TestAssertDirNotEmpty(t *testing.T) {
	DirNotEmpty(t, "testdata")
}

func TestAssertDirNotEmpty_fail(t *testing.T) {
	tempdir, _ := ioutil.TempDir("testdata", "temp")
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		DirNotEmpty(t, tempdir)
	})
	defer os.RemoveAll(tempdir)
}

func TestAssertSubset(t *testing.T) {
	t.Run("Simple integer subset", func(t *testing.T) {
		Subset(t, []int{1, 2, 3}, []int{1, 2})
	})

	t.Run("Simple string subset", func(t *testing.T) {
		Subset(t, []string{"Hello", "World", "Test"}, []string{"Test", "World"})
	})

	t.Run("Complex struct subset", func(t *testing.T) {
		s1 := generateStruct()
		s2 := generateStruct()
		s3 := generateStruct()
		Subset(t, []assertionTestStruct{s1, s2, s3}, []assertionTestStruct{s2, s3})
	})
}

func TestAssertSubset_fail(t *testing.T) {
	t.Run("Simple integer subset", func(t *testing.T) {
		NoSubset(t, []int{1, 2, 3}, []int{1, 7})
	})

	t.Run("Simple string subset", func(t *testing.T) {
		NoSubset(t, []string{"Hello", "World", "Test"}, []string{"Test", "John"})
	})

	t.Run("Complex struct subset", func(t *testing.T) {
		s1 := generateStruct()
		s2 := generateStruct()
		s3 := generateStruct()
		s4 := generateStruct()
		NoSubset(t, []assertionTestStruct{s1, s2, s3}, []assertionTestStruct{s3, s4})
	})
}

// -- Assert output consistency --
func TestOutputConsistency(t *testing.T) {
	var tm testMock

	t.Run("failed test", func(t *testing.T) {
		pterm.DisableStyling()
		Equal(&tm, true, false)
		pterm.EnableStyling()

		err := SnapshotCreateOrValidate(t, t.Name(), tm.ErrorMessage)
		NoError(t, err)
	})

	t.Run("failed test with custom msg", func(t *testing.T) {
		pterm.DisableStyling()
		Equal(&tm, true, false, "Custom message!")
		pterm.EnableStyling()

		err := SnapshotCreateOrValidate(t, t.Name(), tm.ErrorMessage)
		NoError(t, err)
	})

	t.Run("failed test with custom formatted msg", func(t *testing.T) {
		pterm.DisableStyling()
		Equal(&tm, true, false, "Custom %s!", "message")
		pterm.EnableStyling()

		err := SnapshotCreateOrValidate(t, t.Name(), tm.ErrorMessage)
		NoError(t, err)
	})

	t.Run("failed test string", func(t *testing.T) {
		const expected = `Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Fusce consectetur quam id turpis blandit, pulvinar fermentum justo sollicitudin.
Curabitur vehicula eros posuere efficitur egestas.

Duis et lectus in nisi mattis convallis non a odio.
Proin pulvinar felis consectetur condimentum tincidunt.`

		const actual = `Lorem ipsum sit amet, consechbtetur adipiscing elit.z

zFusce consecetur quam iod turpis blandit, pulvnar fermentum justo sollicitudin.
Curdabitur veicular eros posuere hello efficitulr egestas.
Duis et lectus in nisi mattis convallis non a odio.
Proin puvinar feliss consectetur codiementum tincidunt.`

		pterm.DisableStyling()
		Equal(&tm, expected, actual)
		pterm.EnableStyling()

		err := SnapshotCreateOrValidate(t, t.Name(), tm.ErrorMessage)
		NoError(t, err)
	})

	t.Run("failed test struct", func(t *testing.T) {
		pterm.DisableStyling()
		Equal(&tm, assertionTestStruct{
			Name: "John",
			Age:  34,
			Meta: assertionTestStructNested{
				ID:    512345,
				Admin: false,
			},
		}, assertionTestStruct{
			Name: "Bob",
			Age:  22,
			Meta: assertionTestStructNested{
				ID:    123456,
				Admin: true,
			},
		})
		pterm.EnableStyling()

		err := SnapshotCreateOrValidate(t, t.Name(), tm.ErrorMessage)
		NoError(t, err)
	})

	t.Run("failed test slice", func(t *testing.T) {
		pterm.DisableStyling()
		Equal(&tm, nil, []int{1, 2, 3})
		pterm.EnableStyling()

		err := SnapshotCreateOrValidate(t, t.Name(), tm.ErrorMessage)
		NoError(t, err)
	})

	t.Run("failed test newlines", func(t *testing.T) {
		pterm.DisableStyling()
		Equal(&tm, `1234567890`, `1
2
3
4
5
6
7
8
9
0`)
		pterm.EnableStyling()

		err := SnapshotCreateOrValidate(t, t.Name(), tm.ErrorMessage)
		NoError(t, err)
	})

	t.Run("failed test newline start", func(t *testing.T) {
		pterm.DisableStyling()
		Equal(&tm, `1`, `
1`)
		pterm.EnableStyling()

		err := SnapshotCreateOrValidate(t, t.Name(), tm.ErrorMessage)
		NoError(t, err)
	})

	t.Run("failed test newline end", func(t *testing.T) {
		pterm.DisableStyling()
		Equal(&tm, `1`, `1
`)
		pterm.EnableStyling()

		err := SnapshotCreateOrValidate(t, t.Name(), tm.ErrorMessage)
		NoError(t, err)
	})

	t.Run("failed test many matching lines", func(t *testing.T) {
		dataExpected := make([]string, 50)
		dataActual := make([]string, 50)
		for i := 0; i < len(dataExpected); i++ {
			dataExpected[i] = strconv.Itoa(i)
			dataActual[i] = dataExpected[i]
		}

		dataActual[10] = "hello"
		dataActual[20] = "world"
		dataActual[30] = "foo"
		dataActual[40] = "bar"

		pterm.DisableStyling()
		Equal(&tm, strings.Join(dataExpected, "\n"), strings.Join(dataActual, "\n"))
		pterm.EnableStyling()

		err := SnapshotCreateOrValidate(t, t.Name(), tm.ErrorMessage)
		NoError(t, err)
	})
}

func TestAssertUnique(t *testing.T) {
	Unique(t, []int{1, 2, 3, 4, 5})
	Unique(t, []string{"foo", "bar", "baz"})
	Unique(t, []bool{true, false})
	Unique(t, []float64{1.1, 1.2, 1.3, 1.4, 1.5})
}

func TestAssertUnique_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		Unique(t, []int{1, 2, 3, 4, 5, 1})
	})

	TestFails(t, func(t TestingPackageWithFailFunctions) {
		Unique(t, []string{"foo", "bar", "baz", "foo"})
	})

	TestFails(t, func(t TestingPackageWithFailFunctions) {
		Unique(t, []bool{true, false, true})
	})

	TestFails(t, func(t TestingPackageWithFailFunctions) {
		Unique(t, []float64{1.1, 1.2, 1.3, 1.4, 1.5, 1.1})
	})
}

func TestAssertNotUnique(t *testing.T) {
	NotUnique(t, []int{1, 2, 3, 4, 5, 1})
	NotUnique(t, []string{"foo", "bar", "baz", "foo"})
	NotUnique(t, []bool{true, false, true})
	NotUnique(t, []float64{1.1, 1.2, 1.3, 1.4, 1.5, 1.1})
}

func TestAssertNotUnique_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		NotUnique(t, []int{1, 2, 3, 4, 5})
	})

	TestFails(t, func(t TestingPackageWithFailFunctions) {
		NotUnique(t, []string{"foo", "bar", "baz"})
	})

	TestFails(t, func(t TestingPackageWithFailFunctions) {
		NotUnique(t, []bool{true, false})
	})

	TestFails(t, func(t TestingPackageWithFailFunctions) {
		NotUnique(t, []float64{1.1, 1.2, 1.3, 1.4, 1.5})
	})
}

func TestAssertInRange(t *testing.T) {
	InRange(t, 2, 1, 3)
	InRange(t, 1, 1, 3)
}

func TestAssertInRange_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		NotInRange(t, 1, 1, 3)
	})

	TestFails(t, func(t TestingPackageWithFailFunctions) {
		NotInRange(t, 2, 1, 3)
	})
}

func TestJsonEqual(t *testing.T) {
	JSONEqual(t, `{"foo": "bar"}`, `{"foo": "bar"}`)
}

func TestJsonEqual_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		JSONEqual(t, `{"foo": "bar"}`, `{"foo": "baz"}`)
	})
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		JSONEqual(t, `{"foo": "bar"}`, `{"foo": "baa", "bar": "foo"}`)
	})
}

func TestAssertNotInRange(t *testing.T) {
	NotInRange(t, 4, 1, 3)
	NotInRange(t, 0, 1, 3)
}

func TestFailNow(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		FailNow(t, "This is a test")
	})
}

func TestAssertNotInRange_fails(t *testing.T) {
	TestFails(t, func(t TestingPackageWithFailFunctions) {
		InRange(t, 4, 1, 3)
	})

	TestFails(t, func(t TestingPackageWithFailFunctions) {
		InRange(t, 0, 1, 3)
	})
}
