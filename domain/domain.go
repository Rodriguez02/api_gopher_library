package domain

type Book struct {
	ID 		int    `json:"id"`
	Title	string `json:"title"`
	Amount  int	   `json:"amount"`
}

type Loan struct {
	ID		 int   `json:"id"`
	IDBook   int   `json:"idBook"`
	IDUser	 int   `json:"idUser"`
	DueDate	 int64 `json:"dueDate"`
}

func (l Loan) IDValid() bool{
	return l.ID > 0
}

func (l Loan) HasIDBook() bool{
	return l.IDBook > 0
}

func (l Loan) HasIDUser() bool{
	return l.IDUser > 0
}

func (l Loan) HasDueDate() bool{
	return l.DueDate > 0
}

type User struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func (u User) HasName() bool {
	return u.Nombre != ""
}

func (u User) HasSurname() bool {
	return u.Apellido != ""
}

func (u User) IDValid() bool {
	return u.ID > 0
}


