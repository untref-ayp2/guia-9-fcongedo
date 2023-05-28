package ejercicios

type Tarea struct {
	Nombre string
	Tiempo int
}

//Ejemplo:

// t = [4, 10, 2, 20] Si las ejecuto en el orden a1, a2, a3, a4 entonces cada una de las tareas se ejecutarán en los siguientes
// intervamos de tiempo a1: [0, 4), a2: [4, 14), a3:[14, 16), a4:[16, 36)
// y el promedio de los tiempos de finalización es (4 + 14 + 16 + 36) / 4 = 17,5.
// En cambio si se ejecutan a3, a1, a2, a4 entonces el cada una de las tareas se ejecutarán
// en los siguientes intervalos de tiempo a1: [2, 6), a2: [6, 16), a3:[0, 2), a4:[16, 36)
// y el promedio de los tiempos de finalización es (6 + 16 + 2 + 36) / 4 = 15

func Ejercicio2(tareas []Tarea) { // O N**2
	//Ordenar las tareas por tiempo usando burbujeo
	for i := 0; i < len(tareas)-1; i++ {
		for j := 0; j < len(tareas)-1-i; j++ {
			if tareas[j].Tiempo > tareas[j+1].Tiempo {
				tareas[j], tareas[j+1] = tareas[j+1], tareas[j]
				//aux := tareas[j]
				//tareas[j] = tareas[j+1] (explicacion de la línea 21)
				//tareas[j+1] = aux
			}
		}
	}
}
