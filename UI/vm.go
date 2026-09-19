package UI

var (
	vinotinto = "#841A1A"
	marino    = "#08273B"
	olivo     = "#3F461A"
	naranja   = "#A5340D"

	negro  = "#000000"
	blanco = "#F3F4F6"
	gris   = "#EEEFF1"

	uno    = vinotinto
	dos    = marino
	tres   = olivo
	cuatro = naranja

	estiloListaItem = lipGloss.NewStyle().
			Align(lipGloss.Left)
	estiloTitulo = lipGloss.NewStyle().
			Bold(true).
			Align(lipGloss.Center)
	estiloTextarea = lipGloss.NewStyle().
			PaddingTop(2)
	estiloInput = lipGloss.NewStyle().
			PaddingTop(2)
	estiloActivo = lipGloss.NewStyle().
			Foreground(lipGloss.Color(cuatro)).
			Bold(true)
)

var opciones = []string{
	"Habitos",
	"Perfil",
	"Salir",
}
