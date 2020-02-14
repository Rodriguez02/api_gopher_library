package services

import (
	"api_gopher_library/domain"
	"testing"
)

/************************** TEST LOANS ************************/

/* create loan */

func TestCreateLoanNotOk(t *testing.T) {
	r := []struct {
		TestName string
		Loan     domain.Loan
		Expected error
	}{
		{
			TestName: "loan exists",
			Loan: domain.Loan{
				ID:      1,
				IDBook:  "1C",
				IDUser:  1,
				DueDate: "2020-03-12",
				Info:    domain.Information{},
			},
			Expected: ErrorLoanExists,
		},
		{
			TestName: "incorrect id book",
			Loan: domain.Loan{
				ID:      2,
				IDBook:  "",
				IDUser:  1,
				DueDate: "2020-03-12",
				Info:    domain.Information{},
			},
			Expected: ErrorNoIDBook,
		},
		{
			TestName: "incorrect id user",
			Loan: domain.Loan{
				ID:      2,
				IDBook:  "aaaa",
				IDUser:  0,
				DueDate: "2020-03-12",
				Info:    domain.Information{},
			},
			Expected: ErrorNoIDUser,
		},
	}

	l := domain.Loan{
		ID:      1,
		IDBook:  "1B",
		IDUser:  1,
		DueDate: "2020-03-12",
		Info:    domain.Information{},
	}
	loans = append(loans, l)

	for _, c := range r {
		t.Run(c.TestName, func(t *testing.T) {
			_, err := CreateLoan(c.Loan)
			if c.Expected != err {
				t.Fail()
			}
		})
	}
}

/* get loan */

func TestGetAllLoansOk(t *testing.T) {
	CreateUser(domain.User{
		ID:       12,
		Nombre:   "Nombre",
		Apellido: "Apellido",
	})
	CreateLoan(domain.Loan{
		ID:      12,
		IDBook:  "1B",
		IDUser:  12,
		DueDate: "2020-03-12",
		Info:    domain.Information{},
	})

	_, err := GetAllLoans()
	if err == ErrorLoansNotFound {
		t.Fail()
	}
}

func TestGetAllLoansNotOk(t *testing.T) {
	loans = []domain.Loan{}
	_, err := GetAllLoans()
	if err != ErrorLoansNotFound {
		t.Fail()
	}
}

func TestGetLoanOk(t *testing.T) {
	CreateUser(domain.User{
		ID:       10,
		Nombre:   "Nombre",
		Apellido: "Apellido",
	})
	CreateLoan(domain.Loan{
		ID:      10,
		IDBook:  "1B",
		IDUser:  10,
		DueDate: "2020-03-12",
		Info:    domain.Information{},
	})

	id := "10"
	_, err := GetLoan(id)
	if err != nil {
		t.Fail()
	}
}

func TestGetLoanNotOk(t *testing.T) {
	r := []struct {
		TestName string
		IdLoan   string
		Expected error
	}{
		{
			TestName: "ID incorrect",
			IdLoan:   "addaa",
			Expected: ErrorInvalidID,
		},
		{
			TestName: "ID incorrect",
			IdLoan:   "100",
			Expected: ErrorLoanNotFound,
		},
	}

	for _, c := range r {
		t.Run(c.TestName, func(t *testing.T) {
			_, err := GetLoan(c.IdLoan)
			if c.Expected != err {
				t.Fail()
			}
		})
	}
}

/* update loan */

func TestUpdateLoanOk(t *testing.T) {
	CreateUser(domain.User{
		ID:       20,
		Nombre:   "Nombre",
		Apellido: "Apellido",
	})
	CreateLoan(domain.Loan{
		ID:      20,
		IDBook:  "1B",
		IDUser:  20,
		DueDate: "2020-03-12",
		Info:    domain.Information{},
	})
	_, err := UpdateLoan(domain.Loan{
		ID:      20,
		IDBook:  "1R",
		IDUser:  20,
		DueDate: "2020-03-12",
		Info:    domain.Information{},
	})
	if err != nil {
		t.Fail()
	}
}

