/*
* @Author: gustav0
* @Date:   2026-08-24 23:31:02
* @Last Modified by:   Gustav0
* @Last Modified time: 2026-09-16 21:28:02
 */

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	uiPackage "habit-go/UI"
	dataPackage "habit-go/data-integrity"
	habitPackage "habit-go/habits"
	usersPackage "habit-go/users"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var NombreProyecto = "go-habit"
var Proposito = "\033[1m" + NombreProyecto + "\033[0m es una app cli desarrollada para gestionar \033[1mHabitos Personales\033[0m."

func main() {
	archivoHabitos, err := os.OpenFile("habits.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer archivoHabitos.Close()
	var habits []habitPackage.Habit
	info, err := archivoHabitos.Stat()
	if err != nil {
		panic(err)
	}
	if info.Size() != 0 {
		bytes, err := io.ReadAll(archivoHabitos)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &habits)
		if err != nil {
			panic(err)
		}
	} else {
		habits = []habitPackage.Habit{}
	}
	archivoHabitos.Close()

	archivoUsuarios, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer archivoUsuarios.Close()
	var users []usersPackage.User
	infoUsers, err := archivoUsuarios.Stat()
	if err != nil {
		panic(err)
	}
	if infoUsers.Size() != 0 {
		bytes, err := io.ReadAll(archivoUsuarios)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &users)
		if err != nil {
			panic(err)
		}
	} else {
		users = []usersPackage.User{}
	}
	archivoUsuarios.Close()

	if len(os.Args) < 2 || len(os.Args) == 0 {
		fmt.Printf("%s\n", Proposito)
		fmt.Printf("Para iniciar, tipee \033[1m%s init\033[0m.\nConsulte el resto de comandos tipeando \033[1m%s help\033[0m.\n", NombreProyecto, NombreProyecto)
		return
	} else {
		switch os.Args[1] {
		case "init":
			lector := bufio.NewReader(os.Stdin)
			var nick, pass string

			if len(os.Args) >= 4 {
				nick = os.Args[2]
				pass = os.Args[3]

				if usersPackage.ExisteUsuario(nick, pass, users) != 2 {
					fmt.Printf("\033[1mError:\033[0m Credenciales Incorrectas.\n")
					return
				}
				letsGo(nick, users, habits)

			} else if len(os.Args) == 3 {
				nick = os.Args[2]

				fmt.Printf("Ingrese la \033[1mContraseña\033[0m para el perfil de '%s':\n", nick)

				entrada, _ := lector.ReadString('\n')
				entrada = strings.TrimSpace(entrada)
				pass = entrada

				if usersPackage.ExisteUsuario(nick, pass, users) != 2 {
					fmt.Printf("\033[1mError:\033[0m Credenciales Incorrectas.\n")
					return
				}
				letsGo(nick, users, habits)

			} else {
				fmt.Printf("Ingrese su \033[1mNickname\033[0m:\n")
				entradaNick, _ := lector.ReadString('\n')
				nick = strings.TrimSpace(entradaNick)

				fmt.Printf("Ingrese su \033[1mContraseña\033[0m:\n")
				entradaPass, _ := lector.ReadString('\n')
				pass = strings.TrimSpace(entradaPass)

				if usersPackage.ExisteUsuario(nick, pass, users) != 2 {
					fmt.Printf("\033[1mError:\033[0m Credenciales Incorrectas.\n")
					return
				}
				letsGo(nick, users, habits)
			}
		case "new":
			lector := bufio.NewReader(os.Stdin)
			var nombre, nick, fNacimiento, pass string
			fmt.Printf("Ingrese \033[1mNombre del Usuario\033[0m del Nuevo Perfil (Máx. 32 letras):\n")
			for {
				entrada, _ := lector.ReadString('\n')
				entrada = strings.TrimSpace(entrada)
				if !dataPackage.Vacio(entrada) {
					fmt.Printf("\033[1mError:\033[0m Registrar un Nombre es obligatorio..\n")
					continue
				}
				if !dataPackage.Minimo(entrada, 3) {
					fmt.Printf("\033[1mError:\033[0m Un nombre debe tener al menos 3 caracteres.\n")
					continue
				}
				if !dataPackage.Longitud(entrada, 32) {
					fmt.Printf("\033[1mError:\033[0m El máximo de caracteres para un Nombre es 32.\n")
					continue
				}
				if dataPackage.EsNumero(entrada) {
					fmt.Printf("\033[1mError:\033[0m Ingrese sólo letras.\n")
					continue
				}

				nombre = entrada
				fmt.Println("Este perfil lo usará:", nombre)
				break
			}

			fmt.Printf("Ingrese \033[1mNickname\033[0m del Nuevo Perfil (Máx. 12 caracteres):\n")
			for {
				entrada, _ := lector.ReadString('\n')
				entrada = strings.TrimSpace(entrada)

				if !dataPackage.Vacio(entrada) {
					fmt.Printf("\033[1mError:\033[0m Registrar un Nickname es obligatorio.\n")
					continue
				}
				if !dataPackage.Minimo(entrada, 3) {
					fmt.Printf("\033[1mError:\033[0m Un Nickname debe tener al menos 3 caracteres.\n")
					continue
				}
				if !dataPackage.Longitud(entrada, 12) {
					fmt.Printf("\033[1mError:\033[0m El máximo de caracteres para un Nickname es 12.\n")
					continue
				}
				if usersPackage.EvitarDuplicados(entrada, users) {
					fmt.Printf("\033[1mError:\033[0m Este nickname ya está tomado.\n")
					continue
				}

				nick = entrada
				fmt.Println("Ha eligido:", nick)
				break
			}

			fmt.Printf("\033[1m(OPCIONAL)\033[0m Ingrese \033[1mFecha de Nacimiento\033[0m (Formato AAAA-MM-DD):\n")
			for {
				entrada, _ := lector.ReadString('\n')
				entrada = strings.TrimSpace(entrada)

				if entrada != "" && !dataPackage.Fecha(entrada) {
					fmt.Printf("\033[1mError:\033[0m Para guardar una fecha, siga el formato: AAAA-MM-DD.\n")
					continue
				}

				fNacimiento = entrada
				if fNacimiento != "" {
					fmt.Println("Ha elegido:", fNacimiento)
				}
				break
			}

			fmt.Printf("\033[1m(OBLIGATORIO)\033[0m Ingrese la \033[1mContraseña\033[0m del Nuevo Perfil.\nLa misma debe incluir, al menos:\n - Un máximo de 12 caracteres.\n - Al menos, una (01) letra mayuscula y una (01) minuscula.\n - Al menos, un (01) signo.\n - Al menos, un (01) número.\n")
			for {
				entrada, _ := lector.ReadString('\n')
				entrada = strings.TrimSpace(entrada)

				if !dataPackage.Vacio(entrada) {
					fmt.Printf("\033[1mError:\033[0m Una Contraseña es obligatoria para crear un Perfil.\n")
					continue
				}
				if !dataPackage.Minimo(entrada, 3) {
					fmt.Printf("\033[1mError:\033[0m La Contraseña debe tener al menos 3 caracteres.\n")
					continue
				}
				if !dataPackage.Longitud(entrada, 12) {
					fmt.Printf("\033[1mError:\033[0m El máximo de caracteres para una Contraseña es 12.\n")
					continue
				}
				if !usersPackage.Contrasena(entrada) {
					fmt.Printf("\033[1mError:\033[0m La contraseña debe cumplir con las condiciones dadas.\n")
					continue
				}

				pass = usersPackage.Hash(entrada)
				fmt.Println("Ha ingresado una contraseña.")
				break
			}

			if fNacimiento != "" {
				fmt.Printf("¿Esto es correcto?:\n - %s\n - %s\n - %s\nSí(s/S) o No(n/N)\n", nombre, nick, fNacimiento)
			} else {
				fmt.Printf("¿Esto es correcto?:\n - %s\n - %s\nSí(s/S) o No(n/N)\n", nombre, nick)
			}
			for {
				respuesta, _ := lector.ReadString('\n')
				respuesta = strings.TrimSpace(respuesta)
				if respuesta == "s" || respuesta == "S" {
					users = usersPackage.NuevoUsuario(users, nombre, nick, fNacimiento, pass)

					escritura, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0666)
					if err == nil {
						usersPackage.GuardarUsuario(escritura, users)
						escritura.Close()
					}

					fmt.Printf("\033[1mProceso:\033[0m Se ha creado un usuario.\nAhora puede emplear la utilidad.\n")
				} else if respuesta == "n" || respuesta == "N" {
					fmt.Printf("\033[1mProceso:\033[0m No se ha creado un usuario.\nSe cancela el proceso.\n")
				} else if dataPackage.EsNumero(respuesta) || dataPackage.Vacio(respuesta) || dataPackage.Longitud(respuesta, 1) {
					fmt.Printf("\033[1mError:\033[0m Indique si ejecutar el guardado.\n")
					continue
				}
				break
			}
		case "help":
			MostrarUso()
		case "check":
			if len(os.Args) > 2 && os.Args[2] == "users" {
				if len(users) == 0 {
					fmt.Printf("No hay usuarios registrados.\n")
				} else {
					for _, u := range users {
						fmt.Printf(" - %s.\n", u.Nickname)
					}
				}
			} else if len(os.Args) > 2 && os.Args[2] == "data" {
				if len(users) == 0 {
					fmt.Printf("No hay perfiles creados")
					if len(habits) == 0 {
						fmt.Printf(", ni datos asociados")
					}
					fmt.Printf(".\n")
				} else {
					fmt.Printf("Hay perfiles creados")
					if len(habits) != 0 {
						fmt.Printf(", con datos asociados")
					}
					fmt.Printf(".\n")
				}
			} else {
				if len(users) == 0 {
					fmt.Printf("No hay perfiles creados")
					if len(habits) == 0 {
						fmt.Printf(", ni datos asociados")
					}
					fmt.Printf(".\n")
				} else {
					fmt.Printf("Hay perfiles creados")
					if len(habits) != 0 {
						fmt.Printf(", con datos asociados")
					}
					fmt.Printf(".\n")
				}
			}
		default:
			fmt.Printf("\033[1mError:\033[0m Consulte la paleta de comandos disponibles tipeando \033[1m%s help\033[0m.\n", NombreProyecto)
			return
		}
	}
}

