package ejercicios

import "errors"

type Farolas struct {
	Posicion  int
	Radio     int
	Encendida bool
}

// Farolas públicas
// Hay M farolas en las posiciones y1,…,yM de una recta y N puntos x1,…,xN.
// Cada farola tiene un radio de iluminación ri, tal que la i-ésima farola ilumina puntos en
// el intervalo [yi−ri,yi+ri]. Se quiere encender el mínimo número de farolas tales que cada
// uno de los N puntos x1,…xN esté iluminado por al menos una farola.
// Encuentra este mínimo número.
// Entrada: un arreglo de Farolas de tamaño M y
// un arreglo de enteros de tamaño N con los puntos a iluminar.
// Salida: un arreglo de farolas encendidas
// Pre condiciones:
// El arreglo de farolas y puntos está ordenado por posicion de menor a mayor.
// Paso greedy: Dado un punto x, encender la farola más a la derecha que pueda iluminarlo.
func EncenderFarolas(farolas []Farolas, x []int) ([]Farolas, error) {
	encendidas := make([]Farolas, 0)

	for _, punto := range x {
		// Buscar la farola más a la derecha que pueda iluminar el punto actual
		farolaIdx := -1

		for i := len(farolas) - 1; i >= 0; i-- {
			if farolas[i].Posicion-farolas[i].Radio <= punto {
				farolaIdx = i
				break
			}
		}

		if farolaIdx == -1 {
			// No se encontró una farola que pueda iluminar el punto actual
			return nil, errors.New("No es posible iluminar todos los puntos")
		}

		// Encender la farola seleccionada si aún no está encendida
		if !farolas[farolaIdx].Encendida {
			farolas[farolaIdx].Encendida = true
			encendidas = append(encendidas, farolas[farolaIdx])
		}
	}

	return encendidas, nil
}
