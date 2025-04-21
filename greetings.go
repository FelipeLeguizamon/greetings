package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello devuelve un saludo que incluye para la persona especificada
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Nombre vacio")
	}

	//devuelve un saludo que incluye el nombre un mensaje
	message := fmt.Sprintf(ramdomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// ramdonFormat devuelve unode varios formatos de mensajes de salida
// se selecciona de forma aleatoria
func ramdomFormat() string {
	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"¡Que bueno verte ,%v!",
		"¡Saludo,%v!¡Encantado de conocerte!",
	}
	return formats[rand.Intn(len(formats))]
}
