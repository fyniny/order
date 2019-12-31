package data

import (
	"fmt"
	"testing"
)

func TestData(t *testing.T) {
	o, d := RandomArrayInt(4)
	fmt.Println(o)
	fmt.Println(d)
}