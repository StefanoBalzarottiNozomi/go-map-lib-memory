package gomaplibmemory

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	lib1Dll             = windows.NewLazyDLL("libs/libv1.dll")
	procExecuteCallback = lib1Dll.NewProc("ExecuteCallback")
)

const NUMBER_ARRAY_SIZE = 4096
const DEFAULT_BUFFER_SIZE = 16384

type VData struct {
	ones   [NUMBER_ARRAY_SIZE]byte
	twos   [NUMBER_ARRAY_SIZE]byte
	threes [NUMBER_ARRAY_SIZE]byte
	fours  [NUMBER_ARRAY_SIZE]byte
}

type DataType uint32

const (
	VDataT      = DataType(1)
	StringDataT = DataType(2)
	NoDataT     = DataType(3)
)

type Data1WithSize struct {
	dataType   DataType
	bufferSize int32
	buffer     [DEFAULT_BUFFER_SIZE]byte
}

type Data1WithNoSize struct {
	dataType   DataType
	bufferSize int32
	buffer     [0]byte
}

func callbackWithSize(data *Data1WithSize) uintptr {
	switch data.dataType {
	case VDataT:
		vdata := (*VData)(unsafe.Pointer(&data.buffer[0]))
		fmt.Printf("VData: %v\n", vdata)
	case StringDataT:
		str := windows.BytePtrToString(&data.buffer[0])
		fmt.Printf("String data: %s\n", str)
	case NoDataT:
		fmt.Println("No data")
	}
	return 0
}

func callbackWithNoSize(data *Data1WithNoSize) uintptr {
	switch data.dataType {
	case VDataT:
		vdata := (*VData)(unsafe.Pointer(&data.buffer))
		fmt.Printf("VData: %v\n", vdata)
	case StringDataT:
		strPtr := (*byte)(unsafe.Pointer(&data.buffer))
		str := windows.BytePtrToString(strPtr)
		fmt.Printf("String data: %s\n", str)
	case NoDataT:
		fmt.Println("No data")
	}
	return 0
}
