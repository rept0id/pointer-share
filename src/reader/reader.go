package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
	"unsafe"
)

func main() {
	var err error

	// pid

	// pid : input
	var inpPid string

	// pid : pid
	var pid string

	// memory

	// memory : address

	// memory : address : input
	var inpAddrMem string
	// memory : address : address
	var addrMem uint64

	// memory : file

	// memory : file : path
	var pathFileMem string

	// memory : file : value
	var valFileMem *os.File

	// value

	// value : bytes

	var bytesVal []byte

	// value : value
	var val int

	/*** * * ***/

	if len(os.Args) != 3 {
		fmt.Println("Usage: reader <pid> <address>")
		return
	}

	/*** * * ***/

	// pid

	// pid : input
	inpPid = os.Args[1]

	// pid : pid
	pid = inpPid

	// memory

	// memory : address

	// memory : address : input
	inpAddrMem = os.Args[2]

	// memory : address : address
	addrMem, err = strconv.ParseUint(inpAddrMem, 10, 64)
	if err != nil {
		fmt.Println("Invalid address:", err)
		return
	}

	// memory : file

	// memory : file : path
	pathFileMem = fmt.Sprintf("/proc/%s/mem", pid)

	// memory : file : value
	valFileMem, err = os.Open(pathFileMem)
	if err != nil {
		fmt.Println("Error opening memory file:", err)
		return
	}
	defer valFileMem.Close()

	// memory : file : seek addrMem
	_, err = valFileMem.Seek(int64(addrMem), 0)
	if err != nil {
		fmt.Println("Error seeking to address:", err)
		return
	}

	// val

	// val : bytes

	// val : bytes : make
	bytesVal = make([]byte, unsafe.Sizeof(int(0)))

	// val : bytes : read
	_, err = syscall.Pread(int(valFileMem.Fd()), bytesVal, int64(addrMem))
	if err != nil {
		fmt.Println("Error reading memory:", err)
		return
	}

	// val : val
	val = *(*int)(unsafe.Pointer(&bytesVal[0]))

	/*** * * ***/

	fmt.Printf("Value at address %d: %d\n", addrMem, val)
}
