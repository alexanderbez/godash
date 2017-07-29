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

	// Test argument types
	if err := Unique(in, b1); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}
	if err := Unique(in, &b1); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test correct functionality
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

	// Test argument types
	if _, err := SliceEqual(s1, x); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}
	if _, err := SliceEqual(s1, s4); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test correct functionality
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

	// Test argument types
	if _, err := Includes(s1, 1); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}
	if _, err := Includes(1, 1); err == nil {
		t.Errorf("expected an error (got %v)", err)
	}

	// Test correct functionality
	if r, _ := Includes(s1, "a"); !r {
		t.Errorf("expected 'true' (got %v)", r)
	}
	if r, _ := Includes(s1, "d"); r {
		t.Errorf("expected 'false' (got %v)", r)
	}
}

func TestAppendUniq(t *testing.T) {
	// Test argument types
	x := 1
	s1 := []string{"a", "b", "c"}

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

	// Test correct functionality
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

	// Test argument types
	if err := MapKeys(&m, a); err == nil {
		// Test the input variable is a map
		t.Errorf("expected an error (got %v)", err)
	}
	if err := MapKeys(a, a); err == nil {
		// Test the input variable is a map
		t.Errorf("expected an error (got %v)", err)
	}
	if err := MapKeys(m, a); err == nil {
		// Test the output variable is a pointer
		t.Errorf("expected an error (got %v)", err)
	}
	if err := MapKeys(m, &a); err == nil {
		// Test the output variable is a pointer of the valid type
		t.Errorf("expected an error (got %v)", err)
	}
	if err := MapKeys(m, b); err == nil {
		// Test the output variable is a pointer
		t.Errorf("expected an error (got %v)", err)
	}
	if err := MapKeys(m, &b); err == nil {
		// Test the output variable is a pointer
		t.Errorf("expected an error (got %v)", err)
	}

	// Test correct functionality
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

	// Test argument types
	if err := MapValues(&m, a); err == nil {
		// Test the input variable is a map
		t.Errorf("expected an error (got %v)", err)
	}
	if err := MapValues(a, a); err == nil {
		// Test the input variable is a map
		t.Errorf("expected an error (got %v)", err)
	}
	if err := MapValues(m, a); err == nil {
		// Test the output variable is a pointer
		t.Errorf("expected an error (got %v)", err)
	}
	if err := MapValues(m, &a); err == nil {
		// Test the output variable is a pointer of the valid type
		t.Errorf("expected an error (got %v)", err)
	}
	if err := MapValues(m, b); err == nil {
		// Test the output variable is a pointer
		t.Errorf("expected an error (got %v)", err)
	}
	if err := MapValues(m, &b); err == nil {
		// Test the output variable is a pointer
		t.Errorf("expected an error (got %v)", err)
	}

	// Test correct functionality
	expected := []int{3, 6}

	if err := MapValues(m, &c); err != nil {
		t.Errorf("expected nil error (got %v)", err)
	}

	sort.Ints(expected)
	sort.Ints(c)

	if ok := reflect.DeepEqual(c, expected); !ok {
		t.Errorf("expected (%v) (got %v)", expected, c)
	}
}
