// generated by "charlatan -test.v=true -test.coverprofile=/Users/tmh/Percolate/charlatan/build/coverage/test_coverage.out -test.outputdir /Users/tmh/Percolate/charlatan".  DO NOT EDIT.

package main

import (
	"reflect"
	"testing"
)

// PointInvocation represents a single call of FakePointer.Point
type PointInvocation struct {
	Parameters struct {
		Ident1 *string
	}
	Results struct {
		Ident2 *int
	}
}

/*
FakePointer is a mock implementation of Pointer for testing.
Use it in your tests as in this example:

	package example

	func TestWithPointer(t *testing.T) {
		f := &main.FakePointer{
			PointHook: func(ident1 *string) (ident2 *int) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakePoint ...
		f.AssertPointCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a painc if any
unexpected calls are made to FakePoint.
*/
type FakePointer struct {
	PointHook func(*string) *int

	PointCalls []*PointInvocation
}

// NewFakePointerDefaultPanic returns an instance of FakePointer with all hooks configured to panic
func NewFakePointerDefaultPanic() *FakePointer {
	return &FakePointer{
		PointHook: func(*string) (ident2 *int) {
			panic("Unexpected call to Pointer.Point")
			return
		},
	}
}

// NewFakePointerDefaultFatal returns an instance of FakePointer with all hooks configured to call t.Fatal
func NewFakePointerDefaultFatal(t *testing.T) *FakePointer {
	return &FakePointer{
		PointHook: func(*string) (ident2 *int) {
			t.Fatal("Unexpected call to Pointer.Point")
			return
		},
	}
}

// NewFakePointerDefaultError returns an instance of FakePointer with all hooks configured to call t.Error
func NewFakePointerDefaultError(t *testing.T) *FakePointer {
	return &FakePointer{
		PointHook: func(*string) (ident2 *int) {
			t.Error("Unexpected call to Pointer.Point")
			return
		},
	}
}

func (_f1 *FakePointer) Point(ident1 *string) (ident2 *int) {
	invocation := new(PointInvocation)

	invocation.Parameters.Ident1 = ident1

	ident2 = _f1.PointHook(ident1)

	invocation.Results.Ident2 = ident2

	_f1.PointCalls = append(_f1.PointCalls, invocation)

	return
}

// PointCalled returns true if FakePointer.Point was called
func (f *FakePointer) PointCalled() bool {
	return len(f.PointCalls) != 0
}

// AssertPointCalled calls t.Error if FakePointer.Point was not called
func (f *FakePointer) AssertPointCalled(t *testing.T) {
	t.Helper()
	if len(f.PointCalls) == 0 {
		t.Error("FakePointer.Point not called, expected at least one")
	}
}

// PointNotCalled returns true if FakePointer.Point was not called
func (f *FakePointer) PointNotCalled() bool {
	return len(f.PointCalls) == 0
}

// AssertPointNotCalled calls t.Error if FakePointer.Point was called
func (f *FakePointer) AssertPointNotCalled(t *testing.T) {
	t.Helper()
	if len(f.PointCalls) != 0 {
		t.Error("FakePointer.Point called, expected none")
	}
}

// PointCalledOnce returns true if FakePointer.Point was called exactly once
func (f *FakePointer) PointCalledOnce() bool {
	return len(f.PointCalls) == 1
}

// AssertPointCalledOnce calls t.Error if FakePointer.Point was not called exactly once
func (f *FakePointer) AssertPointCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.PointCalls) != 1 {
		t.Errorf("FakePointer.Point called %d times, expected 1", len(f.PointCalls))
	}
}

// PointCalledN returns true if FakePointer.Point was called at least n times
func (f *FakePointer) PointCalledN(n int) bool {
	return len(f.PointCalls) >= n
}

// AssertPointCalledN calls t.Error if FakePointer.Point was called less than n times
func (f *FakePointer) AssertPointCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.PointCalls) < n {
		t.Errorf("FakePointer.Point called %d times, expected >= %d", len(f.PointCalls), n)
	}
}

// PointCalledWith returns true if FakePointer.Point was called with the given values
func (_f2 *FakePointer) PointCalledWith(ident1 *string) (found bool) {
	for _, call := range _f2.PointCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	return
}

// AssertPointCalledWith calls t.Error if FakePointer.Point was not called with the given values
func (_f3 *FakePointer) AssertPointCalledWith(t *testing.T, ident1 *string) {
	t.Helper()
	var found bool
	for _, call := range _f3.PointCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakePointer.Point not called with expected parameters")
	}
}

// PointCalledOnceWith returns true if FakePointer.Point was called exactly once with the given values
func (_f4 *FakePointer) PointCalledOnceWith(ident1 *string) bool {
	var count int
	for _, call := range _f4.PointCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	return count == 1
}

// AssertPointCalledOnceWith calls t.Error if FakePointer.Point was not called exactly once with the given values
func (_f5 *FakePointer) AssertPointOnceCalledWith(t *testing.T, ident1 *string) {
	t.Helper()
	var count int
	for _, call := range _f5.PointCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakePointer.Point called %d times with expected parameters, expected one", count)
	}
}

// PointResultsForCall returns the result values for the first call to FakePointer.Point with the given values
func (_f6 *FakePointer) PointResultsForCall(ident1 *string) (ident2 *int, found bool) {
	for _, call := range _f6.PointCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			ident2 = call.Results.Ident2
			found = true
			break
		}
	}

	return
}
