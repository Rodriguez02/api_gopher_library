package services

import (
	"api_gopher_library/domain"
	//"encoding/json"
	"errors"
	//"io/ioutil"
	//"net/http"
	"strconv"
	"time"
)

var (
	users []domain.User
	loans []domain.Loan
	books []domain.Book
    
	// errores libros
	ErrorBookNotFound 		= errors.New("book not found")

	// errores usuarios
	ErrorNoName             = errors.New("user needs name")
	ErrorNoSurname          = errors.New("user needs surname")
	ErrorInvalidID          = errors.New("id isn't valid")
	ErrorUserExists         = errors.New("this user exists")
	ErrorUsersNotFound      = errors.New("there aren't users")
	ErrorUserNotFound       = errors.New("user not found")
	 
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


func init(){
	
	book1 := domain.Book{
							ID : 1,
							Title : "Libro uno",
							Amount : 5,
						}
	book2 := domain.Book{
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

func availability(book domain.Book) (int, error){
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

func searchBook(id int) (domain.Book, error){	
	for _,b := range books {
		if b.ID == id{
			return b, nil
		}
	}
	return domain.Book{}, ErrorBookNotFound
}

/***************************************************
*********************CRUD USERS*********************
***************************************************/

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