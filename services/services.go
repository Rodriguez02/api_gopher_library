package services

import (
	"api_gopher_library/domain"
	"errors"
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

func CreateUser(user domain.User) (domain.User, error) {
	err := validateUser(user)
	if err != nil {
		return user, err
	}
	users = append(users, user)
	return user, nil
}

func GetAllUsers() ([]domain.User, error) {
	if users == nil {
		return users, ErrorUsersNotFound
	}
	return users, nil
}

func GetUser(user domain.User) (domain.User, error) {
	err := validateUser(user)
	if err == ErrorUserExists {
		return user, nil
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

	for _, u := range users {
		if u.ID == user.ID {
			return ErrorUserExists
		}
	}

	return nil
}
