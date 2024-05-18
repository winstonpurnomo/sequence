# sequence

This Go package is designed to emulate the Swift standard library's sequence operations. It provides a set of generic functions to operate on slices, emulating the behavior of Swift's sequence operations.

```go
aList := []string{"Apple", "Banana", "Coconut"}
itemsWithA := sequence.Filter(aList, func(s string) bool { return strings.Contains(s, "a") }) // "Apple", "Banana"
```