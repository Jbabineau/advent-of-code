package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	data := "yzbqklnj"
	counter := 0
	for {
		hash := md5.Sum([]byte(data + strconv.Itoa(counter)))
		hashString := hex.EncodeToString(hash[:])
		//fmt.Println(counter, hashString)
		if hashString[0:6] == "000000" {
			break
		}

		counter++
	}
	fmt.Println(counter)
}
