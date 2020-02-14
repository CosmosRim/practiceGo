package main

//import "fmt"
//
//func main() {
//    fmt.Println("show me the code")
//}

//practice 2.5 1
//import "fmt"
//
//type Celsius float64
//type Fahrenheit float64
//
//const (
//	AbsoluteZeroC Celsius = -273.15
//	FreezingC     Celsius = 0
//	BoilingC      Celsius = 100
//)
//
//func main() {
//	var currentTemp Celsius = 19
//	fmt.Printf("current temperature is %g Celsius, %g Fahrenheit.\n", currentTemp, CToF(currentTemp))
//	fmt.Printf("different type %g.\n", CToF(BoilingC)-CToF(FreezingC))
//	fmt.Println("try merge two")
//}
//
//func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
//func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//practice 2.6 1
import (
	"fmt"

	"./internal/tempconv"
) 

fmt.Println("zZZ %v\n", tempconv.AbsoluteZeroC)
fmt.Println(tempconv.CToF(tempconv.AbsoluteZeroC))