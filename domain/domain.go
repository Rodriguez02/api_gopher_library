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

type Book struct {
	Nombre string `json:"nombre"`
	Autor  string `json:"autor"`
}

type GoogleBooks struct {
	Items []Items `json:"items"`
}

type Items struct {
	Info Information `json:"volumeInfo"`
}

type Information struct {
	Titulo           string   `json:"title"`
	Subtitulo        string   `json:"subtitle"`
	Autores          []string `json:"authors"`
	FechaPublicacion string   `json:"publishedDate"`
}
