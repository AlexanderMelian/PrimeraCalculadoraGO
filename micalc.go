package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func stringToInt(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func encontrarOperador(entrada string) string {
	operadores := []string{"+", "-", "*", "/"}
	for _, v := range operadores {
		if strings.Contains(entrada, v) {
			return v
		}
	}
	return "p"
}

func calcular(entrada string) int {
	operador := encontrarOperador(entrada)
	if operador == "p" {
		fmt.Println("Operador desconocido")
	} else {
		entradaDividida := strings.Split(entrada, operador)
		operando1 := stringToInt(entradaDividida[0])
		operando2 := stringToInt(entradaDividida[1])
		switch operador {
		case "+":
			return operando1 + operando2
		case "-":
			return operando1 - operando2
		case "*":
			return operando1 * operando2
		case "/":
			return operando1 / operando2
		}
	}
	return 0
}

func main() {
	fmt.Println("Calculadora barata y fea")
	entradaCompleta := leerEntrada()
	resultado := calcular(entradaCompleta)
	fmt.Println(resultado)
}
