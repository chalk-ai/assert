package assert

import (
	"atomicgo.dev/assert"
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"
	"os"
	"reflect"
	"time"

	"github.com/pterm/pterm"

	"github.com/chalk-ai/assert/internal"
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

// KindOf asserts that the object is a type of kind exptectedKind.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertKindOf(t, reflect.Slice, []int{1,2,3})
//	assert.AssertKindOf(t, reflect.Slice, []string{"Hello", "World"})
//	assert.AssertKindOf(t, reflect.Int, 1337)
//	assert.AssertKindOf(t, reflect.Bool, true)
//	assert.AssertKindOf(t, reflect.Map, map[string]bool{})
func KindOf(t testRunner, expectedKind reflect.Kind, object any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !assert.Kind(object, expectedKind) {
		internal.Fail(t,
			fmt.Sprintf("A value that !!should be a type of kind %s!! is a type of kind %s.", expectedKind.String(), reflect.TypeOf(object).Kind().String()),
			internal.NewObjectsExpectedActual(expectedKind, object),
			msg...,
		)
	}
}

// NotKindOf asserts that the object is not a type of kind `kind`.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNotKindOf(t, reflect.Slice, "Hello, World")
//	assert.AssertNotKindOf(t, reflect.Slice, true)
//	assert.AssertNotKindOf(t, reflect.Int, 13.37)
//	assert.AssertNotKindOf(t, reflect.Bool, map[string]bool{})
//	assert.AssertNotKindOf(t, reflect.Map, false)
func NotKindOf(t testRunner, kind reflect.Kind, object any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if assert.Kind(object, kind) {
		internal.Fail(t,
			fmt.Sprintf("A value that !!should not be a type of kind %s!! is a type of kind %s.", kind.String(), reflect.TypeOf(object).Kind().String()),
			internal.Objects{
				internal.NewObjectsSingleNamed("Should not be", kind)[0],
				internal.NewObjectsSingleNamed("Actual", object)[0],
			},
			msg...,
		)
	}
}

// Numeric asserts that the object is a numeric type.
// Numeric types are:
// Int, Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16, Uint32, Uint64, Complex64 and Complex128.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNumeric(t, 123)
//	assert.AssertNumeric(t, 1.23)
//	assert.AssertNumeric(t, uint(123))
func Numeric(t testRunner, object any, msg ...any) {
	if !assert.Number(object) {
		internal.Fail(t, "An object that !!should be a number!! is not of a numeric type.", internal.NewObjectsSingleUnknown(object), msg...)
	}
}

// NotNumeric checks if the object is not a numeric type.
// Numeric types are:
// Int, Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16, Uint32, Uint64, Complex64 and Complex128.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNotNumeric(t, true)
//	assert.AssertNotNumeric(t, "123")
func NotNumeric(t testRunner, object any, msg ...any) {
	if assert.Number(object) {
		internal.Fail(t, "An object that !!should not be a number!! is of a numeric type.", internal.NewObjectsSingleUnknown(object), msg...)
	}
}

// Zero asserts that the value is the zero value for it's type.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertZero(t, 0)
//	assert.AssertZero(t, false)
//	assert.AssertZero(t, "")
func Zero(t testRunner, value any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !assert.Zero(value) {
		internal.Fail(t, "An object that !!should have its zero value!!, does not have its zero value.", internal.NewObjectsSingleUnknown(value), msg...)
	}
}

// NotZero asserts that the value is not the zero value for it's type.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNotZero(t, 1337)
//	assert.AssertNotZero(t, true)
//	assert.AssertNotZero(t, "Hello, World")
func NotZero(t testRunner, value any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if assert.Zero(value) {
		internal.Fail(t, "An object that !!should not have its zero value!!, does have its zero value.", internal.NewObjectsSingleUnknown(value), msg...)
	}
}

// Equal asserts that two objects are equal.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertEqual(t, "Hello, World!", "Hello, World!")
//	assert.AssertEqual(t, true, true)
func Equal(t testRunner, expected any, actual any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !assert.Equal(expected, actual) {
		internal.Fail(t, "Two objects that !!should be equal!!, are not equal.", internal.NewObjectsExpectedActualWithDiff(expected, actual), msg...)
	}
}

// NotEqual asserts that two objects are not equal.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNotEqual(t, true, false)
//	assert.AssertNotEqual(t, "Hello", "World")
func NotEqual(t testRunner, expected any, actual any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if assert.Equal(expected, actual) {
		objects := internal.Objects{
			{
				Name:      "Both Objects",
				NameStyle: pterm.NewStyle(pterm.FgMagenta),
				Data:      expected,
			},
		}
		internal.Fail(t, "Two objects that !!should not be equal!!, are equal.", objects, msg...)
	}
}

// EqualValues asserts that two objects have equal values.
// The order of the values is also validated.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertEqualValues(t, []string{"Hello", "World"}, []string{"Hello", "World"})
//	assert.AssertEqualValues(t, []int{1,2}, []int{1,2})
//	assert.AssertEqualValues(t, []int{1,2}, []int{2,1}) // FAILS (wrong order)
//
// Comparing struct values:
//
//	person1 := Person{
//	  Name:   "Marvin Wendt",
//	  Age:    20,
//	  Gender: "male",
//	}
//
//	person2 := Person{
//	  Name:   "Marvin Wendt",
//	  Age:    20,
//	  Gender: "male",
//	}
//
//	assert.AssertEqualValues(t, person1, person2)
func EqualValues(t testRunner, expected any, actual any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !internal.HasEqualValues(expected, actual) {
		internal.Fail(t, "Two objects that !!should have equal values!!, do not have equal values.", internal.NewObjectsExpectedActualWithDiff(expected, actual), msg...)
	}
}

// NotEqualValues asserts that two objects do not have equal values.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNotEqualValues(t, []int{1,2}, []int{3,4})
//
// Comparing struct values:
//
//	person1 := Person{
//	  Name:   "Marvin Wendt",
//	  Age:    20,
//	  Gender: "male",
//	}
//
//	person2 := Person{
//	  Name:   "Marvin Wendt",
//	  Age:    20,
//	  Gender: "female", // <-- CHANGED
//	}
//
//	assert.AssertNotEqualValues(t, person1, person2)
func NotEqualValues(t testRunner, expected any, actual any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if internal.HasEqualValues(expected, actual) {
		internal.Fail(t, "Two objects that !!should not have equal values!!, do have equal values.", internal.NewObjectsSingleNamed("Both Objects", actual), msg...)
	}
}

// True asserts that an expression or object resolves to true.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertTrue(t, true)
//	assert.AssertTrue(t, 1 == 1)
//	assert.AssertTrue(t, 2 != 3)
//	assert.AssertTrue(t, 1 > 0 && 4 < 5)
func True(t testRunner, value any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if value != true {
		internal.Fail(t, "Value !!should be true!! but is not.", internal.NewObjectsExpectedActual(true, value), msg...)
	}
}

// False asserts that an expression or object resolves to false.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertFalse(t, false)
//	assert.AssertFalse(t, 1 == 2)
//	assert.AssertFalse(t, 2 != 2)
//	assert.AssertFalse(t, 1 > 5 && 4 < 0)
func False(t testRunner, value any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if value == true {
		internal.Fail(t, "Value !!should be false!! but is not.", internal.NewObjectsExpectedActual(false, value), msg...)
	}
}

// Implements asserts that an objects implements an interface.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertImplements(t, (*YourInterface)(nil), new(YourObject))
//	assert.AssertImplements(t, (*fmt.Stringer)(nil), new(types.Const)) => pass
func Implements(t testRunner, interfaceObject, object any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !assert.Implements(object, interfaceObject) {
		internal.Fail(t, fmt.Sprintf("An object that !!should implement %s!! does not implement it.", reflect.TypeOf(interfaceObject).String()), internal.Objects{}, msg...)
	}
}

// NotImplements asserts that an object does not implement an interface.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNotImplements(t, (*YourInterface)(nil), new(YourObject))
//	assert.AssertNotImplements(t, (*fmt.Stringer)(nil), new(types.Const)) => fail, because types.Const does implement fmt.Stringer.
func NotImplements(t testRunner, interfaceObject, object any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if assert.Implements(object, interfaceObject) {
		internal.Fail(t, fmt.Sprintf("An object that !!should not implement %s!! does implement it.", reflect.TypeOf(interfaceObject).String()), internal.Objects{}, msg...)
	}
}

// Contains asserts that a string/list/array/slice/map contains the specified element.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertContains(t, []int{1,2,3}, 2)
//	assert.AssertContains(t, []string{"Hello", "World"}, "World")
//	assert.AssertContains(t, "Hello, World!", "World")
func Contains(t testRunner, object, element any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !assert.Contains(object, element) {
		internal.Fail(t, "An object !!does not contain!! the object it should contain.", internal.Objects{
			internal.NewObjectsSingleNamed("Missing Object", element)[0],
			internal.NewObjectsSingleNamed("Full Object", object)[0],
		}, msg...)
	}
}

// NotContains asserts that a string/list/array/slice/map does not contain the specified element.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNotContains(t, []string{"Hello", "World"}, "Spaceship")
//	assert.AssertNotContains(t, "Hello, World!", "Spaceship")
func NotContains(t testRunner, object, element any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if assert.Contains(object, element) {
		internal.Fail(t, "An object !!does contain!! an object it should not contain.", internal.Objects{
			internal.NewObjectsSingleUnknown(object)[0],
			internal.NewObjectsSingleNamed("Element that should not be in the object", element)[0],
		}, msg...)
	}
}

// Panics asserts that a function panics.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertPanics(t, func() {
//		// ...
//		panic("some panic")
//	}) // => PASS
func Panics(t testRunner, f func(), msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !assert.Panic(f) {
		internal.Fail(t, "A function that !!should panic!! did not panic.", internal.Objects{}, msg...)
	}
}

// NotPanics asserts that a function does not panic.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNotPanics(t, func() {
//		// some code that does not call a panic...
//	}) // => PASS
func NotPanics(t testRunner, f func(), msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if assert.Panic(f) {
		internal.Fail(t, "A function that !!should not panic!! did panic.", internal.Objects{}, msg...)
	}
}

// Nil asserts that an object is nil.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNil(t, nil)
func Nil(t testRunner, object any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !assert.Nil(object) {
		internal.Fail(t, "An object that !!should be nil!! is not nil.", internal.NewObjectsExpectedActual(nil, object), msg...)
	}
}

// NotNil asserts that an object is not nil.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNotNil(t, true)
//	assert.AssertNotNil(t, "Hello, World!")
//	assert.AssertNotNil(t, 0)
func NotNil(t testRunner, object any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if assert.Nil(object) {
		internal.Fail(t, "An object that !!should not be nil!! is nil.", internal.NewObjectsSingleUnknown(object), msg...)
	}
}

// CompletesIn asserts that a function completes in a given time.
// Use this function to test that functions do not take too long to complete.
//
// NOTE: Every system takes a different amount of time to complete a function.
// Do not set the duration too low, if you want consistent results.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertCompletesIn(t, 2 * time.Second, func() {
//		// some code that should take less than 2 seconds...
//	}) // => PASS
func CompletesIn(t testRunner, duration time.Duration, f func(), msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !internal.CompletesIn(duration, f) {
		internal.Fail(t, fmt.Sprintf("The function !!should complete in %s!!, but it did not.", duration), internal.Objects{}, msg...)
	}
}

// NotCompletesIn asserts that a function does not complete in a given time.
// Use this function to test that functions do not complete to quickly.
// For example if your database connection completes in under a millisecond, there might be something wrong.
//
// NOTE: Every system takes a different amount of time to complete a function.
// Do not set the duration too high, if you want consistent results.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNotCompletesIn(t, 2 * time.Second, func() {
//		// some code that should take more than 2 seconds...
//		time.Sleep(3 * time.Second)
//	}) // => PASS
func NotCompletesIn(t testRunner, duration time.Duration, f func(), msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if internal.CompletesIn(duration, f) {
		internal.Fail(t, fmt.Sprintf("The function !!should not complete in %s!!, but it did.", duration), internal.Objects{}, msg...)
	}
}

// NoError asserts that an error is nil.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	err := nil
//	assert.AssertNoError(t, err)
func NoError(t testRunner, err error, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if err != nil {
		internal.Fail(t, "An error that !!should be nil!! is not nil.", internal.Objects{
			{
				Name:      "Error message",
				NameStyle: pterm.NewStyle(pterm.FgLightRed, pterm.Bold),
				Data:      fmt.Sprintf("%q\n", err.Error()),
				DataStyle: pterm.NewStyle(pterm.FgRed),
				Raw:       true,
			},
			{
				Name:      "Error object",
				NameStyle: pterm.NewStyle(pterm.FgLightRed, pterm.Bold),
				Data:      err,
				DataStyle: pterm.NewStyle(pterm.FgRed),
			}}, msg...)
		t.FailNow()
	}
}

// Error asserts that an error is not nil.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	err := errors.New("hello world")
//	assert.AssertError(t, err)
func Error(t testRunner, err error, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if err == nil {
		internal.Fail(t, "An error that !!should not be nil!! is nil.", internal.Objects{}, msg...)
	}
}

// Greater asserts that the first object is greater than the second.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertGreater(t, 5, 1)
//	assert.AssertGreater(t, 10, -10)
func Greater[T constraints.Ordered](t testRunner, object1, object2 T, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if object1 <= object2 {
		internal.Fail(
			t,
			"An object that !!should be greater!! than the second object is not.",
			internal.Objects{
				{Name: "Object 1", Data: object1},
				{Name: "Should be greater than object 2", Data: object2},
			},
			msg...,
		)
	}
}

// GreaterOrEqual asserts that the first object is greater than or equal to the second.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertGreaterOrEqual(t, 5, 1)
//	assert.AssertGreaterOrEqual(t, 10, -10)
//
// assert.AssertGreaterOrEqual(t, 10, 10)
func GreaterOrEqual[T constraints.Ordered](t testRunner, object1, object2 T, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if object1 < object2 {
		internal.Fail(
			t,
			"An object that !!should be greater!! than the second object is not.",
			internal.Objects{
				{Name: "Object 1", Data: object1},
				{Name: "Should be greater than object 2", Data: object2},
			},
			msg...,
		)
	}
}

// Less asserts that the first object is less than the second.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertLess(t, 1, 5)
//	assert.AssertLess(t, -10, 10)
func Less[T constraints.Ordered](t testRunner, object1, object2 T, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !(object1 < object2) {
		internal.Fail(t, "An object that !!should be less!! than the second object is not.", internal.Objects{
			internal.NewObjectsSingleNamed("Should be less than", object1)[0],
			internal.NewObjectsSingleNamed("Actual", object2)[0],
		}, msg...)
	}
}

// LessOrEqual asserts that the first object is less than or equal to the second.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertLessOrEqual(t, 1, 5)
//	assert.AssertLessOrEqual(t, -10, 10)
//	assert.AssertLessOrEqual(t, 1, 1)
func LessOrEqual[T constraints.Ordered](t testRunner, v1, v2 T, msg ...interface{}) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !(v1 <= v2) {
		internal.Fail(t, "An object that !!should be less or equal!! than the second object is not.", internal.Objects{
			internal.NewObjectsSingleNamed("Should be less or equal to", v1)[0],
			internal.NewObjectsSingleNamed("Actual", v2)[0],
		}, msg...)
	}
}

