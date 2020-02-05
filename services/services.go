package services

import (
	"api_gopher_library/domain"
	"errors"
	"strconv"
)

var (
	users []domain.User

	// errores
	ErrorNoName        = errors.New("user needs name")
	ErrorNoSurname     = errors.New("user needs surname")
	ErrorInvalidID     = errors.New("id isn't valid")
	ErrorUserExists    = errors.New("this user exists")
	ErrorUsersNotFound = errors.New("there aren't users")
	ErrorUserNotFound  = errors.New("user not found")
)

/********************************************************/
/******************** CRUD USERS ************************/
/********************************************************/

// Crea un usuario y lo guarda en el arreglo 'users'
func CreateUser(user domain.User) (domain.User, error) {
	err := validateUser(user)
	if err != nil {
		return user, err
	}

	err = existsUser(user)
	if err != nil {
		return user, err
	}

	users = append(users, user)
	return user, nil
}

// Obtengo todos los usuarios guardados en 'users'
// Sino exista ningún usuario se devuelve un error
func GetAllUsers() ([]domain.User, error) {
	if users == nil {
		return users, ErrorUsersNotFound
	}
	return users, nil
}

// Obtengo un usuario en particular que se espcifique
// en el body, si tiene la misma ID devuelve el mismo
// sino un error de que no se encontró el usuario
func GetUser(i string) (domain.User, error) {
	id, err := validateID(i)
	if err != nil {
		return domain.User{}, ErrorInvalidID
	}

	user, err := searchUser(id)
	if err != nil{
		return domain.User{}, err
	}
	
	return user, nil
}

// Actualizo el Nombre y Apellido que se pase por el body
// Sino tiene la misma ID devuelve un error
func UpdateUser(user domain.User) (domain.User, error) {
	err := validateUser(user)
	if err != nil {
		return user, nil
	}

	err = existsUser(user)
	if err != nil {
		for i := 0; i < len(users); i++ {
			if users[i].ID == user.ID {
				users[i].Nombre = user.Nombre
				users[i].Apellido = user.Apellido
				return users[i], nil
			}
		}
	}

	return user, ErrorUserNotFound
}

// Eliminar un usuario del arreglo 'users' según el ID
// del usuario pasada por el body
func DeleteUser(user domain.User) (domain.User, error) {
	err := validateUser(user)
	if err != nil {
		return user, nil
	}

	err = existsUser(user)
	if err != nil {
		for i := 0; i < len(users); i++ {
			if users[i].ID == user.ID {
				users = append(users[:i], users[i+1:]...)
				return user, nil
			}
		}
	}

	return user, ErrorUserNotFound
}

func validateUser(user domain.User) error {
	if !user.HasName() {
		return ErrorNoName
	}

	if !user.HasSurname() {
		return ErrorNoSurname
	}

	if !user.IDValid() {
		return ErrorInvalidID
	}

	return nil
}

func existsUser(user domain.User) error {
	for _, u := range users {
		if u.ID == user.ID {
			return ErrorUserExists
		}
	}
	return nil
}

func validateID(id string) (int, error) {
	num, err := strconv.Atoi(id)
	if err != nil {
		return -1, err
	}
	return num, nil
}

func searchUser(id int) (domain.User, error){	
	for _,u := range users {
		if u.ID == id{
			return u, nil
		}
	}
	return domain.User{}, ErrorUserNotFound
}