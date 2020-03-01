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
//import (
//	"fmt"
//
//	"convTemp/internal/tempconv"
//)
//
//func main() {
//	fmt.Printf("zZZ %v\n", tempconv.AbsoluteZeroC)
//	fmt.Println(tempconv.CToF(tempconv.AbsoluteZeroC))
//	fmt.Printf("absoult zero in kelvin is %v\n", tempconv.CToK(tempconv.AbsoluteZeroC))
//}

//practice 2.6.1 1
import (
	"fmt"
	"os"
	"strconv"

	"convTemp/internal/popcount"
	"convTemp/internal/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		t1 := uint64(t)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", k, tempconv.KToC(k), c, tempconv.CToK(c))
		fmt.Println(popcount.PopCount(t1))
	}
}
