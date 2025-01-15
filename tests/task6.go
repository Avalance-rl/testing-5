package tests

import "fmt"

type Storage interface {
	GetUserByID(id string) (Person, error)
}

type Person struct {
	ID    string
	Name  string
	Email string
}

func GetUserInfo(storage Storage, id string) (string, error) {
	person, err := storage.GetUserByID(id)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("id = %s, name = %s, email = %s", person.ID, person.Name, person.Email), nil
}
