# godash

[![Build Status](https://travis-ci.org/alexanderbez/godash.svg?branch=master)](https://travis-ci.org/alexanderbez/godash)
[![GoDoc](https://godoc.org/github.com/alexanderbez/godash?status.svg)](https://godoc.org/github.com/alexanderbez/godash)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexanderbez/godash)](https://goreportcard.com/report/github.com/alexanderbez/godash)

A collection of handy utility functions for Golang with no external dependencies!
Most of the magic is done via Go's `reflect` package allowing you to get around
having to implement the same functions with different signatures to handle
desired types.

## API

Slice based operations:

```go
// Filter out unique elements from a slice
in := []string{"foo", "bar", "baz", "foo"}
out := []string{}
err := godash.Unique(in, &out)
// out == ["foo", "bar", "baz"]
```

```go
// Determine if two slices are equal by their contents
s1 := []int{1,2,3,4,5}
s2 := []int{5,4,3,2,1}
ok, err := godash.SliceEqual(s1, s2)
// ok, err == <true, nil>
```

```go
// Determine if an element exists in a slice
s := []string{"foo", "bar", "baz"}
ok, err := godash.Includes(s, "baz")
// ok, err == <true, nil>
```

```go
// Append to a slice only if the slice does not already contain the elements
in := []string{"foo", "bar", "baz", "foo"}
err := godash.AppendUniq(&in, "foo", "hello", "world")
// in == ["foo", "bar", "baz", "foo", "hello", "world"]
```

Encoding based operations:

```go
type Person struct {
	Name string `json:"name"`
}

p := Person{Name: "John Doe"}
```

```go
// Encode to pretty JSON (4 space indent)
bytes, err := godash.ToPrettyJSON(p)
```

```go
// Encode to JSON with no indentation (minified)
bytes, err := godash.ToJSON(p)
```

Map based operations:

```go
// Get all the keys in a map
m := map[string]int{"foo": 3, "bar": 6}
o := []string{}
err := MapKeys(m, &o)
// o == ["key", "bar"]
```

```go
// Get all the values in a map
m := map[string]int{"foo": 3, "bar": 6}
o := []int{}
err := MapValues(m, &o)
// o == [3, 6]
```

Visit [godoc](https://godoc.org/github.com/alexanderbez/godash) for further API
documentation as new functions are implemented.

## Tests

```shell
$ make test
```

## Contributing

1. [Fork it](https://github.com/alexanderbez/godash/fork)
2. Create your feature branch (`git checkout -b feature/my-new-feature`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature/my-new-feature`)
5. Create a new Pull Request
