package services

import (
	"api_gopher_library/domain"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	users       []domain.User          // array to storage users
	loans       []domain.Loan          // array to storage loans
	booksAmount = make(map[string]int) // map to storage the amount of book copies on loan

	// Books errors
	ErrorBookNotFound = errors.New("book not found")

	// Users errors
	ErrorNoName             = errors.New("user needs name")
	ErrorNoSurname          = errors.New("user needs surname")
	ErrorInvalidID          = errors.New("id isn't valid")
	ErrorUserExists         = errors.New("this user exists")
	ErrorUsersNotFound      = errors.New("there aren't users")
	ErrorUserNotFound       = errors.New("user not found")
	ErrorRequestExternalAPI = errors.New("error in request to external API")
	ErrorSpecialCharInBooks = errors.New("there're almost a special character in title or author")

	// Loans errors
	ErrorNoIDBook          = errors.New("loan needs idbook")
	ErrorNoIDUser          = errors.New("loan needs iduser")
	ErrorNoDueDate         = errors.New("loan needs due")
	ErrorLoanExists        = errors.New("this loan exists")
	ErrorLoansNotFound     = errors.New("there aren't loans")
	ErrorLoanNotFound      = errors.New("loan not found")
	ErrorNoAvailability    = errors.New("book without availability")
	ErrorExpiredLoans      = errors.New("user has expired loans")
	ErrorInvalidDueDate    = errors.New("due date must be greater than current")
	ErrorInvalidFormatDate = errors.New("invalid due date. example: 2 Jan 2006 = 2006-01-02")
)

const (
	amountBooks = 5                                                                                            // max amount of copies in all books
	based       = "https://www.googleapis.com/books/v1/volumes?key=AIzaSyDVnZCPWXdzNcWiipQ7ng5E-eLRg3xu7MY&q=" // url based to consume API
)

/*************************************************************
************************** CRUD LOANS ************************
*************************************************************/

/*CreateLoan ...
create loans through array of type domain.Loan
search the book, set the Info in loan and control
if exists the book in the loan copies.
If exists increment one, else the add it to the map amountBooks */
func CreateLoan(loan domain.Loan) (domain.Loan, error) {
	err := existsLoan(loan)
	if err != nil {
		return domain.Loan{}, ErrorLoanExists
	}

	err = validateLoan(loan)
	if err != nil {
		return domain.Loan{}, err
	}

	info, _ := searchBook(loan.IDBook)
	amount, _ := availability(loan.IDBook)

	loan.Info.Titulo = info.Titulo
	loan.Info.Subtitulo = info.Subtitulo
	loan.Info.Autores = info.Autores
	loan.Info.FechaPublicacion = info.FechaPublicacion

	if amount >= amountBooks {
		return domain.Loan{}, ErrorNoAvailability
	}
	if amount == 0 {
		booksAmount[loan.IDBook] = 1
	}
	if amount < amountBooks && amount != 0 {
		booksAmount[loan.IDBook]++
	}

	loans = append(loans, loan)
	return loan, nil
}

/*GetAllLoans ...
get all the loans from the array loans
if not exists loans then send an error*/
func GetAllLoans() ([]domain.Loan, error) {
	if loans == nil || len(loans) == 0 {
		return []domain.Loan{}, ErrorLoansNotFound
	}
	return loans, nil
}

/*GetLoan ...
get one specific loan with an id as parameter
if not exists loan then send an error*/
func GetLoan(i string) (domain.Loan, error) {
	id, err := validateID(i)
	if err != nil {
		return domain.Loan{}, ErrorInvalidID
	}

	loan, err := searchLoan(id)
	if err != nil {
		return domain.Loan{}, err
	}

	return loan, nil
}

/*UpdateLoan ..
update a specific loan with a loan as struct as parameter
if not exists or the loan isn't validated then send an error*/
func UpdateLoan(loan domain.Loan) (domain.Loan, error) {
	err := existsLoan(loan)
	if err == nil {
		return domain.Loan{}, ErrorLoanNotFound
	}

	err = validateLoan(loan)
	if err != nil {
		return domain.Loan{}, err
	}

	for i := 0; i < len(loans); i++ {
		if users[i].ID == loan.ID {
			loans[i].IDBook = loan.IDBook
			loans[i].IDUser = loan.IDUser
			loans[i].DueDate = loan.DueDate
			if loan.Info.Titulo != "" {
				loans[i].Info.Titulo = loan.Info.Titulo
			}
			if loan.Info.Titulo != "" {
				loans[i].Info.Autores = loan.Info.Autores
			}
			if loan.Info.Subtitulo != "" {
				loans[i].Info.Subtitulo = loan.Info.Subtitulo
			}
			if loan.Info.Titulo != "" {
				loans[i].Info.FechaPublicacion = loan.Info.FechaPublicacion
			}

			return loans[i], nil
		}
	}

	return domain.Loan{}, err
}

