package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"syscall"
)

func main() {

	if runtime.GOOS == "windows" {
		const dir = "c://isexists_test"
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			log.Fatal(err)
		}

		err = os.MkdirAll(dir, 0777)

		if err == nil {
			log.Fatal("should fail")
		}

		exist := os.IsExist(err)

		var num syscall.Errno
		if errno, ok := err.(syscall.Errno); ok {
			num = errno
		}

		fmt.Printf("%s: Exists: %t Errno: %d", err, exist, num)
	}

}
