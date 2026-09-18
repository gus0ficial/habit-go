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
	TagsIDs       []int    `json:"tag_ids"`
}

type Habit struct {
	ID           int       `json:"id"`
	Nombre       string    `json:"nombre"`
	Tipo         int       `json:"tipo"`
	Descripcion  string    `json:"descripcion"`
	Porque       string    `json:"porque"`
	Paraque      string    `json:"paraque"`
	Como         string    `json:"como"`
	Color        string    `json:"color"`
	Meta         int       `json:"meta"`
	Rutinas      []Routine `json:"rutinas"`
	TematicasIDs []int     `json:"tematicas_ids"`
	IDUsuario    int       `json:"id_usuario"`
	FRegistro    string    `json:"f_registro"`
	Status       bool      `json:"status"`
}

type HabitLog struct {
	IDHabito      int    `json:"id_habito"`
	Fecha         string `json:"fecha"`
	Completado    bool   `json:"completado"`
	Valor         int    `json:"valor"`
	RutinasHechas []int  `json:"rutinas_hechas"`
}
