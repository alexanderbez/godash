// Package godash provides a collection of handy utility functions such as
// working with slices in a generic fashion and various common encoding
// shortcuts.
package godash

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type (
	// Value reflects a placeholder for any type
	Value interface{}
	// Slice reflects a slice or array of any type
	Slice interface{}
	// Pointer reflects a pointer that references any type
	Pointer interface{}
	// Map reflects a map of any type
	Map interface{}
)

// IsPointer returns true if the specified argument is a pointer and false
// otherwise.
func IsPointer(value Value) bool {
	return reflect.ValueOf(value).Kind() == reflect.Ptr
}

// IsFunction returns true if the specified argument is a function and false
// otherwise.
func IsFunction(value Value) bool {
	return reflect.ValueOf(value).Kind() == reflect.Func
}

// IsSlice returns true if the specified argument is a slice or array and false
// otherwise.
func IsSlice(value Value) bool {
	kind := reflect.ValueOf(value).Kind()
	return kind == reflect.Slice || kind == reflect.Array
}

// IsMap returns true if the specified argument is a map and false otherwise.
func IsMap(Value Value) bool {
	return reflect.ValueOf(Value).Kind() == reflect.Map
}

// Unique returns a unique collection of elements found in inSlice. The
// resulting elements are added to a slice referenced by a pointer outPtr.
// If inSlice is not a slice, outPtr not a pointer, or the underlying types
// differ, an error is returned. Operation runs in O(n) time, where n is the
// total number of elements in the source slice.
func Unique(inSlice Slice, outPtr Pointer) error {
	if !IsSlice(inSlice) {
		return fmt.Errorf("argument type '%T' is not a slice", inSlice)
	} else if !IsPointer(outPtr) {
		return fmt.Errorf("argument type '%T' is not a pointer", outPtr)
	}

	inSliceVal := reflect.ValueOf(inSlice)
	inSliceTyp := inSliceVal.Type()

	outValue := reflect.ValueOf(outPtr)
	outType := outValue.Type()

	if !inSliceTyp.AssignableTo(outType.Elem()) {
		return fmt.Errorf("input type '%v' can't be assigned to output type '%v' ", inSliceTyp, outType.Elem())
	}

	outSlice := reflect.MakeSlice(inSliceTyp, 0, 0)
	valCounts := make(map[interface{}]bool)

	for i := 0; i < inSliceVal.Len(); i++ {
		valCounts[inSliceVal.Index(i).Interface()] = true
	}

	for k := range valCounts {
		outSlice = reflect.Append(outSlice, reflect.ValueOf(k))
	}

	// Assign the new unique list to the provided output slice pointer
	outValue.Elem().Set(outSlice)

	return nil
}

// SliceEqual determines if the contents of one slice equals the contents of
// the other. If either parameter is not a slice or the types do not match, an
// error is returned. Operation runs in O(n^2) time, where n is the total
// number of elements found in either slice.
func SliceEqual(slice1, slice2 Slice) (bool, error) {
	if !IsSlice(slice1) {
		return false, fmt.Errorf("argument type '%T' is not a slice", slice1)
	} else if !IsSlice(slice2) {
		return false, fmt.Errorf("argument type '%T' is not a slice", slice2)
	}

	slice1Val := reflect.ValueOf(slice1)
	slice2Val := reflect.ValueOf(slice2)

	if slice1Val.Type().Elem() != slice2Val.Type().Elem() {
		return false, fmt.Errorf("type of '%v' does not match type of '%v'", slice1Val.Type().Elem(), slice2Val.Type().Elem())
	}

	if slice1Val.Len() != slice2Val.Len() {
		return false, nil
	}

	result := true
	i, n := 0, slice1Val.Len()

	for i < n {
		j := 0
		e := false
		for j < n && !e {
			if slice1Val.Index(i).Interface() == slice2Val.Index(j).Interface() {
				e = true
			}
			j++
		}
		if !e {
			result = false
		}
		i++
	}

	return result, nil
}

