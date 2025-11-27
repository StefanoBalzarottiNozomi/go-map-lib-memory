package gomaplibmemory

import (
	"testing"

	"golang.org/x/sys/windows"
)

// Test v1 code with libv1 library
func TestV1CodeWithV1LibWithSize(t *testing.T) {
	// Safe and correct
	cb := windows.NewCallbackCDecl(callbackWithSize)
	procExecuteCallback.Call(cb)
}

func TestV1CodeWithV1LibWithNoSize(t *testing.T) {
	// Safe and correct
	cb := windows.NewCallbackCDecl(callbackWithNoSize)
	procExecuteCallback.Call(cb)
}

// Test v1 code with libv2 library

func TestV1CodeWithV2LibWithSize(t *testing.T) {
	// Safe and correct, but misleading, we set 16384 as size, but size is 65536
	cb := windows.NewCallbackCDecl(callbackWithSize)
	procExecuteCallbackv2.Call(cb)
}

func TestV1CodeWithV2LibWithNoSize(t *testing.T) {
	// Safe and correct
	cb := windows.NewCallbackCDecl(callbackWithNoSize)
	procExecuteCallbackv2.Call(cb)
}
