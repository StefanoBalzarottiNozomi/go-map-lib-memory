package gomaplibmemory

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	lib3Dll               = windows.NewLazyDLL("libs/libv3.dll")
	procExecuteCallbackv3 = lib3Dll.NewProc("ExecuteCallback")
)

const NUMBER_ARRAY_SIZE3 = 16384
const NEW_MAX_BUFFER_SIZE = 131072

type VData3 struct {
	ones   [NUMBER_ARRAY_SIZE]byte
	twos   [NUMBER_ARRAY_SIZE]byte
	threes [NUMBER_ARRAY_SIZE]byte
	fours  [NUMBER_ARRAY_SIZE]byte
	fives  [NUMBER_ARRAY_SIZE2]byte
	sixes  [NUMBER_ARRAY_SIZE2]byte
	eights [NUMBER_ARRAY_SIZE2]byte
	nines  [NUMBER_ARRAY_SIZE3]byte
	tens   [NUMBER_ARRAY_SIZE3]byte
}

const (
	VData3T = DataType(5)
)

type Data3WithSize struct {
	dataType   DataType
	bufferSize int32
	buffer     [NEW_MAX_BUFFER_SIZE]byte
}

func callback3WithSize(data *Data3WithSize) uintptr {
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

func callback3WithNoSize(data *Data1WithNoSize) uintptr {
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