// Includes determines if inSlice contains the element Value. If the specified
// inSlice is not a slice or if element is not the same element type, then an
// error is returned. Operation runs in O(n + k) time, where n is the total
// number of elements in the input slice and k is the number of elements to
// append.
func Includes(inSlice Slice, element Value) (bool, error) {
	if !IsSlice(inSlice) {
		return false, fmt.Errorf("argument type '%T' is not a slice", inSlice)
	}

	inSliceVal := reflect.ValueOf(inSlice)
	elValue := reflect.ValueOf(element)

	if inSliceVal.Type().Elem() != elValue.Type() {
		return false, fmt.Errorf("type of '%v' does not match type of '%v'", inSliceVal.Type().Elem(), elValue.Type())
	}

	result := false
	i, n := 0, inSliceVal.Len()

	for i < n && !result {
		if inSliceVal.Index(i).Interface() == elValue.Interface() {
			result = true
		}
		i++
	}

	return result, nil
}

// AppendUniq appends each element found in elements to the slice referenced by
// inPtr if the element does not already exist in the slice. An error is
// returned and the append considered a no-op if inPtr is not a pointer to a
// slice or if any element found in the elements slice is not of the same type.
func AppendUniq(inPtr Pointer, elements ...Value) error {
	if !IsPointer(inPtr) {
		return fmt.Errorf("argument type '%T' is not a pointer", inPtr)
	}

	inPtrValue := reflect.ValueOf(inPtr)
	inPtrElem := inPtrValue.Elem()

	if !IsSlice(inPtrElem.Interface()) {
		return fmt.Errorf("argument type '%T' is not a pointer to a slice", inPtr)
	}

	inPtrElemType := inPtrElem.Type().Elem()

	elValue := reflect.ValueOf(elements)

	outLen := inPtrElem.Len()
	outSlice := reflect.MakeSlice(inPtrElem.Type(), outLen, outLen)

	reflect.Copy(outSlice, inPtrElem)

	for i := 0; i < elValue.Len(); i++ {
		curr := elValue.Index(i)
		currVal := curr.Interface()
		currValType := reflect.TypeOf(currVal)

		if !inPtrElemType.AssignableTo(currValType) {
			return fmt.Errorf("input type '%v' can't be assigned to output type '%v' ", currValType, inPtrElemType)
		}

		contained, err := Includes(outSlice.Interface(), currVal)
		if err != nil {
			return err
		}
		if !contained {
			outSlice = reflect.Append(outSlice, reflect.ValueOf(currVal))
		}
	}

	// Assign the new unique list to the provided input slice pointer
	inPtrValue.Elem().Set(outSlice)

	return nil
}

// ToPrettyJSON converts a compatible interface to pretty JSON format by using
// four space indentation.
func ToPrettyJSON(value Value) (r []byte, err error) {
	r, err = json.MarshalIndent(value, "", "    ")
	return
}

// ToJSON converts a compatible interface to JSON format (minified).
func ToJSON(value Value) (r []byte, err error) {
	r, err = json.Marshal(value)
	return
}

// MapKeys appends all of the keys in the provided map, inMap, to the slice
// referenced by the pointer outPtr. An error is returned if the argument inMap
// is not a valid map or if the slice that outPtr references is not a slice of
// the same type as the key type in inMap.
func MapKeys(inMap Map, outPtr Pointer) error {
	if !IsMap(inMap) {
		return fmt.Errorf("argument type '%T' is not a map", inMap)
	} else if !IsPointer(outPtr) {
		return fmt.Errorf("argument type '%T' is not a pointer", outPtr)
	}

	inMapVal := reflect.ValueOf(inMap)
	inMapTyp := inMapVal.Type()

	outValue := reflect.ValueOf(outPtr)
	outTypeEl := outValue.Type().Elem()

	if !IsSlice(outValue.Elem().Interface()) {
		return fmt.Errorf("argument type '%T' is not a pointer to a slice", outValue.Elem().Interface())
	}

	if !outTypeEl.Elem().AssignableTo(inMapTyp.Key()) {
		return fmt.Errorf("input type '%v' can't be assigned to output type '%v' ", outTypeEl.Elem(), inMapTyp.Key())
	}

	// Copy keys from map to a temporary slice referenced by outSlice
	outSlice := reflect.MakeSlice(outTypeEl, 0, 0)
	for _, key := range inMapVal.MapKeys() {
		outSlice = reflect.Append(outSlice, key)
	}

	// Set the value of outValue to the pointer referenced by the temporary
	// slice referenced by outSlice (copying the contents).
	outValue.Elem().Set(outSlice)

	return nil
}

