package main

import (
	"fmt"

	"github.com/hultan/roman/internal/roman"
)

func main() {
	fmt.Println("845 =", roman.ToRoman(845))

	fmt.Println("DCCCXLV =", roman.ToArabic("DCCCXLV"))
}