/*DeleteLoan ...
delete a specific loan with an id as parameter
if not exists loan then send an error*/
func DeleteLoan(i string) (domain.Loan, error) {
	id, err := validateID(i)
	if err != nil {
		return domain.Loan{}, ErrorInvalidID
	}

	loan, err := searchLoan(id)
	if err != nil {
		return domain.Loan{}, err
	}

	for i := 0; i < len(loans); i++ {
		if loans[i].ID == id {
			booksAmount[loans[i].IDBook]--
			loans[len(loans)-1], loans[i] = loans[i], loans[len(loans)-1]
			loans = loans[:len(loans)-1]
		}
	}

	return loan, nil
}

/************************************************************/
/************ CRUD LOANS : AUXILIARY FUNCTIONS **************/
/************************************************************/

// search a loan for ID
func searchLoan(id int) (domain.Loan, error) {
	for _, l := range loans {
		if l.ID == id {
			return l, nil
		}
	}
	return domain.Loan{}, ErrorLoanNotFound
}

// control if exists loan in array loans
func existsLoan(loan domain.Loan) error {
	for _, l := range loans {
		if l.ID == loan.ID {
			return ErrorLoanExists
		}
	}
	return nil
}

// control if a loaned book hasn't more copies
// or if exists in map amountBooks
func availability(id string) (int, error) {
	loans, exist := booksAmount[id]

	if exist && loans > amountBooks {
		return -1, ErrorNoAvailability
	}
	if !exist {
		return 0, nil
	}

	return loans, nil
}

// control if the loan expired of a specific user
func expiredLoans(idUser int) error {
	timeNow := time.Now()
	for _, l := range loans {
		if l.IDUser == idUser {
			dueDate, err := time.Parse("2006-01-02", l.DueDate)
			if err != nil {
				return ErrorInvalidFormatDate
			}
			if timeNow.After(dueDate) {
				return ErrorExpiredLoans
			}
		}
	}
	return nil
}

// control if the loan hasn't empty fields
func validEmptyFields(loan domain.Loan, c chan error) {
	err := false
	if !loan.IDValid() {
		err = true
		c <- ErrorInvalidID
	}
	if !loan.HasIDBook() {
		err = true
		c <- ErrorNoIDBook
	}
	if !loan.HasIDUser() {
		err = true
		c <- ErrorNoIDUser
	}
	if !loan.HasDueDate() {
		err = true
		c <- ErrorNoDueDate
	}
	if !err {
		c <- nil
	}
}

// control if the book of the loan is valid and it's availabilited
func validBook(l string, c chan error) {
	_, err := searchBook(l)
	if err != nil {
		c <- err
	}
	_, err = availability(l)
	if err != nil {
		c <- err
	}
	if err == nil {
		c <- nil
	}
}

// control if the user of the loan is valid
func validUser(l int, c chan error) {
	_, err := searchUser(l)
	c <- err
}

// control if the due date of the loan is valid
func validDueDate(dd string, c chan error) {
	timeNow := time.Now()
	dueDate, err := time.Parse("2006-01-02", dd)
	if err != nil {
		c <- ErrorInvalidFormatDate
	}
	if timeNow.After(dueDate) {
		err = ErrorInvalidDueDate
		c <- ErrorInvalidDueDate
	}
	if err == nil {
		c <- nil
	}
}

// control if the loans expired
func validExpiredLoans(idUser int, c chan error) {
	err := expiredLoans(idUser)
	c <- err
}

// combineted all funtions of valid loan with gorutines and channels
func validateLoan(loan domain.Loan) error {

	c := make(chan error)

	go validEmptyFields(loan, c)
	go validBook(loan.IDBook, c)
	go validUser(loan.IDUser, c)
	go validDueDate(loan.DueDate, c)
	go validExpiredLoans(loan.IDUser, c)

	for i := 0; i < 5; i++ {
		value := <-c
		if value != nil {
			return value
		}
	}

	return nil
}

