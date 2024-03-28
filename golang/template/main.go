package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

// Custom function to be used in the template
var funcs = template.FuncMap{"upper": upper}

func upper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	// Declare a new template and parse the template file
	namesTmpl := template.Must(template.New("test.tmpl").Funcs(funcs).ParseFiles("test.tmpl"))
	data := []string{"one", "two", "three"}

	// Execute the template and pass it `data` as it's data
	if err := namesTmpl.Execute(os.Stdout, data); err != nil {
		log.Fatalln(err)
	}
}
