package geometrica

import "fmt"

type figuraGeometrica interface {
	area() float64
	perimetro() float64
}

func Medidas(f figuraGeometrica) {
	fmt.Println(f)
	fmt.Println(f.perimetro())
	fmt.Println(f.area())
}