// TestFails asserts that a unit test fails.
// A unit test fails if one of the following methods is called in the test function: Error, Errorf, Fail, FailNow, Fatal, Fatalf
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertTestFails(t, func(t assert.TestingPackageWithFailFunctions) {
//		assert.AssertTrue(t, false)
//	}) // => Pass
//
//	assert.AssertTestFails(t, func(t assert.TestingPackageWithFailFunctions) {
//		// ...
//		t.Fail() // Or any other failing method.
//	}) // => Pass
func TestFails(t testRunner, test func(t TestingPackageWithFailFunctions), msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	var mock testMock
	test(&mock)

	if !mock.ErrorCalled {
		internal.Fail(t, "A test that !!should fail!! did not fail.", []internal.Object{}, msg...)
	}
}

// ErrorIs asserts that target is inside the error chain of err.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	var testErr = errors.New("hello world")
//	var testErrWrapped = fmt.Errorf("test err: %w", testErr)
//	assert.AssertErrorIs(t, testErrWrapped ,testErr)
func ErrorIs(t testRunner, err, target error, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !errors.Is(err, target) {
		internal.Fail(t, "Target error !!should be in the error chain!! of err.", internal.NewObjectsExpectedActual(target.Error(), err.Error()), msg...)
	}
}

// NotErrorIs
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	var testErr = errors.New("hello world")
//	var test2Err = errors.New("hello world 2")
//	var testErrWrapped = fmt.Errorf("test err: %w", testErr)
//	assert.AssertNotErrorIs(t, testErrWrapped, test2Err)
func NotErrorIs(t testRunner, err, target error, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if errors.Is(err, target) {
		internal.Fail(t, "Target error !!should not be in the error chain!! of err.", internal.NewObjectsExpectedActual(target.Error(), err.Error()), msg...)
	}
}

// Len asserts that the length of an object is equal to the given length.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertLen(t, "abc", 3)
//	assert.AssertLen(t, "Assert", 6)
//	assert.AssertLen(t, []int{1, 2, 1337, 25}, 4)
//	assert.AssertLen(t, map[string]int{"asd": 1, "test": 1337}, 2)
func Len(t testRunner, object any, length int, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	v := reflect.ValueOf(object)
	defer func() {
		if e := recover(); e != nil {
			internal.Fail(t, "The 'object' !!does not!! have a length.", internal.NewObjectsSingleUnknown(object), msg...)
		}
	}()

	if v.Len() != length {
		internal.Fail(t, "The length of 'object' !!is not!! the expected length.", internal.Objects{
			{
				Name:      "Expected length",
				NameStyle: pterm.NewStyle(pterm.FgLightGreen),
				Data:      fmt.Sprint(length) + "\n",
				DataStyle: pterm.NewStyle(pterm.FgGreen),
				Raw:       true,
			},
			{
				Name:      "Actual length",
				NameStyle: pterm.NewStyle(pterm.FgLightRed),
				Data:      fmt.Sprint(v.Len()) + "\n",
				DataStyle: pterm.NewStyle(pterm.FgRed),
				Raw:       true,
			},
			internal.NewObjectsSingleUnknown(object)[0],
		}, msg...)
	}
}

// Increasing asserts that the values in a slice are increasing.
// the test fails if the values are not in a slice or if the values are not comparable.
//
// Valid input kinds are: []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64, []float32, []float64.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertIncreasing(t, []int{1, 2, 137, 1000})
//	assert.AssertIncreasing(t, []float32{-10.3, 0.1, 7, 13.5})
func Increasing(t testRunner, object any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	internal.AssertCompareHelper(t, object, 1, msg...)
}

// Decreasing asserts that the values in a slice are decreasing.
// the test fails if the values are not in a slice or if the values are not comparable.
//
// Valid input kinds are: []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64, []float32, []float64.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertDecreasing(t, []int{1000, 137, 2, 1})
//	assert.AssertDecreasing(t, []float32{13.5, 7, 0.1, -10.3})
func Decreasing(t testRunner, object any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	internal.AssertCompareHelper(t, object, -1, msg...)
}

// Regexp asserts that a string matches a given regexp.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertRegexp(t, "^a.*c$", "abc")
func Regexp(t testRunner, regex any, txt any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	internal.AssertRegexpHelper(t, regex, txt, true, msg...)
}

// NotRegexp asserts that a string does not match a given regexp.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNotRegexp(t, "ab.*", "Hello, World!")
func NotRegexp(t testRunner, regex any, txt any, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	internal.AssertRegexpHelper(t, regex, txt, false, msg...)
}

// FileExists asserts that a file exists.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertFileExists(t, "./test.txt")
//	assert.AssertFileExists(t, "./config.yaml", "the config file is missing")
func FileExists(t testRunner, file string, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	// check if a file does not exists
	if _, err := os.Stat(file); os.IsNotExist(err) {
		internal.Fail(t, "A file !!does not exist!!.", internal.NewObjectsSingleNamed("File", file), msg...)
	}
}

func NoFileExists(t testRunner, file string, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	// check if a file exists
	if _, err := os.Stat(file); !os.IsNotExist(err) {
		internal.Fail(t, "A file that !!should not exist!!, does exist.", internal.NewObjectsSingleUnknown(file), msg...)
	}
}

// DirExists asserts that a directory exists.
// The test will pass when the directory exists, and it's visible to the current user.
// The test will fail, if the path points to a file.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertDirExists(t, "FolderName")
func DirExists(t testRunner, dir string, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		internal.Fail(t, "A directory !!does not exist!!.", internal.NewObjectsSingleNamed("Dir", dir), msg...)
	} else if !stat.IsDir() {
		internal.Fail(t, "A file !!is not a directory!!.", internal.NewObjectsSingleNamed("Dir", dir), msg...)
	}
}

// NoDirExists asserts that a directory does not exists.
// The test will pass, if the path points to a file, as a directory with the same name, cannot exist.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNoDirExists(t, "FolderName")
func NoDirExists(t testRunner, dir string, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return
	}
	if stat.IsDir() {
		internal.Fail(t, "A directory that !!should not exist!!, does exist.", internal.NewObjectsSingleUnknown(dir), msg...)
	}
}

