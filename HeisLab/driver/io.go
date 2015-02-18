package driver
/*
#cgo CFLAGS: -std=c99
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
*/
import "C"
import "fmt"

func Init() bool {
	return bool(int(C.io_init()) != 1)
}

func Set_bit(channel int) {
	C.io_set_bit(C.int(channel))
}

func Clear_bit(channel int) {
	C.io_clear_bit(C.int(channel))
}

func Write_analog(channel, value int) {
	C.io_write_analog(C.int(channel), C.int(value))
}

func Read_bit(channel int) bool {
	temp := int(C.io_read_bit(C.int(channel)))
	fmt.Println("temp:",temp)
	return temp != 0
}

func Read_analog(channel int) int {
	return int(C.io_read_analog(C.int(channel)))
}