func TestUpdateLoanNotOk(t *testing.T) {
	CreateUser(domain.User{
		ID:       20,
		Nombre:   "Nombre",
		Apellido: "Apellido",
	})
	CreateLoan(domain.Loan{
		ID:      20,
		IDBook:  "1B",
		IDUser:  20,
		DueDate: "2020-03-12",
		Info:    domain.Information{},
	})
	
	r := []struct {
		TestName string
		Loan     domain.Loan
		Expected error
	}{
		{
			TestName: "Loan not exist",
			Loan: domain.Loan{
				ID:      1000,
				IDBook:  "1R",
				IDUser:  1000,
				DueDate: "2020-03-12",
				Info:    domain.Information{},
			},
			Expected: ErrorLoanNotFound,
		},
		{
			TestName: "user not exist",
			Loan: domain.Loan{
				ID:      20,
				IDBook:  "jh",
				IDUser:  9,
				DueDate: "2020-03-12",
				Info:    domain.Information{},
			},
			Expected: ErrorUserNotFound,
		},
		{
			TestName: "id book incorrect",
			Loan: domain.Loan{
				ID:      20,
				IDBook:  "",
				IDUser:  20,
				DueDate: "2020-03-12",
				Info:    domain.Information{},
			},
			Expected: ErrorNoIDBook,
		},
		{
			TestName: "id user incorrect",
			Loan: domain.Loan{
				ID:      20,
				IDBook:  "ab",
				IDUser:  0,
				DueDate: "2020-03-12",
				Info:    domain.Information{},
			},
			Expected: ErrorNoIDUser,
		},
	}

	for _, c := range r {
		t.Run(c.TestName, func(t *testing.T) {
			_, err := UpdateLoan(c.Loan)
			if c.Expected != err {
				t.Fail()
			}
		})
	}
}

/* delete loan */

func TestDeleteLoanOk(t *testing.T) {
	CreateUser(domain.User{
		ID:       10,
		Nombre:   "Nombre",
		Apellido: "Apellido",
	})
	CreateLoan(domain.Loan{
		ID:      10,
		IDBook:  "1B",
		IDUser:  10,
		DueDate: "2020-03-12",
		Info:    domain.Information{},
	})

	id := "10"
	_, err := DeleteLoan(id)
	if err != nil {
		t.Fail()
	}
}

func TestDeleteLoanNotOk(t *testing.T) {
	r := []struct {
		TestName string
		IdLoan   string
		Expected error
	}{
		{
			TestName: "ID incorrect",
			IdLoan:   "addaa",
			Expected: ErrorInvalidID,
		},
		{
			TestName: "ID incorrect",
			IdLoan:   "100",
			Expected: ErrorLoanNotFound,
		},
	}

	for _, c := range r {
		t.Run(c.TestName, func(t *testing.T) {
			_, err := DeleteLoan(c.IdLoan)
			if c.Expected != err {
				t.Fail()
			}
		})
	}
}

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
	CreateUser(domain.User{
		ID:       2,
		Nombre:   "Nombre",
		Apellido: "Apellido",
	})
	_, err := GetAllUsers()
	if ErrorUsersNotFound == err {
		t.Fail()
	}
}

func TestGetAllUsersNotOK(t *testing.T) {
	users = []domain.User{}
	_, err := GetAllUsers()
	if ErrorUsersNotFound != err {
		t.Fail()
	}
}

func TestGetUserOK(t *testing.T) {
	r := struct {
		TestName string
		User     domain.User
		Expected error
	}{
		TestName: "Get user",
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

/************************** TEST BOOKS ************************/

func TestValidateBookOk(t *testing.T) {
	book := domain.Book{
		Titulo: "Titulo",
		Autor:  "Autor",
	}
	err := validateBook(book)
	if err != nil {
		t.Fail()
	}
}

func TestValidateBookNotOk(t *testing.T) {
	book := domain.Book{
		Titulo: "Titu/lo",
		Autor:  "Autor",
	}
	err := validateBook(book)
	if err == nil {
		t.Fail()
	}
}
