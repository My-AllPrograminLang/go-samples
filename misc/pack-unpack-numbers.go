package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	buf := new(bytes.Buffer)
	var num uint32 = 256*256 + 256 + 254
	err := binary.Write(buf, binary.LittleEndian, num)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("%X packed to - % 0X\n", num, buf.Bytes())

	var unpacked_num uint32
	err = binary.Read(buf, binary.LittleEndian, &unpacked_num)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Printf("And unpacks back to %X\n", unpacked_num)
}
