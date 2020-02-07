package services

import (
	"api_gopher_library/domain"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	//"encoding/json"
	//"io/ioutil"
	//"net/http"
	"time"
)

var (
	users []domain.User
	loans []domain.Loan
	books []domain.MyBook
    
	// errores libros
	ErrorBookNotFound 		= errors.New("book not found")

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
	ErrorNoIDBook 	        = errors.New("loan needs idbook")
	ErrorNoIDUser	        = errors.New("loan needs iduser")
	ErrorNoDueDate          = errors.New("loan needs due")
	ErrorLoanExists         = errors.New("this loan exists")
	ErrorLoansNotFound      = errors.New("there aren't loans")
	ErrorLoanNotFound       = errors.New("loan not found")
	ErrorNoAvailability     = errors.New("book without availability")
	ErrorExpiredBooksOfUser = errors.New("expired books of user")
	ErrorInvalidDueDate	    = errors.New("due date must be greater than current")
)

const (
	based = "https://www.googleapis.com/books/v1/volumes?key=AIzaSyDVnZCPWXdzNcWiipQ7ng5E-eLRg3xu7MY&fields=items(volumeInfo(title,subtitle,authors,publishedDate))&q="
)

/********************************************************/
/******************** CRUD USERS ************************/
/********************************************************/

func init(){
	
	book1 := domain.MyBook{
							ID : 1,
							Title : "Libro uno",
							Amount : 5,
						}
	book2 := domain.MyBook{
							ID : 2,
							Title : "Libro dos",
							Amount : 2,
						}
	books = append(books, book1)
	books = append(books, book2)
}

/***************************************************
*********************CRUD LOANS*********************
***************************************************/

func CreateLoan(loan domain.Loan) (domain.Loan, error) {
	err := existsLoan(loan)
	if err != nil{
		return domain.Loan{}, ErrorLoanExists
	}

	err = validateLoan(loan)
	if err != nil {
		return loan, err
	}
	loans = append(loans, loan)
	return loan, nil
}

func GetAllLoans() ([]domain.Loan, error){
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
	if err != nil{
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
			loans[i].IDBook  = loan.IDBook
			loans[i].IDUser  = loan.IDUser
			loans[i].DueDate = loan.DueDate
			return loans[i], nil
		}
	}

	return loan, ErrorLoanNotFound
}

func DeleteLoan(i string) (domain.Loan, error) {
	id, err := validateID(i)
	if err != nil {
		return domain.Loan{}, ErrorInvalidID
	}

	loan, err := searchLoan(id)
	if err != nil{
		return domain.Loan{}, err
	}

	for i := 0; i < len(loans); i++ {
		if loans[i].ID == id {
			loans[len(loans)-1], loans[i] = loans[i], loans[len(loans)-1]
			loans = loans[:len(loans)-1]
		}
	}
	return loan, nil
}

func searchLoan(id int) (domain.Loan, error){	
	for _,l := range loans {
		if l.ID == id{
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

func availability(book domain.MyBook) (int, error){
	booksAvailables := book.Amount - amountRentedBooks(book.ID)
	if ( booksAvailables <= 0){
		return 0, ErrorNoAvailability
	}
	return booksAvailables, nil
}

func expiredBooksOfUser(idUser int) bool{
	timeNow := time.Now().UnixNano() / int64(time.Millisecond)
	for _, l := range loans{
		if l.IDUser == idUser {
			if timeNow > l.DueDate {
				return true
			}
		}
	}
	return false
}

func amountRentedBooks(id int) int{
	rentedBooks := 0
	for _,l := range loans {
		if l.IDBook == id{
			rentedBooks++
		}
	}
	return rentedBooks
}


func validateLoan(loan domain.Loan) error{
	if !loan.IDValid(){
		return ErrorInvalidID
	}
	if !loan.HasIDBook() {
		return ErrorNoIDBook
	}
	if !loan.HasIDUser(){
		return ErrorNoIDUser
	}
	if !loan.HasDueDate(){
		return ErrorNoDueDate
	}

	book, err := searchBook(loan.IDBook) 
	if err != nil{
		return ErrorBookNotFound
	}
	_, err = searchUser(loan.IDUser)
	if err != nil {
		return ErrorUserNotFound
	}

	timeNow := time.Now().UnixNano() / int64(time.Millisecond)
	if loan.DueDate <= timeNow {
		return ErrorInvalidDueDate
	}

	_, err = availability(book)
	if err != nil{
		return ErrorNoAvailability
	}

	if expiredBooksOfUser(loan.IDUser) != false {
		return ErrorExpiredBooksOfUser
	}

	return nil
}

func searchBook(id int) (domain.MyBook, error){	
	for _,b := range books {
		if b.ID == id{
			return b, nil
		}
	}
	return domain.MyBook{}, ErrorBookNotFound
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

func GetBook(book domain.Book) ([]domain.Information, error) {
	err := validateBook(book)
	if err != nil {
		return []domain.Information{}, err
	}

	url := based + strings.Replace(book.Titulo, " ", "+", -1) + "+inauthor:" + strings.Replace(book.Autor, " ", "+", -1)
	fmt.Println(url)
	responseExternalAPI, err1 := http.Get(url)
	jsonDataFromHttp, err2 := ioutil.ReadAll(responseExternalAPI.Body)

	var api_book domain.GoogleBooks
	err3 := json.Unmarshal([]byte(jsonDataFromHttp), &api_book)

	if err1 != nil || err2 != nil || err3 != nil {
		return []domain.Information{}, ErrorRequestExternalAPI
	}

	var result_books []domain.Information

	for _, b := range api_book.Items {
		result_books = append(result_books, b.Info)
	}

	return result_books, nil
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

func searchUser(id int) (domain.User, error){	
	for _,u := range users {
		if u.ID == id{
			return u, nil
		}
	}
	return domain.User{}, ErrorUserNotFound
}