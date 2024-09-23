package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"unsafe"
)

func main() {
	var err error

	var num int

	var valMem *int

	var addrMem uintptr

	/*** * * ***/

	if len(os.Args) != 2 {
		fmt.Println("Usage: writer <num>")
		return
	}

	/*** * * ***/

	num, err = strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println(err.Error())
		return
	}

	valMem = new(int)
	*valMem = num

	addrMem = uintptr(unsafe.Pointer(valMem))

	/*** * * ***/

	fmt.Printf("Memory address: %d\n", addrMem)
	fmt.Printf("PID: %d\n", os.Getpid())

	/*** * * ***/

	// Keep the program running
	for {
		time.Sleep(1 * time.Second)
	}
}
