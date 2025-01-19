package variables

import "fmt"

func MuestroEnteros() {
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("int comun = ", intcomun)
	fmt.Println("int de 32 = ", intde32)
	fmt.Println("int de 64 = ", intde64)
}