// DirEmpty asserts that a directory is empty.
// The test will pass when the directory is empty or does not exist.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertDirEmpty(t, "FolderName")
func DirEmpty(t testRunner, dir string, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !internal.AssertDirEmptyHelper(t, dir) {
		internal.Fail(t, "The directory !!is not!! empty.", internal.NewObjectsSingleNamed("Directory", dir), msg...)
	}
}

// DirNotEmpty asserts that a directory is not empty
// The test will pass when the directory is not empty and will fail if the directory does not exist.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertDirNotEmpty(t, "FolderName")
func DirNotEmpty(t testRunner, dir string, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if internal.AssertDirEmptyHelper(t, dir) {
		internal.Fail(t, "The directory !!is!! empty.", internal.NewObjectsSingleNamed("Directory", dir), msg...)
	}
}

// SameElements asserts that two slices contains same elements (including pointers).
// The order is irrelevant.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	 assert.AssertSameElements(t, []string{"Hello", "World"}, []string{"Hello", "World"})
//	 assert.AssertSameElements(t, []int{1,2,3}, []int{1,2,3})
//	 assert.AssertSameElements(t, []int{1,2}, []int{2,1})
//
//	 type A struct {
//		  a string
//	 }
//	 assert.AssertSameElements(t, []*A{{a: "A"}, {a: "B"}, {a: "C"}}, []*A{{a: "A"}, {a: "B"}, {a: "C"}})
func SameElements[T comparable](t testRunner, expected []T, actual []T, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !internal.HasSameElements(expected, actual) {
		internal.Fail(t, "Two objects that !!should have the same elements!!, do not have the same elements.", internal.NewObjectsExpectedActualWithDiff(expected, actual), msg...)
	}
}

