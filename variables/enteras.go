package variables

import "fmt"

func MostrarEnteros() {
	var enteroComun int
	entero32 := int32(10)
	entero64 := int64(100)

	fmt.Println("Entero Comun", enteroComun)
	fmt.Println("entero 32", entero32)
	fmt.Println("entero 64", entero64)
}
