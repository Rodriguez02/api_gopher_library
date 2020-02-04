package services

import "api_gopher_library/domain"

func GetSaludo() domain.Saludo {
	var saludo domain.Saludo
	saludo.Mensaje = "Hello this is The Library Gopher"
	saludo.Integrantes[0] = "Calvacho Emiliano"
	saludo.Integrantes[1] = "Cantarelli Sofia"
	saludo.Integrantes[2] = "Rodriguez José"
	return saludo
}

/* code
.
.
.
*/