// NotSameElements asserts that two slices contains same elements (including pointers).
// The order is irrelevant.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	 assert.AssertNotSameElements(t, []string{"Hello", "World"}, []string{"Hello", "World", "World"})
//	 assert.AssertNotSameElements(t, []int{1,2}, []int{1,2,3})
//
//	 type A struct {
//		  a string
//	 }
//	 assert.AssertNotSameElements(t, []*A{{a: "A"}, {a: "B"}, {a: "C"}}, []*A{{a: "A"}, {a: "B"}, {a: "C"}, {a: "D"}})
func NotSameElements[T comparable](t testRunner, expected []T, actual []T, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if internal.HasSameElements(expected, actual) {
		internal.Fail(t, "Two objects that !!should have the same elements!!, do not have the same elements.", internal.NewObjectsSingleNamed("Both Objects", actual), msg...)
	}
}

// Subset asserts that the second parameter is a subset of the list.
// The order is irrelevant.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertSubset(t, []int{1, 2, 3}, []int{1, 2})
//	assert.AssertSubset(t, []string{"Hello", "World", "Test"}, []string{"Test", "World"})
func Subset[T comparable](t testRunner, list []T, subset []T, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !internal.IsSubset(t, list, subset) {
		internal.Fail(t, "The second parameter !!is not a subset of the list!!, but should be.", internal.Objects{internal.NewObjectsSingleNamed("List", list)[0], internal.NewObjectsSingleNamed("Subset", subset)[0]}, msg...)
	}
}

