package godash

import (
	"reflect"
	"sort"
	"testing"
)

func TestIsPointer(t *testing.T) {
	x := &[]string{}
	y := []string{}

	if r := IsPointer(x); !r {
		t.Errorf("expected 'true' (got %v)", r)
	}
	if r := IsPointer(y); r {
		t.Errorf("expected 'false' (got %v)", r)
	}
}

func TestIsFunction(t *testing.T) {
	x := func() {}
	y := []string{}

	if r := IsFunction(x); !r {
		t.Errorf("expected 'true' (got %v)", r)
	}
	if r := IsPointer(y); r {
		t.Errorf("expected 'false' (got %v)", r)
	}
}

func TestIsSlice(t *testing.T) {
	x := []string{}
	y := func() {}

	if r := IsSlice(x); !r {
		t.Errorf("expected 'true' (got %v)", r)
	}
	if r := IsSlice(y); r {
		t.Errorf("expected 'false' (got %v)", r)
	}
}

func TestIsMap(t *testing.T) {
	x := make(map[string]interface{})
	y := 4
	z := make([]string, 0)

	if r := IsMap(x); !r {
		t.Errorf("expected 'true' (got %v)", r)
	}
	if r := IsMap(y); r {
		t.Errorf("expected 'false' (got %v)", r)
	}
	if r := IsMap(z); r {
		t.Errorf("expected 'false' (got %v)", r)
	}
}

func TestUnique(t *testing.T) {
	in := []string{"a", "a", "c", "d", "c"}
	out := []string{}
	b1 := []int{}

	// Validity tests

	if err := Unique(in, b1); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}
	if err := Unique(in, &b1); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Functionality tests

	if err := Unique(in, &out); err != nil {
		t.Errorf("expected nil error (got %v)", err)
	}
	if len(out) != 3 {
		t.Errorf("expected output slice length of 3 (got %v)", len(out))
	}
}

func TestSliceEqual(t *testing.T) {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"c", "b", "a"}
	s3 := []string{"a", "b", "d", "c"}
	s4 := []int{1, 2, 3}
	x := 0

	// Validity tests

	if _, err := SliceEqual(s1, x); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}
	if _, err := SliceEqual(s1, s4); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Functionality tests

	if r, _ := SliceEqual(s1, x); r {
		t.Errorf("expected 'false' (got %v)", r)
	}
	if r, _ := SliceEqual(s1, s2); !r {
		t.Errorf("expected 'true' (got %v)", r)
	}
	if _, err := SliceEqual(s1, s2); err != nil {
		t.Errorf("expected nil error (got %v)", err)
	}
	if r, _ := SliceEqual(s1, s3); r {
		t.Errorf("expected 'false' (got %v)", r)
	}
	if _, err := SliceEqual(s1, s3); err != nil {
		t.Errorf("expected nil error (got %v)", err)
	}
	if r, _ := SliceEqual(s1, s4); r {
		t.Errorf("expected 'false' (got %v)", r)
	}
}

func TestIncludes(t *testing.T) {
	s1 := []string{"a", "b", "c"}

	// Validity tests

	if _, err := Includes(s1, 1); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}
	if _, err := Includes(1, 1); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Functionality tests

	if r, _ := Includes(s1, "a"); !r {
		t.Errorf("expected 'true' (got %v)", r)
	}
	if r, _ := Includes(s1, "d"); r {
		t.Errorf("expected 'false' (got %v)", r)
	}
}

func TestAppendUniq(t *testing.T) {
	x := 1
	s1 := []string{"a", "b", "c"}

	// Validity tests

	if err := AppendUniq(&x, 1); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}
	if err := AppendUniq(x, 1); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}
	if err := AppendUniq(&s1, 1); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}
	if err := AppendUniq(&s1, "d", 1); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Functionality tests

	s2 := []string{"a", "b", "c"}
	expected := []string{"a", "b", "c", "d"}

	if err := AppendUniq(&s2, "a", "d", "a", "d"); err != nil {
		t.Errorf("expected nil error (got %v)", err)
	}

	sort.Strings(s2)
	sort.Strings(expected)

	if r := reflect.DeepEqual(s2, expected); !r {
		t.Errorf("expected correct slice (got %v)", s2)
	}
}

