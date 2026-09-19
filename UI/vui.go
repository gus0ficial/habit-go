package UI

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lipGloss "github.com/charmbracelet/lipgloss"
	habitPackage "habit-go/habits"
)

type UI struct {
	IdUsuario int
	Paso      int
	Cursor    int
	Entrada   textinput.Model
	Habitos   []habitPackage.Habit

	Altura int
	Ancho  int
}

func (t UI) View() string {
	var contenido string

	if t.Paso == 0 {
		contenido = estiloTitulo.Render("------ MENÚ PRINCIPAL ------")
		contenido += "\n\n"
		for i, op := range opciones {
			prefijo := "  "
			estilo := estiloListaItem
			if t.Cursor == i {
				prefijo = "> "
				estilo = estiloActivo
			}
			contenido += estilo.Render(prefijo+op) + "\n"
		}
		contenido += "\nUsa ↑/↓ para moverte y Enter para seleccionar."
	} else if t.Paso == 1 {
		contenido = estiloTitulo.Render("------ SECCIÓN DE HÁBITOS ------") + "\n\nEscribe algo o presiona Enter:\n" + t.Entrada.View()
	} else if t.Paso == 2 {
		contenido = estiloTitulo.Render("------ PERFIL DE USUARIO ------") + "\n\nEscribe algo o presiona Enter:\n" + t.Entrada.View()
	}

	estiloPantallaCompleta := lipGloss.NewStyle().
		Background(lipGloss.Color(negro)).
		Border(lipGloss.ThickBorder(), true).
		BorderForeground(lipGloss.Color(vinotinto)).
		Width(t.Ancho-2).
		Height(t.Altura-2).
		Align(lipGloss.Center, lipGloss.Center)

	return estiloPantallaCompleta.Render(contenido)
}

func (t *UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		t.Ancho = msg.Width
		t.Altura = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return t, tea.Quit

		case "up", "k":
			if t.Paso == 0 {
				t.Cursor--
				if t.Cursor < 0 {
					t.Cursor = 2
				}
			}
		case "down", "j":
			if t.Paso == 0 {
				t.Cursor++
				if t.Cursor > 2 {
					t.Cursor = 0
				}
			}

		case "enter":
			if t.Paso == 0 {
				switch t.Cursor {
				case 0:
					t.Paso = 1
				case 1:
					t.Paso = 2
				case 2:
					return t, tea.Quit
				}
				t.Cursor = 0
			} else {
				valorTipeado := t.Entrada.Value()
				if valorTipeado == "" {
					return t, nil
				}
				t.Entrada.SetValue("")
			}
		}
	}

	var cmd tea.Cmd
	t.Entrada, cmd = t.Entrada.Update(msg)
	return t, cmd
}

func (t UI) Init() tea.Cmd {
	return nil
}
