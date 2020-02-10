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
	users       []domain.User
	loans       []domain.Loan
	booksAmount = make(map[string]int)

	// errores libros
	ErrorBookNotFound = errors.New("book not found")


	// errores usuarios
	ErrorNoName             = errors.New("user needs name")
	ErrorNoSurname          = errors.New("user needs surname")
	ErrorInvalidID          = errors.New("id isn't valid")
	ErrorUserExists         = errors.New("this user exists")
	ErrorUsersNotFound      = errors.New("there aren't users")
	ErrorUserNotFound       = errors.New("user not found")
	ErrorRequestExternalAPI = errors.New("error in request to external API")
	ErrorSpecialCharInBooks = errors.New("there're almost a special character in title or author")

	// errores prestamos
	ErrorNoIDBook           = errors.New("loan needs idbook")
	ErrorNoIDUser           = errors.New("loan needs iduser")
	ErrorNoDueDate          = errors.New("loan needs due")
	ErrorLoanExists         = errors.New("this loan exists")
	ErrorLoansNotFound      = errors.New("there aren't loans")
	ErrorLoanNotFound       = errors.New("loan not found")
	ErrorNoAvailability     = errors.New("book without availability")
	ErrorExpiredBooksOfUser = errors.New("expired books of user")
	ErrorInvalidDueDate     = errors.New("due date must be greater than current")
	ErrorInvalidFormatDate  = errors.New("invalid due date. example: 2 Jan 2006 = 2006-01-02")
)

const (
	based       = "https://www.googleapis.com/books/v1/volumes?key=AIzaSyDVnZCPWXdzNcWiipQ7ng5E-eLRg3xu7MY&q="
	amountBooks = 2
)

/***************************************************
*********************CRUD LOANS*********************
***************************************************/

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

func GetAllLoans() ([]domain.Loan, error) {
	if loans == nil {
		return loans, ErrorLoansNotFound
	}
	return loans, nil
}

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
			loans[i].Info.Titulo = loan.Info.Titulo
			loans[i].Info.Autores = loan.Info.Autores
			loans[i].Info.Subtitulo = loan.Info.Subtitulo
			return loans[i], nil
		}
	}

	return domain.Loan{}, ErrorLoanNotFound
}

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
			loans[len(loans)-1], loans[i] = loans[i], loans[len(loans)-1]
			loans = loans[:len(loans)-1]
			booksAmount[loans[i].IDBook]--
		}
	}
	return loan, nil
}

func searchLoan(id int) (domain.Loan, error) {
	for _, l := range loans {
		if l.ID == id {
			return l, nil
		}
	}
	return domain.Loan{}, ErrorLoanNotFound
}

func existsLoan(loan domain.Loan) error {
	for _, l := range loans {
		if l.ID == loan.ID {
			return ErrorLoanExists
		}
	}
	return nil
}

/************** CRUD LOANS : FUNCIONES AUXILIARES **************/

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

func expiredLoans(idUser int) error{
	timeNow := time.Now()
	for _, l := range loans{
		if l.IDUser == idUser {
			dueDate, err := time.Parse("2006-01-02", l.DueDate)
			if err != nil {
				return ErrorInvalidFormatDate
			}
			if timeNow.After(dueDate) {
				return ErrorInvalidDueDate
			}
		}
	}
	return nil
}

func validEmptyFields(loan domain.Loan, c chan error){
	c <- nil
	if !loan.IDValid(){
		c <- ErrorInvalidID
	}
	if !loan.HasIDBook() {
		c <- ErrorNoIDBook
	}
	if !loan.HasIDUser(){
		c <- ErrorNoIDUser
	}
	if !loan.HasDueDate(){
		c <- ErrorNoDueDate
	}
}

func validBook(l string, c chan error){
	_, err := searchBook(l) 
	c <- nil
	if err != nil{
		c <- err
	}
	_, err = availability(l)
	if err != nil{
		c <- err
	}
}

func validUser(l int, c chan error){
	_, err := searchUser(l)
	c <- err
}

func validDueDate(dd string, c chan error){
	c <- nil
	timeNow := time.Now()
	dueDate, err := time.Parse("2006-01-02", dd)
	if err != nil {
		c <- ErrorInvalidFormatDate
	}
	if timeNow.After(dueDate) {
		c <- ErrorInvalidDueDate
	}
}

func validExpiredLoans(idUser int, c chan error){
	err := expiredLoans(idUser)
	c <- err
}


func validateLoan(loan domain.Loan) error {

	c := make(chan error)

	go validEmptyFields(loan, c)
	go validBook(loan.IDBook, c)
	go validUser(loan.IDUser, c)
	go validDueDate(loan.DueDate, c)
	go validExpiredLoans(loan.IDUser, c)

	for i := 0 ; i < 5; i++ {
		value := <- c
		if value != nil {
			return value
		}
	}

	return nil
}

func searchBook(id string) (domain.Information, error) {
	url := based + "id=" + id

	api_book, err := apiResponse(url)
	if err != nil {
		return domain.Information{}, err
	}

	if len(api_book.Items) == 0 {
		return domain.Information{}, ErrorBookNotFound
	}

	return api_book.Items[0].Info, nil
}

/***************************************************
*********************CRUD USERS*********************
***************************************************/

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
	if err != nil {
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

func GetBook(book domain.Book) ([]domain.Items, error) {
	err := validateBook(book)
	if err != nil {
		return []domain.Items{}, err
	}

	url := based + strings.Replace(book.Titulo, " ", "+", -1) + "+inauthor:" + strings.Replace(book.Autor, " ", "+", -1)
	url += "&fields=items(id,volumeInfo(title,subtitle,authors,publishedDate))"

	api_book, err := apiResponse(url)
	if err != nil {
		return []domain.Items{}, err
	}

	var result_books []domain.Items
	for _, b := range api_book.Items {
		result_books = append(result_books, b)
	}

	return result_books, nil
}

func apiResponse(url string) (domain.GoogleBooks, error) {
	responseExternalAPI, err1 := http.Get(url)
	jsonDataFromHttp, err2 := ioutil.ReadAll(responseExternalAPI.Body)

	var api_book domain.GoogleBooks
	err3 := json.Unmarshal([]byte(jsonDataFromHttp), &api_book)

	if err1 != nil || err2 != nil || err3 != nil {
		return domain.GoogleBooks{}, ErrorRequestExternalAPI
	}

	return api_book, nil
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

func validateBook(book domain.Book) error {
	if book.SpecialChar() {
		return ErrorSpecialCharInBooks
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

func searchUser(id int) (domain.User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return domain.User{}, ErrorUserNotFound
}
