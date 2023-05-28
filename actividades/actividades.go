package actividades

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

// Selector de actividades iterativo
// Devuelve un slice con las actividades seleccionadas que no se solapan
// Pre condición: las actividades están ordenadas  de menor a mayor por tiempo de finalización

func SelectorActividadesIterativo(actividades []Actividad) []Actividad { // O(N)
	var seleccionadas []Actividad
	n := len(actividades)
	k := 0
	seleccionadas = append(seleccionadas, actividades[k])

	// 0: 1 a 4
	// 2: 4 a 7
	// 5: 10 a 12

	for i := 1; i < n; i++ {
		if actividades[i].Inicio >= actividades[k].Fin {
			k = i
			seleccionadas = append(seleccionadas, actividades[k])
		}
	}
	return seleccionadas
}
