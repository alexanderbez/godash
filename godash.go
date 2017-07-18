// Package godash - Provides a collection of handy utility functions.
//
// Copyright 2017 Aleksandr Bezobchuk. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file.
package godash

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type (
	// Value reflects any type empty interface
	// (named collections of method signatures)
	Value interface{}
	// Slice reflects a slice of any type
	Slice interface{}
	// Pointer reflects a pointer that references any type
	Pointer interface{}
)

// IsPointer returns true if the supplied argument is a pointer or false
// otherwise.
func IsPointer(value Value) bool {
	return reflect.ValueOf(value).Kind() == reflect.Ptr
}

// IsFunction returns true if the supplied argument is a function or false
// otherwise.
func IsFunction(value Value) bool {
	return reflect.ValueOf(value).Kind() == reflect.Func
}

// IsSlice returns true if the supplied argument is a slice or false otherwise.
func IsSlice(value Value) bool {
	kind := reflect.ValueOf(value).Kind()
	return kind == reflect.Slice || kind == reflect.Array
}

// Unique returns a unique collection of elements found in inSlice. The
// resulting elements are added to a pointer referencing a slice of the same
// type. If inSlice is not a slice, outPtr not a pointer, or the underlying
// types differ, an error is returned. Operation runs in O(n) time, where n is
// the total number of elements in the source slice.
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

	if !inSliceTyp.AssignableTo(outValue.Type().Elem()) {
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
// the other. If the either parameter is not a slice or the types do not match,
// an error is returned. Operation runs in O(n^2) time, where n is the total
// number of elements found in either slice.
func SliceEqual(slice1 Slice, slice2 Slice) (bool, error) {
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

// ToPrettyJSON converts a compatible interface to pretty JSON format.
func ToPrettyJSON(value Value) (r []byte, err error) {
	r, err = json.MarshalIndent(value, "", "    ")
	return
}

// ToJSON converts a compatible interface to JSON format (minified).
func ToJSON(value Value) (r []byte, err error) {
	r, err = json.Marshal(value)
	return
}
