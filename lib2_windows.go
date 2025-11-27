package gomaplibmemory

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	lib2Dll               = windows.NewLazyDLL("libs/libv2.dll")
	procExecuteCallbackv2 = lib2Dll.NewProc("ExecuteCallback")
)

const NUMBER_ARRAY_SIZE2 = 8192
const MAX_BUFFER_SIZE = 65535

type VData2 struct {
	ones   [NUMBER_ARRAY_SIZE]byte
	twos   [NUMBER_ARRAY_SIZE]byte
	threes [NUMBER_ARRAY_SIZE]byte
	fours  [NUMBER_ARRAY_SIZE]byte
	fives  [NUMBER_ARRAY_SIZE2]byte
	sixes  [NUMBER_ARRAY_SIZE2]byte
	sevens [NUMBER_ARRAY_SIZE2]byte
	eights [NUMBER_ARRAY_SIZE2]byte
}

const (
	VData2T = DataType(4)
)

type Data2WithSize struct {
	dataType   DataType
	bufferSize int32
	buffer     [MAX_BUFFER_SIZE]byte
}

func callback2WithSize(data *Data2WithSize) uintptr {
	switch data.dataType {
	case VDataT:
		vdata := (*VData)(unsafe.Pointer(&data.buffer[0]))
		fmt.Printf("VData: %v\n", vdata)
	case VData2T:
		vdata := (*VData2)(unsafe.Pointer(&data.buffer[0]))
		fmt.Printf("VData2: %v\n", vdata)
	case StringDataT:
		str := windows.BytePtrToString(&data.buffer[0])
		fmt.Printf("String data: %s\n", str)
	case NoDataT:
		fmt.Println("No data")
	}
	return 0
}

func callback2WithNoSize(data *Data1WithNoSize) uintptr {
	switch data.dataType {
	case VDataT:
		vdata := (*VData)(unsafe.Pointer(&data.buffer))
		fmt.Printf("VData: %v\n", vdata)
	case StringDataT:
		strPtr := (*byte)(unsafe.Pointer(&data.buffer))
		str := windows.BytePtrToString(strPtr)
		fmt.Printf("String data: %s\n", str)
	case VData2T:
		vdata := (*VData2)(unsafe.Pointer(&data.buffer))
		fmt.Printf("VData2: %v\n", vdata)
	case NoDataT:
		fmt.Println("No data")
	}
	return 0
}
