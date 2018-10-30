package main

// #include <reader.h>
import "C"
import (
	"io/ioutil"
	"os"
	"unsafe"
)

func main() {
	var buffer = C.init_buffer()
	bytes, _ := ioutil.ReadFile("reader.h.gch")

	f, _ := os.Create("reader.out")
	for len(bytes) != 0 {
		var write = int(C.buffer_append(&buffer, (*C.uchar)(unsafe.Pointer(&bytes[0])), C.int(len(bytes))))
		bytes = bytes[write:]
		for {
			var bytes = make([]byte, 1024)
			var read = int(C.buffer_read(&buffer, (*C.uchar)(unsafe.Pointer(&bytes[0])), C.int(len(bytes))))
			if read == 0 {
				break
			}
			f.Write(bytes[:read])
		}
	}
	f.Close()
}
