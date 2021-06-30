package internal

// Fail lets the test fail, with a message.
func Fail(t testingT, args ...interface{}) bool {
	if test, ok := t.(helper); ok {
		test.Helper()
	}

	t.Error(args...)
	return false
}
