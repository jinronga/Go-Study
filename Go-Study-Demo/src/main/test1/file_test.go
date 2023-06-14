package test1

import "os"

func readFileValue() {

	buffer := make([]byte, 1024)

	f, error := os.Open("C:\\Users\\14605\\Desktop\\temp\\随时记录04.txt")
	defer f.Close()
	error.Error()
	for {

		n, _ := f.Read(buffer)
		if n == 0 {
			break
		}
		os.Stdout.Write(buffer[:n])
	}

}

func init() {
	readFileValue()

}

var name string
var age int
var isOk bool
