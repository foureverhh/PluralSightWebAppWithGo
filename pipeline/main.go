package main

import (
	"html/template"
	"os"
)

const tax = 6.75 / 100

//Product struct
type Product struct {
	Name  string
	Price float32
}
/* 
//PriceWithTax method
func (p Product) PriceWithTax() float32 {
	return p.Price * (1 + tax)
}
*/

const templateString = `
{{- "Item Information" }}
Name: {{ .Name}}
Price: {{ printf "$%.2f" .Price}}
Price with Tax: {{ calctax .Price | printf "$%.2f"}}
`
//Price with Tax: {{ .PriceWithTax | printf "$%.2f"}}

func main() {

	p := Product{
		Name:  "Lemonade",
		Price: 2.16,
	}

	fm := template.FuncMap{}

	fm["calctax"] = func (price float32) float32 {
		return price * (1 + tax)
	}

	t := template.Must(template.New("").Funcs(fm).Parse(templateString))

	//This is for PriceWithTax
	//t:=template.Must(template.New("").Parse(templateString))
	t.Execute(os.Stdout, p)
}
