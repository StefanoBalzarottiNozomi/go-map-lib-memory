package gomaplibmemory

import (
	"testing"

	"golang.org/x/sys/windows"
)

// Test v4 code with old libv1 library

func TestV4CodeWithV1LibWithSize(t *testing.T) {
	// Safe and correct, but misleading, size is 16384, not 1
	cb := windows.NewCallbackCDecl(callback4WithSize)
	procExecuteCallback.Call(cb)
}

func TestV4CodeWithV1LibWithNoSize(t *testing.T) {
	// Safe and correct
	cb := windows.NewCallbackCDecl(callback4WithNoSize)
	procExecuteCallback.Call(cb)
}

// Test v4 code with libv4 library
func TestV4CodeWithV4LibWithSize(t *testing.T) {
	// Safe and correct, but misleading, size is not 1, it depends on the data structure
	cb := windows.NewCallbackCDecl(callback4WithSize)
	procExecuteCallbackv4v2.Call(cb)
}

func TestV4CodeWithV4LibWithNoSize(t *testing.T) {
	// Safe and correct
	cb := windows.NewCallbackCDecl(callback4WithNoSize)
	procExecuteCallbackv4v2.Call(cb)
}
