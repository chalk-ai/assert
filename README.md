<h1 align="center">Assert</h1>

Fork of the very nice [testza](https://github.com/MarvinJWendt/testza/) with some modifications.
This version uses `comparable` from newer versions of golang and drops the `Assert` prefix from the functions.
It also runs `t.FailNow()` for `assert.NoError`, and introduces json comparisons.

<p align="center">

<a href="https://github.com/chalk-ai/assert/releases">
<img src="https://img.shields.io/github/v/release/chalk-ai/assert?style=flat-square" alt="Latest Release">
</a>

<a href="https://pkg.go.dev/github.com/chalk-ai/assert" target="_blank">
<img src="https://pkg.go.dev/badge/github.com/chalk-ai/assert.svg" alt="Go Reference">
</a>

</p>

---

<p align="center">
<strong><a href="https://github.com/chalk-ai/assert#-installation">Get The Module</a></strong>
|
<strong><a href="https://github.com/chalk-ai/assert#-documentation" target="_blank">Documentation</a></strong>
</p>

---
<br/>

<img height="400" alt="Screenshot of an example test message" src="https://user-images.githubusercontent.com/31022056/161153895-e772bc61-b751-407f-b526-8f6a66d8f8d5.png" />

<br/>

## Installation

```console
# Execute this command inside your project
go get github.com/chalk-ai/assert
```
<br/>
<br/>

## Description

assert is a full-featured testing framework for Go.
It integrates with the default test runner, so you can use it with the standard `go test` tool.
assert contains easy-to-use methods, like assertions, output capturing, fuzzing, and much more.

The main goal of `assert` is to provide an easy and fun experience writing tests and providing a nice, user-friendly output.

## Features

| Feature            | Description                                                                                                                                          |
|--------------------|------------------------------------------------------------------------------------------------------------------------------------------------------|
| Assertions         | Assertions allow you to quickly check objects for expected values.                                                                                   |
| Fuzzing            | Fuzzing allows you to check functions against sets of generated input parameters.<br/>A couple lines of test code can run thousands of sanity tests. |
| Output Capture     | Capture and validate output written to the terminal.<br/>Perfect for CLI tools.                                                                      |
| Snapshots          | Snapshot objects between test runs, to ensure a consistent behaviour.                                                                                |
| Clean Output       | Clean and colorful output provides you the needed details in an easy-to-understand format.                                                           |
| System Information | assert prints information about the system on startup.<br/> You can quickly figure out what's wrong, when a user submits an issue.                   |
| Well Documented    | Every function of assert is well documented and contains an example to make usage super easy.                                                        |
| Customizable       | assert features customizable settings, if you want to change something.                                                                              |
| Test flags         | You can configure assert via flags too!<br/>That makes it super simple to change test runs, or output, without touching code!                        |

## Getting Started

See the examples below for a quick introduction!

```go
// --- Some Examples ---

// - Some assertions -
assert.AssertTrue(t, true) // -> Pass
assert.AssertNoError(t, err) // -> Pass
assert.AssertEqual(t, object, object) // -> Pass
// ...

// - Testing console output -
// Test the output of your CLI tool easily!
terminalOutput, _ := assert.CaptureStdout(func(w io.Writer) error {fmt.Println("Hello"); return nil})
asssert.AssertEqual(t, terminalOutput, "Hello\n") // -> Pass

// - Fuzzing -
// Testing a function that accepts email addresses as a parameter:

// Testset of many different email addresses
emailAddresses := assert.FuzzStringEmailAddresses()

// Run a test for every string in the test set
assert.FuzzStringRunTests(t, emailAddresses, func(t *testing.T, index int, str string) {
  user, domain, err := internal.ParseEmailAddress(str) // Use your function
  assert.AssertNoError(t, err) // Assert that your function does not return an error
  assert.AssertNotZero(t, user) // Assert that the user is returned
  assert.AssertNotZero(t, domain) // Assert that the domain is returned
})

// And that's just a few examples of what you can do with assert!
```

## 📚 Documentation

<!-- docs:start -->
<table>
  <tr>
    <th>Module</th>
    <th>Methods</th>
  </tr><tr>
<td><a href="https://github.com/chalk-ai/assert#Settings">Settings</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [SetColorsEnabled](https://github.com/chalk-ai/assert#SetColorsEnabled)
  - [SetDiffContextLines](https://github.com/chalk-ai/assert#SetDiffContextLines)
  - [SetLineNumbersEnabled](https://github.com/chalk-ai/assert#SetLineNumbersEnabled)
  - [SetRandomSeed](https://github.com/chalk-ai/assert#SetRandomSeed)
  - [SetShowStartupMessage](https://github.com/chalk-ai/assert#SetShowStartupMessage)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/chalk-ai/assert#Assert">Assert</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [AssertCompletesIn](https://github.com/chalk-ai/assert#AssertCompletesIn)
  - [AssertContains](https://github.com/chalk-ai/assert#AssertContains)
  - [AssertDecreasing](https://github.com/chalk-ai/assert#AssertDecreasing)
  - [AssertDirEmpty](https://github.com/chalk-ai/assert#AssertDirEmpty)
  - [AssertDirExists](https://github.com/chalk-ai/assert#AssertDirExists)
  - [AssertDirNotEmpty](https://github.com/chalk-ai/assert#AssertDirNotEmpty)
  - [AssertEqual](https://github.com/chalk-ai/assert#AssertEqual)
  - [AssertEqualValues](https://github.com/chalk-ai/assert#AssertEqualValues)
  - [AssertErrorIs](https://github.com/chalk-ai/assert#AssertErrorIs)
  - [AssertFalse](https://github.com/chalk-ai/assert#AssertFalse)
  - [AssertFileExists](https://github.com/chalk-ai/assert#AssertFileExists)
  - [AssertGreater](https://github.com/chalk-ai/assert#AssertGreater)
  - [AssertGreaterOrEqual](https://github.com/chalk-ai/assert#AssertGreaterOrEqual)
  - [AssertImplements](https://github.com/chalk-ai/assert#AssertImplements)
  - [AssertInRange](https://github.com/chalk-ai/assert#AssertInRange)
  - [AssertIncreasing](https://github.com/chalk-ai/assert#AssertIncreasing)
  - [AssertKindOf](https://github.com/chalk-ai/assert#AssertKindOf)
  - [AssertLen](https://github.com/chalk-ai/assert#AssertLen)
  - [AssertLess](https://github.com/chalk-ai/assert#AssertLess)
  - [AssertLessOrEqual](https://github.com/chalk-ai/assert#AssertLessOrEqual)
  - [AssertNil](https://github.com/chalk-ai/assert#AssertNil)
  - [AssertNoDirExists](https://github.com/chalk-ai/assert#AssertNoDirExists)
  - [AssertNoError](https://github.com/chalk-ai/assert#AssertNoError)
  - [AssertNoFileExists](https://github.com/chalk-ai/assert#AssertNoFileExists)
  - [AssertNoSubset](https://github.com/chalk-ai/assert#AssertNoSubset)
  - [AssertNotCompletesIn](https://github.com/chalk-ai/assert#AssertNotCompletesIn)
  - [AssertNotContains](https://github.com/chalk-ai/assert#AssertNotContains)
  - [AssertNotEqual](https://github.com/chalk-ai/assert#AssertNotEqual)
  - [AssertNotEqualValues](https://github.com/chalk-ai/assert#AssertNotEqualValues)
  - [AssertNotErrorIs](https://github.com/chalk-ai/assert#AssertNotErrorIs)
  - [AssertNotImplements](https://github.com/chalk-ai/assert#AssertNotImplements)
  - [AssertNotInRange](https://github.com/chalk-ai/assert#AssertNotInRange)
  - [AssertNotKindOf](https://github.com/chalk-ai/assert#AssertNotKindOf)
  - [AssertNotNil](https://github.com/chalk-ai/assert#AssertNotNil)
  - [AssertNotNumeric](https://github.com/chalk-ai/assert#AssertNotNumeric)
  - [AssertNotPanics](https://github.com/chalk-ai/assert#AssertNotPanics)
  - [AssertNotRegexp](https://github.com/chalk-ai/assert#AssertNotRegexp)
  - [AssertNotSameElements](https://github.com/chalk-ai/assert#AssertNotSameElements)
  - [AssertNotUnique](https://github.com/chalk-ai/assert#AssertNotUnique)
  - [AssertNotZero](https://github.com/chalk-ai/assert#AssertNotZero)
  - [AssertNumeric](https://github.com/chalk-ai/assert#AssertNumeric)
  - [AssertPanics](https://github.com/chalk-ai/assert#AssertPanics)
  - [AssertRegexp](https://github.com/chalk-ai/assert#AssertRegexp)
  - [AssertSameElements](https://github.com/chalk-ai/assert#AssertSameElements)
  - [AssertSubset](https://github.com/chalk-ai/assert#AssertSubset)
  - [AssertTestFails](https://github.com/chalk-ai/assert#AssertTestFails)
  - [AssertTrue](https://github.com/chalk-ai/assert#AssertTrue)
  - [AssertUnique](https://github.com/chalk-ai/assert#AssertUnique)
  - [AssertZero](https://github.com/chalk-ai/assert#AssertZero)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/chalk-ai/assert#Capture">Capture</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [CaptureStderr](https://github.com/chalk-ai/assert#CaptureStderr)
  - [CaptureStdout](https://github.com/chalk-ai/assert#CaptureStdout)
  - [CaptureStdoutAndStderr](https://github.com/chalk-ai/assert#CaptureStdoutAndStderr)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/chalk-ai/assert#Fuzz-Utils">Fuzz Utils</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [FuzzUtilDistinctSet](https://github.com/chalk-ai/assert#FuzzUtilDistinctSet)
  - [FuzzUtilLimitSet](https://github.com/chalk-ai/assert#FuzzUtilLimitSet)
  - [FuzzUtilMergeSets](https://github.com/chalk-ai/assert#FuzzUtilMergeSets)
  - [FuzzUtilModifySet](https://github.com/chalk-ai/assert#FuzzUtilModifySet)
  - [FuzzUtilRunTests](https://github.com/chalk-ai/assert#FuzzUtilRunTests)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/chalk-ai/assert#Fuzz-Booleans">Fuzz Booleans</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [FuzzBoolFull](https://github.com/chalk-ai/assert#FuzzBoolFull)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/chalk-ai/assert#Fuzz-Strings">Fuzz Strings</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [FuzzStringEmailAddresses](https://github.com/chalk-ai/assert#FuzzStringEmailAddresses)
  - [FuzzStringEmpty](https://github.com/chalk-ai/assert#FuzzStringEmpty)
  - [FuzzStringFull](https://github.com/chalk-ai/assert#FuzzStringFull)
  - [FuzzStringGenerateRandom](https://github.com/chalk-ai/assert#FuzzStringGenerateRandom)
  - [FuzzStringHtmlTags](https://github.com/chalk-ai/assert#FuzzStringHtmlTags)
  - [FuzzStringLong](https://github.com/chalk-ai/assert#FuzzStringLong)
  - [FuzzStringNumeric](https://github.com/chalk-ai/assert#FuzzStringNumeric)
  - [FuzzStringUsernames](https://github.com/chalk-ai/assert#FuzzStringUsernames)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/chalk-ai/assert#Fuzz-Float64s">Fuzz Float64s</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [FuzzFloat64Full](https://github.com/chalk-ai/assert#FuzzFloat64Full)
  - [FuzzFloat64GenerateRandomNegative](https://github.com/chalk-ai/assert#FuzzFloat64GenerateRandomNegative)
  - [FuzzFloat64GenerateRandomPositive](https://github.com/chalk-ai/assert#FuzzFloat64GenerateRandomPositive)
  - [FuzzFloat64GenerateRandomRange](https://github.com/chalk-ai/assert#FuzzFloat64GenerateRandomRange)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/chalk-ai/assert#Fuzz-Integers">Fuzz Integers</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [FuzzIntFull](https://github.com/chalk-ai/assert#FuzzIntFull)
  - [FuzzIntGenerateRandomNegative](https://github.com/chalk-ai/assert#FuzzIntGenerateRandomNegative)
  - [FuzzIntGenerateRandomPositive](https://github.com/chalk-ai/assert#FuzzIntGenerateRandomPositive)
  - [FuzzIntGenerateRandomRange](https://github.com/chalk-ai/assert#FuzzIntGenerateRandomRange)
</td>

</details>

</tr>
<tr>
<td><a href="https://github.com/chalk-ai/assert#Snapshot">Snapshot</a></td>
<td>

<details>
<summary>Click to expand</summary>

  - [SnapshotCreate](https://github.com/chalk-ai/assert#SnapshotCreate)
  - [SnapshotCreateOrValidate](https://github.com/chalk-ai/assert#SnapshotCreateOrValidate)
  - [SnapshotValidate](https://github.com/chalk-ai/assert#SnapshotValidate)
</td>

</details>

</tr>
</table>

### Assert

#### AssertCompletesIn

```go
func AssertCompletesIn(t testRunner, duration time.Duration, f func(), msg ...any)
```

AssertCompletesIn asserts that a function completes in a given time.
Use this function to test that functions do not take too long to complete.

NOTE: Every system takes a different amount of time to complete a function.
Do not set the duration too low, if you want consistent results.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertCompletesIn(t, 2 * time.Second, func() {
    	// some code that should take less than 2 seconds...
    }) // => PASS

#### AssertContains

```go
func AssertContains(t testRunner, object, element any, msg ...any)
```

AssertContains asserts that a string/list/array/slice/map contains the
specified element.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertContains(t, []int{1,2,3}, 2)
    assert.AssertContains(t, []string{"Hello", "World"}, "World")
    assert.AssertContains(t, "Hello, World!", "World")

#### AssertDecreasing

```go
func AssertDecreasing(t testRunner, object any, msg ...any)
```

AssertDecreasing asserts that the values in a slice are decreasing. the test
fails if the values are not in a slice or if the values are not comparable.

Valid input kinds are: []int, []int8, []int16, []int32, []int64, []uint,
[]uint8, []uint16, []uint32, []uint64, []float32, []float64.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertDecreasing(t, []int{1000, 137, 2, 1})
    assert.AssertDecreasing(t, []float32{13.5, 7, 0.1, -10.3})

#### AssertDirEmpty

```go
func AssertDirEmpty(t testRunner, dir string, msg ...any)
```

AssertDirEmpty asserts that a directory is empty. The test will pass when
the directory is empty or does not exist.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertDirEmpty(t, "FolderName")

#### AssertDirExists

```go
func AssertDirExists(t testRunner, dir string, msg ...any)
```

AssertDirExists asserts that a directory exists. The test will pass when the
directory exists, and it's visible to the current user. The test will fail,
if the path points to a file.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertDirExists(t, "FolderName")

#### AssertDirNotEmpty

```go
func AssertDirNotEmpty(t testRunner, dir string, msg ...any)
```

AssertDirNotEmpty asserts that a directory is not empty The test will pass
when the directory is not empty and will fail if the directory does not
exist.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertDirNotEmpty(t, "FolderName")

#### AssertEqual

```go
func AssertEqual(t testRunner, expected any, actual any, msg ...any)
```

AssertEqual asserts that two objects are equal.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertEqual(t, "Hello, World!", "Hello, World!")
    assert.AssertEqual(t, true, true)

#### AssertEqualValues

```go
func AssertEqualValues(t testRunner, expected any, actual any, msg ...any)
```

AssertEqualValues asserts that two objects have equal values. The order of
the values is also validated.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertEqualValues(t, []string{"Hello", "World"}, []string{"Hello", "World"})
    assert.AssertEqualValues(t, []int{1,2}, []int{1,2})
    assert.AssertEqualValues(t, []int{1,2}, []int{2,1}) // FAILS (wrong order)

Comparing struct values:

    person1 := Person{
      Name:   "Marvin Wendt",
      Age:    20,
      Gender: "male",
    }

    person2 := Person{
      Name:   "Marvin Wendt",
      Age:    20,
      Gender: "male",
    }

    assert.AssertEqualValues(t, person1, person2)

#### AssertErrorIs

```go
func AssertErrorIs(t testRunner, err, target error, msg ...any)
```

AssertErrorIs asserts that target is inside the error chain of err.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    var testErr = errors.New("hello world")
    var testErrWrapped = fmt.Errorf("test err: %w", testErr)
    assert.AssertErrorIs(t, testErrWrapped ,testErr)

#### AssertFalse

```go
func AssertFalse(t testRunner, value any, msg ...any)
```

AssertFalse asserts that an expression or object resolves to false.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertFalse(t, false)
    assert.AssertFalse(t, 1 == 2)
    assert.AssertFalse(t, 2 != 2)
    assert.AssertFalse(t, 1 > 5 && 4 < 0)

#### AssertFileExists

```go
func AssertFileExists(t testRunner, file string, msg ...any)
```

AssertFileExists asserts that a file exists.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertFileExists(t, "./test.txt")
    assert.AssertFileExists(t, "./config.yaml", "the config file is missing")

#### AssertGreater

```go
func AssertGreater(t testRunner, object1, object2 any, msg ...any)
```

AssertGreater asserts that the first object is greater than the second.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertGreater(t, 5, 1)
    assert.AssertGreater(t, 10, -10)

#### AssertGreaterOrEqual

```go
func AssertGreaterOrEqual(t testRunner, object1, object2 interface{}, msg ...interface{})
```

AssertGreaterOrEqual asserts that the first object is greater than or equal
to the second.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertGreaterOrEqual(t, 5, 1)
    assert.AssertGreaterOrEqual(t, 10, -10)

assert.AssertGreaterOrEqual(t, 10, 10)

#### AssertImplements

```go
func AssertImplements(t testRunner, interfaceObject, object any, msg ...any)
```

AssertImplements asserts that an objects implements an interface.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertImplements(t, (*YourInterface)(nil), new(YourObject))
    assert.AssertImplements(t, (*fmt.Stringer)(nil), new(types.Const)) => pass

#### AssertInRange

```go
func AssertInRange[T number](t testRunner, value T, min T, max T, msg ...any)
```

AssertInRange asserts that the value is in the range.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertInRange(t, 5, 1, 10)

#### AssertIncreasing

```go
func AssertIncreasing(t testRunner, object any, msg ...any)
```

AssertIncreasing asserts that the values in a slice are increasing. the test
fails if the values are not in a slice or if the values are not comparable.

Valid input kinds are: []int, []int8, []int16, []int32, []int64, []uint,
[]uint8, []uint16, []uint32, []uint64, []float32, []float64.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertIncreasing(t, []int{1, 2, 137, 1000})
    assert.AssertIncreasing(t, []float32{-10.3, 0.1, 7, 13.5})

#### AssertKindOf

```go
func AssertKindOf(t testRunner, expectedKind reflect.Kind, object any, msg ...any)
```

AssertKindOf asserts that the object is a type of kind exptectedKind.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertKindOf(t, reflect.Slice, []int{1,2,3})
    assert.AssertKindOf(t, reflect.Slice, []string{"Hello", "World"})
    assert.AssertKindOf(t, reflect.Int, 1337)
    assert.AssertKindOf(t, reflect.Bool, true)
    assert.AssertKindOf(t, reflect.Map, map[string]bool{})

#### AssertLen

```go
func AssertLen(t testRunner, object any, length int, msg ...any)
```

AssertLen asserts that the length of an object is equal to the given length.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertLen(t, "abc", 3)
    assert.AssertLen(t, "Assert", 6)
    assert.AssertLen(t, []int{1, 2, 1337, 25}, 4)
    assert.AssertLen(t, map[string]int{"asd": 1, "test": 1337}, 2)

#### AssertLess

```go
func AssertLess(t testRunner, object1, object2 any, msg ...any)
```

AssertLess asserts that the first object is less than the second.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertLess(t, 1, 5)
    assert.AssertLess(t, -10, 10)

#### AssertLessOrEqual

```go
func AssertLessOrEqual(t testRunner, object1, object2 interface{}, msg ...interface{})
```

AssertLessOrEqual asserts that the first object is less than or equal to the
second.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertLessOrEqual(t, 1, 5)
    assert.AssertLessOrEqual(t, -10, 10)
    assert.AssertLessOrEqual(t, 1, 1)

#### AssertNil

```go
func AssertNil(t testRunner, object any, msg ...any)
```

AssertNil asserts that an object is nil.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNil(t, nil)

#### AssertNoDirExists

```go
func AssertNoDirExists(t testRunner, dir string, msg ...any)
```

AssertNoDirExists asserts that a directory does not exists. The test will
pass, if the path points to a file, as a directory with the same name,
cannot exist.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNoDirExists(t, "FolderName")

#### AssertNoError

```go
func AssertNoError(t testRunner, err error, msg ...any)
```

AssertNoError asserts that an error is nil.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    err := nil
    assert.AssertNoError(t, err)

#### AssertNoFileExists

```go
func AssertNoFileExists(t testRunner, file string, msg ...any)
```



#### AssertNoSubset

```go
func AssertNoSubset(t testRunner, list any, subset any, msg ...any)
```

AssertNoSubset asserts that the second parameter is not a subset of the
list. The order is irrelevant.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNoSubset(t, []int{1, 2, 3}, []int{1, 7})
    assert.AssertNoSubset(t, []string{"Hello", "World", "Test"}, []string{"Test", "John"})

#### AssertNotCompletesIn

```go
func AssertNotCompletesIn(t testRunner, duration time.Duration, f func(), msg ...any)
```

AssertNotCompletesIn asserts that a function does not complete in a given
time. Use this function to test that functions do not complete to quickly.
For example if your database connection completes in under a millisecond,
there might be something wrong.

NOTE: Every system takes a different amount of time to complete a function.
Do not set the duration too high, if you want consistent results.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNotCompletesIn(t, 2 * time.Second, func() {
    	// some code that should take more than 2 seconds...
    	time.Sleep(3 * time.Second)
    }) // => PASS

#### AssertNotContains

```go
func AssertNotContains(t testRunner, object, element any, msg ...any)
```

AssertNotContains asserts that a string/list/array/slice/map does not
contain the specified element.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNotContains(t, []string{"Hello", "World"}, "Spaceship")
    assert.AssertNotContains(t, "Hello, World!", "Spaceship")

#### AssertNotEqual

```go
func AssertNotEqual(t testRunner, expected any, actual any, msg ...any)
```

AssertNotEqual asserts that two objects are not equal.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNotEqual(t, true, false)
    assert.AssertNotEqual(t, "Hello", "World")

#### AssertNotEqualValues

```go
func AssertNotEqualValues(t testRunner, expected any, actual any, msg ...any)
```

AssertNotEqualValues asserts that two objects do not have equal values.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNotEqualValues(t, []int{1,2}, []int{3,4})

Comparing struct values:

    person1 := Person{
      Name:   "Marvin Wendt",
      Age:    20,
      Gender: "male",
    }

    person2 := Person{
      Name:   "Marvin Wendt",
      Age:    20,
      Gender: "female", // <-- CHANGED
    }

    assert.AssertNotEqualValues(t, person1, person2)

#### AssertNotErrorIs

```go
func AssertNotErrorIs(t testRunner, err, target error, msg ...any)
```

AssertNotErrorIs

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    var testErr = errors.New("hello world")
    var test2Err = errors.New("hello world 2")
    var testErrWrapped = fmt.Errorf("test err: %w", testErr)
    assert.AssertNotErrorIs(t, testErrWrapped, test2Err)

#### AssertNotImplements

```go
func AssertNotImplements(t testRunner, interfaceObject, object any, msg ...any)
```

AssertNotImplements asserts that an object does not implement an interface.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNotImplements(t, (*YourInterface)(nil), new(YourObject))
    assert.AssertNotImplements(t, (*fmt.Stringer)(nil), new(types.Const)) => fail, because types.Const does implement fmt.Stringer.

#### AssertNotInRange

```go
func AssertNotInRange[T number](t testRunner, value T, min T, max T, msg ...any)
```

AssertNotInRange asserts that the value is not in the range.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNotInRange(t, 5, 1, 10)

#### AssertNotKindOf

```go
func AssertNotKindOf(t testRunner, kind reflect.Kind, object any, msg ...any)
```

AssertNotKindOf asserts that the object is not a type of kind `kind`.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNotKindOf(t, reflect.Slice, "Hello, World")
    assert.AssertNotKindOf(t, reflect.Slice, true)
    assert.AssertNotKindOf(t, reflect.Int, 13.37)
    assert.AssertNotKindOf(t, reflect.Bool, map[string]bool{})
    assert.AssertNotKindOf(t, reflect.Map, false)

#### AssertNotNil

```go
func AssertNotNil(t testRunner, object any, msg ...any)
```

AssertNotNil asserts that an object is not nil.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNotNil(t, true)
    assert.AssertNotNil(t, "Hello, World!")
    assert.AssertNotNil(t, 0)

#### AssertNotNumeric

```go
func AssertNotNumeric(t testRunner, object any, msg ...any)
```

AssertNotNumeric checks if the object is not a numeric type. Numeric types
are: Int, Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16,
Uint32, Uint64, Complex64 and Complex128.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNotNumeric(t, true)
    assert.AssertNotNumeric(t, "123")

#### AssertNotPanics

```go
func AssertNotPanics(t testRunner, f func(), msg ...any)
```

AssertNotPanics asserts that a function does not panic.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNotPanics(t, func() {
    	// some code that does not call a panic...
    }) // => PASS

#### AssertNotRegexp

```go
func AssertNotRegexp(t testRunner, regex any, txt any, msg ...any)
```

AssertNotRegexp asserts that a string does not match a given regexp.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNotRegexp(t, "ab.*", "Hello, World!")

#### AssertNotSameElements

```go
func AssertNotSameElements(t testRunner, expected any, actual any, msg ...any)
```

AssertNotSameElements asserts that two slices contains same elements
(including pointers). The order is irrelevant.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

     assert.AssertNotSameElements(t, []string{"Hello", "World"}, []string{"Hello", "World", "World"})
     assert.AssertNotSameElements(t, []int{1,2}, []int{1,2,3})

     type A struct {
    	  a string
     }
     assert.AssertNotSameElements(t, []*A{{a: "A"}, {a: "B"}, {a: "C"}}, []*A{{a: "A"}, {a: "B"}, {a: "C"}, {a: "D"}})

#### AssertNotUnique

```go
func AssertNotUnique[elementType comparable](t testRunner, list []elementType, msg ...any)
```

AssertNotUnique asserts that the elements in a list are not unique.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNotUnique(t, []int{1, 2, 3, 3})

#### AssertNotZero

```go
func AssertNotZero(t testRunner, value any, msg ...any)
```

AssertNotZero asserts that the value is not the zero value for it's type.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNotZero(t, 1337)
    assert.AssertNotZero(t, true)
    assert.AssertNotZero(t, "Hello, World")

#### AssertNumeric

```go
func AssertNumeric(t testRunner, object any, msg ...any)
```

AssertNumeric asserts that the object is a numeric type. Numeric types are:
Int, Int8, Int16, Int32, Int64, Float32, Float64, Uint, Uint8, Uint16,
Uint32, Uint64, Complex64 and Complex128.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertNumeric(t, 123)
    assert.AssertNumeric(t, 1.23)
    assert.AssertNumeric(t, uint(123))

#### AssertPanics

```go
func AssertPanics(t testRunner, f func(), msg ...any)
```

AssertPanics asserts that a function panics.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertPanics(t, func() {
    	// ...
    	panic("some panic")
    }) // => PASS

#### AssertRegexp

```go
func AssertRegexp(t testRunner, regex any, txt any, msg ...any)
```

AssertRegexp asserts that a string matches a given regexp.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertRegexp(t, "^a.*c$", "abc")

#### AssertSameElements

```go
func AssertSameElements(t testRunner, expected any, actual any, msg ...any)
```

AssertSameElements asserts that two slices contains same elements (including
pointers). The order is irrelevant.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

     assert.AssertSameElements(t, []string{"Hello", "World"}, []string{"Hello", "World"})
     assert.AssertSameElements(t, []int{1,2,3}, []int{1,2,3})
     assert.AssertSameElements(t, []int{1,2}, []int{2,1})

     type A struct {
    	  a string
     }
     assert.AssertSameElements(t, []*A{{a: "A"}, {a: "B"}, {a: "C"}}, []*A{{a: "A"}, {a: "B"}, {a: "C"}})

#### AssertSubset

```go
func AssertSubset(t testRunner, list any, subset any, msg ...any)
```

AssertSubset asserts that the second parameter is a subset of the list.
The order is irrelevant.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertSubset(t, []int{1, 2, 3}, []int{1, 2})
    assert.AssertSubset(t, []string{"Hello", "World", "Test"}, []string{"Test", "World"})

#### AssertTestFails

```go
func AssertTestFails(t testRunner, test func(t TestingPackageWithFailFunctions), msg ...any)
```

AssertTestFails asserts that a unit test fails. A unit test fails if one of
the following methods is called in the test function: Error, Errorf, Fail,
FailNow, Fatal, Fatalf

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertTestFails(t, func(t assert.TestingPackageWithFailFunctions) {
    	assert.AssertTrue(t, false)
    }) // => Pass

    assert.AssertTestFails(t, func(t assert.TestingPackageWithFailFunctions) {
    	// ...
    	t.Fail() // Or any other failing method.
    }) // => Pass

#### AssertTrue

```go
func AssertTrue(t testRunner, value any, msg ...any)
```

AssertTrue asserts that an expression or object resolves to true.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertTrue(t, true)
    assert.AssertTrue(t, 1 == 1)
    assert.AssertTrue(t, 2 != 3)
    assert.AssertTrue(t, 1 > 0 && 4 < 5)

#### AssertUnique

```go
func AssertUnique[elementType comparable](t testRunner, list []elementType, msg ...any)
```

AssertUnique asserts that the list contains only unique elements. The order
is irrelevant.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertUnique(t, []int{1, 2, 3})
    assert.AssertUnique(t, []string{"Hello", "World", "!"})

#### AssertZero

```go
func AssertZero(t testRunner, value any, msg ...any)
```

AssertZero asserts that the value is the zero value for it's type.

When using a custom message, the same formatting as with fmt.Sprintf() is
used.

Example:

    assert.AssertZero(t, 0)
    assert.AssertZero(t, false)
    assert.AssertZero(t, "")

### Capture

#### CaptureStderr

```go
func CaptureStderr(capture func(w io.Writer) error) (string, error)
```

CaptureStderr captures everything written to stderr from a specific
function. You can use this method in tests, to validate that your functions
writes a string to the terminal.

Example:

    stderr, err := assert.CaptureStderr(func(w io.Writer) error {
    	_, err := fmt.Fprint(os.Stderr, "Hello, World!")
    	assert.AssertNoError(t, err)
    	return nil
    })

    assert.AssertNoError(t, err)
    assert.AssertEqual(t, "Hello, World!", stderr)

#### CaptureStdout

```go
func CaptureStdout(capture func(w io.Writer) error) (string, error)
```

CaptureStdout captures everything written to stdout from a specific
function. You can use this method in tests, to validate that your functions
writes a string to the terminal.

Example:

    stdout, err := assert.CaptureStdout(func(w io.Writer) error {
    	fmt.Println("Hello, World!")
    	return nil
    })

    assert.AssertNoError(t, err)
    assert.AssertEqual(t, "Hello, World!", stdout)

#### CaptureStdoutAndStderr

```go
func CaptureStdoutAndStderr(capture func(stdoutWriter, stderrWriter io.Writer) error) (stdout, stderr string, err error)
```

CaptureStdoutAndStderr captures everything written to stdout and stderr from
a specific function. You can use this method in tests, to validate that your
functions writes a string to the terminal.

Example:

    stdout, stderr, err := assert.CaptureStdoutAndStderr(func(stdoutWriter, stderrWriter io.Writer) error {
    	fmt.Fprint(os.Stdout, "Hello")
    	fmt.Fprint(os.Stderr, "World")
    	return nil
    })

    assert.AssertNoError(t, err)
    assert.AssertEqual(t, "Hello", stdout)
    assert.AssertEqual(t, "World", stderr)

### Fuzz Booleans

#### FuzzBoolFull

```go
func FuzzBoolFull() []bool
```

FuzzBoolFull returns true and false in a boolean slice.

### Fuzz Float64s

#### FuzzFloat64Full

```go
func FuzzFloat64Full() (floats []float64)
```

FuzzFloat64Full returns a combination of every float64 testset and some
random float64s (positive and negative).

#### FuzzFloat64GenerateRandomNegative

```go
func FuzzFloat64GenerateRandomNegative(count int, min float64) (floats []float64)
```

FuzzFloat64GenerateRandomNegative generates random negative integers with
a minimum of min. If the minimum is positive, it will be converted to a
negative number. If it is set to 0, there is no limit.

#### FuzzFloat64GenerateRandomPositive

```go
func FuzzFloat64GenerateRandomPositive(count int, max float64) (floats []float64)
```

FuzzFloat64GenerateRandomPositive generates random positive integers with
a maximum of max. If the maximum is 0, or below, the maximum will be set to
math.MaxInt64.

#### FuzzFloat64GenerateRandomRange

```go
func FuzzFloat64GenerateRandomRange(count int, min, max float64) (floats []float64)
```

FuzzFloat64GenerateRandomRange generates random positive integers with a
maximum of max. If the maximum is 0, or below, the maximum will be set to
math.MaxInt64.

### Fuzz Integers

#### FuzzIntFull

```go
func FuzzIntFull() (ints []int)
```

FuzzIntFull returns a combination of every integer testset and some random
integers (positive and negative).

#### FuzzIntGenerateRandomNegative

```go
func FuzzIntGenerateRandomNegative(count, min int) (ints []int)
```

FuzzIntGenerateRandomNegative generates random negative integers with a
minimum of min. If the minimum is 0, or above, the maximum will be set to
math.MinInt64.

#### FuzzIntGenerateRandomPositive

```go
func FuzzIntGenerateRandomPositive(count, max int) (ints []int)
```

FuzzIntGenerateRandomPositive generates random positive integers with a
maximum of max. If the maximum is 0, or below, the maximum will be set to
math.MaxInt64.

#### FuzzIntGenerateRandomRange

```go
func FuzzIntGenerateRandomRange(count, min, max int) (ints []int)
```

FuzzIntGenerateRandomRange generates random integers with a range of min to
max.

### Fuzz Strings

#### FuzzStringEmailAddresses

```go
func FuzzStringEmailAddresses() []string
```

FuzzStringEmailAddresses returns a test set with valid email addresses.
The addresses may look like they are invalid, but they are all conform to
RFC 2822 and could be used. You can use this test set to test your email
validation process.

#### FuzzStringEmpty

```go
func FuzzStringEmpty() []string
```

FuzzStringEmpty returns a test set with a single empty string.

#### FuzzStringFull

```go
func FuzzStringFull() (ret []string)
```

FuzzStringFull contains all string test sets plus ten generated random
strings. This test set is huge and should only be used if you want to make
sure that no string, at all, can crash a process.

#### FuzzStringGenerateRandom

```go
func FuzzStringGenerateRandom(count, length int) (result []string)
```

FuzzStringGenerateRandom returns random strings in a test set.

#### FuzzStringHtmlTags

```go
func FuzzStringHtmlTags() []string
```

FuzzStringHtmlTags returns a test set with different html tags.

Example:
  - <script>
  - <script>alert('XSS')</script>
  - <a href="https://github.com/chalk-ai/assert">link</a>

#### FuzzStringLong

```go
func FuzzStringLong() (testSet []string)
```

FuzzStringLong returns a test set with long random strings. Returns: [0]:
Random string (length: 25) [1]: Random string (length: 50) [2]: Random
string (length: 100) [3]: Random string (length: 1,000) [4]: Random string
(length: 100,000)

#### FuzzStringNumeric

```go
func FuzzStringNumeric() []string
```

FuzzStringNumeric returns a test set with strings that are numeric.
The highest number in here is "9223372036854775807", which is equal to the
maxmim int64.

#### FuzzStringUsernames

```go
func FuzzStringUsernames() []string
```

FuzzStringUsernames returns a test set with usernames.

### Fuzz Utils

#### FuzzUtilDistinctSet

```go
func FuzzUtilDistinctSet[setType comparable](testSet []setType) []setType
```

FuzzUtilDistinctSet returns a set with removed duplicates.

Example:

    uniqueSet := assert.FuzzUtilDistinctSet([]string{"A", "C", "A", "B", "A", "B", "C"})
    // uniqueSet => []string{"A", "C", "B"}

#### FuzzUtilLimitSet

```go
func FuzzUtilLimitSet[setType any](testSet []setType, max int) []setType
```

FuzzUtilLimitSet returns a random sample of a test set with a maximal size.

Example:

    limitedSet := assert.FuzzUtilLimitSet(assert.FuzzStringFull(), 10)

#### FuzzUtilMergeSets

```go
func FuzzUtilMergeSets[setType any](sets ...[]setType) (merged []setType)
```

FuzzUtilMergeSets merges multiple test sets into one. All test sets must
have the same type.

Example:

    mergedSet := assert.FuzzUtilMergeSets(assert.FuzzIntGenerateRandomNegative(3, 0), assert.FuzzIntGenerateRandomPositive(2, 0))

#### FuzzUtilModifySet

```go
func FuzzUtilModifySet[setType any](inputSet []setType, modifier func(index int, value setType) setType) (floats []setType)
```

FuzzUtilModifySet returns a modified version of a test set.

Example:

     modifiedSet := assert.FuzzUtilModifySet(assert.FuzzIntFull(), func(i int, value int) int {
    		return i * 2 // double every value in the test set
    	})

#### FuzzUtilRunTests

```go
func FuzzUtilRunTests[setType any](t testRunner, testSet []setType, testFunc func(t *testing.T, index int, f setType))
```

FuzzUtilRunTests runs a test for every value in a test set. You can use
the value as input parameter for your functions, to sanity test against
many different cases. This ensures that your functions have a correct error
handling and enables you to test against hundreds of cases easily.

Example:

    assert.FuzzUtilRunTests(t, assert.FuzzStringEmailAddresses(), func(t *testing.T, index int, emailAddress string) {
    	// Test logic
    	// err := YourFunction(emailAddress)
    	// assert.AssertNoError(t, err)
    	// ...
    })

### 

#### GetColorsEnabled

```go
func GetColorsEnabled() bool
```

GetColorsEnabled returns current value of ColorsEnabled setting.
ColorsEnabled controls if assert should print colored output.

#### GetDiffContextLines

```go
func GetDiffContextLines() int
```

GetDiffContextLines returns current value of DiffContextLines setting.
DiffContextLines setting controls how many lines are shown around a changed
diff line. If set to -1 it will show full diff.

#### GetLineNumbersEnabled

```go
func GetLineNumbersEnabled() bool
```

GetLineNumbersEnabled returns current value of LineNumbersEnabled setting.
LineNumbersEnabled controls if line numbers should be printed in failing
tests.

#### GetRandomSeed

```go
func GetRandomSeed() int64
```

GetRandomSeed returns current value of the random seed setting.

#### GetShowStartupMessage

```go
func GetShowStartupMessage() bool
```

GetShowStartupMessage returns current value of showStartupMessage setting.
showStartupMessage setting controls if the startup message should be
printed.

### Settings

#### SetColorsEnabled

```go
func SetColorsEnabled(enabled bool)
```

SetColorsEnabled controls if assert should print colored output. You should
use this in the init() method of the package, which contains your tests.

> This setting can also be set by the command line flag
--assert.disable-color.

Example:

    init() {
      assert.SetColorsEnabled(false) // Disable colored output
      assert.SetColorsEnabled(true)  // Enable colored output
    }

#### SetDiffContextLines

```go
func SetDiffContextLines(lines int)
```

SetDiffContextLines controls how many lines are shown around a changed diff
line. If set to -1 it will show full diff. You should use this in the init()
method of the package, which contains your tests.

> This setting can also be set by the command line flag
--assert.diff-context-lines.

Example:

    init() {
      assert.SetDiffContextLines(-1) // Show all diff lines
      assert.SetDiffContextLines(3)  // Show 3 lines around every changed line
    }

#### SetLineNumbersEnabled

```go
func SetLineNumbersEnabled(enabled bool)
```

SetLineNumbersEnabled controls if line numbers should be printed in
failing tests. You should use this in the init() method of the package,
which contains your tests.

> This setting can also be set by the command line flag
--assert.disable-line-numbers.

Example:

    init() {
      assert.SetLineNumbersEnabled(false) // Disable line numbers
      assert.SetLineNumbersEnabled(true)  // Enable line numbers
    }

#### SetRandomSeed

```go
func SetRandomSeed(seed int64)
```

SetRandomSeed sets the seed for the random generator used in assert.
Using the same seed will result in the same random sequences each time and
guarantee a reproducible test run. Use this setting, if you want a 100%
deterministic test. You should use this in the init() method of the package,
which contains your tests.

> This setting can also be set by the command line flag --assert.seed.

Example:

    init() {
      assert.SetRandomSeed(1337) // Set the seed to 1337
      assert.SetRandomSeed(time.Now().UnixNano()) // Set the seed back to the current time (default | non-deterministic)
    }

#### SetShowStartupMessage

```go
func SetShowStartupMessage(show bool)
```

SetShowStartupMessage controls if the startup message should be printed.
You should use this in the init() method of the package, which contains your
tests.

> This setting can also be set by the command line flag
--assert.disable-startup-message.

Example:

    init() {
      assert.SetShowStartupMessage(false) // Disable the startup message
      assert.SetShowStartupMessage(true)  // Enable the startup message
    }

### Snapshot

#### SnapshotCreate

```go
func SnapshotCreate(name string, snapshotObject any) error
```

SnapshotCreate creates a snapshot of an object, which can be validated
in future test runs. Using this function directly will override
previous snapshots with the same name. You most likely want to use
SnapshotCreateOrValidate.

NOTICE: \r\n will be replaced with \n to make the files consistent between
operating systems.

Example:

    assert.SnapshotCreate(t.Name(), objectToBeSnapshotted)

#### SnapshotCreateOrValidate

```go
func SnapshotCreateOrValidate(t testRunner, name string, object any, msg ...any) error
```

SnapshotCreateOrValidate creates a snapshot of an object which can be used
in future test runs. It is good practice to name your snapshots the same
as the test they are created in. You can do that automatically by using
t.Name() as the second parameter, if you are using the inbuilt test system
of Go. If a snapshot already exists, the function will not create a new one,
but validate the exisiting one. To re-create a snapshot, you can delete the
according file in /testdata/snapshots/.

NOTICE: \r\n will be replaced with \n to make the files consistent between
operating systems.

Example:

    assert.SnapshotCreateOrValidate(t, t.Name(), object)
    assert.SnapshotCreateOrValidate(t, t.Name(), object, "Optional Message")

#### SnapshotValidate

```go
func SnapshotValidate(t testRunner, name string, actual any, msg ...any) error
```

SnapshotValidate validates an already exisiting snapshot of an object.
You most likely want to use SnapshotCreateOrValidate.

NOTICE: \r\n will be replaced with \n to make the files consistent between
operating systems.

Example:

    assert.SnapshotValidate(t, t.Name(), objectToBeValidated)
    assert.SnapshotValidate(t, t.Name(), objectToBeValidated, "Optional message")


<!-- docs:end -->

---

> Made with ❤️ by [@MarvinJWendt](https://github.com/MarvinJWendt) and contributors! |
> [MarvinJWendt.com](https://marvinjwendt.com)
