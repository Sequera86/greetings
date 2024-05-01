package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

//Hello devuelve un saludo

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Nombre vacio")
	}

	message := fmt.Sprintf(randomformat(), name)
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

func randomformat() string {
	formats := []string{
		"Hola %v , Bienvenido",
		"Que bueno verte, %v",
		"Saludos %v ",
	}
	return formats[rand.Intn(len(formats))]

}
