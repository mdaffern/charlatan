// generated by "charlatan -dir=testdata/importer -output=testdata/importer/importer.go Importer".  DO NOT EDIT.

package main

import "reflect"
import . "fmt"
import z "strings"

// ImporterScanInvocation represents a single call of FakeImporter.Scan
type ImporterScanInvocation struct {
	Parameters struct {
		Ident1 *Scanner
	}
	Results struct {
		Ident2 z.Reader
	}
}

// ImporterTestingT represents the methods of "testing".T used by charlatan Fakes.  It avoids importing the testing package.
type ImporterTestingT interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Helper()
}

/*
FakeImporter is a mock implementation of Importer for testing.
Use it in your tests as in this example:

	package example

	func TestWithImporter(t *testing.T) {
		f := &main.FakeImporter{
			ScanHook: func(ident1 *Scanner) (ident2 z.Reader) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeScan ...
		f.AssertScanCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeScan.
*/
type FakeImporter struct {
	ScanHook func(*Scanner) z.Reader

	ScanCalls []*ImporterScanInvocation
}

// NewFakeImporterDefaultPanic returns an instance of FakeImporter with all hooks configured to panic
func NewFakeImporterDefaultPanic() *FakeImporter {
	return &FakeImporter{
		ScanHook: func(*Scanner) (ident2 z.Reader) {
			panic("Unexpected call to Importer.Scan")
		},
	}
}

// NewFakeImporterDefaultFatal returns an instance of FakeImporter with all hooks configured to call t.Fatal
func NewFakeImporterDefaultFatal(t ImporterTestingT) *FakeImporter {
	return &FakeImporter{
		ScanHook: func(*Scanner) (ident2 z.Reader) {
			t.Fatal("Unexpected call to Importer.Scan")
			return
		},
	}
}

// NewFakeImporterDefaultError returns an instance of FakeImporter with all hooks configured to call t.Error
func NewFakeImporterDefaultError(t ImporterTestingT) *FakeImporter {
	return &FakeImporter{
		ScanHook: func(*Scanner) (ident2 z.Reader) {
			t.Error("Unexpected call to Importer.Scan")
			return
		},
	}
}

func (f *FakeImporter) Reset() {
	f.ScanCalls = []*ImporterScanInvocation{}
}

func (_f1 *FakeImporter) Scan(ident1 *Scanner) (ident2 z.Reader) {
	if _f1.ScanHook == nil {
		panic("Importer.Scan() called but FakeImporter.ScanHook is nil")
	}

	invocation := new(ImporterScanInvocation)
	_f1.ScanCalls = append(_f1.ScanCalls, invocation)

	invocation.Parameters.Ident1 = ident1

	ident2 = _f1.ScanHook(ident1)

	invocation.Results.Ident2 = ident2

	return
}

// SetScanStub configures Importer.Scan to always return the given values
func (_f2 *FakeImporter) SetScanStub(ident2 z.Reader) {
	_f2.ScanHook = func(*Scanner) z.Reader {
		return ident2
	}
}

// SetScanInvocation configures Importer.Scan to return the given results when called with the given parameters
// If no match is found for an invocation the result(s) of the fallback function are returned
func (_f3 *FakeImporter) SetScanInvocation(calls_f4 []*ImporterScanInvocation, fallback_f5 func() z.Reader) {
	_f3.ScanHook = func(ident1 *Scanner) (ident2 z.Reader) {
		for _, call := range calls_f4 {
			if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
				ident2 = call.Results.Ident2

				return
			}
		}

		return fallback_f5()
	}
}

// ScanCalled returns true if FakeImporter.Scan was called
func (f *FakeImporter) ScanCalled() bool {
	return len(f.ScanCalls) != 0
}

// AssertScanCalled calls t.Error if FakeImporter.Scan was not called
func (f *FakeImporter) AssertScanCalled(t ImporterTestingT) {
	t.Helper()
	if len(f.ScanCalls) == 0 {
		t.Error("FakeImporter.Scan not called, expected at least one")
	}
}

// ScanNotCalled returns true if FakeImporter.Scan was not called
func (f *FakeImporter) ScanNotCalled() bool {
	return len(f.ScanCalls) == 0
}

// AssertScanNotCalled calls t.Error if FakeImporter.Scan was called
func (f *FakeImporter) AssertScanNotCalled(t ImporterTestingT) {
	t.Helper()
	if len(f.ScanCalls) != 0 {
		t.Error("FakeImporter.Scan called, expected none")
	}
}

// ScanCalledOnce returns true if FakeImporter.Scan was called exactly once
func (f *FakeImporter) ScanCalledOnce() bool {
	return len(f.ScanCalls) == 1
}

// AssertScanCalledOnce calls t.Error if FakeImporter.Scan was not called exactly once
func (f *FakeImporter) AssertScanCalledOnce(t ImporterTestingT) {
	t.Helper()
	if len(f.ScanCalls) != 1 {
		t.Errorf("FakeImporter.Scan called %d times, expected 1", len(f.ScanCalls))
	}
}

// ScanCalledN returns true if FakeImporter.Scan was called at least n times
func (f *FakeImporter) ScanCalledN(n int) bool {
	return len(f.ScanCalls) >= n
}

// AssertScanCalledN calls t.Error if FakeImporter.Scan was called less than n times
func (f *FakeImporter) AssertScanCalledN(t ImporterTestingT, n int) {
	t.Helper()
	if len(f.ScanCalls) < n {
		t.Errorf("FakeImporter.Scan called %d times, expected >= %d", len(f.ScanCalls), n)
	}
}

// ScanCalledWith returns true if FakeImporter.Scan was called with the given values
func (_f6 *FakeImporter) ScanCalledWith(ident1 *Scanner) (found bool) {
	for _, call := range _f6.ScanCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	return
}

// AssertScanCalledWith calls t.Error if FakeImporter.Scan was not called with the given values
func (_f7 *FakeImporter) AssertScanCalledWith(t ImporterTestingT, ident1 *Scanner) {
	t.Helper()
	var found bool
	for _, call := range _f7.ScanCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeImporter.Scan not called with expected parameters")
	}
}

// ScanCalledOnceWith returns true if FakeImporter.Scan was called exactly once with the given values
func (_f8 *FakeImporter) ScanCalledOnceWith(ident1 *Scanner) bool {
	var count int
	for _, call := range _f8.ScanCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	return count == 1
}

// AssertScanCalledOnceWith calls t.Error if FakeImporter.Scan was not called exactly once with the given values
func (_f9 *FakeImporter) AssertScanCalledOnceWith(t ImporterTestingT, ident1 *Scanner) {
	t.Helper()
	var count int
	for _, call := range _f9.ScanCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeImporter.Scan called %d times with expected parameters, expected one", count)
	}
}

// ScanResultsForCall returns the result values for the first call to FakeImporter.Scan with the given values
func (_f10 *FakeImporter) ScanResultsForCall(ident1 *Scanner) (ident2 z.Reader, found bool) {
	for _, call := range _f10.ScanCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			ident2 = call.Results.Ident2
			found = true
			break
		}
	}

	return
}
