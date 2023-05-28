package ejercicios

import "fmt"

type Objeto struct {
	Nombre string
	Peso   int
	Valor  int
}

func Ejercicio3(objetos []Objeto, capacidad int) map[Objeto]float32 {
	// Ordenar los objetos por valor/peso usando burbujeo
	for i := 0; i < len(objetos)-1; i++ {
		for j := 0; j < len(objetos)-i-1; j++ {
			if float32(objetos[j].Valor)/float32(objetos[j].Peso) < float32(objetos[j+1].Valor)/float32(objetos[j+1].Peso) {
				objetos[j], objetos[j+1] = objetos[j+1], objetos[j]
			}
		}
	}
	mochilaCargada := make(map[Objeto]float32)
	pesoActual := 0
	valorTotal := 0
	// Agrego el objeto completo sin que sobrepase la capacidad o fraccion del objeto
	for _, objeto := range objetos {
		if pesoActual+objeto.Peso <= capacidad {
			pesoActual += objeto.Peso
			valorTotal += objeto.Valor
			mochilaCargada[objeto] = 1.0
		} else {
			fraccion := float32(capacidad-pesoActual) / float32(objeto.Peso)
			pesoActual += int(fraccion * float32(objeto.Peso))
			valorTotal += int(fraccion * float32(objeto.Valor))
			mochilaCargada[objeto] = fraccion
			break
		}
	}
	fmt.Printf("Valor total: %d\n", valorTotal)
	fmt.Printf("Peso total: %d\n", pesoActual)

	return mochilaCargada
}
