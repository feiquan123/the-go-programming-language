package main

import "fmt"

// Celsius ℃
type Celsius float64

func (c Celsius) String() string { return fmt.Sprintf("%g℃", c) }

// Fahrenheit ℉
type Fahrenheit float64

func (f Fahrenheit) String() string { return fmt.Sprintf("%g℉", f) }

// ℃
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	fmt.Printf("%v = %v\n", AbsoluteZeroC, CToF(AbsoluteZeroC))
	fmt.Printf("%v = %v\n", FreezingC, CToF(FreezingC))
	fmt.Printf("%v = %v\n", BoilingC, CToF(BoilingC))

	fmt.Printf("%v = %v\n", CToF(AbsoluteZeroC), AbsoluteZeroC)
	fmt.Printf("%v = %v\n", CToF(FreezingC), FreezingC)
	fmt.Printf("%v = %v\n", CToF(BoilingC), BoilingC)

	fmt.Printf("%g\n", BoilingC)   // 不调用字符串
	fmt.Println(float64(BoilingC)) // 不调用字符串
}

// FToC ℉ -> ℃
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5.0 / 9.0)
}

// CToF ℃ -> ℉
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(9.0/5.0*c + 32)
}