// NoSubset asserts that the second parameter is not a subset of the list.
// The order is irrelevant.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNoSubset(t, []int{1, 2, 3}, []int{1, 7})
//	assert.AssertNoSubset(t, []string{"Hello", "World", "Test"}, []string{"Test", "John"})
func NoSubset[T comparable](t testRunner, list []T, subset []T, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if internal.IsSubset(t, list, subset) {
		internal.Fail(t, "The second parameter !!is a subset of the list!!, but should not be.", internal.Objects{internal.NewObjectsSingleNamed("List", list)[0], internal.NewObjectsSingleNamed("Subset", subset)[0]}, msg...)
	}
}

// Unique asserts that the list contains only unique elements.
// The order is irrelevant.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertUnique(t, []int{1, 2, 3})
//	assert.AssertUnique(t, []string{"Hello", "World", "!"})
func Unique[T comparable](t testRunner, list []T, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if !assert.Unique(list) {
		internal.Fail(t, "The list is !!not unique!!.", internal.NewObjectsSingleNamed("List", list), msg...)
	}
}

// NotUnique asserts that the elements in a list are not unique.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNotUnique(t, []int{1, 2, 3, 3})
func NotUnique[elementType comparable](t testRunner, list []elementType, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if assert.Unique(list) {
		internal.Fail(t, "The list !!is unique!!, but should not.", internal.NewObjectsSingleNamed("List", list), msg...)
	}
}

