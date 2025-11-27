package gomaplibmemory

import (
	"testing"

	"golang.org/x/sys/windows"
)

// Test v3 code with old libv1 library

func TestV3CodeWithV1LibWithSize(t *testing.T) {
	// Unsafe memory acess
	cb := windows.NewCallbackCDecl(callback3WithSize)
	procExecuteCallback.Call(cb)
}

func TestV3CodeWithV1LibWithNoSize(t *testing.T) {
	// Safe and correct
	cb := windows.NewCallbackCDecl(callback3WithNoSize)
	procExecuteCallback.Call(cb)
}

// Test v3 code with libv3 library
func TestV3CodeWithV3LibWithSize(t *testing.T) {
	// Safe and correct
	cb := windows.NewCallbackCDecl(callback3WithSize)
	procExecuteCallbackv3.Call(cb)
}

func TestV3CodeWithV3LibWithNoSize(t *testing.T) {
	// Safe and correct
	cb := windows.NewCallbackCDecl(callback3WithNoSize)
	procExecuteCallbackv3.Call(cb)
}

// Test v3 code with libv4 library
func TestV3CodeWithV4LibWithSize(t *testing.T) {
	// Safe and correct, procExecuteCallbackv4 is comptible with v3
	// procExecuteCallbackv4v2 is a breaking change and not comptible with old code
	cb := windows.NewCallbackCDecl(callback3WithSize)
	procExecuteCallbackv4.Call(cb)
}

func TestV3CodeWithV4LibWithNoSize(t *testing.T) {
	// Safe and correct, procExecuteCallbackv4 is comptible with v3
	// procExecuteCallbackv4v2 is a breaking change and not comptible with old code
	cb := windows.NewCallbackCDecl(callback2WithNoSize)
	procExecuteCallbackv4.Call(cb)
}
