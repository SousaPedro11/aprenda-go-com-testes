package formas

import "math"

type Forma interface {
	Area() float64
	Perimetro() float64
}

type Retangulo struct {
	Altura float64
	Base   float64
}

type Circulo struct {
	Raio float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

func (retangulo Retangulo) Perimetro() float64 {
	return 2 * (retangulo.Base + retangulo.Altura)
}

func (circulo Circulo) Perimetro() float64 {
	return 2 * math.Pi * circulo.Raio
}

func (triangulo Triangulo) Perimetro() float64 {
	hipotenusa := math.Sqrt(math.Pow(triangulo.Base, 2) + math.Pow(triangulo.Altura, 2))
	return triangulo.Base + triangulo.Altura + hipotenusa
}

func (retangulo Retangulo) Area() float64 {
	return retangulo.Base * retangulo.Altura
}

func (circulo Circulo) Area() float64 {
	return math.Pi * math.Pow(circulo.Raio, 2)
}

func (triangulo Triangulo) Area() float64 {
	return (triangulo.Base * triangulo.Altura) / 2
}
