package domain

import (
	"testing"
)

/*********************** LOAN **************************/
func TestIDValidTrue(t *testing.T){
	r := Loan {
		ID: 1,
	}
	a := r.IDValid()
	if a != true {
		t.Error("An error occurred while validating loan's id")
		return
	}
	
}

func TestIDValidFalse(t *testing.T){
	r := Loan {
		ID: -1,
	}
	a := r.IDValid()
	if a == true {
		t.Error("An error occurred while validating loan's id")
		return
	}
}

func TestHasIDBookTrue(t *testing.T){
	r := Loan {
		IDBook: "1a2",
	}
	a := r.HasIDBook()
	if a != true {
		t.Error("An error occurred while validating book's id")
		return
	}
}

func TestHasIDBookFalse(t *testing.T){
	r := Loan {
		IDBook: "",
	}
	a := r.HasIDBook()
	if a == true {
		t.Error("An error occurred while validating book's id")
		return
	}
}

func TestHasIDUserTrue(t *testing.T){
	r := Loan {
		IDUser: 1,
	}
	a := r.HasIDUser()
	if a != true {
		t.Error("An error occurred while validating user's id")
		return
	}
}

func TestHasIDUserFalse(t *testing.T){
	r := Loan {
		IDUser: -1,
	}
	a := r.HasIDUser()
	if a == true {
		t.Error("An error occurred while validating user's id")
		return
	}
}
func TestHasDueDateTrue(t *testing.T){
	r := Loan {
		DueDate: "20-12-2020",
	}
	a := r.HasDueDate()
	if a != true {
		t.Error("An error occurred while validating due's date")
		return
	}
}

func TestHasIDDueDateFalse(t *testing.T){
	r := Loan {
		DueDate: "",
	}
	a := r.HasDueDate()
	if a == true {
		t.Error("An error occurred while validating due's date")
		return
	}
}


/*********************** USER **************************/

func TestHasNameTrue(t *testing.T){
	r := User {
		Nombre: "Uno",
	}
	a := r.HasName()
	if a != true {
		t.Error("An error occurred while validating name")
		return
	}
}

func TestHasNameFalse(t *testing.T){
	r := User {
		Nombre: "",
	}
	a := r.HasName()
	if a == true {
		t.Error("An error occurred while validating name")
		return
	}
}

func TestHasSurnameTrue(t *testing.T){
	r := User {
		Apellido: "Uno",
	}
	a := r.HasSurname()
	if a != true {
		t.Error("An error occurred while validating surname")
		return
	}
}

func TestHasSurnameFalse(t *testing.T){
	r := User {
		Apellido: "",
	}
	a := r.HasSurname()
	if a == true {
		t.Error("An error occurred while validating surname")
		return
	}
}

func TestHasIDValidTrue(t *testing.T){
	r := User {
		ID: 1,
	}
	a := r.IDValid()
	if a != true {
		t.Error("An error occurred while validating user's id")
		return
	}
}

func TestHasIDValidFalse(t *testing.T){
	r := User {
		ID: -1,
	}
	a := r.IDValid()
	if a == true {
		t.Error("An error occurred while validating user's id")
		return
	}
}


/*********************** BOOK **************************/

func TestHasTitleTrue(t *testing.T){
	r := Book {
		Titulo: "Uno",
	}
	a := r.HasTitle()
	if a != true {
		t.Error("An error occurred while validating title")
		return
	}
}

func TestHasTitleFalse(t *testing.T){
	r := Book {
		Titulo: "",
	}
	a := r.HasTitle()
	if a == true {
		t.Error("An error occurred while validating title")
		return
	}
}

func TestHasAuthorTrue(t *testing.T){
	r := Book {
		Autor: "Uno",
	}
	a := r.HasAuthor()
	if a != true {
		t.Error("An error occurred while validating author")
		return
	}
}

func TestHasAuthorFalse(t *testing.T){
	r := Book {
		Autor: "",
	}
	a := r.HasAuthor()
	if a == true {
		t.Error("An error occurred while validating author")
		return
	}
}