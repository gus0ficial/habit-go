package users

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"time"
)

type User struct {
	ID          int    `json:"id"`
	Nombres     string `json:"nombre"`
	Nickname    string `json:"nickname"`
	FNacimiento string `json:"f_nacimiento"`
	Contrasena  string `json:"pass"`
	FRegistro   string `json:"f_registro"`
	Rol         int    `json:"rol"`
	Status      bool   `json:"status"`
}

func NuevoUsuario(users []User, nombre string, nick string, date string, pass string) []User {
	usuarioNuevo := User{
		ID:          SiguienteID(users),
		Nombres:     nombre,
		Nickname:    nick,
		FNacimiento: date,
		Contrasena:  pass,
		FRegistro:   time.Now().Format("2006-01-02"),
		Rol:         1,
		Status:      true,
	}
	return append(users, usuarioNuevo)
}

func GuardarUsuario(file *os.File, users []User) {
	bytes, err := json.MarshalIndent(users, "", "  ")
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

func SiguienteID(users []User) int {
	if len(users) == 0 {
		return 1
	}
	return users[len(users)-1].ID + 1
}

func EvitarDuplicados(valor string, listaUsuarios []User) bool {
	for _, usuario := range listaUsuarios {
		if usuario.Nickname == valor {
			return true
		}
	}
	return false
}

func ExisteUsuario(nick string, pass string, users []User) int {
	for _, usuario := range users {
		if nick == usuario.Nickname {
			if Verificar(usuario.Contrasena, pass) {
				return 2
			}
			return 1
		}
	}
	return 0
}

/* Datos del Usuario */

func Hash(plana string) string {
	var hash uint32 = 5381

	for i := 0; i < len(plana); i++ {
		hash = ((hash << 5) + hash) + uint32(plana[i])
	}

	return fmt.Sprintf("%x", hash)
}

func Verificar(hash string, plana string) bool {
	hashNuevo := Hash(plana)

	return hash == hashNuevo
}

func Contrasena(valor string) bool {
	tieneMinuscula, _ := regexp.MatchString(`[a-z]`, valor)
	tieneMayuscula, _ := regexp.MatchString(`[A-Z]`, valor)
	tieneNumero, _ := regexp.MatchString(`[0-9]`, valor)
	tieneSigno, _ := regexp.MatchString(`[^a-zA-Z0-9]`, valor)

	if tieneMinuscula && tieneMayuscula && tieneNumero && tieneSigno {
		return true
	}
	return false
}
