package main

import "fmt"

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	var choice int
	var temperature float64

	fmt.Println("Conversor de Temperaturas")
	fmt.Println("-------------------------")
	fmt.Println("Selecione a operação:")
	fmt.Println("1 - Celsius para Fahrenheit")
	fmt.Println("2 - Fahrenheit para Celsius")
	fmt.Print("Escolha: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Print("Digite a temperatura em Celsius: ")
		fmt.Scanln(&temperature)
		converted := celsiusToFahrenheit(temperature)
		fmt.Printf("%.2f °C equivale a %.2f °F\n", temperature, converted)
	case 2:
		fmt.Print("Digite a temperatura em Fahrenheit: ")
		fmt.Scanln(&temperature)
		converted := fahrenheitToCelsius(temperature)
		fmt.Printf("%.2f °F equivale a %.2f °C\n", temperature, converted)
	default:
		fmt.Println("Escolha inválida.")
	}

	fmt.Println("-------------------------")
}
