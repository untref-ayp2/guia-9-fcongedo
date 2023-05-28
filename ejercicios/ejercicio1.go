package ejercicios

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

// Devuelve un slice con las actividades seleccionadas que no se solapan
// Pre condición: las actividades están ordenadas de menor a mayor por tiempo de finalización

func SelectorRecursivo(actividades []Actividad) []Actividad {
	resultado := []Actividad{}
	aux := []Actividad{}

	if len(actividades) == 0 {
		return aux
	}
	// Seleccionar la actividad de menor fin
	resultado = append(resultado, actividades[0])

	// Eliminar las actividadess que se solapan con la seleccionada
	for i := 1; i < len(actividades); i++ {
		if actividades[i].Inicio >= actividades[0].Fin {
			aux = append(aux, actividades[i])
		}
	}
	// Llamar a la función recursivamente
	resultado = append(resultado, SelectorRecursivo(aux)...)
	return resultado
}
