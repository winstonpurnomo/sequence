package sequence

import "fmt"

// Returns an array containing the results of mapping the given closure over the sequence’s elements.
func Map[T ~[]In, In, Out any](s T, fn func(In) Out) []Out {
	res := make([]Out, 0, len(s))
	for _, v := range s {
		res = append(res, fn(v))
	}
	return res
}

// Returns an array containing the results of mapping the given closure over the sequence’s elements.
func TryMap[T ~[]In, In, Out any](s T, fn func(In) (Out, error)) ([]Out, error) {
	res := make([]Out, 0, len(s))
	for _, v := range s {
		as, err := fn(v)
		if err != nil {
			return nil, fmt.Errorf("tryMap(): %w", err)
		}
		res = append(res, as)
	}
	return res, nil
}

// Returns an array containing the results of mapping the given closure over the sequence’s elements.
func CollectMap[T ~[]In, In, Out any](s T, fn func(In) (Out, error)) ([]Out, []error) {
	res := make([]Out, 0, len(s))
	errs := make([]error, 0, len(s))
	for _, v := range s {
		as, err := fn(v)
		if err != nil {
			errs = append(errs, err)
		} else {
			res = append(res, as)
		}
	}
	return res, errs
}

// Returns an array containing the non-nil results of calling the given transformation with each element of this sequence.
func CompactMap[T ~[]In, In, Out any](s T, fn func(In) *Out) []Out {
	res := make([]Out, 0, len(s))
	for _, v := range s {
		if result := fn(v); result != nil {
			res = append(res, *result)
		}
	}
	return res
}

// Returns the result of combining the elements of the sequence using the given function.
func Reduce[T ~[]In, In, Out any](ss T, into Out, reducer func(Out, In) Out) Out {
	for _, s := range ss {
		into = reducer(into, s)
	}
	return into
}
