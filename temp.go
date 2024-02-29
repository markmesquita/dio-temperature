package main

import "fmt"

const ebulicaoK = 373.0

func main() {
	var tempK float64 = ebulicaoK
	var tempC float64 = tempK - 273

	fmt.Printf("A ebulição da água em ºK é %g \ne ebulição da água em ºC é = %g .", tempK, tempC)
}