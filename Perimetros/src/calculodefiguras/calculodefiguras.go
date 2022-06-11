package calculodefiguras

import (
	"errors"
	"fmt"
)

type Figuras interface {
	perimetro() float64
	area() float64
}

type Cuadrado struct {
	base float64
}

type Rectangulo struct {
	base   float64
	altura float64
}

type Triangulo struct {
	lado   float64
	altura float64
}

func (c *Cuadrado) area() float64 {
	return c.base * c.base
}

func (r *Rectangulo) area() float64 {
	return r.altura * r.base
}

func (t *Triangulo) area() float64 {
	return t.lado * t.altura / 2
}

func (c Cuadrado) perimetro() float64 {
	return c.base * 4
}

func (r *Rectangulo) perimetro() float64 {
	return r.base*2 + r.altura*2
}

func (t *Triangulo) perimetro() float64 {
	return t.lado * 3
}

func printArea(a Figuras) {
	fmt.Println("Area: ", a.area())
}

func printPerimetro(p Figuras) {
	fmt.Println("Perímetro: ", p.perimetro())
}

func Cal() {
	ask :=
		`
	Hola:), ingresa que quieres hacer:
	1. Calculo de perimetro
	2. Calculo de area
	`

	fmt.Print(ask)

	var eleccion string
	fmt.Scanln(&eleccion)

	err := errors.New("Syntax Error *.*")

	if eleccion != "1" && eleccion != "2" {
		fmt.Println(err)
	} else {

		if eleccion == "1" {
			ask :=
				`
			De que figura quieres el perimetro?:
			1. Cuadrado
			2. Rectangulo
			3. Triangulo
			`

			fmt.Print(ask)

			var eleccion string
			fmt.Scanln(&eleccion)

			if eleccion == "1" {
				ask :=
					`
			Ingresa la base
			`

				fmt.Print(ask)

				var eleccion float64
				fmt.Scanln(&eleccion)

				c := Cuadrado{base: eleccion}
				printPerimetro(&c)
			}
			if eleccion == "2" {

				var b, h float64

				bask :=
					`
				Ingresa la base
				`
				fmt.Print(bask)

				fmt.Scanln(&b)

				hask := `
				Ingresa la altura
				`
				fmt.Print(hask)
				fmt.Scanln(&h)

				r := Rectangulo{base: b, altura: h}
				printPerimetro(&r)
			}

			if eleccion == "3" {

				var b float64

				bask :=
					`
				Ingresa la base
				`
				fmt.Print(bask)

				fmt.Scanln(&b)

				t := Triangulo{lado: b}
				printPerimetro(&t)
			}

		}
		if eleccion == "2" {
			ask :=
				`
			De que figura quieres el area?:
			1. Cuadrado
			2. Rectangulo
			3. Triangulo
			`

			fmt.Print(ask)

			var eleccion string
			fmt.Scanln(&eleccion)

			if eleccion == "1" {
				ask :=
					`
			Ingresa la base
			`

				fmt.Print(ask)

				var eleccion float64
				fmt.Scanln(&eleccion)

				c := Cuadrado{base: eleccion}
				printArea(&c)
			}
			if eleccion == "2" {

				var b, h float64

				bask :=
					`
				Ingresa la base
				`
				fmt.Print(bask)

				fmt.Scanln(&b)

				hask := `
				Ingresa la altura
				`
				fmt.Print(hask)
				fmt.Scanln(&h)

				r := Rectangulo{base: b, altura: h}
				printArea(&r)
			}

			if eleccion == "3" {

				var b, h float64

				bask :=
					`
				Ingresa la base
				`

				fmt.Print(bask)

				fmt.Scanln(&b)

				hask :=
					`
				Ingresa la altura
				`
				fmt.Print(hask)

				fmt.Scanln(&h)

				t := Triangulo{lado: b, altura: h}
				printArea(&t)
			}
		}
	}

}
