package assert

func AssertNotNil(val any) {
	if val == nil {
		panic("Given value was nil")
	}
}

func AssertEq[T any](val1, val2 any) {
	if val1 != val2 {
		panic("val1 != val2")
	}
}
