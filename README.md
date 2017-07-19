# godash

[![Build Status](https://travis-ci.org/alexanderbez/godash.svg?branch=master)](https://travis-ci.org/alexanderbez/godash)
[![GoDoc](https://godoc.org/github.com/alexanderbez/godash?status.svg)](https://godoc.org/github.com/alexanderbez/godash)


A collection of handy utility functions for Golang with no external dependencies!
Most of the magic is done via Go's `reflect` package. While this does allow you to
get around having to implement the same functions with different signatures to handle
desired types, there are performance implications.

## API

Slice based operations:

```go
import (
    // ...
    "github.com/alexanderbez/godash"
)
// Filter out unique elements from a slice
in := []string{"foo", "bar", "baz", "foo"}
out := []string{}
godash.Unique(in, &out)

// Determine if two slices are equal by their contents
s1 := []int{1,2,3,4,5}
s2 := []int{5,4,3,2,1}
godash.SliceEqual(s1, s2)

// Determine if an element exists in a slice
s := []string{"foo", "bar", "baz"}
godash.Includes(s, "baz")

// Append to a slice only if the slice does not already contain the elements
in := []string{"foo", "bar", "baz", "foo"}
godash.AppendUniq(&in, "foo", "hello", "world")
```

Encoding based operations:

```go
import (
    // ...
    "github.com/alexanderbez/godash"
)

type Person struct {
    Name string `json:"name"`
}
s := Person{Name: "John Doe"}

// Encode to pretty JSON (4 space indent)
pBytes, err := godash.ToPrettyJSON(s)

// Encode to JSON with no indentation (minified)
bytes, err := godash.ToJSON(s)
```

Visit [godoc](https://godoc.org/github.com/alexanderbez/godash) for further API documentation as new functions are implemented.

## Tests

```shell
$ make test
```

## Benchmarks

Coming soon...

## TODO:

- Benchmark tests
- More utility functions

## Contributing

1. [Fork it](https://github.com/alexanderbez/godash/fork)
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
