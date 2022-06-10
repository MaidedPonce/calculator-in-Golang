package calculadora

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calculadora struct {
	output float64
}

func (out *Calculadora) obtenerDatos(o string, p string) {

	values := strings.Split(o, p)
	x := []int{}
	s := 0
	for i := 0; i < len(values); i++ {
		newValues, _ := strconv.Atoi(values[i])
		x = append(x, newValues)
	}
	err := errors.New("Así no!")

	if len(x) == 1 {
		fmt.Println(err)
	} else {
		if p == "+" {
			for _, v := range x {
				s += v
			}
			out.output = float64(s)
			fmt.Println(out.output)
		}
		if p == "-" {
			for _, v := range x {

				if s == 0 {
					s = x[0]
				}
				s -= v
			}
			out.output = float64(s)
			fmt.Println(out.output)
		}
		if p == "*" {
			for _, v := range x {

				if s == 0 {
					s = x[0]
				} else {
					s *= v
				}
			}
			out.output = float64(s)
			fmt.Println(out.output)
		}
		if p == "/" {
			for _, v := range x {

				if s == 0 {
					s = x[0]
				} else {
					s /= v
				}
			}
			out.output = float64(s)
			fmt.Println(out.output)
		}
	}

}

func (out *Calculadora) sumar(operacion string) {

	o := &operacion

	out.obtenerDatos(*o, "+")
}

func (out *Calculadora) restar(operacion string) {

	o := &operacion

	out.obtenerDatos(*o, "-")

}

func (out *Calculadora) multiplicar(operacion string) {

	o := &operacion

	out.obtenerDatos(*o, "*")

}

func (out *Calculadora) dividir(operacion string) {

	o := &operacion

	out.obtenerDatos(*o, "/")

}

func Calcular() {

	ask :=
		`
	Hola:), ingresa que operacion haras:
	1. Suma
	2. Res
	3. Multiplicacion
	4. Division
	`

	fmt.Print(ask)

	var eleccion string
	fmt.Scanln(&eleccion)

	out := Calculadora{}
	err := errors.New("Así no!")

	if eleccion > "4" || len(eleccion) > 2 {
		fmt.Println(err)
	} else {
		fmt.Print("Ingresa tu operacion con el siguiente formato. Ejemplo: 1+1+1 \n")

		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		operacion := scan.Text()

		if eleccion == "1" {
			out.sumar(operacion)
		}
		if eleccion == "2" {
			out.restar(operacion)
		}
		if eleccion == "3" {
			out.multiplicar(operacion)
		}
		if eleccion == "4" {
			out.dividir(operacion)
		}
	}

}
