package main

import (
	"fmt"
	"strings"
)

func main() {

	if strings.ContainsAny("d'", "!@#$%^&*()'\"|{}[];:><?/`~,") {
		fmt.Println("kk")
	}

}
