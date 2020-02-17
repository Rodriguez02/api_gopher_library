package services

import (
	"api_gopher_library/domain"
	"testing"
)

var (
	requestLoan = domain.Loan{
		ID:      1,
		IDBook:  "1C",
		IDUser:  1,
		DueDate: "2020-03-12",
		Info:    domain.Information{},
	}
	expectedLoan = domain.Loan{
		ID:      1,
		IDBook:  "1C",
		IDUser:  1,
		DueDate: "2020-03-12",
		Info: domain.Information{
			Titulo:           "The City",
			Subtitulo:        "",
			Autores:          [5]string{"Robert E. Park", "Ernest W. Burgess", "", "", ""},
			FechaPublicacion: "2019-04-19",
		},
	}
	expectedUser = domain.User{
		ID:       1,
		Nombre:   "Nombre",
		Apellido: "Apellido",
	}
	url = based + "Lord" + "+inauthor:" + "tolkien" + "&fields=items(id,volumeInfo(title,subtitle,authors,publishedDate))"
)

func ClearArrays() {
	users = []domain.User{}
	loans = []domain.Loan{}
}

func CreateDataUser() {
	CreateUser(domain.User{
		ID:       1,
		Nombre:   "Nombre",
		Apellido: "Apellido",
	})
}

/************************** TEST LOANS ************************/

/* create loan */
func TestCreateLoanOk(t *testing.T) {
	ClearArrays()
	CreateDataUser()

	_, err := CreateLoan(requestLoan)
	if err != nil {
		t.Fail()
	}
}

func TestCreateLoanExist(t *testing.T) {
	ClearArrays()
	CreateDataUser()
	loans = append(loans, expectedLoan)

	_, err := CreateLoan(requestLoan)
	if err != ErrorLoanExists {
		t.Fail()
	}
}

func TestCreateLoanNoAvailability(t *testing.T) {
	ClearArrays()
	CreateDataUser()
	booksAmount["1C"] = 5

	_, err := CreateLoan(requestLoan)
	if err != ErrorNoAvailability {
		t.Fail()
	}
}

/* get loan */

func TestGetAllLoansOk(t *testing.T) {
	ClearArrays()
	CreateDataUser()
	loans = append(loans, expectedLoan)

	_, err := GetAllLoans()
	if err == ErrorLoansNotFound {
		t.Fail()
	}
}

func TestGetAllLoansNotOk(t *testing.T) {
	ClearArrays()

	_, err := GetAllLoans()
	if err != ErrorLoansNotFound {
		t.Fail()
	}
}

func TestGetLoanOk(t *testing.T) {
	ClearArrays()
	CreateDataUser()
	loans = append(loans, expectedLoan)

	_, err := GetLoan("1")
	if err != nil {
		t.Fail()
	}
}

func TestGetLoanNotOk(t *testing.T) {
	ClearArrays()

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
	ClearArrays()
	CreateDataUser()
	loans = append(loans, expectedLoan)

	_, err := UpdateLoan(domain.Loan{
		ID:      1,
		IDBook:  "1R",
		IDUser:  1,
		DueDate: "2020-03-12",
		Info:    domain.Information{},
	})
	if err != nil {
		t.Fail()
	}
}

func TestUpdateLoanNotFound(t *testing.T) {
	ClearArrays()
	CreateDataUser()

	_, err := UpdateLoan(expectedLoan)
	if err != ErrorLoanNotFound {
		t.Fail()
	}
}

/* delete loan */

func TestDeleteLoanOk(t *testing.T) {
	ClearArrays()
	CreateDataUser()
	loans = append(loans, expectedLoan)

	_, err := DeleteLoan("1")
	if err != nil {
		t.Fail()
	}
}

func TestDeleteLoanNotOk(t *testing.T) {
	ClearArrays()

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
	ClearArrays()

	r, _ := CreateUser(expectedUser)
	if r != expectedUser {
		t.Fail()
	}
}

func TestCreateUserNotOk(t *testing.T) {
	ClearArrays()
	CreateDataUser()

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
			_, err := CreateUser(c.User)
			if c.Expected != err {
				t.Fail()
			}
		})
	}
}

/* get users */

