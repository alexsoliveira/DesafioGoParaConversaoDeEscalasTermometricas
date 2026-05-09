package main

import "fmt"

//declarar uma constante para a temperatura de ebulição da água em Kelvin
const ebulicaoK = 373.15

func main() {
	tempKelvin := ebulicaoK
	tempCelsius := (tempKelvin - 273.15)

	fmt.Printf("A temperatura da ebulicao da agua em Kelvin é: %.2f e a temperatura da ebulicao da agua em Celsius é: %.2f", tempKelvin, tempCelsius)
}
