package main

//import "fmt"
//
//func main() {
//    fmt.Println("show me the code")
//}

//practice 2.5 1
import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	var currentTemp Celsius = 19
	fmt.Printf("current temperature is %g Celsius, %g Fahrenheit.", currentTemp, CToF(currentTemp))
}

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