/*************************************************************
************************** CRUD USERS ************************
*************************************************************/

/*CreateUser ...
create an user through of array users, and
control if is valid and if exists the user*/
func CreateUser(user domain.User) (domain.User, error) {
	err := validateUser(user)
	if err != nil {
		return domain.User{}, err
	}

	err = existsUser(user)
	if err != nil {
		return domain.User{}, err
	}

	users = append(users, user)
	return user, nil
}

/*GetAllUsers ...
get all users storage in array users*/
func GetAllUsers() ([]domain.User, error) {
	if users == nil || len(users) == 0 {
		return users, ErrorUsersNotFound
	}
	return users, nil
}

/*GetUser ...
get an specific user with an id as parameter
if not exists the user then send an error*/
func GetUser(i string) (domain.User, error) {
	id, err := validateID(i)
	if err != nil {
		return domain.User{}, ErrorInvalidID
	}

	user, err := searchUser(id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

/*UpdateUser ...
update the name and surname of the specific user
if not exists the user then send an error*/
func UpdateUser(user domain.User) (domain.User, error) {
	err := validateUser(user)
	if err != nil {
		return domain.User{}, err
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

	return domain.User{}, ErrorUserNotFound
}

/*DeleteUser ...
delete the specific user with an id as parameter
if not exists the user then send an error*/
func DeleteUser(i string) (domain.User, error) {
	id, err := validateID(i)
	if err != nil {
		return domain.User{}, ErrorInvalidID
	}

	user, err := searchUser(id)
	if err != nil {
		return domain.User{}, err
	}

	for i := 0; i < len(users); i++ {
		if users[i].ID == user.ID {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	return user, nil
}

/***********************************************************/
/************ CRUD USER : AUXILIARY FUNCTIONS **************/
/***********************************************************/

// control that user fields are valids
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

// control if exists the user
func existsUser(user domain.User) error {
	for _, u := range users {
		if u.ID == user.ID {
			return ErrorUserExists
		}
	}
	return nil
}

// control if id is valid
func validateID(id string) (int, error) {
	num, err := strconv.Atoi(id)
	if err != nil {
		return -1, err
	}
	return num, nil
}

// search the user for id
func searchUser(id int) (domain.User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return domain.User{}, ErrorUserNotFound
}

/*********************************************************
************ CONSUME GOOGLE BOOKS API *******************
********************************************************/

/*GetBook ...
consume the API of Google Books with the specific URI
to consult about the title and the author, and I bring
the complete title, subtitle, author/s and date of publication */
func GetBook(book domain.Book) ([]domain.Items, error) {
	err := validateBook(book)
	if err != nil {
		return []domain.Items{}, err
	}

	url := based + strings.Replace(book.Titulo, " ", "+", -1) + "+inauthor:" + strings.Replace(book.Autor, " ", "+", -1)
	url += "&fields=items(id,volumeInfo(title,subtitle,authors,publishedDate))"

	apiBook, err := apiResponse(url)
	if err != nil {
		return []domain.Items{}, err
	}

	var resultBooks []domain.Items
	for _, b := range apiBook.Items {
		resultBooks = append(resultBooks, b)
	}

	return resultBooks, nil
}

// consume the API of Google Books through of
// the params ID volume of the book
func searchBook(id string) (domain.Information, error) {
	url := based + "id=" + id

	apiBook, err := apiResponse(url)
	if err != nil {
		return domain.Information{}, err
	}

	if len(apiBook.Items) == 0 {
		return domain.Information{}, ErrorBookNotFound
	}

	return apiBook.Items[0].Info, nil
}

// generic funtions to consume URI and traslate the response
// to a struct to work the data in the API library gopher
func apiResponse(url string) (domain.GoogleBooks, error) {
	responseExternalAPI, err1 := http.Get(url)
	jsonDataFromHTTP, err2 := ioutil.ReadAll(responseExternalAPI.Body)

	var apiBook domain.GoogleBooks
	err3 := json.Unmarshal([]byte(jsonDataFromHTTP), &apiBook)

	if err1 != nil || err2 != nil {
		return domain.GoogleBooks{}, ErrorRequestExternalAPI
	}
	if err3 != nil {
		return domain.GoogleBooks{}, ErrorBookNotFound
	}

	return apiBook, nil
}

// control that title and author in the request hasn't
// any special character
func validateBook(book domain.Book) error {
	if book.SpecialChar() {
		return ErrorSpecialCharInBooks
	}
	return nil
}
