package viewmodel

//Base struct
type Base struct{
	Title string
}

//NewBase definition
func NewBase() Base{
	return Base{
		Title: "Lemonad Stand ",
	}
}