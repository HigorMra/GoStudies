package basic

import "fmt"

//Floats diferentes não conseguem operar juntos, por exemplo um cálculo a partir de float32 e float64. 
func Floatnumber1() {
	//var pi floatNumber float32 = 1.1
	var pi float32 = 3.14
	var raio float32 = 2.5
	var area = pi * raio * raio

	fmt.Println("Area do círculo:", area)
}
