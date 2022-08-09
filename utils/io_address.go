package utils

import "fmt"

var digitalMemory int = 0
var analogMemory int = 100

func NextDigital() string {
	address := fmt.Sprintf("M%d.%d", digitalMemory/8, digitalMemory%8)
	digitalMemory++
	return address
}

func NextAnalog() string {
	address := fmt.Sprintf("MW%d", analogMemory)
	analogMemory += 2
	return address
}
