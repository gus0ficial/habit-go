/*
* @Author: Gustav0
* @Date:   2026-09-16 22:04:26
* @Last Modified by:   Gustav0
* @Last Modified time: 2026-09-16 22:11:13
 */
package habits

import (
	"bufio"
	"encoding/json"
	"os"
	"time"
)

// Funciones de los Habitos
func NuevoHabito(habitos []Habit, nombre string, descripcion string, porque string, paraque string, como string, idUsuario int) []Habit {
	habitoNuevo := Habit{
		ID:          SiguienteID(habitos),
		Nombre:      nombre,
		Descripcion: descripcion,
		Porque:      porque,
		Paraque:     paraque,
		Como:        como,
		IDUsuario:   idUsuario,
		FRegistro:   time.Now().Format("2006-01-02"),
		Status:      true,
	}
	/*
	   package habits

	   const (
	   	TipoBooleano    = 1
	   	TipoConteo      = 2
	   	TipoMetaSemanal = 3
	   )

	   type Tag struct {
	   	ID     int    `json:"id"`
	   	Nombre string `json:"nombre"`
	   }

	   type Theme struct {
	   	ID     int    `json:"id"`
	   	Nombre string `json:"nombre"`
	   }

	   type Action struct {
	   	ID          int    `json:"id"`
	   	Nombre      string `json:"nombre"`
	   	Cantidad    string `json:"cantidad"`
	   	Descripcion string `json:"descripcion"`
	   	Duracion    int    `json:"duracion"`
	   }

	   type Routine struct {
	   	ID            int      `json:"id"`
	   	Nombre        string   `json:"nombre"`
	   	Utilidad      string   `json:"utilidad"`
	   	DuracionTotal int      `json:"duracion_total"`
	   	Acciones      []Action `json:"acciones"`
	   	Tags          []Tag    `json:"etiquetas"`
	   }

	   type Habit struct {
	   	ID          int       `json:"id"`
	   	Nombre      string    `json:"nombre"`
	   	Tipo        int       `json:"tipo"`
	   	Descripcion string    `json:"descripcion"`
	   	Porque      string    `json:"porque"`
	   	Paraque     string    `json:"paraque"`
	   	Como        string    `json:"como"`
	   	Color       string    `json:"color"`
	   	Meta        int       `json:"meta"`
	   	Rutinas     []Routine `json:"rutinas"`
	   	Tematicas   []Theme   `json:"tematicas"`
	   	IDUsuario   int       `json:"id_usuario"`
	   	FRegistro   string    `json:"f_registro"`
	   	Status      bool      `json:"status"`
	   }

	   type HabitLog struct {
	   	IDHabito   int    `json:"id_habito"`
	   	Fecha      string `json:"fecha"`
	   	Completado bool   `json:"completado"`
	   	Valor      int    `json:"valor"`
	   }

	*/
	return append(habitos, habitoNuevo)
}

func GuardarHabito(file *os.File, habitos []Habit) {
	bytes, err := json.MarshalIndent(habitos, "", "  ")
	if err != nil {
		panic(err)
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}

func SiguienteID(habitos []Habit) int {
	if len(habitos) == 0 {
		return 1
	}
	return habitos[len(habitos)-1].ID + 1
}

func EvitarDuplicados(valor string, idUsuario int, listaHabitos []Habit) bool {
	for _, habito := range listaHabitos {
		if habito.Nombre == valor && habito.IDUsuario == idUsuario {
			return true
		}
	}
	return false
}

func DeshabilitarHabito(idHabito int, idUsuario int, listaHabitos []Habit) []Habit {
	for i, habito := range listaHabitos {
		if habito.ID == idHabito && habito.IDUsuario == idUsuario {
			listaHabitos[i].Status = false
			break
		}
	}
	return listaHabitos
}

func ModificarHabito(idHabito int, idUsuario int, listaHabitos []Habit, nombreNuevo string, descripcionNuevo string, porqueNuevo string, paraqueNuevo string, comoNuevo string) []Habit {
	for i, habito := range listaHabitos {
		if habito.ID == idHabito && habito.IDUsuario == idUsuario {
			listaHabitos[i].Nombre = nombreNuevo
			listaHabitos[i].Descripcion = descripcionNuevo
			listaHabitos[i].Porque = porqueNuevo
			listaHabitos[i].Paraque = paraqueNuevo
			listaHabitos[i].Como = comoNuevo
			break
		}
	}
	return listaHabitos
}

// Funciones de las Rutinas

/*


Plan de Desarrollo

Métodos:
- Mostrar Uso: Hecho.

- Agregar Hábitos: Hecho
- Quitar Hábitos: Hecho
- Modificar Hábitos: Hecho
- Mostrar Hábitos:

- Mostrar Estado de hoy:
- Marcar:
- Desmarcar:

- Mostrar Seguimiento Semanal:
- Mostrar Seguimiento Mensual:
- Mostrar Seguimiento Anual:
*/