// InRange asserts that the value is in the range.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertInRange(t, 5, 1, 10)
func InRange[T number](t testRunner, value T, min T, max T, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if min >= max {
		internal.Fail(t, "The minimum value is greater than or equal to the maximum value.", internal.Objects{internal.NewObjectsSingleNamed("Min", min)[0], internal.NewObjectsSingleNamed("Max", max)[0]}, msg...)
	}

	if value < min || value > max {
		internal.Fail(t, "The value is !!not in range!!, but should be.", internal.Objects{internal.NewObjectsSingleNamed("Value", value)[0], internal.NewObjectsSingleNamed("Min", min)[0], internal.NewObjectsSingleNamed("Max", max)[0]}, msg...)
	}
}

// NotInRange asserts that the value is not in the range.
//
// When using a custom message, the same formatting as with fmt.Sprintf() is used.
//
// Example:
//
//	assert.AssertNotInRange(t, 5, 1, 10)
func NotInRange[T number](t testRunner, value T, min T, max T, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	if min >= max {
		internal.Fail(t, "The minimum value is greater than or equal to the maximum value.", internal.Objects{internal.NewObjectsSingleNamed("Min", min)[0], internal.NewObjectsSingleNamed("Max", max)[0]}, msg...)
	}

	if value >= min && value <= max {
		internal.Fail(t, "The value is in range, but should not be.", internal.Objects{internal.NewObjectsSingleNamed("Value", value)[0], internal.NewObjectsSingleNamed("Min", min)[0], internal.NewObjectsSingleNamed("Max", max)[0]}, msg...)
	}
}

func FailNow(t testRunner, msg ...any) {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	internal.Fail(t, "The test should fail now.", internal.Objects{}, msg...)
	t.FailNow()
}
