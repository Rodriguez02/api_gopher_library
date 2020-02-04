package domain

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
