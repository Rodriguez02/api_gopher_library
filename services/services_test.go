package services

import (
	"api_gopher_library/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

/************************** TEST LOANS ************************/
func TestConstants(t *testing.T) {
	assert.EqualValues(t, 2, amountBooks)
	assert.EqualValues(t, "https://www.googleapis.com/books/v1/volumes?key=AIzaSyDVnZCPWXdzNcWiipQ7ng5E-eLRg3xu7MY&q=", based)
}

/* Func aux */



/************************** TEST USERS ************************/

/* create users */
func TestCreateUserOK(t *testing.T) {
	r := struct {
		TestName string
		User     domain.User
		Expected domain.User
	}{
		TestName: "Correct user",
		User: domain.User{
			ID:       1,
			Nombre:   "Nombre",
			Apellido: "Apellido",
		},
		Expected: domain.User{
			ID:       1,
			Nombre:   "Nombre",
			Apellido: "Apellido",
		},
	}
	t.Run(r.TestName, func(t *testing.T) {
		result, _ := CreateUser(r.User)
		if r.Expected != result {
			t.Fail()
		}
	})
}

func TestCreateUserNotOk(t *testing.T) {
	r := []struct {
		TestName string
		User     domain.User
		Expected error
	}{
		{
			TestName: "ID incorrect",
			User: domain.User{
				ID:       -1,
				Nombre:   "Nombre",
				Apellido: "Apellido",
			},
			Expected: ErrorInvalidID,
		},
		{
			TestName: "Name incorrect",
			User: domain.User{
				ID:       1,
				Nombre:   "",
				Apellido: "Apellido",
			},
			Expected: ErrorNoName,
		},
		{
			TestName: "Surname incorrect",
			User: domain.User{
				ID:       1,
				Nombre:   "Nombre",
				Apellido: "",
			},
			Expected: ErrorNoSurname,
		},
		{
			TestName: "User exist",
			User: domain.User{
				ID:       1,
				Nombre:   "Nombre",
				Apellido: "Apellido",
			},
			Expected: ErrorUserExists,
		},
	}

	for _, c := range r {
		t.Run(c.TestName, func(t *testing.T) {
			u := domain.User{
				ID:       1,
				Nombre:   "Nombre",
				Apellido: "Apellido",
			}
			users = append(users, u)
			_, err := CreateUser(c.User)
			if c.Expected != err {
				t.Fail()
			}
		})
	}
}

/* get users */
func TestGetAllUsersOK(t *testing.T) {
	r := struct {
		TestName string
		User     domain.User
		Expected error
	}{
		TestName: "Get all users",
		User: domain.User{
			ID:       1,
			Nombre:   "Nombre",
			Apellido: "Apellido",
		},
		Expected: nil,
	}

	t.Run(r.TestName, func(t *testing.T) {
		CreateUser(r.User)
		_, err := GetAllUsers()
		if r.Expected != err {
			t.Fail()
		}
	})
}

func TestGetAllUsersNotOK(t *testing.T) {
	users = []domain.User{}
	_, err := GetAllUsers()
	if  ErrorUsersNotFound != err {
		t.Fail()
	}
}

func TestGetUserOK(t *testing.T) {
	r := struct {
		TestName string
		User     domain.User
		Expected error
	}{
		TestName: "Get all users",
		User: domain.User{
			ID:       1,
			Nombre:   "Nombre",
			Apellido: "Apellido",
		},
		Expected: nil,
	}

	t.Run(r.TestName, func(t *testing.T) {
		CreateUser(r.User)
		id := "1"
		_, err := GetUser(id)
		if r.Expected != err {
			t.Fail()
		}
	})
}

func TestGetUserNotOK(t *testing.T) {
	r := struct {
		TestName string
		User     domain.User
		Expected error
	}{
		TestName: "Get all users",
		User: domain.User{
			ID:       1,
			Nombre:   "Nombre",
			Apellido: "Apellido",
		},
		Expected: ErrorUserNotFound,
	}

	t.Run(r.TestName, func(t *testing.T) {
		CreateUser(r.User)
		id := "2"
		_, err := GetUser(id)
		if r.Expected != err {
			t.Fail()
		}
	})
}

/* update users */

func TestUpdateUserOK(t *testing.T) {
	r := struct {
		TestName string
		User     domain.User
		Expected error
	}{
		TestName: "Update user",
		User: domain.User{
			ID:       1,
			Nombre:   "Nombre",
			Apellido: "Apellido",
		},
		Expected: nil,
	}

	t.Run(r.TestName, func(t *testing.T) {
		u := domain.User{
			ID:       1,
			Nombre:   "Nombre1",
			Apellido: "Apellido1",
		}
		CreateUser(r.User)
		_, err := UpdateUser(u)
		if r.Expected != err {
			t.Fail()
		}

	})
}

func TestUpdateUserNotOK(t *testing.T) {
	r := struct {
		TestName string
		User     domain.User
		Expected error
	}{
		TestName: "No update user",
		User: domain.User{
			ID:       1,
			Nombre:   "Nombre",
			Apellido: "Apellido",
		},
		Expected: ErrorUserNotFound,
	}

	t.Run(r.TestName, func(t *testing.T) {
		u := domain.User{
			ID:       2,
			Nombre:   "Nombre1",
			Apellido: "Apellido1",
		}
		CreateUser(r.User)
		_, err := UpdateUser(u)
		if r.Expected != err {
			t.Fail()
		}

	})
}

/* delete users */

func TestDeleteUserOk(t *testing.T) {
	r := struct {
		TestName string
		User     domain.User
		Expected error
	}{
		TestName: "Delete user",
		User: domain.User{
			ID:       1,
			Nombre:   "Nombre",
			Apellido: "Apellido",
		},
		Expected: nil,
	}

	t.Run(r.TestName, func(t *testing.T) {
		CreateUser(r.User)
		id := "1"
		_, err := DeleteUser(id)
		if r.Expected != err {
			t.Fail()
		}
	})
}

func TestDeleteUserNotOk(t *testing.T) {
	r := struct {
		TestName string
		User     domain.User
		Expected error
	}{
		TestName: "Delete user failed",
		User: domain.User{
			ID:       1,
			Nombre:   "Nombre",
			Apellido: "Apellido",
		},
		Expected: ErrorUserNotFound,
	}

	t.Run(r.TestName, func(t *testing.T) {
		CreateUser(r.User)
		id := "2"
		_, err := DeleteUser(id)
		if r.Expected != err {
			t.Fail()
		}
	})
}


func TestValidateBookOk(t *testing.T){
	book := domain.Book {
		Titulo: "Titulo",
		Autor: "Autor",
	}
	err := validateBook(book)
	if err != nil {
		t.Fail()
	}
}

func TestValidateBookNotOk(t *testing.T){
	book := domain.Book {
		Titulo: "Titu/lo",
		Autor: "Autor",
	}
	err := validateBook(book)
	if err == nil {
		t.Fail()
	}
}