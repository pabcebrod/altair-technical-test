package data

type Person struct {
	Id       string
	Nombre   string
	Apellido string
	Equipo   string
	Edad     int
	Palabra  string
}

//Enum type for teams
const (
		RED string = "rojo"
		BLUE string = "azul"
	)