func TestKeys(t *testing.T) {
	m := map[string]interface{}{"a": 3, "b": false}
	a := []int{}
	b := 3
	c := []string{}

	// Validity tests

	// Test the input variable is a map
	if err := MapKeys(&m, a); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test the input variable is a map
	if err := MapKeys(a, a); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test the output variable is a pointer
	if err := MapKeys(m, a); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test the output variable is a pointer of the valid type
	if err := MapKeys(m, &a); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test the output variable is a pointer
	if err := MapKeys(m, b); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test the output variable is a pointer
	if err := MapKeys(m, &b); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Functionality tests

	expected := []string{"a", "b"}

	if err := MapKeys(m, &c); err != nil {
		t.Errorf("expected nil error (got %v)", err)
	}

	sort.Strings(expected)
	sort.Strings(c)

	if ok := reflect.DeepEqual(c, expected); !ok {
		t.Errorf("expected (%v) (got %v)", expected, c)
	}
}

func TestMapValues(t *testing.T) {
	m := map[string]int{"foo": 3, "bar": 6}
	a := []string{}
	b := 3
	c := []int{}

	// Validity tests

	// Test the input variable is a map
	if err := MapValues(&m, a); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test the input variable is a map
	if err := MapValues(a, a); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test the output variable is a pointer
	if err := MapValues(m, a); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test the output variable is a pointer of the valid type
	if err := MapValues(m, &a); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test the output variable is a pointer
	if err := MapValues(m, b); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test the output variable is a pointer
	if err := MapValues(m, &b); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Functionality tests

	expected := []int{3, 6}

	// Test when no error should be returned
	if err := MapValues(m, &c); err != nil {
		t.Errorf("expected nil error (got %v)", err)
	}

	sort.Ints(expected)
	sort.Ints(c)

	// Test the correct values were returned
	if ok := reflect.DeepEqual(c, expected); !ok {
		t.Errorf("expected (%v) (got %v)", expected, c)
	}
}

func TestIntersect(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{3, 4, 5, 1}
	s3 := []int{7, 8, 9, 0}
	s4 := []int{3, 4, 4, 3}
	x := 3
	y := []string{}
	out := []int{}

	// Validity tests

	// Test if the first slice is actually a slice
	if err := Intersect(x, s2, &out); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test if the second slice is actually a slice
	if err := Intersect(s1, x, &out); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test if the output argument is a pointer
	if err := Intersect(x, s2, out); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test if the output argument is a pointer to a slice
	if err := Intersect(s1, s2, &x); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test if the two slices are of the same type
	if err := Intersect(s1, y, &out); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}
	if err := Intersect(y, s2, &out); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test if the output slice reference is of the same type
	if err := Intersect(s1, s2, &y); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Functionality tests

	var (
		expected []int
		err      error
	)

	// Test when there is an intersection
	out = []int{}
	expected = []int{3, 4, 1}
	err = Intersect(s1, s2, &out)

	sort.Ints(expected)
	sort.Ints(out)

	if err != nil {
		t.Errorf("expected nil error (got %v)", err)
	}
	if ok := reflect.DeepEqual(out, expected); !ok {
		t.Errorf("expected (%v) (got %v)", expected, out)
	}

	// Test when there is no intersection
	out = []int{}
	expected = []int{}
	err = Intersect(s1, s3, &out)

	if err != nil {
		t.Errorf("expected nil error (got %v)", err)
	}
	if ok := reflect.DeepEqual(out, expected); !ok {
		t.Errorf("expected (%v) (got %v)", expected, out)
	}

	// Test intersection with duplicates
	out = []int{}
	expected = []int{3, 4}
	err = Intersect(s1, s4, &out)

	sort.Ints(expected)
	sort.Ints(out)

	if err != nil {
		t.Errorf("expected nil error (got %v)", err)
	}
	if ok := reflect.DeepEqual(out, expected); !ok {
		t.Errorf("expected (%v) (got %v)", expected, out)
	}
}
