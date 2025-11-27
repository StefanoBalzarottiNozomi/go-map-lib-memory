package gomaplibmemory

import (
	"testing"

	"golang.org/x/sys/windows"
)

// Test v2 code with old libv1 library

func TestV2CodeWithV1LibWithSize(t *testing.T) {
	// Unsafe memory acess
	cb := windows.NewCallbackCDecl(callback2WithSize)
	procExecuteCallback.Call(cb)
}

func TestV2CodeWithV1LibWithNoSize(t *testing.T) {
	// Safe and correct
	cb := windows.NewCallbackCDecl(callback2WithNoSize)
	procExecuteCallback.Call(cb)
}

// Test v2 code with libv2 library
func TestV2CodeWithV2LibWithSize(t *testing.T) {
	// Safe and correct
	cb := windows.NewCallbackCDecl(callback2WithSize)
	procExecuteCallbackv2.Call(cb)
}

func TestV2CodeWithV2LibWithNoSize(t *testing.T) {
	// Safe and correct
	cb := windows.NewCallbackCDecl(callback2WithNoSize)
	procExecuteCallbackv2.Call(cb)
}

// Test v2 code with libv3 library
func TestV2CodeWithV3LibWithSize(t *testing.T) {
	// Safe and correct but misleading, we set 65535 as size, but size is 131072
	cb := windows.NewCallbackCDecl(callback2WithSize)
	procExecuteCallbackv3.Call(cb)
}

func TestV2CodeWithV3LibWithNoSize(t *testing.T) {
	// Safe and correct
	cb := windows.NewCallbackCDecl(callback2WithNoSize)
	procExecuteCallbackv3.Call(cb)
}
