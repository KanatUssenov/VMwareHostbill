package main

import "fmt"

func main() {
	var sliceAcc []string = hbAccTaker()
	for key, value := range hbAccDetTaker(sliceAcc) {
		fmt.Println(key, value)
	}
	defer vmwposreq()
}
