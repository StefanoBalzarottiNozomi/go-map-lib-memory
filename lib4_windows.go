package gomaplibmemory

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	lib4Dll                 = windows.NewLazyDLL("libs/libv4.dll")
	procExecuteCallbackv4   = lib4Dll.NewProc("ExecuteCallback")
	procExecuteCallbackv4v2 = lib4Dll.NewProc("ExecuteCallbackv2")
)

const ANY_SIZE = 1 // this is how is defined in windows

type Data4WithSize struct {
	dataType   DataType
	bufferSize int32
	buffer     [ANY_SIZE]byte
}

func callback4WithSize(data *Data4WithSize) uintptr {
	switch data.dataType {
	case VDataT:
		vdata := (*VData)(unsafe.Pointer(&data.buffer[0]))
		fmt.Printf("VData: %v\n", vdata)
	case VData2T:
		vdata := (*VData2)(unsafe.Pointer(&data.buffer[0]))
		fmt.Printf("VData2: %v\n", vdata)
	case VData3T:
		vdata := (*VData3)(unsafe.Pointer(&data.buffer[0]))
		fmt.Printf("VData3: %v\n", vdata)
	case StringDataT:
		str := windows.BytePtrToString(&data.buffer[0])
		fmt.Printf("String data: %s\n", str)
	case NoDataT:
		fmt.Println("No data")
	}
	return 0
}

func callback4WithNoSize(data *Data1WithNoSize) uintptr {
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
	case VData3T:
		vdata := (*VData3)(unsafe.Pointer(&data.buffer))
		fmt.Printf("VData3: %v\n", vdata)
	case NoDataT:
		fmt.Println("No data")
	}
	return 0
}
