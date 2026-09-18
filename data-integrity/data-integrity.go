package data_integrity

import (
	"regexp"
	"strconv"
	"time"
)

func Longitud(valor string, longitud int) bool {
	if len(valor) > longitud {
		return false
	}
	return true
}

func Minimo(valor string, longitud int) bool {
	if len(valor) < longitud {
		return false
	}
	return true
}

func EsNumero(valor string) bool {
	_, err := strconv.Atoi(valor)
	if err != nil {
		return false
	}
	return true
}

func EsLetra(valor string) bool {
	entrada, _ := regexp.MatchString("^[a-zA-ZáéíóúÁÉÍÓÚñÑ ]+$", valor)
	return entrada
}

func Fecha(valor string) bool {
	_, err := time.Parse("2006-01-02", valor)
	if err != nil {
		return false
	}
	return true
}

func Vacio(valor any) bool {
	if valor != "" {
		return true
	}
	return false
}

