package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer1, buffer2 bytes.Buffer
	var st_temp string
	b_temp := make([]byte, 30)

	for i := 0; i < 200; i++ {
		st_temp = fmt.Sprintf("%04d", i) + "-1234567890A"
		if _, err := buffer1.WriteString(st_temp); err != nil {
			break
		}
		if _, err := buffer2.ReadFrom(&buffer1); err != nil {
			break
		}
		buffer1.Reset()
	}

	//	buffer1.Write([]byte("Hello World (Buffer)\n")) // 溜め
	//	fmt.Println(buffer2.String())

	//	63-123456789  64-123456789
	var idx int = 0
	for {
		idx = idx + 1
		if _, err := buffer2.Read(b_temp); err != nil {
			fmt.Printf("Errored %s\n", string(b_temp))
			break
		}
		fmt.Printf("No.%03d %s\n", idx, string(b_temp))
		for j := 0; j < len(b_temp); j++ {
			b_temp[j] = 0x00
		}
	}

	fmt.Printf("No.%03d %s\n", idx, string(b_temp))

}