func MostrarUso() {
	arrayComandos := []string{
		/* Lista de Comandos */
		"	\033[1mLista de Comandos:\033[0m",
		/* Familia init */
		"	\033[1minit:\033[0m",
		"		init -> Inicializa la app",
		"		init [NOMBRE-DEL-PERFIL] -> Inicializa la app, debe ingresar contraseña",
		"		init [NOMBRE-DEL-PERFIL] [CONTRASEÑA-DEL-PERFIL] -> Inicializa la app, sin verificación de credenciales (MENOS SEGURO)",
		/* Familia new */
		"	\033[1mnew:\033[0m",
		"		new -> Crea un nuevo Perfil, debe ingresar credenciales",
		/* Familia check */
		"	\033[1mcheck:\033[0m",
		"		check -> Verifica si hay datos guardados",
		"		check data -> Verifica si hay datos guardados",
		"		check users -> Lista los Perfiles Registrados",
		/* Familia lang */
		"	\033[1mlang:\033[0m",
		"		lang -> Indica el idioma actual de la App",
		"		lang es -> Cambia el idioma de la App a Español",
		"		lang en -> Cambia el idioma de la App a Inglés",
		/* Huerfanos */
		"	\033[1motros:\033[0m",
		"		credits -> Desarollado por gustav0",
	}

	fmt.Printf("%s\n", Proposito)
	for _, comando := range arrayComandos {
		fmt.Printf("%s \n", comando)
	}
}

func letsGo(nick string, usuarios []usersPackage.User, habitos []habitPackage.Habit) {
	var idUsuarioActivo int

	for _, usuario := range usuarios {
		if usuario.Nickname == nick {
			idUsuarioActivo = usuario.ID
			break
		}
	}
	fmt.Printf("\n\033[1mProceso:\033[0m Inicio de sesión exitoso.\n")
	fmt.Println("\033[1mProceso:\033[0m Cargando panel de gestión de hábitos...")

	IniciarMenu(idUsuarioActivo, habitos)
}

func IniciarMenu(idUsuarioActivo int, habitos []habitPackage.Habit) {
	cuadro := textinput.New()
	cuadro.Placeholder = "Escribe tu nombre..."
	cuadro.Focus()

	modelo := uiPackage.UI{
		IdUsuario: idUsuarioActivo,
		Paso:      0,
		Cursor:    0,
		Entrada:   cuadro,
		Habitos:   habitos,
		Altura:    0,
		Ancho:     0,
	}

	p := tea.NewProgram(&modelo, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error al iniciar la interfaz interactiva: %v\n", err)
		os.Exit(1)
	}
}