func TestGetAllUsersOK(t *testing.T) {
	ClearArrays()
	CreateDataUser()

	_, err := GetAllUsers()
	if ErrorUsersNotFound == err {
		t.Fail()
	}
}

func TestGetAllUsersNotOK(t *testing.T) {
	ClearArrays()

	_, err := GetAllUsers()
	if ErrorUsersNotFound != err {
		t.Fail()
	}
}

func TestGetUserOK(t *testing.T) {
	ClearArrays()
	CreateDataUser()

	u, _ := GetUser("1")
	if u != expectedUser {
		t.Fail()
	}
}

func TestGetUserNotOK(t *testing.T) {
	ClearArrays()
	CreateDataUser()

	_, err := GetUser("2")
	if ErrorUserNotFound != err {
		t.Fail()
	}
}

/* update users */

func TestUpdateUserOK(t *testing.T) {
	ClearArrays()
	CreateDataUser()

	_, err := UpdateUser(domain.User{
		ID:       1,
		Nombre:   "Nombre1",
		Apellido: "Apellido1",
	})
	if err != nil {
		t.Fail()
	}
}

func TestUpdateUserNotOK(t *testing.T) {
	ClearArrays()
	CreateDataUser()

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
			TestName: "User not found",
			User: domain.User{
				ID:       2,
				Nombre:   "Nombre",
				Apellido: "Apellido",
			},
			Expected: ErrorUserNotFound,
		},
	}

	for _, c := range r {
		t.Run(c.TestName, func(t *testing.T) {
			_, err := UpdateUser(c.User)
			if c.Expected != err {
				t.Fail()
			}
		})
	}
}

/* delete users */

func TestDeleteUserOk(t *testing.T) {
	ClearArrays()
	CreateDataUser()

	_, err := DeleteUser("1")
	if err != nil {
		t.Fail()
	}
}

func TestDeleteUserNotOk(t *testing.T) {
	ClearArrays()
	CreateDataUser()

	r := []struct {
		TestName string
		User     string
		Expected error
	}{
		{
			TestName: "ID incorrect 1",
			User:     "assd",
			Expected: ErrorInvalidID,
		},
		{
			TestName: "User not found",
			User:     "2",
			Expected: ErrorUserNotFound,
		},
	}

	for _, c := range r {
		t.Run(c.TestName, func(t *testing.T) {
			_, err := DeleteUser(c.User)
			if c.Expected != err {
				t.Fail()
			}
		})
	}
}

/************************** TEST BOOKS ************************/

func TestGetBookOk(t *testing.T) {
	_, err := GetBook(domain.Book{
		Titulo: "The Lord",
		Autor:  "Tolkien",
	})
	if err != nil {
		t.Fail()
	}
}

func TestGetBookNotOk(t *testing.T) {
	r := []struct {
		TestName string
		Book     domain.Book
		Expected error
	}{
		{
			TestName: "book special char",
			Book: domain.Book{
				Titulo: "The Lord",
				Autor:  "Tol-?kien",
			},
			Expected: ErrorSpecialCharInBooks,
		},
		{
			TestName: "book not found",
			Book: domain.Book{
				Titulo: "123n1dasmdp12",
				Autor:  "zsouo23eo120",
			},
			Expected: ErrorBookNotFound,
		},
	}

	for _, c := range r {
		t.Run(c.TestName, func(t *testing.T) {
			_, err := GetBook(c.Book)
			if c.Expected != err {
				t.Fail()
			}
		})
	}
}

func TestApiResponseOk(t *testing.T) {
	_, err := apiResponse(url)
	if err != nil {
		t.Fail()
	}
}

func TestApiResponseNotOk(t *testing.T) {
	r := []struct {
		TestName string
		Url      string
		Expected error
	}{
		{
			TestName: "book not found",
			Url:      (based + "asd1231" + "+inauthor:" + "1231jb23kd1" + "&fields=items(id,volumeInfo(title,subtitle,authors,publishedDate))"),
			Expected: ErrorBookNotFound,
		},
	}

	for _, c := range r {
		t.Run(c.TestName, func(t *testing.T) {
			_, err := apiResponse(c.Url)
			if c.Expected != err {
				t.Fail()
			}
		})
	}
}
