package sequence

// Returns a new slice of the same type containing, in order, the elements of the original collection that satisfy the given test.
func Filter[T ~[]In, In any](ss T, test func(In) bool) (ret T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

// Returns the first element of the sequence that satisfies the given test. Functionally equivalent to calling Filter(ss, func(T) bool)[0].
func First[T ~[]In, In any](ss T, test func(In) bool) (ret In) {
	for _, s := range ss {
		if test(s) {
			return s
		}
	}

	// Return zero value of T if not found.
	var zero In
	return zero
}