// MapValues appends all of the values in the provided map, inMap, to the slice
// referenced by the pointer outPtr. An error is returned if the argument inMap
// is not a valid map or if the slice that outPtr references is not a slice of
// the same type as the value type in inMap.
func MapValues(inMap Map, outPtr Pointer) error {
	if !IsMap(inMap) {
		return fmt.Errorf("argument type '%T' is not a map", inMap)
	} else if !IsPointer(outPtr) {
		return fmt.Errorf("argument type '%T' is not a pointer", outPtr)
	}

	inMapVal := reflect.ValueOf(inMap)
	inMapTyp := inMapVal.Type()
	inMapValType := inMapTyp.Elem()

	outValue := reflect.ValueOf(outPtr)
	outTypeEl := outValue.Type().Elem()

	if !IsSlice(outValue.Elem().Interface()) {
		return fmt.Errorf("argument type '%T' is not a pointer to a slice", outValue.Elem().Interface())
	}

	if !inMapValType.AssignableTo(outTypeEl.Elem()) {
		return fmt.Errorf("input type '%v' can't be assigned to output type '%v' ", inMapValType, outTypeEl.Elem())
	}

	// Copy values from map to a temporary slice referenced by outSlice
	outSlice := reflect.MakeSlice(outTypeEl, 0, 0)
	for _, key := range inMapVal.MapKeys() {
		outSlice = reflect.Append(outSlice, inMapVal.MapIndex(key))
	}

	// Set the value of outValue to the pointer referenced by the temporary
	// slice referenced by outSlice (copying the contents).
	outValue.Elem().Set(outSlice)

	return nil
}

// Intersect appends all the common distinct elements found in slice1 and
// slice2 to the slice referenced by outPtr. An error is returned if any of the
// following conditions are met, slice1 or slice2 are not slices, they are not
// slices of the same type, or the value referenced by outPtr is not a slice,
// or a slice of the same type as slice1 and slice2.
func Intersect(slice1, slice2 Slice, outPtr Pointer) error {
	if !IsSlice(slice1) {
		return fmt.Errorf("argument type '%T' is not a slice", slice1)
	} else if !IsSlice(slice2) {
		return fmt.Errorf("argument type '%T' is not a slice", slice2)
	} else if !IsPointer(outPtr) {
		return fmt.Errorf("argument type '%T' is not a pointer", outPtr)
	}

	slice1Val := reflect.ValueOf(slice1)
	slice1Type := slice1Val.Type()

	slice2Val := reflect.ValueOf(slice2)
	slice2Type := slice2Val.Type()

	if !slice1Type.AssignableTo(slice2Type) {
		return fmt.Errorf("incompatible slice types '%v' and '%v'", slice1Type, slice2Type)
	}

	outPtrValue := reflect.ValueOf(outPtr)
	outPtrType := outPtrValue.Type()

	if !slice1Type.AssignableTo(outPtrType.Elem()) {
		return fmt.Errorf("input type '%v' can't be assigned to output type '%v'", slice1Type, outPtrType.Elem())
	}

	var (
		len         int
		iterSlice   reflect.Value
		searchSlice Slice
	)

	// Determine which slice to iterate over and which slice to search through
	if slice1Val.Len() < slice2Val.Len() {
		len = slice1Val.Len()
		iterSlice = slice1Val
		searchSlice = slice2
	} else {
		len = slice2Val.Len()
		iterSlice = slice2Val
		searchSlice = slice1
	}

	distinct := map[interface{}]bool{}

	// Iterate over iterSlice and only append the current element to outSlice
	// if it exists in searchSlice.
	outSlice := reflect.MakeSlice(slice1Type, 0, 0)
	for i := 0; i < len; i++ {
		intrVal := iterSlice.Index(i).Interface()

		if ok, _ := Includes(searchSlice, intrVal); ok {
			if _, ok = distinct[intrVal]; !ok {
				outSlice = reflect.Append(outSlice, reflect.ValueOf(intrVal))
			}

			distinct[intrVal] = true
		}
	}

	// Assign the new unique list to the provided output slice pointer
	outPtrValue.Elem().Set(outSlice)

	return nil
}
