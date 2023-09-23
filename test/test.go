package main

import (
	"fmt"
)

func main() {
	employees := make(map[string]string)

	employees["Josn"] = "Software Developer"
	employees["James"] = "Fisher"
	employees["Ariana"] = "Management Consultant"

	fmt.Println(employees)
}